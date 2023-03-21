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
	EntryPrice *big.Int
	Position   *big.Int
}

// AccountsInterfaceSigningKey is an auto generated low-level Go binding around an user-defined struct.
type AccountsInterfaceSigningKey struct {
	SigningKey common.Address
	Expiry     *big.Int
}

// AccountsMetaData contains all meta data concerning the Accounts contract.
var AccountsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_domain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadRegisterSigningKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadRevokeSigningKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotKeeper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Credit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"}],\"name\":\"CreditInstrument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Debit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"}],\"name\":\"DebitInstrument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"InstrumentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"InstrumentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"KeeperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NewAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"RegisteredSigningKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"revokeHash\",\"type\":\"bytes32\"}],\"name\":\"RevokedSigningKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"VersionInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"WithdrawProxyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountsList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"addInstrument\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"addSigningKey\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorities\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collaterals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"creditInstrument\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"debitInstrument\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Balance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Position[]\",\"name\":\"positions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structAccountsInterface.SigningKey[]\",\"name\":\"signingKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getCollateralBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollaterals\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"getInstrumentEntryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"getInstrumentPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInstruments\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInsuranceFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getQuoteBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getSigningKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"}],\"name\":\"hasSigningKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"instruments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keepers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signingKeySig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountSig\",\"type\":\"bytes\"}],\"name\":\"registerSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"removeCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"removeInstrument\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"resetInstrumentPosition\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"revokeSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signingKeyExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalPositions\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawProxy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateWithdrawProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawInsuranceFund\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawProxies\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101a0604052306080523480156200001657600080fd5b506040516200508d3803806200508d833981016040819052620000399162000433565b6001600160a01b038116620000615760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811660a08190526040805163313ce56760e01b8152905163313ce567916004808201926020929091908290030181865afa158015620000ac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000d2919062000509565b60ff1660c052620000f08383620002c3602090811b6200319417901c565b60e052505060408051680aed2e8d0c8e4c2ee560bb1b6020808301919091527f6164647265737320636f6c6c61746572616c2c0000000000000000000000000060298301526a1859191c995cdcc81d1bcb60aa1b603c8301526e1d5a5b9d0c8d4d88185b5bdd5b9d0b608a1b60478301526c1d5a5b9d0c8d4d881cd85b1d0b609a1b60568301526b62797465733332206461746160a01b6063830152602960f81b606f83018190528351605081850301815260708401855280519083012061010052670a6d2cedc96caf2560c31b60908401526e1859191c995cdcc81858d8dbdd5b9d608a1b609884015260a783018190528351608881850301815260a88401855280519083012061012052680a4caced2e6e8cae4560bb1b60c88401526b1859191c995cdcc81ad95e4b60a21b60d18401526d75696e743235362065787069727960901b60dd84015260eb8301819052835160cc81850301815260ec8401855280519083012061014052660a4caecded6ca560cb1b61010c8401526a61646472657373206b657960a81b61011384015261011e830152825160ff81840301815261011f90920190925280519101206101605250736cdc77af264be2349f98f5d2cb58a453757caae76101805262000553565b6040516c08a92a06e626488dedac2d2dc5609b1b60208201526b1cdd1c9a5b99c81b985b594b60a21b602d8201526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398201526e1d5a5b9d0c8d4d8818da185a5b9259608a1b6048820152602960f81b605782015260009060580160405160208183030381529060405280519060200120836040516020016200035a919062000535565b60408051601f198184030181528282528051602091820120908301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201526080810183905260a00160405160208183030381529060405280519060200120905092915050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101562000400578181015183820152602001620003e6565b8381111562000410576000848401525b50505050565b80516001600160a01b03811681146200042e57600080fd5b919050565b6000806000606084860312156200044957600080fd5b83516001600160401b03808211156200046157600080fd5b818601915086601f8301126200047657600080fd5b8151818111156200048b576200048b620003cd565b604051601f8201601f19908116603f01168101908382118183101715620004b657620004b6620003cd565b81604052828152896020848701011115620004d057600080fd5b620004e3836020830160208801620003e3565b809750505050505060208401519150620005006040850162000416565b90509250925092565b6000602082840312156200051c57600080fd5b815160ff811681146200052e57600080fd5b9392505050565b6000825162000549818460208701620003e3565b9190910192915050565b60805160a05160c05160e0516101005161012051610140516101605161018051614a5962000634600039600081816103f4015281816129f701528181612a7501528181612c0d01526138b901526000610c270152600061131e015260006114bd015260006136ef015260008181610ba401528181610c06015281816113630152818161149c01526136ce0152600061054b015260008181610843015281816108de0152818161212a015281816121a3015261226101526000818161123e0152818161127e015281816116290152818161166901526116fc0152614a596000f3fe6080604052600436106103355760003560e01c80638869a3f2116101ab578063cf9b3c8d116100f7578063ea547fb811610095578063f0d2d5a81161006f578063f0d2d5a814610b52578063f160d36914610b72578063f698da2514610b92578063fbcbc0f114610bc657600080fd5b8063ea547fb814610ab8578063eeb97d3b14610b02578063efd1952314610b3257600080fd5b8063d658d2e9116100d1578063d658d2e914610a28578063dfa8ea4514610a58578063e30e838414610a78578063e89b0e1e14610a9857600080fd5b8063cf9b3c8d146109c8578063d3057877146109e8578063d5c2073014610a0857600080fd5b80639aeddeff11610164578063b7d5820b1161013e578063b7d5820b1461090e578063b8eb94e914610958578063c4d66de814610988578063c99d3a06146109a857600080fd5b80639aeddeff14610865578063ac759c7114610885578063af1aa3b1146108a557600080fd5b80638869a3f2146107705780638a48ac031461079f5780638da5cb5b146107b45780638e991606146107d457806391223d6914610801578063999b93af1461083157600080fd5b8063469048401161028557806362946d3b1161022357806372d39ed6116101fd57806372d39ed6146106fb57806378b926361461071b5780637de182c5146107305780638340f5491461075057600080fd5b806362946d3b146106615780636cd22eaf146106ae5780636eacd398146106ce57600080fd5b80634f1ef2861161025f5780634f1ef286146105df57806352d1902d146105f257806361d21920146106075780636205ed5e1461063457600080fd5b8063469048401461057f57806348e4ccbe1461059f57806349ad755b146105bf57600080fd5b806322bbad0b116102f25780633659cfe6116102cc5780633659cfe6146104e95780633bbd64bc146105095780633fd1e2bd146105395780634614be9f146103e557600080fd5b806322bbad0b1461044e578063241d400d1461048e578063262709e6146104bb57600080fd5b806306fdde031461033a57806307b1b8b21461038457806313af4035146103a657806314f326a1146103c6578063158626f7146103e55780631635717c1461042c575b600080fd5b34801561034657600080fd5b5061036e604051806040016040528060088152602001674163636f756e747360c01b81525081565b60405161037b9190614152565b60405180910390f35b34801561039057600080fd5b506103a461039f36600461421f565b610bf7565b005b3480156103b257600080fd5b506103a46103c136600461426d565b610ff9565b3480156103d257600080fd5b50600c545b60405190815260200161037b565b3480156103f157600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b03909116815260200161037b565b34801561043857600080fd5b5061044161100d565b60405161037b9190614288565b34801561045a57600080fd5b5061047e6104693660046142cc565b600b6020526000908152604090205460ff1681565b604051901515815260200161037b565b34801561049a57600080fd5b506104ae6104a936600461426d565b611065565b60405161037b91906142e5565b3480156104c757600080fd5b506104db6104d6366004614326565b6110de565b60405161037b929190614384565b3480156104f557600080fd5b506103a461050436600461426d565b611234565b34801561051557600080fd5b5061047e61052436600461426d565b60046020526000908152604090205460ff1681565b34801561054557600080fd5b5061056d7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff909116815260200161037b565b34801561058b57600080fd5b50600554610414906001600160a01b031681565b3480156105ab57600080fd5b506103a46105ba366004614398565b611310565b3480156105cb57600080fd5b506103d76105da366004614416565b6115f1565b6103a46105ed36600461421f565b61161f565b3480156105fe57600080fd5b506103d76116ef565b34801561061357600080fd5b506103d76106223660046142cc565b60106020526000908152604090205481565b34801561064057600080fd5b506103d761064f36600461426d565b600e6020526000908152604090205481565b34801561066d57600080fd5b5061047e61067c366004614440565b6001600160a01b039182166000908152600d602090815260408083209390941682526002909201909152205460ff1690565b3480156106ba57600080fd5b506103a46106c9366004614473565b6117a2565b3480156106da57600080fd5b506103d76106e936600461426d565b600f6020526000908152604090205481565b34801561070757600080fd5b506103d76107163660046144af565b611825565b34801561072757600080fd5b506104ae6118fb565b34801561073c57600080fd5b506104db61074b366004614326565b61195c565b34801561075c57600080fd5b506103d761076b366004614326565b611a98565b34801561077c57600080fd5b5061079061078b366004614551565b611adf565b60405161037b9392919061458a565b3480156107ab57600080fd5b506104ae611c43565b3480156107c057600080fd5b50600054610414906001600160a01b031681565b3480156107e057600080fd5b506107f46107ef3660046142cc565b611ca3565b60405161037b91906145a5565b34801561080d57600080fd5b5061047e61081c36600461426d565b60036020526000908152604090205460ff1681565b34801561083d57600080fd5b506104147f000000000000000000000000000000000000000000000000000000000000000081565b34801561087157600080fd5b506107f4610880366004614416565b611d8b565b34801561089157600080fd5b506107906108a0366004614551565b611f4a565b3480156108b157600080fd5b506103d76108c036600461426d565b6001600160a01b039081166000908152600d602090815260408083207f0000000000000000000000000000000000000000000000000000000000000000909416835260039093019052205490565b34801561091a57600080fd5b506103d7610929366004614440565b6001600160a01b039182166000908152600d602090815260408083209390941682526003909201909152205490565b34801561096457600080fd5b5061047e61097336600461426d565b60066020526000908152604090205460ff1681565b34801561099457600080fd5b506103a46109a336600461426d565b612094565b3480156109b457600080fd5b506103a46109c336600461426d565b61221e565b3480156109d457600080fd5b506104146109e33660046142cc565b612463565b3480156109f457600080fd5b506103a4610a03366004614473565b61248d565b348015610a1457600080fd5b506103a4610a23366004614473565b612510565b348015610a3457600080fd5b5061047e610a433660046142cc565b60076020526000908152604090205460ff1681565b348015610a6457600080fd5b506107f4610a73366004614326565b612593565b348015610a8457600080fd5b506107f4610a933660046142cc565b6126f5565b348015610aa457600080fd5b5061047e610ab336600461426d565b6128a5565b348015610ac457600080fd5b506103d7610ad3366004614416565b6001600160a01b03919091166000908152600d6020908152604080832093835260049093019052206001015490565b348015610b0e57600080fd5b5061047e610b1d36600461426d565b60096020526000908152604090205460ff1681565b348015610b3e57600080fd5b506103d7610b4d366004614326565b6128e6565b348015610b5e57600080fd5b506103a4610b6d36600461426d565b612c6a565b348015610b7e57600080fd5b506103a4610b8d36600461426d565b612d62565b348015610b9e57600080fd5b506103d77f000000000000000000000000000000000000000000000000000000000000000081565b348015610bd257600080fd5b50610be6610be136600461426d565b612ddb565b60405161037b959493929190614648565b610bff61329c565b6000610ca07f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000085604051602001610c6a9291909182526001600160a01b0316602082015260400190565b6040516020818303038152906040528051906020012061190160f01b600090815260029290925260229081526042822091905290565b3360009081526004602052604090205490915060ff16610ce4578060016040516001622b6d3960e21b03198152600401610cdb929190614384565b60405180910390fd5b6001600160a01b038316610d13578060036040516001622b6d3960e21b03198152600401610cdb929190614384565b6001600160a01b0383166000908152600e60205260408120549003610d535780600a6040516001622b6d3960e21b03198152600401610cdb929190614384565b6000610d5f82846132c5565b90506001600160a01b038116610d90578160076040516001622b6d3960e21b03198152600401610cdb929190614384565b6001600160a01b038082166000908152600d60209081526040808320938816835260029093019052205460ff16610de25781600a6040516001622b6d3960e21b03198152600401610cdb929190614384565b6001600160a01b0381166000908152600d6020908152604080832060010180548251818502810185019093528083528493830182828015610e4c57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e2e575b505050505090505b8051821015610e9b57856001600160a01b0316818381518110610e7957610e796146dc565b60200260200101516001600160a01b03160315610e9b57816001019150610e54565b80518210610ec45783600a6040516001622b6d3960e21b03198152600401610cdb929190614384565b8060018251610ed39190614708565b81518110610ee357610ee36146dc565b6020026020010151600d6000856001600160a01b03166001600160a01b031681526020019081526020016000206001018381548110610f2457610f246146dc565b600091825260208083209190910180546001600160a01b0319166001600160a01b039485161790559185168152600d90915260409020600101805480610f6c57610f6c61471f565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03858116808452600d83526040808520928b168086526002909301909352828420805460ff19169055915187939192917ff0d72ac27c7920d0a69f77bbf64f8509a0307c49e54eee377861c46101ef25cb91a45050600160025550505050565b6110016133e8565b61100a81613414565b50565b6060600a80548060200260200160405190810160405280929190818152602001828054801561105b57602002820191906000526020600020905b815481526020019060010190808311611047575b5050505050905090565b6001600160a01b0381166000908152600d60209081526040918290206001018054835181840281018401909452808452606093928301828280156110d257602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116110b4575b50505050509050919050565b33600090815260036020526040812054819060ff1661111057604051631890934360e11b815260040160405180910390fd5b61111861329c565b6001600160a01b03841660009081526009602052604090205460ff166111445750600090506004611225565b61114d85613483565b506001600160a01b038086166000908152600d6020908152604080832093881683526003909301905290812054611185908590614735565b6001600160a01b038088166000908152600d60209081526040808320938a16835260039093018152828220849055600f9052908120805492935086929091906111cf908490614735565b909155505060408051858152602081018390526001600160a01b0380881692908916917fcb9aff7bd66be0077d57437ece119a9f166aad37de4491ac4f92fdb2c1ad8d3491015b60405180910390a39150600090505b60016002559094909350915050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361127c5760405162461bcd60e51b8152600401610cdb90614774565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166112c56000805160206149e6833981519152546001600160a01b031690565b6001600160a01b0316146112eb5760405162461bcd60e51b8152600401610cdb906147c0565b6112f48161354c565b6040805160008082526020820190925261100a91839190613554565b61131861329c565b604080517f000000000000000000000000000000000000000000000000000000000000000060208201526001600160a01b038616918101919091526060810184905260009061138b907f000000000000000000000000000000000000000000000000000000000000000090608001610c6a565b3360009081526004602052604090205490915060ff166113c357806001604051634e7391e160e11b8152600401610cdb929190614384565b6001600160a01b0385166113ef57806003604051634e7391e160e11b8152600401610cdb929190614384565b8360000361141557806002604051634e7391e160e11b8152600401610cdb929190614384565b6001600160a01b0385166000908152600e6020526040902054156114515780600a604051634e7391e160e11b8152600401610cdb929190614384565b600061145d82846132c5565b90506001600160a01b03811661148b57816007604051634e7391e160e11b8152600401610cdb929190614384565b61149481613483565b5060006115007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000084604051602001610c6a9291909182526001600160a01b0316602082015260400190565b9050600061150e82876132c5565b9050876001600160a01b0316816001600160a01b0316146115475783600b604051634e7391e160e11b8152600401610cdb929190614384565b6001600160a01b038881166000818152600e602090815260408083208c9055938716808352600d82528483206001808201805480830182559086528486200180546001600160a01b031916871790558585526002909101835292859020805460ff191690931790925592518a815287937f1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac910160405180910390a450506001600255505050505050565b6001600160a01b0382166000908152600d602090815260408083208484526004019091529020545b92915050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036116675760405162461bcd60e51b8152600401610cdb90614774565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166116b06000805160206149e6833981519152546001600160a01b031690565b6001600160a01b0316146116d65760405162461bcd60e51b8152600401610cdb906147c0565b6116df8261354c565b6116eb82826001613554565b5050565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461178f5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610cdb565b506000805160206149e683398151915290565b6117aa6133e8565b6001600160a01b0382166117d15760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff191685151590811790915590519092917fc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc91a35050565b600061182f61329c565b6000806118418a8a8a8a8a8a8a6136c4565b9150915061184e82613483565b506118636001600160a01b038b168a89613aab565b6001600160a01b03891660009081526006602052604090205460ff16156118e957604051631e6fd1a760e31b81526001600160a01b038a169063f37e8d38906118b69085908e908c908b9060040161480c565b600060405180830381600087803b1580156118d057600080fd5b505af11580156118e4573d6000803e3d6000fd5b505050505b60016002559998505050505050505050565b6060600880548060200260200160405190810160405280929190818152602001828054801561105b57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611935575050505050905090565b33600090815260036020526040812054819060ff1661198e57604051631890934360e11b815260040160405180910390fd5b61199661329c565b6001600160a01b03841660009081526009602052604090205460ff166119c25750600090506004611225565b6119cb85613483565b506001600160a01b038086166000908152600d6020908152604080832093881683526003909301905290812054611a0390859061483f565b6001600160a01b038088166000908152600d60209081526040808320938a16835260039093018152828220849055600f905290812080549293508692909190611a4d90849061483f565b909155505060408051858152602081018390526001600160a01b0380881692908916917fc9a1b9e3b11def04d81baa28220fa980997040aef1abec95d1cd2bcfd4ba8a659101611216565b6000611aa261329c565b611aab84613483565b506000611ab9858585613b29565b9050611ad06001600160a01b038516333086613c85565b90505b60016002559392505050565b336000908152600360205260408120548190819060ff16611b1357604051631890934360e11b815260040160405180910390fd5b611b1b61329c565b6000868152600b602052604090205460ff16611b405750600091508190506005611c31565b6001600160ff1b03851115611b5e5750600091508190506002611c31565b611b6787613483565b506001600160a01b0387166000908152600d60209081526040808320898452600401909152812081908190611b9d908989613d0f565b60008c815260106020526040812080549497509295509093508a92611bc390849061483f565b90915550506040805189815260208101899052908101849052606081018390526080810182905289906001600160a01b038c16907f27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae653639060a0015b60405180910390a391945090925060009150505b60016002819055509450945094915050565b6060600c80548060200260200160405190810160405280929190818152602001828054801561105b576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611935575050505050905090565b3360009081526003602052604081205460ff16611cd357604051631890934360e11b815260040160405180910390fd5b611cdb61329c565b81600003611ceb57506002611d81565b6000828152600b602052604090205460ff1615611d0a57506005611d81565b600a805460018181019092557fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8018390556000838152600b6020526040808220805460ff1916909317909255905183917fb8597a358785eb73ad0a81079eb5e50ba0ac43dd023f47c0e0ff47e58a9495c891a25060005b6001600255919050565b3360009081526003602052604081205460ff16611dbb57604051631890934360e11b815260040160405180910390fd5b611dc361329c565b6000828152600b602052604090205460ff16611de157506005611f3f565b611dea83613483565b506001600160a01b0383166000908152600d6020908152604080832085845260040190915281208054828255600190910182905590811315611ea95760008381526010602052604081208054839290611e44908490614735565b909155505060408051828152600060208201819052818301819052606082018190526080820152905184916001600160a01b038716917fc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d19181900360a00190a3611f39565b6000811215611f395760008381526010602052604081208054839290611ed0908490614735565b909155508390506001600160a01b0385167f27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae65363611f0b84614880565b6040805191825260006020830181905290820181905260608201819052608082015260a00160405180910390a35b60009150505b600160025592915050565b336000908152600360205260408120548190819060ff16611f7e57604051631890934360e11b815260040160405180910390fd5b611f8661329c565b6000868152600b602052604090205460ff16611fab5750600091508190506005611c31565b6001600160ff1b03851115611fc95750600091508190506002611c31565b611fd287613483565b5060008080612010611fe389614880565b6001600160a01b038c166000908152600d602090815260408083208e845260040190915290209089613d0f565b60008c815260106020526040812080549497509295509093508a92612036908490614735565b90915550506040805189815260208101899052908101849052606081018390526080810182905289906001600160a01b038c16907fc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d19060a001611c1d565b60016120a08180614708565b600154146120c1576040516302ed543d60e51b815260040160405180910390fd5b806001036120cf5760016002555b6120d882613414565b600580546001600160a01b038085166001600160a01b031992831681179093556008805460018181019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180547f0000000000000000000000000000000000000000000000000000000000000000909316929093168217909255600090815260096020526040808220805460ff191690931790925590517f7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc29190a26040516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016907f7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f90600090a2600181905560405181907f7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d90600090a25050565b6122266133e8565b6001600160a01b03811660009081526009602052604090205460ff1661225f576040516368f7a67560e11b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316036122b1576040516368f7a67560e11b815260040160405180910390fd5b600080600880548060200260200160405190810160405280929190818152602001828054801561230a57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116122ec575b505050505090505b805182101561235957826001600160a01b0316818381518110612337576123376146dc565b60200260200101516001600160a01b0316031561235957816001019150612312565b8051821061237a576040516368f7a67560e11b815260040160405180910390fd5b80600182516123899190614708565b81518110612399576123996146dc565b6020026020010151600883815481106123b4576123b46146dc565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060088054806123f3576123f361471f565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03851680835260099091526040808320805460ff191690555190917fd89d2ee68ab04dca0193f48a4aff55e20fa5ec0429a8a8c1c51b8dad6178a59391a2505050565b600c818154811061247357600080fd5b6000918252602090912001546001600160a01b0316905081565b6124956133e8565b6001600160a01b0382166124bc5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260046020526040808220805460ff191685151590811790915590519092917f786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c0091a35050565b6125186133e8565b6001600160a01b03821661253f5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260066020526040808220805460ff191685151590811790915590519092917f097359b7e2dca2adc06bfc6c3e4e8fcc69a2e916105499e8e4c60a569598cbc091a35050565b3360009081526003602052604081205460ff166125c357604051631890934360e11b815260040160405180910390fd5b6125cb61329c565b6001600160a01b0383166000908152600e6020526040902054156125f15750600b611ad3565b6001600160a01b0383161580612605575081155b1561261257506002611ad3565b6001600160a01b038085166000908152600d60209081526040808320938716835260029093019052205460ff161561264c57506007611ad3565b6001600160a01b038381166000818152600e60209081526040808320879055938816808352600d82528483206001808201805480830182559086528486200180546001600160a01b0319168717905585855260029091018352858420805460ff1916909117905593518681529193917f1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac910160405180910390a450600060016002559392505050565b3360009081526003602052604081205460ff1661272557604051631890934360e11b815260040160405180910390fd5b61272d61329c565b6000828152600b602052604090205460ff1661274b57506005611d81565b600080600a80548060200260200160405190810160405280929190818152602001828054801561279a57602002820191906000526020600020905b815481526020019060010190808311612786575b505050505090505b80518210156127d757838183815181106127be576127be6146dc565b602002602001015103156127d7578160010191506127a2565b805182106127ea57600592505050611d81565b80600182516127f99190614708565b81518110612809576128096146dc565b6020026020010151600a8381548110612824576128246146dc565b600091825260209091200155600a8054806128415761284161471f565b600082815260208082208301600019908101839055909201909255858252600b90526040808220805460ff191690555185917f08b9963dddd96e14a775b44c82852b5a482f5060ae86706c04fc5878c3fe3d8d91a250506001600255506000919050565b3360009081526003602052604081205460ff166128d557604051631890934360e11b815260040160405180910390fd5b6128dd61329c565b611f3f82613483565b60006128f061329c565b3360009081526004602052604090205460ff1661293357600080516020614a2d833981519152600160405163c802b7f360e01b8152600401610cdb929190614384565b8160000361296757600080516020614a2d833981519152600260405163c802b7f360e01b8152600401610cdb929190614384565b6001600160a01b0383166129a157600080516020614a2d833981519152600360405163c802b7f360e01b8152600401610cdb929190614384565b6001600160a01b03841660009081526009602052604090205460ff166129ed57600080516020614a2d833981519152600460405163c802b7f360e01b8152600401610cdb929190614384565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166000908152600d60209081526040808320938816835260039093019052205482811215612a6b57600080516020614a2d833981519152600860405163c802b7f360e01b8152600401610cdb929190614384565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166000908152600d602090815260408083209389168352600390930181528282209386900393849055600f90529081208054859290612ad5908490614735565b90915550612aef90506001600160a01b0386168585613aab565b6001600160a01b03841660009081526006602052604090205460ff1615612b9157600554604051631e6fd1a760e31b81526001600160a01b03918216600482015286821660248201526044810185905260806064820152600060848201529085169063f37e8d389060a401600060405180830381600087803b158015612b7457600080fd5b505af1158015612b88573d6000803e3d6000fd5b50505050612bd2565b6005546001600160a01b03858116911614612bd257600080516020614a2d833981519152600c60405163c802b7f360e01b8152600401610cdb929190614384565b604080516001600160a01b0386811682526020820186905291810185905260608101839052600080516020614a2d83398151915291808816917f0000000000000000000000000000000000000000000000000000000000000000909116907f66efaee033f9c338e1e0c145c8653403e5f05eddfa1e35f67931c6ba3e231ea09060800160405180910390a46001600255949350505050565b612c726133e8565b6001600160a01b038116612c995760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811660009081526009602052604090205460ff1615612cd3576040516368f7a67560e11b815260040160405180910390fd5b6008805460018082019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180546001600160a01b0319166001600160a01b038416908117909155600081815260096020526040808220805460ff1916909417909355915190917f7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f91a250565b612d6a6133e8565b6001600160a01b038116612d915760405163d92e233d60e01b815260040160405180910390fd5b600580546001600160a01b0319166001600160a01b0383169081179091556040517f7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc290600090a250565b6001600160a01b0381166000908152600d602052604081205460085460ff90911691606091829182918067ffffffffffffffff811115612e1d57612e1d61417c565b604051908082528060200260200182016040528015612e6257816020015b6040805180820190915260008082526020820152815260200190600190039081612e3b5790505b50945060005b81811015612ef357600060088281548110612e8557612e856146dc565b60009182526020808320909101546040805180820182526001600160a01b03928316808252928e168552600d84528185208386526003018452932054918301919091528851909250889084908110612edf57612edf6146dc565b602090810291909101015250600101612e68565b50600a548067ffffffffffffffff811115612f1057612f1061417c565b604051908082528060200260200182016040528015612f6557816020015b612f5260405180606001604052806000815260200160008152602001600081525090565b815260200190600190039081612f2e5790505b50945060005b81811015613007576000600a8281548110612f8857612f886146dc565b6000918252602080832090910154604080516060810182528281526001600160a01b038f168552600d845281852083865260040180855282862060018101548387015295849052909352925492820192909252885191925090889084908110612ff357612ff36146dc565b602090810291909101015250600101612f6b565b506001600160a01b0388166000908152600d602090815260408083206001018054825181850281018501909352808352919290919083018282801561307557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613057575b50505050509050805167ffffffffffffffff8111156130965761309661417c565b6040519080825280602002602001820160405280156130db57816020015b60408051808201909152600080825260208201528152602001906001900390816130b45790505b50945060005b8151811015613184576040518060400160405280838381518110613107576131076146dc565b60200260200101516001600160a01b03168152602001600e6000858581518110613133576131336146dc565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054815250868281518110613171576131716146dc565b60209081029190910101526001016130e1565b5043935050505091939590929450565b6040516c08a92a06e626488dedac2d2dc5609b1b60208201526b1cdd1c9a5b99c81b985b594b60a21b602d8201526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398201526e1d5a5b9d0c8d4d8818da185a5b9259608a1b6048820152602960f81b60578201526000906058016040516020818303038152906040528051906020012083604051602001613229919061489c565b60408051601f198184030181528282528051602091820120908301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201526080810183905260a00160405160208183030381529060405280519060200120905092915050565b6002546001146132bf5760405163558a1e0360e11b815260040160405180910390fd5b60028055565b60008060008084516040036132f757505050602082015160408301516001600160ff1b0381169060ff1c601b0161334c565b84516041036133405750505060208201516040830151606084015160001a601b811480159061332a57508060ff16601c14155b1561333b5760009350505050611619565b61334c565b60009350505050611619565b7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156133805760009350505050611619565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa1580156133d3573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6000546001600160a01b03163314613412576040516282b42960e81b815260040160405180910390fd5b565b6001600160a01b03811661343b5760405163d92e233d60e01b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b038316908117825560405190917f4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b91a250565b60006001600160a01b03821615613547576001600160a01b0382166000908152600d602052604090205460ff16156134bd57506001919050565b6001600160a01b0382166000818152600d6020526040808220805460ff19166001908117909155600c8054918201815583527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c70180546001600160a01b03191684179055517fef4ab4f35cd2027fcc6364f430a86765b6bbd24462cd31f5a6d09bb74241aaf19190a25b919050565b61100a6133e8565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561358c5761358783613e97565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156135e6575060408051601f3d908101601f191682019092526135e3918101906148b8565b60015b6136495760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610cdb565b6000805160206149e683398151915281146136b85760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610cdb565b50613587838383613f33565b600080600061375e7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008c8c8c8b8b80519060200120604051602001610c6a969594939291909586526001600160a01b0394851660208701529290931660408501526060840152608083019190915260a082015260c00190565b3360009081526004602052604090205490915060ff166137965780600160405163c802b7f360e01b8152600401610cdb929190614384565b8715806137a1575086155b156137c45780600260405163c802b7f360e01b8152600401610cdb929190614384565b6001600160a01b0389166137f05780600360405163c802b7f360e01b8152600401610cdb929190614384565b6001600160a01b038a1660009081526009602052604090205460ff1661382e5780600460405163c802b7f360e01b8152600401610cdb929190614384565b60008181526007602052604090205460ff16156138635780600660405163c802b7f360e01b8152600401610cdb929190614384565b86881415801561387557506000198814155b156138985780600960405163c802b7f360e01b8152600401610cdb929190614384565b60006138a482866132c5565b90506001600160a01b03811615806138ed57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316145b156139105781600760405163c802b7f360e01b8152600401610cdb929190614384565b6001600160a01b03808c166000818152600f60209081526040808320549486168352600d82528083209383526003909301905220548982128061395257508981125b156139755783600860405163c802b7f360e01b8152600401610cdb929190614384565b8982039150898103905081600f60008f6001600160a01b03166001600160a01b031681526020019081526020016000208190555080600d6000856001600160a01b03166001600160a01b0316815260200190815260200160002060030160008f6001600160a01b03166001600160a01b031681526020019081526020016000208190555060016007600086815260200190815260200160002060006101000a81548160ff021916908315150217905550838d6001600160a01b0316846001600160a01b03167f66efaee033f9c338e1e0c145c8653403e5f05eddfa1e35f67931c6ba3e231ea08f8f8f87604051613a9094939291906001600160a01b0394909416845260208401929092526040830152606082015260800190565b60405180910390a4919c919b50909950505050505050505050565b600060405163a9059cbb60e01b8152836004820152826024820152602060006044836000895af13d15601f3d1160016000511416171691505080613b235760405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b6044820152606401610cdb565b50505050565b60006001600160a01b038416613b525760405163d92e233d60e01b815260040160405180910390fd5b81600003613b7357604051631f2a200560e01b815260040160405180910390fd5b6001600160a01b03831660009081526009602052604090205460ff16613bac576040516368f7a67560e11b815260040160405180910390fd5b6001600160a01b038085166000908152600d6020908152604080832093871683526003909301905290812054613be390849061483f565b6001600160a01b038087166000908152600d60209081526040808320938916835260039093018152828220849055600f905290812080549293508592909190613c2d90849061483f565b909155505060408051848152602081018390526001600160a01b0380871692908816917f8cd158929f48bc4a1b49c8c3873d0561de29c5b993767bbe115997adeb1113cf910160405180910390a390505b9392505050565b60006040516323b872dd60e01b81528460048201528360248201528260448201526020600060648360008a5af13d15601f3d1160016000511416171691505080613d085760405162461bcd60e51b81526020600482015260146024820152731514905394d1915497d19493d357d1905253115160621b6044820152606401610cdb565b5050505050565b600080600084600003613d2f575050835460018501549091506000613e8e565b8554613d3b868261483f565b935080600003613d545760018701859055849250613e89565b83600003613d7957613d6b81886001015487613f58565b600060018901559150613e89565b600081128015613d895750600084135b80613d9f5750600081138015613d9f5750600084125b15613dc457613db381886001015487613f58565b600188018690558593509150613e89565b600081138015613dd357508084135b15613e145783613de386886148d1565b6001890154613df290846148d1565b613dfc91906148f0565b613e06919061491e565b600188018190559250613e89565b600081128015613e2357508084125b15613e5d57613e3184614880565b85613e3b88614880565b613e4591906148d1565b6001890154613e5384614880565b613df291906148d1565b60018701549250620f4240613e728685614735565b613e7c9088614932565b613e8691906149b7565b91505b508286555b93509350939050565b6001600160a01b0381163b613f045760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610cdb565b6000805160206149e683398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b613f3c83613f84565b600082511180613f495750805b1561358757613b238383613fc4565b6000620f4240613f688484614735565b613f729086614932565b613f7c91906149b7565b949350505050565b613f8d81613e97565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060613c7e8383604051806060016040528060278152602001614a066027913960606001600160a01b0384163b61404c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610cdb565b600080856001600160a01b031685604051614067919061489c565b600060405180830381855af49150503d80600081146140a2576040519150601f19603f3d011682016040523d82523d6000602084013e6140a7565b606091505b50915091506140b78282866140c1565b9695505050505050565b606083156140d0575081613c7e565b8251156140e05782518084602001fd5b8160405162461bcd60e51b8152600401610cdb9190614152565b60005b838110156141155781810151838201526020016140fd565b83811115613b235750506000910152565b6000815180845261413e8160208601602086016140fa565b601f01601f19169290920160200192915050565b602081526000613c7e6020830184614126565b80356001600160a01b038116811461354757600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f8301126141a357600080fd5b813567ffffffffffffffff808211156141be576141be61417c565b604051601f8301601f19908116603f011681019082821181831017156141e6576141e661417c565b816040528381528660208588010111156141ff57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561423257600080fd5b61423b83614165565b9150602083013567ffffffffffffffff81111561425757600080fd5b61426385828601614192565b9150509250929050565b60006020828403121561427f57600080fd5b613c7e82614165565b6020808252825182820181905260009190848201906040850190845b818110156142c0578351835292840192918401916001016142a4565b50909695505050505050565b6000602082840312156142de57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156142c05783516001600160a01b031683529284019291840191600101614301565b60008060006060848603121561433b57600080fd5b61434484614165565b925061435260208501614165565b9150604084013590509250925092565b600d811061438057634e487b7160e01b600052602160045260246000fd5b9052565b82815260408101613c7e6020830184614362565b600080600080608085870312156143ae57600080fd5b6143b785614165565b935060208501359250604085013567ffffffffffffffff808211156143db57600080fd5b6143e788838901614192565b935060608701359150808211156143fd57600080fd5b5061440a87828801614192565b91505092959194509250565b6000806040838503121561442957600080fd5b61443283614165565b946020939093013593505050565b6000806040838503121561445357600080fd5b61445c83614165565b915061446a60208401614165565b90509250929050565b6000806040838503121561448657600080fd5b61448f83614165565b9150602083013580151581146144a457600080fd5b809150509250929050565b600080600080600080600060e0888a0312156144ca57600080fd5b6144d388614165565b96506144e160208901614165565b955060408801359450606088013593506080880135925060a088013567ffffffffffffffff8082111561451357600080fd5b61451f8b838c01614192565b935060c08a013591508082111561453557600080fd5b506145428a828b01614192565b91505092959891949750929550565b6000806000806080858703121561456757600080fd5b61457085614165565b966020860135965060408601359560600135945092505050565b8381526020810183905260608101613f7c6040830184614362565b602081016116198284614362565b600081518084526020808501945080840160005b838110156145f957815180518852838101518489015260409081015190880152606090960195908201906001016145c7565b509495945050505050565b600081518084526020808501945080840160005b838110156145f957815180516001600160a01b031688528301518388015260409096019590820190600101614618565b600060a082018715158352602060a08185015281885180845260c086019150828a01935060005b8181101561469f57845180516001600160a01b03168452840151848401529383019360409092019160010161466f565b505084810360408601526146b381896145b3565b9250505082810360608401526146c98186614604565b9150508260808301529695505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008282101561471a5761471a6146f2565b500390565b634e487b7160e01b600052603160045260246000fd5b60008083128015600160ff1b850184121615614753576147536146f2565b6001600160ff1b038401831381161561476e5761476e6146f2565b50500390565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906140b790830184614126565b600080821280156001600160ff1b0384900385131615614861576148616146f2565b600160ff1b839003841281161561487a5761487a6146f2565b50500190565b6000600160ff1b8201614895576148956146f2565b5060000390565b600082516148ae8184602087016140fa565b9190910192915050565b6000602082840312156148ca57600080fd5b5051919050565b60008160001904831182151516156148eb576148eb6146f2565b500290565b60008219821115614903576149036146f2565b500190565b634e487b7160e01b600052601260045260246000fd5b60008261492d5761492d614908565b500490565b60006001600160ff1b0381841382841380821686840486111615614958576149586146f2565b600160ff1b6000871282811687830589121615614977576149776146f2565b60008712925087820587128484161615614993576149936146f2565b878505871281841616156149a9576149a96146f2565b505050929093029392505050565b6000826149c6576149c6614908565b600160ff1b8214600019841416156149e0576149e06146f2565b50059056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65647d551b4e977cb76b47865bfb6cdc77af264be2349f98f5d2cb58a453757caae7a164736f6c634300080f000a",
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

