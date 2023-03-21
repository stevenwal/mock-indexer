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

// ExecutorInterfaceCall is an auto generated low-level Go binding around an user-defined struct.
type ExecutorInterfaceCall struct {
	Index  *big.Int
	Fail   bool
	Target common.Address
	Data   []byte
}

// ExecutorInterfaceResult is an auto generated low-level Go binding around an user-defined struct.
type ExecutorInterfaceResult struct {
	Called  bool
	Success bool
	Data    []byte
}

// ExecutorMetaData contains all meta data concerning the Executor contract.
var ExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotKeeper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"UnknownRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"IndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"KeeperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"VersionInitialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorities\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fail\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structExecutorInterface.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"called\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structExecutorInterface.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keepers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"loadIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161140f61004c6000396000818161032d0152818161037601528181610470015281816104b00152610543015261140f6000f3fe6080604052600436106100dd5760003560e01c806352d1902d1161007f5780638da5cb5b116100595780638da5cb5b1461026757806391223d691461029f578063c4d66de8146102cf578063d3057877146102ef57600080fd5b806352d1902d146102055780636cd22eaf1461021a5780638a8ef0b31461023a57600080fd5b80633659cfe6116100bb5780633659cfe61461017257806338ece48b146101925780633bbd64bc146101b25780634f1ef286146101f257600080fd5b806306fdde03146100e257806313af40351461012c5780632986c0e51461014e575b600080fd5b3480156100ee57600080fd5b506101166040518060400160405280600881526020016722bc32b1baba37b960c11b81525081565b6040516101239190610f8d565b60405180910390f35b34801561013857600080fd5b5061014c610147366004610fbc565b61030f565b005b34801561015a57600080fd5b5061016460055481565b604051908152602001610123565b34801561017e57600080fd5b5061014c61018d366004610fbc565b610323565b34801561019e57600080fd5b5061014c6101ad366004610fd7565b610408565b3480156101be57600080fd5b506101e26101cd366004610fbc565b60046020526000908152604090205460ff1681565b6040519015158152602001610123565b61014c610200366004611006565b610466565b34801561021157600080fd5b50610164610536565b34801561022657600080fd5b5061014c6102353660046110d8565b6105e9565b34801561024657600080fd5b5061025a61025536600461110b565b61066c565b6040516101239190611180565b34801561027357600080fd5b50600054610287906001600160a01b031681565b6040516001600160a01b039091168152602001610123565b3480156102ab57600080fd5b506101e26102ba366004610fbc565b60036020526000908152604090205460ff1681565b3480156102db57600080fd5b5061014c6102ea366004610fbc565b610985565b3480156102fb57600080fd5b5061014c61030a3660046110d8565b610a2d565b610317610ab0565b61032081610adc565b50565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036103745760405162461bcd60e51b815260040161036b90611203565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166103bd6000805160206113bc833981519152546001600160a01b031690565b6001600160a01b0316146103e35760405162461bcd60e51b815260040161036b9061124f565b6103ec81610b4b565b6040805160008082526020820190925261032091839190610b53565b610410610ab0565b600160055414610433576040516363df817160e01b815260040160405180910390fd5b600581905560405181907facfc085c9be45d2b3f9e5c09a19d4a95749cc16939519c13e090de3a4cb192c690600090a250565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036104ae5760405162461bcd60e51b815260040161036b90611203565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166104f76000805160206113bc833981519152546001600160a01b031690565b6001600160a01b03161461051d5760405162461bcd60e51b815260040161036b9061124f565b61052682610b4b565b61053282826001610b53565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105d65760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000606482015260840161036b565b506000805160206113bc83398151915290565b6105f1610ab0565b6001600160a01b0382166106185760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff191685151590811790915590519092917fc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc91a35050565b6060610676610cc3565b3360009081526004602052604090205460ff166106a657604051631ea2564f60e31b815260040160405180910390fd5b818067ffffffffffffffff8111156106c0576106c0610ff0565b60405190808252806020026020018201604052801561070d57816020015b604080516060808201835260008083526020830152918101919091528152602001906001900390816106de5790505b50915060005b81811015610978576005548585838181106107305761073061129b565b905060200281019061074291906112b1565b35036109685760008086868481811061075d5761075d61129b565b905060200281019061076f91906112b1565b610780906060810190604001610fbc565b6001600160a01b031687878581811061079b5761079b61129b565b90506020028101906107ad91906112b1565b6107bb9060608101906112d1565b6040516107c992919061131f565b6000604051808303816000865af19150503d8060008114610806576040519150601f19603f3d011682016040523d82523d6000602084013e61080b565b606091505b509150915081158015610850575086868481811061082b5761082b61129b565b905060200281019061083d91906112b1565b61084e90604081019060200161132f565b155b1561088257805160000361087a5760405163dacad91760e01b81526004810184905260240161036b565b805181602001fd5b60018584815181106108965761089661129b565b60209081029190910101519015159052845182908690859081106108bc576108bc61129b565b60200260200101516020019015159081151581525050808584815181106108e5576108e561129b565b6020026020010151604001819052508686848181106109065761090661129b565b905060200281019061091891906112b1565b60405183151581529035907ffa9788a6b2f6e96ab2d85a5b4ed3668e72b91e06c5eec6b3e87afd95fe5368639060200160405180910390a26005805490600061096083611360565b919050555050505b61097181611360565b9050610713565b5050600160025592915050565b60016109918180611379565b600154146109b2576040516302ed543d60e51b815260040160405180910390fd5b806001036109c05760016002555b6109c982610adc565b600160058190556040517facfc085c9be45d2b3f9e5c09a19d4a95749cc16939519c13e090de3a4cb192c690600090a2600181905560405181907f7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d90600090a25050565b610a35610ab0565b6001600160a01b038216610a5c5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260046020526040808220805460ff191685151590811790915590519092917f786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c0091a35050565b6000546001600160a01b03163314610ada576040516282b42960e81b815260040160405180910390fd5b565b6001600160a01b038116610b035760405163d92e233d60e01b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b038316908117825560405190917f4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b91a250565b610320610ab0565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610b8b57610b8683610cec565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610be5575060408051601f3d908101601f19168201909252610be291810190611390565b60015b610c485760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b606482015260840161036b565b6000805160206113bc8339815191528114610cb75760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b606482015260840161036b565b50610b86838383610d88565b600254600114610ce65760405163558a1e0360e11b815260040160405180910390fd5b60028055565b6001600160a01b0381163b610d595760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161036b565b6000805160206113bc83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b610d9183610db3565b600082511180610d9e5750805b15610b8657610dad8383610df3565b50505050565b610dbc81610cec565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610e1883836040518060600160405280602781526020016113dc60279139610e1f565b9392505050565b60606001600160a01b0384163b610e875760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161036b565b600080856001600160a01b031685604051610ea291906113a9565b600060405180830381855af49150503d8060008114610edd576040519150601f19603f3d011682016040523d82523d6000602084013e610ee2565b606091505b5091509150610ef2828286610efc565b9695505050505050565b60608315610f0b575081610e18565b825115610f1b5782518084602001fd5b8160405162461bcd60e51b815260040161036b9190610f8d565b60005b83811015610f50578181015183820152602001610f38565b83811115610dad5750506000910152565b60008151808452610f79816020860160208601610f35565b601f01601f19169290920160200192915050565b602081526000610e186020830184610f61565b80356001600160a01b0381168114610fb757600080fd5b919050565b600060208284031215610fce57600080fd5b610e1882610fa0565b600060208284031215610fe957600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561101957600080fd5b61102283610fa0565b9150602083013567ffffffffffffffff8082111561103f57600080fd5b818501915085601f83011261105357600080fd5b81358181111561106557611065610ff0565b604051601f8201601f19908116603f0116810190838211818310171561108d5761108d610ff0565b816040528281528860208487010111156110a657600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b80358015158114610fb757600080fd5b600080604083850312156110eb57600080fd5b6110f483610fa0565b9150611102602084016110c8565b90509250929050565b6000806020838503121561111e57600080fd5b823567ffffffffffffffff8082111561113657600080fd5b818501915085601f83011261114a57600080fd5b81358181111561115957600080fd5b8660208260051b850101111561116e57600080fd5b60209290920196919550909350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156111f557888303603f1901855281518051151584528781015115158885015286015160608785018190526111e181860183610f61565b9689019694505050908601906001016111a7565b509098975050505050505050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008235607e198336030181126112c757600080fd5b9190910192915050565b6000808335601e198436030181126112e857600080fd5b83018035915067ffffffffffffffff82111561130357600080fd5b60200191503681900382131561131857600080fd5b9250929050565b8183823760009101908152919050565b60006020828403121561134157600080fd5b610e18826110c8565b634e487b7160e01b600052601160045260246000fd5b6000600182016113725761137261134a565b5060010190565b60008282101561138b5761138b61134a565b500390565b6000602082840312156113a257600080fd5b5051919050565b600082516112c7818460208701610f3556fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300080f000a",
}

// ExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutorMetaData.ABI instead.
var ExecutorABI = ExecutorMetaData.ABI

// ExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecutorMetaData.Bin instead.
var ExecutorBin = ExecutorMetaData.Bin

// DeployExecutor deploys a new Ethereum contract, binding an instance of Executor to it.
func DeployExecutor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Executor, error) {
	parsed, err := ExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Executor{ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

// Executor is an auto generated Go binding around an Ethereum contract.
type Executor struct {
	ExecutorCaller     // Read-only binding to the contract
	ExecutorTransactor // Write-only binding to the contract
	ExecutorFilterer   // Log filterer for contract events
}

// ExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutorSession struct {
	Contract     *Executor         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutorCallerSession struct {
	Contract *ExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutorTransactorSession struct {
	Contract     *ExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutorRaw struct {
	Contract *Executor // Generic contract binding to access the raw methods on
}

// ExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutorCallerRaw struct {
	Contract *ExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutorTransactorRaw struct {
	Contract *ExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutor creates a new instance of Executor, bound to a specific deployed contract.
func NewExecutor(address common.Address, backend bind.ContractBackend) (*Executor, error) {
	contract, err := bindExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Executor{ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

// NewExecutorCaller creates a new read-only instance of Executor, bound to a specific deployed contract.
func NewExecutorCaller(address common.Address, caller bind.ContractCaller) (*ExecutorCaller, error) {
	contract, err := bindExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorCaller{contract: contract}, nil
}

// NewExecutorTransactor creates a new write-only instance of Executor, bound to a specific deployed contract.
func NewExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorTransactor, error) {
	contract, err := bindExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorTransactor{contract: contract}, nil
}

// NewExecutorFilterer creates a new log filterer instance of Executor, bound to a specific deployed contract.
func NewExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorFilterer, error) {
	contract, err := bindExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorFilterer{contract: contract}, nil
}

// bindExecutor binds a generic wrapper to an already deployed contract.
func bindExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.ExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transact(opts, method, params...)
}

// Authorities is a free data retrieval call binding the contract method 0x91223d69.
//
// Solidity: function authorities(address ) view returns(bool)
func (_Executor *ExecutorCaller) Authorities(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "authorities", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Authorities is a free data retrieval call binding the contract method 0x91223d69.
//
// Solidity: function authorities(address ) view returns(bool)
func (_Executor *ExecutorSession) Authorities(arg0 common.Address) (bool, error) {
	return _Executor.Contract.Authorities(&_Executor.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x91223d69.
//
// Solidity: function authorities(address ) view returns(bool)
func (_Executor *ExecutorCallerSession) Authorities(arg0 common.Address) (bool, error) {
	return _Executor.Contract.Authorities(&_Executor.CallOpts, arg0)
}

// Index is a free data retrieval call binding the contract method 0x2986c0e5.
//
// Solidity: function index() view returns(uint256)
func (_Executor *ExecutorCaller) Index(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "index")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index is a free data retrieval call binding the contract method 0x2986c0e5.
//
// Solidity: function index() view returns(uint256)
func (_Executor *ExecutorSession) Index() (*big.Int, error) {
	return _Executor.Contract.Index(&_Executor.CallOpts)
}

// Index is a free data retrieval call binding the contract method 0x2986c0e5.
//
// Solidity: function index() view returns(uint256)
func (_Executor *ExecutorCallerSession) Index() (*big.Int, error) {
	return _Executor.Contract.Index(&_Executor.CallOpts)
}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Executor *ExecutorCaller) Keepers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "keepers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Executor *ExecutorSession) Keepers(arg0 common.Address) (bool, error) {
	return _Executor.Contract.Keepers(&_Executor.CallOpts, arg0)
}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Executor *ExecutorCallerSession) Keepers(arg0 common.Address) (bool, error) {
	return _Executor.Contract.Keepers(&_Executor.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Executor *ExecutorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Executor *ExecutorSession) Name() (string, error) {
	return _Executor.Contract.Name(&_Executor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Executor *ExecutorCallerSession) Name() (string, error) {
	return _Executor.Contract.Name(&_Executor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Executor *ExecutorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Executor *ExecutorSession) Owner() (common.Address, error) {
	return _Executor.Contract.Owner(&_Executor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Executor *ExecutorCallerSession) Owner() (common.Address, error) {
	return _Executor.Contract.Owner(&_Executor.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Executor *ExecutorCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Executor *ExecutorSession) ProxiableUUID() ([32]byte, error) {
	return _Executor.Contract.ProxiableUUID(&_Executor.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Executor *ExecutorCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Executor.Contract.ProxiableUUID(&_Executor.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x8a8ef0b3.
//
// Solidity: function execute((uint256,bool,address,bytes)[] calls) returns((bool,bool,bytes)[] results)
func (_Executor *ExecutorTransactor) Execute(opts *bind.TransactOpts, calls []ExecutorInterfaceCall) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "execute", calls)
}

// Execute is a paid mutator transaction binding the contract method 0x8a8ef0b3.
//
// Solidity: function execute((uint256,bool,address,bytes)[] calls) returns((bool,bool,bytes)[] results)
func (_Executor *ExecutorSession) Execute(calls []ExecutorInterfaceCall) (*types.Transaction, error) {
	return _Executor.Contract.Execute(&_Executor.TransactOpts, calls)
}

// Execute is a paid mutator transaction binding the contract method 0x8a8ef0b3.
//
// Solidity: function execute((uint256,bool,address,bytes)[] calls) returns((bool,bool,bytes)[] results)
func (_Executor *ExecutorTransactorSession) Execute(calls []ExecutorInterfaceCall) (*types.Transaction, error) {
	return _Executor.Contract.Execute(&_Executor.TransactOpts, calls)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Executor *ExecutorTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Executor *ExecutorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Executor.Contract.Initialize(&_Executor.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Executor *ExecutorTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Executor.Contract.Initialize(&_Executor.TransactOpts, _owner)
}

// LoadIndex is a paid mutator transaction binding the contract method 0x38ece48b.
//
// Solidity: function loadIndex(uint256 _index) returns()
func (_Executor *ExecutorTransactor) LoadIndex(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "loadIndex", _index)
}

// LoadIndex is a paid mutator transaction binding the contract method 0x38ece48b.
//
// Solidity: function loadIndex(uint256 _index) returns()
func (_Executor *ExecutorSession) LoadIndex(_index *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.LoadIndex(&_Executor.TransactOpts, _index)
}

// LoadIndex is a paid mutator transaction binding the contract method 0x38ece48b.
//
// Solidity: function loadIndex(uint256 _index) returns()
func (_Executor *ExecutorTransactorSession) LoadIndex(_index *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.LoadIndex(&_Executor.TransactOpts, _index)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Executor *ExecutorTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Executor *ExecutorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Executor.Contract.SetOwner(&_Executor.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Executor *ExecutorTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Executor.Contract.SetOwner(&_Executor.TransactOpts, newOwner)
}

// UpdateAuthority is a paid mutator transaction binding the contract method 0x6cd22eaf.
//
// Solidity: function updateAuthority(address authority, bool allowed) returns()
func (_Executor *ExecutorTransactor) UpdateAuthority(opts *bind.TransactOpts, authority common.Address, allowed bool) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "updateAuthority", authority, allowed)
}

// UpdateAuthority is a paid mutator transaction binding the contract method 0x6cd22eaf.
//
// Solidity: function updateAuthority(address authority, bool allowed) returns()
func (_Executor *ExecutorSession) UpdateAuthority(authority common.Address, allowed bool) (*types.Transaction, error) {
	return _Executor.Contract.UpdateAuthority(&_Executor.TransactOpts, authority, allowed)
}

// UpdateAuthority is a paid mutator transaction binding the contract method 0x6cd22eaf.
//
// Solidity: function updateAuthority(address authority, bool allowed) returns()
func (_Executor *ExecutorTransactorSession) UpdateAuthority(authority common.Address, allowed bool) (*types.Transaction, error) {
	return _Executor.Contract.UpdateAuthority(&_Executor.TransactOpts, authority, allowed)
}

// UpdateKeeper is a paid mutator transaction binding the contract method 0xd3057877.
//
// Solidity: function updateKeeper(address keeper, bool allowed) returns()
func (_Executor *ExecutorTransactor) UpdateKeeper(opts *bind.TransactOpts, keeper common.Address, allowed bool) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "updateKeeper", keeper, allowed)
}

// UpdateKeeper is a paid mutator transaction binding the contract method 0xd3057877.
//
// Solidity: function updateKeeper(address keeper, bool allowed) returns()
func (_Executor *ExecutorSession) UpdateKeeper(keeper common.Address, allowed bool) (*types.Transaction, error) {
	return _Executor.Contract.UpdateKeeper(&_Executor.TransactOpts, keeper, allowed)
}

// UpdateKeeper is a paid mutator transaction binding the contract method 0xd3057877.
//
// Solidity: function updateKeeper(address keeper, bool allowed) returns()
func (_Executor *ExecutorTransactorSession) UpdateKeeper(keeper common.Address, allowed bool) (*types.Transaction, error) {
	return _Executor.Contract.UpdateKeeper(&_Executor.TransactOpts, keeper, allowed)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Executor *ExecutorTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Executor *ExecutorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Executor.Contract.UpgradeTo(&_Executor.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Executor *ExecutorTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Executor.Contract.UpgradeTo(&_Executor.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Executor *ExecutorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Executor *ExecutorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Executor.Contract.UpgradeToAndCall(&_Executor.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Executor *ExecutorTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Executor.Contract.UpgradeToAndCall(&_Executor.TransactOpts, newImplementation, data)
}

// ExecutorAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Executor contract.
type ExecutorAdminChangedIterator struct {
	Event *ExecutorAdminChanged // Event containing the contract specifics and raw log

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
func (it *ExecutorAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorAdminChanged)
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
		it.Event = new(ExecutorAdminChanged)
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
func (it *ExecutorAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorAdminChanged represents a AdminChanged event raised by the Executor contract.
type ExecutorAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Executor *ExecutorFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ExecutorAdminChangedIterator, error) {

	logs, sub, err := _Executor.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ExecutorAdminChangedIterator{contract: _Executor.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Executor *ExecutorFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ExecutorAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Executor.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorAdminChanged)
				if err := _Executor.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Executor *ExecutorFilterer) ParseAdminChanged(log types.Log) (*ExecutorAdminChanged, error) {
	event := new(ExecutorAdminChanged)
	if err := _Executor.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the Executor contract.
type ExecutorAuthorityUpdatedIterator struct {
	Event *ExecutorAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutorAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorAuthorityUpdated)
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
		it.Event = new(ExecutorAuthorityUpdated)
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
func (it *ExecutorAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorAuthorityUpdated represents a AuthorityUpdated event raised by the Executor contract.
type ExecutorAuthorityUpdated struct {
	Authority common.Address
	Allowed   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0xc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc.
//
// Solidity: event AuthorityUpdated(address indexed authority, bool indexed allowed)
func (_Executor *ExecutorFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts, authority []common.Address, allowed []bool) (*ExecutorAuthorityUpdatedIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "AuthorityUpdated", authorityRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorAuthorityUpdatedIterator{contract: _Executor.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0xc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc.
//
// Solidity: event AuthorityUpdated(address indexed authority, bool indexed allowed)
func (_Executor *ExecutorFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorAuthorityUpdated, authority []common.Address, allowed []bool) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "AuthorityUpdated", authorityRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorAuthorityUpdated)
				if err := _Executor.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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

// ParseAuthorityUpdated is a log parse operation binding the contract event 0xc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc.
//
// Solidity: event AuthorityUpdated(address indexed authority, bool indexed allowed)
func (_Executor *ExecutorFilterer) ParseAuthorityUpdated(log types.Log) (*ExecutorAuthorityUpdated, error) {
	event := new(ExecutorAuthorityUpdated)
	if err := _Executor.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Executor contract.
type ExecutorBeaconUpgradedIterator struct {
	Event *ExecutorBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ExecutorBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorBeaconUpgraded)
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
		it.Event = new(ExecutorBeaconUpgraded)
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
func (it *ExecutorBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorBeaconUpgraded represents a BeaconUpgraded event raised by the Executor contract.
type ExecutorBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Executor *ExecutorFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ExecutorBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorBeaconUpgradedIterator{contract: _Executor.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Executor *ExecutorFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ExecutorBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorBeaconUpgraded)
				if err := _Executor.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Executor *ExecutorFilterer) ParseBeaconUpgraded(log types.Log) (*ExecutorBeaconUpgraded, error) {
	event := new(ExecutorBeaconUpgraded)
	if err := _Executor.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Executor contract.
type ExecutorExecutedIterator struct {
	Event *ExecutorExecuted // Event containing the contract specifics and raw log

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
func (it *ExecutorExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorExecuted)
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
		it.Event = new(ExecutorExecuted)
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
func (it *ExecutorExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorExecuted represents a Executed event raised by the Executor contract.
type ExecutorExecuted struct {
	Index   *big.Int
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xfa9788a6b2f6e96ab2d85a5b4ed3668e72b91e06c5eec6b3e87afd95fe536863.
//
// Solidity: event Executed(uint256 indexed index, bool success)
func (_Executor *ExecutorFilterer) FilterExecuted(opts *bind.FilterOpts, index []*big.Int) (*ExecutorExecutedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "Executed", indexRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorExecutedIterator{contract: _Executor.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xfa9788a6b2f6e96ab2d85a5b4ed3668e72b91e06c5eec6b3e87afd95fe536863.
//
// Solidity: event Executed(uint256 indexed index, bool success)
func (_Executor *ExecutorFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ExecutorExecuted, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "Executed", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorExecuted)
				if err := _Executor.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xfa9788a6b2f6e96ab2d85a5b4ed3668e72b91e06c5eec6b3e87afd95fe536863.
//
// Solidity: event Executed(uint256 indexed index, bool success)
func (_Executor *ExecutorFilterer) ParseExecuted(log types.Log) (*ExecutorExecuted, error) {
	event := new(ExecutorExecuted)
	if err := _Executor.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorIndexUpdatedIterator is returned from FilterIndexUpdated and is used to iterate over the raw logs and unpacked data for IndexUpdated events raised by the Executor contract.
type ExecutorIndexUpdatedIterator struct {
	Event *ExecutorIndexUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutorIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorIndexUpdated)
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
		it.Event = new(ExecutorIndexUpdated)
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
func (it *ExecutorIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorIndexUpdated represents a IndexUpdated event raised by the Executor contract.
type ExecutorIndexUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIndexUpdated is a free log retrieval operation binding the contract event 0xacfc085c9be45d2b3f9e5c09a19d4a95749cc16939519c13e090de3a4cb192c6.
//
// Solidity: event IndexUpdated(uint256 indexed index)
func (_Executor *ExecutorFilterer) FilterIndexUpdated(opts *bind.FilterOpts, index []*big.Int) (*ExecutorIndexUpdatedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "IndexUpdated", indexRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorIndexUpdatedIterator{contract: _Executor.contract, event: "IndexUpdated", logs: logs, sub: sub}, nil
}

// WatchIndexUpdated is a free log subscription operation binding the contract event 0xacfc085c9be45d2b3f9e5c09a19d4a95749cc16939519c13e090de3a4cb192c6.
//
// Solidity: event IndexUpdated(uint256 indexed index)
func (_Executor *ExecutorFilterer) WatchIndexUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorIndexUpdated, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "IndexUpdated", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorIndexUpdated)
				if err := _Executor.contract.UnpackLog(event, "IndexUpdated", log); err != nil {
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

// ParseIndexUpdated is a log parse operation binding the contract event 0xacfc085c9be45d2b3f9e5c09a19d4a95749cc16939519c13e090de3a4cb192c6.
//
// Solidity: event IndexUpdated(uint256 indexed index)
func (_Executor *ExecutorFilterer) ParseIndexUpdated(log types.Log) (*ExecutorIndexUpdated, error) {
	event := new(ExecutorIndexUpdated)
	if err := _Executor.contract.UnpackLog(event, "IndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorKeeperUpdatedIterator is returned from FilterKeeperUpdated and is used to iterate over the raw logs and unpacked data for KeeperUpdated events raised by the Executor contract.
type ExecutorKeeperUpdatedIterator struct {
	Event *ExecutorKeeperUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutorKeeperUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorKeeperUpdated)
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
		it.Event = new(ExecutorKeeperUpdated)
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
func (it *ExecutorKeeperUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorKeeperUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorKeeperUpdated represents a KeeperUpdated event raised by the Executor contract.
type ExecutorKeeperUpdated struct {
	Keeper  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeeperUpdated is a free log retrieval operation binding the contract event 0x786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c00.
//
// Solidity: event KeeperUpdated(address indexed keeper, bool indexed allowed)
func (_Executor *ExecutorFilterer) FilterKeeperUpdated(opts *bind.FilterOpts, keeper []common.Address, allowed []bool) (*ExecutorKeeperUpdatedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "KeeperUpdated", keeperRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorKeeperUpdatedIterator{contract: _Executor.contract, event: "KeeperUpdated", logs: logs, sub: sub}, nil
}

// WatchKeeperUpdated is a free log subscription operation binding the contract event 0x786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c00.
//
// Solidity: event KeeperUpdated(address indexed keeper, bool indexed allowed)
func (_Executor *ExecutorFilterer) WatchKeeperUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorKeeperUpdated, keeper []common.Address, allowed []bool) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "KeeperUpdated", keeperRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorKeeperUpdated)
				if err := _Executor.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
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

// ParseKeeperUpdated is a log parse operation binding the contract event 0x786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c00.
//
// Solidity: event KeeperUpdated(address indexed keeper, bool indexed allowed)
func (_Executor *ExecutorFilterer) ParseKeeperUpdated(log types.Log) (*ExecutorKeeperUpdated, error) {
	event := new(ExecutorKeeperUpdated)
	if err := _Executor.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorOwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the Executor contract.
type ExecutorOwnerUpdatedIterator struct {
	Event *ExecutorOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutorOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOwnerUpdated)
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
		it.Event = new(ExecutorOwnerUpdated)
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
func (it *ExecutorOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorOwnerUpdated represents a OwnerUpdated event raised by the Executor contract.
type ExecutorOwnerUpdated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed newOwner)
func (_Executor *ExecutorFilterer) FilterOwnerUpdated(opts *bind.FilterOpts, newOwner []common.Address) (*ExecutorOwnerUpdatedIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "OwnerUpdated", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOwnerUpdatedIterator{contract: _Executor.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed newOwner)
func (_Executor *ExecutorFilterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorOwnerUpdated, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "OwnerUpdated", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorOwnerUpdated)
				if err := _Executor.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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

// ParseOwnerUpdated is a log parse operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed newOwner)
func (_Executor *ExecutorFilterer) ParseOwnerUpdated(log types.Log) (*ExecutorOwnerUpdated, error) {
	event := new(ExecutorOwnerUpdated)
	if err := _Executor.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Executor contract.
type ExecutorUpgradedIterator struct {
	Event *ExecutorUpgraded // Event containing the contract specifics and raw log

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
func (it *ExecutorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorUpgraded)
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
		it.Event = new(ExecutorUpgraded)
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
func (it *ExecutorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorUpgraded represents a Upgraded event raised by the Executor contract.
type ExecutorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Executor *ExecutorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ExecutorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorUpgradedIterator{contract: _Executor.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Executor *ExecutorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ExecutorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorUpgraded)
				if err := _Executor.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Executor *ExecutorFilterer) ParseUpgraded(log types.Log) (*ExecutorUpgraded, error) {
	event := new(ExecutorUpgraded)
	if err := _Executor.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutorVersionInitializedIterator is returned from FilterVersionInitialized and is used to iterate over the raw logs and unpacked data for VersionInitialized events raised by the Executor contract.
type ExecutorVersionInitializedIterator struct {
	Event *ExecutorVersionInitialized // Event containing the contract specifics and raw log

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
func (it *ExecutorVersionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorVersionInitialized)
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
		it.Event = new(ExecutorVersionInitialized)
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
func (it *ExecutorVersionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutorVersionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutorVersionInitialized represents a VersionInitialized event raised by the Executor contract.
type ExecutorVersionInitialized struct {
	NewVersion *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVersionInitialized is a free log retrieval operation binding the contract event 0x7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d.
//
// Solidity: event VersionInitialized(uint256 indexed newVersion)
func (_Executor *ExecutorFilterer) FilterVersionInitialized(opts *bind.FilterOpts, newVersion []*big.Int) (*ExecutorVersionInitializedIterator, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "VersionInitialized", newVersionRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorVersionInitializedIterator{contract: _Executor.contract, event: "VersionInitialized", logs: logs, sub: sub}, nil
}

// WatchVersionInitialized is a free log subscription operation binding the contract event 0x7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d.
//
// Solidity: event VersionInitialized(uint256 indexed newVersion)
func (_Executor *ExecutorFilterer) WatchVersionInitialized(opts *bind.WatchOpts, sink chan<- *ExecutorVersionInitialized, newVersion []*big.Int) (event.Subscription, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "VersionInitialized", newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutorVersionInitialized)
				if err := _Executor.contract.UnpackLog(event, "VersionInitialized", log); err != nil {
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

// ParseVersionInitialized is a log parse operation binding the contract event 0x7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d.
//
// Solidity: event VersionInitialized(uint256 indexed newVersion)
func (_Executor *ExecutorFilterer) ParseVersionInitialized(log types.Log) (*ExecutorVersionInitialized, error) {
	event := new(ExecutorVersionInitialized)
	if err := _Executor.contract.UnpackLog(event, "VersionInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
