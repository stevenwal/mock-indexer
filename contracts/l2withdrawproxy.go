// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// L2WithdrawProxyMetaData contains all meta data concerning the L2WithdrawProxy contract.
var L2WithdrawProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accounts\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"contractIL2ERC20Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161046438038061046483398101604081905261002f91610062565b6001600160a01b039182166080521660a052610095565b80516001600160a01b038116811461005d57600080fd5b919050565b6000806040838503121561007557600080fd5b61007e83610046565b915061008c60208401610046565b90509250929050565b60805160a05161039f6100c560003960008181604b015260ea015260008181609b01526101de015261039f6000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806368cd03f614610046578063ae1f6aaf14610096578063f37e8d38146100bd575b600080fd5b61006d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b61006d7f000000000000000000000000000000000000000000000000000000000000000081565b6100d06100cb366004610298565b6100d2565b005b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610175576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4e4f545f4143434f554e54530000000000000000000000000000000000000000604482015260640160405180910390fd5b6040517fa3a7954800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015285811660248301526044820184905260006064830181905260a0608484015260a48301527f0000000000000000000000000000000000000000000000000000000000000000169063a3a795489060c401600060405180830381600087803b15801561022257600080fd5b505af1158015610236573d6000803e3d6000fd5b5050505050505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461026457600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600080608085870312156102ae57600080fd5b6102b785610240565b93506102c560208601610240565b925060408501359150606085013567ffffffffffffffff808211156102e957600080fd5b818701915087601f8301126102fd57600080fd5b81358181111561030f5761030f610269565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561035557610355610269565b816040528281528a602084870101111561036e57600080fd5b8260208601602083013760006020848301015280955050505050509295919450925056fea164736f6c634300080f000a",
}

// L2WithdrawProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use L2WithdrawProxyMetaData.ABI instead.
var L2WithdrawProxyABI = L2WithdrawProxyMetaData.ABI

// L2WithdrawProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2WithdrawProxyMetaData.Bin instead.
var L2WithdrawProxyBin = L2WithdrawProxyMetaData.Bin

// DeployL2WithdrawProxy deploys a new Ethereum contract, binding an instance of L2WithdrawProxy to it.
func DeployL2WithdrawProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _l2Bridge common.Address, _accounts common.Address) (common.Address, *types.Transaction, *L2WithdrawProxy, error) {
	parsed, err := L2WithdrawProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2WithdrawProxyBin), backend, _l2Bridge, _accounts)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2WithdrawProxy{L2WithdrawProxyCaller: L2WithdrawProxyCaller{contract: contract}, L2WithdrawProxyTransactor: L2WithdrawProxyTransactor{contract: contract}, L2WithdrawProxyFilterer: L2WithdrawProxyFilterer{contract: contract}}, nil
}

// L2WithdrawProxy is an auto generated Go binding around an Ethereum contract.
type L2WithdrawProxy struct {
	L2WithdrawProxyCaller     // Read-only binding to the contract
	L2WithdrawProxyTransactor // Write-only binding to the contract
	L2WithdrawProxyFilterer   // Log filterer for contract events
}

// L2WithdrawProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2WithdrawProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WithdrawProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2WithdrawProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WithdrawProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2WithdrawProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WithdrawProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2WithdrawProxySession struct {
	Contract     *L2WithdrawProxy  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2WithdrawProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2WithdrawProxyCallerSession struct {
	Contract *L2WithdrawProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L2WithdrawProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2WithdrawProxyTransactorSession struct {
	Contract     *L2WithdrawProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L2WithdrawProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2WithdrawProxyRaw struct {
	Contract *L2WithdrawProxy // Generic contract binding to access the raw methods on
}

// L2WithdrawProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2WithdrawProxyCallerRaw struct {
	Contract *L2WithdrawProxyCaller // Generic read-only contract binding to access the raw methods on
}

// L2WithdrawProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2WithdrawProxyTransactorRaw struct {
	Contract *L2WithdrawProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2WithdrawProxy creates a new instance of L2WithdrawProxy, bound to a specific deployed contract.
func NewL2WithdrawProxy(address common.Address, backend bind.ContractBackend) (*L2WithdrawProxy, error) {
	contract, err := bindL2WithdrawProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawProxy{L2WithdrawProxyCaller: L2WithdrawProxyCaller{contract: contract}, L2WithdrawProxyTransactor: L2WithdrawProxyTransactor{contract: contract}, L2WithdrawProxyFilterer: L2WithdrawProxyFilterer{contract: contract}}, nil
}

// NewL2WithdrawProxyCaller creates a new read-only instance of L2WithdrawProxy, bound to a specific deployed contract.
func NewL2WithdrawProxyCaller(address common.Address, caller bind.ContractCaller) (*L2WithdrawProxyCaller, error) {
	contract, err := bindL2WithdrawProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawProxyCaller{contract: contract}, nil
}

// NewL2WithdrawProxyTransactor creates a new write-only instance of L2WithdrawProxy, bound to a specific deployed contract.
func NewL2WithdrawProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*L2WithdrawProxyTransactor, error) {
	contract, err := bindL2WithdrawProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawProxyTransactor{contract: contract}, nil
}