// AccountsList is a free data retrieval call binding the contract method 0xcf9b3c8d.
//
// Solidity: function accountsList(uint256 ) view returns(address)
func (_Accounts *AccountsCaller) AccountsList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "accountsList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountsList is a free data retrieval call binding the contract method 0xcf9b3c8d.
//
// Solidity: function accountsList(uint256 ) view returns(address)
func (_Accounts *AccountsSession) AccountsList(arg0 *big.Int) (common.Address, error) {
	return _Accounts.Contract.AccountsList(&_Accounts.CallOpts, arg0)
}

// AccountsList is a free data retrieval call binding the contract method 0xcf9b3c8d.
//
// Solidity: function accountsList(uint256 ) view returns(address)
func (_Accounts *AccountsCallerSession) AccountsList(arg0 *big.Int) (common.Address, error) {
	return _Accounts.Contract.AccountsList(&_Accounts.CallOpts, arg0)
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

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Accounts *AccountsCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Accounts *AccountsSession) FeeRecipient() (common.Address, error) {
	return _Accounts.Contract.FeeRecipient(&_Accounts.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Accounts *AccountsCallerSession) FeeRecipient() (common.Address, error) {
	return _Accounts.Contract.FeeRecipient(&_Accounts.CallOpts)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address account) view returns(bool exists, (address,int256)[] balances, (uint256,uint256,int256)[] positions, (address,uint256)[] signingKeys, uint256 blockNumber)
func (_Accounts *AccountsCaller) GetAccount(opts *bind.CallOpts, account common.Address) (struct {
	Exists      bool
	Balances    []AccountsInterfaceBalance
	Positions   []AccountsInterfacePosition
	SigningKeys []AccountsInterfaceSigningKey
	BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getAccount", account)

	outstruct := new(struct {
		Exists      bool
		Balances    []AccountsInterfaceBalance
		Positions   []AccountsInterfacePosition
		SigningKeys []AccountsInterfaceSigningKey
		BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Balances = *abi.ConvertType(out[1], new([]AccountsInterfaceBalance)).(*[]AccountsInterfaceBalance)
	outstruct.Positions = *abi.ConvertType(out[2], new([]AccountsInterfacePosition)).(*[]AccountsInterfacePosition)
	outstruct.SigningKeys = *abi.ConvertType(out[3], new([]AccountsInterfaceSigningKey)).(*[]AccountsInterfaceSigningKey)
	outstruct.BlockNumber = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address account) view returns(bool exists, (address,int256)[] balances, (uint256,uint256,int256)[] positions, (address,uint256)[] signingKeys, uint256 blockNumber)
func (_Accounts *AccountsSession) GetAccount(account common.Address) (struct {
	Exists      bool
	Balances    []AccountsInterfaceBalance
	Positions   []AccountsInterfacePosition
	SigningKeys []AccountsInterfaceSigningKey
	BlockNumber *big.Int
}, error) {
	return _Accounts.Contract.GetAccount(&_Accounts.CallOpts, account)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address account) view returns(bool exists, (address,int256)[] balances, (uint256,uint256,int256)[] positions, (address,uint256)[] signingKeys, uint256 blockNumber)
func (_Accounts *AccountsCallerSession) GetAccount(account common.Address) (struct {
	Exists      bool
	Balances    []AccountsInterfaceBalance
	Positions   []AccountsInterfacePosition
	SigningKeys []AccountsInterfaceSigningKey
	BlockNumber *big.Int
}, error) {
	return _Accounts.Contract.GetAccount(&_Accounts.CallOpts, account)
}

// GetAccounts is a free data retrieval call binding the contract method 0x8a48ac03.
//
// Solidity: function getAccounts() view returns(address[])
func (_Accounts *AccountsCaller) GetAccounts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getAccounts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAccounts is a free data retrieval call binding the contract method 0x8a48ac03.
//
// Solidity: function getAccounts() view returns(address[])
func (_Accounts *AccountsSession) GetAccounts() ([]common.Address, error) {
	return _Accounts.Contract.GetAccounts(&_Accounts.CallOpts)
}

// GetAccounts is a free data retrieval call binding the contract method 0x8a48ac03.
//
// Solidity: function getAccounts() view returns(address[])
func (_Accounts *AccountsCallerSession) GetAccounts() ([]common.Address, error) {
	return _Accounts.Contract.GetAccounts(&_Accounts.CallOpts)
}

// GetAccountsLength is a free data retrieval call binding the contract method 0x14f326a1.
//
// Solidity: function getAccountsLength() view returns(uint256)
func (_Accounts *AccountsCaller) GetAccountsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getAccountsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountsLength is a free data retrieval call binding the contract method 0x14f326a1.
//
// Solidity: function getAccountsLength() view returns(uint256)
func (_Accounts *AccountsSession) GetAccountsLength() (*big.Int, error) {
	return _Accounts.Contract.GetAccountsLength(&_Accounts.CallOpts)
}

// GetAccountsLength is a free data retrieval call binding the contract method 0x14f326a1.
//
// Solidity: function getAccountsLength() view returns(uint256)
func (_Accounts *AccountsCallerSession) GetAccountsLength() (*big.Int, error) {
	return _Accounts.Contract.GetAccountsLength(&_Accounts.CallOpts)
}

// GetCollateralBalance is a free data retrieval call binding the contract method 0xb7d5820b.
//
// Solidity: function getCollateralBalance(address account, address collateral) view returns(int256)
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
// Solidity: function getCollateralBalance(address account, address collateral) view returns(int256)
func (_Accounts *AccountsSession) GetCollateralBalance(account common.Address, collateral common.Address) (*big.Int, error) {
	return _Accounts.Contract.GetCollateralBalance(&_Accounts.CallOpts, account, collateral)
}

// GetCollateralBalance is a free data retrieval call binding the contract method 0xb7d5820b.
//
// Solidity: function getCollateralBalance(address account, address collateral) view returns(int256)
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
// Solidity: function getFeeAccount() view returns(address)
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
// Solidity: function getFeeAccount() view returns(address)
func (_Accounts *AccountsSession) GetFeeAccount() (common.Address, error) {
	return _Accounts.Contract.GetFeeAccount(&_Accounts.CallOpts)
}

// GetFeeAccount is a free data retrieval call binding the contract method 0x4614be9f.
//
// Solidity: function getFeeAccount() view returns(address)
func (_Accounts *AccountsCallerSession) GetFeeAccount() (common.Address, error) {
	return _Accounts.Contract.GetFeeAccount(&_Accounts.CallOpts)
}

// GetInstrumentEntryPrice is a free data retrieval call binding the contract method 0xea547fb8.
//
// Solidity: function getInstrumentEntryPrice(address account, uint256 instrument) view returns(uint256)
func (_Accounts *AccountsCaller) GetInstrumentEntryPrice(opts *bind.CallOpts, account common.Address, instrument *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Accounts.contract.Call(opts, &out, "getInstrumentEntryPrice", account, instrument)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInstrumentEntryPrice is a free data retrieval call binding the contract method 0xea547fb8.
//
// Solidity: function getInstrumentEntryPrice(address account, uint256 instrument) view returns(uint256)
func (_Accounts *AccountsSession) GetInstrumentEntryPrice(account common.Address, instrument *big.Int) (*big.Int, error) {
	return _Accounts.Contract.GetInstrumentEntryPrice(&_Accounts.CallOpts, account, instrument)
}

// GetInstrumentEntryPrice is a free data retrieval call binding the contract method 0xea547fb8.
//
// Solidity: function getInstrumentEntryPrice(address account, uint256 instrument) view returns(uint256)
func (_Accounts *AccountsCallerSession) GetInstrumentEntryPrice(account common.Address, instrument *big.Int) (*big.Int, error) {
	return _Accounts.Contract.GetInstrumentEntryPrice(&_Accounts.CallOpts, account, instrument)
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
// Solidity: function getInsuranceFund() view returns(address)
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
// Solidity: function getInsuranceFund() view returns(address)
func (_Accounts *AccountsSession) GetInsuranceFund() (common.Address, error) {
	return _Accounts.Contract.GetInsuranceFund(&_Accounts.CallOpts)
}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() view returns(address)
func (_Accounts *AccountsCallerSession) GetInsuranceFund() (common.Address, error) {
	return _Accounts.Contract.GetInsuranceFund(&_Accounts.CallOpts)
}

// GetQuoteBalance is a free data retrieval call binding the contract method 0xaf1aa3b1.
//
// Solidity: function getQuoteBalance(address account) view returns(int256)
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
// Solidity: function getQuoteBalance(address account) view returns(int256)
func (_Accounts *AccountsSession) GetQuoteBalance(account common.Address) (*big.Int, error) {
	return _Accounts.Contract.GetQuoteBalance(&_Accounts.CallOpts, account)
}

// GetQuoteBalance is a free data retrieval call binding the contract method 0xaf1aa3b1.
//
// Solidity: function getQuoteBalance(address account) view returns(int256)
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
// Solidity: function totalBalance(address ) view returns(int256)
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
// Solidity: function totalBalance(address ) view returns(int256)
func (_Accounts *AccountsSession) TotalBalance(arg0 common.Address) (*big.Int, error) {
	return _Accounts.Contract.TotalBalance(&_Accounts.CallOpts, arg0)
}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address ) view returns(int256)
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

// AddAccount is a paid mutator transaction binding the contract method 0xe89b0e1e.
//
// Solidity: function addAccount(address account) returns(bool)
func (_Accounts *AccountsTransactor) AddAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "addAccount", account)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe89b0e1e.
//
// Solidity: function addAccount(address account) returns(bool)
func (_Accounts *AccountsSession) AddAccount(account common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.AddAccount(&_Accounts.TransactOpts, account)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe89b0e1e.
//
// Solidity: function addAccount(address account) returns(bool)
func (_Accounts *AccountsTransactorSession) AddAccount(account common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.AddAccount(&_Accounts.TransactOpts, account)
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

// AddSigningKey is a paid mutator transaction binding the contract method 0xdfa8ea45.
//
// Solidity: function addSigningKey(address account, address signingKey, uint256 expiry) returns(uint8)
func (_Accounts *AccountsTransactor) AddSigningKey(opts *bind.TransactOpts, account common.Address, signingKey common.Address, expiry *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "addSigningKey", account, signingKey, expiry)
}

// AddSigningKey is a paid mutator transaction binding the contract method 0xdfa8ea45.
//
// Solidity: function addSigningKey(address account, address signingKey, uint256 expiry) returns(uint8)
func (_Accounts *AccountsSession) AddSigningKey(account common.Address, signingKey common.Address, expiry *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.AddSigningKey(&_Accounts.TransactOpts, account, signingKey, expiry)
}

// AddSigningKey is a paid mutator transaction binding the contract method 0xdfa8ea45.
//
// Solidity: function addSigningKey(address account, address signingKey, uint256 expiry) returns(uint8)
func (_Accounts *AccountsTransactorSession) AddSigningKey(account common.Address, signingKey common.Address, expiry *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.AddSigningKey(&_Accounts.TransactOpts, account, signingKey, expiry)
}

// Credit is a paid mutator transaction binding the contract method 0x7de182c5.
//
// Solidity: function credit(address account, address collateral, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsTransactor) Credit(opts *bind.TransactOpts, account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "credit", account, collateral, amount)
}

// Credit is a paid mutator transaction binding the contract method 0x7de182c5.
//
// Solidity: function credit(address account, address collateral, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsSession) Credit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Credit(&_Accounts.TransactOpts, account, collateral, amount)
}

// Credit is a paid mutator transaction binding the contract method 0x7de182c5.
//
// Solidity: function credit(address account, address collateral, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsTransactorSession) Credit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Credit(&_Accounts.TransactOpts, account, collateral, amount)
}

// CreditInstrument is a paid mutator transaction binding the contract method 0x8869a3f2.
//
// Solidity: function creditInstrument(address account, uint256 instrument, uint256 amount, uint256 price) returns(int256, int256, uint8)
func (_Accounts *AccountsTransactor) CreditInstrument(opts *bind.TransactOpts, account common.Address, instrument *big.Int, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "creditInstrument", account, instrument, amount, price)
}

// CreditInstrument is a paid mutator transaction binding the contract method 0x8869a3f2.
//
// Solidity: function creditInstrument(address account, uint256 instrument, uint256 amount, uint256 price) returns(int256, int256, uint8)
func (_Accounts *AccountsSession) CreditInstrument(account common.Address, instrument *big.Int, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.CreditInstrument(&_Accounts.TransactOpts, account, instrument, amount, price)
}

// CreditInstrument is a paid mutator transaction binding the contract method 0x8869a3f2.
//
// Solidity: function creditInstrument(address account, uint256 instrument, uint256 amount, uint256 price) returns(int256, int256, uint8)
func (_Accounts *AccountsTransactorSession) CreditInstrument(account common.Address, instrument *big.Int, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.CreditInstrument(&_Accounts.TransactOpts, account, instrument, amount, price)
}

// Debit is a paid mutator transaction binding the contract method 0x262709e6.
//
// Solidity: function debit(address account, address collateral, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsTransactor) Debit(opts *bind.TransactOpts, account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "debit", account, collateral, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x262709e6.
//
// Solidity: function debit(address account, address collateral, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsSession) Debit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Debit(&_Accounts.TransactOpts, account, collateral, amount)
}

