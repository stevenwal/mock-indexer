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

// AccountsInterfaceBalance is an auto generated low-level Go binding around an user-defined struct.
type AccountsInterfaceBalance struct {
	Collateral common.Address
	Balance    *big.Int
}

// AccountsInterfacePosition is an auto generated low-level Go binding around an user-defined struct.
type AccountsInterfacePosition struct {
	Instrument *big.Int
	Position   *big.Int
}

// AccountsInterfaceSigningKey is an auto generated low-level Go binding around an user-defined struct.
type AccountsInterfaceSigningKey struct {
	SigningKey common.Address
	Expiry     *big.Int
}

// AccountsMetaData contains all meta data concerning the Accounts contract.
var AccountsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_domain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadRegisterSigningKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadRevokeSigningKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotKeeper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Credit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"}],\"name\":\"CreditInstrument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Debit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"}],\"name\":\"DebitInstrument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"InstrumentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"InstrumentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"KeeperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"RegisteredSigningKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"revokeHash\",\"type\":\"bytes32\"}],\"name\":\"RevokedSigningKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"VersionInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"WithdrawProxyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"addInstrument\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorities\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collaterals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"creditInstrument\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debitInstrument\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structAccountsInterface.Balance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Position[]\",\"name\":\"positions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structAccountsInterface.SigningKey[]\",\"name\":\"signingKeys\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getCollateralBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollaterals\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"getInstrumentPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInstruments\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInsuranceFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getQuoteBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getSigningKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"}],\"name\":\"hasSigningKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"instruments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keepers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signingKeySig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountSig\",\"type\":\"bytes\"}],\"name\":\"registerSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"removeCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"removeInstrument\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"resetInstrumentPosition\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"revokeSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signingKeyExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalPositions\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawProxy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateWithdrawProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawProxies\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610180604052306080523480156200001657600080fd5b50604051620058333803806200583383398101604081905262000039916200041a565b6001600160a01b038116620000615760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811660a08190526040805163313ce56760e01b8152905163313ce567916004808201926020929091908290030181865afa158015620000ac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000d29190620004f0565b60ff1660c052620000f08383620002aa602090811b6200384317901c565b60e052505060408051680aed2e8d0c8e4c2ee560bb1b6020808301919091527f6164647265737320636f6c6c61746572616c2c0000000000000000000000000060298301526a1859191c995cdcc81d1bcb60aa1b603c8301526e1d5a5b9d0c8d4d88185b5bdd5b9d0b608a1b60478301526c1d5a5b9d0c8d4d881cd85b1d0b609a1b60568301526b62797465733332206461746160a01b6063830152602960f81b606f83018190528351605081850301815260708401855280519083012061010052670a6d2cedc96caf2560c31b60908401526e1859191c995cdcc81858d8dbdd5b9d608a1b609884015260a783018190528351608881850301815260a88401855280519083012061012052680a4caced2e6e8cae4560bb1b60c88401526b1859191c995cdcc81ad95e4b60a21b60d18401526d75696e743235362065787069727960901b60dd84015260eb8301819052835160cc81850301815260ec8401855280519083012061014052660a4caecded6ca560cb1b61010c8401526a61646472657373206b657960a81b61011384015261011e830152825160ff81840301815261011f909201909252805191012061016052506200053a565b6040516c08a92a06e626488dedac2d2dc5609b1b60208201526b1cdd1c9a5b99c81b985b594b60a21b602d8201526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398201526e1d5a5b9d0c8d4d8818da185a5b9259608a1b6048820152602960f81b605782015260009060580160405160208183030381529060405280519060200120836040516020016200034191906200051c565b60408051601f198184030181528282528051602091820120908301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201526080810183905260a00160405160208183030381529060405280519060200120905092915050565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620003e7578181015183820152602001620003cd565b83811115620003f7576000848401525b50505050565b80516001600160a01b03811681146200041557600080fd5b919050565b6000806000606084860312156200043057600080fd5b83516001600160401b03808211156200044857600080fd5b818601915086601f8301126200045d57600080fd5b815181811115620004725762000472620003b4565b604051601f8201601f19908116603f011681019083821181831017156200049d576200049d620003b4565b81604052828152896020848701011115620004b757600080fd5b620004ca836020830160208801620003ca565b809750505050505060208401519150620004e760408501620003fd565b90509250925092565b6000602082840312156200050357600080fd5b815160ff811681146200051557600080fd5b9392505050565b6000825162000530818460208701620003ca565b9190910192915050565b60805160a05160c05160e05161010051610120516101405161016051615246620005ed6000396000610b0f015260006117da01526000611a2001526000613e09015260008181610a7801528181610b5a0152818161182c015281816119ff0152613de80152600061052b0152600081816108070152818161088f015281816129c30152612b440152600081816115a90152818161165901528181611c2601528181611cd60152611e1f01526152466000f3fe6080604052600436106102f25760003560e01c806372d39ed61161018f578063b7d5820b116100e1578063d658d2e91161008a578063f0d2d5a811610064578063f0d2d5a814610a46578063f698da2514610a66578063fbcbc0f114610a9a57600080fd5b8063d658d2e9146109c6578063e30e8384146109f6578063eeb97d3b14610a1657600080fd5b8063c99d3a06116100bb578063c99d3a0614610966578063d305787714610986578063d5c20730146109a657600080fd5b8063b7d5820b146108bf578063b8eb94e914610916578063c4d66de81461094657600080fd5b80638da5cb5b11610143578063999b93af1161011d578063999b93af146107f55780639aeddeff14610829578063af1aa3b11461084957600080fd5b80638da5cb5b1461076b5780638e9916061461079857806391223d69146107c557600080fd5b80637de182c5116101745780637de182c51461070b5780637df5b7011461072b5780638340f5491461074b57600080fd5b806372d39ed6146106d657806378b92636146106f657600080fd5b80633fd1e2bd1161024857806352d1902d116101fc57806362946d3b116101d657806362946d3b1461062f5780636cd22eaf146106895780636eacd398146106a957600080fd5b806352d1902d146105c057806361d21920146105d55780636205ed5e1461060257600080fd5b806348e4ccbe1161022d57806348e4ccbe1461055f57806349ad755b1461057f5780634f1ef286146105ad57600080fd5b80633fd1e2bd146105195780634614be9f1461039857600080fd5b80631dcc7d5c116102aa578063262709e611610284578063262709e6146104a95780633659cfe6146104c95780633bbd64bc146104e957600080fd5b80631dcc7d5c1461040e57806322bbad0b1461043c578063241d400d1461047c57600080fd5b806313af4035116102db57806313af403514610378578063158626f7146103985780631635717c146103ec57600080fd5b806306fdde03146102f757806307b1b8b214610356575b600080fd5b34801561030357600080fd5b506103406040518060400160405280600881526020017f4163636f756e747300000000000000000000000000000000000000000000000081525081565b60405161034d91906149eb565b60405180910390f35b34801561036257600080fd5b50610376610371366004614b01565b610ac9565b005b34801561038457600080fd5b50610376610393366004614b4f565b6110af565b3480156103a457600080fd5b507f7d551b4e977cb76b47865bfb6cdc77af264be2349f98f5d2cb58a453757caae75b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161034d565b3480156103f857600080fd5b5061040161110c565b60405161034d9190614b6a565b34801561041a57600080fd5b5061042e610429366004614bae565b611164565b60405161034d929190614c1c565b34801561044857600080fd5b5061046c610457366004614c30565b600a6020526000908152604090205460ff1681565b604051901515815260200161034d565b34801561048857600080fd5b5061049c610497366004614b4f565b611345565b60405161034d9190614c49565b3480156104b557600080fd5b5061042e6104c4366004614c97565b6113d5565b3480156104d557600080fd5b506103766104e4366004614b4f565b611592565b3480156104f557600080fd5b5061046c610504366004614b4f565b60046020526000908152604090205460ff1681565b34801561052557600080fd5b5061054d7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff909116815260200161034d565b34801561056b57600080fd5b5061037661057a366004614cd3565b611794565b34801561058b57600080fd5b5061059f61059a366004614d51565b611bd4565b60405190815260200161034d565b6103766105bb366004614b01565b611c0f565b3480156105cc57600080fd5b5061059f611e05565b3480156105e157600080fd5b5061059f6105f0366004614c30565b600e6020526000908152604090205481565b34801561060e57600080fd5b5061059f61061d366004614b4f565b600c6020526000908152604090205481565b34801561063b57600080fd5b5061046c61064a366004614d7b565b73ffffffffffffffffffffffffffffffffffffffff9182166000908152600b602090815260408083209390941682526001909201909152205460ff1690565b34801561069557600080fd5b506103766106a4366004614dae565b611ef1565b3480156106b557600080fd5b5061059f6106c4366004614b4f565b600d6020526000908152604090205481565b3480156106e257600080fd5b5061059f6106f1366004614dea565b61200e565b34801561070257600080fd5b5061049c612153565b34801561071757600080fd5b5061042e610726366004614c97565b6121c1565b34801561073757600080fd5b5061042e610746366004614bae565b612378565b34801561075757600080fd5b5061059f610766366004614c97565b61253f565b34801561077757600080fd5b506000546103c79073ffffffffffffffffffffffffffffffffffffffff1681565b3480156107a457600080fd5b506107b86107b3366004614c30565b6125bf565b60405161034d9190614e8c565b3480156107d157600080fd5b5061046c6107e0366004614b4f565b60036020526000908152604090205460ff1681565b34801561080157600080fd5b506103c77f000000000000000000000000000000000000000000000000000000000000000081565b34801561083557600080fd5b506107b8610844366004614d51565b612717565b34801561085557600080fd5b5061059f610864366004614b4f565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600b602090815260408083207f0000000000000000000000000000000000000000000000000000000000000000909416835260029093019052205490565b3480156108cb57600080fd5b5061059f6108da366004614d7b565b73ffffffffffffffffffffffffffffffffffffffff9182166000908152600b602090815260408083209390941682526002909201909152205490565b34801561092257600080fd5b5061046c610931366004614b4f565b60056020526000908152604090205460ff1681565b34801561095257600080fd5b50610376610961366004614b4f565b612912565b34801561097257600080fd5b50610376610981366004614b4f565b612a92565b34801561099257600080fd5b506103766109a1366004614dae565b612e34565b3480156109b257600080fd5b506103766109c1366004614dae565b612f51565b3480156109d257600080fd5b5061046c6109e1366004614c30565b60066020526000908152604090205460ff1681565b348015610a0257600080fd5b506107b8610a11366004614c30565b61306e565b348015610a2257600080fd5b5061046c610a31366004614b4f565b60086020526000908152604090205460ff1681565b348015610a5257600080fd5b50610376610a61366004614b4f565b6132ab565b348015610a7257600080fd5b5061059f7f000000000000000000000000000000000000000000000000000000000000000081565b348015610aa657600080fd5b50610aba610ab5366004614b4f565b61347b565b60405161034d93929190614eff565b600254600114610b05576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028055604080517f0000000000000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff841691810191909152600090610bcf907f0000000000000000000000000000000000000000000000000000000000000000906060015b604051602081830303815290604052805190602001207f1901000000000000000000000000000000000000000000000000000000000000600090815260029290925260229081526042822091905290565b3360009081526004602052604090205490915060ff16610c29578060016040517fff524b1c000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316610c7b578060036040517fff524b1c000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600c60205260408120549003610cde5780600a6040517fff524b1c000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b6000610cea82846139c2565b905073ffffffffffffffffffffffffffffffffffffffff8116610d3e578160076040517fff524b1c000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600b60209081526040808320938816835260019093019052205460ff16610db35781600a6040517fff524b1c000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600b6020908152604080832080548251818502810185019093528083528493830182828015610e3457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610e09575b505050505090505b8051821015610e9d578573ffffffffffffffffffffffffffffffffffffffff16818381518110610e6e57610e6e614fc2565b602002602001015173ffffffffffffffffffffffffffffffffffffffff160315610e9d57816001019150610e3c565b80518210610edc5783600a6040517fff524b1c000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b8060018251610eeb9190615020565b81518110610efb57610efb614fc2565b6020026020010151600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018381548110610f5657610f56614fc2565b600091825260208083209190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9485161790559185168152600b90915260409020805480610fc057610fc0615037565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff858116808452600b83526040808520928b1680865260019390930190935282842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055915187939192917ff0d72ac27c7920d0a69f77bbf64f8509a0307c49e54eee377861c46101ef25cb91a45050600160025550505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611100576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61110981613b1c565b50565b6060600980548060200260200160405190810160405280929190818152602001828054801561115a57602002820191906000526020600020905b815481526020019060010190808311611146575b5050505050905090565b33600090815260036020526040812054819060ff166111af576040517f3121268600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001146111eb576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280556000848152600a602052604090205460ff166112115750600090506005611336565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156112455750600090506002611336565b73ffffffffffffffffffffffffffffffffffffffff85166000908152600b60209081526040808320878452600301909152812054611284908590615066565b73ffffffffffffffffffffffffffffffffffffffff87166000908152600b602090815260408083208984526003018252808320849055600e9091528120805492935086929091906112d6908490615066565b90915550506040805185815260208101839052869173ffffffffffffffffffffffffffffffffffffffff8916917f8437001e64ef3652dd8a3935534d9d6cfdb1fd4cbe6db8454c3369845cb45c5e91015b60405180910390a39150600090505b60016002559094909350915050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600b60209081526040918290208054835181840281018401909452808452606093928301828280156113c957602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161139e575b50505050509050919050565b33600090815260036020526040812054819060ff16611420576040517f3121268600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025460011461145c576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805573ffffffffffffffffffffffffffffffffffffffff841660009081526008602052604090205460ff166114995750600090506004611336565b73ffffffffffffffffffffffffffffffffffffffff8086166000908152600b602090815260408083209388168352600290930190522054838110156114e357915060029050611336565b73ffffffffffffffffffffffffffffffffffffffff8087166000908152600b602090815260408083209389168352600290930181528282209387900393849055600d9052908120805486929061153a908490615020565b9091555050604080518581526020810183905273ffffffffffffffffffffffffffffffffffffffff80881692908916917faa31acf5383e18b9a4ae2b3ccfa0faf705cf8e968ee6dd291db4b1bc1a2a8ebd9101611327565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163003611657576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c00000000000000000000000000000000000000006064820152608401610c20565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166116cc7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161461176f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f6163746976652070726f787900000000000000000000000000000000000000006064820152608401610c20565b61177881613b89565b6040805160008082526020820190925261110991839190613bda565b6002546001146117d0576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028055604080517f0000000000000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff86169181019190915260608101849052600090611854907f000000000000000000000000000000000000000000000000000000000000000090608001610b7e565b3360009081526004602052604090205490915060ff166118a5578060016040517f9ce723c2000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b73ffffffffffffffffffffffffffffffffffffffff85166118f7578060036040517f9ce723c2000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b83600003611936578060026040517f9ce723c2000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b73ffffffffffffffffffffffffffffffffffffffff85166000908152600c6020526040902054156119985780600a6040517f9ce723c2000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b60006119a482846139c2565b905073ffffffffffffffffffffffffffffffffffffffff81166119f8578160076040517f9ce723c2000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b6000611a707f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000084604051602001610b7e92919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b90506000611a7e82876139c2565b90508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611aea5783600b6040517f9ce723c2000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b73ffffffffffffffffffffffffffffffffffffffff8881166000818152600c602090815260408083208c9055938716808352600b825284832080546001808201835582865284862090910180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168717905585855290810183529285902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690931790925592518a815287937f1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac910160405180910390a450506001600255505050505050565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600b602090815260408083208484526003019091529020545b92915050565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163003611cd4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c00000000000000000000000000000000000000006064820152608401610c20565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611d497f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614611dec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f6163746976652070726f787900000000000000000000000000000000000000006064820152608401610c20565b611df582613b89565b611e0182826001613bda565b5050565b60003073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611ecc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610c20565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b60005473ffffffffffffffffffffffffffffffffffffffff163314611f42576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611f8f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660008181526003602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915590519092917fc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc91a35050565b600060025460011461204c576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280556000806120628a8a8a8a8a8a8a613dde565b909250905061208873ffffffffffffffffffffffffffffffffffffffff8b168a886142f1565b73ffffffffffffffffffffffffffffffffffffffff891660009081526005602052604090205460ff1615612141576040517ff37e8d3800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a169063f37e8d389061210e9085908e908b908b906004016150da565b600060405180830381600087803b15801561212857600080fd5b505af115801561213c573d6000803e3d6000fd5b505050505b60016002559998505050505050505050565b6060600780548060200260200160405190810160405280929190818152602001828054801561115a57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161218d575050505050905090565b33600090815260036020526040812054819060ff1661220c576040517f3121268600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254600114612248576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805573ffffffffffffffffffffffffffffffffffffffff841660009081526008602052604090205460ff166122855750600090506004611336565b73ffffffffffffffffffffffffffffffffffffffff8086166000908152600b60209081526040808320938816835260029093019052908120546122c9908590615119565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600b60209081526040808320938a16835260029093018152828220849055600d905290812080549293508692909190612320908490615119565b9091555050604080518581526020810183905273ffffffffffffffffffffffffffffffffffffffff80881692908916917f27f2026b6afc5d271934dacdaaa950d69fd5fed4982a55b6c921e1667916f06c9101611327565b33600090815260036020526040812054819060ff166123c3576040517f3121268600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001146123ff576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280556000848152600a602052604090205460ff166124255750600090506005611336565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156124595750600090506002611336565b73ffffffffffffffffffffffffffffffffffffffff85166000908152600b60209081526040808320878452600301909152812054612498908590615131565b73ffffffffffffffffffffffffffffffffffffffff87166000908152600b602090815260408083208984526003018252808320849055600e9091528120805492935086929091906124ea908490615131565b90915550506040805185815260208101839052869173ffffffffffffffffffffffffffffffffffffffff8916917f151a955d0b994ca6227ecca24a04f6bdbc5e39d3b71fba7df15fafda915a59789101611327565b600060025460011461257d576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028055600061258e8585856143b0565b90506125b273ffffffffffffffffffffffffffffffffffffffff8516333086614598565b6001600255949350505050565b3360009081526003602052604081205460ff16612608576040517f3121268600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254600114612644576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805560008290036126595750600261270d565b6000828152600a602052604090205460ff16156126785750600561270d565b6009805460018181019092557f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af018390556000838152600a602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909317909255905183917fb8597a358785eb73ad0a81079eb5e50ba0ac43dd023f47c0e0ff47e58a9495c891a25060005b6001600255919050565b3360009081526003602052604081205460ff16612760576040517f3121268600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025460011461279c576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280556000828152600a602052604090205460ff166127be57506005612907565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600b6020908152604080832085845260030190915281208054908290559081131561287a576000838152600e60205260408120805483929061281d908490615066565b90915550506040805182815260006020820152849173ffffffffffffffffffffffffffffffffffffffff8716917f8437001e64ef3652dd8a3935534d9d6cfdb1fd4cbe6db8454c3369845cb45c5e910160405180910390a3612901565b6000811215612901576000838152600e6020526040812080548392906128a1908490615066565b9091555083905073ffffffffffffffffffffffffffffffffffffffff85167f151a955d0b994ca6227ecca24a04f6bdbc5e39d3b71fba7df15fafda915a59786128e9846151a5565b60408051918252600060208301520160405180910390a35b60009150505b600160025592915050565b600161291e8180615020565b60015414612958576040517f5daa87a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001036129665760016002555b61296f82613b1c565b6007805460018082019092557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180547fffffffffffffffffffffffff0000000000000000000000000000000000000000167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690811790915560008181526008602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909417909355915190917f7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f91a2600181905560405181907f7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d90600090a25050565b60005473ffffffffffffffffffffffffffffffffffffffff163314612ae3576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811660009081526008602052604090205460ff16612b42576040517fd1ef4cea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612bc7576040517fd1ef4cea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806007805480602002602001604051908101604052809291908181526020018280548015612c2d57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311612c02575b505050505090505b8051821015612c96578273ffffffffffffffffffffffffffffffffffffffff16818381518110612c6757612c67614fc2565b602002602001015173ffffffffffffffffffffffffffffffffffffffff160315612c9657816001019150612c35565b80518210612cd0576040517fd1ef4cea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060018251612cdf9190615020565b81518110612cef57612cef614fc2565b602002602001015160078381548110612d0a57612d0a614fc2565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506007805480612d6357612d63615037565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff85168083526008909152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555190917fd89d2ee68ab04dca0193f48a4aff55e20fa5ec0429a8a8c1c51b8dad6178a59391a2505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314612e85576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216612ed2576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660008181526004602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915590519092917f786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c0091a35050565b60005473ffffffffffffffffffffffffffffffffffffffff163314612fa2576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216612fef576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660008181526005602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915590519092917f097359b7e2dca2adc06bfc6c3e4e8fcc69a2e916105499e8e4c60a569598cbc091a35050565b3360009081526003602052604081205460ff166130b7576040517f3121268600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001146130f3576040517fab143c0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280556000828152600a602052604090205460ff166131155750600561270d565b600080600980548060200260200160405190810160405280929190818152602001828054801561316457602002820191906000526020600020905b815481526020019060010190808311613150575b505050505090505b80518210156131a1578381838151811061318857613188614fc2565b602002602001015103156131a15781600101915061316c565b805182106131b45760059250505061270d565b80600182516131c39190615020565b815181106131d3576131d3614fc2565b6020026020010151600983815481106131ee576131ee614fc2565b600091825260209091200155600980548061320b5761320b615037565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101839055909201909255858252600a9052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555185917f08b9963dddd96e14a775b44c82852b5a482f5060ae86706c04fc5878c3fe3d8d91a250506001600255506000919050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146132fc576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116613349576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811660009081526008602052604090205460ff16156133a9576040517fd1ef4cea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007805460018082019092557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff841690811790915560008181526008602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909417909355915190917f7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f91a250565b600754606090819081908067ffffffffffffffff81111561349e5761349e614a27565b6040519080825280602002602001820160405280156134e357816020015b60408051808201909152600080825260208201528152602001906001900390816134bc5790505b50935060005b818110156135815760006007828154811061350657613506614fc2565b600091825260208083209091015460408051808201825273ffffffffffffffffffffffffffffffffffffffff928316808252928c168552600b8452818520838652600201845293205491830191909152875190925087908490811061356d5761356d614fc2565b6020908102919091010152506001016134e9565b506009548067ffffffffffffffff81111561359e5761359e614a27565b6040519080825280602002602001820160405280156135e357816020015b60408051808201909152600080825260208201528152602001906001900390816135bc5790505b50935060005b8181101561367d5760006009828154811061360657613606614fc2565b600091825260208083209091015460408051808201825282815273ffffffffffffffffffffffffffffffffffffffff8d168552600b8452818520838652600301845293205491830191909152875190925087908490811061366957613669614fc2565b6020908102919091010152506001016135e9565b5073ffffffffffffffffffffffffffffffffffffffff86166000908152600b602090815260408083208054825181850281018501909352808352919290919083018282801561370257602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116136d7575b50505050509050805167ffffffffffffffff81111561372357613723614a27565b60405190808252806020026020018201604052801561376857816020015b60408051808201909152600080825260208201528152602001906001900390816137415790505b50935060005b815181101561383857604051806040016040528083838151811061379457613794614fc2565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168152602001600c60008585815181106137cd576137cd614fc2565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481525085828151811061382557613825614fc2565b602090810291909101015260010161376e565b505050509193909250565b6040517f454950373132446f6d61696e280000000000000000000000000000000000000060208201527f737472696e67206e616d652c0000000000000000000000000000000000000000602d8201527f737472696e672076657273696f6e2c000000000000000000000000000000000060398201527f75696e7432353620636861696e4964000000000000000000000000000000000060488201527f29000000000000000000000000000000000000000000000000000000000000006057820152600090605801604051602081830303815290604052805190602001208360405160200161393191906151dd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120908301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201526080810183905260a00160405160208183030381529060405280519060200120905092915050565b6000806000808451604003613a0d57505050602082015160408301517f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81169060ff1c601b01613a62565b8451604103613a565750505060208201516040830151606084015160001a601b8114801590613a4057508060ff16601c14155b15613a515760009350505050611c09565b613a62565b60009350505050611c09565b7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115613a965760009350505050611c09565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015613ae9573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00151979650505050505050565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117825560405190917f4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b91a250565b60005473ffffffffffffffffffffffffffffffffffffffff163314611109576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615613c1257613c0d8361465e565b505050565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613c97575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252613c94918101906151f9565b60015b613d23576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201527f6f6e206973206e6f7420555550530000000000000000000000000000000000006064820152608401610c20565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114613dd2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f7860448201527f6961626c655555494400000000000000000000000000000000000000000000006064820152608401610c20565b50613c0d838383614768565b6000806000613e857f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008c8c8c8c8b80519060200120604051602001610b7e9695949392919095865273ffffffffffffffffffffffffffffffffffffffff94851660208701529290931660408501526060840152608083019190915260a082015260c00190565b3360009081526004602052604090205490915060ff16613ed6578060016040517fc802b7f3000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b871580613ee1575085155b15613f1d578060026040517fc802b7f3000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b73ffffffffffffffffffffffffffffffffffffffff8916613f6f578060036040517fc802b7f3000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b73ffffffffffffffffffffffffffffffffffffffff8a1660009081526008602052604090205460ff16613fd3578060046040517fc802b7f3000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b60008181526006602052604090205460ff1615614021578060066040517fc802b7f3000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b85881415801561405157507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8814155b1561408d578060096040517fc802b7f3000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b600061409982866139c2565b905073ffffffffffffffffffffffffffffffffffffffff81166140ed578160076040517fc802b7f3000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600b60209081526040808320938f16835260029093019052205487811015614163578260086040517fc802b7f3000000000000000000000000000000000000000000000000000000008152600401610c20929190614c1c565b878103905080600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555087600d60008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461423b9190615020565b909155505060008381526006602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055815173ffffffffffffffffffffffffffffffffffffffff8e811682529181018d90529182018a9052606082018390528491818f16918516907fc0bd00c18050bbd2c9991c938b45550d2ad9fd31f67ba9d85597d8ef87df1f689060800160405180910390a4909b909a5098505050505050505050565b60006040517fa9059cbb000000000000000000000000000000000000000000000000000000008152836004820152826024820152602060006044836000895af13d15601f3d11600160005114161716915050806143aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f5452414e534645525f4641494c454400000000000000000000000000000000006044820152606401610c20565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff84166143ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600003614439576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831660009081526008602052604090205460ff16614498576040517fd1ef4cea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600b60209081526040808320938716835260029093019052908120546144dc908490615119565b73ffffffffffffffffffffffffffffffffffffffff8087166000908152600b60209081526040808320938916835260029093018152828220849055600d905290812080549293508592909190614533908490615119565b9091555050604080518481526020810183905273ffffffffffffffffffffffffffffffffffffffff80871692908816917fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7910160405180910390a390505b9392505050565b60006040517f23b872dd0000000000000000000000000000000000000000000000000000000081528460048201528360248201528260448201526020600060648360008a5af13d15601f3d1160016000511416171691505080614657576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5452414e534645525f46524f4d5f4641494c45440000000000000000000000006044820152606401610c20565b5050505050565b73ffffffffffffffffffffffffffffffffffffffff81163b614702576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610c20565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6147718361478d565b60008251118061477e5750805b15613c0d576143aa83836147da565b6147968161465e565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060614591838360405180606001604052806027815260200161521360279139606073ffffffffffffffffffffffffffffffffffffffff84163b6148a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e747261637400000000000000000000000000000000000000000000000000006064820152608401610c20565b6000808573ffffffffffffffffffffffffffffffffffffffff16856040516148c891906151dd565b600060405180830381855af49150503d8060008114614903576040519150601f19603f3d011682016040523d82523d6000602084013e614908565b606091505b5091509150614918828286614922565b9695505050505050565b60608315614931575081614591565b8251156149415782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2091906149eb565b60005b83811015614990578181015183820152602001614978565b838111156143aa5750506000910152565b600081518084526149b9816020860160208601614975565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061459160208301846149a1565b803573ffffffffffffffffffffffffffffffffffffffff81168114614a2257600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112614a6757600080fd5b813567ffffffffffffffff80821115614a8257614a82614a27565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715614ac857614ac8614a27565b81604052838152866020858801011115614ae157600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215614b1457600080fd5b614b1d836149fe565b9150602083013567ffffffffffffffff811115614b3957600080fd5b614b4585828601614a56565b9150509250929050565b600060208284031215614b6157600080fd5b614591826149fe565b6020808252825182820181905260009190848201906040850190845b81811015614ba257835183529284019291840191600101614b86565b50909695505050505050565b600080600060608486031215614bc357600080fd5b614bcc846149fe565b95602085013595506040909401359392505050565b600c8110614c18577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b828152604081016145916020830184614be1565b600060208284031215614c4257600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015614ba257835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101614c65565b600080600060608486031215614cac57600080fd5b614cb5846149fe565b9250614cc3602085016149fe565b9150604084013590509250925092565b60008060008060808587031215614ce957600080fd5b614cf2856149fe565b935060208501359250604085013567ffffffffffffffff80821115614d1657600080fd5b614d2288838901614a56565b93506060870135915080821115614d3857600080fd5b50614d4587828801614a56565b91505092959194509250565b60008060408385031215614d6457600080fd5b614d6d836149fe565b946020939093013593505050565b60008060408385031215614d8e57600080fd5b614d97836149fe565b9150614da5602084016149fe565b90509250929050565b60008060408385031215614dc157600080fd5b614dca836149fe565b915060208301358015158114614ddf57600080fd5b809150509250929050565b600080600080600080600060e0888a031215614e0557600080fd5b614e0e886149fe565b9650614e1c602089016149fe565b955060408801359450606088013593506080880135925060a088013567ffffffffffffffff80821115614e4e57600080fd5b614e5a8b838c01614a56565b935060c08a0135915080821115614e7057600080fd5b50614e7d8a828b01614a56565b91505092959891949750929550565b60208101611c098284614be1565b600081518084526020808501945080840160005b83811015614ef457614ee1878351805173ffffffffffffffffffffffffffffffffffffffff168252602090810151910152565b6040969096019590820190600101614eae565b509495945050505050565b606080825284519082018190526000906020906080840190828801845b82811015614f6257614f4f848351805173ffffffffffffffffffffffffffffffffffffffff168252602090810151910152565b6040939093019290840190600101614f1c565b5050508381038285015285518082528683019183019060005b81811015614fa2578351805184528501518584015292840192604090920191600101614f7b565b50508481036040860152614fb68187614e9a565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561503257615032614ff1565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000808312837f8000000000000000000000000000000000000000000000000000000000000000018312811516156150a0576150a0614ff1565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0183138116156150d4576150d4614ff1565b50500390565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261491860808301846149a1565b6000821982111561512c5761512c614ff1565b500190565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0384138115161561516b5761516b614ff1565b827f800000000000000000000000000000000000000000000000000000000000000003841281161561519f5761519f614ff1565b50500190565b60007f800000000000000000000000000000000000000000000000000000000000000082036151d6576151d6614ff1565b5060000390565b600082516151ef818460208701614975565b9190910192915050565b60006020828403121561520b57600080fd5b505191905056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300080f000a",
}

// AccountsABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountsMetaData.ABI instead.
var AccountsABI = AccountsMetaData.ABI

// AccountsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountsMetaData.Bin instead.
var AccountsBin = AccountsMetaData.Bin

// DeployAccounts deploys a new Ethereum contract, binding an instance of Accounts to it.
func DeployAccounts(auth *bind.TransactOpts, backend bind.ContractBackend, _domain string, _chainId *big.Int, _quote common.Address) (common.Address, *types.Transaction, *Accounts, error) {
	parsed, err := AccountsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountsBin), backend, _domain, _chainId, _quote)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Accounts{AccountsCaller: AccountsCaller{contract: contract}, AccountsTransactor: AccountsTransactor{contract: contract}, AccountsFilterer: AccountsFilterer{contract: contract}}, nil
}

// Accounts is an auto generated Go binding around an Ethereum contract.
type Accounts struct {
	AccountsCaller     // Read-only binding to the contract
	AccountsTransactor // Write-only binding to the contract
	AccountsFilterer   // Log filterer for contract events
}

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountsSession struct {
	Contract     *Accounts         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountsCallerSession struct {
	Contract *AccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountsTransactorSession struct {
	Contract     *AccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountsRaw struct {
	Contract *Accounts // Generic contract binding to access the raw methods on
}

// AccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountsCallerRaw struct {
	Contract *AccountsCaller // Generic read-only contract binding to access the raw methods on
}

// AccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountsTransactorRaw struct {
	Contract *AccountsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccounts creates a new instance of Accounts, bound to a specific deployed contract.
func NewAccounts(address common.Address, backend bind.ContractBackend) (*Accounts, error) {
	contract, err := bindAccounts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accounts{AccountsCaller: AccountsCaller{contract: contract}, AccountsTransactor: AccountsTransactor{contract: contract}, AccountsFilterer: AccountsFilterer{contract: contract}}, nil
}

// NewAccountsCaller creates a new read-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsCaller(address common.Address, caller bind.ContractCaller) (*AccountsCaller, error) {
	contract, err := bindAccounts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsCaller{contract: contract}, nil
}

// NewAccountsTransactor creates a new write-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountsTransactor, error) {
	contract, err := bindAccounts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsTransactor{contract: contract}, nil
}

