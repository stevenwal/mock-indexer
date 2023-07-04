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

// OptionsOption is an auto generated low-level Go binding around an user-defined struct.
type OptionsOption struct {
	Asset  common.Address
	IsPut  bool
	Strike *big.Int
	Expiry *big.Int
}

// InstrumentsMetaData contains all meta data concerning the Instruments contract.
var InstrumentsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"enumInstrumentsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadAddOption\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"enumInstrumentsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadAddPerpetual\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"enumInstrumentsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadChargeFunding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumInstrumentsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadLoadAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"enumInstrumentsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadSetExpiryPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"enumInstrumentsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadSettleOption\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodOver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotKeeper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accounts\",\"type\":\"address\"}],\"name\":\"AccountsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DisputePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ExpiryPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementFee\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oi\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingRate\",\"type\":\"int256\"}],\"name\":\"Funding\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"InstrumentAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"openInterest\",\"type\":\"uint256\"}],\"name\":\"InstrumentTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"KeeperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oi\",\"type\":\"uint256\"}],\"name\":\"OpenInterestUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPut\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"}],\"name\":\"OptionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rate\",\"type\":\"int256\"}],\"name\":\"Payout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"PerpetualAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oi\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Settlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"VersionInitialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isPut\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structOptions.Option\",\"name\":\"option\",\"type\":\"tuple\"}],\"name\":\"addOption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"addPerpetual\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorities\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"fundingRate\",\"type\":\"int256\"}],\"name\":\"chargeFunding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"fundingRate\",\"type\":\"int256\"}],\"name\":\"chargeFunding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"disputePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"disputePeriodOver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"epochInstrument\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"expiryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"getInstrumentAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"instrumentAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"instruments\",\"outputs\":[{\"internalType\":\"enumInstrumentsInterface.InstrumentType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keepers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Balance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Position[]\",\"name\":\"positions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structAccountsInterface.SigningKey[]\",\"name\":\"signingKeys\",\"type\":\"tuple[]\"}],\"name\":\"loadAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oi\",\"type\":\"uint256\"}],\"name\":\"loadOpenInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isPut\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structOptions.Option\",\"name\":\"option\",\"type\":\"tuple\"}],\"name\":\"loadOption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"loadPerpetual\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"options\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isPut\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"overrideExpiryPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"perpetuals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accounts\",\"type\":\"address\"}],\"name\":\"setAccounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setDisputePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setExpiryPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementFee\",\"type\":\"uint256\"}],\"name\":\"settleOption\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementFee\",\"type\":\"uint256\"}],\"name\":\"settleOption\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"settled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"settledInstrumentAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"settledLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"settling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"totalOptionPayout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"transferInstrument\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumInstrumentsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161513961004c60003960008181610a7201528181610abb01528181610b5701528181610b970152610c2a01526151396000f3fe60806040526004361061025c5760003560e01c80636cd22eaf11610144578063bda9e302116100b6578063dee52b641161007a578063dee52b6414610887578063dfac3836146108db578063e6a97e44146108fb578063ebfb62b31461091b578063ee2061691461093b578063ee5314091461097657600080fd5b8063bda9e302146107d7578063c4d66de8146107f7578063c7f9df8b14610817578063d1d49e7c14610847578063d30578771461086757600080fd5b80638f5cafcc116101085780638f5cafcc146106f857806391223d6914610718578063991f94a0146107485780639d56921f146107775780639d5bb5c014610797578063a557b067146107b757600080fd5b80636cd22eaf1461064b57806388e53ec81461066b5780638b3cddaf146106985780638cffa230146106b85780638da5cb5b146106d857600080fd5b80633bbd64bc116101dd57806353f9ead2116101a157806353f9ead21461055d57806359a1ecaa1461057d5780635a0eb264146105cb5780635af9ac0a146105eb5780635f92e55f1461060b57806368cd03f61461062b57600080fd5b80633bbd64bc1461045d578063409e22051461048d5780634e1b7775146105085780634f1ef2861461053557806352d1902d1461054857600080fd5b806319b57f9b1161022457806319b57f9b14610358578063205e4373146103a357806322bbad0b146103d0578063283457801461040d5780633659cfe61461043d57600080fd5b806304a34c2a14610261578063059f6b071461029757806306fdde03146102d2578063082743821461031657806313af403514610338575b600080fd5b34801561026d57600080fd5b5061028161027c36600461444d565b610996565b60405161028e9190614466565b60405180910390f35b3480156102a357600080fd5b506102c46102b23660046144c8565b60066020526000908152604090205481565b60405190815260200161028e565b3480156102de57600080fd5b506103096040518060400160405280600b81526020016a496e737472756d656e747360a81b81525081565b60405161028e9190614511565b34801561032257600080fd5b506103366103313660046144c8565b610a02565b005b34801561034457600080fd5b506103366103533660046144c8565b610a54565b34801561036457600080fd5b50610393610373366004614544565b600b60209081526000928352604080842090915290825290205460ff1681565b604051901515815260200161028e565b3480156103af57600080fd5b506102c46103be36600461444d565b600e6020526000908152604090205481565b3480156103dc57600080fd5b506104006103eb36600461444d565b60076020526000908152604090205460ff1681565b60405161028e919061458a565b34801561041957600080fd5b5061039361042836600461444d565b600d6020526000908152604090205460ff1681565b34801561044957600080fd5b506103366104583660046144c8565b610a68565b34801561046957600080fd5b506103936104783660046144c8565b60046020526000908152604090205460ff1681565b34801561049957600080fd5b506104de6104a836600461444d565b6010602052600090815260409020805460018201546002909201546001600160a01b03821692600160a01b90920460ff16919084565b604080516001600160a01b039095168552921515602085015291830152606082015260800161028e565b34801561051457600080fd5b506102c461052336600461444d565b60126020526000908152604090205481565b610336610543366004614637565b610b4d565b34801561055457600080fd5b506102c4610c1d565b34801561056957600080fd5b50610336610578366004614544565b610cd0565b34801561058957600080fd5b506105b361059836600461444d565b6011602052600090815260409020546001600160a01b031681565b6040516001600160a01b03909116815260200161028e565b3480156105d757600080fd5b506102c46105e63660046146df565b610e25565b3480156105f757600080fd5b506102c461060636600461444d565b610e54565b34801561061757600080fd5b50610336610626366004614701565b610f90565b34801561063757600080fd5b506005546105b3906001600160a01b031681565b34801561065757600080fd5b50610336610666366004614747565b6111aa565b34801561067757600080fd5b506102c461068636600461444d565b60086020526000908152604090205481565b3480156106a457600080fd5b506103366106b3366004614775565b61122d565b3480156106c457600080fd5b506103936106d33660046147a1565b61127c565b3480156106e457600080fd5b506000546105b3906001600160a01b031681565b34801561070457600080fd5b506103366107133660046146df565b61171f565b34801561072457600080fd5b506103936107333660046144c8565b60036020526000908152604090205460ff1681565b34801561075457600080fd5b506107686107633660046147dc565b6117a0565b60405161028e93929190614842565b34801561078357600080fd5b506103366107923660046149ad565b6119c9565b3480156107a357600080fd5b506103936107b23660046146df565b61253d565b3480156107c357600080fd5b506103936107d2366004614acd565b612634565b3480156107e357600080fd5b506103366107f2366004614aff565b612647565b34801561080357600080fd5b506103366108123660046144c8565b612712565b34801561082357600080fd5b5061039361083236600461444d565b600c6020526000908152604090205460ff1681565b34801561085357600080fd5b50610393610862366004614b34565b61278a565b34801561087357600080fd5b50610336610882366004614747565b612c1f565b34801561089357600080fd5b506108c66108a2366004614775565b60096020908152600092835260408084209091529082529020805460019091015482565b6040805192835260208301919091520161028e565b3480156108e757600080fd5b506103936108f6366004614775565b612ca2565b34801561090757600080fd5b50610336610916366004614701565b612d07565b34801561092757600080fd5b50610336610936366004614544565b612e75565b34801561094757600080fd5b50610393610956366004614544565b600f60209081526000928352604080842090915290825290205460ff1681565b34801561098257600080fd5b50610336610991366004614aff565b612ff9565b6000818152600a60209081526040918290208054835181840281018401909452808452606093928301828280156109f657602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116109d8575b50505050509050919050565b610a0a6130a4565b600580546001600160a01b0319166001600160a01b0383169081179091556040517fbecde7fe690c73ba54232f00eb06c31464f65b45aff18984febaa80df22dcb8d90600090a250565b610a5c6130a4565b610a65816130d0565b50565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610ab95760405162461bcd60e51b8152600401610ab090614b77565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610b026000805160206150e6833981519152546001600160a01b031690565b6001600160a01b031614610b285760405162461bcd60e51b8152600401610ab090614bc3565b610b318161313f565b60408051600080825260208201909252610a6591839190613147565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610b955760405162461bcd60e51b8152600401610ab090614b77565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610bde6000805160206150e6833981519152546001600160a01b031690565b6001600160a01b031614610c045760405162461bcd60e51b8152600401610ab090614bc3565b610c0d8261313f565b610c1982826001613147565b5050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cbd5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610ab0565b506000805160206150e683398151915290565b3360009081526004602052604090205460ff16610d0057604051631ea2564f60e31b815260040160405180910390fd5b6000828152601160205260409020546001600160a01b0390811690821603610d405781600560405163e2a5e23160e01b8152600401610ab0929190614c0f565b6000828152600760209081526040808320805460ff1916600217905560119091529081902080546001600160a01b038481166001600160a01b031990921691909117909155600554915163474c8b0360e11b815260048101859052911690638e991606906024016020604051808303816000875af1158015610dc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dea9190614c32565b506040516001600160a01b0382169083907f65f3016ffa801ac65a87ed4a8c83abfe84ba1f2ffd38229df8912360535a033990600090a35050565b604080516020808201859052818301849052825180830384018152606090920190925280519101205b92915050565b6000600160008381526007602052604090205460ff166002811115610e7b57610e7b614574565b14610e8857506000919050565b6000828152601060209081526040808320815160808101835281546001600160a01b038116808352600160a01b90910460ff1615158286015260018301548285015260029092015460608201908152918552600984528285209151855292528220549091819003610efd575060009392505050565b600084815260086020908152604091829020546005548351633fd1e2bd60e01b81529351610f8894929386936001600160a01b0390931692633fd1e2bd92600480830193928290030181865afa158015610f5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f7f9190614c4d565b859291906132b7565b949350505050565b3360009081526004602052604090205460ff16610fc057604051631ea2564f60e31b815260040160405180910390fd5b600082815260106020908152604091829020825160808101845281546001600160a01b0381168252600160a01b900460ff16151581840152600182015493810193909352600201546060830152611019908301836144c8565b6001600160a01b031681600001516001600160a01b031614801561105357506110486040830160208401614c70565b151581602001511515145b8015611066575081604001358160400151145b8015611079575081606001358160600151145b1561109c57826005604051631e23755760e11b8152600401610ab0929190614c0f565b6000838152600760209081526040808320805460ff191660011790556010909152902082906110cb8282614c8d565b505060055460405163474c8b0360e11b8152600481018590526001600160a01b0390911690638e991606906024016020604051808303816000875af1158015611118573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061113c9190614c32565b50606082013561114f60208401846144c8565b6001600160a01b0316847f8a3e5ee639fc2822a7d3969b04dabf44aa1249f1565d8b09b4940c64a1fa24d661118a6040870160208801614c70565b6040805191151582528088013560208301520160405180910390a4505050565b6111b26130a4565b6001600160a01b0382166111d95760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff191685151590811790915590519092917fc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc91a35050565b6112356130a4565b6001600160a01b038216600081815260066020526040808220849055518392917f6650c9b75cf9b0e48611d4e590a8bd06f11be77c6b005b76114d42955d95e3a791a35050565b604080516101a0810182526000610120820181815261014083018290526101608301829052610180830182905282526020808301829052828401829052606083018290526080830182905260a0830182905260c0830182905260e083018290526101008301829052338252600490529182205460ff16611316578685600460405163153365cf60e01b8152600401610ab093929190614842565b6000198403611331576000878152600a602052604090205493505b6005546001600160a01b031660208083018290526040805163999b93af60e01b8152905163999b93af926004808401939192918290030181865afa15801561137d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a19190614cfb565b6001600160a01b0390811660408381019190915260608084018a905260808085018a905260a0850189905260008b8152600a60209081528482205461010088015260c087018981528d835260108252918590208551938401865280549687168452600160a01b90960460ff161515908301526001850154938201939093526002909301549083015290825251620f4240101561145f5780606001518160a00151600960405163153365cf60e01b8152600401610ab093929190614842565b6001606082015160009081526007602052604090205460ff16600281111561148957611489614574565b146114b65780606001518160a00151600560405163153365cf60e01b8152600401610ab093929190614842565b60608101516000908152600d602052604090205460ff16156114fa5780606001518160a00151601060405163153365cf60e01b8152600401610ab093929190614842565b60608101516000908152600c602052604090205460ff166115345760608101516000908152600c60205260409020805460ff191660011790555b8061010001516000036115635761155981606001518260a00151888460c001516132ea565b6001915050611716565b6115db620f42408783602001516001600160a01b0316633fd1e2bd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d19190614c4d565b84519291906132b7565b60e082018190526000036116015761155981606001518260a00151888460c001516132ea565b845b848110156116dd5760608201516000908152600a6020526040812080548390811061163057611630614d18565b600091825260209182902001549084015160608501516040516349ad755b60e01b81526001600160a01b039384166004820181905260248201929092529093506116ca92869285929116906349ad755b90604401602060405180830381865afa1580156116a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116c59190614d2e565b61340e565b50806116d581614d5d565b915050611603565b506101008101516000888152600e6020526040902054036117105761155981606001518260a00151888460c001516132ea565b60009150505b95945050505050565b3360009081526004602052604090205460ff1661174f57604051631ea2564f60e31b815260040160405180910390fd5b600082815260086020526040908190208290555182907fec9d3d15d27b45ec272eca821314dbc7dfac39e122618a8cbfb23ca99d19c883906117949084815260200190565b60405180910390a25050565b336000908152600360205260408120548190819060ff166117d457604051631890934360e11b815260040160405180910390fd5b6117dc613867565b856001600160a01b0316876001600160a01b031603611803575060009150819050806119b6565b6005546001600160a01b0316600061181d828b8b8b613890565b9050600080836001600160a01b031663ac759c718c8e8c8c6040518563ffffffff1660e01b81526004016118549493929190614d76565b6060604051808303816000875af1158015611873573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118979190614d9c565b90935091506000905081600d8111156118b2576118b2614574565b146118d2576000806118c3836139b6565b965096509650505050506119b6565b506118dd8b8b613a69565b600080846001600160a01b0316638869a3f28c8f8d8d6040518563ffffffff1660e01b81526004016119129493929190614d76565b6060604051808303816000875af1158015611931573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119559190614d9c565b90935091506000905081600d81111561197057611970614574565b1461199157600080611981836139b6565b97509750975050505050506119b6565b5061199c8c8b613a69565b6119aa848d8d8d8d88613b1b565b90955093506000925050505b6001600281905550955095509592505050565b3360009081526004602052604090205460ff166119f957604051631ea2564f60e31b815260040160405180910390fd5b60055460405163744d870f60e11b81526001600160a01b03868116600483015290911690819063e89b0e1e906024016020604051808303816000875af1158015611a47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a6b9190614dd1565b506000805b8551811015611f3657826001600160a01b031663eeb97d3b878381518110611a9a57611a9a614d18565b6020908102919091010151516040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015611aea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b0e9190614dd1565b611b2e57600b60405163eb1b129160e01b8152600401610ab09190614dee565b6000836001600160a01b031663b7d5820b89898581518110611b5257611b52614d18565b6020908102919091010151516040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604401602060405180830381865afa158015611ba9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bcd9190614d2e565b90506000811315611c7057836001600160a01b031663262709e689898581518110611bfa57611bfa614d18565b602002602001015160000151846040518463ffffffff1660e01b8152600401611c2593929190614dfc565b60408051808303816000875af1158015611c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c679190614e20565b9350611d159050565b6000811215611d1557836001600160a01b0316637de182c589898581518110611c9b57611c9b614d18565b60200260200101516000015184611cb190614e4c565b6040518463ffffffff1660e01b8152600401611ccf93929190614dfc565b60408051808303816000875af1158015611ced573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d119190614e20565b9350505b600083600d811115611d2957611d29614574565b14611d5157611d37836139b6565b60405163eb1b129160e01b8152600401610ab09190614dee565b6000878381518110611d6557611d65614d18565b6020026020010151602001511315611e2c57836001600160a01b0316637de182c589898581518110611d9957611d99614d18565b6020026020010151600001518a8681518110611db757611db7614d18565b6020026020010151602001516040518463ffffffff1660e01b8152600401611de193929190614dfc565b60408051808303816000875af1158015611dff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e239190614e20565b9350611f0b9050565b6000878381518110611e4057611e40614d18565b6020026020010151602001511215611f0b57836001600160a01b031663262709e689898581518110611e7457611e74614d18565b6020026020010151600001518a8681518110611e9257611e92614d18565b602002602001015160200151611ea790614e4c565b6040518463ffffffff1660e01b8152600401611ec593929190614dfc565b60408051808303816000875af1158015611ee3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f079190614e20565b9350505b600083600d811115611f1f57611f1f614574565b14611f2d57611d37836139b6565b50600101611a70565b5060005b845181101561244e57826001600160a01b03166322bbad0b868381518110611f6457611f64614d18565b6020026020010151600001516040518263ffffffff1660e01b8152600401611f8e91815260200190565b602060405180830381865afa158015611fab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fcf9190614dd1565b611fef57600560405163eb1b129160e01b8152600401610ab09190614dee565b6000836001600160a01b03166349ad755b8988858151811061201357612013614d18565b6020908102919091010151516040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401602060405180830381865afa158015612068573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061208c9190614d2e565b9050600081131561213457836001600160a01b031663ac759c71898885815181106120b9576120b9614d18565b6020026020010151600001518460006040518563ffffffff1660e01b81526004016120e79493929190614d76565b6060604051808303816000875af1158015612106573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061212a9190614d9c565b94506121d5915050565b836001600160a01b0316638869a3f28988858151811061215657612156614d18565b6020026020010151600001518461216c90614e4c565b60006040518563ffffffff1660e01b815260040161218d9493929190614d76565b6060604051808303816000875af11580156121ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121d09190614d9c565b945050505b60008683815181106121e9576121e9614d18565b60200260200101516040015113156122d157836001600160a01b0316638869a3f28988858151811061221d5761221d614d18565b60200260200101516000015189868151811061223b5761223b614d18565b6020026020010151604001518a878151811061225957612259614d18565b6020026020010151602001516040518563ffffffff1660e01b81526004016122849493929190614d76565b6060604051808303816000875af11580156122a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122c79190614d9c565b94506123ab915050565b836001600160a01b031663ac759c71898885815181106122f3576122f3614d18565b60200260200101516000015189868151811061231157612311614d18565b60200260200101516040015161232690614e4c565b8a878151811061233857612338614d18565b6020026020010151602001516040518563ffffffff1660e01b81526004016123639493929190614d76565b6060604051808303816000875af1158015612382573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123a69190614d9c565b945050505b600083600d8111156123bf576123bf614574565b146123cd57611d37836139b6565b8582815181106123df576123df614d18565b60200260200101516040015160000361241e5761241986838151811061240757612407614d18565b60200260200101516000015189613bdc565b612445565b61244586838151811061243357612433614d18565b60200260200101516000015189613a69565b50600101611f3a565b5060005b835181101561253457826001600160a01b031663dfa8ea458886848151811061247d5761247d614d18565b60200260200101516000015187858151811061249b5761249b614d18565b6020026020010151602001516040518463ffffffff1660e01b81526004016124c593929190614dfc565b6020604051808303816000875af11580156124e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125089190614c32565b9150600282600d81111561251e5761251e614574565b0361252c57611d37826139b6565b600101612452565b50505050505050565b6000828152601060209081526040808320815160808101835281546001600160a01b038116808352600160a01b90910460ff16151582860152600180840154838601526002909301546060830190815290865260098552838620905186528452828520835180850190945280548452909101549282018390529183036125de57846000600660405163153365cf60e01b8152600401610ab093929190614842565b6125f082600001518360600151612ca2565b61261557846000600760405163153365cf60e01b8152600401610ab093929190614842565b80516000868152600a602052604081205461171692889290918861127c565b600061171685858560006000198761278a565b61264f6130a4565b428211156126705760405163d0404f8560e01b815260040160405180910390fd5b61267a8383612ca2565b1561269857604051631e3b157760e01b815260040160405180910390fd5b6040805180820182528281524260208083019182526001600160a01b038716600081815260098352858120888252835285902093518455915160019093019290925591518381528492917f8cc2700c35138467f8087f0f5101b52a163c8554838ef019be2546e67eebb9c7910160405180910390a3505050565b600161271e8180614e68565b6001541461273f576040516302ed543d60e51b815260040160405180910390fd5b8060010361274d5760016002555b612756826130d0565b600181905560405181907f7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d90600090a25050565b60006127f460405180610140016040528060006001600160a01b0316815260200160006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b3360009081526004602052604090205460ff1661282d57878786600460405163cf5c195f60e01b8152600401610ab09493929190614e7f565b6005546001600160a01b03168082526040805163999b93af60e01b8152905163999b93af916004808201926020929091908290030181865afa158015612877573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061289b9190614cfb565b6001600160a01b0316602082015260408101889052606081018790526128c18888610e25565b60a082015260c081018590526080810186905260e08101839052620f42406128e98488614ea1565b6128f39190614f3c565b61010082015260a08101516000908152601260205260408120549003612936576000888152600a602090815260408083205460a085015184526012909252909120555b60a08101516000908152601260205260409020546101208201526001840161296f5760a081015160009081526012602052604090205493505b60008160e001511380156129895750620f42408160e00151135b806129b1575060008160e001511280156129b157506129aa620f4240614e4c565b8160e00151125b156129e257604080820151606083015160c0840151925163cf5c195f60e01b8152610ab09390601190600401614e7f565b600260408083015160009081526007602052205460ff166002811115612a0a57612a0a614574565b14612a3b57604080820151606083015160c0840151925163cf5c195f60e01b8152610ab09390600590600401614e7f565b60a08101516000908152600d602052604090205460ff1615612a8357604080820151606083015160c0840151925163cf5c195f60e01b8152610ab09390601090600401614e7f565b60a08101516000908152600c602052604090205460ff16612abd5760a08101516000908152600c60205260409020805460ff191660011790555b806101000151600003612af557612aeb816040015182606001518360a0015184608001518560e00151613c35565b6001915050612c15565b845b84811015612bce576040808301516000908152600a602052908120805483908110612b2457612b24614d18565b600091825260209091200154835160408086015190516349ad755b60e01b81526001600160a01b03938416600482018190526024820192909252909350612bbb92869285929116906349ad755b90604401602060405180830381865afa158015612b92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bb69190614d2e565b613c9c565b5080612bc681614d5d565b915050612af7565b5061012081015160a08201516000908152600e602052604090205403612c0f57612aeb816040015182606001518360a0015184608001518560e00151613c35565b60009150505b9695505050505050565b612c276130a4565b6001600160a01b038216612c4e5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260046020526040808220805460ff191685151590811790915590519092917f786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c0091a35050565b6001600160a01b0382166000908152600960209081526040808320848452909152812060010154808203612cda576000915050610e4e565b6001600160a01b038416600090815260066020526040902054612cfd9082614f6a565b4211949350505050565b3360009081526004602052604090205460ff16612d3c57816004604051631e23755760e11b8152600401610ab0929190614c0f565b60008281526007602052604081205460ff166002811115612d5f57612d5f614574565b14612d8257816005604051631e23755760e11b8152600401610ab0929190614c0f565b6000828152600760209081526040808320805460ff19166001179055601090915290208190612db18282614c8d565b505060055460405163474c8b0360e11b8152600481018490526000916001600160a01b031690638e991606906024016020604051808303816000875af1158015612dff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e239190614c32565b9050600081600d811115612e3957612e39614574565b14612e635782612e48826139b6565b604051631e23755760e11b8152600401610ab0929190614c0f565b606082013561114f60208401846144c8565b3360009081526004602052604090205460ff16612eaa5781600460405163e2a5e23160e01b8152600401610ab0929190614c0f565b60008281526007602052604081205460ff166002811115612ecd57612ecd614574565b14612ef05781600560405163e2a5e23160e01b8152600401610ab0929190614c0f565b6000828152600760209081526040808320805460ff19166002179055601190915280822080546001600160a01b038581166001600160a01b031990921691909117909155600554915163474c8b0360e11b815260048101869052911690638e991606906024016020604051808303816000875af1158015612f75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f999190614c32565b9050600081600d811115612faf57612faf614574565b14612fbe5782612e48826139b6565b6040516001600160a01b0383169084907f65f3016ffa801ac65a87ed4a8c83abfe84ba1f2ffd38229df8912360535a033990600090a3505050565b3360009081526004602052604090205460ff166130305782826004604051638ccd20bb60e01b8152600401610ab093929190614f82565b428211156130585782826006604051638ccd20bb60e01b8152600401610ab093929190614f82565b6001600160a01b0383166000908152600960209081526040808320858452909152902060010154156126985782826008604051638ccd20bb60e01b8152600401610ab093929190614f82565b6000546001600160a01b031633146130ce576040516282b42960e81b815260040160405180910390fd5b565b6001600160a01b0381166130f75760405163d92e233d60e01b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b038316908117825560405190917f4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b91a250565b610a656130a4565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561317f5761317a83613f5c565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156131d9575060408051601f3d908101601f191682019092526131d691810190614d2e565b60015b61323c5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610ab0565b6000805160206150e683398151915281146132ab5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610ab0565b5061317a838383613ff8565b60008460200151156132da576132d3856040015184868561401d565b9050610f88565b6132d385604001518486856140d1565b6000848152600d6020526040808220805460ff1916600117905560055490516338c3a0e160e21b8152600481018790526001600160a01b039091169063e30e8384906024016020604051808303816000875af115801561334e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133729190614c32565b9050600081600d81111561338857613388614574565b146133b4578484613398836139b6565b60405163153365cf60e01b8152600401610ab093929190614842565b60008581526008602090815260409182902054825190815290810185905290810183905285907f0f3455c580c87ae485cf1fd9108e0d31f59b529eb91f338403886bd23c9d96ac9060600160405180910390a25050505050565b60608301516000908152600f602090815260408083206001600160a01b038616845290915290205460ff161561344357505050565b60608301516000908152600e60205260408120805460019290613467908490614f6a565b909155505060608301516000908152600f602090815260408083206001600160a01b03861684529091529020805460ff19166001179055801561317a5760208301516060840151604051639aeddeff60e01b81526000926001600160a01b031691639aeddeff916134f09187916004016001600160a01b03929092168252602082015260400190565b6020604051808303816000875af115801561350f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135339190614c32565b9050600081600d81111561354957613549614574565b146135615783606001518460a00151613398836139b6565b6000821315613709576000620f42408560e00151846135809190614fa6565b61358a9190614fc5565b905060006135a7866060015186848960c001518a6040015161415d565b93509050600083600d8111156135bf576135bf614574565b146135d75785606001518660a00151613398856139b6565b85602001516001600160a01b0316637de182c586886040015184866135fc9190614e68565b6040518463ffffffff1660e01b815260040161361a93929190614dfc565b60408051808303816000875af1158015613638573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061365c9190614e20565b93506000905083600d81111561367457613674614574565b1461368c5785606001518660a00151613398856139b6565b60608087015187518201516080808a015160e08b0151604080519485526020850192909252908301899052938201869052810184905260a08101929092526001600160a01b038716917f9564f1ed9c02902fc9b9d1d3a30c1bb7320ac24c0d4454bcb53ec71f67f292649060c00160405180910390a35050613861565b6000620f42408560e001518461371e90614e4c565b6137289190614fa6565b6137329190614fc5565b905084602001516001600160a01b031663262709e6858760400151846040518463ffffffff1660e01b815260040161376c93929190614dfc565b60408051808303816000875af115801561378a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137ae9190614e20565b92506000905082600d8111156137c6576137c6614574565b146137de5784606001518560a00151613398846139b6565b836001600160a01b031685606001517f9564f1ed9c02902fc9b9d1d3a30c1bb7320ac24c0d4454bcb53ec71f67f292648760000151606001518860800151878661382790614e4c565b60e08c01516040805195865260208601949094529284019190915260608301526000608083015260a082015260c0015b60405180910390a3505b50505050565b60025460011461388a5760405163558a1e0360e11b815260040160405180910390fd5b60028055565b6040516349ad755b60e01b81526001600160a01b0383811660048301526024820185905260009182918716906349ad755b90604401602060405180830381865afa1580156138e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139069190614d2e565b6040516349ad755b60e01b81526001600160a01b038581166004830152602482018890529192506000918816906349ad755b90604401602060405180830381865afa158015613959573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061397d9190614d2e565b90506000811361398e576000613990565b805b6000831361399f5760006139a1565b825b6139ab9190614f6a565b979650505050505050565b60008082600d8111156139cb576139cb614574565b036139d857506000919050565b600582600d8111156139ec576139ec614574565b036139f957506002919050565b600282600d811115613a0d57613a0d614574565b03613a1a57506003919050565b600b82600d811115613a2e57613a2e614574565b03613a3b5750600e919050565b600782600d811115613a4f57613a4f614574565b03613a5c5750600f919050565b506001919050565b919050565b6000828152600b602090815260408083206001600160a01b038516845290915290205460ff16610c19576000828152600a602090815260408083208054600180820183559185528385200180546001600160a01b0319166001600160a01b038716908117909155868552600b8452828520818652909352818420805460ff1916909117905551909184917fd56cbb684c2af0cb6004dfc76579372e9989c539a6140d9622ab4e4264f1f75f9190a35050565b6000613b2987878787613890565b60008781526008602052604090205490915082821015613b5e57613b4d8284614e68565b613b579082614e68565b9050613b75565b613b688383614e68565b613b729082614f6a565b90505b60008781526008602090815260409182902083905581518681529081018390526001600160a01b0380881692908916918a917fea4600feed794ec67dfe54a044b2501a9ec05001c578ee10a1f347d26d8ef46c910160405180910390a45050505050505050565b6000828152600b602090815260408083206001600160a01b038516845290915290205460ff1615610c19576000828152600b602090815260408083206001600160a01b03851684529091529020805460ff191690555050565b6000838152600d60209081526040808320805460ff191660011790558783526008825291829020548251908152908101849052908101829052849086907fdf8220724a6098d8315f430588d7692863861c4805f1006a31d7c5bbabc4842690606001613857565b60a08301516000908152600f602090815260408083206001600160a01b038616845290915290205460ff1615613cd157505050565b60a08301516000908152600e60205260408120805460019290613cf5908490614f6a565b909155505060a08301516000908152600f602090815260408083206001600160a01b03861684529091529020805460ff19166001179055801561317a576000620f424084610100015183613d499190614ea1565b613d5290614e4c565b613d5c9190614f3c565b90506000811315613e165783516020850151604051637de182c560e01b81526000926001600160a01b031691637de182c591613d9e9188918790600401614dfc565b60408051808303816000875af1158015613dbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613de09190614e20565b91506000905081600d811115613df857613df8614574565b14613e105784604001518560c00151613398836139b6565b50613ed6565b6000811215613ed657600084600001516001600160a01b031663262709e685876020015185613e4490614e4c565b6040518463ffffffff1660e01b8152600401613e6293929190614dfc565b60408051808303816000875af1158015613e80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ea49190614e20565b91506000905081600d811115613ebc57613ebc614574565b14613ed45784604001518560c00151613398836139b6565b505b826001600160a01b031684604001517f9564f1ed9c02902fc9b9d1d3a30c1bb7320ac24c0d4454bcb53ec71f67f2926486606001518760800151868660008b60e00151604051613f4e96959493929190958652602086019490945260408501929092526060840152608083015260a082015260c00190565b60405180910390a350505050565b6001600160a01b0381163b613fc95760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610ab0565b6000805160206150e683398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b614001836142da565b60008251118061400e5750805b1561317a57613861838361431a565b600082158061402c5750848410155b1561403957506000610f88565b614044600680614f6a565b8260ff1610156140905760ff821661405d600680614f6a565b6140679190614e68565b61407290600a6150bd565b61407c8587614e68565b6140869085614fa6565b6132d39190614fc5565b61409b600680614f6a565b6140a89060ff8416614e68565b6140b390600a6150bd565b6140bd8587614e68565b6140c79085614fa6565b6132d39190614fa6565b60008215806140e05750848411155b156140ed57506000610f88565b6140f8600680614f6a565b8260ff1610156141305760ff8216614111600680614f6a565b61411b9190614e68565b61412690600a6150bd565b61407c8686614e68565b61413b600680614f6a565b6141489060ff8416614e68565b61415390600a6150bd565b6140bd8686614e68565b60008080620f424061416f8688614fa6565b6141799190614fc5565b905080156142855760055460408051634614be9f60e01b815290516000926001600160a01b031691637de182c5918391634614be9f9160048083019260209291908290030181865afa1580156141d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141f79190614cfb565b87856040518463ffffffff1660e01b815260040161421793929190614dfc565b60408051808303816000875af1158015614235573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142599190614e20565b91506000905081600d81111561427157614271614574565b14614283576000935091506142d09050565b505b866001600160a01b0316887fc90f6e366e1943d6ee7f9f5d64d27a53bf4a64e438dc757f9b7e1e96b9aaaab2836040516142c191815260200190565b60405180910390a39150600090505b9550959350505050565b6142e381613f5c565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b606061433f838360405180606001604052806027815260200161510660279139614346565b9392505050565b60606001600160a01b0384163b6143ae5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610ab0565b600080856001600160a01b0316856040516143c991906150c9565b600060405180830381855af49150503d8060008114614404576040519150601f19603f3d011682016040523d82523d6000602084013e614409565b606091505b5091509150612c158282866060831561442357508161433f565b8251156144335782518084602001fd5b8160405162461bcd60e51b8152600401610ab09190614511565b60006020828403121561445f57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156144a75783516001600160a01b031683529284019291840191600101614482565b50909695505050505050565b6001600160a01b0381168114610a6557600080fd5b6000602082840312156144da57600080fd5b813561433f816144b3565b60005b838110156145005781810151838201526020016144e8565b838111156138615750506000910152565b60208152600082518060208401526145308160408501602087016144e5565b601f01601f19169190910160400192915050565b6000806040838503121561455757600080fd5b823591506020830135614569816144b3565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b602081016003831061459e5761459e614574565b91905290565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156145dd576145dd6145a4565b60405290565b6040805190810167ffffffffffffffff811182821017156145dd576145dd6145a4565b604051601f8201601f1916810167ffffffffffffffff8111828210171561462f5761462f6145a4565b604052919050565b6000806040838503121561464a57600080fd5b8235614655816144b3565b915060208381013567ffffffffffffffff8082111561467357600080fd5b818601915086601f83011261468757600080fd5b813581811115614699576146996145a4565b6146ab601f8201601f19168501614606565b915080825287848285010111156146c157600080fd5b80848401858401376000848284010152508093505050509250929050565b600080604083850312156146f257600080fd5b50508035926020909101359150565b60008082840360a081121561471557600080fd5b833592506080601f198201121561472b57600080fd5b506020830190509250929050565b8015158114610a6557600080fd5b6000806040838503121561475a57600080fd5b8235614765816144b3565b9150602083013561456981614739565b6000806040838503121561478857600080fd5b8235614793816144b3565b946020939093013593505050565b600080600080600060a086880312156147b957600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b600080600080600060a086880312156147f457600080fd5b853594506020860135614806816144b3565b93506040860135614816816144b3565b94979396509394606081013594506080013592915050565b6012811061483e5761483e614574565b9052565b8381526020810183905260608101610f88604083018461482e565b600067ffffffffffffffff821115614877576148776145a4565b5060051b60200190565b600082601f83011261489257600080fd5b813560206148a76148a28361485d565b614606565b828152606092830285018201928282019190878511156148c657600080fd5b8387015b8581101561490d5781818a0312156148e25760008081fd5b6148ea6145ba565b8135815285820135868201526040808301359082015284529284019281016148ca565b5090979650505050505050565b600082601f83011261492b57600080fd5b8135602061493b6148a28361485d565b82815260069290921b8401810191818101908684111561495a57600080fd5b8286015b848110156149a257604081890312156149775760008081fd5b61497f6145e3565b813561498a816144b3565b8152818501358582015283529183019160400161495e565b509695505050505050565b600080600080608085870312156149c357600080fd5b84356149ce816144b3565b935060208581013567ffffffffffffffff808211156149ec57600080fd5b818801915088601f830112614a0057600080fd5b8135614a0e6148a28261485d565b81815260069190911b8301840190848101908b831115614a2d57600080fd5b938501935b82851015614a79576040858d031215614a4b5760008081fd5b614a536145e3565b8535614a5e816144b3565b81528587013587820152825260409094019390850190614a32565b975050506040880135925080831115614a9157600080fd5b614a9d89848a01614881565b94506060880135925080831115614ab357600080fd5b5050614ac18782880161491a565b91505092959194509250565b60008060008060808587031215614ae357600080fd5b5050823594602084013594506040840135936060013592509050565b600080600060608486031215614b1457600080fd5b8335614b1f816144b3565b95602085013595506040909401359392505050565b60008060008060008060c08789031215614b4d57600080fd5b505084359660208601359650604086013595606081013595506080810135945060a0013592509050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b8281526040810161433f602083018461482e565b8051600e8110613a6457600080fd5b600060208284031215614c4457600080fd5b61433f82614c23565b600060208284031215614c5f57600080fd5b815160ff8116811461433f57600080fd5b600060208284031215614c8257600080fd5b813561433f81614739565b8135614c98816144b3565b81546001600160a01b031981166001600160a01b039290921691821783556020840135614cc481614739565b6001600160a81b03199190911690911790151560a01b60ff60a01b1617815560408201356001820155606090910135600290910155565b600060208284031215614d0d57600080fd5b815161433f816144b3565b634e487b7160e01b600052603260045260246000fd5b600060208284031215614d4057600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600060018201614d6f57614d6f614d47565b5060010190565b6001600160a01b0394909416845260208401929092526040830152606082015260800190565b600080600060608486031215614db157600080fd5b8351925060208401519150614dc860408501614c23565b90509250925092565b600060208284031215614de357600080fd5b815161433f81614739565b60208101610e4e828461482e565b6001600160a01b039384168152919092166020820152604081019190915260600190565b60008060408385031215614e3357600080fd5b82519150614e4360208401614c23565b90509250929050565b6000600160ff1b8201614e6157614e61614d47565b5060000390565b600082821015614e7a57614e7a614d47565b500390565b848152602081018490526040810183905260808101611716606083018461482e565b60006001600160ff1b0381841382841380821686840486111615614ec757614ec7614d47565b600160ff1b6000871282811687830589121615614ee657614ee6614d47565b60008712925087820587128484161615614f0257614f02614d47565b87850587128184161615614f1857614f18614d47565b505050929093029392505050565b634e487b7160e01b600052601260045260246000fd5b600082614f4b57614f4b614f26565b600160ff1b821460001984141615614f6557614f65614d47565b500590565b60008219821115614f7d57614f7d614d47565b500190565b6001600160a01b03841681526020810183905260608101610f88604083018461482e565b6000816000190483118215151615614fc057614fc0614d47565b500290565b600082614fd457614fd4614f26565b500490565b600181815b80851115615014578160001904821115614ffa57614ffa614d47565b8085161561500757918102915b93841c9390800290614fde565b509250929050565b60008261502b57506001610e4e565b8161503857506000610e4e565b816001811461504e576002811461505857615074565b6001915050610e4e565b60ff84111561506957615069614d47565b50506001821b610e4e565b5060208310610133831016604e8410600b8410161715615097575081810a610e4e565b6150a18383614fd9565b80600019048211156150b5576150b5614d47565b029392505050565b600061433f838361501c565b600082516150db8184602087016144e5565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300080f000a",
}

