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
	_ = abi.ConvertType
)

// MockL2ERC20BridgeMetaData contains all meta data concerning the MockL2ERC20Bridge contract.
var MockL2ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DepositFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"}],\"name\":\"MessagePassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_withdrawHash\",\"type\":\"bytes32\"}],\"name\":\"MockWithdrawalInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSAGE_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"supportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supported\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516108f53803806108f583398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b60805161085d6100986000396000818160af0152818161017301526103cb015261085d6000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806392b9f8f51161006657806392b9f8f514610130578063969b53da1461016e578063a3a7954814610195578063ec151431146101a8578063ecc70428146101bb57600080fd5b806332b7006d1461009857806336c717c1146100ad5780633f827a5a146100ec5780638597b7c914610107575b600080fd5b6100ab6100a6366004610578565b6101d1565b005b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020015b60405180910390f35b6100f4600181565b60405161ffff90911681526020016100e3565b6100cf6101153660046105e7565b6001602052600090815260409020546001600160a01b031681565b6100ab61013e366004610609565b6001600160a01b0390811660009081526001602052604090208054919092166001600160a01b0319909116179055565b6100cf7f000000000000000000000000000000000000000000000000000000000000000081565b6100ab6101a336600461063c565b6101e7565b6100ab6101b63660046106bb565b6101fe565b6101c36102f0565b6040519081526020016100e3565b6101e085338660008686610305565b5050505050565b6101f686868660008686610305565b505050505050565b6040516340c10f1960e01b81526001600160a01b038681166004830152602482018690528816906340c10f1990604401600060405180830381600087803b15801561024857600080fd5b505af115801561025c573d6000803e3d6000fd5b50506040518392507f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c9150600090a2856001600160a01b0316876001600160a01b0316896001600160a01b03167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89888888886040516102de949392919061077b565b60405180910390a45050505050505050565b6000546001600160f01b0316600160f01b1790565b604051632770a7eb60e21b8152336004820152602481018590526001600160a01b03871690639dc29fac90604401600060405180830381600087803b15801561034d57600080fd5b505af1158015610361573d6000803e3d6000fd5b50505050600063a3a7954860e01b8787876000878760405160240161038b969594939291906107ad565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217825251902090506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016306103f46102f0565b7f02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a274495505460008088888860405161042c9594939291906107f6565b60405180910390a46001600160a01b0387811660008181526001602052604090819020549051339391909116907f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e9061048c908b908b908a908a9061077b565b60405180910390a46001600160a01b0387811660008181526001602052604090819020549051339391909116907fd83a3e216f3103a8a243c34844058cff9b93b3eadd57cf0d7baf683c5618611c906104ee908b908b908a908a908a90610828565b60405180910390a450505050505050565b80356001600160a01b038116811461051657600080fd5b919050565b803563ffffffff8116811461051657600080fd5b60008083601f84011261054157600080fd5b50813567ffffffffffffffff81111561055957600080fd5b60208301915083602082850101111561057157600080fd5b9250929050565b60008060008060006080868803121561059057600080fd5b610599866104ff565b9450602086013593506105ae6040870161051b565b9250606086013567ffffffffffffffff8111156105ca57600080fd5b6105d68882890161052f565b969995985093965092949392505050565b6000602082840312156105f957600080fd5b610602826104ff565b9392505050565b6000806040838503121561061c57600080fd5b610625836104ff565b9150610633602084016104ff565b90509250929050565b60008060008060008060a0878903121561065557600080fd5b61065e876104ff565b955061066c602088016104ff565b9450604087013593506106816060880161051b565b9250608087013567ffffffffffffffff81111561069d57600080fd5b6106a989828a0161052f565b979a9699509497509295939492505050565b60008060008060008060008060e0898b0312156106d757600080fd5b6106e0896104ff565b97506106ee60208a016104ff565b96506106fc60408a016104ff565b955061070a60608a016104ff565b94506080890135935060a089013567ffffffffffffffff81111561072d57600080fd5b6107398b828c0161052f565b999c989b50969995989497949560c00135949350505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60018060a01b03851681528360208201526060604082015260006107a3606083018486610752565b9695505050505050565b6001600160a01b038781168252861660208201526040810185905260ff8416606082015260a0608082018190526000906107ea9083018486610752565b98975050505050505050565b858152846020820152608060408201526000610816608083018587610752565b90508260608301529695505050505050565b60018060a01b038616815284602082015260806040820152600061081660808301858761075256fea164736f6c634300080f000a",
}

// MockL2ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use MockL2ERC20BridgeMetaData.ABI instead.
var MockL2ERC20BridgeABI = MockL2ERC20BridgeMetaData.ABI

// MockL2ERC20BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockL2ERC20BridgeMetaData.Bin instead.
var MockL2ERC20BridgeBin = MockL2ERC20BridgeMetaData.Bin

