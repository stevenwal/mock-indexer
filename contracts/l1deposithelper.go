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

// L1DepositHelperMetaData contains all meta data concerning the L1DepositHelper contract.
var L1DepositHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"depositERC20ToWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"depositERC20WithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Bridge\",\"outputs\":[{\"internalType\":\"contractIL1ERC20Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161092838038061092883398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516108826100a6600039600081816060015281816101a7015281816102090152818161037201526103d401526108826000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630e47e9e314610046578063969b53da1461005b578063b4b62609146100ab575b600080fd5b610059610054366004610674565b6100be565b005b6100827f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b9366004610733565b610289565b6040517fd505accf000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018990526064810185905260ff8416608482015260a4810183905260c4810182905273ffffffffffffffffffffffffffffffffffffffff8c169063d505accf9060e401600060405180830381600087803b15801561015057600080fd5b505af1158015610164573d6000803e3d6000fd5b5061018b9250505073ffffffffffffffffffffffffffffffffffffffff8c1633308b610453565b6101cc73ffffffffffffffffffffffffffffffffffffffff8c167f00000000000000000000000000000000000000000000000000000000000000008a61051e565b6040517f838b252000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063838b25209061024a908e908e908e908e908e908e908e906004016107e0565b600060405180830381600087803b15801561026457600080fd5b505af1158015610278573d6000803e3d6000fd5b505050505050505050505050505050565b6040517fd505accf000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018990526064810185905260ff8416608482015260a4810183905260c4810182905273ffffffffffffffffffffffffffffffffffffffff8b169063d505accf9060e401600060405180830381600087803b15801561031b57600080fd5b505af115801561032f573d6000803e3d6000fd5b506103569250505073ffffffffffffffffffffffffffffffffffffffff8b1633308b610453565b61039773ffffffffffffffffffffffffffffffffffffffff8b167f00000000000000000000000000000000000000000000000000000000000000008a61051e565b6040517f838b252000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063838b252090610415908d908d9033908e908e908e908e906004016107e0565b600060405180830381600087803b15801561042f57600080fd5b505af1158015610443573d6000803e3d6000fd5b5050505050505050505050505050565b60006040517f23b872dd0000000000000000000000000000000000000000000000000000000081528460048201528360248201528260448201526020600060648360008a5af13d15601f3d1160016000511416171691505080610517576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5452414e534645525f46524f4d5f4641494c454400000000000000000000000060448201526064015b60405180910390fd5b5050505050565b60006040517f095ea7b3000000000000000000000000000000000000000000000000000000008152836004820152826024820152602060006044836000895af13d15601f3d11600160005114161716915050806105d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f415050524f56455f4641494c4544000000000000000000000000000000000000604482015260640161050e565b50505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461060157600080fd5b919050565b803563ffffffff8116811461060157600080fd5b60008083601f84011261062c57600080fd5b50813567ffffffffffffffff81111561064457600080fd5b60208301915083602082850101111561065c57600080fd5b9250929050565b803560ff8116811461060157600080fd5b60008060008060008060008060008060006101408c8e03121561069657600080fd5b61069f8c6105dd565b9a506106ad60208d016105dd565b99506106bb60408d016105dd565b985060608c013597506106d060808d01610606565b965060a08c013567ffffffffffffffff8111156106ec57600080fd5b6106f88e828f0161061a565b90975095505060c08c0135935061071160e08d01610663565b92506101008c013591506101208c013590509295989b509295989b9093969950565b6000806000806000806000806000806101208b8d03121561075357600080fd5b61075c8b6105dd565b995061076a60208c016105dd565b985060408b0135975061077f60608c01610606565b965060808b013567ffffffffffffffff81111561079b57600080fd5b6107a78d828e0161061a565b90975095505060a08b013593506107c060c08c01610663565b925060e08b013591506101008b013590509295989b9194979a5092959850565b600073ffffffffffffffffffffffffffffffffffffffff808a168352808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a08301528260c0830152828460e0840137600060e0848401015260e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011683010190509897505050505050505056fea164736f6c634300080f000a",
}

