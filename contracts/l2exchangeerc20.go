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

// L2ExchangeERC20MetaData contains all meta data concerning the L2ExchangeERC20 contract.
var L2ExchangeERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accounts\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"contractAccountsInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620018633803806200186383398101604081905262000035916200023b565b858585600062000046848262000384565b50600162000055838262000384565b5060ff81166080524660a0526200006b620000b5565b60c0525050506001600160a01b0392831660e052908216610100521661012081905230600090815260046020908152604080832093835292905220600019905550620004ce915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051620000e9919062000450565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200017957600080fd5b81516001600160401b038082111562000196576200019662000151565b604051601f8301601f19908116603f01168101908282118183101715620001c157620001c162000151565b81604052838152602092508683858801011115620001de57600080fd5b600091505b83821015620002025785820183015181830184015290820190620001e3565b83821115620002145760008385830101525b9695505050505050565b80516001600160a01b03811681146200023657600080fd5b919050565b60008060008060008060c087890312156200025557600080fd5b86516001600160401b03808211156200026d57600080fd5b6200027b8a838b0162000167565b975060208901519150808211156200029257600080fd5b50620002a189828a0162000167565b955050604087015160ff81168114620002b957600080fd5b9350620002c9606088016200021e565b9250620002d9608088016200021e565b9150620002e960a088016200021e565b90509295509295509295565b600181811c908216806200030a57607f821691505b6020821081036200032b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200037f57600081815260208120601f850160051c810160208610156200035a5750805b601f850160051c820191505b818110156200037b5782815560010162000366565b5050505b505050565b81516001600160401b03811115620003a057620003a062000151565b620003b881620003b18454620002f5565b8462000331565b602080601f831160018114620003f05760008415620003d75750858301515b600019600386901b1c1916600185901b1785556200037b565b600085815260208120601f198616915b82811015620004215788860151825594840194600190910190840162000400565b5085821015620004405787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008083546200046081620002f5565b600182811680156200047b57600181146200049157620004c2565b60ff1984168752821515830287019450620004c2565b8760005260208060002060005b85811015620004b95781548a8201529084019082016200049e565b50505082870194505b50929695505050505050565b60805160a05160c05160e0516101005161012051611330620005336000396000818161021001526108790152600081816102ca015281816107910152610960015260006102f10152600061075701526000610722015260006101ba01526113306000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c806370a08231116100b2578063a9059cbb11610081578063c01e1bd611610066578063c01e1bd6146102ec578063d505accf14610313578063dd62ed3e1461032657600080fd5b8063a9059cbb146102b2578063ae1f6aaf146102c557600080fd5b806370a08231146102575780637ecebe001461027757806395d89b41146102975780639dc29fac1461029f57600080fd5b806323b872dd116101095780633644e515116100ee5780633644e515146101ee57806340c10f19146101f657806368cd03f61461020b57600080fd5b806323b872dd146101a2578063313ce567146101b557600080fd5b806301ffc9a71461013b57806306fdde0314610163578063095ea7b31461017857806318160ddd1461018b575b600080fd5b61014e610149366004610f7e565b610351565b60405190151581526020015b60405180910390f35b61016b6104d3565b60405161015a9190610fc0565b61014e61018636600461105c565b610561565b61019460025481565b60405190815260200161015a565b61014e6101b0366004611086565b6105da565b6101dc7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff909116815260200161015a565b61019461071e565b61020961020436600461105c565b610779565b005b6102327f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b6101946102653660046110c2565b60036020526000908152604090205481565b6101946102853660046110c2565b60056020526000908152604090205481565b61016b61093b565b6102096102ad36600461105c565b610948565b61014e6102c036600461105c565b610a39565b6102327f000000000000000000000000000000000000000000000000000000000000000081565b6102327f000000000000000000000000000000000000000000000000000000000000000081565b6102096103213660046110dd565b610abe565b610194610334366004611150565b600460209081526000928352604080842090915290825290205481565b60007f1d1d8b63000000000000000000000000000000000000000000000000000000007f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000841614806103e857507fffffffff00000000000000000000000000000000000000000000000000000000838116908216145b8061043457507fffffffff0000000000000000000000000000000000000000000000000000000083167fc01e1bd600000000000000000000000000000000000000000000000000000000145b8061048057507fffffffff0000000000000000000000000000000000000000000000000000000083167f40c10f1900000000000000000000000000000000000000000000000000000000145b806104cc57507fffffffff0000000000000000000000000000000000000000000000000000000083167f9dc29fac00000000000000000000000000000000000000000000000000000000145b9392505050565b600080546104e090611183565b80601f016020809104026020016040519081016040528092919081815260200182805461050c90611183565b80156105595780601f1061052e57610100808354040283529160200191610559565b820191906000526020600020905b81548152906001019060200180831161053c57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906105c99086815260200190565b60405180910390a350600192915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526004602090815260408083203384529091528120547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461066e5761063c8382611205565b73ffffffffffffffffffffffffffffffffffffffff861660009081526004602090815260408083203384529091529020555b73ffffffffffffffffffffffffffffffffffffffff8516600090815260036020526040812080548592906106a3908490611205565b909155505073ffffffffffffffffffffffffffffffffffffffff808516600081815260036020526040908190208054870190555190918716907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9061070b9087815260200190565b60405180910390a3506001949350505050565b60007f000000000000000000000000000000000000000000000000000000000000000046146107545761074f610ddd565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461081d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e4f545f4c325f4252494447450000000000000000000000000000000000000060448201526064015b60405180910390fd5b6108273082610e77565b6040517f8340f54900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152306024830152604482018390527f00000000000000000000000000000000000000000000000000000000000000001690638340f549906064016020604051808303816000875af11580156108c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e6919061121c565b508173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161092f91815260200190565b60405180910390a25050565b600180546104e090611183565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146109e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e4f545f4c325f425249444745000000000000000000000000000000000000006044820152606401610814565b6109f18282610ef0565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161092f91815260200190565b33600090815260036020526040812080548391908390610a5a908490611205565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600081815260036020526040908190208054850190555133907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906105c99086815260200190565b42841015610b28576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f5045524d49545f444541444c494e455f455850495245440000000000000000006044820152606401610814565b60006001610b3461071e565b73ffffffffffffffffffffffffffffffffffffffff8a811660008181526005602090815260409182902080546001810190915582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98184015280840194909452938d166060840152608083018c905260a083019390935260c08083018b90528151808403909101815260e0830190915280519201919091207f190100000000000000000000000000000000000000000000000000000000000061010083015261010282019290925261012281019190915261014201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff871690820152606081018590526080810184905260a0016020604051602081039080840390855afa158015610c86573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615801590610d0157508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610d67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f494e56414c49445f5349474e45520000000000000000000000000000000000006044820152606401610814565b73ffffffffffffffffffffffffffffffffffffffff90811660009081526004602090815260408083208a8516808552908352928190208990555188815291928a16917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a350505050505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051610e0f9190611235565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b8060026000828254610e89919061130b565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000818152600360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91015b60405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602052604081208054839290610f25908490611205565b909155505060028054829003905560405181815260009073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610ee4565b600060208284031215610f9057600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146104cc57600080fd5b600060208083528351808285015260005b81811015610fed57858101830151858201604001528201610fd1565b81811115610fff576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461105757600080fd5b919050565b6000806040838503121561106f57600080fd5b61107883611033565b946020939093013593505050565b60008060006060848603121561109b57600080fd5b6110a484611033565b92506110b260208501611033565b9150604084013590509250925092565b6000602082840312156110d457600080fd5b6104cc82611033565b600080600080600080600060e0888a0312156110f857600080fd5b61110188611033565b965061110f60208901611033565b95506040880135945060608801359350608088013560ff8116811461113357600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561116357600080fd5b61116c83611033565b915061117a60208401611033565b90509250929050565b600181811c9082168061119757607f821691505b6020821081036111d0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611217576112176111d6565b500390565b60006020828403121561122e57600080fd5b5051919050565b600080835481600182811c91508083168061125157607f831692505b60208084108203611289577f4e487b710000000000000000000000000000000000000000000000000000000086526022600452602486fd5b81801561129d57600181146112d0576112fd565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00861689528415158502890196506112fd565b60008a81526020902060005b868110156112f55781548b8201529085019083016112dc565b505084890196505b509498975050505050505050565b6000821982111561131e5761131e6111d6565b50019056fea164736f6c634300080f000a",
}

