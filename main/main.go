package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/stevenwal/mock-indexer/contracts"
)

var (
	key                          = ""
	l1Rpc                        = "ws://anvil-l1:8545"
	l2Rpc                        = "ws://anvil-l2:8545"
	arbRpc                       = "ws://anvil-arb:8545"
	set                          = false
	host                         = ":1010"
	l1ERC20BridgeAddress         = common.Address{}
	l2BridgeAddress              = common.Address{}
	socketControllerAddress      = common.Address{}
	socketVaultAddress           = common.Address{}
	L1ERC20BridgeInterface, _    = abi.JSON(strings.NewReader(contracts.MockL1ERC20BridgeABI))
	L2ERC20BridgeInterface, _    = abi.JSON(strings.NewReader(contracts.MockL2ERC20BridgeABI))
	SocketVaultInterface, _      = abi.JSON(strings.NewReader(contracts.MockL1SocketVaultABI))
	SocketControllerInterface, _ = abi.JSON(strings.NewReader(contracts.MockL2SocketControllerABI))
	DepositInitiated             = L1ERC20BridgeInterface.Events["MockDepositInitiated"].ID
	WithdrawalInitiated          = L2ERC20BridgeInterface.Events["MockWithdrawalInitiated"].ID
	MessageOutbound              = SocketVaultInterface.Events["MessageOutbound"].ID
)

// Timeout when connecting to the RPC
const clientTimeout = 5 * time.Second

// ContractBackend defines the methods needed to work with contracts on a read-write basis
type ContractBackend interface {
	bind.ContractBackend

	// BlockNumber returns the most recent block number
	BlockNumber(ctx context.Context) (uint64, error)
	ChainID(ctx context.Context) (*big.Int, error)
}

// Create an ETH RPC client
func CreateClient(ctx context.Context, rpc string) ContractBackend {
	for {
		log.Info("Connecting to " + rpc + "...")
		if client, err := ethclient.DialContext(ctx, rpc); err != nil {
			time.Sleep(clientTimeout)
		} else {
			log.Info("Connected.")
			return client
		}
	}
}

func CreateDBConn(ctx context.Context) driver.Conn {
	opt := &clickhouse.Options{
		Addr: []string{"clickhouse:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "",
			Password: "",
		},
		DialTimeout: time.Minute,
	}

	for {
		log.Info("Connecting to clickhouse...")
		if conn, err := clickhouse.Open(opt); err != nil {
			time.Sleep(clientTimeout)
		} else {
			log.Info("Connected.")
			return conn
		}
	}
}

// Create a subscription to the client to listen for on-chain events
func StartSubscription(
	ctx context.Context,
	client ContractBackend,
	query ethereum.FilterQuery,
	sink chan<- types.Log,
) {
	// Start the background filtering
	logs := make(chan types.Log, 128)

	// Create the subscription
	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return
	}

	for {
		select {
		case l := <-logs:
			select {
			case sink <- l:
			case err := <-sub.Err():
				log.Info(err)
			}
		case err := <-sub.Err():
			log.Info(err)
		case <-ctx.Done():
			return
		}
	}
}

func cacheLocation() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	return filepath.Join(basepath, "addresses.json")
}