// NewAccountsFilterer creates a new log filterer instance of Accounts, bound to a specific deployed contract.
func NewAccountsFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountsFilterer, error) {
	contract, err := bindAccounts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountsFilterer{contract: contract}, nil
}

// bindAccounts binds a generic wrapper to an already deployed contract.
func bindAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.AccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transact(opts, method, params...)
}

// Authorities is a free data retrieval call binding the contract method 0x91223d69.
//
// Solidity: function authorities(address ) view returns(bool)
func (_Accounts *AccountsCaller) Authorities(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "authorities", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Authorities is a free data retrieval call binding the contract method 0x91223d69.
//
// Solidity: function authorities(address ) view returns(bool)
func (_Accounts *AccountsSession) Authorities(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.Authorities(&_Accounts.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x91223d69.
//
// Solidity: function authorities(address ) view returns(bool)
func (_Accounts *AccountsCallerSession) Authorities(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.Authorities(&_Accounts.CallOpts, arg0)
}

// Collaterals is a free data retrieval call binding the contract method 0xeeb97d3b.
//
// Solidity: function collaterals(address ) view returns(bool)
func (_Accounts *AccountsCaller) Collaterals(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "collaterals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Collaterals is a free data retrieval call binding the contract method 0xeeb97d3b.
//
// Solidity: function collaterals(address ) view returns(bool)
func (_Accounts *AccountsSession) Collaterals(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.Collaterals(&_Accounts.CallOpts, arg0)
}

// Collaterals is a free data retrieval call binding the contract method 0xeeb97d3b.
//
// Solidity: function collaterals(address ) view returns(bool)
func (_Accounts *AccountsCallerSession) Collaterals(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.Collaterals(&_Accounts.CallOpts, arg0)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Accounts *AccountsCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Accounts *AccountsSession) DomainSeparator() ([32]byte, error) {
	return _Accounts.Contract.DomainSeparator(&_Accounts.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Accounts *AccountsCallerSession) DomainSeparator() ([32]byte, error) {
	return _Accounts.Contract.DomainSeparator(&_Accounts.CallOpts)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address account) view returns((address,uint256)[] balances, (uint256,int256)[] positions, (address,uint256)[] signingKeys)
func (_Accounts *AccountsCaller) GetAccount(opts *bind.CallOpts, account common.Address) (struct {
	Balances    []AccountsInterfaceBalance
	Positions   []AccountsInterfacePosition
	SigningKeys []AccountsInterfaceSigningKey
}, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getAccount", account)

	outstruct := new(struct {
		Balances    []AccountsInterfaceBalance
		Positions   []AccountsInterfacePosition
		SigningKeys []AccountsInterfaceSigningKey
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balances = *abi.ConvertType(out[0], new([]AccountsInterfaceBalance)).(*[]AccountsInterfaceBalance)
	outstruct.Positions = *abi.ConvertType(out[1], new([]AccountsInterfacePosition)).(*[]AccountsInterfacePosition)
	outstruct.SigningKeys = *abi.ConvertType(out[2], new([]AccountsInterfaceSigningKey)).(*[]AccountsInterfaceSigningKey)

	return *outstruct, err

}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address account) view returns((address,uint256)[] balances, (uint256,int256)[] positions, (address,uint256)[] signingKeys)
func (_Accounts *AccountsSession) GetAccount(account common.Address) (struct {
	Balances    []AccountsInterfaceBalance
	Positions   []AccountsInterfacePosition
	SigningKeys []AccountsInterfaceSigningKey
}, error) {
	return _Accounts.Contract.GetAccount(&_Accounts.CallOpts, account)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address account) view returns((address,uint256)[] balances, (uint256,int256)[] positions, (address,uint256)[] signingKeys)
func (_Accounts *AccountsCallerSession) GetAccount(account common.Address) (struct {
	Balances    []AccountsInterfaceBalance
	Positions   []AccountsInterfacePosition
	SigningKeys []AccountsInterfaceSigningKey
}, error) {
	return _Accounts.Contract.GetAccount(&_Accounts.CallOpts, account)
}

// GetCollateralBalance is a free data retrieval call binding the contract method 0xb7d5820b.
//
// Solidity: function getCollateralBalance(address account, address collateral) view returns(uint256)
func (_Accounts *AccountsCaller) GetCollateralBalance(opts *bind.CallOpts, account common.Address, collateral common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getCollateralBalance", account, collateral)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralBalance is a free data retrieval call binding the contract method 0xb7d5820b.
//
// Solidity: function getCollateralBalance(address account, address collateral) view returns(uint256)
func (_Accounts *AccountsSession) GetCollateralBalance(account common.Address, collateral common.Address) (*big.Int, error) {
	return _Accounts.Contract.GetCollateralBalance(&_Accounts.CallOpts, account, collateral)
}

// GetCollateralBalance is a free data retrieval call binding the contract method 0xb7d5820b.
//
// Solidity: function getCollateralBalance(address account, address collateral) view returns(uint256)
func (_Accounts *AccountsCallerSession) GetCollateralBalance(account common.Address, collateral common.Address) (*big.Int, error) {
	return _Accounts.Contract.GetCollateralBalance(&_Accounts.CallOpts, account, collateral)
}

// GetCollaterals is a free data retrieval call binding the contract method 0x78b92636.
//
// Solidity: function getCollaterals() view returns(address[])
func (_Accounts *AccountsCaller) GetCollaterals(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getCollaterals")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCollaterals is a free data retrieval call binding the contract method 0x78b92636.
//
// Solidity: function getCollaterals() view returns(address[])
func (_Accounts *AccountsSession) GetCollaterals() ([]common.Address, error) {
	return _Accounts.Contract.GetCollaterals(&_Accounts.CallOpts)
}

// GetCollaterals is a free data retrieval call binding the contract method 0x78b92636.
//
// Solidity: function getCollaterals() view returns(address[])
func (_Accounts *AccountsCallerSession) GetCollaterals() ([]common.Address, error) {
	return _Accounts.Contract.GetCollaterals(&_Accounts.CallOpts)
}

// GetFeeAccount is a free data retrieval call binding the contract method 0x4614be9f.
//
// Solidity: function getFeeAccount() pure returns(address)
func (_Accounts *AccountsCaller) GetFeeAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getFeeAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeAccount is a free data retrieval call binding the contract method 0x4614be9f.
//
// Solidity: function getFeeAccount() pure returns(address)
func (_Accounts *AccountsSession) GetFeeAccount() (common.Address, error) {
	return _Accounts.Contract.GetFeeAccount(&_Accounts.CallOpts)
}

// GetFeeAccount is a free data retrieval call binding the contract method 0x4614be9f.
//
// Solidity: function getFeeAccount() pure returns(address)
func (_Accounts *AccountsCallerSession) GetFeeAccount() (common.Address, error) {
	return _Accounts.Contract.GetFeeAccount(&_Accounts.CallOpts)
}

// GetInstrumentPosition is a free data retrieval call binding the contract method 0x49ad755b.
//
// Solidity: function getInstrumentPosition(address account, uint256 instrument) view returns(int256)
func (_Accounts *AccountsCaller) GetInstrumentPosition(opts *bind.CallOpts, account common.Address, instrument *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getInstrumentPosition", account, instrument)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInstrumentPosition is a free data retrieval call binding the contract method 0x49ad755b.
//
// Solidity: function getInstrumentPosition(address account, uint256 instrument) view returns(int256)
func (_Accounts *AccountsSession) GetInstrumentPosition(account common.Address, instrument *big.Int) (*big.Int, error) {
	return _Accounts.Contract.GetInstrumentPosition(&_Accounts.CallOpts, account, instrument)
}

// GetInstrumentPosition is a free data retrieval call binding the contract method 0x49ad755b.
//
// Solidity: function getInstrumentPosition(address account, uint256 instrument) view returns(int256)
func (_Accounts *AccountsCallerSession) GetInstrumentPosition(account common.Address, instrument *big.Int) (*big.Int, error) {
	return _Accounts.Contract.GetInstrumentPosition(&_Accounts.CallOpts, account, instrument)
}

// GetInstruments is a free data retrieval call binding the contract method 0x1635717c.
//
// Solidity: function getInstruments() view returns(uint256[])
func (_Accounts *AccountsCaller) GetInstruments(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getInstruments")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetInstruments is a free data retrieval call binding the contract method 0x1635717c.
//
// Solidity: function getInstruments() view returns(uint256[])
func (_Accounts *AccountsSession) GetInstruments() ([]*big.Int, error) {
	return _Accounts.Contract.GetInstruments(&_Accounts.CallOpts)
}

// GetInstruments is a free data retrieval call binding the contract method 0x1635717c.
//
// Solidity: function getInstruments() view returns(uint256[])
func (_Accounts *AccountsCallerSession) GetInstruments() ([]*big.Int, error) {
	return _Accounts.Contract.GetInstruments(&_Accounts.CallOpts)
}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() pure returns(address)
func (_Accounts *AccountsCaller) GetInsuranceFund(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getInsuranceFund")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() pure returns(address)
func (_Accounts *AccountsSession) GetInsuranceFund() (common.Address, error) {
	return _Accounts.Contract.GetInsuranceFund(&_Accounts.CallOpts)
}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() pure returns(address)
func (_Accounts *AccountsCallerSession) GetInsuranceFund() (common.Address, error) {
	return _Accounts.Contract.GetInsuranceFund(&_Accounts.CallOpts)
}

// GetQuoteBalance is a free data retrieval call binding the contract method 0xaf1aa3b1.
//
// Solidity: function getQuoteBalance(address account) view returns(uint256)
func (_Accounts *AccountsCaller) GetQuoteBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getQuoteBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteBalance is a free data retrieval call binding the contract method 0xaf1aa3b1.
//
// Solidity: function getQuoteBalance(address account) view returns(uint256)
func (_Accounts *AccountsSession) GetQuoteBalance(account common.Address) (*big.Int, error) {
	return _Accounts.Contract.GetQuoteBalance(&_Accounts.CallOpts, account)
}

// GetQuoteBalance is a free data retrieval call binding the contract method 0xaf1aa3b1.
//
// Solidity: function getQuoteBalance(address account) view returns(uint256)
func (_Accounts *AccountsCallerSession) GetQuoteBalance(account common.Address) (*big.Int, error) {
	return _Accounts.Contract.GetQuoteBalance(&_Accounts.CallOpts, account)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x241d400d.
//
// Solidity: function getSigningKeys(address account) view returns(address[])
func (_Accounts *AccountsCaller) GetSigningKeys(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getSigningKeys", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigningKeys is a free data retrieval call binding the contract method 0x241d400d.
//
// Solidity: function getSigningKeys(address account) view returns(address[])
func (_Accounts *AccountsSession) GetSigningKeys(account common.Address) ([]common.Address, error) {
	return _Accounts.Contract.GetSigningKeys(&_Accounts.CallOpts, account)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x241d400d.
//
// Solidity: function getSigningKeys(address account) view returns(address[])
func (_Accounts *AccountsCallerSession) GetSigningKeys(account common.Address) ([]common.Address, error) {
	return _Accounts.Contract.GetSigningKeys(&_Accounts.CallOpts, account)
}

// HasSigningKey is a free data retrieval call binding the contract method 0x62946d3b.
//
// Solidity: function hasSigningKey(address account, address signingKey) view returns(bool)
func (_Accounts *AccountsCaller) HasSigningKey(opts *bind.CallOpts, account common.Address, signingKey common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hasSigningKey", account, signingKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSigningKey is a free data retrieval call binding the contract method 0x62946d3b.
//
// Solidity: function hasSigningKey(address account, address signingKey) view returns(bool)
func (_Accounts *AccountsSession) HasSigningKey(account common.Address, signingKey common.Address) (bool, error) {
	return _Accounts.Contract.HasSigningKey(&_Accounts.CallOpts, account, signingKey)
}

// HasSigningKey is a free data retrieval call binding the contract method 0x62946d3b.
//
// Solidity: function hasSigningKey(address account, address signingKey) view returns(bool)
func (_Accounts *AccountsCallerSession) HasSigningKey(account common.Address, signingKey common.Address) (bool, error) {
	return _Accounts.Contract.HasSigningKey(&_Accounts.CallOpts, account, signingKey)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_Accounts *AccountsCaller) Hashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_Accounts *AccountsSession) Hashes(arg0 [32]byte) (bool, error) {
	return _Accounts.Contract.Hashes(&_Accounts.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_Accounts *AccountsCallerSession) Hashes(arg0 [32]byte) (bool, error) {
	return _Accounts.Contract.Hashes(&_Accounts.CallOpts, arg0)
}

// Instruments is a free data retrieval call binding the contract method 0x22bbad0b.
//
// Solidity: function instruments(uint256 ) view returns(bool)
func (_Accounts *AccountsCaller) Instruments(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "instruments", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Instruments is a free data retrieval call binding the contract method 0x22bbad0b.
//
// Solidity: function instruments(uint256 ) view returns(bool)
func (_Accounts *AccountsSession) Instruments(arg0 *big.Int) (bool, error) {
	return _Accounts.Contract.Instruments(&_Accounts.CallOpts, arg0)
}

// Instruments is a free data retrieval call binding the contract method 0x22bbad0b.
//
// Solidity: function instruments(uint256 ) view returns(bool)
func (_Accounts *AccountsCallerSession) Instruments(arg0 *big.Int) (bool, error) {
	return _Accounts.Contract.Instruments(&_Accounts.CallOpts, arg0)
}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Accounts *AccountsCaller) Keepers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "keepers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Accounts *AccountsSession) Keepers(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.Keepers(&_Accounts.CallOpts, arg0)
}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Accounts *AccountsCallerSession) Keepers(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.Keepers(&_Accounts.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Accounts *AccountsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Accounts *AccountsSession) Name() (string, error) {
	return _Accounts.Contract.Name(&_Accounts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Accounts *AccountsCallerSession) Name() (string, error) {
	return _Accounts.Contract.Name(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accounts *AccountsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accounts *AccountsSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accounts *AccountsCallerSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Accounts *AccountsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Accounts *AccountsSession) ProxiableUUID() ([32]byte, error) {
	return _Accounts.Contract.ProxiableUUID(&_Accounts.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Accounts *AccountsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Accounts.Contract.ProxiableUUID(&_Accounts.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0x999b93af.
//
// Solidity: function quote() view returns(address)
func (_Accounts *AccountsCaller) Quote(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "quote")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0x999b93af.
//
// Solidity: function quote() view returns(address)
func (_Accounts *AccountsSession) Quote() (common.Address, error) {
	return _Accounts.Contract.Quote(&_Accounts.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0x999b93af.
//
// Solidity: function quote() view returns(address)
func (_Accounts *AccountsCallerSession) Quote() (common.Address, error) {
	return _Accounts.Contract.Quote(&_Accounts.CallOpts)
}

// QuoteDecimals is a free data retrieval call binding the contract method 0x3fd1e2bd.
//
// Solidity: function quoteDecimals() view returns(uint8)
func (_Accounts *AccountsCaller) QuoteDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "quoteDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuoteDecimals is a free data retrieval call binding the contract method 0x3fd1e2bd.
//
// Solidity: function quoteDecimals() view returns(uint8)
func (_Accounts *AccountsSession) QuoteDecimals() (uint8, error) {
	return _Accounts.Contract.QuoteDecimals(&_Accounts.CallOpts)
}

// QuoteDecimals is a free data retrieval call binding the contract method 0x3fd1e2bd.
//
// Solidity: function quoteDecimals() view returns(uint8)
func (_Accounts *AccountsCallerSession) QuoteDecimals() (uint8, error) {
	return _Accounts.Contract.QuoteDecimals(&_Accounts.CallOpts)
}

// SigningKeyExpiry is a free data retrieval call binding the contract method 0x6205ed5e.
//
// Solidity: function signingKeyExpiry(address ) view returns(uint256)
func (_Accounts *AccountsCaller) SigningKeyExpiry(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "signingKeyExpiry", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SigningKeyExpiry is a free data retrieval call binding the contract method 0x6205ed5e.
//
// Solidity: function signingKeyExpiry(address ) view returns(uint256)
func (_Accounts *AccountsSession) SigningKeyExpiry(arg0 common.Address) (*big.Int, error) {
	return _Accounts.Contract.SigningKeyExpiry(&_Accounts.CallOpts, arg0)
}

// SigningKeyExpiry is a free data retrieval call binding the contract method 0x6205ed5e.
//
// Solidity: function signingKeyExpiry(address ) view returns(uint256)
func (_Accounts *AccountsCallerSession) SigningKeyExpiry(arg0 common.Address) (*big.Int, error) {
	return _Accounts.Contract.SigningKeyExpiry(&_Accounts.CallOpts, arg0)
}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address ) view returns(uint256)
func (_Accounts *AccountsCaller) TotalBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "totalBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address ) view returns(uint256)
func (_Accounts *AccountsSession) TotalBalance(arg0 common.Address) (*big.Int, error) {
	return _Accounts.Contract.TotalBalance(&_Accounts.CallOpts, arg0)
}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address ) view returns(uint256)
func (_Accounts *AccountsCallerSession) TotalBalance(arg0 common.Address) (*big.Int, error) {
	return _Accounts.Contract.TotalBalance(&_Accounts.CallOpts, arg0)
}

// TotalPositions is a free data retrieval call binding the contract method 0x61d21920.
//
// Solidity: function totalPositions(uint256 ) view returns(int256)
func (_Accounts *AccountsCaller) TotalPositions(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "totalPositions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPositions is a free data retrieval call binding the contract method 0x61d21920.
//
// Solidity: function totalPositions(uint256 ) view returns(int256)
func (_Accounts *AccountsSession) TotalPositions(arg0 *big.Int) (*big.Int, error) {
	return _Accounts.Contract.TotalPositions(&_Accounts.CallOpts, arg0)
}

// TotalPositions is a free data retrieval call binding the contract method 0x61d21920.
//
// Solidity: function totalPositions(uint256 ) view returns(int256)
func (_Accounts *AccountsCallerSession) TotalPositions(arg0 *big.Int) (*big.Int, error) {
	return _Accounts.Contract.TotalPositions(&_Accounts.CallOpts, arg0)
}

// WithdrawProxies is a free data retrieval call binding the contract method 0xb8eb94e9.
//
// Solidity: function withdrawProxies(address ) view returns(bool)
func (_Accounts *AccountsCaller) WithdrawProxies(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "withdrawProxies", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawProxies is a free data retrieval call binding the contract method 0xb8eb94e9.
//
// Solidity: function withdrawProxies(address ) view returns(bool)
func (_Accounts *AccountsSession) WithdrawProxies(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.WithdrawProxies(&_Accounts.CallOpts, arg0)
}

// WithdrawProxies is a free data retrieval call binding the contract method 0xb8eb94e9.
//
// Solidity: function withdrawProxies(address ) view returns(bool)
func (_Accounts *AccountsCallerSession) WithdrawProxies(arg0 common.Address) (bool, error) {
	return _Accounts.Contract.WithdrawProxies(&_Accounts.CallOpts, arg0)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xf0d2d5a8.
//
// Solidity: function addCollateral(address collateral) returns()
func (_Accounts *AccountsTransactor) AddCollateral(opts *bind.TransactOpts, collateral common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "addCollateral", collateral)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xf0d2d5a8.
//
// Solidity: function addCollateral(address collateral) returns()
func (_Accounts *AccountsSession) AddCollateral(collateral common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.AddCollateral(&_Accounts.TransactOpts, collateral)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xf0d2d5a8.
//
// Solidity: function addCollateral(address collateral) returns()
func (_Accounts *AccountsTransactorSession) AddCollateral(collateral common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.AddCollateral(&_Accounts.TransactOpts, collateral)
}

// AddInstrument is a paid mutator transaction binding the contract method 0x8e991606.
//
// Solidity: function addInstrument(uint256 instrument) returns(uint8)
func (_Accounts *AccountsTransactor) AddInstrument(opts *bind.TransactOpts, instrument *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "addInstrument", instrument)
}

// AddInstrument is a paid mutator transaction binding the contract method 0x8e991606.
//
// Solidity: function addInstrument(uint256 instrument) returns(uint8)
func (_Accounts *AccountsSession) AddInstrument(instrument *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.AddInstrument(&_Accounts.TransactOpts, instrument)
}

// AddInstrument is a paid mutator transaction binding the contract method 0x8e991606.
//
// Solidity: function addInstrument(uint256 instrument) returns(uint8)
func (_Accounts *AccountsTransactorSession) AddInstrument(instrument *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.AddInstrument(&_Accounts.TransactOpts, instrument)
}

// Credit is a paid mutator transaction binding the contract method 0x7de182c5.
//
// Solidity: function credit(address account, address collateral, uint256 amount) returns(uint256, uint8)
func (_Accounts *AccountsTransactor) Credit(opts *bind.TransactOpts, account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "credit", account, collateral, amount)
}

// Credit is a paid mutator transaction binding the contract method 0x7de182c5.
//
// Solidity: function credit(address account, address collateral, uint256 amount) returns(uint256, uint8)
func (_Accounts *AccountsSession) Credit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Credit(&_Accounts.TransactOpts, account, collateral, amount)
}

// Credit is a paid mutator transaction binding the contract method 0x7de182c5.
//
// Solidity: function credit(address account, address collateral, uint256 amount) returns(uint256, uint8)
func (_Accounts *AccountsTransactorSession) Credit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Credit(&_Accounts.TransactOpts, account, collateral, amount)
}

// CreditInstrument is a paid mutator transaction binding the contract method 0x7df5b701.
//
// Solidity: function creditInstrument(address account, uint256 instrument, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsTransactor) CreditInstrument(opts *bind.TransactOpts, account common.Address, instrument *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "creditInstrument", account, instrument, amount)
}

// CreditInstrument is a paid mutator transaction binding the contract method 0x7df5b701.
//
// Solidity: function creditInstrument(address account, uint256 instrument, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsSession) CreditInstrument(account common.Address, instrument *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.CreditInstrument(&_Accounts.TransactOpts, account, instrument, amount)
}

// CreditInstrument is a paid mutator transaction binding the contract method 0x7df5b701.
//
// Solidity: function creditInstrument(address account, uint256 instrument, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsTransactorSession) CreditInstrument(account common.Address, instrument *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.CreditInstrument(&_Accounts.TransactOpts, account, instrument, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x262709e6.
//
// Solidity: function debit(address account, address collateral, uint256 amount) returns(uint256, uint8)
func (_Accounts *AccountsTransactor) Debit(opts *bind.TransactOpts, account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "debit", account, collateral, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x262709e6.
//
// Solidity: function debit(address account, address collateral, uint256 amount) returns(uint256, uint8)
func (_Accounts *AccountsSession) Debit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Debit(&_Accounts.TransactOpts, account, collateral, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x262709e6.
//
// Solidity: function debit(address account, address collateral, uint256 amount) returns(uint256, uint8)
func (_Accounts *AccountsTransactorSession) Debit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Debit(&_Accounts.TransactOpts, account, collateral, amount)
}

// DebitInstrument is a paid mutator transaction binding the contract method 0x1dcc7d5c.
//
// Solidity: function debitInstrument(address account, uint256 instrument, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsTransactor) DebitInstrument(opts *bind.TransactOpts, account common.Address, instrument *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "debitInstrument", account, instrument, amount)
}

// DebitInstrument is a paid mutator transaction binding the contract method 0x1dcc7d5c.
//
// Solidity: function debitInstrument(address account, uint256 instrument, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsSession) DebitInstrument(account common.Address, instrument *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.DebitInstrument(&_Accounts.TransactOpts, account, instrument, amount)
}

// DebitInstrument is a paid mutator transaction binding the contract method 0x1dcc7d5c.
//
// Solidity: function debitInstrument(address account, uint256 instrument, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsTransactorSession) DebitInstrument(account common.Address, instrument *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.DebitInstrument(&_Accounts.TransactOpts, account, instrument, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address account, address collateral, uint256 amount) returns(uint256)
func (_Accounts *AccountsTransactor) Deposit(opts *bind.TransactOpts, account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "deposit", account, collateral, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address account, address collateral, uint256 amount) returns(uint256)
func (_Accounts *AccountsSession) Deposit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Deposit(&_Accounts.TransactOpts, account, collateral, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address account, address collateral, uint256 amount) returns(uint256)
func (_Accounts *AccountsTransactorSession) Deposit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Deposit(&_Accounts.TransactOpts, account, collateral, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Accounts *AccountsTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Accounts *AccountsSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.Initialize(&_Accounts.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Accounts *AccountsTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.Initialize(&_Accounts.TransactOpts, _owner)
}

// RegisterSigningKey is a paid mutator transaction binding the contract method 0x48e4ccbe.
//
// Solidity: function registerSigningKey(address signingKey, uint256 expiry, bytes signingKeySig, bytes accountSig) returns()
func (_Accounts *AccountsTransactor) RegisterSigningKey(opts *bind.TransactOpts, signingKey common.Address, expiry *big.Int, signingKeySig []byte, accountSig []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "registerSigningKey", signingKey, expiry, signingKeySig, accountSig)
}

// RegisterSigningKey is a paid mutator transaction binding the contract method 0x48e4ccbe.
//
// Solidity: function registerSigningKey(address signingKey, uint256 expiry, bytes signingKeySig, bytes accountSig) returns()
func (_Accounts *AccountsSession) RegisterSigningKey(signingKey common.Address, expiry *big.Int, signingKeySig []byte, accountSig []byte) (*types.Transaction, error) {
	return _Accounts.Contract.RegisterSigningKey(&_Accounts.TransactOpts, signingKey, expiry, signingKeySig, accountSig)
}

// RegisterSigningKey is a paid mutator transaction binding the contract method 0x48e4ccbe.
//
// Solidity: function registerSigningKey(address signingKey, uint256 expiry, bytes signingKeySig, bytes accountSig) returns()
func (_Accounts *AccountsTransactorSession) RegisterSigningKey(signingKey common.Address, expiry *big.Int, signingKeySig []byte, accountSig []byte) (*types.Transaction, error) {
	return _Accounts.Contract.RegisterSigningKey(&_Accounts.TransactOpts, signingKey, expiry, signingKeySig, accountSig)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0xc99d3a06.
//
// Solidity: function removeCollateral(address collateral) returns()
func (_Accounts *AccountsTransactor) RemoveCollateral(opts *bind.TransactOpts, collateral common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeCollateral", collateral)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0xc99d3a06.
//
// Solidity: function removeCollateral(address collateral) returns()
func (_Accounts *AccountsSession) RemoveCollateral(collateral common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveCollateral(&_Accounts.TransactOpts, collateral)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0xc99d3a06.
//
// Solidity: function removeCollateral(address collateral) returns()
func (_Accounts *AccountsTransactorSession) RemoveCollateral(collateral common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveCollateral(&_Accounts.TransactOpts, collateral)
}

// RemoveInstrument is a paid mutator transaction binding the contract method 0xe30e8384.
//
// Solidity: function removeInstrument(uint256 instrument) returns(uint8)
func (_Accounts *AccountsTransactor) RemoveInstrument(opts *bind.TransactOpts, instrument *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeInstrument", instrument)
}

// RemoveInstrument is a paid mutator transaction binding the contract method 0xe30e8384.
//
// Solidity: function removeInstrument(uint256 instrument) returns(uint8)
func (_Accounts *AccountsSession) RemoveInstrument(instrument *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveInstrument(&_Accounts.TransactOpts, instrument)
}

// RemoveInstrument is a paid mutator transaction binding the contract method 0xe30e8384.
//
// Solidity: function removeInstrument(uint256 instrument) returns(uint8)
func (_Accounts *AccountsTransactorSession) RemoveInstrument(instrument *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.RemoveInstrument(&_Accounts.TransactOpts, instrument)
}

// ResetInstrumentPosition is a paid mutator transaction binding the contract method 0x9aeddeff.
//
// Solidity: function resetInstrumentPosition(address account, uint256 instrument) returns(uint8)
func (_Accounts *AccountsTransactor) ResetInstrumentPosition(opts *bind.TransactOpts, account common.Address, instrument *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "resetInstrumentPosition", account, instrument)
}

// ResetInstrumentPosition is a paid mutator transaction binding the contract method 0x9aeddeff.
//
// Solidity: function resetInstrumentPosition(address account, uint256 instrument) returns(uint8)
func (_Accounts *AccountsSession) ResetInstrumentPosition(account common.Address, instrument *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.ResetInstrumentPosition(&_Accounts.TransactOpts, account, instrument)
}

// ResetInstrumentPosition is a paid mutator transaction binding the contract method 0x9aeddeff.
//
// Solidity: function resetInstrumentPosition(address account, uint256 instrument) returns(uint8)
func (_Accounts *AccountsTransactorSession) ResetInstrumentPosition(account common.Address, instrument *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.ResetInstrumentPosition(&_Accounts.TransactOpts, account, instrument)
}

// RevokeSigningKey is a paid mutator transaction binding the contract method 0x07b1b8b2.
//
// Solidity: function revokeSigningKey(address signingKey, bytes signature) returns()
func (_Accounts *AccountsTransactor) RevokeSigningKey(opts *bind.TransactOpts, signingKey common.Address, signature []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "revokeSigningKey", signingKey, signature)
}

// RevokeSigningKey is a paid mutator transaction binding the contract method 0x07b1b8b2.
//
// Solidity: function revokeSigningKey(address signingKey, bytes signature) returns()
func (_Accounts *AccountsSession) RevokeSigningKey(signingKey common.Address, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.RevokeSigningKey(&_Accounts.TransactOpts, signingKey, signature)
}

// RevokeSigningKey is a paid mutator transaction binding the contract method 0x07b1b8b2.
//
// Solidity: function revokeSigningKey(address signingKey, bytes signature) returns()
func (_Accounts *AccountsTransactorSession) RevokeSigningKey(signingKey common.Address, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.RevokeSigningKey(&_Accounts.TransactOpts, signingKey, signature)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Accounts *AccountsTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Accounts *AccountsSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetOwner(&_Accounts.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Accounts *AccountsTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetOwner(&_Accounts.TransactOpts, newOwner)
}

// UpdateAuthority is a paid mutator transaction binding the contract method 0x6cd22eaf.
//
// Solidity: function updateAuthority(address authority, bool allowed) returns()
func (_Accounts *AccountsTransactor) UpdateAuthority(opts *bind.TransactOpts, authority common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "updateAuthority", authority, allowed)
}

// UpdateAuthority is a paid mutator transaction binding the contract method 0x6cd22eaf.
//
// Solidity: function updateAuthority(address authority, bool allowed) returns()
func (_Accounts *AccountsSession) UpdateAuthority(authority common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateAuthority(&_Accounts.TransactOpts, authority, allowed)
}

// UpdateAuthority is a paid mutator transaction binding the contract method 0x6cd22eaf.
//
// Solidity: function updateAuthority(address authority, bool allowed) returns()
func (_Accounts *AccountsTransactorSession) UpdateAuthority(authority common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateAuthority(&_Accounts.TransactOpts, authority, allowed)
}

// UpdateKeeper is a paid mutator transaction binding the contract method 0xd3057877.
//
// Solidity: function updateKeeper(address keeper, bool allowed) returns()
func (_Accounts *AccountsTransactor) UpdateKeeper(opts *bind.TransactOpts, keeper common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "updateKeeper", keeper, allowed)
}

// UpdateKeeper is a paid mutator transaction binding the contract method 0xd3057877.
//
// Solidity: function updateKeeper(address keeper, bool allowed) returns()
func (_Accounts *AccountsSession) UpdateKeeper(keeper common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateKeeper(&_Accounts.TransactOpts, keeper, allowed)
}

// UpdateKeeper is a paid mutator transaction binding the contract method 0xd3057877.
//
// Solidity: function updateKeeper(address keeper, bool allowed) returns()
func (_Accounts *AccountsTransactorSession) UpdateKeeper(keeper common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateKeeper(&_Accounts.TransactOpts, keeper, allowed)
}

// UpdateWithdrawProxy is a paid mutator transaction binding the contract method 0xd5c20730.
//
// Solidity: function updateWithdrawProxy(address withdrawProxy, bool allowed) returns()
func (_Accounts *AccountsTransactor) UpdateWithdrawProxy(opts *bind.TransactOpts, withdrawProxy common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "updateWithdrawProxy", withdrawProxy, allowed)
}

// UpdateWithdrawProxy is a paid mutator transaction binding the contract method 0xd5c20730.
//
// Solidity: function updateWithdrawProxy(address withdrawProxy, bool allowed) returns()
func (_Accounts *AccountsSession) UpdateWithdrawProxy(withdrawProxy common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateWithdrawProxy(&_Accounts.TransactOpts, withdrawProxy, allowed)
}

// UpdateWithdrawProxy is a paid mutator transaction binding the contract method 0xd5c20730.
//
// Solidity: function updateWithdrawProxy(address withdrawProxy, bool allowed) returns()
func (_Accounts *AccountsTransactorSession) UpdateWithdrawProxy(withdrawProxy common.Address, allowed bool) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateWithdrawProxy(&_Accounts.TransactOpts, withdrawProxy, allowed)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Accounts *AccountsTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Accounts *AccountsSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.UpgradeTo(&_Accounts.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Accounts *AccountsTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.UpgradeTo(&_Accounts.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Accounts *AccountsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Accounts *AccountsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Accounts.Contract.UpgradeToAndCall(&_Accounts.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Accounts *AccountsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Accounts.Contract.UpgradeToAndCall(&_Accounts.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x72d39ed6.
//
// Solidity: function withdraw(address collateral, address to, uint256 amount, uint256 salt, uint256 withdrawAmount, bytes data, bytes signature) returns(uint256)
func (_Accounts *AccountsTransactor) Withdraw(opts *bind.TransactOpts, collateral common.Address, to common.Address, amount *big.Int, salt *big.Int, withdrawAmount *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "withdraw", collateral, to, amount, salt, withdrawAmount, data, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x72d39ed6.
//
// Solidity: function withdraw(address collateral, address to, uint256 amount, uint256 salt, uint256 withdrawAmount, bytes data, bytes signature) returns(uint256)
func (_Accounts *AccountsSession) Withdraw(collateral common.Address, to common.Address, amount *big.Int, salt *big.Int, withdrawAmount *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Withdraw(&_Accounts.TransactOpts, collateral, to, amount, salt, withdrawAmount, data, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x72d39ed6.
//
// Solidity: function withdraw(address collateral, address to, uint256 amount, uint256 salt, uint256 withdrawAmount, bytes data, bytes signature) returns(uint256)
func (_Accounts *AccountsTransactorSession) Withdraw(collateral common.Address, to common.Address, amount *big.Int, salt *big.Int, withdrawAmount *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Withdraw(&_Accounts.TransactOpts, collateral, to, amount, salt, withdrawAmount, data, signature)
}

// AccountsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Accounts contract.
type AccountsAdminChangedIterator struct {
	Event *AccountsAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccountsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAdminChanged)
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
		it.Event = new(AccountsAdminChanged)
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
func (it *AccountsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAdminChanged represents a AdminChanged event raised by the Accounts contract.
type AccountsAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Accounts *AccountsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AccountsAdminChangedIterator, error) {

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AccountsAdminChangedIterator{contract: _Accounts.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Accounts *AccountsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AccountsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAdminChanged)
				if err := _Accounts.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseAdminChanged(log types.Log) (*AccountsAdminChanged, error) {
	event := new(AccountsAdminChanged)
	if err := _Accounts.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the Accounts contract.
type AccountsAuthorityUpdatedIterator struct {
	Event *AccountsAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *AccountsAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAuthorityUpdated)
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
		it.Event = new(AccountsAuthorityUpdated)
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
func (it *AccountsAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAuthorityUpdated represents a AuthorityUpdated event raised by the Accounts contract.
type AccountsAuthorityUpdated struct {
	Authority common.Address
	Allowed   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0xc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc.
//
// Solidity: event AuthorityUpdated(address indexed authority, bool indexed allowed)
func (_Accounts *AccountsFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts, authority []common.Address, allowed []bool) (*AccountsAuthorityUpdatedIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AuthorityUpdated", authorityRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAuthorityUpdatedIterator{contract: _Accounts.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0xc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc.
//
// Solidity: event AuthorityUpdated(address indexed authority, bool indexed allowed)
func (_Accounts *AccountsFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *AccountsAuthorityUpdated, authority []common.Address, allowed []bool) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AuthorityUpdated", authorityRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAuthorityUpdated)
				if err := _Accounts.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseAuthorityUpdated(log types.Log) (*AccountsAuthorityUpdated, error) {
	event := new(AccountsAuthorityUpdated)
	if err := _Accounts.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Accounts contract.
type AccountsBeaconUpgradedIterator struct {
	Event *AccountsBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AccountsBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsBeaconUpgraded)
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
		it.Event = new(AccountsBeaconUpgraded)
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
func (it *AccountsBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsBeaconUpgraded represents a BeaconUpgraded event raised by the Accounts contract.
type AccountsBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Accounts *AccountsFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AccountsBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AccountsBeaconUpgradedIterator{contract: _Accounts.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Accounts *AccountsFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AccountsBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsBeaconUpgraded)
				if err := _Accounts.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseBeaconUpgraded(log types.Log) (*AccountsBeaconUpgraded, error) {
	event := new(AccountsBeaconUpgraded)
	if err := _Accounts.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the Accounts contract.
type AccountsCollateralAddedIterator struct {
	Event *AccountsCollateralAdded // Event containing the contract specifics and raw log

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
func (it *AccountsCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsCollateralAdded)
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
		it.Event = new(AccountsCollateralAdded)
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
func (it *AccountsCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsCollateralAdded represents a CollateralAdded event raised by the Accounts contract.
type AccountsCollateralAdded struct {
	Collateral common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f.
//
// Solidity: event CollateralAdded(address indexed collateral)
func (_Accounts *AccountsFilterer) FilterCollateralAdded(opts *bind.FilterOpts, collateral []common.Address) (*AccountsCollateralAddedIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "CollateralAdded", collateralRule)
	if err != nil {
		return nil, err
	}
	return &AccountsCollateralAddedIterator{contract: _Accounts.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f.
//
// Solidity: event CollateralAdded(address indexed collateral)
func (_Accounts *AccountsFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *AccountsCollateralAdded, collateral []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "CollateralAdded", collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsCollateralAdded)
				if err := _Accounts.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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

// ParseCollateralAdded is a log parse operation binding the contract event 0x7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f.
//
// Solidity: event CollateralAdded(address indexed collateral)
func (_Accounts *AccountsFilterer) ParseCollateralAdded(log types.Log) (*AccountsCollateralAdded, error) {
	event := new(AccountsCollateralAdded)
	if err := _Accounts.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsCollateralRemovedIterator is returned from FilterCollateralRemoved and is used to iterate over the raw logs and unpacked data for CollateralRemoved events raised by the Accounts contract.
type AccountsCollateralRemovedIterator struct {
	Event *AccountsCollateralRemoved // Event containing the contract specifics and raw log

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
func (it *AccountsCollateralRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsCollateralRemoved)
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
		it.Event = new(AccountsCollateralRemoved)
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
func (it *AccountsCollateralRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsCollateralRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsCollateralRemoved represents a CollateralRemoved event raised by the Accounts contract.
type AccountsCollateralRemoved struct {
	Collateral common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralRemoved is a free log retrieval operation binding the contract event 0xd89d2ee68ab04dca0193f48a4aff55e20fa5ec0429a8a8c1c51b8dad6178a593.
//
// Solidity: event CollateralRemoved(address indexed collateral)
func (_Accounts *AccountsFilterer) FilterCollateralRemoved(opts *bind.FilterOpts, collateral []common.Address) (*AccountsCollateralRemovedIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "CollateralRemoved", collateralRule)
	if err != nil {
		return nil, err
	}
	return &AccountsCollateralRemovedIterator{contract: _Accounts.contract, event: "CollateralRemoved", logs: logs, sub: sub}, nil
}

// WatchCollateralRemoved is a free log subscription operation binding the contract event 0xd89d2ee68ab04dca0193f48a4aff55e20fa5ec0429a8a8c1c51b8dad6178a593.
//
// Solidity: event CollateralRemoved(address indexed collateral)
func (_Accounts *AccountsFilterer) WatchCollateralRemoved(opts *bind.WatchOpts, sink chan<- *AccountsCollateralRemoved, collateral []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "CollateralRemoved", collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsCollateralRemoved)
				if err := _Accounts.contract.UnpackLog(event, "CollateralRemoved", log); err != nil {
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

// ParseCollateralRemoved is a log parse operation binding the contract event 0xd89d2ee68ab04dca0193f48a4aff55e20fa5ec0429a8a8c1c51b8dad6178a593.
//
// Solidity: event CollateralRemoved(address indexed collateral)
func (_Accounts *AccountsFilterer) ParseCollateralRemoved(log types.Log) (*AccountsCollateralRemoved, error) {
	event := new(AccountsCollateralRemoved)
	if err := _Accounts.contract.UnpackLog(event, "CollateralRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsCreditIterator is returned from FilterCredit and is used to iterate over the raw logs and unpacked data for Credit events raised by the Accounts contract.
type AccountsCreditIterator struct {
	Event *AccountsCredit // Event containing the contract specifics and raw log

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
func (it *AccountsCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsCredit)
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
		it.Event = new(AccountsCredit)
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
func (it *AccountsCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsCredit represents a Credit event raised by the Accounts contract.
type AccountsCredit struct {
	Account    common.Address
	Collateral common.Address
	Amount     *big.Int
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCredit is a free log retrieval operation binding the contract event 0x27f2026b6afc5d271934dacdaaa950d69fd5fed4982a55b6c921e1667916f06c.
//
// Solidity: event Credit(address indexed account, address indexed collateral, uint256 amount, uint256 balance)
func (_Accounts *AccountsFilterer) FilterCredit(opts *bind.FilterOpts, account []common.Address, collateral []common.Address) (*AccountsCreditIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "Credit", accountRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return &AccountsCreditIterator{contract: _Accounts.contract, event: "Credit", logs: logs, sub: sub}, nil
}

// WatchCredit is a free log subscription operation binding the contract event 0x27f2026b6afc5d271934dacdaaa950d69fd5fed4982a55b6c921e1667916f06c.
//
// Solidity: event Credit(address indexed account, address indexed collateral, uint256 amount, uint256 balance)
func (_Accounts *AccountsFilterer) WatchCredit(opts *bind.WatchOpts, sink chan<- *AccountsCredit, account []common.Address, collateral []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "Credit", accountRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsCredit)
				if err := _Accounts.contract.UnpackLog(event, "Credit", log); err != nil {
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

// ParseCredit is a log parse operation binding the contract event 0x27f2026b6afc5d271934dacdaaa950d69fd5fed4982a55b6c921e1667916f06c.
//
// Solidity: event Credit(address indexed account, address indexed collateral, uint256 amount, uint256 balance)
func (_Accounts *AccountsFilterer) ParseCredit(log types.Log) (*AccountsCredit, error) {
	event := new(AccountsCredit)
	if err := _Accounts.contract.UnpackLog(event, "Credit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsCreditInstrumentIterator is returned from FilterCreditInstrument and is used to iterate over the raw logs and unpacked data for CreditInstrument events raised by the Accounts contract.
type AccountsCreditInstrumentIterator struct {
	Event *AccountsCreditInstrument // Event containing the contract specifics and raw log

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
func (it *AccountsCreditInstrumentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsCreditInstrument)
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
		it.Event = new(AccountsCreditInstrument)
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
func (it *AccountsCreditInstrumentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsCreditInstrumentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsCreditInstrument represents a CreditInstrument event raised by the Accounts contract.
type AccountsCreditInstrument struct {
	Account    common.Address
	Instrument *big.Int
	Amount     *big.Int
	Position   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreditInstrument is a free log retrieval operation binding the contract event 0x151a955d0b994ca6227ecca24a04f6bdbc5e39d3b71fba7df15fafda915a5978.
//
// Solidity: event CreditInstrument(address indexed account, uint256 indexed instrument, uint256 amount, int256 position)
func (_Accounts *AccountsFilterer) FilterCreditInstrument(opts *bind.FilterOpts, account []common.Address, instrument []*big.Int) (*AccountsCreditInstrumentIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "CreditInstrument", accountRule, instrumentRule)
	if err != nil {
		return nil, err
	}
	return &AccountsCreditInstrumentIterator{contract: _Accounts.contract, event: "CreditInstrument", logs: logs, sub: sub}, nil
}

// WatchCreditInstrument is a free log subscription operation binding the contract event 0x151a955d0b994ca6227ecca24a04f6bdbc5e39d3b71fba7df15fafda915a5978.
//
// Solidity: event CreditInstrument(address indexed account, uint256 indexed instrument, uint256 amount, int256 position)
func (_Accounts *AccountsFilterer) WatchCreditInstrument(opts *bind.WatchOpts, sink chan<- *AccountsCreditInstrument, account []common.Address, instrument []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "CreditInstrument", accountRule, instrumentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsCreditInstrument)
				if err := _Accounts.contract.UnpackLog(event, "CreditInstrument", log); err != nil {
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

// ParseCreditInstrument is a log parse operation binding the contract event 0x151a955d0b994ca6227ecca24a04f6bdbc5e39d3b71fba7df15fafda915a5978.
//
// Solidity: event CreditInstrument(address indexed account, uint256 indexed instrument, uint256 amount, int256 position)
func (_Accounts *AccountsFilterer) ParseCreditInstrument(log types.Log) (*AccountsCreditInstrument, error) {
	event := new(AccountsCreditInstrument)
	if err := _Accounts.contract.UnpackLog(event, "CreditInstrument", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsDebitIterator is returned from FilterDebit and is used to iterate over the raw logs and unpacked data for Debit events raised by the Accounts contract.
type AccountsDebitIterator struct {
	Event *AccountsDebit // Event containing the contract specifics and raw log

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
func (it *AccountsDebitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsDebit)
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
		it.Event = new(AccountsDebit)
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
func (it *AccountsDebitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsDebitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsDebit represents a Debit event raised by the Accounts contract.
type AccountsDebit struct {
	Account    common.Address
	Collateral common.Address
	Amount     *big.Int
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDebit is a free log retrieval operation binding the contract event 0xaa31acf5383e18b9a4ae2b3ccfa0faf705cf8e968ee6dd291db4b1bc1a2a8ebd.
//
// Solidity: event Debit(address indexed account, address indexed collateral, uint256 amount, uint256 balance)
func (_Accounts *AccountsFilterer) FilterDebit(opts *bind.FilterOpts, account []common.Address, collateral []common.Address) (*AccountsDebitIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "Debit", accountRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return &AccountsDebitIterator{contract: _Accounts.contract, event: "Debit", logs: logs, sub: sub}, nil
}

// WatchDebit is a free log subscription operation binding the contract event 0xaa31acf5383e18b9a4ae2b3ccfa0faf705cf8e968ee6dd291db4b1bc1a2a8ebd.
//
// Solidity: event Debit(address indexed account, address indexed collateral, uint256 amount, uint256 balance)
func (_Accounts *AccountsFilterer) WatchDebit(opts *bind.WatchOpts, sink chan<- *AccountsDebit, account []common.Address, collateral []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "Debit", accountRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsDebit)
				if err := _Accounts.contract.UnpackLog(event, "Debit", log); err != nil {
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

// ParseDebit is a log parse operation binding the contract event 0xaa31acf5383e18b9a4ae2b3ccfa0faf705cf8e968ee6dd291db4b1bc1a2a8ebd.
//
// Solidity: event Debit(address indexed account, address indexed collateral, uint256 amount, uint256 balance)
func (_Accounts *AccountsFilterer) ParseDebit(log types.Log) (*AccountsDebit, error) {
	event := new(AccountsDebit)
	if err := _Accounts.contract.UnpackLog(event, "Debit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsDebitInstrumentIterator is returned from FilterDebitInstrument and is used to iterate over the raw logs and unpacked data for DebitInstrument events raised by the Accounts contract.
type AccountsDebitInstrumentIterator struct {
	Event *AccountsDebitInstrument // Event containing the contract specifics and raw log

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
func (it *AccountsDebitInstrumentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsDebitInstrument)
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
		it.Event = new(AccountsDebitInstrument)
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
func (it *AccountsDebitInstrumentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsDebitInstrumentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsDebitInstrument represents a DebitInstrument event raised by the Accounts contract.
type AccountsDebitInstrument struct {
	Account    common.Address
	Instrument *big.Int
	Amount     *big.Int
	Position   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDebitInstrument is a free log retrieval operation binding the contract event 0x8437001e64ef3652dd8a3935534d9d6cfdb1fd4cbe6db8454c3369845cb45c5e.
//
// Solidity: event DebitInstrument(address indexed account, uint256 indexed instrument, uint256 amount, int256 position)
func (_Accounts *AccountsFilterer) FilterDebitInstrument(opts *bind.FilterOpts, account []common.Address, instrument []*big.Int) (*AccountsDebitInstrumentIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "DebitInstrument", accountRule, instrumentRule)
	if err != nil {
		return nil, err
	}
	return &AccountsDebitInstrumentIterator{contract: _Accounts.contract, event: "DebitInstrument", logs: logs, sub: sub}, nil
}

// WatchDebitInstrument is a free log subscription operation binding the contract event 0x8437001e64ef3652dd8a3935534d9d6cfdb1fd4cbe6db8454c3369845cb45c5e.
//
// Solidity: event DebitInstrument(address indexed account, uint256 indexed instrument, uint256 amount, int256 position)
func (_Accounts *AccountsFilterer) WatchDebitInstrument(opts *bind.WatchOpts, sink chan<- *AccountsDebitInstrument, account []common.Address, instrument []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "DebitInstrument", accountRule, instrumentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsDebitInstrument)
				if err := _Accounts.contract.UnpackLog(event, "DebitInstrument", log); err != nil {
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

// ParseDebitInstrument is a log parse operation binding the contract event 0x8437001e64ef3652dd8a3935534d9d6cfdb1fd4cbe6db8454c3369845cb45c5e.
//
// Solidity: event DebitInstrument(address indexed account, uint256 indexed instrument, uint256 amount, int256 position)
func (_Accounts *AccountsFilterer) ParseDebitInstrument(log types.Log) (*AccountsDebitInstrument, error) {
	event := new(AccountsDebitInstrument)
	if err := _Accounts.contract.UnpackLog(event, "DebitInstrument", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Accounts contract.
type AccountsDepositIterator struct {
	Event *AccountsDeposit // Event containing the contract specifics and raw log

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
func (it *AccountsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsDeposit)
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
		it.Event = new(AccountsDeposit)
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
func (it *AccountsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsDeposit represents a Deposit event raised by the Accounts contract.
type AccountsDeposit struct {
	Account    common.Address
	Collateral common.Address
	Amount     *big.Int
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed account, address indexed collateral, uint256 amount, uint256 balance)
func (_Accounts *AccountsFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address, collateral []common.Address) (*AccountsDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "Deposit", accountRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return &AccountsDepositIterator{contract: _Accounts.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed account, address indexed collateral, uint256 amount, uint256 balance)
func (_Accounts *AccountsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AccountsDeposit, account []common.Address, collateral []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "Deposit", accountRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsDeposit)
				if err := _Accounts.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed account, address indexed collateral, uint256 amount, uint256 balance)
func (_Accounts *AccountsFilterer) ParseDeposit(log types.Log) (*AccountsDeposit, error) {
	event := new(AccountsDeposit)
	if err := _Accounts.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsInstrumentAddedIterator is returned from FilterInstrumentAdded and is used to iterate over the raw logs and unpacked data for InstrumentAdded events raised by the Accounts contract.
type AccountsInstrumentAddedIterator struct {
	Event *AccountsInstrumentAdded // Event containing the contract specifics and raw log

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
func (it *AccountsInstrumentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsInstrumentAdded)
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
		it.Event = new(AccountsInstrumentAdded)
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
func (it *AccountsInstrumentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsInstrumentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsInstrumentAdded represents a InstrumentAdded event raised by the Accounts contract.
type AccountsInstrumentAdded struct {
	Instrument *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInstrumentAdded is a free log retrieval operation binding the contract event 0xb8597a358785eb73ad0a81079eb5e50ba0ac43dd023f47c0e0ff47e58a9495c8.
//
// Solidity: event InstrumentAdded(uint256 indexed instrument)
func (_Accounts *AccountsFilterer) FilterInstrumentAdded(opts *bind.FilterOpts, instrument []*big.Int) (*AccountsInstrumentAddedIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "InstrumentAdded", instrumentRule)
	if err != nil {
		return nil, err
	}
	return &AccountsInstrumentAddedIterator{contract: _Accounts.contract, event: "InstrumentAdded", logs: logs, sub: sub}, nil
}

// WatchInstrumentAdded is a free log subscription operation binding the contract event 0xb8597a358785eb73ad0a81079eb5e50ba0ac43dd023f47c0e0ff47e58a9495c8.
//
// Solidity: event InstrumentAdded(uint256 indexed instrument)
func (_Accounts *AccountsFilterer) WatchInstrumentAdded(opts *bind.WatchOpts, sink chan<- *AccountsInstrumentAdded, instrument []*big.Int) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "InstrumentAdded", instrumentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsInstrumentAdded)
				if err := _Accounts.contract.UnpackLog(event, "InstrumentAdded", log); err != nil {
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

// ParseInstrumentAdded is a log parse operation binding the contract event 0xb8597a358785eb73ad0a81079eb5e50ba0ac43dd023f47c0e0ff47e58a9495c8.
//
// Solidity: event InstrumentAdded(uint256 indexed instrument)
func (_Accounts *AccountsFilterer) ParseInstrumentAdded(log types.Log) (*AccountsInstrumentAdded, error) {
	event := new(AccountsInstrumentAdded)
	if err := _Accounts.contract.UnpackLog(event, "InstrumentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsInstrumentRemovedIterator is returned from FilterInstrumentRemoved and is used to iterate over the raw logs and unpacked data for InstrumentRemoved events raised by the Accounts contract.
type AccountsInstrumentRemovedIterator struct {
	Event *AccountsInstrumentRemoved // Event containing the contract specifics and raw log

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
func (it *AccountsInstrumentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsInstrumentRemoved)
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
		it.Event = new(AccountsInstrumentRemoved)
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
func (it *AccountsInstrumentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsInstrumentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsInstrumentRemoved represents a InstrumentRemoved event raised by the Accounts contract.
type AccountsInstrumentRemoved struct {
	Instrument *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInstrumentRemoved is a free log retrieval operation binding the contract event 0x08b9963dddd96e14a775b44c82852b5a482f5060ae86706c04fc5878c3fe3d8d.
//
// Solidity: event InstrumentRemoved(uint256 indexed instrument)
func (_Accounts *AccountsFilterer) FilterInstrumentRemoved(opts *bind.FilterOpts, instrument []*big.Int) (*AccountsInstrumentRemovedIterator, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "InstrumentRemoved", instrumentRule)
	if err != nil {
		return nil, err
	}
	return &AccountsInstrumentRemovedIterator{contract: _Accounts.contract, event: "InstrumentRemoved", logs: logs, sub: sub}, nil
}

// WatchInstrumentRemoved is a free log subscription operation binding the contract event 0x08b9963dddd96e14a775b44c82852b5a482f5060ae86706c04fc5878c3fe3d8d.
//
// Solidity: event InstrumentRemoved(uint256 indexed instrument)
func (_Accounts *AccountsFilterer) WatchInstrumentRemoved(opts *bind.WatchOpts, sink chan<- *AccountsInstrumentRemoved, instrument []*big.Int) (event.Subscription, error) {

	var instrumentRule []interface{}
	for _, instrumentItem := range instrument {
		instrumentRule = append(instrumentRule, instrumentItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "InstrumentRemoved", instrumentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsInstrumentRemoved)
				if err := _Accounts.contract.UnpackLog(event, "InstrumentRemoved", log); err != nil {
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

// ParseInstrumentRemoved is a log parse operation binding the contract event 0x08b9963dddd96e14a775b44c82852b5a482f5060ae86706c04fc5878c3fe3d8d.
//
// Solidity: event InstrumentRemoved(uint256 indexed instrument)
func (_Accounts *AccountsFilterer) ParseInstrumentRemoved(log types.Log) (*AccountsInstrumentRemoved, error) {
	event := new(AccountsInstrumentRemoved)
	if err := _Accounts.contract.UnpackLog(event, "InstrumentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsKeeperUpdatedIterator is returned from FilterKeeperUpdated and is used to iterate over the raw logs and unpacked data for KeeperUpdated events raised by the Accounts contract.
type AccountsKeeperUpdatedIterator struct {
	Event *AccountsKeeperUpdated // Event containing the contract specifics and raw log

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
func (it *AccountsKeeperUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsKeeperUpdated)
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
		it.Event = new(AccountsKeeperUpdated)
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
func (it *AccountsKeeperUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsKeeperUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsKeeperUpdated represents a KeeperUpdated event raised by the Accounts contract.
type AccountsKeeperUpdated struct {
	Keeper  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeeperUpdated is a free log retrieval operation binding the contract event 0x786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c00.
//
// Solidity: event KeeperUpdated(address indexed keeper, bool indexed allowed)
func (_Accounts *AccountsFilterer) FilterKeeperUpdated(opts *bind.FilterOpts, keeper []common.Address, allowed []bool) (*AccountsKeeperUpdatedIterator, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "KeeperUpdated", keeperRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &AccountsKeeperUpdatedIterator{contract: _Accounts.contract, event: "KeeperUpdated", logs: logs, sub: sub}, nil
}

// WatchKeeperUpdated is a free log subscription operation binding the contract event 0x786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c00.
//
// Solidity: event KeeperUpdated(address indexed keeper, bool indexed allowed)
func (_Accounts *AccountsFilterer) WatchKeeperUpdated(opts *bind.WatchOpts, sink chan<- *AccountsKeeperUpdated, keeper []common.Address, allowed []bool) (event.Subscription, error) {

	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "KeeperUpdated", keeperRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsKeeperUpdated)
				if err := _Accounts.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseKeeperUpdated(log types.Log) (*AccountsKeeperUpdated, error) {
	event := new(AccountsKeeperUpdated)
	if err := _Accounts.contract.UnpackLog(event, "KeeperUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsOwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the Accounts contract.
type AccountsOwnerUpdatedIterator struct {
	Event *AccountsOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *AccountsOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOwnerUpdated)
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
		it.Event = new(AccountsOwnerUpdated)
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
func (it *AccountsOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOwnerUpdated represents a OwnerUpdated event raised by the Accounts contract.
type AccountsOwnerUpdated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed newOwner)
func (_Accounts *AccountsFilterer) FilterOwnerUpdated(opts *bind.FilterOpts, newOwner []common.Address) (*AccountsOwnerUpdatedIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OwnerUpdated", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOwnerUpdatedIterator{contract: _Accounts.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address indexed newOwner)
func (_Accounts *AccountsFilterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *AccountsOwnerUpdated, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OwnerUpdated", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOwnerUpdated)
				if err := _Accounts.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseOwnerUpdated(log types.Log) (*AccountsOwnerUpdated, error) {
	event := new(AccountsOwnerUpdated)
	if err := _Accounts.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsRegisteredSigningKeyIterator is returned from FilterRegisteredSigningKey and is used to iterate over the raw logs and unpacked data for RegisteredSigningKey events raised by the Accounts contract.
type AccountsRegisteredSigningKeyIterator struct {
	Event *AccountsRegisteredSigningKey // Event containing the contract specifics and raw log

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
func (it *AccountsRegisteredSigningKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsRegisteredSigningKey)
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
		it.Event = new(AccountsRegisteredSigningKey)
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
func (it *AccountsRegisteredSigningKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsRegisteredSigningKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsRegisteredSigningKey represents a RegisteredSigningKey event raised by the Accounts contract.
type AccountsRegisteredSigningKey struct {
	Account      common.Address
	SigningKey   common.Address
	RegisterHash [32]byte
	Expiry       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredSigningKey is a free log retrieval operation binding the contract event 0x1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac.
//
// Solidity: event RegisteredSigningKey(address indexed account, address indexed signingKey, bytes32 indexed registerHash, uint256 expiry)
func (_Accounts *AccountsFilterer) FilterRegisteredSigningKey(opts *bind.FilterOpts, account []common.Address, signingKey []common.Address, registerHash [][32]byte) (*AccountsRegisteredSigningKeyIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var signingKeyRule []interface{}
	for _, signingKeyItem := range signingKey {
		signingKeyRule = append(signingKeyRule, signingKeyItem)
	}
	var registerHashRule []interface{}
	for _, registerHashItem := range registerHash {
		registerHashRule = append(registerHashRule, registerHashItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "RegisteredSigningKey", accountRule, signingKeyRule, registerHashRule)
	if err != nil {
		return nil, err
	}
	return &AccountsRegisteredSigningKeyIterator{contract: _Accounts.contract, event: "RegisteredSigningKey", logs: logs, sub: sub}, nil
}

// WatchRegisteredSigningKey is a free log subscription operation binding the contract event 0x1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac.
//
// Solidity: event RegisteredSigningKey(address indexed account, address indexed signingKey, bytes32 indexed registerHash, uint256 expiry)
func (_Accounts *AccountsFilterer) WatchRegisteredSigningKey(opts *bind.WatchOpts, sink chan<- *AccountsRegisteredSigningKey, account []common.Address, signingKey []common.Address, registerHash [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var signingKeyRule []interface{}
	for _, signingKeyItem := range signingKey {
		signingKeyRule = append(signingKeyRule, signingKeyItem)
	}
	var registerHashRule []interface{}
	for _, registerHashItem := range registerHash {
		registerHashRule = append(registerHashRule, registerHashItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "RegisteredSigningKey", accountRule, signingKeyRule, registerHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsRegisteredSigningKey)
				if err := _Accounts.contract.UnpackLog(event, "RegisteredSigningKey", log); err != nil {
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

// ParseRegisteredSigningKey is a log parse operation binding the contract event 0x1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac.
//
// Solidity: event RegisteredSigningKey(address indexed account, address indexed signingKey, bytes32 indexed registerHash, uint256 expiry)
func (_Accounts *AccountsFilterer) ParseRegisteredSigningKey(log types.Log) (*AccountsRegisteredSigningKey, error) {
	event := new(AccountsRegisteredSigningKey)
	if err := _Accounts.contract.UnpackLog(event, "RegisteredSigningKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsRevokedSigningKeyIterator is returned from FilterRevokedSigningKey and is used to iterate over the raw logs and unpacked data for RevokedSigningKey events raised by the Accounts contract.
type AccountsRevokedSigningKeyIterator struct {
	Event *AccountsRevokedSigningKey // Event containing the contract specifics and raw log

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
func (it *AccountsRevokedSigningKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsRevokedSigningKey)
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
		it.Event = new(AccountsRevokedSigningKey)
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
func (it *AccountsRevokedSigningKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsRevokedSigningKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsRevokedSigningKey represents a RevokedSigningKey event raised by the Accounts contract.
type AccountsRevokedSigningKey struct {
	Account    common.Address
	SigningKey common.Address
	RevokeHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevokedSigningKey is a free log retrieval operation binding the contract event 0xf0d72ac27c7920d0a69f77bbf64f8509a0307c49e54eee377861c46101ef25cb.
//
// Solidity: event RevokedSigningKey(address indexed account, address indexed signingKey, bytes32 indexed revokeHash)
func (_Accounts *AccountsFilterer) FilterRevokedSigningKey(opts *bind.FilterOpts, account []common.Address, signingKey []common.Address, revokeHash [][32]byte) (*AccountsRevokedSigningKeyIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var signingKeyRule []interface{}
	for _, signingKeyItem := range signingKey {
		signingKeyRule = append(signingKeyRule, signingKeyItem)
	}
	var revokeHashRule []interface{}
	for _, revokeHashItem := range revokeHash {
		revokeHashRule = append(revokeHashRule, revokeHashItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "RevokedSigningKey", accountRule, signingKeyRule, revokeHashRule)
	if err != nil {
		return nil, err
	}
	return &AccountsRevokedSigningKeyIterator{contract: _Accounts.contract, event: "RevokedSigningKey", logs: logs, sub: sub}, nil
}

// WatchRevokedSigningKey is a free log subscription operation binding the contract event 0xf0d72ac27c7920d0a69f77bbf64f8509a0307c49e54eee377861c46101ef25cb.
//
// Solidity: event RevokedSigningKey(address indexed account, address indexed signingKey, bytes32 indexed revokeHash)
func (_Accounts *AccountsFilterer) WatchRevokedSigningKey(opts *bind.WatchOpts, sink chan<- *AccountsRevokedSigningKey, account []common.Address, signingKey []common.Address, revokeHash [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var signingKeyRule []interface{}
	for _, signingKeyItem := range signingKey {
		signingKeyRule = append(signingKeyRule, signingKeyItem)
	}
	var revokeHashRule []interface{}
	for _, revokeHashItem := range revokeHash {
		revokeHashRule = append(revokeHashRule, revokeHashItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "RevokedSigningKey", accountRule, signingKeyRule, revokeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsRevokedSigningKey)
				if err := _Accounts.contract.UnpackLog(event, "RevokedSigningKey", log); err != nil {
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

// ParseRevokedSigningKey is a log parse operation binding the contract event 0xf0d72ac27c7920d0a69f77bbf64f8509a0307c49e54eee377861c46101ef25cb.
//
// Solidity: event RevokedSigningKey(address indexed account, address indexed signingKey, bytes32 indexed revokeHash)
func (_Accounts *AccountsFilterer) ParseRevokedSigningKey(log types.Log) (*AccountsRevokedSigningKey, error) {
	event := new(AccountsRevokedSigningKey)
	if err := _Accounts.contract.UnpackLog(event, "RevokedSigningKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Accounts contract.
type AccountsUpgradedIterator struct {
	Event *AccountsUpgraded // Event containing the contract specifics and raw log

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
func (it *AccountsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsUpgraded)
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
		it.Event = new(AccountsUpgraded)
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
func (it *AccountsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsUpgraded represents a Upgraded event raised by the Accounts contract.
type AccountsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Accounts *AccountsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AccountsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AccountsUpgradedIterator{contract: _Accounts.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Accounts *AccountsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AccountsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsUpgraded)
				if err := _Accounts.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseUpgraded(log types.Log) (*AccountsUpgraded, error) {
	event := new(AccountsUpgraded)
	if err := _Accounts.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsVersionInitializedIterator is returned from FilterVersionInitialized and is used to iterate over the raw logs and unpacked data for VersionInitialized events raised by the Accounts contract.
type AccountsVersionInitializedIterator struct {
	Event *AccountsVersionInitialized // Event containing the contract specifics and raw log

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
func (it *AccountsVersionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsVersionInitialized)
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
		it.Event = new(AccountsVersionInitialized)
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
func (it *AccountsVersionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsVersionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsVersionInitialized represents a VersionInitialized event raised by the Accounts contract.
type AccountsVersionInitialized struct {
	NewVersion *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVersionInitialized is a free log retrieval operation binding the contract event 0x7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d.
//
// Solidity: event VersionInitialized(uint256 indexed newVersion)
func (_Accounts *AccountsFilterer) FilterVersionInitialized(opts *bind.FilterOpts, newVersion []*big.Int) (*AccountsVersionInitializedIterator, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "VersionInitialized", newVersionRule)
	if err != nil {
		return nil, err
	}
	return &AccountsVersionInitializedIterator{contract: _Accounts.contract, event: "VersionInitialized", logs: logs, sub: sub}, nil
}

// WatchVersionInitialized is a free log subscription operation binding the contract event 0x7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d.
//
// Solidity: event VersionInitialized(uint256 indexed newVersion)
func (_Accounts *AccountsFilterer) WatchVersionInitialized(opts *bind.WatchOpts, sink chan<- *AccountsVersionInitialized, newVersion []*big.Int) (event.Subscription, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "VersionInitialized", newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsVersionInitialized)
				if err := _Accounts.contract.UnpackLog(event, "VersionInitialized", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseVersionInitialized(log types.Log) (*AccountsVersionInitialized, error) {
	event := new(AccountsVersionInitialized)
	if err := _Accounts.contract.UnpackLog(event, "VersionInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Accounts contract.
type AccountsWithdrawIterator struct {
	Event *AccountsWithdraw // Event containing the contract specifics and raw log

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
func (it *AccountsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsWithdraw)
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
		it.Event = new(AccountsWithdraw)
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
func (it *AccountsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsWithdraw represents a Withdraw event raised by the Accounts contract.
type AccountsWithdraw struct {
	Account        common.Address
	Collateral     common.Address
	WithdrawHash   [32]byte
	To             common.Address
	Amount         *big.Int
	WithdrawAmount *big.Int
	Balance        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xc0bd00c18050bbd2c9991c938b45550d2ad9fd31f67ba9d85597d8ef87df1f68.
//
// Solidity: event Withdraw(address indexed account, address indexed collateral, bytes32 indexed withdrawHash, address to, uint256 amount, uint256 withdrawAmount, uint256 balance)
func (_Accounts *AccountsFilterer) FilterWithdraw(opts *bind.FilterOpts, account []common.Address, collateral []common.Address, withdrawHash [][32]byte) (*AccountsWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var withdrawHashRule []interface{}
	for _, withdrawHashItem := range withdrawHash {
		withdrawHashRule = append(withdrawHashRule, withdrawHashItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "Withdraw", accountRule, collateralRule, withdrawHashRule)
	if err != nil {
		return nil, err
	}
	return &AccountsWithdrawIterator{contract: _Accounts.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xc0bd00c18050bbd2c9991c938b45550d2ad9fd31f67ba9d85597d8ef87df1f68.
//
// Solidity: event Withdraw(address indexed account, address indexed collateral, bytes32 indexed withdrawHash, address to, uint256 amount, uint256 withdrawAmount, uint256 balance)
func (_Accounts *AccountsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AccountsWithdraw, account []common.Address, collateral []common.Address, withdrawHash [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}
	var withdrawHashRule []interface{}
	for _, withdrawHashItem := range withdrawHash {
		withdrawHashRule = append(withdrawHashRule, withdrawHashItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "Withdraw", accountRule, collateralRule, withdrawHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsWithdraw)
				if err := _Accounts.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xc0bd00c18050bbd2c9991c938b45550d2ad9fd31f67ba9d85597d8ef87df1f68.
//
// Solidity: event Withdraw(address indexed account, address indexed collateral, bytes32 indexed withdrawHash, address to, uint256 amount, uint256 withdrawAmount, uint256 balance)
func (_Accounts *AccountsFilterer) ParseWithdraw(log types.Log) (*AccountsWithdraw, error) {
	event := new(AccountsWithdraw)
	if err := _Accounts.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsWithdrawProxyUpdatedIterator is returned from FilterWithdrawProxyUpdated and is used to iterate over the raw logs and unpacked data for WithdrawProxyUpdated events raised by the Accounts contract.
type AccountsWithdrawProxyUpdatedIterator struct {
	Event *AccountsWithdrawProxyUpdated // Event containing the contract specifics and raw log

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
func (it *AccountsWithdrawProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsWithdrawProxyUpdated)
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
		it.Event = new(AccountsWithdrawProxyUpdated)
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
func (it *AccountsWithdrawProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsWithdrawProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsWithdrawProxyUpdated represents a WithdrawProxyUpdated event raised by the Accounts contract.
type AccountsWithdrawProxyUpdated struct {
	WithdrawProxy common.Address
	Allowed       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawProxyUpdated is a free log retrieval operation binding the contract event 0x097359b7e2dca2adc06bfc6c3e4e8fcc69a2e916105499e8e4c60a569598cbc0.
//
// Solidity: event WithdrawProxyUpdated(address indexed withdrawProxy, bool indexed allowed)
func (_Accounts *AccountsFilterer) FilterWithdrawProxyUpdated(opts *bind.FilterOpts, withdrawProxy []common.Address, allowed []bool) (*AccountsWithdrawProxyUpdatedIterator, error) {

	var withdrawProxyRule []interface{}
	for _, withdrawProxyItem := range withdrawProxy {
		withdrawProxyRule = append(withdrawProxyRule, withdrawProxyItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "WithdrawProxyUpdated", withdrawProxyRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &AccountsWithdrawProxyUpdatedIterator{contract: _Accounts.contract, event: "WithdrawProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawProxyUpdated is a free log subscription operation binding the contract event 0x097359b7e2dca2adc06bfc6c3e4e8fcc69a2e916105499e8e4c60a569598cbc0.
//
// Solidity: event WithdrawProxyUpdated(address indexed withdrawProxy, bool indexed allowed)
func (_Accounts *AccountsFilterer) WatchWithdrawProxyUpdated(opts *bind.WatchOpts, sink chan<- *AccountsWithdrawProxyUpdated, withdrawProxy []common.Address, allowed []bool) (event.Subscription, error) {

	var withdrawProxyRule []interface{}
	for _, withdrawProxyItem := range withdrawProxy {
		withdrawProxyRule = append(withdrawProxyRule, withdrawProxyItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "WithdrawProxyUpdated", withdrawProxyRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsWithdrawProxyUpdated)
				if err := _Accounts.contract.UnpackLog(event, "WithdrawProxyUpdated", log); err != nil {
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

// ParseWithdrawProxyUpdated is a log parse operation binding the contract event 0x097359b7e2dca2adc06bfc6c3e4e8fcc69a2e916105499e8e4c60a569598cbc0.
//
// Solidity: event WithdrawProxyUpdated(address indexed withdrawProxy, bool indexed allowed)
func (_Accounts *AccountsFilterer) ParseWithdrawProxyUpdated(log types.Log) (*AccountsWithdrawProxyUpdated, error) {
	event := new(AccountsWithdrawProxyUpdated)
	if err := _Accounts.contract.UnpackLog(event, "WithdrawProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
