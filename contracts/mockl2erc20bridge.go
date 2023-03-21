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

// MockL2ERC20BridgeMetaData contains all meta data concerning the MockL2ERC20Bridge contract.
var MockL2ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DepositFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516105b03803806105b083398101604081905261002f91610062565b6001600160a01b039182166080521660a052610095565b80516001600160a01b038116811461005d57600080fd5b919050565b6000806040838503121561007557600080fd5b61007e83610046565b915061008c60208401610046565b90509250929050565b60805160a0516104e36100cd6000396000818161010c01528181610144015261022a015260008181607e015260d201526104e36000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806332b7006d1461006757806336c717c11461007c578063662a633a146100ba578063969b53da146100cd578063a3a79548146100f4578063c01e1bd614610107575b600080fd5b61007a610075366004610310565b61012e565b005b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b03909116815260200160405180910390f35b61007a6100c836600461037f565b6101b0565b61009e7f000000000000000000000000000000000000000000000000000000000000000081565b61007a61010236600461040f565b610214565b61009e7f000000000000000000000000000000000000000000000000000000000000000081565b336001600160a01b0316856001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e338887876040516101a1949392919061048e565b60405180910390a45050505050565b846001600160a01b0316866001600160a01b0316886001600160a01b03167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd8987878787604051610203949392919061048e565b60405180910390a450505050505050565b336001600160a01b0316866001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e88888787604051610287949392919061048e565b60405180910390a4505050505050565b80356001600160a01b03811681146102ae57600080fd5b919050565b803563ffffffff811681146102ae57600080fd5b60008083601f8401126102d957600080fd5b50813567ffffffffffffffff8111156102f157600080fd5b60208301915083602082850101111561030957600080fd5b9250929050565b60008060008060006080868803121561032857600080fd5b61033186610297565b945060208601359350610346604087016102b3565b9250606086013567ffffffffffffffff81111561036257600080fd5b61036e888289016102c7565b969995985093965092949392505050565b600080600080600080600060c0888a03121561039a57600080fd5b6103a388610297565b96506103b160208901610297565b95506103bf60408901610297565b94506103cd60608901610297565b93506080880135925060a088013567ffffffffffffffff8111156103f057600080fd5b6103fc8a828b016102c7565b989b979a50959850939692959293505050565b60008060008060008060a0878903121561042857600080fd5b61043187610297565b955061043f60208801610297565b945060408701359350610454606088016102b3565b9250608087013567ffffffffffffffff81111561047057600080fd5b61047c89828a016102c7565b979a9699509497509295939492505050565b6001600160a01b0385168152602081018490526060604082018190528101829052818360808301376000818301608090810191909152601f909201601f19160101939250505056fea164736f6c634300080f000a",
}

// MockL2ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use MockL2ERC20BridgeMetaData.ABI instead.
var MockL2ERC20BridgeABI = MockL2ERC20BridgeMetaData.ABI

// MockL2ERC20BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockL2ERC20BridgeMetaData.Bin instead.
var MockL2ERC20BridgeBin = MockL2ERC20BridgeMetaData.Bin

// DeployMockL2ERC20Bridge deploys a new Ethereum contract, binding an instance of MockL2ERC20Bridge to it.
func DeployMockL2ERC20Bridge(auth *bind.TransactOpts, backend bind.ContractBackend, _l1Bridge common.Address, _l1Token common.Address) (common.Address, *types.Transaction, *MockL2ERC20Bridge, error) {
	parsed, err := MockL2ERC20BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockL2ERC20BridgeBin), backend, _l1Bridge, _l1Token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockL2ERC20Bridge{MockL2ERC20BridgeCaller: MockL2ERC20BridgeCaller{contract: contract}, MockL2ERC20BridgeTransactor: MockL2ERC20BridgeTransactor{contract: contract}, MockL2ERC20BridgeFilterer: MockL2ERC20BridgeFilterer{contract: contract}}, nil
}