// InstrumentsABI is the input ABI used to generate the binding from.
// Deprecated: Use InstrumentsMetaData.ABI instead.
var InstrumentsABI = InstrumentsMetaData.ABI

// InstrumentsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InstrumentsMetaData.Bin instead.
var InstrumentsBin = InstrumentsMetaData.Bin

// DeployInstruments deploys a new Ethereum contract, binding an instance of Instruments to it.
func DeployInstruments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Instruments, error) {
	parsed, err := InstrumentsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InstrumentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Instruments{InstrumentsCaller: InstrumentsCaller{contract: contract}, InstrumentsTransactor: InstrumentsTransactor{contract: contract}, InstrumentsFilterer: InstrumentsFilterer{contract: contract}}, nil
}

// Instruments is an auto generated Go binding around an Ethereum contract.
type Instruments struct {
	InstrumentsCaller     // Read-only binding to the contract
	InstrumentsTransactor // Write-only binding to the contract
	InstrumentsFilterer   // Log filterer for contract events
}

// InstrumentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InstrumentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstrumentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InstrumentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstrumentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InstrumentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstrumentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InstrumentsSession struct {
	Contract     *Instruments      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InstrumentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InstrumentsCallerSession struct {
	Contract *InstrumentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InstrumentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InstrumentsTransactorSession struct {
	Contract     *InstrumentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InstrumentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InstrumentsRaw struct {
	Contract *Instruments // Generic contract binding to access the raw methods on
}

// InstrumentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InstrumentsCallerRaw struct {
	Contract *InstrumentsCaller // Generic read-only contract binding to access the raw methods on
}

// InstrumentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InstrumentsTransactorRaw struct {
	Contract *InstrumentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInstruments creates a new instance of Instruments, bound to a specific deployed contract.
func NewInstruments(address common.Address, backend bind.ContractBackend) (*Instruments, error) {
	contract, err := bindInstruments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Instruments{InstrumentsCaller: InstrumentsCaller{contract: contract}, InstrumentsTransactor: InstrumentsTransactor{contract: contract}, InstrumentsFilterer: InstrumentsFilterer{contract: contract}}, nil
}

// NewInstrumentsCaller creates a new read-only instance of Instruments, bound to a specific deployed contract.
func NewInstrumentsCaller(address common.Address, caller bind.ContractCaller) (*InstrumentsCaller, error) {
	contract, err := bindInstruments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InstrumentsCaller{contract: contract}, nil
}

// NewInstrumentsTransactor creates a new write-only instance of Instruments, bound to a specific deployed contract.
func NewInstrumentsTransactor(address common.Address, transactor bind.ContractTransactor) (*InstrumentsTransactor, error) {
	contract, err := bindInstruments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InstrumentsTransactor{contract: contract}, nil
}

// NewInstrumentsFilterer creates a new log filterer instance of Instruments, bound to a specific deployed contract.
func NewInstrumentsFilterer(address common.Address, filterer bind.ContractFilterer) (*InstrumentsFilterer, error) {
	contract, err := bindInstruments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InstrumentsFilterer{contract: contract}, nil
}

// bindInstruments binds a generic wrapper to an already deployed contract.
func bindInstruments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InstrumentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Instruments *InstrumentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Instruments.Contract.InstrumentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Instruments *InstrumentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Instruments.Contract.InstrumentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Instruments *InstrumentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Instruments.Contract.InstrumentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Instruments *InstrumentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Instruments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Instruments *InstrumentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Instruments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Instruments *InstrumentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Instruments.Contract.contract.Transact(opts, method, params...)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address)
func (_Instruments *InstrumentsCaller) Accounts(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "accounts")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address)
func (_Instruments *InstrumentsSession) Accounts() (common.Address, error) {
	return _Instruments.Contract.Accounts(&_Instruments.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() view returns(address)
func (_Instruments *InstrumentsCallerSession) Accounts() (common.Address, error) {
	return _Instruments.Contract.Accounts(&_Instruments.CallOpts)
}

// AccountsLength is a free data retrieval call binding the contract method 0x4e1b7775.
//
// Solidity: function accountsLength(uint256 ) view returns(uint256)
func (_Instruments *InstrumentsCaller) AccountsLength(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "accountsLength", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountsLength is a free data retrieval call binding the contract method 0x4e1b7775.
//
// Solidity: function accountsLength(uint256 ) view returns(uint256)
func (_Instruments *InstrumentsSession) AccountsLength(arg0 *big.Int) (*big.Int, error) {
	return _Instruments.Contract.AccountsLength(&_Instruments.CallOpts, arg0)
}

// AccountsLength is a free data retrieval call binding the contract method 0x4e1b7775.
//
// Solidity: function accountsLength(uint256 ) view returns(uint256)
func (_Instruments *InstrumentsCallerSession) AccountsLength(arg0 *big.Int) (*big.Int, error) {
	return _Instruments.Contract.AccountsLength(&_Instruments.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x91223d69.
//
// Solidity: function authorities(address ) view returns(bool)
func (_Instruments *InstrumentsCaller) Authorities(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "authorities", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Authorities is a free data retrieval call binding the contract method 0x91223d69.
//
// Solidity: function authorities(address ) view returns(bool)
func (_Instruments *InstrumentsSession) Authorities(arg0 common.Address) (bool, error) {
	return _Instruments.Contract.Authorities(&_Instruments.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x91223d69.
//
// Solidity: function authorities(address ) view returns(bool)
func (_Instruments *InstrumentsCallerSession) Authorities(arg0 common.Address) (bool, error) {
	return _Instruments.Contract.Authorities(&_Instruments.CallOpts, arg0)
}

// DisputePeriod is a free data retrieval call binding the contract method 0x059f6b07.
//
// Solidity: function disputePeriod(address ) view returns(uint256)
func (_Instruments *InstrumentsCaller) DisputePeriod(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "disputePeriod", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputePeriod is a free data retrieval call binding the contract method 0x059f6b07.
//
// Solidity: function disputePeriod(address ) view returns(uint256)
func (_Instruments *InstrumentsSession) DisputePeriod(arg0 common.Address) (*big.Int, error) {
	return _Instruments.Contract.DisputePeriod(&_Instruments.CallOpts, arg0)
}

// DisputePeriod is a free data retrieval call binding the contract method 0x059f6b07.
//
// Solidity: function disputePeriod(address ) view returns(uint256)
func (_Instruments *InstrumentsCallerSession) DisputePeriod(arg0 common.Address) (*big.Int, error) {
	return _Instruments.Contract.DisputePeriod(&_Instruments.CallOpts, arg0)
}

// DisputePeriodOver is a free data retrieval call binding the contract method 0xdfac3836.
//
// Solidity: function disputePeriodOver(address asset, uint256 expiry) view returns(bool)
func (_Instruments *InstrumentsCaller) DisputePeriodOver(opts *bind.CallOpts, asset common.Address, expiry *big.Int) (bool, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "disputePeriodOver", asset, expiry)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DisputePeriodOver is a free data retrieval call binding the contract method 0xdfac3836.
//
// Solidity: function disputePeriodOver(address asset, uint256 expiry) view returns(bool)
func (_Instruments *InstrumentsSession) DisputePeriodOver(asset common.Address, expiry *big.Int) (bool, error) {
	return _Instruments.Contract.DisputePeriodOver(&_Instruments.CallOpts, asset, expiry)
}

// DisputePeriodOver is a free data retrieval call binding the contract method 0xdfac3836.
//
// Solidity: function disputePeriodOver(address asset, uint256 expiry) view returns(bool)
func (_Instruments *InstrumentsCallerSession) DisputePeriodOver(asset common.Address, expiry *big.Int) (bool, error) {
	return _Instruments.Contract.DisputePeriodOver(&_Instruments.CallOpts, asset, expiry)
}

// EpochInstrument is a free data retrieval call binding the contract method 0x5a0eb264.
//
// Solidity: function epochInstrument(uint256 instrument, uint256 epoch) pure returns(uint256)
func (_Instruments *InstrumentsCaller) EpochInstrument(opts *bind.CallOpts, instrument *big.Int, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "epochInstrument", instrument, epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochInstrument is a free data retrieval call binding the contract method 0x5a0eb264.
//
// Solidity: function epochInstrument(uint256 instrument, uint256 epoch) pure returns(uint256)
func (_Instruments *InstrumentsSession) EpochInstrument(instrument *big.Int, epoch *big.Int) (*big.Int, error) {
	return _Instruments.Contract.EpochInstrument(&_Instruments.CallOpts, instrument, epoch)
}

// EpochInstrument is a free data retrieval call binding the contract method 0x5a0eb264.
//
// Solidity: function epochInstrument(uint256 instrument, uint256 epoch) pure returns(uint256)
func (_Instruments *InstrumentsCallerSession) EpochInstrument(instrument *big.Int, epoch *big.Int) (*big.Int, error) {
	return _Instruments.Contract.EpochInstrument(&_Instruments.CallOpts, instrument, epoch)
}

// ExpiryPrice is a free data retrieval call binding the contract method 0xdee52b64.
//
// Solidity: function expiryPrice(address , uint256 ) view returns(uint256 price, uint256 timestamp)
func (_Instruments *InstrumentsCaller) ExpiryPrice(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "expiryPrice", arg0, arg1)

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ExpiryPrice is a free data retrieval call binding the contract method 0xdee52b64.
//
// Solidity: function expiryPrice(address , uint256 ) view returns(uint256 price, uint256 timestamp)
func (_Instruments *InstrumentsSession) ExpiryPrice(arg0 common.Address, arg1 *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _Instruments.Contract.ExpiryPrice(&_Instruments.CallOpts, arg0, arg1)
}

// ExpiryPrice is a free data retrieval call binding the contract method 0xdee52b64.
//
// Solidity: function expiryPrice(address , uint256 ) view returns(uint256 price, uint256 timestamp)
func (_Instruments *InstrumentsCallerSession) ExpiryPrice(arg0 common.Address, arg1 *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _Instruments.Contract.ExpiryPrice(&_Instruments.CallOpts, arg0, arg1)
}

// GetInstrumentAccounts is a free data retrieval call binding the contract method 0x04a34c2a.
//
// Solidity: function getInstrumentAccounts(uint256 instrument) view returns(address[])
func (_Instruments *InstrumentsCaller) GetInstrumentAccounts(opts *bind.CallOpts, instrument *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "getInstrumentAccounts", instrument)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInstrumentAccounts is a free data retrieval call binding the contract method 0x04a34c2a.
//
// Solidity: function getInstrumentAccounts(uint256 instrument) view returns(address[])
func (_Instruments *InstrumentsSession) GetInstrumentAccounts(instrument *big.Int) ([]common.Address, error) {
	return _Instruments.Contract.GetInstrumentAccounts(&_Instruments.CallOpts, instrument)
}

// GetInstrumentAccounts is a free data retrieval call binding the contract method 0x04a34c2a.
//
// Solidity: function getInstrumentAccounts(uint256 instrument) view returns(address[])
func (_Instruments *InstrumentsCallerSession) GetInstrumentAccounts(instrument *big.Int) ([]common.Address, error) {
	return _Instruments.Contract.GetInstrumentAccounts(&_Instruments.CallOpts, instrument)
}

// InstrumentAccounts is a free data retrieval call binding the contract method 0x19b57f9b.
//
// Solidity: function instrumentAccounts(uint256 , address ) view returns(bool)
func (_Instruments *InstrumentsCaller) InstrumentAccounts(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "instrumentAccounts", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InstrumentAccounts is a free data retrieval call binding the contract method 0x19b57f9b.
//
// Solidity: function instrumentAccounts(uint256 , address ) view returns(bool)
func (_Instruments *InstrumentsSession) InstrumentAccounts(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Instruments.Contract.InstrumentAccounts(&_Instruments.CallOpts, arg0, arg1)
}

// InstrumentAccounts is a free data retrieval call binding the contract method 0x19b57f9b.
//
// Solidity: function instrumentAccounts(uint256 , address ) view returns(bool)
func (_Instruments *InstrumentsCallerSession) InstrumentAccounts(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Instruments.Contract.InstrumentAccounts(&_Instruments.CallOpts, arg0, arg1)
}

// Instruments is a free data retrieval call binding the contract method 0x22bbad0b.
//
// Solidity: function instruments(uint256 ) view returns(uint8)
func (_Instruments *InstrumentsCaller) Instruments(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "instruments", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Instruments is a free data retrieval call binding the contract method 0x22bbad0b.
//
// Solidity: function instruments(uint256 ) view returns(uint8)
func (_Instruments *InstrumentsSession) Instruments(arg0 *big.Int) (uint8, error) {
	return _Instruments.Contract.Instruments(&_Instruments.CallOpts, arg0)
}

// Instruments is a free data retrieval call binding the contract method 0x22bbad0b.
//
// Solidity: function instruments(uint256 ) view returns(uint8)
func (_Instruments *InstrumentsCallerSession) Instruments(arg0 *big.Int) (uint8, error) {
	return _Instruments.Contract.Instruments(&_Instruments.CallOpts, arg0)
}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Instruments *InstrumentsCaller) Keepers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "keepers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Instruments *InstrumentsSession) Keepers(arg0 common.Address) (bool, error) {
	return _Instruments.Contract.Keepers(&_Instruments.CallOpts, arg0)
}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Instruments *InstrumentsCallerSession) Keepers(arg0 common.Address) (bool, error) {
	return _Instruments.Contract.Keepers(&_Instruments.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Instruments *InstrumentsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Instruments *InstrumentsSession) Name() (string, error) {
	return _Instruments.Contract.Name(&_Instruments.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Instruments *InstrumentsCallerSession) Name() (string, error) {
	return _Instruments.Contract.Name(&_Instruments.CallOpts)
}

// OpenInterest is a free data retrieval call binding the contract method 0x88e53ec8.
//
// Solidity: function openInterest(uint256 ) view returns(uint256)
func (_Instruments *InstrumentsCaller) OpenInterest(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "openInterest", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenInterest is a free data retrieval call binding the contract method 0x88e53ec8.
//
// Solidity: function openInterest(uint256 ) view returns(uint256)
func (_Instruments *InstrumentsSession) OpenInterest(arg0 *big.Int) (*big.Int, error) {
	return _Instruments.Contract.OpenInterest(&_Instruments.CallOpts, arg0)
}

// OpenInterest is a free data retrieval call binding the contract method 0x88e53ec8.
//
// Solidity: function openInterest(uint256 ) view returns(uint256)
func (_Instruments *InstrumentsCallerSession) OpenInterest(arg0 *big.Int) (*big.Int, error) {
	return _Instruments.Contract.OpenInterest(&_Instruments.CallOpts, arg0)
}

// Options is a free data retrieval call binding the contract method 0x409e2205.
//
// Solidity: function options(uint256 ) view returns(address asset, bool isPut, uint256 strike, uint256 expiry)
func (_Instruments *InstrumentsCaller) Options(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Asset  common.Address
	IsPut  bool
	Strike *big.Int
	Expiry *big.Int
}, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "options", arg0)

	outstruct := new(struct {
		Asset  common.Address
		IsPut  bool
		Strike *big.Int
		Expiry *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Asset = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsPut = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Strike = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Expiry = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Options is a free data retrieval call binding the contract method 0x409e2205.
//
// Solidity: function options(uint256 ) view returns(address asset, bool isPut, uint256 strike, uint256 expiry)
func (_Instruments *InstrumentsSession) Options(arg0 *big.Int) (struct {
	Asset  common.Address
	IsPut  bool
	Strike *big.Int
	Expiry *big.Int
}, error) {
	return _Instruments.Contract.Options(&_Instruments.CallOpts, arg0)
}

// Options is a free data retrieval call binding the contract method 0x409e2205.
//
// Solidity: function options(uint256 ) view returns(address asset, bool isPut, uint256 strike, uint256 expiry)
func (_Instruments *InstrumentsCallerSession) Options(arg0 *big.Int) (struct {
	Asset  common.Address
	IsPut  bool
	Strike *big.Int
	Expiry *big.Int
}, error) {
	return _Instruments.Contract.Options(&_Instruments.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Instruments *InstrumentsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Instruments *InstrumentsSession) Owner() (common.Address, error) {
	return _Instruments.Contract.Owner(&_Instruments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Instruments *InstrumentsCallerSession) Owner() (common.Address, error) {
	return _Instruments.Contract.Owner(&_Instruments.CallOpts)
}

// Perpetuals is a free data retrieval call binding the contract method 0x59a1ecaa.
//
// Solidity: function perpetuals(uint256 ) view returns(address)
func (_Instruments *InstrumentsCaller) Perpetuals(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "perpetuals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Perpetuals is a free data retrieval call binding the contract method 0x59a1ecaa.
//
// Solidity: function perpetuals(uint256 ) view returns(address)
func (_Instruments *InstrumentsSession) Perpetuals(arg0 *big.Int) (common.Address, error) {
	return _Instruments.Contract.Perpetuals(&_Instruments.CallOpts, arg0)
}

// Perpetuals is a free data retrieval call binding the contract method 0x59a1ecaa.
//
// Solidity: function perpetuals(uint256 ) view returns(address)
func (_Instruments *InstrumentsCallerSession) Perpetuals(arg0 *big.Int) (common.Address, error) {
	return _Instruments.Contract.Perpetuals(&_Instruments.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Instruments *InstrumentsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Instruments *InstrumentsSession) ProxiableUUID() ([32]byte, error) {
	return _Instruments.Contract.ProxiableUUID(&_Instruments.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Instruments *InstrumentsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Instruments.Contract.ProxiableUUID(&_Instruments.CallOpts)
}

// Settled is a free data retrieval call binding the contract method 0x28345780.
//
// Solidity: function settled(uint256 ) view returns(bool)
func (_Instruments *InstrumentsCaller) Settled(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "settled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Settled is a free data retrieval call binding the contract method 0x28345780.
//
// Solidity: function settled(uint256 ) view returns(bool)
func (_Instruments *InstrumentsSession) Settled(arg0 *big.Int) (bool, error) {
	return _Instruments.Contract.Settled(&_Instruments.CallOpts, arg0)
}

// Settled is a free data retrieval call binding the contract method 0x28345780.
//
// Solidity: function settled(uint256 ) view returns(bool)
func (_Instruments *InstrumentsCallerSession) Settled(arg0 *big.Int) (bool, error) {
	return _Instruments.Contract.Settled(&_Instruments.CallOpts, arg0)
}

// SettledInstrumentAccounts is a free data retrieval call binding the contract method 0xee206169.
//
// Solidity: function settledInstrumentAccounts(uint256 , address ) view returns(bool)
func (_Instruments *InstrumentsCaller) SettledInstrumentAccounts(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "settledInstrumentAccounts", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SettledInstrumentAccounts is a free data retrieval call binding the contract method 0xee206169.
//
// Solidity: function settledInstrumentAccounts(uint256 , address ) view returns(bool)
func (_Instruments *InstrumentsSession) SettledInstrumentAccounts(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Instruments.Contract.SettledInstrumentAccounts(&_Instruments.CallOpts, arg0, arg1)
}

// SettledInstrumentAccounts is a free data retrieval call binding the contract method 0xee206169.
//
// Solidity: function settledInstrumentAccounts(uint256 , address ) view returns(bool)
func (_Instruments *InstrumentsCallerSession) SettledInstrumentAccounts(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Instruments.Contract.SettledInstrumentAccounts(&_Instruments.CallOpts, arg0, arg1)
}

// SettledLength is a free data retrieval call binding the contract method 0x205e4373.
//
// Solidity: function settledLength(uint256 ) view returns(uint256)
func (_Instruments *InstrumentsCaller) SettledLength(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "settledLength", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SettledLength is a free data retrieval call binding the contract method 0x205e4373.
//
// Solidity: function settledLength(uint256 ) view returns(uint256)
func (_Instruments *InstrumentsSession) SettledLength(arg0 *big.Int) (*big.Int, error) {
	return _Instruments.Contract.SettledLength(&_Instruments.CallOpts, arg0)
}

// SettledLength is a free data retrieval call binding the contract method 0x205e4373.
//
// Solidity: function settledLength(uint256 ) view returns(uint256)
func (_Instruments *InstrumentsCallerSession) SettledLength(arg0 *big.Int) (*big.Int, error) {
	return _Instruments.Contract.SettledLength(&_Instruments.CallOpts, arg0)
}

// Settling is a free data retrieval call binding the contract method 0xc7f9df8b.
//
// Solidity: function settling(uint256 ) view returns(bool)
func (_Instruments *InstrumentsCaller) Settling(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "settling", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Settling is a free data retrieval call binding the contract method 0xc7f9df8b.
//
// Solidity: function settling(uint256 ) view returns(bool)
func (_Instruments *InstrumentsSession) Settling(arg0 *big.Int) (bool, error) {
	return _Instruments.Contract.Settling(&_Instruments.CallOpts, arg0)
}

// Settling is a free data retrieval call binding the contract method 0xc7f9df8b.
//
// Solidity: function settling(uint256 ) view returns(bool)
func (_Instruments *InstrumentsCallerSession) Settling(arg0 *big.Int) (bool, error) {
	return _Instruments.Contract.Settling(&_Instruments.CallOpts, arg0)
}

// TotalOptionPayout is a free data retrieval call binding the contract method 0x5af9ac0a.
//
// Solidity: function totalOptionPayout(uint256 instrument) view returns(uint256)
func (_Instruments *InstrumentsCaller) TotalOptionPayout(opts *bind.CallOpts, instrument *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Instruments.contract.Call(opts, &out, "totalOptionPayout", instrument)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOptionPayout is a free data retrieval call binding the contract method 0x5af9ac0a.
//
// Solidity: function totalOptionPayout(uint256 instrument) view returns(uint256)
func (_Instruments *InstrumentsSession) TotalOptionPayout(instrument *big.Int) (*big.Int, error) {
	return _Instruments.Contract.TotalOptionPayout(&_Instruments.CallOpts, instrument)
}

// TotalOptionPayout is a free data retrieval call binding the contract method 0x5af9ac0a.
//
// Solidity: function totalOptionPayout(uint256 instrument) view returns(uint256)
func (_Instruments *InstrumentsCallerSession) TotalOptionPayout(instrument *big.Int) (*big.Int, error) {
	return _Instruments.Contract.TotalOptionPayout(&_Instruments.CallOpts, instrument)
}

// AddOption is a paid mutator transaction binding the contract method 0xe6a97e44.
//
// Solidity: function addOption(uint256 instrument, (address,bool,uint256,uint256) option) returns()
func (_Instruments *InstrumentsTransactor) AddOption(opts *bind.TransactOpts, instrument *big.Int, option OptionsOption) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "addOption", instrument, option)
}

// AddOption is a paid mutator transaction binding the contract method 0xe6a97e44.
//
// Solidity: function addOption(uint256 instrument, (address,bool,uint256,uint256) option) returns()
func (_Instruments *InstrumentsSession) AddOption(instrument *big.Int, option OptionsOption) (*types.Transaction, error) {
	return _Instruments.Contract.AddOption(&_Instruments.TransactOpts, instrument, option)
}

// AddOption is a paid mutator transaction binding the contract method 0xe6a97e44.
//
// Solidity: function addOption(uint256 instrument, (address,bool,uint256,uint256) option) returns()
func (_Instruments *InstrumentsTransactorSession) AddOption(instrument *big.Int, option OptionsOption) (*types.Transaction, error) {
	return _Instruments.Contract.AddOption(&_Instruments.TransactOpts, instrument, option)
}

// AddPerpetual is a paid mutator transaction binding the contract method 0xebfb62b3.
//
// Solidity: function addPerpetual(uint256 instrument, address asset) returns()
func (_Instruments *InstrumentsTransactor) AddPerpetual(opts *bind.TransactOpts, instrument *big.Int, asset common.Address) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "addPerpetual", instrument, asset)
}

// AddPerpetual is a paid mutator transaction binding the contract method 0xebfb62b3.
//
// Solidity: function addPerpetual(uint256 instrument, address asset) returns()
func (_Instruments *InstrumentsSession) AddPerpetual(instrument *big.Int, asset common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.AddPerpetual(&_Instruments.TransactOpts, instrument, asset)
}

// AddPerpetual is a paid mutator transaction binding the contract method 0xebfb62b3.
//
// Solidity: function addPerpetual(uint256 instrument, address asset) returns()
func (_Instruments *InstrumentsTransactorSession) AddPerpetual(instrument *big.Int, asset common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.AddPerpetual(&_Instruments.TransactOpts, instrument, asset)
}

// ChargeFunding is a paid mutator transaction binding the contract method 0xa557b067.
//
// Solidity: function chargeFunding(uint256 instrument, uint256 epoch, uint256 price, int256 fundingRate) returns(bool)
func (_Instruments *InstrumentsTransactor) ChargeFunding(opts *bind.TransactOpts, instrument *big.Int, epoch *big.Int, price *big.Int, fundingRate *big.Int) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "chargeFunding", instrument, epoch, price, fundingRate)
}

// ChargeFunding is a paid mutator transaction binding the contract method 0xa557b067.
//
// Solidity: function chargeFunding(uint256 instrument, uint256 epoch, uint256 price, int256 fundingRate) returns(bool)
func (_Instruments *InstrumentsSession) ChargeFunding(instrument *big.Int, epoch *big.Int, price *big.Int, fundingRate *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.ChargeFunding(&_Instruments.TransactOpts, instrument, epoch, price, fundingRate)
}

// ChargeFunding is a paid mutator transaction binding the contract method 0xa557b067.
//
// Solidity: function chargeFunding(uint256 instrument, uint256 epoch, uint256 price, int256 fundingRate) returns(bool)
func (_Instruments *InstrumentsTransactorSession) ChargeFunding(instrument *big.Int, epoch *big.Int, price *big.Int, fundingRate *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.ChargeFunding(&_Instruments.TransactOpts, instrument, epoch, price, fundingRate)
}

// ChargeFunding0 is a paid mutator transaction binding the contract method 0xd1d49e7c.
//
// Solidity: function chargeFunding(uint256 instrument, uint256 epoch, uint256 price, uint256 start, uint256 end, int256 fundingRate) returns(bool)
func (_Instruments *InstrumentsTransactor) ChargeFunding0(opts *bind.TransactOpts, instrument *big.Int, epoch *big.Int, price *big.Int, start *big.Int, end *big.Int, fundingRate *big.Int) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "chargeFunding0", instrument, epoch, price, start, end, fundingRate)
}

// ChargeFunding0 is a paid mutator transaction binding the contract method 0xd1d49e7c.
//
// Solidity: function chargeFunding(uint256 instrument, uint256 epoch, uint256 price, uint256 start, uint256 end, int256 fundingRate) returns(bool)
func (_Instruments *InstrumentsSession) ChargeFunding0(instrument *big.Int, epoch *big.Int, price *big.Int, start *big.Int, end *big.Int, fundingRate *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.ChargeFunding0(&_Instruments.TransactOpts, instrument, epoch, price, start, end, fundingRate)
}

// ChargeFunding0 is a paid mutator transaction binding the contract method 0xd1d49e7c.
//
// Solidity: function chargeFunding(uint256 instrument, uint256 epoch, uint256 price, uint256 start, uint256 end, int256 fundingRate) returns(bool)
func (_Instruments *InstrumentsTransactorSession) ChargeFunding0(instrument *big.Int, epoch *big.Int, price *big.Int, start *big.Int, end *big.Int, fundingRate *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.ChargeFunding0(&_Instruments.TransactOpts, instrument, epoch, price, start, end, fundingRate)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Instruments *InstrumentsTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Instruments *InstrumentsSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.Initialize(&_Instruments.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Instruments *InstrumentsTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.Initialize(&_Instruments.TransactOpts, _owner)
}

// LoadAccount is a paid mutator transaction binding the contract method 0x9d56921f.
//
// Solidity: function loadAccount(address account, (address,int256)[] balances, (uint256,uint256,int256)[] positions, (address,uint256)[] signingKeys) returns()
func (_Instruments *InstrumentsTransactor) LoadAccount(opts *bind.TransactOpts, account common.Address, balances []AccountsInterfaceBalance, positions []AccountsInterfacePosition, signingKeys []AccountsInterfaceSigningKey) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "loadAccount", account, balances, positions, signingKeys)
}

// LoadAccount is a paid mutator transaction binding the contract method 0x9d56921f.
//
// Solidity: function loadAccount(address account, (address,int256)[] balances, (uint256,uint256,int256)[] positions, (address,uint256)[] signingKeys) returns()
func (_Instruments *InstrumentsSession) LoadAccount(account common.Address, balances []AccountsInterfaceBalance, positions []AccountsInterfacePosition, signingKeys []AccountsInterfaceSigningKey) (*types.Transaction, error) {
	return _Instruments.Contract.LoadAccount(&_Instruments.TransactOpts, account, balances, positions, signingKeys)
}

// LoadAccount is a paid mutator transaction binding the contract method 0x9d56921f.
//
// Solidity: function loadAccount(address account, (address,int256)[] balances, (uint256,uint256,int256)[] positions, (address,uint256)[] signingKeys) returns()
func (_Instruments *InstrumentsTransactorSession) LoadAccount(account common.Address, balances []AccountsInterfaceBalance, positions []AccountsInterfacePosition, signingKeys []AccountsInterfaceSigningKey) (*types.Transaction, error) {
	return _Instruments.Contract.LoadAccount(&_Instruments.TransactOpts, account, balances, positions, signingKeys)
}

// LoadOpenInterest is a paid mutator transaction binding the contract method 0x8f5cafcc.
//
// Solidity: function loadOpenInterest(uint256 instrument, uint256 oi) returns()
func (_Instruments *InstrumentsTransactor) LoadOpenInterest(opts *bind.TransactOpts, instrument *big.Int, oi *big.Int) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "loadOpenInterest", instrument, oi)
}

// LoadOpenInterest is a paid mutator transaction binding the contract method 0x8f5cafcc.
//
// Solidity: function loadOpenInterest(uint256 instrument, uint256 oi) returns()
func (_Instruments *InstrumentsSession) LoadOpenInterest(instrument *big.Int, oi *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.LoadOpenInterest(&_Instruments.TransactOpts, instrument, oi)
}

// LoadOpenInterest is a paid mutator transaction binding the contract method 0x8f5cafcc.
//
// Solidity: function loadOpenInterest(uint256 instrument, uint256 oi) returns()
func (_Instruments *InstrumentsTransactorSession) LoadOpenInterest(instrument *big.Int, oi *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.LoadOpenInterest(&_Instruments.TransactOpts, instrument, oi)
}

// LoadOption is a paid mutator transaction binding the contract method 0x5f92e55f.
//
// Solidity: function loadOption(uint256 instrument, (address,bool,uint256,uint256) option) returns()
func (_Instruments *InstrumentsTransactor) LoadOption(opts *bind.TransactOpts, instrument *big.Int, option OptionsOption) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "loadOption", instrument, option)
}

// LoadOption is a paid mutator transaction binding the contract method 0x5f92e55f.
//
// Solidity: function loadOption(uint256 instrument, (address,bool,uint256,uint256) option) returns()
func (_Instruments *InstrumentsSession) LoadOption(instrument *big.Int, option OptionsOption) (*types.Transaction, error) {
	return _Instruments.Contract.LoadOption(&_Instruments.TransactOpts, instrument, option)
}

// LoadOption is a paid mutator transaction binding the contract method 0x5f92e55f.
//
// Solidity: function loadOption(uint256 instrument, (address,bool,uint256,uint256) option) returns()
func (_Instruments *InstrumentsTransactorSession) LoadOption(instrument *big.Int, option OptionsOption) (*types.Transaction, error) {
	return _Instruments.Contract.LoadOption(&_Instruments.TransactOpts, instrument, option)
}

// LoadPerpetual is a paid mutator transaction binding the contract method 0x53f9ead2.
//
// Solidity: function loadPerpetual(uint256 instrument, address asset) returns()
func (_Instruments *InstrumentsTransactor) LoadPerpetual(opts *bind.TransactOpts, instrument *big.Int, asset common.Address) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "loadPerpetual", instrument, asset)
}

// LoadPerpetual is a paid mutator transaction binding the contract method 0x53f9ead2.
//
// Solidity: function loadPerpetual(uint256 instrument, address asset) returns()
func (_Instruments *InstrumentsSession) LoadPerpetual(instrument *big.Int, asset common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.LoadPerpetual(&_Instruments.TransactOpts, instrument, asset)
}

// LoadPerpetual is a paid mutator transaction binding the contract method 0x53f9ead2.
//
// Solidity: function loadPerpetual(uint256 instrument, address asset) returns()
func (_Instruments *InstrumentsTransactorSession) LoadPerpetual(instrument *big.Int, asset common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.LoadPerpetual(&_Instruments.TransactOpts, instrument, asset)
}

// OverrideExpiryPrice is a paid mutator transaction binding the contract method 0xbda9e302.
//
// Solidity: function overrideExpiryPrice(address asset, uint256 expiry, uint256 price) returns()
func (_Instruments *InstrumentsTransactor) OverrideExpiryPrice(opts *bind.TransactOpts, asset common.Address, expiry *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "overrideExpiryPrice", asset, expiry, price)
}

// OverrideExpiryPrice is a paid mutator transaction binding the contract method 0xbda9e302.
//
// Solidity: function overrideExpiryPrice(address asset, uint256 expiry, uint256 price) returns()
func (_Instruments *InstrumentsSession) OverrideExpiryPrice(asset common.Address, expiry *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.OverrideExpiryPrice(&_Instruments.TransactOpts, asset, expiry, price)
}

// OverrideExpiryPrice is a paid mutator transaction binding the contract method 0xbda9e302.
//
// Solidity: function overrideExpiryPrice(address asset, uint256 expiry, uint256 price) returns()
func (_Instruments *InstrumentsTransactorSession) OverrideExpiryPrice(asset common.Address, expiry *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.OverrideExpiryPrice(&_Instruments.TransactOpts, asset, expiry, price)
}

// SetAccounts is a paid mutator transaction binding the contract method 0x08274382.
//
// Solidity: function setAccounts(address _accounts) returns()
func (_Instruments *InstrumentsTransactor) SetAccounts(opts *bind.TransactOpts, _accounts common.Address) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "setAccounts", _accounts)
}

// SetAccounts is a paid mutator transaction binding the contract method 0x08274382.
//
// Solidity: function setAccounts(address _accounts) returns()
func (_Instruments *InstrumentsSession) SetAccounts(_accounts common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.SetAccounts(&_Instruments.TransactOpts, _accounts)
}

// SetAccounts is a paid mutator transaction binding the contract method 0x08274382.
//
// Solidity: function setAccounts(address _accounts) returns()
func (_Instruments *InstrumentsTransactorSession) SetAccounts(_accounts common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.SetAccounts(&_Instruments.TransactOpts, _accounts)
}

// SetDisputePeriod is a paid mutator transaction binding the contract method 0x8b3cddaf.
//
// Solidity: function setDisputePeriod(address asset, uint256 period) returns()
func (_Instruments *InstrumentsTransactor) SetDisputePeriod(opts *bind.TransactOpts, asset common.Address, period *big.Int) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "setDisputePeriod", asset, period)
}

// SetDisputePeriod is a paid mutator transaction binding the contract method 0x8b3cddaf.
//
// Solidity: function setDisputePeriod(address asset, uint256 period) returns()
func (_Instruments *InstrumentsSession) SetDisputePeriod(asset common.Address, period *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.SetDisputePeriod(&_Instruments.TransactOpts, asset, period)
}

// SetDisputePeriod is a paid mutator transaction binding the contract method 0x8b3cddaf.
//
// Solidity: function setDisputePeriod(address asset, uint256 period) returns()
func (_Instruments *InstrumentsTransactorSession) SetDisputePeriod(asset common.Address, period *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.SetDisputePeriod(&_Instruments.TransactOpts, asset, period)
}

// SetExpiryPrice is a paid mutator transaction binding the contract method 0xee531409.
//
// Solidity: function setExpiryPrice(address asset, uint256 expiry, uint256 price) returns()
func (_Instruments *InstrumentsTransactor) SetExpiryPrice(opts *bind.TransactOpts, asset common.Address, expiry *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "setExpiryPrice", asset, expiry, price)
}

// SetExpiryPrice is a paid mutator transaction binding the contract method 0xee531409.
//
// Solidity: function setExpiryPrice(address asset, uint256 expiry, uint256 price) returns()
func (_Instruments *InstrumentsSession) SetExpiryPrice(asset common.Address, expiry *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.SetExpiryPrice(&_Instruments.TransactOpts, asset, expiry, price)
}

// SetExpiryPrice is a paid mutator transaction binding the contract method 0xee531409.
//
// Solidity: function setExpiryPrice(address asset, uint256 expiry, uint256 price) returns()
func (_Instruments *InstrumentsTransactorSession) SetExpiryPrice(asset common.Address, expiry *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.SetExpiryPrice(&_Instruments.TransactOpts, asset, expiry, price)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Instruments *InstrumentsTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Instruments *InstrumentsSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.SetOwner(&_Instruments.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Instruments *InstrumentsTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.SetOwner(&_Instruments.TransactOpts, newOwner)
}

// SettleOption is a paid mutator transaction binding the contract method 0x8cffa230.
//
// Solidity: function settleOption(uint256 instrument, uint256 price, uint256 start, uint256 end, uint256 settlementFee) returns(bool)
func (_Instruments *InstrumentsTransactor) SettleOption(opts *bind.TransactOpts, instrument *big.Int, price *big.Int, start *big.Int, end *big.Int, settlementFee *big.Int) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "settleOption", instrument, price, start, end, settlementFee)
}

// SettleOption is a paid mutator transaction binding the contract method 0x8cffa230.
//
// Solidity: function settleOption(uint256 instrument, uint256 price, uint256 start, uint256 end, uint256 settlementFee) returns(bool)
func (_Instruments *InstrumentsSession) SettleOption(instrument *big.Int, price *big.Int, start *big.Int, end *big.Int, settlementFee *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.SettleOption(&_Instruments.TransactOpts, instrument, price, start, end, settlementFee)
}

// SettleOption is a paid mutator transaction binding the contract method 0x8cffa230.
//
// Solidity: function settleOption(uint256 instrument, uint256 price, uint256 start, uint256 end, uint256 settlementFee) returns(bool)
func (_Instruments *InstrumentsTransactorSession) SettleOption(instrument *big.Int, price *big.Int, start *big.Int, end *big.Int, settlementFee *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.SettleOption(&_Instruments.TransactOpts, instrument, price, start, end, settlementFee)
}

// SettleOption0 is a paid mutator transaction binding the contract method 0x9d5bb5c0.
//
// Solidity: function settleOption(uint256 instrument, uint256 settlementFee) returns(bool)
func (_Instruments *InstrumentsTransactor) SettleOption0(opts *bind.TransactOpts, instrument *big.Int, settlementFee *big.Int) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "settleOption0", instrument, settlementFee)
}

