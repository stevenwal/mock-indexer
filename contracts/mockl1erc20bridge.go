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

// MockL1ERC20BridgeMetaData contains all meta data concerning the MockL1ERC20Bridge contract.
var MockL1ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Bridge\",\"type\":\"address\"}],\"name\":\"setl2TokenBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061059f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c806391c49bf81161005057806391c49bf8146100f4578063a9f9e67514610137578063ae1f6aaf1461014a57600080fd5b8063425993111461007757806358a997f6146100ce578063838b2520146100e1575b600080fd5b6100cc610085366004610322565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b005b6100cc6100dc3660046103a1565b61016a565b6100cc6100ef366004610420565b6101f4565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100cc6101453660046104b0565b61027f565b60005461010e9073ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396338887876040516101e49493929190610521565b60405180910390a4505050505050565b3373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d03968888878760405161026e9493929190610521565b60405180910390a450505050505050565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b38787878760405161026e9493929190610521565b803573ffffffffffffffffffffffffffffffffffffffff8116811461031d57600080fd5b919050565b60006020828403121561033457600080fd5b61033d826102f9565b9392505050565b803563ffffffff8116811461031d57600080fd5b60008083601f84011261036a57600080fd5b50813567ffffffffffffffff81111561038257600080fd5b60208301915083602082850101111561039a57600080fd5b9250929050565b60008060008060008060a087890312156103ba57600080fd5b6103c3876102f9565b95506103d1602088016102f9565b9450604087013593506103e660608801610344565b9250608087013567ffffffffffffffff81111561040257600080fd5b61040e89828a01610358565b979a9699509497509295939492505050565b600080600080600080600060c0888a03121561043b57600080fd5b610444886102f9565b9650610452602089016102f9565b9550610460604089016102f9565b94506060880135935061047560808901610344565b925060a088013567ffffffffffffffff81111561049157600080fd5b61049d8a828b01610358565b989b979a50959850939692959293505050565b600080600080600080600060c0888a0312156104cb57600080fd5b6104d4886102f9565b96506104e2602089016102f9565b95506104f0604089016102f9565b94506104fe606089016102f9565b93506080880135925060a088013567ffffffffffffffff81111561049157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152816060820152818360808301376000818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101939250505056fea164736f6c634300080f000a",
}

// MockL1ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use MockL1ERC20BridgeMetaData.ABI instead.
var MockL1ERC20BridgeABI = MockL1ERC20BridgeMetaData.ABI

// MockL1ERC20BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockL1ERC20BridgeMetaData.Bin instead.
var MockL1ERC20BridgeBin = MockL1ERC20BridgeMetaData.Bin

// DeployMockL1ERC20Bridge deploys a new Ethereum contract, binding an instance of MockL1ERC20Bridge to it.
func DeployMockL1ERC20Bridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockL1ERC20Bridge, error) {
	parsed, err := MockL1ERC20BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockL1ERC20BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockL1ERC20Bridge{MockL1ERC20BridgeCaller: MockL1ERC20BridgeCaller{contract: contract}, MockL1ERC20BridgeTransactor: MockL1ERC20BridgeTransactor{contract: contract}, MockL1ERC20BridgeFilterer: MockL1ERC20BridgeFilterer{contract: contract}}, nil
}

// MockL1ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type MockL1ERC20Bridge struct {
	MockL1ERC20BridgeCaller     // Read-only binding to the contract
	MockL1ERC20BridgeTransactor // Write-only binding to the contract
	MockL1ERC20BridgeFilterer   // Log filterer for contract events
}

// MockL1ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockL1ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL1ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockL1ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL1ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockL1ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL1ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockL1ERC20BridgeSession struct {
	Contract     *MockL1ERC20Bridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MockL1ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockL1ERC20BridgeCallerSession struct {
	Contract *MockL1ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MockL1ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockL1ERC20BridgeTransactorSession struct {
	Contract     *MockL1ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MockL1ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockL1ERC20BridgeRaw struct {
	Contract *MockL1ERC20Bridge // Generic contract binding to access the raw methods on
}

// MockL1ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockL1ERC20BridgeCallerRaw struct {
	Contract *MockL1ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// MockL1ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockL1ERC20BridgeTransactorRaw struct {
	Contract *MockL1ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockL1ERC20Bridge creates a new instance of MockL1ERC20Bridge, bound to a specific deployed contract.
func NewMockL1ERC20Bridge(address common.Address, backend bind.ContractBackend) (*MockL1ERC20Bridge, error) {
	contract, err := bindMockL1ERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockL1ERC20Bridge{MockL1ERC20BridgeCaller: MockL1ERC20BridgeCaller{contract: contract}, MockL1ERC20BridgeTransactor: MockL1ERC20BridgeTransactor{contract: contract}, MockL1ERC20BridgeFilterer: MockL1ERC20BridgeFilterer{contract: contract}}, nil
}

// NewMockL1ERC20BridgeCaller creates a new read-only instance of MockL1ERC20Bridge, bound to a specific deployed contract.
func NewMockL1ERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*MockL1ERC20BridgeCaller, error) {
	contract, err := bindMockL1ERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockL1ERC20BridgeCaller{contract: contract}, nil
}

// NewMockL1ERC20BridgeTransactor creates a new write-only instance of MockL1ERC20Bridge, bound to a specific deployed contract.
func NewMockL1ERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*MockL1ERC20BridgeTransactor, error) {
	contract, err := bindMockL1ERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockL1ERC20BridgeTransactor{contract: contract}, nil
}

// NewMockL1ERC20BridgeFilterer creates a new log filterer instance of MockL1ERC20Bridge, bound to a specific deployed contract.
func NewMockL1ERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*MockL1ERC20BridgeFilterer, error) {
	contract, err := bindMockL1ERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockL1ERC20BridgeFilterer{contract: contract}, nil
}