// MockL2ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type MockL2ERC20Bridge struct {
	MockL2ERC20BridgeCaller     // Read-only binding to the contract
	MockL2ERC20BridgeTransactor // Write-only binding to the contract
	MockL2ERC20BridgeFilterer   // Log filterer for contract events
}

// MockL2ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockL2ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL2ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockL2ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL2ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockL2ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL2ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockL2ERC20BridgeSession struct {
	Contract     *MockL2ERC20Bridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MockL2ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockL2ERC20BridgeCallerSession struct {
	Contract *MockL2ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MockL2ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockL2ERC20BridgeTransactorSession struct {
	Contract     *MockL2ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MockL2ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockL2ERC20BridgeRaw struct {
	Contract *MockL2ERC20Bridge // Generic contract binding to access the raw methods on
}

// MockL2ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockL2ERC20BridgeCallerRaw struct {
	Contract *MockL2ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// MockL2ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockL2ERC20BridgeTransactorRaw struct {
	Contract *MockL2ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockL2ERC20Bridge creates a new instance of MockL2ERC20Bridge, bound to a specific deployed contract.
func NewMockL2ERC20Bridge(address common.Address, backend bind.ContractBackend) (*MockL2ERC20Bridge, error) {
	contract, err := bindMockL2ERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20Bridge{MockL2ERC20BridgeCaller: MockL2ERC20BridgeCaller{contract: contract}, MockL2ERC20BridgeTransactor: MockL2ERC20BridgeTransactor{contract: contract}, MockL2ERC20BridgeFilterer: MockL2ERC20BridgeFilterer{contract: contract}}, nil
}

// NewMockL2ERC20BridgeCaller creates a new read-only instance of MockL2ERC20Bridge, bound to a specific deployed contract.
func NewMockL2ERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*MockL2ERC20BridgeCaller, error) {
	contract, err := bindMockL2ERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20BridgeCaller{contract: contract}, nil
}

// NewMockL2ERC20BridgeTransactor creates a new write-only instance of MockL2ERC20Bridge, bound to a specific deployed contract.
func NewMockL2ERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*MockL2ERC20BridgeTransactor, error) {
	contract, err := bindMockL2ERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20BridgeTransactor{contract: contract}, nil
}

// NewMockL2ERC20BridgeFilterer creates a new log filterer instance of MockL2ERC20Bridge, bound to a specific deployed contract.
func NewMockL2ERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*MockL2ERC20BridgeFilterer, error) {
	contract, err := bindMockL2ERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20BridgeFilterer{contract: contract}, nil
}