// SettleOption0 is a paid mutator transaction binding the contract method 0x9d5bb5c0.
//
// Solidity: function settleOption(uint256 instrument, uint256 settlementFee) returns(bool)
func (_Instruments *InstrumentsSession) SettleOption0(instrument *big.Int, settlementFee *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.SettleOption0(&_Instruments.TransactOpts, instrument, settlementFee)
}

// SettleOption0 is a paid mutator transaction binding the contract method 0x9d5bb5c0.
//
// Solidity: function settleOption(uint256 instrument, uint256 settlementFee) returns(bool)
func (_Instruments *InstrumentsTransactorSession) SettleOption0(instrument *big.Int, settlementFee *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.SettleOption0(&_Instruments.TransactOpts, instrument, settlementFee)
}

// TransferInstrument is a paid mutator transaction binding the contract method 0x991f94a0.
//
// Solidity: function transferInstrument(uint256 instrument, address from, address to, uint256 amount, uint256 price) returns(int256, int256, uint8)
func (_Instruments *InstrumentsTransactor) TransferInstrument(opts *bind.TransactOpts, instrument *big.Int, from common.Address, to common.Address, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "transferInstrument", instrument, from, to, amount, price)
}

// TransferInstrument is a paid mutator transaction binding the contract method 0x991f94a0.
//
// Solidity: function transferInstrument(uint256 instrument, address from, address to, uint256 amount, uint256 price) returns(int256, int256, uint8)
func (_Instruments *InstrumentsSession) TransferInstrument(instrument *big.Int, from common.Address, to common.Address, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.TransferInstrument(&_Instruments.TransactOpts, instrument, from, to, amount, price)
}

