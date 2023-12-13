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

// ISocketFees is an auto generated low-level Go binding around an user-defined struct.
type ISocketFees struct {
	TransmissionFees *big.Int
	ExecutionFee     *big.Int
	SwitchboardFees  *big.Int
}

// MockL1SocketVaultMetaData contains all meta data concerning the MockL1SocketVault contract.
var MockL1SocketVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"localChainSlug\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localPlug\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstChainSlug\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstPlug\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minMsgGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionParams\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transmissionParams\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"transmissionFees\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"executionFee\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"switchboardFees\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structISocket.Fees\",\"name\":\"fees\",\"type\":\"tuple\"}],\"name\":\"MessageOutbound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"}],\"name\":\"SocketMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"TokensDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAmount\",\"type\":\"uint256\"}],\"name\":\"TokensUnlocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositToAppChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_msgId\",\"type\":\"bytes32\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161054338038061054383398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b6104b0806100936000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80637e44144b14610051578063864f6a7a14610066578063ecc7042814610079578063fc0c546a1461008f575b600080fd5b61006461005f36600461042a565b6100ba565b005b61006461007436600461045d565b610152565b6000546040519081526020015b60405180910390f35b6001546100a2906001600160a01b031681565b6040516001600160a01b039091168152602001610086565b6001546100d1906001600160a01b03168484610301565b604080513081526001600160a01b03851660208201529081018390527fece684e11f49f06d351439e63189ad1703238b8040d90cf994901ca2b3da8d449060600160405180910390a16040518181527fdc29884a71d2bb98d3c53dc09718be05c7bfd142b7773a5c5cf2517629290ac09060200160405180910390a1505050565b60015461016a906001600160a01b0316333086610384565b604080513081526001600160a01b03861660208201819052818301526060810185905290517f9474e087d8a0e83962ac44e292b4aba027426203ea66adfb1dd9f65795ff599a9181900360800190a17f45d70e73f9a26cea18d4009d1ad25ffbdc4114ce9d8ceb15959ae9eadb0a3a2c6000806000806101e960005490565b604080516060808201835260008083526020808401828152848601838152865163ffffffff9c8d1681526001600160a01b039b8c169381019390935298909a16818601529590971690850152608084019290925260a0830185905260c0830185905260e083018590526101806101008401819052830194909452516fffffffffffffffffffffffffffffffff90811661012083015293518416610140820152905190921661016083015251908190036101a00190a17fc419bcb5e6804701c0d077263cd1a3a021af5c35e88c64784f6981388248217d84846102ca60005490565b604080516001600160a01b03909416845260208401929092529082015260600160405180910390a150506000805460010190555050565b600060405163a9059cbb60e01b8152836004820152826024820152602060006044836000895af13d15601f3d116001600051141617169150508061037e5760405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b60448201526064015b60405180910390fd5b50505050565b60006040516323b872dd60e01b81528460048201528360248201528260448201526020600060648360008a5af13d15601f3d11600160005114161716915050806104075760405162461bcd60e51b81526020600482015260146024820152731514905394d1915497d19493d357d1905253115160621b6044820152606401610375565b5050505050565b80356001600160a01b038116811461042557600080fd5b919050565b60008060006060848603121561043f57600080fd5b6104488461040e565b95602085013595506040909401359392505050565b6000806000806080858703121561047357600080fd5b61047c8561040e565b935060208501359250604085013591506104986060860161040e565b90509295919450925056fea164736f6c634300080f000a",
}

// MockL1SocketVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use MockL1SocketVaultMetaData.ABI instead.
var MockL1SocketVaultABI = MockL1SocketVaultMetaData.ABI

// MockL1SocketVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockL1SocketVaultMetaData.Bin instead.
var MockL1SocketVaultBin = MockL1SocketVaultMetaData.Bin

// DeployMockL1SocketVault deploys a new Ethereum contract, binding an instance of MockL1SocketVault to it.
func DeployMockL1SocketVault(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *MockL1SocketVault, error) {
	parsed, err := MockL1SocketVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockL1SocketVaultBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockL1SocketVault{MockL1SocketVaultCaller: MockL1SocketVaultCaller{contract: contract}, MockL1SocketVaultTransactor: MockL1SocketVaultTransactor{contract: contract}, MockL1SocketVaultFilterer: MockL1SocketVaultFilterer{contract: contract}}, nil
}