// DeployMockL2ERC20Bridge deploys a new Ethereum contract, binding an instance of MockL2ERC20Bridge to it.
func DeployMockL2ERC20Bridge(auth *bind.TransactOpts, backend bind.ContractBackend, _l1Bridge common.Address) (common.Address, *types.Transaction, *MockL2ERC20Bridge, error) {
	parsed, err := MockL2ERC20BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockL2ERC20BridgeBin), backend, _l1Bridge)
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
	parsed, err := MockL2ERC20BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCaller) MESSAGEVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MockL2ERC20Bridge.contract.Call(opts, &out, "MESSAGE_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) MESSAGEVERSION() (uint16, error) {
	return _MockL2ERC20Bridge.Contract.MESSAGEVERSION(&_MockL2ERC20Bridge.CallOpts)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCallerSession) MESSAGEVERSION() (uint16, error) {
	return _MockL2ERC20Bridge.Contract.MESSAGEVERSION(&_MockL2ERC20Bridge.CallOpts)
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

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockL2ERC20Bridge.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) MessageNonce() (*big.Int, error) {
	return _MockL2ERC20Bridge.Contract.MessageNonce(&_MockL2ERC20Bridge.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCallerSession) MessageNonce() (*big.Int, error) {
	return _MockL2ERC20Bridge.Contract.MessageNonce(&_MockL2ERC20Bridge.CallOpts)
}

// Supported is a free data retrieval call binding the contract method 0x8597b7c9.
//
// Solidity: function supported(address ) view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCaller) Supported(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MockL2ERC20Bridge.contract.Call(opts, &out, "supported", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Supported is a free data retrieval call binding the contract method 0x8597b7c9.
//
// Solidity: function supported(address ) view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) Supported(arg0 common.Address) (common.Address, error) {
	return _MockL2ERC20Bridge.Contract.Supported(&_MockL2ERC20Bridge.CallOpts, arg0)
}

// Supported is a free data retrieval call binding the contract method 0x8597b7c9.
//
// Solidity: function supported(address ) view returns(address)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeCallerSession) Supported(arg0 common.Address) (common.Address, error) {
	return _MockL2ERC20Bridge.Contract.Supported(&_MockL2ERC20Bridge.CallOpts, arg0)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xec151431.
//
// Solidity: function finalizeDeposit(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data, bytes32 msgHash) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactor) FinalizeDeposit(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte, msgHash [32]byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.contract.Transact(opts, "finalizeDeposit", _l1Token, _l2Token, _from, _to, _amount, _data, msgHash)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xec151431.
//
// Solidity: function finalizeDeposit(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data, bytes32 msgHash) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte, msgHash [32]byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.FinalizeDeposit(&_MockL2ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data, msgHash)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xec151431.
//
// Solidity: function finalizeDeposit(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data, bytes32 msgHash) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactorSession) FinalizeDeposit(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte, msgHash [32]byte) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.FinalizeDeposit(&_MockL2ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data, msgHash)
}

// SupportToken is a paid mutator transaction binding the contract method 0x92b9f8f5.
//
// Solidity: function supportToken(address _l1Token, address _l2Token) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactor) SupportToken(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.contract.Transact(opts, "supportToken", _l1Token, _l2Token)
}

// SupportToken is a paid mutator transaction binding the contract method 0x92b9f8f5.
//
// Solidity: function supportToken(address _l1Token, address _l2Token) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeSession) SupportToken(_l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.SupportToken(&_MockL2ERC20Bridge.TransactOpts, _l1Token, _l2Token)
}