// L1DepositHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use L1DepositHelperMetaData.ABI instead.
var L1DepositHelperABI = L1DepositHelperMetaData.ABI

// L1DepositHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1DepositHelperMetaData.Bin instead.
var L1DepositHelperBin = L1DepositHelperMetaData.Bin

// DeployL1DepositHelper deploys a new Ethereum contract, binding an instance of L1DepositHelper to it.
func DeployL1DepositHelper(auth *bind.TransactOpts, backend bind.ContractBackend, _l1Bridge common.Address) (common.Address, *types.Transaction, *L1DepositHelper, error) {
	parsed, err := L1DepositHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1DepositHelperBin), backend, _l1Bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1DepositHelper{L1DepositHelperCaller: L1DepositHelperCaller{contract: contract}, L1DepositHelperTransactor: L1DepositHelperTransactor{contract: contract}, L1DepositHelperFilterer: L1DepositHelperFilterer{contract: contract}}, nil
}

// L1DepositHelper is an auto generated Go binding around an Ethereum contract.
type L1DepositHelper struct {
	L1DepositHelperCaller     // Read-only binding to the contract
	L1DepositHelperTransactor // Write-only binding to the contract
	L1DepositHelperFilterer   // Log filterer for contract events
}

// L1DepositHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1DepositHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1DepositHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1DepositHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1DepositHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1DepositHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1DepositHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1DepositHelperSession struct {
	Contract     *L1DepositHelper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1DepositHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1DepositHelperCallerSession struct {
	Contract *L1DepositHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L1DepositHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1DepositHelperTransactorSession struct {
	Contract     *L1DepositHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L1DepositHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1DepositHelperRaw struct {
	Contract *L1DepositHelper // Generic contract binding to access the raw methods on
}

// L1DepositHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1DepositHelperCallerRaw struct {
	Contract *L1DepositHelperCaller // Generic read-only contract binding to access the raw methods on
}

// L1DepositHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1DepositHelperTransactorRaw struct {
	Contract *L1DepositHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1DepositHelper creates a new instance of L1DepositHelper, bound to a specific deployed contract.
func NewL1DepositHelper(address common.Address, backend bind.ContractBackend) (*L1DepositHelper, error) {
	contract, err := bindL1DepositHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1DepositHelper{L1DepositHelperCaller: L1DepositHelperCaller{contract: contract}, L1DepositHelperTransactor: L1DepositHelperTransactor{contract: contract}, L1DepositHelperFilterer: L1DepositHelperFilterer{contract: contract}}, nil
}

// NewL1DepositHelperCaller creates a new read-only instance of L1DepositHelper, bound to a specific deployed contract.
func NewL1DepositHelperCaller(address common.Address, caller bind.ContractCaller) (*L1DepositHelperCaller, error) {
	contract, err := bindL1DepositHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1DepositHelperCaller{contract: contract}, nil
}

// NewL1DepositHelperTransactor creates a new write-only instance of L1DepositHelper, bound to a specific deployed contract.
func NewL1DepositHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*L1DepositHelperTransactor, error) {
	contract, err := bindL1DepositHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1DepositHelperTransactor{contract: contract}, nil
}

// NewL1DepositHelperFilterer creates a new log filterer instance of L1DepositHelper, bound to a specific deployed contract.
func NewL1DepositHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*L1DepositHelperFilterer, error) {
	contract, err := bindL1DepositHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1DepositHelperFilterer{contract: contract}, nil
}

// bindL1DepositHelper binds a generic wrapper to an already deployed contract.
func bindL1DepositHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1DepositHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1DepositHelper *L1DepositHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1DepositHelper.Contract.L1DepositHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1DepositHelper *L1DepositHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1DepositHelper.Contract.L1DepositHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1DepositHelper *L1DepositHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1DepositHelper.Contract.L1DepositHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1DepositHelper *L1DepositHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1DepositHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1DepositHelper *L1DepositHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1DepositHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1DepositHelper *L1DepositHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1DepositHelper.Contract.contract.Transact(opts, method, params...)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_L1DepositHelper *L1DepositHelperCaller) L1Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1DepositHelper.contract.Call(opts, &out, "l1Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_L1DepositHelper *L1DepositHelperSession) L1Bridge() (common.Address, error) {
	return _L1DepositHelper.Contract.L1Bridge(&_L1DepositHelper.CallOpts)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_L1DepositHelper *L1DepositHelperCallerSession) L1Bridge() (common.Address, error) {
	return _L1DepositHelper.Contract.L1Bridge(&_L1DepositHelper.CallOpts)
}