// MockL1SocketVault is an auto generated Go binding around an Ethereum contract.
type MockL1SocketVault struct {
	MockL1SocketVaultCaller     // Read-only binding to the contract
	MockL1SocketVaultTransactor // Write-only binding to the contract
	MockL1SocketVaultFilterer   // Log filterer for contract events
}

// MockL1SocketVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockL1SocketVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL1SocketVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockL1SocketVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL1SocketVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockL1SocketVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockL1SocketVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockL1SocketVaultSession struct {
	Contract     *MockL1SocketVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MockL1SocketVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockL1SocketVaultCallerSession struct {
	Contract *MockL1SocketVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MockL1SocketVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockL1SocketVaultTransactorSession struct {
	Contract     *MockL1SocketVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MockL1SocketVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockL1SocketVaultRaw struct {
	Contract *MockL1SocketVault // Generic contract binding to access the raw methods on
}

// MockL1SocketVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockL1SocketVaultCallerRaw struct {
	Contract *MockL1SocketVaultCaller // Generic read-only contract binding to access the raw methods on
}

// MockL1SocketVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockL1SocketVaultTransactorRaw struct {
	Contract *MockL1SocketVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockL1SocketVault creates a new instance of MockL1SocketVault, bound to a specific deployed contract.
func NewMockL1SocketVault(address common.Address, backend bind.ContractBackend) (*MockL1SocketVault, error) {
	contract, err := bindMockL1SocketVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockL1SocketVault{MockL1SocketVaultCaller: MockL1SocketVaultCaller{contract: contract}, MockL1SocketVaultTransactor: MockL1SocketVaultTransactor{contract: contract}, MockL1SocketVaultFilterer: MockL1SocketVaultFilterer{contract: contract}}, nil
}

// NewMockL1SocketVaultCaller creates a new read-only instance of MockL1SocketVault, bound to a specific deployed contract.
func NewMockL1SocketVaultCaller(address common.Address, caller bind.ContractCaller) (*MockL1SocketVaultCaller, error) {
	contract, err := bindMockL1SocketVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockL1SocketVaultCaller{contract: contract}, nil
}

// NewMockL1SocketVaultTransactor creates a new write-only instance of MockL1SocketVault, bound to a specific deployed contract.
func NewMockL1SocketVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*MockL1SocketVaultTransactor, error) {
	contract, err := bindMockL1SocketVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockL1SocketVaultTransactor{contract: contract}, nil
}

// NewMockL1SocketVaultFilterer creates a new log filterer instance of MockL1SocketVault, bound to a specific deployed contract.
func NewMockL1SocketVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*MockL1SocketVaultFilterer, error) {
	contract, err := bindMockL1SocketVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockL1SocketVaultFilterer{contract: contract}, nil
}