// SupportToken is a paid mutator transaction binding the contract method 0x92b9f8f5.
//
// Solidity: function supportToken(address _l1Token, address _l2Token) returns()
func (_MockL2ERC20Bridge *MockL2ERC20BridgeTransactorSession) SupportToken(_l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _MockL2ERC20Bridge.Contract.SupportToken(&_MockL2ERC20Bridge.TransactOpts, _l1Token, _l2Token)
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

// MockL2ERC20BridgeMessagePassedIterator is returned from FilterMessagePassed and is used to iterate over the raw logs and unpacked data for MessagePassed events raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeMessagePassedIterator struct {
	Event *MockL2ERC20BridgeMessagePassed // Event containing the contract specifics and raw log

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
func (it *MockL2ERC20BridgeMessagePassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2ERC20BridgeMessagePassed)
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
		it.Event = new(MockL2ERC20BridgeMessagePassed)
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
func (it *MockL2ERC20BridgeMessagePassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2ERC20BridgeMessagePassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2ERC20BridgeMessagePassed represents a MessagePassed event raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeMessagePassed struct {
	Nonce          *big.Int
	Sender         common.Address
	Target         common.Address
	Value          *big.Int
	GasLimit       *big.Int
	Data           []byte
	WithdrawalHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessagePassed is a free log retrieval operation binding the contract event 0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) FilterMessagePassed(opts *bind.FilterOpts, nonce []*big.Int, sender []common.Address, target []common.Address) (*MockL2ERC20BridgeMessagePassedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _MockL2ERC20Bridge.contract.FilterLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20BridgeMessagePassedIterator{contract: _MockL2ERC20Bridge.contract, event: "MessagePassed", logs: logs, sub: sub}, nil
}

// WatchMessagePassed is a free log subscription operation binding the contract event 0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) WatchMessagePassed(opts *bind.WatchOpts, sink chan<- *MockL2ERC20BridgeMessagePassed, nonce []*big.Int, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _MockL2ERC20Bridge.contract.WatchLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2ERC20BridgeMessagePassed)
				if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "MessagePassed", log); err != nil {
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

// ParseMessagePassed is a log parse operation binding the contract event 0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) ParseMessagePassed(log types.Log) (*MockL2ERC20BridgeMessagePassed, error) {
	event := new(MockL2ERC20BridgeMessagePassed)
	if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "MessagePassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL2ERC20BridgeMockWithdrawalInitiatedIterator is returned from FilterMockWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for MockWithdrawalInitiated events raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeMockWithdrawalInitiatedIterator struct {
	Event *MockL2ERC20BridgeMockWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *MockL2ERC20BridgeMockWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2ERC20BridgeMockWithdrawalInitiated)
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
		it.Event = new(MockL2ERC20BridgeMockWithdrawalInitiated)
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
func (it *MockL2ERC20BridgeMockWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2ERC20BridgeMockWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2ERC20BridgeMockWithdrawalInitiated represents a MockWithdrawalInitiated event raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeMockWithdrawalInitiated struct {
	L1Token      common.Address
	L2Token      common.Address
	From         common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	WithdrawHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMockWithdrawalInitiated is a free log retrieval operation binding the contract event 0xd83a3e216f3103a8a243c34844058cff9b93b3eadd57cf0d7baf683c5618611c.
//
// Solidity: event MockWithdrawalInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data, bytes32 _withdrawHash)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) FilterMockWithdrawalInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*MockL2ERC20BridgeMockWithdrawalInitiatedIterator, error) {

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

	logs, sub, err := _MockL2ERC20Bridge.contract.FilterLogs(opts, "MockWithdrawalInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20BridgeMockWithdrawalInitiatedIterator{contract: _MockL2ERC20Bridge.contract, event: "MockWithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchMockWithdrawalInitiated is a free log subscription operation binding the contract event 0xd83a3e216f3103a8a243c34844058cff9b93b3eadd57cf0d7baf683c5618611c.
//
// Solidity: event MockWithdrawalInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data, bytes32 _withdrawHash)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) WatchMockWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *MockL2ERC20BridgeMockWithdrawalInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MockL2ERC20Bridge.contract.WatchLogs(opts, "MockWithdrawalInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2ERC20BridgeMockWithdrawalInitiated)
				if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "MockWithdrawalInitiated", log); err != nil {
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

// ParseMockWithdrawalInitiated is a log parse operation binding the contract event 0xd83a3e216f3103a8a243c34844058cff9b93b3eadd57cf0d7baf683c5618611c.
//
// Solidity: event MockWithdrawalInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data, bytes32 _withdrawHash)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) ParseMockWithdrawalInitiated(log types.Log) (*MockL2ERC20BridgeMockWithdrawalInitiated, error) {
	event := new(MockL2ERC20BridgeMockWithdrawalInitiated)
	if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "MockWithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL2ERC20BridgeRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeRelayedMessageIterator struct {
	Event *MockL2ERC20BridgeRelayedMessage // Event containing the contract specifics and raw log

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
func (it *MockL2ERC20BridgeRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL2ERC20BridgeRelayedMessage)
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
		it.Event = new(MockL2ERC20BridgeRelayedMessage)
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
func (it *MockL2ERC20BridgeRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL2ERC20BridgeRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL2ERC20BridgeRelayedMessage represents a RelayedMessage event raised by the MockL2ERC20Bridge contract.
type MockL2ERC20BridgeRelayedMessage struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) FilterRelayedMessage(opts *bind.FilterOpts, msgHash [][32]byte) (*MockL2ERC20BridgeRelayedMessageIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _MockL2ERC20Bridge.contract.FilterLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &MockL2ERC20BridgeRelayedMessageIterator{contract: _MockL2ERC20Bridge.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *MockL2ERC20BridgeRelayedMessage, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _MockL2ERC20Bridge.contract.WatchLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL2ERC20BridgeRelayedMessage)
				if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_MockL2ERC20Bridge *MockL2ERC20BridgeFilterer) ParseRelayedMessage(log types.Log) (*MockL2ERC20BridgeRelayedMessage, error) {
	event := new(MockL2ERC20BridgeRelayedMessage)
	if err := _MockL2ERC20Bridge.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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
