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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotKeeper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"UnknownRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"callHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"KeeperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"VersionInitialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorities\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fail\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structExecutorInterface.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"called\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structExecutorInterface.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keepers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523060805234801561001457600080fd5b5060405164086c2d8d8560db1b60208201526d1d5a5b9d0c8d4d881a5b99195e0b60921b602582015269189bdbdb0819985a5b0b60b21b60338201526e1859191c995cdcc81d185c99d95d0b608a1b603d8201526b62797465733332206461746160a01b604c820152602960f81b605882015260590160408051601f19818403018152919052805160209091012060a05260805160a051611bb36100e6600039600061129c0152600081816103c101528181610476015281816105c80152818161067801526107c10152611bb36000f3fe6080604052600436106100d25760003560e01c80636cd22eaf1161007f57806391223d691161005957806391223d691461028d578063a9fcfb33146102bd578063c4d66de81461030d578063d30578771461032d57600080fd5b80636cd22eaf146101ee5780638a8ef0b31461020e5780638da5cb5b1461023b57600080fd5b80633bbd64bc116100b05780633bbd64bc146101785780634f1ef286146101b857806352d1902d146101cb57600080fd5b806306fdde03146100d757806313af4035146101365780633659cfe614610158575b600080fd5b3480156100e357600080fd5b506101206040518060400160405280600881526020017f4578656375746f7200000000000000000000000000000000000000000000000081525081565b60405161012d91906116fa565b60405180910390f35b34801561014257600080fd5b50610156610151366004611736565b61034d565b005b34801561016457600080fd5b50610156610173366004611736565b6103aa565b34801561018457600080fd5b506101a8610193366004611736565b60046020526000908152604090205460ff1681565b604051901515815260200161012d565b6101566101c6366004611780565b6105b1565b3480156101d757600080fd5b506101e06107a7565b60405190815260200161012d565b3480156101fa57600080fd5b50610156610209366004611870565b610893565b34801561021a57600080fd5b5061022e6102293660046118a3565b6109b0565b60405161012d9190611918565b34801561024757600080fd5b506000546102689073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012d565b34801561029957600080fd5b506101a86102a8366004611736565b60036020526000908152604090205460ff1681565b3480156102c957600080fd5b506102f66102d83660046119b9565b60056020526000908152604090205460ff8082169161010090041682565b60408051921515835290151560208301520161012d565b34801561031957600080fd5b50610156610328366004611736565b610e28565b34801561033957600080fd5b50610156610348366004611870565b610eb9565b60005473ffffffffffffffffffffffffffffffffffffffff16331461039e576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103a781610fd6565b50565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163003610474576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c000000000000000000000000000000000000000060648201526084015b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166104e97f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161461058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f6163746976652070726f78790000000000000000000000000000000000000000606482015260840161046b565b61059581611043565b604080516000808252602082019092526103a791839190611094565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163003610676576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c0000000000000000000000000000000000000000606482015260840161046b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166106eb7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161461078e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f6163746976652070726f78790000000000000000000000000000000000000000606482015260840161046b565b61079782611043565b6107a382826001611094565b5050565b60003073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461086e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000606482015260840161046b565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b60005473ffffffffffffffffffffffffffffffffffffffff1633146108e4576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610931576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660008181526003602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915590519092917fc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc91a35050565b60606002546001146109ee576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280553360009081526004602052604090205460ff16610a3b576040517ff512b27800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b818067ffffffffffffffff811115610a5557610a55611751565b604051908082528060200260200182016040528015610aa257816020015b60408051606080820183526000808352602083015291810191909152815260200190600190039081610a735790505b50915060005b81811015610e1b576000610ade868684818110610ac757610ac76119d2565b9050602002810190610ad99190611a01565b611298565b60008181526005602090815260409182902082518084019093525460ff80821615801585526101009092041615159183019190915291925090610b72576001858481518110610b2f57610b2f6119d2565b6020908102919091018101519115159091528101518551869085908110610b5857610b586119d2565b60209081029190910181015191151591015250610e0b9050565b600080888886818110610b8757610b876119d2565b9050602002810190610b999190611a01565b610baa906060810190604001611736565b73ffffffffffffffffffffffffffffffffffffffff16898987818110610bd257610bd26119d2565b9050602002810190610be49190611a01565b610bf2906060810190611a3f565b604051610c00929190611aab565b6000604051808303816000865af19150503d8060008114610c3d576040519150601f19603f3d011682016040523d82523d6000602084013e610c42565b606091505b509150915081158015610c875750888886818110610c6257610c626119d2565b9050602002810190610c749190611a01565b610c85906040810190602001611abb565b155b15610cd2578051600003610cca576040517fdacad9170000000000000000000000000000000000000000000000000000000081526004810186905260240161046b565b805181602001fd5b60408051808201825260018152831515602080830191825260008881526005909152929092209051815492517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009093169015157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16176101009215159290920291909117905586518190889087908110610d6e57610d6e6119d2565b60200260200101516040018190525081878681518110610d9057610d906119d2565b602090810291909101810151911515910152888886818110610db457610db46119d2565b9050602002810190610dc69190611a01565b60000135847fccaac8e08856c6267e4b7f91eefa649b1c19cbdeb351f42bd21b09d44c78d39f84604051610dfe911515815260200190565b60405180910390a3505050505b610e1481611b05565b9050610aa8565b5050600160025592915050565b6001610e348180611b3d565b60015414610e6e576040517f5daa87a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600103610e7c5760016002555b610e8582610fd6565b600181905560405181907f7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d90600090a25050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610f0a576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610f57576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660008181526004602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915590519092917f786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c0091a35050565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117825560405190917f4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b91a250565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103a7576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156110cc576110c78361135b565b505050565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611151575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261114e91810190611b54565b60015b6111dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201527f6f6e206973206e6f742055555053000000000000000000000000000000000000606482015260840161046b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461128c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f7860448201527f6961626c65555549440000000000000000000000000000000000000000000000606482015260840161046b565b506110c7838383611465565b60007f000000000000000000000000000000000000000000000000000000000000000082356112cd6040850160208601611abb565b6112dd6060860160408701611736565b6112ea6060870187611a3f565b6040516112f8929190611aab565b604080519182900382206020830196909652810193909352901515606083015273ffffffffffffffffffffffffffffffffffffffff16608082015260a081019190915260c001604051602081830303815290604052805190602001209050919050565b73ffffffffffffffffffffffffffffffffffffffff81163b6113ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e747261637400000000000000000000000000000000000000606482015260840161046b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b61146e83611490565b60008251118061147b5750805b156110c75761148a83836114dd565b50505050565b6114998161135b565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606115028383604051806060016040528060278152602001611b8060279139611509565b9392505050565b606073ffffffffffffffffffffffffffffffffffffffff84163b6115af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e74726163740000000000000000000000000000000000000000000000000000606482015260840161046b565b6000808573ffffffffffffffffffffffffffffffffffffffff16856040516115d79190611b6d565b600060405180830381855af49150503d8060008114611612576040519150601f19603f3d011682016040523d82523d6000602084013e611617565b606091505b5091509150611627828286611631565b9695505050505050565b60608315611640575081611502565b8251156116505782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046b91906116fa565b60005b8381101561169f578181015183820152602001611687565b8381111561148a5750506000910152565b600081518084526116c8816020860160208601611684565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061150260208301846116b0565b803573ffffffffffffffffffffffffffffffffffffffff8116811461173157600080fd5b919050565b60006020828403121561174857600080fd5b6115028261170d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561179357600080fd5b61179c8361170d565b9150602083013567ffffffffffffffff808211156117b957600080fd5b818501915085601f8301126117cd57600080fd5b8135818111156117df576117df611751565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561182557611825611751565b8160405282815288602084870101111561183e57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b8035801515811461173157600080fd5b6000806040838503121561188357600080fd5b61188c8361170d565b915061189a60208401611860565b90509250929050565b600080602083850312156118b657600080fd5b823567ffffffffffffffff808211156118ce57600080fd5b818501915085601f8301126118e257600080fd5b8135818111156118f157600080fd5b8660208260051b850101111561190657600080fd5b60209290920196919550909350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156119ab578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc00185528151805115158452878101511515888501528601516060878501819052611997818601836116b0565b96890196945050509086019060010161193f565b509098975050505050505050565b6000602082840312156119cb57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112611a3557600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611a7457600080fd5b83018035915067ffffffffffffffff821115611a8f57600080fd5b602001915036819003821315611aa457600080fd5b9250929050565b8183823760009101908152919050565b600060208284031215611acd57600080fd5b61150282611860565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b3657611b36611ad6565b5060010190565b600082821015611b4f57611b4f611ad6565b500390565b600060208284031215611b6657600080fd5b5051919050565b60008251611a3581846020870161168456fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300080f000a",
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

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool executed, bool success)
func (_Executor *ExecutorCaller) Executed(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Executed bool
	Success  bool
}, error) {
	var out []interface{}
	err := _Executor.contract.Call(opts, &out, "executed", arg0)

	outstruct := new(struct {
		Executed bool
		Success  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Executed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Success = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool executed, bool success)
func (_Executor *ExecutorSession) Executed(arg0 [32]byte) (struct {
	Executed bool
	Success  bool
}, error) {
	return _Executor.Contract.Executed(&_Executor.CallOpts, arg0)
}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool executed, bool success)
func (_Executor *ExecutorCallerSession) Executed(arg0 [32]byte) (struct {
	Executed bool
	Success  bool
}, error) {
	return _Executor.Contract.Executed(&_Executor.CallOpts, arg0)
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
	CallHash [32]byte
	Index    *big.Int
	Success  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xccaac8e08856c6267e4b7f91eefa649b1c19cbdeb351f42bd21b09d44c78d39f.
//
// Solidity: event Executed(bytes32 indexed callHash, uint256 indexed index, bool success)
func (_Executor *ExecutorFilterer) FilterExecuted(opts *bind.FilterOpts, callHash [][32]byte, index []*big.Int) (*ExecutorExecutedIterator, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Executor.contract.FilterLogs(opts, "Executed", callHashRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorExecutedIterator{contract: _Executor.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xccaac8e08856c6267e4b7f91eefa649b1c19cbdeb351f42bd21b09d44c78d39f.
//
// Solidity: event Executed(bytes32 indexed callHash, uint256 indexed index, bool success)
func (_Executor *ExecutorFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ExecutorExecuted, callHash [][32]byte, index []*big.Int) (event.Subscription, error) {

	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Executor.contract.WatchLogs(opts, "Executed", callHashRule, indexRule)
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

// ParseExecuted is a log parse operation binding the contract event 0xccaac8e08856c6267e4b7f91eefa649b1c19cbdeb351f42bd21b09d44c78d39f.
//
// Solidity: event Executed(bytes32 indexed callHash, uint256 indexed index, bool success)
func (_Executor *ExecutorFilterer) ParseExecuted(log types.Log) (*ExecutorExecuted, error) {
	event := new(ExecutorExecuted)
	if err := _Executor.contract.UnpackLog(event, "Executed", log); err != nil {
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