// bindMockL1ERC20Bridge binds a generic wrapper to an already deployed contract.
func bindMockL1ERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockL1ERC20BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockL1ERC20Bridge *MockL1ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL1ERC20Bridge.Contract.MockL1ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockL1ERC20Bridge *MockL1ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.MockL1ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockL1ERC20Bridge *MockL1ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.MockL1ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockL1ERC20Bridge *MockL1ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL1ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockL1ERC20Bridge.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) L2Bridge() (common.Address, error) {
	return _MockL1ERC20Bridge.Contract.L2Bridge(&_MockL1ERC20Bridge.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeCallerSession) L2Bridge() (common.Address, error) {
	return _MockL1ERC20Bridge.Contract.L2Bridge(&_MockL1ERC20Bridge.CallOpts)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeCaller) L2TokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockL1ERC20Bridge.contract.Call(opts, &out, "l2TokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) L2TokenBridge() (common.Address, error) {
	return _MockL1ERC20Bridge.Contract.L2TokenBridge(&_MockL1ERC20Bridge.CallOpts)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeCallerSession) L2TokenBridge() (common.Address, error) {
	return _MockL1ERC20Bridge.Contract.L2TokenBridge(&_MockL1ERC20Bridge.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactor) DepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _amount *big.Int, arg3 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.contract.Transact(opts, "depositERC20", _l1Token, _l2Token, _amount, arg3, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, arg3 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.DepositERC20(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _amount, arg3, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactorSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, arg3 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.DepositERC20(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _amount, arg3, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, arg4 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, arg4, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, arg4 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.DepositERC20To(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _to, _amount, arg4, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, arg4 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.DepositERC20To(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _to, _amount, arg4, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.FinalizeERC20Withdrawal(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.FinalizeERC20Withdrawal(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Setl2TokenBridge is a paid mutator transaction binding the contract method 0x42599311.
//
// Solidity: function setl2TokenBridge(address _l2Bridge) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactor) Setl2TokenBridge(opts *bind.TransactOpts, _l2Bridge common.Address) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.contract.Transact(opts, "setl2TokenBridge", _l2Bridge)
}

// Setl2TokenBridge is a paid mutator transaction binding the contract method 0x42599311.
//
// Solidity: function setl2TokenBridge(address _l2Bridge) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) Setl2TokenBridge(_l2Bridge common.Address) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.Setl2TokenBridge(&_MockL1ERC20Bridge.TransactOpts, _l2Bridge)
}

// Setl2TokenBridge is a paid mutator transaction binding the contract method 0x42599311.
//
// Solidity: function setl2TokenBridge(address _l2Bridge) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactorSession) Setl2TokenBridge(_l2Bridge common.Address) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.Setl2TokenBridge(&_MockL1ERC20Bridge.TransactOpts, _l2Bridge)
}

// MockL1ERC20BridgeERC20DepositInitiatedIterator is returned from FilterERC20DepositInitiated and is used to iterate over the raw logs and unpacked data for ERC20DepositInitiated events raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeERC20DepositInitiatedIterator struct {
	Event *MockL1ERC20BridgeERC20DepositInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockL1ERC20BridgeERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1ERC20BridgeERC20DepositInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockL1ERC20BridgeERC20DepositInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockL1ERC20BridgeERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1ERC20BridgeERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1ERC20BridgeERC20DepositInitiated represents a ERC20DepositInitiated event raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeERC20DepositInitiated struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositInitiated is a free log retrieval operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) FilterERC20DepositInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*MockL1ERC20BridgeERC20DepositInitiatedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _MockL1ERC20Bridge.contract.FilterLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &MockL1ERC20BridgeERC20DepositInitiatedIterator{contract: _MockL1ERC20Bridge.contract, event: "ERC20DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20DepositInitiated is a free log subscription operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) WatchERC20DepositInitiated(opts *bind.WatchOpts, sink chan<- *MockL1ERC20BridgeERC20DepositInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _MockL1ERC20Bridge.contract.WatchLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1ERC20BridgeERC20DepositInitiated)
				if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20DepositInitiated is a log parse operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) ParseERC20DepositInitiated(log types.Log) (*MockL1ERC20BridgeERC20DepositInitiated, error) {
	event := new(MockL1ERC20BridgeERC20DepositInitiated)
	if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL1ERC20BridgeERC20WithdrawalFinalizedIterator is returned from FilterERC20WithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ERC20WithdrawalFinalized events raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeERC20WithdrawalFinalizedIterator struct {
	Event *MockL1ERC20BridgeERC20WithdrawalFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockL1ERC20BridgeERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1ERC20BridgeERC20WithdrawalFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockL1ERC20BridgeERC20WithdrawalFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockL1ERC20BridgeERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1ERC20BridgeERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1ERC20BridgeERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeERC20WithdrawalFinalized struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC20WithdrawalFinalized is a free log retrieval operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) FilterERC20WithdrawalFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*MockL1ERC20BridgeERC20WithdrawalFinalizedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _MockL1ERC20Bridge.contract.FilterLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &MockL1ERC20BridgeERC20WithdrawalFinalizedIterator{contract: _MockL1ERC20Bridge.contract, event: "ERC20WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20WithdrawalFinalized is a free log subscription operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) WatchERC20WithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *MockL1ERC20BridgeERC20WithdrawalFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _MockL1ERC20Bridge.contract.WatchLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1ERC20BridgeERC20WithdrawalFinalized)
				if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20WithdrawalFinalized is a log parse operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) ParseERC20WithdrawalFinalized(log types.Log) (*MockL1ERC20BridgeERC20WithdrawalFinalized, error) {
	event := new(MockL1ERC20BridgeERC20WithdrawalFinalized)
	if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