// L2ExchangeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ExchangeERC20MetaData.ABI instead.
var L2ExchangeERC20ABI = L2ExchangeERC20MetaData.ABI

// L2ExchangeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ExchangeERC20MetaData.Bin instead.
var L2ExchangeERC20Bin = L2ExchangeERC20MetaData.Bin

// DeployL2ExchangeERC20 deploys a new Ethereum contract, binding an instance of L2ExchangeERC20 to it.
func DeployL2ExchangeERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8, _l1Token common.Address, _l2Bridge common.Address, _accounts common.Address) (common.Address, *types.Transaction, *L2ExchangeERC20, error) {
	parsed, err := L2ExchangeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ExchangeERC20Bin), backend, _name, _symbol, _decimals, _l1Token, _l2Bridge, _accounts)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ExchangeERC20{L2ExchangeERC20Caller: L2ExchangeERC20Caller{contract: contract}, L2ExchangeERC20Transactor: L2ExchangeERC20Transactor{contract: contract}, L2ExchangeERC20Filterer: L2ExchangeERC20Filterer{contract: contract}}, nil
}

// L2ExchangeERC20 is an auto generated Go binding around an Ethereum contract.
type L2ExchangeERC20 struct {
	L2ExchangeERC20Caller     // Read-only binding to the contract
	L2ExchangeERC20Transactor // Write-only binding to the contract
	L2ExchangeERC20Filterer   // Log filterer for contract events
}