// TransferInstrument is a paid mutator transaction binding the contract method 0x991f94a0.
//
// Solidity: function transferInstrument(uint256 instrument, address from, address to, uint256 amount, uint256 price) returns(int256, int256, uint8)
func (_Instruments *InstrumentsTransactorSession) TransferInstrument(instrument *big.Int, from common.Address, to common.Address, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Instruments.Contract.TransferInstrument(&_Instruments.TransactOpts, instrument, from, to, amount, price)
}

// UpdateAuthority is a paid mutator transaction binding the contract method 0x6cd22eaf.
//
// Solidity: function updateAuthority(address authority, bool allowed) returns()
func (_Instruments *InstrumentsTransactor) UpdateAuthority(opts *bind.TransactOpts, authority common.Address, allowed bool) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "updateAuthority", authority, allowed)
}

// UpdateAuthority is a paid mutator transaction binding the contract method 0x6cd22eaf.
//
// Solidity: function updateAuthority(address authority, bool allowed) returns()
func (_Instruments *InstrumentsSession) UpdateAuthority(authority common.Address, allowed bool) (*types.Transaction, error) {
	return _Instruments.Contract.UpdateAuthority(&_Instruments.TransactOpts, authority, allowed)
}

// UpdateAuthority is a paid mutator transaction binding the contract method 0x6cd22eaf.
//
// Solidity: function updateAuthority(address authority, bool allowed) returns()
func (_Instruments *InstrumentsTransactorSession) UpdateAuthority(authority common.Address, allowed bool) (*types.Transaction, error) {
	return _Instruments.Contract.UpdateAuthority(&_Instruments.TransactOpts, authority, allowed)
}