// Debit is a paid mutator transaction binding the contract method 0x262709e6.
//
// Solidity: function debit(address account, address collateral, uint256 amount) returns(int256, uint8)
func (_Accounts *AccountsTransactorSession) Debit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Debit(&_Accounts.TransactOpts, account, collateral, amount)
}

// DebitInstrument is a paid mutator transaction binding the contract method 0xac759c71.
//
// Solidity: function debitInstrument(address account, uint256 instrument, uint256 amount, uint256 price) returns(int256, int256, uint8)
func (_Accounts *AccountsTransactor) DebitInstrument(opts *bind.TransactOpts, account common.Address, instrument *big.Int, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "debitInstrument", account, instrument, amount, price)
}

// DebitInstrument is a paid mutator transaction binding the contract method 0xac759c71.
//
// Solidity: function debitInstrument(address account, uint256 instrument, uint256 amount, uint256 price) returns(int256, int256, uint8)
func (_Accounts *AccountsSession) DebitInstrument(account common.Address, instrument *big.Int, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.DebitInstrument(&_Accounts.TransactOpts, account, instrument, amount, price)
}

// DebitInstrument is a paid mutator transaction binding the contract method 0xac759c71.
//
// Solidity: function debitInstrument(address account, uint256 instrument, uint256 amount, uint256 price) returns(int256, int256, uint8)
func (_Accounts *AccountsTransactorSession) DebitInstrument(account common.Address, instrument *big.Int, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.DebitInstrument(&_Accounts.TransactOpts, account, instrument, amount, price)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address account, address collateral, uint256 amount) returns(int256)
func (_Accounts *AccountsTransactor) Deposit(opts *bind.TransactOpts, account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "deposit", account, collateral, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address account, address collateral, uint256 amount) returns(int256)
func (_Accounts *AccountsSession) Deposit(account common.Address, collateral common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Deposit(&_Accounts.TransactOpts, account, collateral, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address account, address collateral, uint256 amount) returns(int256)
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

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newFeeRecipient) returns()
func (_Accounts *AccountsTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "updateFeeRecipient", newFeeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newFeeRecipient) returns()
func (_Accounts *AccountsSession) UpdateFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateFeeRecipient(&_Accounts.TransactOpts, newFeeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newFeeRecipient) returns()
func (_Accounts *AccountsTransactorSession) UpdateFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateFeeRecipient(&_Accounts.TransactOpts, newFeeRecipient)
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
// Solidity: function withdraw(address collateral, address to, uint256 amount, uint256 withdrawAmount, uint256 salt, bytes data, bytes signature) returns(int256)
func (_Accounts *AccountsTransactor) Withdraw(opts *bind.TransactOpts, collateral common.Address, to common.Address, amount *big.Int, withdrawAmount *big.Int, salt *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "withdraw", collateral, to, amount, withdrawAmount, salt, data, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x72d39ed6.
//
// Solidity: function withdraw(address collateral, address to, uint256 amount, uint256 withdrawAmount, uint256 salt, bytes data, bytes signature) returns(int256)
func (_Accounts *AccountsSession) Withdraw(collateral common.Address, to common.Address, amount *big.Int, withdrawAmount *big.Int, salt *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Withdraw(&_Accounts.TransactOpts, collateral, to, amount, withdrawAmount, salt, data, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x72d39ed6.
//
// Solidity: function withdraw(address collateral, address to, uint256 amount, uint256 withdrawAmount, uint256 salt, bytes data, bytes signature) returns(int256)
func (_Accounts *AccountsTransactorSession) Withdraw(collateral common.Address, to common.Address, amount *big.Int, withdrawAmount *big.Int, salt *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Withdraw(&_Accounts.TransactOpts, collateral, to, amount, withdrawAmount, salt, data, signature)
}

// WithdrawInsuranceFund is a paid mutator transaction binding the contract method 0xefd19523.
//
// Solidity: function withdrawInsuranceFund(address collateral, address to, uint256 amount) returns(int256)
func (_Accounts *AccountsTransactor) WithdrawInsuranceFund(opts *bind.TransactOpts, collateral common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "withdrawInsuranceFund", collateral, to, amount)
}

// WithdrawInsuranceFund is a paid mutator transaction binding the contract method 0xefd19523.
//
// Solidity: function withdrawInsuranceFund(address collateral, address to, uint256 amount) returns(int256)
func (_Accounts *AccountsSession) WithdrawInsuranceFund(collateral common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.WithdrawInsuranceFund(&_Accounts.TransactOpts, collateral, to, amount)
}

// WithdrawInsuranceFund is a paid mutator transaction binding the contract method 0xefd19523.
//
// Solidity: function withdrawInsuranceFund(address collateral, address to, uint256 amount) returns(int256)
func (_Accounts *AccountsTransactorSession) WithdrawInsuranceFund(collateral common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.WithdrawInsuranceFund(&_Accounts.TransactOpts, collateral, to, amount)
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

// FilterCredit is a free log retrieval operation binding the contract event 0xc9a1b9e3b11def04d81baa28220fa980997040aef1abec95d1cd2bcfd4ba8a65.
//
// Solidity: event Credit(address indexed account, address indexed collateral, uint256 amount, int256 balance)
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

// WatchCredit is a free log subscription operation binding the contract event 0xc9a1b9e3b11def04d81baa28220fa980997040aef1abec95d1cd2bcfd4ba8a65.
//
// Solidity: event Credit(address indexed account, address indexed collateral, uint256 amount, int256 balance)
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

// ParseCredit is a log parse operation binding the contract event 0xc9a1b9e3b11def04d81baa28220fa980997040aef1abec95d1cd2bcfd4ba8a65.
//
// Solidity: event Credit(address indexed account, address indexed collateral, uint256 amount, int256 balance)
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
	Price      *big.Int
	Position   *big.Int
	EntryPrice *big.Int
	Pnl        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreditInstrument is a free log retrieval operation binding the contract event 0x27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae65363.
//
// Solidity: event CreditInstrument(address indexed account, uint256 indexed instrument, uint256 amount, uint256 price, int256 position, uint256 entryPrice, int256 pnl)
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

// WatchCreditInstrument is a free log subscription operation binding the contract event 0x27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae65363.
//
// Solidity: event CreditInstrument(address indexed account, uint256 indexed instrument, uint256 amount, uint256 price, int256 position, uint256 entryPrice, int256 pnl)
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

// ParseCreditInstrument is a log parse operation binding the contract event 0x27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae65363.
//
// Solidity: event CreditInstrument(address indexed account, uint256 indexed instrument, uint256 amount, uint256 price, int256 position, uint256 entryPrice, int256 pnl)
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

// FilterDebit is a free log retrieval operation binding the contract event 0xcb9aff7bd66be0077d57437ece119a9f166aad37de4491ac4f92fdb2c1ad8d34.
//
// Solidity: event Debit(address indexed account, address indexed collateral, uint256 amount, int256 balance)
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

// WatchDebit is a free log subscription operation binding the contract event 0xcb9aff7bd66be0077d57437ece119a9f166aad37de4491ac4f92fdb2c1ad8d34.
//
// Solidity: event Debit(address indexed account, address indexed collateral, uint256 amount, int256 balance)
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

// ParseDebit is a log parse operation binding the contract event 0xcb9aff7bd66be0077d57437ece119a9f166aad37de4491ac4f92fdb2c1ad8d34.
//
// Solidity: event Debit(address indexed account, address indexed collateral, uint256 amount, int256 balance)
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
	Price      *big.Int
	Position   *big.Int
	EntryPrice *big.Int
	Pnl        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDebitInstrument is a free log retrieval operation binding the contract event 0xc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d1.
//
// Solidity: event DebitInstrument(address indexed account, uint256 indexed instrument, uint256 amount, uint256 price, int256 position, uint256 entryPrice, int256 pnl)
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

// WatchDebitInstrument is a free log subscription operation binding the contract event 0xc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d1.
//
// Solidity: event DebitInstrument(address indexed account, uint256 indexed instrument, uint256 amount, uint256 price, int256 position, uint256 entryPrice, int256 pnl)
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

// ParseDebitInstrument is a log parse operation binding the contract event 0xc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d1.
//
// Solidity: event DebitInstrument(address indexed account, uint256 indexed instrument, uint256 amount, uint256 price, int256 position, uint256 entryPrice, int256 pnl)
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

// FilterDeposit is a free log retrieval operation binding the contract event 0x8cd158929f48bc4a1b49c8c3873d0561de29c5b993767bbe115997adeb1113cf.
//
// Solidity: event Deposit(address indexed account, address indexed collateral, uint256 amount, int256 balance)
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

// WatchDeposit is a free log subscription operation binding the contract event 0x8cd158929f48bc4a1b49c8c3873d0561de29c5b993767bbe115997adeb1113cf.
//
// Solidity: event Deposit(address indexed account, address indexed collateral, uint256 amount, int256 balance)
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

// ParseDeposit is a log parse operation binding the contract event 0x8cd158929f48bc4a1b49c8c3873d0561de29c5b993767bbe115997adeb1113cf.
//
// Solidity: event Deposit(address indexed account, address indexed collateral, uint256 amount, int256 balance)
func (_Accounts *AccountsFilterer) ParseDeposit(log types.Log) (*AccountsDeposit, error) {
	event := new(AccountsDeposit)
	if err := _Accounts.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountsFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the Accounts contract.
type AccountsFeeRecipientUpdatedIterator struct {
	Event *AccountsFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *AccountsFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsFeeRecipientUpdated)
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
		it.Event = new(AccountsFeeRecipientUpdated)
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
func (it *AccountsFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the Accounts contract.
type AccountsFeeRecipientUpdated struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0x7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc2.
//
// Solidity: event FeeRecipientUpdated(address indexed feeRecipient)
func (_Accounts *AccountsFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts, feeRecipient []common.Address) (*AccountsFeeRecipientUpdatedIterator, error) {

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "FeeRecipientUpdated", feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &AccountsFeeRecipientUpdatedIterator{contract: _Accounts.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0x7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc2.
//
// Solidity: event FeeRecipientUpdated(address indexed feeRecipient)
func (_Accounts *AccountsFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *AccountsFeeRecipientUpdated, feeRecipient []common.Address) (event.Subscription, error) {

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "FeeRecipientUpdated", feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsFeeRecipientUpdated)
				if err := _Accounts.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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

// ParseFeeRecipientUpdated is a log parse operation binding the contract event 0x7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc2.
//
// Solidity: event FeeRecipientUpdated(address indexed feeRecipient)
func (_Accounts *AccountsFilterer) ParseFeeRecipientUpdated(log types.Log) (*AccountsFeeRecipientUpdated, error) {
	event := new(AccountsFeeRecipientUpdated)
	if err := _Accounts.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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

// AccountsNewAccountIterator is returned from FilterNewAccount and is used to iterate over the raw logs and unpacked data for NewAccount events raised by the Accounts contract.
type AccountsNewAccountIterator struct {
	Event *AccountsNewAccount // Event containing the contract specifics and raw log

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
func (it *AccountsNewAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsNewAccount)
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
		it.Event = new(AccountsNewAccount)
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
func (it *AccountsNewAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsNewAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsNewAccount represents a NewAccount event raised by the Accounts contract.
type AccountsNewAccount struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewAccount is a free log retrieval operation binding the contract event 0xef4ab4f35cd2027fcc6364f430a86765b6bbd24462cd31f5a6d09bb74241aaf1.
//
// Solidity: event NewAccount(address indexed account)
func (_Accounts *AccountsFilterer) FilterNewAccount(opts *bind.FilterOpts, account []common.Address) (*AccountsNewAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "NewAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsNewAccountIterator{contract: _Accounts.contract, event: "NewAccount", logs: logs, sub: sub}, nil
}

// WatchNewAccount is a free log subscription operation binding the contract event 0xef4ab4f35cd2027fcc6364f430a86765b6bbd24462cd31f5a6d09bb74241aaf1.
//
// Solidity: event NewAccount(address indexed account)
func (_Accounts *AccountsFilterer) WatchNewAccount(opts *bind.WatchOpts, sink chan<- *AccountsNewAccount, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "NewAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsNewAccount)
				if err := _Accounts.contract.UnpackLog(event, "NewAccount", log); err != nil {
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

// ParseNewAccount is a log parse operation binding the contract event 0xef4ab4f35cd2027fcc6364f430a86765b6bbd24462cd31f5a6d09bb74241aaf1.
//
// Solidity: event NewAccount(address indexed account)
func (_Accounts *AccountsFilterer) ParseNewAccount(log types.Log) (*AccountsNewAccount, error) {
	event := new(AccountsNewAccount)
	if err := _Accounts.contract.UnpackLog(event, "NewAccount", log); err != nil {
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

// FilterWithdraw is a free log retrieval operation binding the contract event 0x66efaee033f9c338e1e0c145c8653403e5f05eddfa1e35f67931c6ba3e231ea0.
//
// Solidity: event Withdraw(address indexed account, address indexed collateral, bytes32 indexed withdrawHash, address to, uint256 amount, uint256 withdrawAmount, int256 balance)
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

// WatchWithdraw is a free log subscription operation binding the contract event 0x66efaee033f9c338e1e0c145c8653403e5f05eddfa1e35f67931c6ba3e231ea0.
//
// Solidity: event Withdraw(address indexed account, address indexed collateral, bytes32 indexed withdrawHash, address to, uint256 amount, uint256 withdrawAmount, int256 balance)
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

// ParseWithdraw is a log parse operation binding the contract event 0x66efaee033f9c338e1e0c145c8653403e5f05eddfa1e35f67931c6ba3e231ea0.
//
// Solidity: event Withdraw(address indexed account, address indexed collateral, bytes32 indexed withdrawHash, address to, uint256 amount, uint256 withdrawAmount, int256 balance)
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