// L2ExchangeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type L2ExchangeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ExchangeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ExchangeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ExchangeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ExchangeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ExchangeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ExchangeERC20Session struct {
	Contract     *L2ExchangeERC20  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ExchangeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ExchangeERC20CallerSession struct {
	Contract *L2ExchangeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L2ExchangeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ExchangeERC20TransactorSession struct {
	Contract     *L2ExchangeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L2ExchangeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type L2ExchangeERC20Raw struct {
	Contract *L2ExchangeERC20 // Generic contract binding to access the raw methods on
}

// L2ExchangeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ExchangeERC20CallerRaw struct {
	Contract *L2ExchangeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// L2ExchangeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ExchangeERC20TransactorRaw struct {
	Contract *L2ExchangeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ExchangeERC20 creates a new instance of L2ExchangeERC20, bound to a specific deployed contract.
func NewL2ExchangeERC20(address common.Address, backend bind.ContractBackend) (*L2ExchangeERC20, error) {
	contract, err := bindL2ExchangeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ExchangeERC20{L2ExchangeERC20Caller: L2ExchangeERC20Caller{contract: contract}, L2ExchangeERC20Transactor: L2ExchangeERC20Transactor{contract: contract}, L2ExchangeERC20Filterer: L2ExchangeERC20Filterer{contract: contract}}, nil
}

// NewL2ExchangeERC20Caller creates a new read-only instance of L2ExchangeERC20, bound to a specific deployed contract.
func NewL2ExchangeERC20Caller(address common.Address, caller bind.ContractCaller) (*L2ExchangeERC20Caller, error) {
	contract, err := bindL2ExchangeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ExchangeERC20Caller{contract: contract}, nil
}

// NewL2ExchangeERC20Transactor creates a new write-only instance of L2ExchangeERC20, bound to a specific deployed contract.
func NewL2ExchangeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*L2ExchangeERC20Transactor, error) {
	contract, err := bindL2ExchangeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ExchangeERC20Transactor{contract: contract}, nil
}

// NewL2ExchangeERC20Filterer creates a new log filterer instance of L2ExchangeERC20, bound to a specific deployed contract.
func NewL2ExchangeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*L2ExchangeERC20Filterer, error) {
	contract, err := bindL2ExchangeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ExchangeERC20Filterer{contract: contract}, nil
}