// bindMockL2ERC20Bridge binds a generic wrapper to an already deployed contract.
func bindMockL2ERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockL2ERC20BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockL2ERC20Bridge *MockL2ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL2ERC20Bridge.Contract.MockL2ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockL2ERC20Bridge *MockL2ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.MockL2ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockL2ERC20Bridge *MockL2ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.MockL2ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL2ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCaller) L1Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockL2ERC20Bridge.contract.Call(opts, &out, "l1Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) L1Bridge() (common.Address, error) {
	return _MockL2ERC20Bridge.Contract.L1Bridge(&_MockL2ERC20Bridge.CallOpts)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCallerSession) L1Bridge() (common.Address, error) {
	return _MockL2ERC20Bridge.Contract.L1Bridge(&_MockL2ERC20Bridge.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockL2ERC20Bridge.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) L1Token() (common.Address, error) {
	return _MockL2ERC20Bridge.Contract.L1Token(&_MockL2ERC20Bridge.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCallerSession) L1Token() (common.Address, error) {
	return _MockL2ERC20Bridge.Contract.L1Token(&_MockL2ERC20Bridge.CallOpts)
}

// L1TokenBridge is a free data retrieval call binding the contract method 0x36c717c1.
//
// Solidity: function l1TokenBridge() view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCaller) L1TokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockL2ERC20Bridge.contract.Call(opts, &out, "l1TokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1TokenBridge is a free data retrieval call binding the contract method 0x36c717c1.
//
// Solidity: function l1TokenBridge() view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) L1TokenBridge() (common.Address, error) {
	return _MockL2ERC20Bridge.Contract.L1TokenBridge(&_MockL2ERC20Bridge.CallOpts)
}

// L1TokenBridge is a free data retrieval call binding the contract method 0x36c717c1.
//
// Solidity: function l1TokenBridge() view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCallerSession) L1TokenBridge() (common.Address, error) {
	return _MockL2ERC20Bridge.Contract.L1TokenBridge(&_MockL2ERC20Bridge.CallOpts)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x662a633a.
//
// Solidity: function finalizeDeposit(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactor) FinalizeDeposit(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.contract.Transact(opts, "finalizeDeposit", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x662a633a.
//
// Solidity: function finalizeDeposit(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.FinalizeDeposit(&_MockL2ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x662a633a.
//
// Solidity: function finalizeDeposit(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactorSession) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.FinalizeDeposit(&_MockL2ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32b7006d.
//
// Solidity: function withdraw(address _l2Token, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactor) Withdraw(opts *bind.TransactOpts, _l2Token common.Address, _amount *big.Int, arg2 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.contract.Transact(opts, "withdraw", _l2Token, _amount, arg2, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32b7006d.
//
// Solidity: function withdraw(address _l2Token, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) Withdraw(_l2Token common.Address, _amount *big.Int, arg2 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.Withdraw(&_MockL2ERC20Bridge.TransactOpts, _l2Token, _amount, arg2, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32b7006d.
//
// Solidity: function withdraw(address _l2Token, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactorSession) Withdraw(_l2Token common.Address, _amount *big.Int, arg2 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.Withdraw(&_MockL2ERC20Bridge.TransactOpts, _l2Token, _amount, arg2, _data)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xa3a79548.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactor) WithdrawTo(opts *bind.TransactOpts, _l2Token common.Address, _to common.Address, _amount *big.Int, arg3 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.contract.Transact(opts, "withdrawTo", _l2Token, _to, _amount, arg3, _data)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xa3a79548.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, arg3 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.WithdrawTo(&_MockL2ERC20Bridge.TransactOpts, _l2Token, _to, _amount, arg3, _data)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xa3a79548.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 , bytes _data) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactorSession) WithdrawTo(_l2Token common.Address, _to common.Address, _amount *big.Int, arg3 uint32, _data []byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.WithdrawTo(&_MockL2ERC20Bridge.TransactOpts, _l2Token, _to, _amount, arg3, _data)
}