// bindMockL1SocketVault binds a generic wrapper to an already deployed contract.
func bindMockL1SocketVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockL1SocketVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockL1SocketVault *MockL1SocketVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL1SocketVault.Contract.MockL1SocketVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockL1SocketVault *MockL1SocketVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL1SocketVault.Contract.MockL1SocketVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockL1SocketVault *MockL1SocketVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL1SocketVault.Contract.MockL1SocketVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockL1SocketVault *MockL1SocketVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockL1SocketVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockL1SocketVault *MockL1SocketVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockL1SocketVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockL1SocketVault *MockL1SocketVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockL1SocketVault.Contract.contract.Transact(opts, method, params...)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(bytes32)
func (_MockL1SocketVault *MockL1SocketVaultCaller) MessageNonce(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockL1SocketVault.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(bytes32)
func (_MockL1SocketVault *MockL1SocketVaultSession) MessageNonce() ([32]byte, error) {
	return _MockL1SocketVault.Contract.MessageNonce(&_MockL1SocketVault.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(bytes32)
func (_MockL1SocketVault *MockL1SocketVaultCallerSession) MessageNonce() ([32]byte, error) {
	return _MockL1SocketVault.Contract.MessageNonce(&_MockL1SocketVault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MockL1SocketVault *MockL1SocketVaultCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockL1SocketVault.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MockL1SocketVault *MockL1SocketVaultSession) Token() (common.Address, error) {
	return _MockL1SocketVault.Contract.Token(&_MockL1SocketVault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MockL1SocketVault *MockL1SocketVaultCallerSession) Token() (common.Address, error) {
	return _MockL1SocketVault.Contract.Token(&_MockL1SocketVault.CallOpts)
}

// DepositToAppChain is a paid mutator transaction binding the contract method 0x864f6a7a.
//
// Solidity: function depositToAppChain(address receiver_, uint256 amount_, uint256 , address ) returns()
func (_MockL1SocketVault *MockL1SocketVaultTransactor) DepositToAppChain(opts *bind.TransactOpts, receiver_ common.Address, amount_ *big.Int, arg2 *big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _MockL1SocketVault.contract.Transact(opts, "depositToAppChain", receiver_, amount_, arg2, arg3)
}

// DepositToAppChain is a paid mutator transaction binding the contract method 0x864f6a7a.
//
// Solidity: function depositToAppChain(address receiver_, uint256 amount_, uint256 , address ) returns()
func (_MockL1SocketVault *MockL1SocketVaultSession) DepositToAppChain(receiver_ common.Address, amount_ *big.Int, arg2 *big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _MockL1SocketVault.Contract.DepositToAppChain(&_MockL1SocketVault.TransactOpts, receiver_, amount_, arg2, arg3)
}

// DepositToAppChain is a paid mutator transaction binding the contract method 0x864f6a7a.
//
// Solidity: function depositToAppChain(address receiver_, uint256 amount_, uint256 , address ) returns()
func (_MockL1SocketVault *MockL1SocketVaultTransactorSession) DepositToAppChain(receiver_ common.Address, amount_ *big.Int, arg2 *big.Int, arg3 common.Address) (*types.Transaction, error) {
	return _MockL1SocketVault.Contract.DepositToAppChain(&_MockL1SocketVault.TransactOpts, receiver_, amount_, arg2, arg3)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0x7e44144b.
//
// Solidity: function finalizeERC20Withdrawal(address _to, uint256 _amount, bytes32 _msgId) returns()
func (_MockL1SocketVault *MockL1SocketVaultTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _msgId [32]byte) (*types.Transaction, error) {
	return _MockL1SocketVault.contract.Transact(opts, "finalizeERC20Withdrawal", _to, _amount, _msgId)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0x7e44144b.
//
// Solidity: function finalizeERC20Withdrawal(address _to, uint256 _amount, bytes32 _msgId) returns()
func (_MockL1SocketVault *MockL1SocketVaultSession) FinalizeERC20Withdrawal(_to common.Address, _amount *big.Int, _msgId [32]byte) (*types.Transaction, error) {
	return _MockL1SocketVault.Contract.FinalizeERC20Withdrawal(&_MockL1SocketVault.TransactOpts, _to, _amount, _msgId)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0x7e44144b.
//
// Solidity: function finalizeERC20Withdrawal(address _to, uint256 _amount, bytes32 _msgId) returns()
func (_MockL1SocketVault *MockL1SocketVaultTransactorSession) FinalizeERC20Withdrawal(_to common.Address, _amount *big.Int, _msgId [32]byte) (*types.Transaction, error) {
	return _MockL1SocketVault.Contract.FinalizeERC20Withdrawal(&_MockL1SocketVault.TransactOpts, _to, _amount, _msgId)
}

// MockL1SocketVaultExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the MockL1SocketVault contract.
type MockL1SocketVaultExecutionSuccessIterator struct {
	Event *MockL1SocketVaultExecutionSuccess // Event containing the contract specifics and raw log

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
func (it *MockL1SocketVaultExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1SocketVaultExecutionSuccess)
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
		it.Event = new(MockL1SocketVaultExecutionSuccess)
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
func (it *MockL1SocketVaultExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1SocketVaultExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1SocketVaultExecutionSuccess represents a ExecutionSuccess event raised by the MockL1SocketVault contract.
type MockL1SocketVaultExecutionSuccess struct {
	MsgId [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0xdc29884a71d2bb98d3c53dc09718be05c7bfd142b7773a5c5cf2517629290ac0.
//
// Solidity: event ExecutionSuccess(bytes32 msgId)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) FilterExecutionSuccess(opts *bind.FilterOpts) (*MockL1SocketVaultExecutionSuccessIterator, error) {

	logs, sub, err := _MockL1SocketVault.contract.FilterLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return &MockL1SocketVaultExecutionSuccessIterator{contract: _MockL1SocketVault.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0xdc29884a71d2bb98d3c53dc09718be05c7bfd142b7773a5c5cf2517629290ac0.
//
// Solidity: event ExecutionSuccess(bytes32 msgId)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *MockL1SocketVaultExecutionSuccess) (event.Subscription, error) {

	logs, sub, err := _MockL1SocketVault.contract.WatchLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1SocketVaultExecutionSuccess)
				if err := _MockL1SocketVault.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
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

// ParseExecutionSuccess is a log parse operation binding the contract event 0xdc29884a71d2bb98d3c53dc09718be05c7bfd142b7773a5c5cf2517629290ac0.
//
// Solidity: event ExecutionSuccess(bytes32 msgId)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) ParseExecutionSuccess(log types.Log) (*MockL1SocketVaultExecutionSuccess, error) {
	event := new(MockL1SocketVaultExecutionSuccess)
	if err := _MockL1SocketVault.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL1SocketVaultMessageOutboundIterator is returned from FilterMessageOutbound and is used to iterate over the raw logs and unpacked data for MessageOutbound events raised by the MockL1SocketVault contract.
type MockL1SocketVaultMessageOutboundIterator struct {
	Event *MockL1SocketVaultMessageOutbound // Event containing the contract specifics and raw log

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
func (it *MockL1SocketVaultMessageOutboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1SocketVaultMessageOutbound)
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
		it.Event = new(MockL1SocketVaultMessageOutbound)
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
func (it *MockL1SocketVaultMessageOutboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1SocketVaultMessageOutboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1SocketVaultMessageOutbound represents a MessageOutbound event raised by the MockL1SocketVault contract.
type MockL1SocketVaultMessageOutbound struct {
	LocalChainSlug     uint32
	LocalPlug          common.Address
	DstChainSlug       uint32
	DstPlug            common.Address
	MsgId              [32]byte
	MinMsgGasLimit     *big.Int
	ExecutionParams    [32]byte
	TransmissionParams [32]byte
	Payload            []byte
	Fees               ISocketFees
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageOutbound is a free log retrieval operation binding the contract event 0x45d70e73f9a26cea18d4009d1ad25ffbdc4114ce9d8ceb15959ae9eadb0a3a2c.
//
// Solidity: event MessageOutbound(uint32 localChainSlug, address localPlug, uint32 dstChainSlug, address dstPlug, bytes32 msgId, uint256 minMsgGasLimit, bytes32 executionParams, bytes32 transmissionParams, bytes payload, (uint128,uint128,uint128) fees)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) FilterMessageOutbound(opts *bind.FilterOpts) (*MockL1SocketVaultMessageOutboundIterator, error) {

	logs, sub, err := _MockL1SocketVault.contract.FilterLogs(opts, "MessageOutbound")
	if err != nil {
		return nil, err
	}
	return &MockL1SocketVaultMessageOutboundIterator{contract: _MockL1SocketVault.contract, event: "MessageOutbound", logs: logs, sub: sub}, nil
}

// WatchMessageOutbound is a free log subscription operation binding the contract event 0x45d70e73f9a26cea18d4009d1ad25ffbdc4114ce9d8ceb15959ae9eadb0a3a2c.
//
// Solidity: event MessageOutbound(uint32 localChainSlug, address localPlug, uint32 dstChainSlug, address dstPlug, bytes32 msgId, uint256 minMsgGasLimit, bytes32 executionParams, bytes32 transmissionParams, bytes payload, (uint128,uint128,uint128) fees)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) WatchMessageOutbound(opts *bind.WatchOpts, sink chan<- *MockL1SocketVaultMessageOutbound) (event.Subscription, error) {

	logs, sub, err := _MockL1SocketVault.contract.WatchLogs(opts, "MessageOutbound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1SocketVaultMessageOutbound)
				if err := _MockL1SocketVault.contract.UnpackLog(event, "MessageOutbound", log); err != nil {
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

// ParseMessageOutbound is a log parse operation binding the contract event 0x45d70e73f9a26cea18d4009d1ad25ffbdc4114ce9d8ceb15959ae9eadb0a3a2c.
//
// Solidity: event MessageOutbound(uint32 localChainSlug, address localPlug, uint32 dstChainSlug, address dstPlug, bytes32 msgId, uint256 minMsgGasLimit, bytes32 executionParams, bytes32 transmissionParams, bytes payload, (uint128,uint128,uint128) fees)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) ParseMessageOutbound(log types.Log) (*MockL1SocketVaultMessageOutbound, error) {
	event := new(MockL1SocketVaultMessageOutbound)
	if err := _MockL1SocketVault.contract.UnpackLog(event, "MessageOutbound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL1SocketVaultSocketMessageIterator is returned from FilterSocketMessage and is used to iterate over the raw logs and unpacked data for SocketMessage events raised by the MockL1SocketVault contract.
type MockL1SocketVaultSocketMessageIterator struct {
	Event *MockL1SocketVaultSocketMessage // Event containing the contract specifics and raw log

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
func (it *MockL1SocketVaultSocketMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1SocketVaultSocketMessage)
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
		it.Event = new(MockL1SocketVaultSocketMessage)
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
func (it *MockL1SocketVaultSocketMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1SocketVaultSocketMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1SocketVaultSocketMessage represents a SocketMessage event raised by the MockL1SocketVault contract.
type MockL1SocketVaultSocketMessage struct {
	To     common.Address
	Amount *big.Int
	MsgId  [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSocketMessage is a free log retrieval operation binding the contract event 0xc419bcb5e6804701c0d077263cd1a3a021af5c35e88c64784f6981388248217d.
//
// Solidity: event SocketMessage(address to, uint256 amount, bytes32 msgId)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) FilterSocketMessage(opts *bind.FilterOpts) (*MockL1SocketVaultSocketMessageIterator, error) {

	logs, sub, err := _MockL1SocketVault.contract.FilterLogs(opts, "SocketMessage")
	if err != nil {
		return nil, err
	}
	return &MockL1SocketVaultSocketMessageIterator{contract: _MockL1SocketVault.contract, event: "SocketMessage", logs: logs, sub: sub}, nil
}

// WatchSocketMessage is a free log subscription operation binding the contract event 0xc419bcb5e6804701c0d077263cd1a3a021af5c35e88c64784f6981388248217d.
//
// Solidity: event SocketMessage(address to, uint256 amount, bytes32 msgId)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) WatchSocketMessage(opts *bind.WatchOpts, sink chan<- *MockL1SocketVaultSocketMessage) (event.Subscription, error) {

	logs, sub, err := _MockL1SocketVault.contract.WatchLogs(opts, "SocketMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1SocketVaultSocketMessage)
				if err := _MockL1SocketVault.contract.UnpackLog(event, "SocketMessage", log); err != nil {
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

// ParseSocketMessage is a log parse operation binding the contract event 0xc419bcb5e6804701c0d077263cd1a3a021af5c35e88c64784f6981388248217d.
//
// Solidity: event SocketMessage(address to, uint256 amount, bytes32 msgId)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) ParseSocketMessage(log types.Log) (*MockL1SocketVaultSocketMessage, error) {
	event := new(MockL1SocketVaultSocketMessage)
	if err := _MockL1SocketVault.contract.UnpackLog(event, "SocketMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL1SocketVaultTokensDepositedIterator is returned from FilterTokensDeposited and is used to iterate over the raw logs and unpacked data for TokensDeposited events raised by the MockL1SocketVault contract.
type MockL1SocketVaultTokensDepositedIterator struct {
	Event *MockL1SocketVaultTokensDeposited // Event containing the contract specifics and raw log

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
func (it *MockL1SocketVaultTokensDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1SocketVaultTokensDeposited)
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
		it.Event = new(MockL1SocketVaultTokensDeposited)
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
func (it *MockL1SocketVaultTokensDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1SocketVaultTokensDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1SocketVaultTokensDeposited represents a TokensDeposited event raised by the MockL1SocketVault contract.
type MockL1SocketVaultTokensDeposited struct {
	Connector     common.Address
	Depositor     common.Address
	Receiver      common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensDeposited is a free log retrieval operation binding the contract event 0x9474e087d8a0e83962ac44e292b4aba027426203ea66adfb1dd9f65795ff599a.
//
// Solidity: event TokensDeposited(address connector, address depositor, address receiver, uint256 depositAmount)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) FilterTokensDeposited(opts *bind.FilterOpts) (*MockL1SocketVaultTokensDepositedIterator, error) {

	logs, sub, err := _MockL1SocketVault.contract.FilterLogs(opts, "TokensDeposited")
	if err != nil {
		return nil, err
	}
	return &MockL1SocketVaultTokensDepositedIterator{contract: _MockL1SocketVault.contract, event: "TokensDeposited", logs: logs, sub: sub}, nil
}

// WatchTokensDeposited is a free log subscription operation binding the contract event 0x9474e087d8a0e83962ac44e292b4aba027426203ea66adfb1dd9f65795ff599a.
//
// Solidity: event TokensDeposited(address connector, address depositor, address receiver, uint256 depositAmount)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) WatchTokensDeposited(opts *bind.WatchOpts, sink chan<- *MockL1SocketVaultTokensDeposited) (event.Subscription, error) {

	logs, sub, err := _MockL1SocketVault.contract.WatchLogs(opts, "TokensDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1SocketVaultTokensDeposited)
				if err := _MockL1SocketVault.contract.UnpackLog(event, "TokensDeposited", log); err != nil {
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

// ParseTokensDeposited is a log parse operation binding the contract event 0x9474e087d8a0e83962ac44e292b4aba027426203ea66adfb1dd9f65795ff599a.
//
// Solidity: event TokensDeposited(address connector, address depositor, address receiver, uint256 depositAmount)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) ParseTokensDeposited(log types.Log) (*MockL1SocketVaultTokensDeposited, error) {
	event := new(MockL1SocketVaultTokensDeposited)
	if err := _MockL1SocketVault.contract.UnpackLog(event, "TokensDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockL1SocketVaultTokensUnlockedIterator is returned from FilterTokensUnlocked and is used to iterate over the raw logs and unpacked data for TokensUnlocked events raised by the MockL1SocketVault contract.
type MockL1SocketVaultTokensUnlockedIterator struct {
	Event *MockL1SocketVaultTokensUnlocked // Event containing the contract specifics and raw log

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
func (it *MockL1SocketVaultTokensUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockL1SocketVaultTokensUnlocked)
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
		it.Event = new(MockL1SocketVaultTokensUnlocked)
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
func (it *MockL1SocketVaultTokensUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockL1SocketVaultTokensUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockL1SocketVaultTokensUnlocked represents a TokensUnlocked event raised by the MockL1SocketVault contract.
type MockL1SocketVaultTokensUnlocked struct {
	Connector      common.Address
	Receiver       common.Address
	UnlockedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokensUnlocked is a free log retrieval operation binding the contract event 0xece684e11f49f06d351439e63189ad1703238b8040d90cf994901ca2b3da8d44.
//
// Solidity: event TokensUnlocked(address connector, address receiver, uint256 unlockedAmount)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) FilterTokensUnlocked(opts *bind.FilterOpts) (*MockL1SocketVaultTokensUnlockedIterator, error) {

	logs, sub, err := _MockL1SocketVault.contract.FilterLogs(opts, "TokensUnlocked")
	if err != nil {
		return nil, err
	}
	return &MockL1SocketVaultTokensUnlockedIterator{contract: _MockL1SocketVault.contract, event: "TokensUnlocked", logs: logs, sub: sub}, nil
}

// WatchTokensUnlocked is a free log subscription operation binding the contract event 0xece684e11f49f06d351439e63189ad1703238b8040d90cf994901ca2b3da8d44.
//
// Solidity: event TokensUnlocked(address connector, address receiver, uint256 unlockedAmount)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) WatchTokensUnlocked(opts *bind.WatchOpts, sink chan<- *MockL1SocketVaultTokensUnlocked) (event.Subscription, error) {

	logs, sub, err := _MockL1SocketVault.contract.WatchLogs(opts, "TokensUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockL1SocketVaultTokensUnlocked)
				if err := _MockL1SocketVault.contract.UnpackLog(event, "TokensUnlocked", log); err != nil {
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

// ParseTokensUnlocked is a log parse operation binding the contract event 0xece684e11f49f06d351439e63189ad1703238b8040d90cf994901ca2b3da8d44.
//
// Solidity: event TokensUnlocked(address connector, address receiver, uint256 unlockedAmount)
func (_MockL1SocketVault *MockL1SocketVaultFilterer) ParseTokensUnlocked(log types.Log) (*MockL1SocketVaultTokensUnlocked, error) {
	event := new(MockL1SocketVaultTokensUnlocked)
	if err := _MockL1SocketVault.contract.UnpackLog(event, "TokensUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