func readCache() {
	jsonFile, err := os.Open(cacheLocation())

	if err != nil {
		log.Error("cache does not exist")
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var deployments ContractAddresses

	err = json.Unmarshal(byteValue, &deployments)
	if err != nil {
		log.Error("fail to parse addresses")
	}

	l1ERC20BridgeAddress = deployments.L1StandardBridge
	l2BridgeAddress = deployments.L2StandardBridge
	socketVaultAddress = deployments.ArbitrumUSDCVault
	socketControllerAddress = deployments.SocketController
}

func main() {
	var (
		delay int64
	)

	flag.Int64Var(&delay, "delay", 0, "Delay duration in seconds")
	flag.Parse()

	log.Info(fmt.Sprintf("Delay set at %d seconds", delay))

	cache := cacheLocation()
	_, err := os.Open(cache)
	if err == nil {
		readCache()
		set = true
	}

	ctx := context.Background()
	key = os.Getenv("OWNER_KEY")
	router := httprouter.New()
	router.POST("/addresses", func(_ http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		var deployments ContractAddresses

		b, err := io.ReadAll(req.Body)
		if err != nil {
			log.Error("fail to parse addresses")
			return
		}

		err = json.Unmarshal(b, &deployments)
		if err != nil {
			log.Error("fail to parse addresses")
			return
		}

		if deployments.L1StandardBridge.Hex() == (common.Address{}).Hex() {
			log.Error("empty account")
			return
		}
		if deployments.L2StandardBridge.Hex() == (common.Address{}).Hex() {
			log.Error("empty account")
			return
		}

		l1ERC20BridgeAddress = deployments.L1StandardBridge
		l2BridgeAddress = deployments.L2StandardBridge
		socketVaultAddress = deployments.ArbitrumUSDCVault
		socketControllerAddress = deployments.SocketController

		file, _ := json.MarshalIndent(deployments, "", "  ")
		err = os.WriteFile(cacheLocation(), file, 0644)
		if err != nil {
			panic(err)
		}

		set = true
	})

	server := &http.Server{
		Addr:    host,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Error(err)
		}
	}()

	for !set {
		time.Sleep(2 * time.Second)
		log.Info("Waiting for deployment address")
	}

	l1Client := CreateClient(ctx, l1Rpc)
	l2Client := CreateClient(ctx, l2Rpc)
	arbClient := CreateClient(ctx, arbRpc)

	sink := make(chan types.Log, 1000)

	conn := CreateDBConn(ctx)

	l1ChainId, _ := l1Client.ChainID(ctx)
	l2ChainId, _ := l2Client.ChainID(ctx)
	arbChainId, _ := arbClient.ChainID(ctx)

	ownerl1 := wallet(key, l1ChainId)
	ownerl2 := wallet(key, l2ChainId)
	ownerArb := wallet(key, arbChainId)
	l1Bridge, _ := contracts.NewMockL1ERC20Bridge(l1ERC20BridgeAddress, l1Client)
	l2Bridge, _ := contracts.NewMockL2ERC20Bridge(l2BridgeAddress, l2Client)
	socketController, _ := contracts.NewMockL2SocketController(socketControllerAddress, arbClient)
	socketVault, _ := contracts.NewMockL1SocketVault(socketVaultAddress, arbClient)

	l1Topics, _ := abi.MakeTopics([]interface{}{DepositInitiated})
	l1Filter := ethereum.FilterQuery{
		Addresses: []common.Address{l1ERC20BridgeAddress},
		Topics:    l1Topics,
	}

	l2Topics, _ := abi.MakeTopics([]interface{}{WithdrawalInitiated})
	l2Filter := ethereum.FilterQuery{
		Addresses: []common.Address{l2BridgeAddress, socketControllerAddress},
		Topics:    l2Topics,
	}

	arbTopics, _ := abi.MakeTopics([]interface{}{MessageOutbound})
	arbFilter := ethereum.FilterQuery{
		Addresses: []common.Address{socketVaultAddress},
		Topics:    arbTopics,
	}

	events, err := l1Client.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{l1ERC20BridgeAddress},
		Topics:    l1Topics,
		FromBlock: new(big.Int).SetUint64(0),
	})
	for _, event := range events {
		sink <- event
	}
	if err != nil {
		log.Fatal(err)
	}

	events, err = l2Client.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{l2BridgeAddress},
		Topics:    l2Topics,
		FromBlock: new(big.Int).SetUint64(0),
	})
	for _, event := range events {
		sink <- event
	}
	if err != nil {
		log.Fatal(err)
	}

	events, err = arbClient.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{socketVaultAddress},
		Topics:    arbTopics,
		FromBlock: new(big.Int).SetUint64(0),
	})
	for _, event := range events {
		sink <- event
	}
	if err != nil {
		log.Fatal(err)
	}

	go StartSubscription(ctx, l1Client, l1Filter, sink)
	go StartSubscription(ctx, l2Client, l2Filter, sink)
	go StartSubscription(ctx, arbClient, arbFilter, sink)

	log.Info("Indexing...")

	for {
		select {
		case event := <-sink:
			switch event.Address {
			case socketControllerAddress:
				switch event.Topics[0] {
				case MessageOutbound:
					// execute on l2
					withdraw, _ := socketController.MockL2SocketControllerFilterer.ParseSocketMessage(event)
					log.WithFields(log.Fields{
						"msg": withdraw.MsgId,
					}).Info("Withdraw initiated.")

					time.Sleep(time.Duration(delay) * time.Second)

					// Finalize deposit
					_, err := socketVault.FinalizeERC20Withdrawal(ownerArb.signer, withdraw.To, withdraw.Amount, withdraw.MsgId)
					if err != nil {
						log.WithField("err", err).Error("finalize failed")
					}
				}
			case socketVaultAddress:
				switch event.Topics[0] {
				case MessageOutbound:
					// execute on l2
					deposit, _ := socketVault.MockL1SocketVaultFilterer.ParseSocketMessage(event)
					log.WithFields(log.Fields{
						"msg": deposit.MsgId,
					}).Info("Deposit initiated.")

					time.Sleep(time.Duration(delay) * time.Second)

					// Finalize deposit
					_, err := socketController.FinalizeDeposit(ownerl2.signer, deposit.To, deposit.Amount, deposit.MsgId)
					if err != nil {
						log.WithField("err", err).Error("deposit failed")
					}
				}
			case l2BridgeAddress:
				switch event.Topics[0] {
				case WithdrawalInitiated:
					// execute on l1
					withdraw, _ := l2Bridge.MockL2ERC20BridgeFilterer.ParseMockWithdrawalInitiated(event)
					log.WithFields(log.Fields{
						"account": withdraw.From.Hex(),
						"to":      withdraw.To.Hex(),
						"amount":  withdraw.Amount.String(),
					}).Info("Withdrawal initiated.")

					time.Sleep(time.Duration(delay) * time.Second)

					tx, err := l1Bridge.FinalizeERC20Withdrawal(
						ownerl1.signer,
						withdraw.L1Token,
						withdraw.L2Token,
						withdraw.From,
						withdraw.To,
						withdraw.Amount,
						withdraw.Data,
						withdraw.WithdrawHash,
					)
					if err != nil {
						log.Error(err)
					}
					err = conn.Exec(ctx,
						"INSERT INTO processed_withdraw_history VALUES ($1, $2, $3, $4, $5)",
						event.TxHash.Hex(),
						tx.Hash().Hex(),
						int64(event.BlockNumber),
						l1ChainId,
						time.Now().Unix(),
					)
					if err != nil {
						log.Error(err)
					}

				default:
					continue
				}
			case l1ERC20BridgeAddress:
				switch event.Topics[0] {
				case DepositInitiated:
					// execute on l2
					deposit, _ := l1Bridge.MockL1ERC20BridgeFilterer.ParseMockDepositInitiated(event)
					log.WithFields(log.Fields{
						"account": deposit.From.Hex(),
						"to":      deposit.To.Hex(),
						"amount":  deposit.Amount.String(),
						"data":    common.BytesToHash(deposit.Data),
					}).Info("Deposit initiated.")

					time.Sleep(time.Duration(delay) * time.Second)

					// Finalize deposit
					_, err := l2Bridge.FinalizeDeposit(ownerl2.signer, deposit.L1Token, deposit.L2Token, deposit.From, deposit.To, deposit.Amount, deposit.Data, deposit.Message)
					if err != nil {
						log.WithField("err", err).Error("finalize failed")
					}
				default:
					continue
				}
			}
		case <-ctx.Done():
			// Shutdown server
			if err := server.Shutdown(ctx); err != nil {
				log.Fatal(err)
			}
			return
		}
	}
}

type Wallet struct {
	address common.Address
	signer  *bind.TransactOpts
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func wallet(key string, chainId *big.Int) Wallet {
	privateKey, err := crypto.HexToECDSA(key)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	check(err)

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	check(err)

	return Wallet{
		address: address,
		signer:  signer,
	}
}

type ContractAddresses struct {
	L1StandardBridge  common.Address `json:"l1StandardBridge"`
	L2StandardBridge  common.Address `json:"l2StandardBridge"`
	SocketController  common.Address `json:"socketController"`
	ArbitrumUSDCVault common.Address `json:"arbitrumUSDCVault"`
}