// bindL2ExchangeERC20 binds a generic wrapper to an already deployed contract.
func bindL2ExchangeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ExchangeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ExchangeERC20 *L2ExchangeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ExchangeERC20.Contract.L2ExchangeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ExchangeERC20 *L2ExchangeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.L2ExchangeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ExchangeERC20 *L2ExchangeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.L2ExchangeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ExchangeERC20 *L2ExchangeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ExchangeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ExchangeERC20 *L2ExchangeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ExchangeERC20 *L2ExchangeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _L2ExchangeERC20.Contract.DOMAINSEPARATOR(&_L2ExchangeERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _L2ExchangeERC20.Contract.DOMAINSEPARATOR(&_L2ExchangeERC20.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) Accounts(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "accounts")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Accounts() (common.Address, error) {
	return _L2ExchangeERC20.Contract.Accounts(&_L2ExchangeERC20.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) Accounts() (common.Address, error) {
	return _L2ExchangeERC20.Contract.Accounts(&_L2ExchangeERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L2ExchangeERC20.Contract.Allowance(&_L2ExchangeERC20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L2ExchangeERC20.Contract.Allowance(&_L2ExchangeERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _L2ExchangeERC20.Contract.BalanceOf(&_L2ExchangeERC20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _L2ExchangeERC20.Contract.BalanceOf(&_L2ExchangeERC20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Decimals() (uint8, error) {
	return _L2ExchangeERC20.Contract.Decimals(&_L2ExchangeERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) Decimals() (uint8, error) {
	return _L2ExchangeERC20.Contract.Decimals(&_L2ExchangeERC20.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) L1Token() (common.Address, error) {
	return _L2ExchangeERC20.Contract.L1Token(&_L2ExchangeERC20.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) L1Token() (common.Address, error) {
	return _L2ExchangeERC20.Contract.L1Token(&_L2ExchangeERC20.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) L2Bridge() (common.Address, error) {
	return _L2ExchangeERC20.Contract.L2Bridge(&_L2ExchangeERC20.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) L2Bridge() (common.Address, error) {
	return _L2ExchangeERC20.Contract.L2Bridge(&_L2ExchangeERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Name() (string, error) {
	return _L2ExchangeERC20.Contract.Name(&_L2ExchangeERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) Name() (string, error) {
	return _L2ExchangeERC20.Contract.Name(&_L2ExchangeERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _L2ExchangeERC20.Contract.Nonces(&_L2ExchangeERC20.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _L2ExchangeERC20.Contract.Nonces(&_L2ExchangeERC20.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2ExchangeERC20.Contract.SupportsInterface(&_L2ExchangeERC20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2ExchangeERC20.Contract.SupportsInterface(&_L2ExchangeERC20.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Symbol() (string, error) {
	return _L2ExchangeERC20.Contract.Symbol(&_L2ExchangeERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) Symbol() (string, error) {
	return _L2ExchangeERC20.Contract.Symbol(&_L2ExchangeERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ExchangeERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) TotalSupply() (*big.Int, error) {
	return _L2ExchangeERC20.Contract.TotalSupply(&_L2ExchangeERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2ExchangeERC20 *L2ExchangeERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _L2ExchangeERC20.Contract.TotalSupply(&_L2ExchangeERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Approve(&_L2ExchangeERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Approve(&_L2ExchangeERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_L2ExchangeERC20 *L2ExchangeERC20Transactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Burn(&_L2ExchangeERC20.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_L2ExchangeERC20 *L2ExchangeERC20TransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Burn(&_L2ExchangeERC20.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_L2ExchangeERC20 *L2ExchangeERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Mint(&_L2ExchangeERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_L2ExchangeERC20 *L2ExchangeERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Mint(&_L2ExchangeERC20.TransactOpts, to, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_L2ExchangeERC20 *L2ExchangeERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _L2ExchangeERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Permit(&_L2ExchangeERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_L2ExchangeERC20 *L2ExchangeERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Permit(&_L2ExchangeERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Transfer(&_L2ExchangeERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.Transfer(&_L2ExchangeERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.TransferFrom(&_L2ExchangeERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_L2ExchangeERC20 *L2ExchangeERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2ExchangeERC20.Contract.TransferFrom(&_L2ExchangeERC20.TransactOpts, from, to, amount)
}

// L2ExchangeERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the L2ExchangeERC20 contract.
type L2ExchangeERC20ApprovalIterator struct {
	Event *L2ExchangeERC20Approval // Event containing the contract specifics and raw log

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
func (it *L2ExchangeERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ExchangeERC20Approval)
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
		it.Event = new(L2ExchangeERC20Approval)
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
func (it *L2ExchangeERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ExchangeERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ExchangeERC20Approval represents a Approval event raised by the L2ExchangeERC20 contract.
type L2ExchangeERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*L2ExchangeERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L2ExchangeERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &L2ExchangeERC20ApprovalIterator{contract: _L2ExchangeERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *L2ExchangeERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L2ExchangeERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ExchangeERC20Approval)
				if err := _L2ExchangeERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) ParseApproval(log types.Log) (*L2ExchangeERC20Approval, error) {
	event := new(L2ExchangeERC20Approval)
	if err := _L2ExchangeERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ExchangeERC20BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the L2ExchangeERC20 contract.
type L2ExchangeERC20BurnIterator struct {
	Event *L2ExchangeERC20Burn // Event containing the contract specifics and raw log

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
func (it *L2ExchangeERC20BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ExchangeERC20Burn)
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
		it.Event = new(L2ExchangeERC20Burn)
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
func (it *L2ExchangeERC20BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ExchangeERC20BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ExchangeERC20Burn represents a Burn event raised by the L2ExchangeERC20 contract.
type L2ExchangeERC20Burn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _account, uint256 _amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) FilterBurn(opts *bind.FilterOpts, _account []common.Address) (*L2ExchangeERC20BurnIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2ExchangeERC20.contract.FilterLogs(opts, "Burn", _accountRule)
	if err != nil {
		return nil, err
	}
	return &L2ExchangeERC20BurnIterator{contract: _L2ExchangeERC20.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _account, uint256 _amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *L2ExchangeERC20Burn, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2ExchangeERC20.contract.WatchLogs(opts, "Burn", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ExchangeERC20Burn)
				if err := _L2ExchangeERC20.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _account, uint256 _amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) ParseBurn(log types.Log) (*L2ExchangeERC20Burn, error) {
	event := new(L2ExchangeERC20Burn)
	if err := _L2ExchangeERC20.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ExchangeERC20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the L2ExchangeERC20 contract.
type L2ExchangeERC20MintIterator struct {
	Event *L2ExchangeERC20Mint // Event containing the contract specifics and raw log

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
func (it *L2ExchangeERC20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ExchangeERC20Mint)
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
		it.Event = new(L2ExchangeERC20Mint)
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
func (it *L2ExchangeERC20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ExchangeERC20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ExchangeERC20Mint represents a Mint event raised by the L2ExchangeERC20 contract.
type L2ExchangeERC20Mint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _account, uint256 _amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) FilterMint(opts *bind.FilterOpts, _account []common.Address) (*L2ExchangeERC20MintIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2ExchangeERC20.contract.FilterLogs(opts, "Mint", _accountRule)
	if err != nil {
		return nil, err
	}
	return &L2ExchangeERC20MintIterator{contract: _L2ExchangeERC20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _account, uint256 _amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *L2ExchangeERC20Mint, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2ExchangeERC20.contract.WatchLogs(opts, "Mint", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ExchangeERC20Mint)
				if err := _L2ExchangeERC20.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _account, uint256 _amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) ParseMint(log types.Log) (*L2ExchangeERC20Mint, error) {
	event := new(L2ExchangeERC20Mint)
	if err := _L2ExchangeERC20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ExchangeERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the L2ExchangeERC20 contract.
type L2ExchangeERC20TransferIterator struct {
	Event *L2ExchangeERC20Transfer // Event containing the contract specifics and raw log

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
func (it *L2ExchangeERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ExchangeERC20Transfer)
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
		it.Event = new(L2ExchangeERC20Transfer)
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
func (it *L2ExchangeERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ExchangeERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ExchangeERC20Transfer represents a Transfer event raised by the L2ExchangeERC20 contract.
type L2ExchangeERC20Transfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2ExchangeERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ExchangeERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2ExchangeERC20TransferIterator{contract: _L2ExchangeERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *L2ExchangeERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2ExchangeERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ExchangeERC20Transfer)
				if err := _L2ExchangeERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_L2ExchangeERC20 *L2ExchangeERC20Filterer) ParseTransfer(log types.Log) (*L2ExchangeERC20Transfer, error) {
	event := new(L2ExchangeERC20Transfer)
	if err := _L2ExchangeERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
