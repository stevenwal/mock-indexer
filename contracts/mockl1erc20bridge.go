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

// MockL1ERC20BridgeMetaData contains all meta data concerning the MockL1ERC20Bridge contract.
var MockL1ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"}],\"name\":\"MockDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSAGE_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Bridge\",\"type\":\"address\"}],\"name\":\"setl2TokenBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061097b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063838b25201161005b578063838b25201461010557806391c49bf814610118578063ae1f6aaf1461013d578063ecc704281461015057600080fd5b8063185f8f1a1461008d5780633f827a5a146100a257806342599311146100c257806358a997f6146100f2575b600080fd5b6100a061009b3660046105d5565b610166565b005b6100aa600181565b60405161ffff90911681526020015b60405180910390f35b6100a06100d036600461066c565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6100a06101003660046106a2565b610215565b6100a0610113366004610721565b61022d565b6001546001600160a01b03165b6040516001600160a01b0390911681526020016100b9565b600154610125906001600160a01b031681565b610158610245565b6040519081526020016100b9565b61017a6001600160a01b038916868661025a565b6040516001815281907fdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b9060200160405180910390a2856001600160a01b0316876001600160a01b0316896001600160a01b03167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b38888888860405161020394939291906107da565b60405180910390a45050505050505050565b61022586863387600087876102dd565b505050505050565b61023c878787878787876102dd565b50505050505050565b6000546001600160f01b0316600160f01b1790565b600060405163a9059cbb60e01b8152836004820152826024820152602060006044836000895af13d15601f3d11600160005114161716915050806102d75760405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b60448201526064015b60405180910390fd5b50505050565b6102f26001600160a01b0388163330876104e6565b600063ec15143160e01b88883389898888604051602401610319979695949392919061080c565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915290506000610357610245565b60015460405161037f929130916001600160a01b03909116906000908a9088906024016108aa565b60408051601f198184030181529190526020810180516001600160e01b031663d764ad0b60e01b178152905190206001549091506001600160a01b03167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a30846103e7610245565b60006040516103f994939291906108fd565b60405180910390a2336001600160a01b0316886001600160a01b03168a6001600160a01b03167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d03968a8a898960405161045494939291906107da565b60405180910390a4336001600160a01b0316886001600160a01b03168a6001600160a01b03167f6feb6160b2038e98bd583f08016fafafa6de12d546e0e5834d7d80a128085b1a8a8a8989886040516104b1959493929190610934565b60405180910390a45050600080546001600160f01b03808216600101166001600160f01b031990911617905550505050505050565b60006040516323b872dd60e01b81528460048201528360248201528260448201526020600060648360008a5af13d15601f3d11600160005114161716915050806105695760405162461bcd60e51b81526020600482015260146024820152731514905394d1915497d19493d357d1905253115160621b60448201526064016102ce565b5050505050565b80356001600160a01b038116811461058757600080fd5b919050565b60008083601f84011261059e57600080fd5b50813567ffffffffffffffff8111156105b657600080fd5b6020830191508360208285010111156105ce57600080fd5b9250929050565b60008060008060008060008060e0898b0312156105f157600080fd5b6105fa89610570565b975061060860208a01610570565b965061061660408a01610570565b955061062460608a01610570565b94506080890135935060a089013567ffffffffffffffff81111561064757600080fd5b6106538b828c0161058c565b999c989b50969995989497949560c00135949350505050565b60006020828403121561067e57600080fd5b61068782610570565b9392505050565b803563ffffffff8116811461058757600080fd5b60008060008060008060a087890312156106bb57600080fd5b6106c487610570565b95506106d260208801610570565b9450604087013593506106e76060880161068e565b9250608087013567ffffffffffffffff81111561070357600080fd5b61070f89828a0161058c565b979a9699509497509295939492505050565b600080600080600080600060c0888a03121561073c57600080fd5b61074588610570565b965061075360208901610570565b955061076160408901610570565b9450606088013593506107766080890161068e565b925060a088013567ffffffffffffffff81111561079257600080fd5b61079e8a828b0161058c565b989b979a50959850939692959293505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60018060a01b03851681528360208201526060604082015260006108026060830184866107b1565b9695505050505050565b6001600160a01b03888116825287811660208301528681166040830152851660608201526080810184905260c060a0820181905260009061085090830184866107b1565b9998505050505050505050565b6000815180845260005b8181101561088357602081850181015186830182015201610867565b81811115610895576000602083870101525b50601f01601f19169290920160200192915050565b8681526001600160a01b0386811660208301528516604082015260ff8416606082015263ffffffff8316608082015260c060a082018190526000906108f19083018461085d565b98975050505050505050565b6001600160a01b03851681526080602082018190526000906109219083018661085d565b6040830194909452506060015292915050565b60018060a01b038616815284602082015260806040820152600061095c6080830185876107b1565b9050826060830152969550505050505056fea164736f6c634300080f000a",
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
	parsed, err := MockL1ERC20BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeCaller) MESSAGEVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MockL1ERC20Bridge.contract.Call(opts, &out, "MESSAGE_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) MESSAGEVERSION() (uint16, error) {
	return _MockL1ERC20Bridge.Contract.MESSAGEVERSION(&_MockL1ERC20Bridge.CallOpts)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeCallerSession) MESSAGEVERSION() (uint16, error) {
	return _MockL1ERC20Bridge.Contract.MESSAGEVERSION(&_MockL1ERC20Bridge.CallOpts)
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

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockL1ERC20Bridge.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) MessageNonce() (*big.Int, error) {
	return _MockL1ERC20Bridge.Contract.MessageNonce(&_MockL1ERC20Bridge.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeCallerSession) MessageNonce() (*big.Int, error) {
	return _MockL1ERC20Bridge.Contract.MessageNonce(&_MockL1ERC20Bridge.CallOpts)
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
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _gasLimit, bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _gasLimit uint32, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, _gasLimit, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _gasLimit, bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _gasLimit uint32, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.DepositERC20To(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _gasLimit, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _gasLimit, bytes _data) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _gasLimit uint32, _data []byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.DepositERC20To(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _gasLimit, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0x185f8f1a.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data, bytes32 withdrawalHash) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte, withdrawalHash [32]byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _data, withdrawalHash)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0x185f8f1a.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data, bytes32 withdrawalHash) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte, withdrawalHash [32]byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.FinalizeERC20Withdrawal(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data, withdrawalHash)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0x185f8f1a.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data, bytes32 withdrawalHash) returns()
func (_MockL1ERC20Bridge *MockL1ERC20BridgeTransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte, withdrawalHash [32]byte) (*types.Transaction, error) {
	return _MockL1ERC20Bridge.Contract.FinalizeERC20Withdrawal(&_MockL1ERC20Bridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data, withdrawalHash)
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

// MockL1ERC20BridgeMockDepositInitiatedIterator is returned from FilterMockDepositInitiated and is used to iterate over the raw logs and unpacked data for MockDepositInitiated events raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeMockDepositInitiatedIterator struct {
	Event *MockL1ERC20BridgeMockDepositInitiated // Event containing the contract specifics and raw log

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
func (it *MockL1ERC20BridgeMockDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1ERC20BridgeMockDepositInitiated)
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
		it.Event = new(MockL1ERC20BridgeMockDepositInitiated)
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
func (it *MockL1ERC20BridgeMockDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1ERC20BridgeMockDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1ERC20BridgeMockDepositInitiated represents a MockDepositInitiated event raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeMockDepositInitiated struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Message [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMockDepositInitiated is a free log retrieval operation binding the contract event 0x6feb6160b2038e98bd583f08016fafafa6de12d546e0e5834d7d80a128085b1a.
//
// Solidity: event MockDepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data, bytes32 _message)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) FilterMockDepositInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*MockL1ERC20BridgeMockDepositInitiatedIterator, error) {

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

	logs, sub, err := _MockL1ERC20Bridge.contract.FilterLogs(opts, "MockDepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &MockL1ERC20BridgeMockDepositInitiatedIterator{contract: _MockL1ERC20Bridge.contract, event: "MockDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchMockDepositInitiated is a free log subscription operation binding the contract event 0x6feb6160b2038e98bd583f08016fafafa6de12d546e0e5834d7d80a128085b1a.
//
// Solidity: event MockDepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data, bytes32 _message)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) WatchMockDepositInitiated(opts *bind.WatchOpts, sink chan<- *MockL1ERC20BridgeMockDepositInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MockL1ERC20Bridge.contract.WatchLogs(opts, "MockDepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1ERC20BridgeMockDepositInitiated)
				if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "MockDepositInitiated", log); err != nil {
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

// ParseMockDepositInitiated is a log parse operation binding the contract event 0x6feb6160b2038e98bd583f08016fafafa6de12d546e0e5834d7d80a128085b1a.
//
// Solidity: event MockDepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data, bytes32 _message)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) ParseMockDepositInitiated(log types.Log) (*MockL1ERC20BridgeMockDepositInitiated, error) {
	event := new(MockL1ERC20BridgeMockDepositInitiated)
	if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "MockDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL1ERC20BridgeSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeSentMessageIterator struct {
	Event *MockL1ERC20BridgeSentMessage // Event containing the contract specifics and raw log

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
func (it *MockL1ERC20BridgeSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1ERC20BridgeSentMessage)
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
		it.Event = new(MockL1ERC20BridgeSentMessage)
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
func (it *MockL1ERC20BridgeSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1ERC20BridgeSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1ERC20BridgeSentMessage represents a SentMessage event raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeSentMessage struct {
	Target       common.Address
	Sender       common.Address
	Message      []byte
	MessageNonce *big.Int
	GasLimit     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) FilterSentMessage(opts *bind.FilterOpts, target []common.Address) (*MockL1ERC20BridgeSentMessageIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _MockL1ERC20Bridge.contract.FilterLogs(opts, "SentMessage", targetRule)
	if err != nil {
		return nil, err
	}
	return &MockL1ERC20BridgeSentMessageIterator{contract: _MockL1ERC20Bridge.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *MockL1ERC20BridgeSentMessage, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _MockL1ERC20Bridge.contract.WatchLogs(opts, "SentMessage", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1ERC20BridgeSentMessage)
				if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "SentMessage", log); err != nil {
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

// ParseSentMessage is a log parse operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) ParseSentMessage(log types.Log) (*MockL1ERC20BridgeSentMessage, error) {
	event := new(MockL1ERC20BridgeSentMessage)
	if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL1ERC20BridgeWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeWithdrawalFinalizedIterator struct {
	Event *MockL1ERC20BridgeWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *MockL1ERC20BridgeWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1ERC20BridgeWithdrawalFinalized)
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
		it.Event = new(MockL1ERC20BridgeWithdrawalFinalized)
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
func (it *MockL1ERC20BridgeWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1ERC20BridgeWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1ERC20BridgeWithdrawalFinalized represents a WithdrawalFinalized event raised by the MockL1ERC20Bridge contract.
type MockL1ERC20BridgeWithdrawalFinalized struct {
	WithdrawalHash [32]byte
	Success        bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, withdrawalHash [][32]byte) (*MockL1ERC20BridgeWithdrawalFinalizedIterator, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _MockL1ERC20Bridge.contract.FilterLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return &MockL1ERC20BridgeWithdrawalFinalizedIterator{contract: _MockL1ERC20Bridge.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *MockL1ERC20BridgeWithdrawalFinalized, withdrawalHash [][32]byte) (event.Subscription, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _MockL1ERC20Bridge.contract.WatchLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1ERC20BridgeWithdrawalFinalized)
				if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_MockL1ERC20Bridge *MockL1ERC20BridgeFilterer) ParseWithdrawalFinalized(log types.Log) (*MockL1ERC20BridgeWithdrawalFinalized, error) {
	event := new(MockL1ERC20BridgeWithdrawalFinalized)
	if err := _MockL1ERC20Bridge.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
