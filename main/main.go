package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
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
	log "github.com/sirupsen/logrus"
	"github.com/stevenwal/mock-indexer/contracts"
)

var (
	cache                     = "addresses.json"
	key                       = ""
	l1Rpc                     = "ws://anvil-l1:8545"
	l2Rpc                     = "ws://anvil-l2:8545"
	accountsAddress           = common.Address{}
	l1ERC20BridgeAddress      = common.Address{}
	l2BridgeAddress           = common.Address{}
	l1UsdcAddress             = common.Address{}
	l2UsdcAddress             = common.Address{}
	AccountsInterface, _      = abi.JSON(strings.NewReader(contracts.AccountsABI))
	L1ERC20BridgeInterface, _ = abi.JSON(strings.NewReader(contracts.IL1ERC20BridgeABI))
	ERC20DepositInitiated     = L1ERC20BridgeInterface.Events["ERC20DepositInitiated"].ID
	ERC20WithdrawalFinalized  = L1ERC20BridgeInterface.Events["ERC20WithdrawalFinalized"].ID
	Deposit                   = AccountsInterface.Events["Deposit"].ID
	Withdraw                  = AccountsInterface.Events["Withdraw"].ID
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

func init() {
	key = os.Getenv("OWNER_KEY")

	file, err := os.Open(cache)

	if err != nil {
		log.Fatal("cache does not exist")
	}

	byteValue, _ := io.ReadAll(file)

	var deployments ContractAddresses

	err = json.Unmarshal(byteValue, &deployments)
	if err != nil {
		log.Fatal("fail to parse addresses")
	}

	accountsAddress = deployments.Account
	l1ERC20BridgeAddress = deployments.L1StandardBridge
	l2BridgeAddress = deployments.L2StandardBridge
	l1UsdcAddress = deployments.USDC
	l2UsdcAddress = deployments.L2USDC
}

func main() {
	ctx := context.Background()

	var (
		delay int64
	)

	flag.Int64Var(&delay, "delay", 0, "Delay duration in seconds")
	flag.Parse()

	log.Info(fmt.Sprintf("Delay set at %d seconds", delay))

	sink := make(chan types.Log)

	l1Client := CreateClient(ctx, l1Rpc)
	l2Client := CreateClient(ctx, l2Rpc)
	conn := CreateDBConn(ctx)

	l1ChainId, _ := l1Client.ChainID(ctx)
	l2ChainId, _ := l2Client.ChainID(ctx)

	ownerl1 := wallet(key, l1ChainId)
	ownerl2 := wallet(key, l2ChainId)
	l1Bridge, _ := contracts.NewIL1ERC20Bridge(l1ERC20BridgeAddress, l1Client)
	l2Bridge, _ := contracts.NewMockL2ERC20Bridge(l2BridgeAddress, l2Client)

	accounts, _ := contracts.NewAccounts(accountsAddress, l2Client)
	l2Usdc, _ := contracts.NewTestERC20(l2UsdcAddress, l2Client)

	l1Topics, _ := abi.MakeTopics([]interface{}{ERC20DepositInitiated, ERC20WithdrawalFinalized})
	l1Filter := ethereum.FilterQuery{
		Addresses: []common.Address{l1ERC20BridgeAddress},
		Topics:    l1Topics,
	}

	l2Topics, _ := abi.MakeTopics([]interface{}{Deposit, Withdraw})
	l2Filter := ethereum.FilterQuery{
		Addresses: []common.Address{accountsAddress},
		Topics:    l2Topics,
	}

	go StartSubscription(ctx, l1Client, l1Filter, sink)
	go StartSubscription(ctx, l2Client, l2Filter, sink)

	log.Info("Indexing...")

	for {
		select {
		case event := <-sink:
			switch event.Address {
			case accountsAddress:
				switch event.Topics[0] {
				case Withdraw:
					// execute on l1
					withdraw, _ := accounts.AccountsFilterer.ParseWithdraw(event)
					log.WithFields(log.Fields{
						"account": withdraw.Account.Hex(),
						"to":      withdraw.To.Hex(),
						"amount":  withdraw.Amount.String(),
					}).Info("Withdrawal initiated.")

					time.Sleep(time.Duration(delay) * time.Second)

					tx, err := l1Bridge.FinalizeERC20Withdrawal(ownerl1.signer, l1UsdcAddress, withdraw.Collateral, withdraw.Account, withdraw.Account, withdraw.Amount, []byte{})
					if err != nil {
						log.Error(err)
					}
					err = conn.Exec(ctx,
						"INSERT INTO processed_withdraw_history VALUES ($1, $2, $3, $4)",
						event.TxHash.Hex(),
						tx.Hash().Hex(),
						int64(event.BlockNumber),
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
				case ERC20DepositInitiated:
					// execute on l2
					deposit, _ := l1Bridge.IL1ERC20BridgeFilterer.ParseERC20DepositInitiated(event)
					log.WithFields(log.Fields{
						"account": deposit.From.Hex(),
						"to":      deposit.To.Hex(),
						"amount":  deposit.Amount.String(),
						"data":    common.BytesToHash(deposit.Data),
					}).Info("Deposit initiated.")

					time.Sleep(time.Duration(delay) * time.Second)

					// Finalize deposit
					_, err := l2Bridge.FinalizeDeposit(ownerl2.signer, deposit.L1Token, deposit.L2Token, deposit.From, deposit.To, deposit.Amount, deposit.Data)
					if err != nil {
						log.WithField("err", err).Error("finalize failed")
					}

					// Mint the amount
					_, err = l2Usdc.Mint(ownerl2.signer, ownerl2.address, deposit.Amount)
					if err != nil {
						log.WithField("err", err).Error("mint failed")
					}

					// Approve the amount
					_, err = l2Usdc.Approve(ownerl2.signer, accountsAddress, deposit.Amount)
					if err != nil {
						log.WithField("err", err).Error("approve failed")
					}

					// Deposit
					_, err = accounts.Deposit(ownerl2.signer, deposit.To, deposit.L2Token, deposit.Amount)
					if err != nil {
						log.WithField("err", err).Error("deposit failed")
					}
				default:
					continue
				}
			}
		case <-ctx.Done():
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
	Account          common.Address `json:"account"`
	USDC             common.Address `json:"usdc"`
	L1DepositHelper  common.Address `json:"l1DepositHelper"`
	L1StandardBridge common.Address `json:"l1StandardBridge"`
	L2StandardBridge common.Address `json:"l2StandardBridge"`
	L2WithdrawProxy  common.Address `json:"l2WithdrawProxy"`
	L2USDC           common.Address `json:"l2usdc"`

	Instruments common.Address `json:"instruments"`
	Exchange    common.Address `json:"exchange"`
	Executor    common.Address `json:"executor"`
	L2WETH      common.Address `json:"l2weth"`
	L2WBTC      common.Address `json:"l2wbtc"`
}