// UpdateKeeper is a paid mutator transaction binding the contract method 0xd3057877.
//
// Solidity: function updateKeeper(address keeper, bool allowed) returns()
func (_Instruments *InstrumentsTransactor) UpdateKeeper(opts *bind.TransactOpts, keeper common.Address, allowed bool) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "updateKeeper", keeper, allowed)
}

// UpdateKeeper is a paid mutator transaction binding the contract method 0xd3057877.
//
// Solidity: function updateKeeper(address keeper, bool allowed) returns()
func (_Instruments *InstrumentsSession) UpdateKeeper(keeper common.Address, allowed bool) (*types.Transaction, error) {
	return _Instruments.Contract.UpdateKeeper(&_Instruments.TransactOpts, keeper, allowed)
}

// UpdateKeeper is a paid mutator transaction binding the contract method 0xd3057877.
//
// Solidity: function updateKeeper(address keeper, bool allowed) returns()
func (_Instruments *InstrumentsTransactorSession) UpdateKeeper(keeper common.Address, allowed bool) (*types.Transaction, error) {
	return _Instruments.Contract.UpdateKeeper(&_Instruments.TransactOpts, keeper, allowed)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Instruments *InstrumentsTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Instruments *InstrumentsSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.UpgradeTo(&_Instruments.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Instruments *InstrumentsTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Instruments.Contract.UpgradeTo(&_Instruments.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Instruments *InstrumentsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Instruments.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Instruments *InstrumentsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Instruments.Contract.UpgradeToAndCall(&_Instruments.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Instruments *InstrumentsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Instruments.Contract.UpgradeToAndCall(&_Instruments.TransactOpts, newImplementation, data)
}

// InstrumentsAccountsUpdatedIterator is returned from FilterAccountsUpdated and is used to iterate over the raw logs and unpacked data for AccountsUpdated events raised by the Instruments contract.
type InstrumentsAccountsUpdatedIterator struct {
	Event *InstrumentsAccountsUpdated // Event containing the contract specifics and raw log

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
func (it *InstrumentsAccountsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsAccountsUpdated)
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
		it.Event = new(InstrumentsAccountsUpdated)
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
func (it *InstrumentsAccountsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsAccountsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsAccountsUpdated represents a AccountsUpdated event raised by the Instruments contract.
type InstrumentsAccountsUpdated struct {
	Accounts common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccountsUpdated is a free log retrieval operation binding the contract event 0xbecde7fe690c73ba54232f00eb06c31464f65b45aff18984febaa80df22dcb8d.
//
// Solidity: event AccountsUpdated(address indexed accounts)
func (_Instruments *InstrumentsFilterer) FilterAccountsUpdated(opts *bind.FilterOpts, accounts []common.Address) (*InstrumentsAccountsUpdatedIterator, error) {

	var accountsRule []interface{}
	for _, accountsItem := range accounts {
		accountsRule = append(accountsRule, accountsItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "AccountsUpdated", accountsRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsAccountsUpdatedIterator{contract: _Instruments.contract, event: "AccountsUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountsUpdated is a free log subscription operation binding the contract event 0xbecde7fe690c73ba54232f00eb06c31464f65b45aff18984febaa80df22dcb8d.
//
// Solidity: event AccountsUpdated(address indexed accounts)
func (_Instruments *InstrumentsFilterer) WatchAccountsUpdated(opts *bind.WatchOpts, sink chan<- *InstrumentsAccountsUpdated, accounts []common.Address) (event.Subscription, error) {

	var accountsRule []interface{}
	for _, accountsItem := range accounts {
		accountsRule = append(accountsRule, accountsItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "AccountsUpdated", accountsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsAccountsUpdated)
				if err := _Instruments.contract.UnpackLog(event, "AccountsUpdated", log); err != nil {
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

// ParseAccountsUpdated is a log parse operation binding the contract event 0xbecde7fe690c73ba54232f00eb06c31464f65b45aff18984febaa80df22dcb8d.
//
// Solidity: event AccountsUpdated(address indexed accounts)
func (_Instruments *InstrumentsFilterer) ParseAccountsUpdated(log types.Log) (*InstrumentsAccountsUpdated, error) {
	event := new(InstrumentsAccountsUpdated)
	if err := _Instruments.contract.UnpackLog(event, "AccountsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Instruments contract.
type InstrumentsAdminChangedIterator struct {
	Event *InstrumentsAdminChanged // Event containing the contract specifics and raw log

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
func (it *InstrumentsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsAdminChanged)
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
		it.Event = new(InstrumentsAdminChanged)
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
func (it *InstrumentsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsAdminChanged represents a AdminChanged event raised by the Instruments contract.
type InstrumentsAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Instruments *InstrumentsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*InstrumentsAdminChangedIterator, error) {

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &InstrumentsAdminChangedIterator{contract: _Instruments.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Instruments *InstrumentsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *InstrumentsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsAdminChanged)
				if err := _Instruments.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Instruments *InstrumentsFilterer) ParseAdminChanged(log types.Log) (*InstrumentsAdminChanged, error) {
	event := new(InstrumentsAdminChanged)
	if err := _Instruments.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the Instruments contract.
type InstrumentsAuthorityUpdatedIterator struct {
	Event *InstrumentsAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *InstrumentsAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsAuthorityUpdated)
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
		it.Event = new(InstrumentsAuthorityUpdated)
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
func (it *InstrumentsAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsAuthorityUpdated represents a AuthorityUpdated event raised by the Instruments contract.
type InstrumentsAuthorityUpdated struct {
	Authority common.Address
	Allowed   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0xc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc.
//
// Solidity: event AuthorityUpdated(address indexed authority, bool indexed allowed)
func (_Instruments *InstrumentsFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts, authority []common.Address, allowed []bool) (*InstrumentsAuthorityUpdatedIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "AuthorityUpdated", authorityRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsAuthorityUpdatedIterator{contract: _Instruments.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0xc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc.
//
// Solidity: event AuthorityUpdated(address indexed authority, bool indexed allowed)
func (_Instruments *InstrumentsFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *InstrumentsAuthorityUpdated, authority []common.Address, allowed []bool) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "AuthorityUpdated", authorityRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsAuthorityUpdated)
				if err := _Instruments.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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
func (_Instruments *InstrumentsFilterer) ParseAuthorityUpdated(log types.Log) (*InstrumentsAuthorityUpdated, error) {
	event := new(InstrumentsAuthorityUpdated)
	if err := _Instruments.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Instruments contract.
type InstrumentsBeaconUpgradedIterator struct {
	Event *InstrumentsBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *InstrumentsBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsBeaconUpgraded)
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
		it.Event = new(InstrumentsBeaconUpgraded)
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
func (it *InstrumentsBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsBeaconUpgraded represents a BeaconUpgraded event raised by the Instruments contract.
type InstrumentsBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Instruments *InstrumentsFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*InstrumentsBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsBeaconUpgradedIterator{contract: _Instruments.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Instruments *InstrumentsFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *InstrumentsBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsBeaconUpgraded)
				if err := _Instruments.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Instruments *InstrumentsFilterer) ParseBeaconUpgraded(log types.Log) (*InstrumentsBeaconUpgraded, error) {
	event := new(InstrumentsBeaconUpgraded)
	if err := _Instruments.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsDisputePeriodUpdatedIterator is returned from FilterDisputePeriodUpdated and is used to iterate over the raw logs and unpacked data for DisputePeriodUpdated events raised by the Instruments contract.
type InstrumentsDisputePeriodUpdatedIterator struct {
	Event *InstrumentsDisputePeriodUpdated // Event containing the contract specifics and raw log

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
func (it *InstrumentsDisputePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsDisputePeriodUpdated)
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
		it.Event = new(InstrumentsDisputePeriodUpdated)
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
func (it *InstrumentsDisputePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsDisputePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsDisputePeriodUpdated represents a DisputePeriodUpdated event raised by the Instruments contract.
type InstrumentsDisputePeriodUpdated struct {
	Asset  common.Address
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisputePeriodUpdated is a free log retrieval operation binding the contract event 0x6650c9b75cf9b0e48611d4e590a8bd06f11be77c6b005b76114d42955d95e3a7.
//
// Solidity: event DisputePeriodUpdated(address indexed asset, uint256 indexed period)
func (_Instruments *InstrumentsFilterer) FilterDisputePeriodUpdated(opts *bind.FilterOpts, asset []common.Address, period []*big.Int) (*InstrumentsDisputePeriodUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "DisputePeriodUpdated", assetRule, periodRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsDisputePeriodUpdatedIterator{contract: _Instruments.contract, event: "DisputePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputePeriodUpdated is a free log subscription operation binding the contract event 0x6650c9b75cf9b0e48611d4e590a8bd06f11be77c6b005b76114d42955d95e3a7.
//
// Solidity: event DisputePeriodUpdated(address indexed asset, uint256 indexed period)
func (_Instruments *InstrumentsFilterer) WatchDisputePeriodUpdated(opts *bind.WatchOpts, sink chan<- *InstrumentsDisputePeriodUpdated, asset []common.Address, period []*big.Int) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "DisputePeriodUpdated", assetRule, periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsDisputePeriodUpdated)
				if err := _Instruments.contract.UnpackLog(event, "DisputePeriodUpdated", log); err != nil {
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

// ParseDisputePeriodUpdated is a log parse operation binding the contract event 0x6650c9b75cf9b0e48611d4e590a8bd06f11be77c6b005b76114d42955d95e3a7.
//
// Solidity: event DisputePeriodUpdated(address indexed asset, uint256 indexed period)
func (_Instruments *InstrumentsFilterer) ParseDisputePeriodUpdated(log types.Log) (*InstrumentsDisputePeriodUpdated, error) {
	event := new(InstrumentsDisputePeriodUpdated)
	if err := _Instruments.contract.UnpackLog(event, "DisputePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsExpiryPriceSetIterator is returned from FilterExpiryPriceSet and is used to iterate over the raw logs and unpacked data for ExpiryPriceSet events raised by the Instruments contract.
type InstrumentsExpiryPriceSetIterator struct {
	Event *InstrumentsExpiryPriceSet // Event containing the contract specifics and raw log

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
func (it *InstrumentsExpiryPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsExpiryPriceSet)
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
		it.Event = new(InstrumentsExpiryPriceSet)
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
func (it *InstrumentsExpiryPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsExpiryPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsExpiryPriceSet represents a ExpiryPriceSet event raised by the Instruments contract.
type InstrumentsExpiryPriceSet struct {
	Asset  common.Address
	Expiry *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExpiryPriceSet is a free log retrieval operation binding the contract event 0x8cc2700c35138467f8087f0f5101b52a163c8554838ef019be2546e67eebb9c7.
//
// Solidity: event ExpiryPriceSet(address indexed asset, uint256 indexed expiry, uint256 price)
func (_Instruments *InstrumentsFilterer) FilterExpiryPriceSet(opts *bind.FilterOpts, asset []common.Address, expiry []*big.Int) (*InstrumentsExpiryPriceSetIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var expiryRule []interface{}
	for _, expiryItem := range expiry {
		expiryRule = append(expiryRule, expiryItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "ExpiryPriceSet", assetRule, expiryRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsExpiryPriceSetIterator{contract: _Instruments.contract, event: "ExpiryPriceSet", logs: logs, sub: sub}, nil
}

// WatchExpiryPriceSet is a free log subscription operation binding the contract event 0x8cc2700c35138467f8087f0f5101b52a163c8554838ef019be2546e67eebb9c7.
//
// Solidity: event ExpiryPriceSet(address indexed asset, uint256 indexed expiry, uint256 price)
func (_Instruments *InstrumentsFilterer) WatchExpiryPriceSet(opts *bind.WatchOpts, sink chan<- *InstrumentsExpiryPriceSet, asset []common.Address, expiry []*big.Int) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var expiryRule []interface{}
	for _, expiryItem := range expiry {
		expiryRule = append(expiryRule, expiryItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "ExpiryPriceSet", assetRule, expiryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsExpiryPriceSet)
				if err := _Instruments.contract.UnpackLog(event, "ExpiryPriceSet", log); err != nil {
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

// ParseExpiryPriceSet is a log parse operation binding the contract event 0x8cc2700c35138467f8087f0f5101b52a163c8554838ef019be2546e67eebb9c7.
//
// Solidity: event ExpiryPriceSet(address indexed asset, uint256 indexed expiry, uint256 price)
func (_Instruments *InstrumentsFilterer) ParseExpiryPriceSet(log types.Log) (*InstrumentsExpiryPriceSet, error) {
	event := new(InstrumentsExpiryPriceSet)
	if err := _Instruments.contract.UnpackLog(event, "ExpiryPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the Instruments contract.
type InstrumentsFeeIterator struct {
	Event *InstrumentsFee // Event containing the contract specifics and raw log

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
func (it *InstrumentsFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsFee)
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
		it.Event = new(InstrumentsFee)
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
func (it *InstrumentsFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsFee represents a Fee event raised by the Instruments contract.
type InstrumentsFee struct {
	Instrument    *big.Int
	Account       common.Address
	SettlementFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xc90f6e366e1943d6ee7f9f5d64d27a53bf4a64e438dc757f9b7e1e96b9aaaab2.
//
// Solidity: event Fee(uint256 indexed instrument, address indexed account, uint256 settlementFee)
func (_Instruments *InstrumentsFilterer) FilterFee(opts *bind.FilterOpts, instrument []*big.Int, account []common.Address) (*InstrumentsFeeIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "Fee", instrumentRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsFeeIterator{contract: _Instruments.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xc90f6e366e1943d6ee7f9f5d64d27a53bf4a64e438dc757f9b7e1e96b9aaaab2.
//
// Solidity: event Fee(uint256 indexed instrument, address indexed account, uint256 settlementFee)
func (_Instruments *InstrumentsFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *InstrumentsFee, instrument []*big.Int, account []common.Address) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "Fee", instrumentRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsFee)
				if err := _Instruments.contract.UnpackLog(event, "Fee", log); err != nil {
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

// ParseFee is a log parse operation binding the contract event 0xc90f6e366e1943d6ee7f9f5d64d27a53bf4a64e438dc757f9b7e1e96b9aaaab2.
//
// Solidity: event Fee(uint256 indexed instrument, address indexed account, uint256 settlementFee)
func (_Instruments *InstrumentsFilterer) ParseFee(log types.Log) (*InstrumentsFee, error) {
	event := new(InstrumentsFee)
	if err := _Instruments.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsFundingIterator is returned from FilterFunding and is used to iterate over the raw logs and unpacked data for Funding events raised by the Instruments contract.
type InstrumentsFundingIterator struct {
	Event *InstrumentsFunding // Event containing the contract specifics and raw log

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
func (it *InstrumentsFundingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsFunding)
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
		it.Event = new(InstrumentsFunding)
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
func (it *InstrumentsFundingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsFundingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsFunding represents a Funding event raised by the Instruments contract.
type InstrumentsFunding struct {
	Instrument  *big.Int
	Epoch       *big.Int
	Oi          *big.Int
	Price       *big.Int
	FundingRate *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFunding is a free log retrieval operation binding the contract event 0xdf8220724a6098d8315f430588d7692863861c4805f1006a31d7c5bbabc48426.
//
// Solidity: event Funding(uint256 indexed instrument, uint256 indexed epoch, uint256 oi, uint256 price, int256 fundingRate)
func (_Instruments *InstrumentsFilterer) FilterFunding(opts *bind.FilterOpts, instrument []*big.Int, epoch []*big.Int) (*InstrumentsFundingIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "Funding", instrumentRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsFundingIterator{contract: _Instruments.contract, event: "Funding", logs: logs, sub: sub}, nil
}

// WatchFunding is a free log subscription operation binding the contract event 0xdf8220724a6098d8315f430588d7692863861c4805f1006a31d7c5bbabc48426.
//
// Solidity: event Funding(uint256 indexed instrument, uint256 indexed epoch, uint256 oi, uint256 price, int256 fundingRate)
func (_Instruments *InstrumentsFilterer) WatchFunding(opts *bind.WatchOpts, sink chan<- *InstrumentsFunding, instrument []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "Funding", instrumentRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsFunding)
				if err := _Instruments.contract.UnpackLog(event, "Funding", log); err != nil {
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

// ParseFunding is a log parse operation binding the contract event 0xdf8220724a6098d8315f430588d7692863861c4805f1006a31d7c5bbabc48426.
//
// Solidity: event Funding(uint256 indexed instrument, uint256 indexed epoch, uint256 oi, uint256 price, int256 fundingRate)
func (_Instruments *InstrumentsFilterer) ParseFunding(log types.Log) (*InstrumentsFunding, error) {
	event := new(InstrumentsFunding)
	if err := _Instruments.contract.UnpackLog(event, "Funding", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsInstrumentAccountAddedIterator is returned from FilterInstrumentAccountAdded and is used to iterate over the raw logs and unpacked data for InstrumentAccountAdded events raised by the Instruments contract.
type InstrumentsInstrumentAccountAddedIterator struct {
	Event *InstrumentsInstrumentAccountAdded // Event containing the contract specifics and raw log

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
func (it *InstrumentsInstrumentAccountAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsInstrumentAccountAdded)
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
		it.Event = new(InstrumentsInstrumentAccountAdded)
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
func (it *InstrumentsInstrumentAccountAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsInstrumentAccountAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsInstrumentAccountAdded represents a InstrumentAccountAdded event raised by the Instruments contract.
type InstrumentsInstrumentAccountAdded struct {
	Instrument *big.Int
	Account    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInstrumentAccountAdded is a free log retrieval operation binding the contract event 0xd56cbb684c2af0cb6004dfc76579372e9989c539a6140d9622ab4e4264f1f75f.
//
// Solidity: event InstrumentAccountAdded(uint256 indexed instrument, address indexed account)
func (_Instruments *InstrumentsFilterer) FilterInstrumentAccountAdded(opts *bind.FilterOpts, instrument []*big.Int, account []common.Address) (*InstrumentsInstrumentAccountAddedIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "InstrumentAccountAdded", instrumentRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsInstrumentAccountAddedIterator{contract: _Instruments.contract, event: "InstrumentAccountAdded", logs: logs, sub: sub}, nil
}

// WatchInstrumentAccountAdded is a free log subscription operation binding the contract event 0xd56cbb684c2af0cb6004dfc76579372e9989c539a6140d9622ab4e4264f1f75f.
//
// Solidity: event InstrumentAccountAdded(uint256 indexed instrument, address indexed account)
func (_Instruments *InstrumentsFilterer) WatchInstrumentAccountAdded(opts *bind.WatchOpts, sink chan<- *InstrumentsInstrumentAccountAdded, instrument []*big.Int, account []common.Address) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "InstrumentAccountAdded", instrumentRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsInstrumentAccountAdded)
				if err := _Instruments.contract.UnpackLog(event, "InstrumentAccountAdded", log); err != nil {
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

// ParseInstrumentAccountAdded is a log parse operation binding the contract event 0xd56cbb684c2af0cb6004dfc76579372e9989c539a6140d9622ab4e4264f1f75f.
//
// Solidity: event InstrumentAccountAdded(uint256 indexed instrument, address indexed account)
func (_Instruments *InstrumentsFilterer) ParseInstrumentAccountAdded(log types.Log) (*InstrumentsInstrumentAccountAdded, error) {
	event := new(InstrumentsInstrumentAccountAdded)
	if err := _Instruments.contract.UnpackLog(event, "InstrumentAccountAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsInstrumentTransferredIterator is returned from FilterInstrumentTransferred and is used to iterate over the raw logs and unpacked data for InstrumentTransferred events raised by the Instruments contract.
type InstrumentsInstrumentTransferredIterator struct {
	Event *InstrumentsInstrumentTransferred // Event containing the contract specifics and raw log

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
func (it *InstrumentsInstrumentTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsInstrumentTransferred)
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
		it.Event = new(InstrumentsInstrumentTransferred)
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
func (it *InstrumentsInstrumentTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsInstrumentTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsInstrumentTransferred represents a InstrumentTransferred event raised by the Instruments contract.
type InstrumentsInstrumentTransferred struct {
	Instrument   *big.Int
	From         common.Address
	To           common.Address
	Amount       *big.Int
	OpenInterest *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInstrumentTransferred is a free log retrieval operation binding the contract event 0xea4600feed794ec67dfe54a044b2501a9ec05001c578ee10a1f347d26d8ef46c.
//
// Solidity: event InstrumentTransferred(uint256 indexed instrument, address indexed from, address indexed to, uint256 amount, uint256 openInterest)
func (_Instruments *InstrumentsFilterer) FilterInstrumentTransferred(opts *bind.FilterOpts, instrument []*big.Int, from []common.Address, to []common.Address) (*InstrumentsInstrumentTransferredIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "InstrumentTransferred", instrumentRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsInstrumentTransferredIterator{contract: _Instruments.contract, event: "InstrumentTransferred", logs: logs, sub: sub}, nil
}

// WatchInstrumentTransferred is a free log subscription operation binding the contract event 0xea4600feed794ec67dfe54a044b2501a9ec05001c578ee10a1f347d26d8ef46c.
//
// Solidity: event InstrumentTransferred(uint256 indexed instrument, address indexed from, address indexed to, uint256 amount, uint256 openInterest)
func (_Instruments *InstrumentsFilterer) WatchInstrumentTransferred(opts *bind.WatchOpts, sink chan<- *InstrumentsInstrumentTransferred, instrument []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "InstrumentTransferred", instrumentRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsInstrumentTransferred)
				if err := _Instruments.contract.UnpackLog(event, "InstrumentTransferred", log); err != nil {
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

// ParseInstrumentTransferred is a log parse operation binding the contract event 0xea4600feed794ec67dfe54a044b2501a9ec05001c578ee10a1f347d26d8ef46c.
//
// Solidity: event InstrumentTransferred(uint256 indexed instrument, address indexed from, address indexed to, uint256 amount, uint256 openInterest)
func (_Instruments *InstrumentsFilterer) ParseInstrumentTransferred(log types.Log) (*InstrumentsInstrumentTransferred, error) {
	event := new(InstrumentsInstrumentTransferred)
	if err := _Instruments.contract.UnpackLog(event, "InstrumentTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsKeeperUpdatedIterator is returned from FilterKeeperUpdated and is used to iterate over the raw logs and unpacked data for KeeperUpdated events raised by the Instruments contract.
type InstrumentsKeeperUpdatedIterator struct {
	Event *InstrumentsKeeperUpdated // Event containing the contract specifics and raw log

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
func (it *InstrumentsKeeperUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsKeeperUpdated)
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
		it.Event = new(InstrumentsKeeperUpdated)
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
func (it *InstrumentsKeeperUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsKeeperUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsKeeperUpdated represents a KeeperUpdated event raised by the Instruments contract.
type InstrumentsKeeperUpdated struct {
	Keeper  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeeperUpdated is a free log retrieval operation binding the contract event 0x786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c00.
//
// Solidity: event KeeperUpdated(address indexed keeper, bool indexed allowed)
func (_Instruments *InstrumentsFilterer) FilterKeeperUpdated(opts *bind.FilterOpts, keeper []common.Address, allowed []bool) (*InstrumentsKeeperUpdatedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "KeeperUpdated", keeperRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsKeeperUpdatedIterator{contract: _Instruments.contract, event: "KeeperUpdated", logs: logs, sub: sub}, nil
}

// WatchKeeperUpdated is a free log subscription operation binding the contract event 0x786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c00.
//
// Solidity: event KeeperUpdated(address indexed keeper, bool indexed allowed)
func (_Instruments *InstrumentsFilterer) WatchKeeperUpdated(opts *bind.WatchOpts, sink chan<- *InstrumentsKeeperUpdated, keeper []common.Address, allowed []bool) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "KeeperUpdated", keeperRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsKeeperUpdated)
				if err := _Instruments.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
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
func (_Instruments *InstrumentsFilterer) ParseKeeperUpdated(log types.Log) (*InstrumentsKeeperUpdated, error) {
	event := new(InstrumentsKeeperUpdated)
	if err := _Instruments.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsOpenInterestUpdatedIterator is returned from FilterOpenInterestUpdated and is used to iterate over the raw logs and unpacked data for OpenInterestUpdated events raised by the Instruments contract.
type InstrumentsOpenInterestUpdatedIterator struct {
	Event *InstrumentsOpenInterestUpdated // Event containing the contract specifics and raw log

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
func (it *InstrumentsOpenInterestUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsOpenInterestUpdated)
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
		it.Event = new(InstrumentsOpenInterestUpdated)
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
func (it *InstrumentsOpenInterestUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsOpenInterestUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsOpenInterestUpdated represents a OpenInterestUpdated event raised by the Instruments contract.
type InstrumentsOpenInterestUpdated struct {
	Instrument *big.Int
	Oi         *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOpenInterestUpdated is a free log retrieval operation binding the contract event 0xec9d3d15d27b45ec272eca821314dbc7dfac39e122618a8cbfb23ca99d19c883.
//
// Solidity: event OpenInterestUpdated(uint256 indexed instrument, uint256 oi)
func (_Instruments *InstrumentsFilterer) FilterOpenInterestUpdated(opts *bind.FilterOpts, instrument []*big.Int) (*InstrumentsOpenInterestUpdatedIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "OpenInterestUpdated", instrumentRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsOpenInterestUpdatedIterator{contract: _Instruments.contract, event: "OpenInterestUpdated", logs: logs, sub: sub}, nil
}

// WatchOpenInterestUpdated is a free log subscription operation binding the contract event 0xec9d3d15d27b45ec272eca821314dbc7dfac39e122618a8cbfb23ca99d19c883.
//
// Solidity: event OpenInterestUpdated(uint256 indexed instrument, uint256 oi)
func (_Instruments *InstrumentsFilterer) WatchOpenInterestUpdated(opts *bind.WatchOpts, sink chan<- *InstrumentsOpenInterestUpdated, instrument []*big.Int) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "OpenInterestUpdated", instrumentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsOpenInterestUpdated)
				if err := _Instruments.contract.UnpackLog(event, "OpenInterestUpdated", log); err != nil {
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

// ParseOpenInterestUpdated is a log parse operation binding the contract event 0xec9d3d15d27b45ec272eca821314dbc7dfac39e122618a8cbfb23ca99d19c883.
//
// Solidity: event OpenInterestUpdated(uint256 indexed instrument, uint256 oi)
func (_Instruments *InstrumentsFilterer) ParseOpenInterestUpdated(log types.Log) (*InstrumentsOpenInterestUpdated, error) {
	event := new(InstrumentsOpenInterestUpdated)
	if err := _Instruments.contract.UnpackLog(event, "OpenInterestUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsOptionAddedIterator is returned from FilterOptionAdded and is used to iterate over the raw logs and unpacked data for OptionAdded events raised by the Instruments contract.
type InstrumentsOptionAddedIterator struct {
	Event *InstrumentsOptionAdded // Event containing the contract specifics and raw log

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
func (it *InstrumentsOptionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsOptionAdded)
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
		it.Event = new(InstrumentsOptionAdded)
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
func (it *InstrumentsOptionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsOptionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsOptionAdded represents a OptionAdded event raised by the Instruments contract.
type InstrumentsOptionAdded struct {
	Instrument *big.Int
	Asset      common.Address
	Expiry     *big.Int
	IsPut      bool
	Strike     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOptionAdded is a free log retrieval operation binding the contract event 0x8a3e5ee639fc2822a7d3969b04dabf44aa1249f1565d8b09b4940c64a1fa24d6.
//
// Solidity: event OptionAdded(uint256 indexed instrument, address indexed asset, uint256 indexed expiry, bool isPut, uint256 strike)
func (_Instruments *InstrumentsFilterer) FilterOptionAdded(opts *bind.FilterOpts, instrument []*big.Int, asset []common.Address, expiry []*big.Int) (*InstrumentsOptionAddedIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var expiryRule []interface{}
	for _, expiryItem := range expiry {
		expiryRule = append(expiryRule, expiryItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "OptionAdded", instrumentRule, assetRule, expiryRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsOptionAddedIterator{contract: _Instruments.contract, event: "OptionAdded", logs: logs, sub: sub}, nil
}

// WatchOptionAdded is a free log subscription operation binding the contract event 0x8a3e5ee639fc2822a7d3969b04dabf44aa1249f1565d8b09b4940c64a1fa24d6.
//
// Solidity: event OptionAdded(uint256 indexed instrument, address indexed asset, uint256 indexed expiry, bool isPut, uint256 strike)
func (_Instruments *InstrumentsFilterer) WatchOptionAdded(opts *bind.WatchOpts, sink chan<- *InstrumentsOptionAdded, instrument []*big.Int, asset []common.Address, expiry []*big.Int) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var expiryRule []interface{}
	for _, expiryItem := range expiry {
		expiryRule = append(expiryRule, expiryItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "OptionAdded", instrumentRule, assetRule, expiryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsOptionAdded)
				if err := _Instruments.contract.UnpackLog(event, "OptionAdded", log); err != nil {
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

// ParseOptionAdded is a log parse operation binding the contract event 0x8a3e5ee639fc2822a7d3969b04dabf44aa1249f1565d8b09b4940c64a1fa24d6.
//
// Solidity: event OptionAdded(uint256 indexed instrument, address indexed asset, uint256 indexed expiry, bool isPut, uint256 strike)
func (_Instruments *InstrumentsFilterer) ParseOptionAdded(log types.Log) (*InstrumentsOptionAdded, error) {
	event := new(InstrumentsOptionAdded)
	if err := _Instruments.contract.UnpackLog(event, "OptionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsOwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the Instruments contract.
type InstrumentsOwnerUpdatedIterator struct {
	Event *InstrumentsOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *InstrumentsOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsOwnerUpdated)
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
		it.Event = new(InstrumentsOwnerUpdated)
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
func (it *InstrumentsOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsOwnerUpdated represents a OwnerUpdated event raised by the Instruments contract.
type InstrumentsOwnerUpdated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed newOwner)
func (_Instruments *InstrumentsFilterer) FilterOwnerUpdated(opts *bind.FilterOpts, newOwner []common.Address) (*InstrumentsOwnerUpdatedIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "OwnerUpdated", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsOwnerUpdatedIterator{contract: _Instruments.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed newOwner)
func (_Instruments *InstrumentsFilterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *InstrumentsOwnerUpdated, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "OwnerUpdated", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsOwnerUpdated)
				if err := _Instruments.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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
func (_Instruments *InstrumentsFilterer) ParseOwnerUpdated(log types.Log) (*InstrumentsOwnerUpdated, error) {
	event := new(InstrumentsOwnerUpdated)
	if err := _Instruments.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsPayoutIterator is returned from FilterPayout and is used to iterate over the raw logs and unpacked data for Payout events raised by the Instruments contract.
type InstrumentsPayoutIterator struct {
	Event *InstrumentsPayout // Event containing the contract specifics and raw log

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
func (it *InstrumentsPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsPayout)
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
		it.Event = new(InstrumentsPayout)
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
func (it *InstrumentsPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsPayout represents a Payout event raised by the Instruments contract.
type InstrumentsPayout struct {
	Instrument *big.Int
	Account    common.Address
	Epoch      *big.Int
	Price      *big.Int
	Position   *big.Int
	Amount     *big.Int
	Fee        *big.Int
	Rate       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPayout is a free log retrieval operation binding the contract event 0x9564f1ed9c02902fc9b9d1d3a30c1bb7320ac24c0d4454bcb53ec71f67f29264.
//
// Solidity: event Payout(uint256 indexed instrument, address indexed account, uint256 epoch, uint256 price, int256 position, int256 amount, uint256 fee, int256 rate)
func (_Instruments *InstrumentsFilterer) FilterPayout(opts *bind.FilterOpts, instrument []*big.Int, account []common.Address) (*InstrumentsPayoutIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "Payout", instrumentRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsPayoutIterator{contract: _Instruments.contract, event: "Payout", logs: logs, sub: sub}, nil
}

// WatchPayout is a free log subscription operation binding the contract event 0x9564f1ed9c02902fc9b9d1d3a30c1bb7320ac24c0d4454bcb53ec71f67f29264.
//
// Solidity: event Payout(uint256 indexed instrument, address indexed account, uint256 epoch, uint256 price, int256 position, int256 amount, uint256 fee, int256 rate)
func (_Instruments *InstrumentsFilterer) WatchPayout(opts *bind.WatchOpts, sink chan<- *InstrumentsPayout, instrument []*big.Int, account []common.Address) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "Payout", instrumentRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsPayout)
				if err := _Instruments.contract.UnpackLog(event, "Payout", log); err != nil {
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

// ParsePayout is a log parse operation binding the contract event 0x9564f1ed9c02902fc9b9d1d3a30c1bb7320ac24c0d4454bcb53ec71f67f29264.
//
// Solidity: event Payout(uint256 indexed instrument, address indexed account, uint256 epoch, uint256 price, int256 position, int256 amount, uint256 fee, int256 rate)
func (_Instruments *InstrumentsFilterer) ParsePayout(log types.Log) (*InstrumentsPayout, error) {
	event := new(InstrumentsPayout)
	if err := _Instruments.contract.UnpackLog(event, "Payout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsPerpetualAddedIterator is returned from FilterPerpetualAdded and is used to iterate over the raw logs and unpacked data for PerpetualAdded events raised by the Instruments contract.
type InstrumentsPerpetualAddedIterator struct {
	Event *InstrumentsPerpetualAdded // Event containing the contract specifics and raw log

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
func (it *InstrumentsPerpetualAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsPerpetualAdded)
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
		it.Event = new(InstrumentsPerpetualAdded)
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
func (it *InstrumentsPerpetualAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsPerpetualAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsPerpetualAdded represents a PerpetualAdded event raised by the Instruments contract.
type InstrumentsPerpetualAdded struct {
	Instrument *big.Int
	Asset      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPerpetualAdded is a free log retrieval operation binding the contract event 0x65f3016ffa801ac65a87ed4a8c83abfe84ba1f2ffd38229df8912360535a0339.
//
// Solidity: event PerpetualAdded(uint256 indexed instrument, address indexed asset)
func (_Instruments *InstrumentsFilterer) FilterPerpetualAdded(opts *bind.FilterOpts, instrument []*big.Int, asset []common.Address) (*InstrumentsPerpetualAddedIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "PerpetualAdded", instrumentRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsPerpetualAddedIterator{contract: _Instruments.contract, event: "PerpetualAdded", logs: logs, sub: sub}, nil
}

// WatchPerpetualAdded is a free log subscription operation binding the contract event 0x65f3016ffa801ac65a87ed4a8c83abfe84ba1f2ffd38229df8912360535a0339.
//
// Solidity: event PerpetualAdded(uint256 indexed instrument, address indexed asset)
func (_Instruments *InstrumentsFilterer) WatchPerpetualAdded(opts *bind.WatchOpts, sink chan<- *InstrumentsPerpetualAdded, instrument []*big.Int, asset []common.Address) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "PerpetualAdded", instrumentRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsPerpetualAdded)
				if err := _Instruments.contract.UnpackLog(event, "PerpetualAdded", log); err != nil {
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

// ParsePerpetualAdded is a log parse operation binding the contract event 0x65f3016ffa801ac65a87ed4a8c83abfe84ba1f2ffd38229df8912360535a0339.
//
// Solidity: event PerpetualAdded(uint256 indexed instrument, address indexed asset)
func (_Instruments *InstrumentsFilterer) ParsePerpetualAdded(log types.Log) (*InstrumentsPerpetualAdded, error) {
	event := new(InstrumentsPerpetualAdded)
	if err := _Instruments.contract.UnpackLog(event, "PerpetualAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsSettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the Instruments contract.
type InstrumentsSettlementIterator struct {
	Event *InstrumentsSettlement // Event containing the contract specifics and raw log

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
func (it *InstrumentsSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsSettlement)
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
		it.Event = new(InstrumentsSettlement)
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
func (it *InstrumentsSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsSettlement represents a Settlement event raised by the Instruments contract.
type InstrumentsSettlement struct {
	Instrument *big.Int
	Oi         *big.Int
	Price      *big.Int
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0x0f3455c580c87ae485cf1fd9108e0d31f59b529eb91f338403886bd23c9d96ac.
//
// Solidity: event Settlement(uint256 indexed instrument, uint256 oi, uint256 price, uint256 fee)
func (_Instruments *InstrumentsFilterer) FilterSettlement(opts *bind.FilterOpts, instrument []*big.Int) (*InstrumentsSettlementIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "Settlement", instrumentRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsSettlementIterator{contract: _Instruments.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0x0f3455c580c87ae485cf1fd9108e0d31f59b529eb91f338403886bd23c9d96ac.
//
// Solidity: event Settlement(uint256 indexed instrument, uint256 oi, uint256 price, uint256 fee)
func (_Instruments *InstrumentsFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *InstrumentsSettlement, instrument []*big.Int) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "Settlement", instrumentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsSettlement)
				if err := _Instruments.contract.UnpackLog(event, "Settlement", log); err != nil {
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

// ParseSettlement is a log parse operation binding the contract event 0x0f3455c580c87ae485cf1fd9108e0d31f59b529eb91f338403886bd23c9d96ac.
//
// Solidity: event Settlement(uint256 indexed instrument, uint256 oi, uint256 price, uint256 fee)
func (_Instruments *InstrumentsFilterer) ParseSettlement(log types.Log) (*InstrumentsSettlement, error) {
	event := new(InstrumentsSettlement)
	if err := _Instruments.contract.UnpackLog(event, "Settlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Instruments contract.
type InstrumentsUpgradedIterator struct {
	Event *InstrumentsUpgraded // Event containing the contract specifics and raw log

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
func (it *InstrumentsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsUpgraded)
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
		it.Event = new(InstrumentsUpgraded)
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
func (it *InstrumentsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsUpgraded represents a Upgraded event raised by the Instruments contract.
type InstrumentsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Instruments *InstrumentsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*InstrumentsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsUpgradedIterator{contract: _Instruments.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Instruments *InstrumentsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *InstrumentsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsUpgraded)
				if err := _Instruments.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Instruments *InstrumentsFilterer) ParseUpgraded(log types.Log) (*InstrumentsUpgraded, error) {
	event := new(InstrumentsUpgraded)
	if err := _Instruments.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstrumentsVersionInitializedIterator is returned from FilterVersionInitialized and is used to iterate over the raw logs and unpacked data for VersionInitialized events raised by the Instruments contract.
type InstrumentsVersionInitializedIterator struct {
	Event *InstrumentsVersionInitialized // Event containing the contract specifics and raw log

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
func (it *InstrumentsVersionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstrumentsVersionInitialized)
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
		it.Event = new(InstrumentsVersionInitialized)
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
func (it *InstrumentsVersionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstrumentsVersionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstrumentsVersionInitialized represents a VersionInitialized event raised by the Instruments contract.
type InstrumentsVersionInitialized struct {
	NewVersion *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVersionInitialized is a free log retrieval operation binding the contract event 0x7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d.
//
// Solidity: event VersionInitialized(uint256 indexed newVersion)
func (_Instruments *InstrumentsFilterer) FilterVersionInitialized(opts *bind.FilterOpts, newVersion []*big.Int) (*InstrumentsVersionInitializedIterator, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Instruments.contract.FilterLogs(opts, "VersionInitialized", newVersionRule)
	if err != nil {
		return nil, err
	}
	return &InstrumentsVersionInitializedIterator{contract: _Instruments.contract, event: "VersionInitialized", logs: logs, sub: sub}, nil
}

// WatchVersionInitialized is a free log subscription operation binding the contract event 0x7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d.
//
// Solidity: event VersionInitialized(uint256 indexed newVersion)
func (_Instruments *InstrumentsFilterer) WatchVersionInitialized(opts *bind.WatchOpts, sink chan<- *InstrumentsVersionInitialized, newVersion []*big.Int) (event.Subscription, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Instruments.contract.WatchLogs(opts, "VersionInitialized", newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstrumentsVersionInitialized)
				if err := _Instruments.contract.UnpackLog(event, "VersionInitialized", log); err != nil {
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
func (_Instruments *InstrumentsFilterer) ParseVersionInitialized(log types.Log) (*InstrumentsVersionInitialized, error) {
	event := new(InstrumentsVersionInitialized)
	if err := _Instruments.contract.UnpackLog(event, "VersionInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