// NewL2WithdrawProxyFilterer creates a new log filterer instance of L2WithdrawProxy, bound to a specific deployed contract.
func NewL2WithdrawProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*L2WithdrawProxyFilterer, error) {
	contract, err := bindL2WithdrawProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2WithdrawProxyFilterer{contract: contract}, nil
}

// bindL2WithdrawProxy binds a generic wrapper to an already deployed contract.
func bindL2WithdrawProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2WithdrawProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2WithdrawProxy *L2WithdrawProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2WithdrawProxy.Contract.L2WithdrawProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2WithdrawProxy *L2WithdrawProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WithdrawProxy.Contract.L2WithdrawProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2WithdrawProxy *L2WithdrawProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2WithdrawProxy.Contract.L2WithdrawProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2WithdrawProxy *L2WithdrawProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2WithdrawProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2WithdrawProxy *L2WithdrawProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WithdrawProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2WithdrawProxy *L2WithdrawProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2WithdrawProxy.Contract.contract.Transact(opts, method, params...)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address)
func (_L2WithdrawProxy *L2WithdrawProxyCaller) Accounts(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WithdrawProxy.contract.Call(opts, &out, "accounts")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address)
func (_L2WithdrawProxy *L2WithdrawProxySession) Accounts() (common.Address, error) {
	return _L2WithdrawProxy.Contract.Accounts(&_L2WithdrawProxy.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address)
func (_L2WithdrawProxy *L2WithdrawProxyCallerSession) Accounts() (common.Address, error) {
	return _L2WithdrawProxy.Contract.Accounts(&_L2WithdrawProxy.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2WithdrawProxy *L2WithdrawProxyCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WithdrawProxy.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2WithdrawProxy *L2WithdrawProxySession) L2Bridge() (common.Address, error) {
	return _L2WithdrawProxy.Contract.L2Bridge(&_L2WithdrawProxy.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2WithdrawProxy *L2WithdrawProxyCallerSession) L2Bridge() (common.Address, error) {
	return _L2WithdrawProxy.Contract.L2Bridge(&_L2WithdrawProxy.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf37e8d38.
//
// Solidity: function withdraw(address account, address collateral, uint256 amount, bytes ) returns()
func (_L2WithdrawProxy *L2WithdrawProxyTransactor) Withdraw(opts *bind.TransactOpts, account common.Address, collateral common.Address, amount *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2WithdrawProxy.contract.Transact(opts, "withdraw", account, collateral, amount, arg3)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf37e8d38.
//
// Solidity: function withdraw(address account, address collateral, uint256 amount, bytes ) returns()
func (_L2WithdrawProxy *L2WithdrawProxySession) Withdraw(account common.Address, collateral common.Address, amount *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2WithdrawProxy.Contract.Withdraw(&_L2WithdrawProxy.TransactOpts, account, collateral, amount, arg3)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf37e8d38.
//
// Solidity: function withdraw(address account, address collateral, uint256 amount, bytes ) returns()
func (_L2WithdrawProxy *L2WithdrawProxyTransactorSession) Withdraw(account common.Address, collateral common.Address, amount *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2WithdrawProxy.Contract.Withdraw(&_L2WithdrawProxy.TransactOpts, account, collateral, amount, arg3)
}