// MockL2ERC20BridgeDepositFailedIterator is returned from FilterDepositFailed and is used to iterate over the raw logs and unpacked data for DepositFailed events raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeDepositFailedIterator struct {
	Event *MockL2ERC20BridgeDepositFailed // Event containing the contract specifics and raw log

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
func (it *MockL2ERC20BridgeDepositFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2ERC20BridgeDepositFailed)
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
		it.Event = new(MockL2ERC20BridgeDepositFailed)
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
func (it *MockL2ERC20BridgeDepositFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2ERC20BridgeDepositFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2ERC20BridgeDepositFailed represents a DepositFailed event raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeDepositFailed struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositFailed is a free log retrieval operation binding the contract event 0x7ea89a4591614515571c2b51f5ea06494056f261c10ab1ed8c03c7590d87bce0.
//
// Solidity: event DepositFailed(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) FilterDepositFailed(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*MockL2ERC20BridgeDepositFailedIterator, error) {

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

	logs, sub, err := _MockL2ERC20Bridge.contract.FilterLogs(opts, "DepositFailed", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20BridgeDepositFailedIterator{contract: _MockL2ERC20Bridge.contract, event: "DepositFailed", logs: logs, sub: sub}, nil
}

// WatchDepositFailed is a free log subscription operation binding the contract event 0x7ea89a4591614515571c2b51f5ea06494056f261c10ab1ed8c03c7590d87bce0.
//
// Solidity: event DepositFailed(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) WatchDepositFailed(opts *bind.WatchOpts, sink chan<- *MockL2ERC20BridgeDepositFailed, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MockL2ERC20Bridge.contract.WatchLogs(opts, "DepositFailed", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2ERC20BridgeDepositFailed)
				if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "DepositFailed", log); err != nil {
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

// ParseDepositFailed is a log parse operation binding the contract event 0x7ea89a4591614515571c2b51f5ea06494056f261c10ab1ed8c03c7590d87bce0.
//
// Solidity: event DepositFailed(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) ParseDepositFailed(log types.Log) (*MockL2ERC20BridgeDepositFailed, error) {
	event := new(MockL2ERC20BridgeDepositFailed)
	if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "DepositFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL2ERC20BridgeDepositFinalizedIterator is returned from FilterDepositFinalized and is used to iterate over the raw logs and unpacked data for DepositFinalized events raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeDepositFinalizedIterator struct {
	Event *MockL2ERC20BridgeDepositFinalized // Event containing the contract specifics and raw log

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
func (it *MockL2ERC20BridgeDepositFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2ERC20BridgeDepositFinalized)
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
		it.Event = new(MockL2ERC20BridgeDepositFinalized)
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
func (it *MockL2ERC20BridgeDepositFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2ERC20BridgeDepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2ERC20BridgeDepositFinalized represents a DepositFinalized event raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeDepositFinalized struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositFinalized is a free log retrieval operation binding the contract event 0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89.
//
// Solidity: event DepositFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) FilterDepositFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*MockL2ERC20BridgeDepositFinalizedIterator, error) {

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

	logs, sub, err := _MockL2ERC20Bridge.contract.FilterLogs(opts, "DepositFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20BridgeDepositFinalizedIterator{contract: _MockL2ERC20Bridge.contract, event: "DepositFinalized", logs: logs, sub: sub}, nil
}

// WatchDepositFinalized is a free log subscription operation binding the contract event 0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89.
//
// Solidity: event DepositFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) WatchDepositFinalized(opts *bind.WatchOpts, sink chan<- *MockL2ERC20BridgeDepositFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MockL2ERC20Bridge.contract.WatchLogs(opts, "DepositFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2ERC20BridgeDepositFinalized)
				if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
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

// ParseDepositFinalized is a log parse operation binding the contract event 0xb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89.
//
// Solidity: event DepositFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) ParseDepositFinalized(log types.Log) (*MockL2ERC20BridgeDepositFinalized, error) {
	event := new(MockL2ERC20BridgeDepositFinalized)
	if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL2ERC20BridgeWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeWithdrawalInitiatedIterator struct {
	Event *MockL2ERC20BridgeWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *MockL2ERC20BridgeWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2ERC20BridgeWithdrawalInitiated)
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
		it.Event = new(MockL2ERC20BridgeWithdrawalInitiated)
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
func (it *MockL2ERC20BridgeWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2ERC20BridgeWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2ERC20BridgeWithdrawalInitiated represents a WithdrawalInitiated event raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeWithdrawalInitiated struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e.
//
// Solidity: event WithdrawalInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*MockL2ERC20BridgeWithdrawalInitiatedIterator, error) {

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

	logs, sub, err := _MockL2ERC20Bridge.contract.FilterLogs(opts, "WithdrawalInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20BridgeWithdrawalInitiatedIterator{contract: _MockL2ERC20Bridge.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e.
//
// Solidity: event WithdrawalInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *MockL2ERC20BridgeWithdrawalInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MockL2ERC20Bridge.contract.WatchLogs(opts, "WithdrawalInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2ERC20BridgeWithdrawalInitiated)
				if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e.
//
// Solidity: event WithdrawalInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) ParseWithdrawalInitiated(log types.Log) (*MockL2ERC20BridgeWithdrawalInitiated, error) {
	event := new(MockL2ERC20BridgeWithdrawalInitiated)
	if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