// DepositERC20ToWithPermit is a paid mutator transaction binding the contract method 0x0e47e9e3.
//
// Solidity: function depositERC20ToWithPermit(address l1Token, address l2Token, address to, uint256 amount, uint32 l2Gas, bytes data, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_L1DepositHelper *L1DepositHelperTransactor) DepositERC20ToWithPermit(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, to common.Address, amount *big.Int, l2Gas uint32, data []byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _L1DepositHelper.contract.Transact(opts, "depositERC20ToWithPermit", l1Token, l2Token, to, amount, l2Gas, data, deadline, v, r, s)
}

// DepositERC20ToWithPermit is a paid mutator transaction binding the contract method 0x0e47e9e3.
//
// Solidity: function depositERC20ToWithPermit(address l1Token, address l2Token, address to, uint256 amount, uint32 l2Gas, bytes data, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_L1DepositHelper *L1DepositHelperSession) DepositERC20ToWithPermit(l1Token common.Address, l2Token common.Address, to common.Address, amount *big.Int, l2Gas uint32, data []byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _L1DepositHelper.Contract.DepositERC20ToWithPermit(&_L1DepositHelper.TransactOpts, l1Token, l2Token, to, amount, l2Gas, data, deadline, v, r, s)
}

// DepositERC20ToWithPermit is a paid mutator transaction binding the contract method 0x0e47e9e3.
//
// Solidity: function depositERC20ToWithPermit(address l1Token, address l2Token, address to, uint256 amount, uint32 l2Gas, bytes data, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_L1DepositHelper *L1DepositHelperTransactorSession) DepositERC20ToWithPermit(l1Token common.Address, l2Token common.Address, to common.Address, amount *big.Int, l2Gas uint32, data []byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _L1DepositHelper.Contract.DepositERC20ToWithPermit(&_L1DepositHelper.TransactOpts, l1Token, l2Token, to, amount, l2Gas, data, deadline, v, r, s)
}

// DepositERC20WithPermit is a paid mutator transaction binding the contract method 0xb4b62609.
//
// Solidity: function depositERC20WithPermit(address l1Token, address l2Token, uint256 amount, uint32 l2Gas, bytes data, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_L1DepositHelper *L1DepositHelperTransactor) DepositERC20WithPermit(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, amount *big.Int, l2Gas uint32, data []byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _L1DepositHelper.contract.Transact(opts, "depositERC20WithPermit", l1Token, l2Token, amount, l2Gas, data, deadline, v, r, s)
}

// DepositERC20WithPermit is a paid mutator transaction binding the contract method 0xb4b62609.
//
// Solidity: function depositERC20WithPermit(address l1Token, address l2Token, uint256 amount, uint32 l2Gas, bytes data, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_L1DepositHelper *L1DepositHelperSession) DepositERC20WithPermit(l1Token common.Address, l2Token common.Address, amount *big.Int, l2Gas uint32, data []byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _L1DepositHelper.Contract.DepositERC20WithPermit(&_L1DepositHelper.TransactOpts, l1Token, l2Token, amount, l2Gas, data, deadline, v, r, s)
}

// DepositERC20WithPermit is a paid mutator transaction binding the contract method 0xb4b62609.
//
// Solidity: function depositERC20WithPermit(address l1Token, address l2Token, uint256 amount, uint32 l2Gas, bytes data, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_L1DepositHelper *L1DepositHelperTransactorSession) DepositERC20WithPermit(l1Token common.Address, l2Token common.Address, amount *big.Int, l2Gas uint32, data []byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _L1DepositHelper.Contract.DepositERC20WithPermit(&_L1DepositHelper.TransactOpts, l1Token, l2Token, amount, l2Gas, data, deadline, v, r, s)
}
