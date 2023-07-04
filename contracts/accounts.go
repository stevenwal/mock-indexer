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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_domain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadRegisterSigningKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadRevokeSigningKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotKeeper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Credit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"}],\"name\":\"CreditInstrument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Debit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"}],\"name\":\"DebitInstrument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"InstrumentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"InstrumentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"KeeperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NewAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"RegisteredSigningKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"revokeHash\",\"type\":\"bytes32\"}],\"name\":\"RevokedSigningKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferHash\",\"type\":\"bytes32\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"VersionInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"WithdrawProxyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawlFeeRecipient\",\"type\":\"address\"}],\"name\":\"WithdrawalFeeRecipientUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountsList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"addInstrument\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"addSigningKey\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorities\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collaterals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"creditInstrument\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"debitInstrument\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Balance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Position[]\",\"name\":\"positions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structAccountsInterface.SigningKey[]\",\"name\":\"signingKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getCollateralBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollaterals\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"getInstrumentEntryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"getInstrumentPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInstruments\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInsuranceFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getQuoteBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getSigningKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"}],\"name\":\"hasSigningKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"instruments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keepers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signingKeySig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountSig\",\"type\":\"bytes\"}],\"name\":\"registerSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"removeCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"removeInstrument\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"resetInstrumentPosition\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"revokeSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signingKeyExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalPositions\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferInsuranceFund\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"fromBalance\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"toBalance\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawProxy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateWithdrawProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWithdrawalFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateWithdrawalFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawInsuranceFund\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawProxies\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101c0604052306080523480156200001657600080fd5b5060405162005a0d38038062005a0d8339810160408190526200003991620004d4565b6001600160a01b038116620000615760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811660a08190526040805163313ce56760e01b8152905163313ce567916004808201926020929091908290030181865afa158015620000ac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000d29190620005aa565b60ff1660c052620000f0838362000364602090811b620035cc17901c565b60e052604051680aed2e8d0c8e4c2ee560bb1b6020820152721859191c995cdcc818dbdb1b185d195c985b0b606a1b60298201526a1859191c995cdcc81d1bcb60aa1b603c8201526e1d5a5b9d0c8d4d88185b5bdd5b9d0b608a1b60478201526c1d5a5b9d0c8d4d881cd85b1d0b609a1b60568201526b62797465733332206461746160a01b6063820152602960f81b606f82015260700160408051808303601f1901815290829052805160209182012061010052670a6d2cedc96caf2560c31b908201526e1859191c995cdcc81858d8dbdd5b9d608a1b6028820152602960f81b603782015260380160408051808303601f1901815290829052805160209182012061012052680a4caced2e6e8cae4560bb1b908201526b1859191c995cdcc81ad95e4b60a21b60298201526d75696e743235362065787069727960901b6035820152602960f81b604382015260440160408051808303601f1901815290829052805160209182012061014052660a4caecded6ca560cb1b908201526a61646472657373206b657960a81b6027820152602960f81b603282015260330160408051808303601f1901815290829052805160209182012061016052680a8e4c2dce6cccae4560bb1b90820152721859191c995cdcc818dbdb1b185d195c985b0b606a1b60298201526a1859191c995cdcc81d1bcb60aa1b603c8201526e1d5a5b9d0c8d4d88185b5bdd5b9d0b608a1b60478201526b1d5a5b9d0c8d4d881cd85b1d60a21b6056820152602960f81b606282015260630160408051601f1981840301815291905280516020909101206101a0525050736cdc77af264be2349f98f5d2cb58a453757caae76101805250620005f4565b6040516c08a92a06e626488dedac2d2dc5609b1b60208201526b1cdd1c9a5b99c81b985b594b60a21b602d8201526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398201526e1d5a5b9d0c8d4d8818da185a5b9259608a1b6048820152602960f81b60578201526000906058016040516020818303038152906040528051906020012083604051602001620003fb9190620005d6565b60408051601f198184030181528282528051602091820120908301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201526080810183905260a00160405160208183030381529060405280519060200120905092915050565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620004a157818101518382015260200162000487565b83811115620004b1576000848401525b50505050565b80516001600160a01b0381168114620004cf57600080fd5b919050565b600080600060608486031215620004ea57600080fd5b83516001600160401b03808211156200050257600080fd5b818601915086601f8301126200051757600080fd5b8151818111156200052c576200052c6200046e565b604051601f8201601f19908116603f011681019083821181831017156200055757620005576200046e565b816040528281528960208487010111156200057157600080fd5b6200058483602083016020880162000484565b809750505050505060208401519150620005a160408501620004b7565b90509250925092565b600060208284031215620005bd57600080fd5b815160ff81168114620005cf57600080fd5b9392505050565b60008251620005ea81846020870162000484565b9190910192915050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a05161530a6200070360003960006137f201526000818161046a01528181611487015281816115170152818161158401528181612e2801528181612ea6015281816130450152818161399301526144a901526000610cdd0152600061172d015260006118cc015260006142a4015260008181610c5a01528181610cbc01528181611772015281816118ab0152818161384d0152614283015260006106010152600081816108d90152818161097401528181612463015281816124dc015261259a01526000818161164d0152818161168d01528181611a3801528181611a780152611b0b015261530a6000f3fe6080604052600436106103765760003560e01c80638340f549116101d1578063cf9b3c8d11610102578063e89b0e1e116100a0578063f0d2d5a81161006f578063f0d2d5a814610c08578063f160d36914610c28578063f698da2514610c48578063fbcbc0f114610c7c57600080fd5b8063e89b0e1e14610b4e578063ea547fb814610b6e578063eeb97d3b14610bb8578063efd1952314610be857600080fd5b8063d658d2e9116100dc578063d658d2e914610abe578063d858348414610aee578063dfa8ea4514610b0e578063e30e838414610b2e57600080fd5b8063cf9b3c8d14610a5e578063d305787714610a7e578063d5c2073014610a9e57600080fd5b80639aeddeff1161016f578063b7d5820b11610149578063b7d5820b146109a4578063b8eb94e9146109ee578063c4d66de814610a1e578063c99d3a0614610a3e57600080fd5b80639aeddeff146108fb578063ac759c711461091b578063af1aa3b11461093b57600080fd5b80638da5cb5b116101ab5780638da5cb5b1461084a5780638e9916061461086a57806391223d6914610897578063999b93af146108c757600080fd5b80638340f549146107e65780638869a3f2146108065780638a48ac031461083557600080fd5b80633fd1e2bd116102ab57806361d21920116102495780636cd22eaf116102235780636cd22eaf146107645780636eacd3981461078457806378b92636146107b15780637de182c5146107c657600080fd5b806361d21920146106bd5780636205ed5e146106ea57806362946d3b1461071757600080fd5b806348e4ccbe1161028557806348e4ccbe1461065557806349ad755b146106755780634f1ef2861461069557806352d1902d146106a857600080fd5b80633fd1e2bd146105ef5780634614be9f1461045b578063469048401461063557600080fd5b806322bbad0b1161031857806327eba4ae116102f257806327eba4ae1461055f578063298dc9d51461057f5780633659cfe61461059f5780633bbd64bc146105bf57600080fd5b806322bbad0b146104c4578063241d400d14610504578063262709e61461053157600080fd5b806313af40351161035457806313af40351461041c57806314f326a11461043c578063158626f71461045b5780631635717c146104a257600080fd5b806306fdde031461037b57806307b1b8b2146103c5578063127217d6146103e7575b600080fd5b34801561038757600080fd5b506103af604051806040016040528060088152602001674163636f756e747360c01b81525081565b6040516103bc9190614987565b60405180910390f35b3480156103d157600080fd5b506103e56103e0366004614a54565b610cad565b005b3480156103f357600080fd5b50610407610402366004614aa2565b6110af565b604080519283526020830191909152016103bc565b34801561042857600080fd5b506103e5610437366004614b14565b6110fa565b34801561044857600080fd5b50600c545b6040519081526020016103bc565b34801561046757600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016103bc565b3480156104ae57600080fd5b506104b761110e565b6040516103bc9190614b2f565b3480156104d057600080fd5b506104f46104df366004614b73565b600b6020526000908152604090205460ff1681565b60405190151581526020016103bc565b34801561051057600080fd5b5061052461051f366004614b14565b611166565b6040516103bc9190614b8c565b34801561053d57600080fd5b5061055161054c366004614bcd565b6111df565b6040516103bc929190614c2b565b34801561056b57600080fd5b506103e561057a366004614b14565b611335565b34801561058b57600080fd5b5061040761059a366004614bcd565b6113ae565b3480156105ab57600080fd5b506103e56105ba366004614b14565b611643565b3480156105cb57600080fd5b506104f46105da366004614b14565b60046020526000908152604090205460ff1681565b3480156105fb57600080fd5b506106237f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016103bc565b34801561064157600080fd5b5060055461048a906001600160a01b031681565b34801561066157600080fd5b506103e5610670366004614c3f565b61171f565b34801561068157600080fd5b5061044d610690366004614cbd565b611a00565b6103e56106a3366004614a54565b611a2e565b3480156106b457600080fd5b5061044d611afe565b3480156106c957600080fd5b5061044d6106d8366004614b73565b60106020526000908152604090205481565b3480156106f657600080fd5b5061044d610705366004614b14565b600e6020526000908152604090205481565b34801561072357600080fd5b506104f4610732366004614ce7565b6001600160a01b039182166000908152600d602090815260408083209390941682526002909201909152205460ff1690565b34801561077057600080fd5b506103e561077f366004614d1a565b611bb1565b34801561079057600080fd5b5061044d61079f366004614b14565b600f6020526000908152604090205481565b3480156107bd57600080fd5b50610524611c34565b3480156107d257600080fd5b506105516107e1366004614bcd565b611c95565b3480156107f257600080fd5b5061044d610801366004614bcd565b611dd1565b34801561081257600080fd5b50610826610821366004614d56565b611e18565b6040516103bc93929190614d8f565b34801561084157600080fd5b50610524611f7c565b34801561085657600080fd5b5060005461048a906001600160a01b031681565b34801561087657600080fd5b5061088a610885366004614b73565b611fdc565b6040516103bc9190614daa565b3480156108a357600080fd5b506104f46108b2366004614b14565b60036020526000908152604090205460ff1681565b3480156108d357600080fd5b5061048a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561090757600080fd5b5061088a610916366004614cbd565b6120c4565b34801561092757600080fd5b50610826610936366004614d56565b612283565b34801561094757600080fd5b5061044d610956366004614b14565b6001600160a01b039081166000908152600d602090815260408083207f0000000000000000000000000000000000000000000000000000000000000000909416835260039093019052205490565b3480156109b057600080fd5b5061044d6109bf366004614ce7565b6001600160a01b039182166000908152600d602090815260408083209390941682526003909201909152205490565b3480156109fa57600080fd5b506104f4610a09366004614b14565b60066020526000908152604090205460ff1681565b348015610a2a57600080fd5b506103e5610a39366004614b14565b6123cd565b348015610a4a57600080fd5b506103e5610a59366004614b14565b612557565b348015610a6a57600080fd5b5061048a610a79366004614b73565b61279c565b348015610a8a57600080fd5b506103e5610a99366004614d1a565b6127c6565b348015610aaa57600080fd5b506103e5610ab9366004614d1a565b612849565b348015610aca57600080fd5b506104f4610ad9366004614b73565b60076020526000908152604090205460ff1681565b348015610afa57600080fd5b5061044d610b09366004614db8565b6128cc565b348015610b1a57600080fd5b5061088a610b29366004614bcd565b6129c4565b348015610b3a57600080fd5b5061088a610b49366004614b73565b612b26565b348015610b5a57600080fd5b506104f4610b69366004614b14565b612cd6565b348015610b7a57600080fd5b5061044d610b89366004614cbd565b6001600160a01b03919091166000908152600d6020908152604080832093835260049093019052206001015490565b348015610bc457600080fd5b506104f4610bd3366004614b14565b60096020526000908152604090205460ff1681565b348015610bf457600080fd5b5061044d610c03366004614bcd565b612d17565b348015610c1457600080fd5b506103e5610c23366004614b14565b6130a2565b348015610c3457600080fd5b506103e5610c43366004614b14565b61319a565b348015610c5457600080fd5b5061044d7f000000000000000000000000000000000000000000000000000000000000000081565b348015610c8857600080fd5b50610c9c610c97366004614b14565b613213565b6040516103bc959493929190614ef9565b610cb56136d4565b6000610d567f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000085604051602001610d209291909182526001600160a01b0316602082015260400190565b6040516020818303038152906040528051906020012061190160f01b600090815260029290925260229081526042822091905290565b3360009081526004602052604090205490915060ff16610d9a578060016040516001622b6d3960e21b03198152600401610d91929190614c2b565b60405180910390fd5b6001600160a01b038316610dc9578060036040516001622b6d3960e21b03198152600401610d91929190614c2b565b6001600160a01b0383166000908152600e60205260408120549003610e095780600a6040516001622b6d3960e21b03198152600401610d91929190614c2b565b6000610e1582846136fd565b90506001600160a01b038116610e46578160076040516001622b6d3960e21b03198152600401610d91929190614c2b565b6001600160a01b038082166000908152600d60209081526040808320938816835260029093019052205460ff16610e985781600a6040516001622b6d3960e21b03198152600401610d91929190614c2b565b6001600160a01b0381166000908152600d6020908152604080832060010180548251818502810185019093528083528493830182828015610f0257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610ee4575b505050505090505b8051821015610f5157856001600160a01b0316818381518110610f2f57610f2f614f8d565b60200260200101516001600160a01b03160315610f5157816001019150610f0a565b80518210610f7a5783600a6040516001622b6d3960e21b03198152600401610d91929190614c2b565b8060018251610f899190614fb9565b81518110610f9957610f99614f8d565b6020026020010151600d6000856001600160a01b03166001600160a01b031681526020019081526020016000206001018381548110610fda57610fda614f8d565b600091825260208083209190910180546001600160a01b0319166001600160a01b039485161790559185168152600d9091526040902060010180548061102257611022614fd0565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03858116808452600d83526040808520928b168086526002909301909352828420805460ff19169055915187939192917ff0d72ac27c7920d0a69f77bbf64f8509a0307c49e54eee377861c46101ef25cb91a45050600160025550505050565b6000806110ba6136d4565b60008060006110cc8a8a8a8a8a6137ec565b9250925092506110db83613c2f565b506110e589613c2f565b50600160025590999098509650505050505050565b611102613cf8565b61110b81613d24565b50565b6060600a80548060200260200160405190810160405280929190818152602001828054801561115c57602002820191906000526020600020905b815481526020019060010190808311611148575b5050505050905090565b6001600160a01b0381166000908152600d60209081526040918290206001018054835181840281018401909452808452606093928301828280156111d357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111b5575b50505050509050919050565b33600090815260036020526040812054819060ff1661121157604051631890934360e11b815260040160405180910390fd5b6112196136d4565b6001600160a01b03841660009081526009602052604090205460ff166112455750600090506004611326565b61124e85613c2f565b506001600160a01b038086166000908152600d6020908152604080832093881683526003909301905290812054611286908590614fe6565b6001600160a01b038088166000908152600d60209081526040808320938a16835260039093018152828220849055600f9052908120805492935086929091906112d0908490614fe6565b909155505060408051858152602081018390526001600160a01b0380881692908916917fcb9aff7bd66be0077d57437ece119a9f166aad37de4491ac4f92fdb2c1ad8d3491015b60405180910390a39150600090505b60016002559094909350915050565b61133d613cf8565b6001600160a01b0381166113645760405163d92e233d60e01b815260040160405180910390fd5b601180546001600160a01b0319166001600160a01b0383169081179091556040517fcbc6a0c385161960535b52c682db3a205870921d9a6cefbd50ada57faadb697690600090a250565b6000806113b96136d4565b6113c1613cf8565b826000036113f5576000805160206152de8339815191526002604051630149876960e41b8152600401610d91929190614c2b565b6001600160a01b03841661142f576000805160206152de8339815191526003604051630149876960e41b8152600401610d91929190614c2b565b6001600160a01b03851660009081526009602052604090205460ff1661147b576000805160206152de8339815191526004604051630149876960e41b8152600401610d91929190614c2b565b50506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166000908152600d602081815260408084208886168086526003918201845282862054968916865293835281852093855292909201905290205482821215611515576000805160206152de8339815191526008604051630149876960e41b8152600401610d91929190614c2b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316846001600160a01b03160361157a576000805160206152de833981519152600c604051630149876960e41b8152600401610d91929190614c2b565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166000818152600d602081815260408084208b871680865260039182018452828620998b9003998a9055968b16808652938352818520878652018252928390209588019586905582518881526000805160206152de833981519152918101919091529092917fecf06959d9b6442f0e285add9df1ac6dd0730c4c7b7c33aa8bff6583cb3090e3910160405180910390a460016002559094909350915050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361168b5760405162461bcd60e51b8152600401610d9190615025565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166116d4600080516020615297833981519152546001600160a01b031690565b6001600160a01b0316146116fa5760405162461bcd60e51b8152600401610d9190615071565b61170381613d93565b6040805160008082526020820190925261110b91839190613d9b565b6117276136d4565b604080517f000000000000000000000000000000000000000000000000000000000000000060208201526001600160a01b038616918101919091526060810184905260009061179a907f000000000000000000000000000000000000000000000000000000000000000090608001610d20565b3360009081526004602052604090205490915060ff166117d257806001604051634e7391e160e11b8152600401610d91929190614c2b565b6001600160a01b0385166117fe57806003604051634e7391e160e11b8152600401610d91929190614c2b565b8360000361182457806002604051634e7391e160e11b8152600401610d91929190614c2b565b6001600160a01b0385166000908152600e6020526040902054156118605780600a604051634e7391e160e11b8152600401610d91929190614c2b565b600061186c82846136fd565b90506001600160a01b03811661189a57816007604051634e7391e160e11b8152600401610d91929190614c2b565b6118a381613c2f565b50600061190f7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000084604051602001610d209291909182526001600160a01b0316602082015260400190565b9050600061191d82876136fd565b9050876001600160a01b0316816001600160a01b0316146119565783600b604051634e7391e160e11b8152600401610d91929190614c2b565b6001600160a01b038881166000818152600e602090815260408083208c9055938716808352600d82528483206001808201805480830182559086528486200180546001600160a01b031916871790558585526002909101835292859020805460ff191690931790925592518a815287937f1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac910160405180910390a450506001600255505050505050565b6001600160a01b0382166000908152600d602090815260408083208484526004019091529020545b92915050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003611a765760405162461bcd60e51b8152600401610d9190615025565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611abf600080516020615297833981519152546001600160a01b031690565b6001600160a01b031614611ae55760405162461bcd60e51b8152600401610d9190615071565b611aee82613d93565b611afa82826001613d9b565b5050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611b9e5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610d91565b5060008051602061529783398151915290565b611bb9613cf8565b6001600160a01b038216611be05760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff191685151590811790915590519092917fc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc91a35050565b6060600880548060200260200160405190810160405280929190818152602001828054801561115c57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611c6e575050505050905090565b33600090815260036020526040812054819060ff16611cc757604051631890934360e11b815260040160405180910390fd5b611ccf6136d4565b6001600160a01b03841660009081526009602052604090205460ff16611cfb5750600090506004611326565b611d0485613c2f565b506001600160a01b038086166000908152600d6020908152604080832093881683526003909301905290812054611d3c9085906150bd565b6001600160a01b038088166000908152600d60209081526040808320938a16835260039093018152828220849055600f905290812080549293508692909190611d869084906150bd565b909155505060408051858152602081018390526001600160a01b0380881692908916917fc9a1b9e3b11def04d81baa28220fa980997040aef1abec95d1cd2bcfd4ba8a659101611317565b6000611ddb6136d4565b611de484613c2f565b506000611df2858585613f0b565b9050611e096001600160a01b038516333086614067565b90505b60016002559392505050565b336000908152600360205260408120548190819060ff16611e4c57604051631890934360e11b815260040160405180910390fd5b611e546136d4565b6000868152600b602052604090205460ff16611e795750600091508190506005611f6a565b6001600160ff1b03851115611e975750600091508190506002611f6a565b611ea087613c2f565b506001600160a01b0387166000908152600d60209081526040808320898452600401909152812081908190611ed69089896140f1565b60008c815260106020526040812080549497509295509093508a92611efc9084906150bd565b90915550506040805189815260208101899052908101849052606081018390526080810182905289906001600160a01b038c16907f27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae653639060a0015b60405180910390a391945090925060009150505b60016002819055509450945094915050565b6060600c80548060200260200160405190810160405280929190818152602001828054801561115c576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611c6e575050505050905090565b3360009081526003602052604081205460ff1661200c57604051631890934360e11b815260040160405180910390fd5b6120146136d4565b81600003612024575060026120ba565b6000828152600b602052604090205460ff1615612043575060056120ba565b600a805460018181019092557fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8018390556000838152600b6020526040808220805460ff1916909317909255905183917fb8597a358785eb73ad0a81079eb5e50ba0ac43dd023f47c0e0ff47e58a9495c891a25060005b6001600255919050565b3360009081526003602052604081205460ff166120f457604051631890934360e11b815260040160405180910390fd5b6120fc6136d4565b6000828152600b602052604090205460ff1661211a57506005612278565b61212383613c2f565b506001600160a01b0383166000908152600d60209081526040808320858452600401909152812080548282556001909101829055908113156121e2576000838152601060205260408120805483929061217d908490614fe6565b909155505060408051828152600060208201819052818301819052606082018190526080820152905184916001600160a01b038716917fc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d19181900360a00190a3612272565b60008112156122725760008381526010602052604081208054839290612209908490614fe6565b909155508390506001600160a01b0385167f27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae65363612244846150fe565b6040805191825260006020830181905290820181905260608201819052608082015260a00160405180910390a35b60009150505b600160025592915050565b336000908152600360205260408120548190819060ff166122b757604051631890934360e11b815260040160405180910390fd5b6122bf6136d4565b6000868152600b602052604090205460ff166122e45750600091508190506005611f6a565b6001600160ff1b038511156123025750600091508190506002611f6a565b61230b87613c2f565b506000808061234961231c896150fe565b6001600160a01b038c166000908152600d602090815260408083208e8452600401909152902090896140f1565b60008c815260106020526040812080549497509295509093508a9261236f908490614fe6565b90915550506040805189815260208101899052908101849052606081018390526080810182905289906001600160a01b038c16907fc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d19060a001611f56565b60016123d98180614fb9565b600154146123fa576040516302ed543d60e51b815260040160405180910390fd5b806001036124085760016002555b61241182613d24565b600580546001600160a01b038085166001600160a01b031992831681179093556008805460018181019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180547f0000000000000000000000000000000000000000000000000000000000000000909316929093168217909255600090815260096020526040808220805460ff191690931790925590517f7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc29190a26040516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016907f7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f90600090a2600181905560405181907f7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d90600090a25050565b61255f613cf8565b6001600160a01b03811660009081526009602052604090205460ff16612598576040516368f7a67560e11b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316036125ea576040516368f7a67560e11b815260040160405180910390fd5b600080600880548060200260200160405190810160405280929190818152602001828054801561264357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612625575b505050505090505b805182101561269257826001600160a01b031681838151811061267057612670614f8d565b60200260200101516001600160a01b031603156126925781600101915061264b565b805182106126b3576040516368f7a67560e11b815260040160405180910390fd5b80600182516126c29190614fb9565b815181106126d2576126d2614f8d565b6020026020010151600883815481106126ed576126ed614f8d565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600880548061272c5761272c614fd0565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03851680835260099091526040808320805460ff191690555190917fd89d2ee68ab04dca0193f48a4aff55e20fa5ec0429a8a8c1c51b8dad6178a59391a2505050565b600c81815481106127ac57600080fd5b6000918252602090912001546001600160a01b0316905081565b6127ce613cf8565b6001600160a01b0382166127f55760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260046020526040808220805460ff191685151590811790915590519092917f786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c0091a35050565b612851613cf8565b6001600160a01b0382166128785760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260066020526040808220805460ff191685151590811790915590519092917f097359b7e2dca2adc06bfc6c3e4e8fcc69a2e916105499e8e4c60a569598cbc091a35050565b60006128d66136d4565b6000806128e98b8b8b8b8b8b8b8b614279565b915091506128f682613c2f565b5061290b6001600160a01b038c168b8a614523565b6001600160a01b038a1660009081526006602052604090205460ff161561299157604051631e6fd1a760e31b81526001600160a01b038b169063f37e8d389061295e9085908f908d908b9060040161511a565b600060405180830381600087803b15801561297857600080fd5b505af115801561298c573d6000803e3d6000fd5b505050505b85156129b1576011546129b1906001600160a01b038d8116911688614523565b60016002559a9950505050505050505050565b3360009081526003602052604081205460ff166129f457604051631890934360e11b815260040160405180910390fd5b6129fc6136d4565b6001600160a01b0383166000908152600e602052604090205415612a225750600b611e0c565b6001600160a01b0383161580612a36575081155b15612a4357506002611e0c565b6001600160a01b038085166000908152600d60209081526040808320938716835260029093019052205460ff1615612a7d57506007611e0c565b6001600160a01b038381166000818152600e60209081526040808320879055938816808352600d82528483206001808201805480830182559086528486200180546001600160a01b0319168717905585855260029091018352858420805460ff1916909117905593518681529193917f1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac910160405180910390a450600060016002559392505050565b3360009081526003602052604081205460ff16612b5657604051631890934360e11b815260040160405180910390fd5b612b5e6136d4565b6000828152600b602052604090205460ff16612b7c575060056120ba565b600080600a805480602002602001604051908101604052809291908181526020018280548015612bcb57602002820191906000526020600020905b815481526020019060010190808311612bb7575b505050505090505b8051821015612c085783818381518110612bef57612bef614f8d565b60200260200101510315612c0857816001019150612bd3565b80518210612c1b576005925050506120ba565b8060018251612c2a9190614fb9565b81518110612c3a57612c3a614f8d565b6020026020010151600a8381548110612c5557612c55614f8d565b600091825260209091200155600a805480612c7257612c72614fd0565b600082815260208082208301600019908101839055909201909255858252600b90526040808220805460ff191690555185917f08b9963dddd96e14a775b44c82852b5a482f5060ae86706c04fc5878c3fe3d8d91a250506001600255506000919050565b3360009081526003602052604081205460ff16612d0657604051631890934360e11b815260040160405180910390fd5b612d0e6136d4565b61227882613c2f565b6000612d216136d4565b3360009081526004602052604090205460ff16612d64576000805160206152de833981519152600160405163c802b7f360e01b8152600401610d91929190614c2b565b81600003612d98576000805160206152de833981519152600260405163c802b7f360e01b8152600401610d91929190614c2b565b6001600160a01b038316612dd2576000805160206152de833981519152600360405163c802b7f360e01b8152600401610d91929190614c2b565b6001600160a01b03841660009081526009602052604090205460ff16612e1e576000805160206152de833981519152600460405163c802b7f360e01b8152600401610d91929190614c2b565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166000908152600d60209081526040808320938816835260039093019052205482811215612e9c576000805160206152de833981519152600860405163c802b7f360e01b8152600401610d91929190614c2b565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166000908152600d602090815260408083209389168352600390930181528282209386900393849055600f90529081208054859290612f06908490614fe6565b90915550612f2090506001600160a01b0386168585614523565b6001600160a01b03841660009081526006602052604090205460ff1615612fc257600554604051631e6fd1a760e31b81526001600160a01b03918216600482015286821660248201526044810185905260806064820152600060848201529085169063f37e8d389060a401600060405180830381600087803b158015612fa557600080fd5b505af1158015612fb9573d6000803e3d6000fd5b50505050613003565b6005546001600160a01b03858116911614613003576000805160206152de833981519152600c60405163c802b7f360e01b8152600401610d91929190614c2b565b604080516001600160a01b0386811682526020820186905291810185905260006060820152608081018390526000805160206152de83398151915291808816917f0000000000000000000000000000000000000000000000000000000000000000909116907f9eaa8b17b9a084ea901c6a072e094ed435b7a049ca673db6b17e98203886944e9060a00160405180910390a46001600255949350505050565b6130aa613cf8565b6001600160a01b0381166130d15760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811660009081526009602052604090205460ff161561310b576040516368f7a67560e11b815260040160405180910390fd5b6008805460018082019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180546001600160a01b0319166001600160a01b038416908117909155600081815260096020526040808220805460ff1916909417909355915190917f7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f91a250565b6131a2613cf8565b6001600160a01b0381166131c95760405163d92e233d60e01b815260040160405180910390fd5b600580546001600160a01b0319166001600160a01b0383169081179091556040517f7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc290600090a250565b6001600160a01b0381166000908152600d602052604081205460085460ff90911691606091829182918067ffffffffffffffff811115613255576132556149b1565b60405190808252806020026020018201604052801561329a57816020015b60408051808201909152600080825260208201528152602001906001900390816132735790505b50945060005b8181101561332b576000600882815481106132bd576132bd614f8d565b60009182526020808320909101546040805180820182526001600160a01b03928316808252928e168552600d8452818520838652600301845293205491830191909152885190925088908490811061331757613317614f8d565b6020908102919091010152506001016132a0565b50600a548067ffffffffffffffff811115613348576133486149b1565b60405190808252806020026020018201604052801561339d57816020015b61338a60405180606001604052806000815260200160008152602001600081525090565b8152602001906001900390816133665790505b50945060005b8181101561343f576000600a82815481106133c0576133c0614f8d565b6000918252602080832090910154604080516060810182528281526001600160a01b038f168552600d84528185208386526004018085528286206001810154838701529584905290935292549282019290925288519192509088908490811061342b5761342b614f8d565b6020908102919091010152506001016133a3565b506001600160a01b0388166000908152600d60209081526040808320600101805482518185028101850190935280835291929091908301828280156134ad57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161348f575b50505050509050805167ffffffffffffffff8111156134ce576134ce6149b1565b60405190808252806020026020018201604052801561351357816020015b60408051808201909152600080825260208201528152602001906001900390816134ec5790505b50945060005b81518110156135bc57604051806040016040528083838151811061353f5761353f614f8d565b60200260200101516001600160a01b03168152602001600e600085858151811061356b5761356b614f8d565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020548152508682815181106135a9576135a9614f8d565b6020908102919091010152600101613519565b5043935050505091939590929450565b6040516c08a92a06e626488dedac2d2dc5609b1b60208201526b1cdd1c9a5b99c81b985b594b60a21b602d8201526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398201526e1d5a5b9d0c8d4d8818da185a5b9259608a1b6048820152602960f81b60578201526000906058016040516020818303038152906040528051906020012083604051602001613661919061514d565b60408051601f198184030181528282528051602091820120908301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201526080810183905260a00160405160208183030381529060405280519060200120905092915050565b6002546001146136f75760405163558a1e0360e11b815260040160405180910390fd5b60028055565b600080600080845160400361372f57505050602082015160408301516001600160ff1b0381169060ff1c601b01613784565b84516041036137785750505060208201516040830151606084015160001a601b811480159061376257508060ff16601c14155b156137735760009350505050611a28565b613784565b60009350505050611a28565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa1580156137d7573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b604080517f000000000000000000000000000000000000000000000000000000000000000060208201526001600160a01b038088169282019290925290851660608201526080810184905260a08101839052600090819081908190613875907f00000000000000000000000000000000000000000000000000000000000000009060c001610d20565b3360009081526004602052604090205490915060ff166138ad57806001604051630149876960e41b8152600401610d91929190614c2b565b866000036138d357806002604051630149876960e41b8152600401610d91929190614c2b565b6001600160a01b0388166138ff57806003604051630149876960e41b8152600401610d91929190614c2b565b6001600160a01b03891660009081526009602052604090205460ff1661393d57806004604051630149876960e41b8152600401610d91929190614c2b565b60008181526007602052604090205460ff161561397257806006604051630149876960e41b8152600401610d91929190614c2b565b600061397e82876136fd565b90506001600160a01b03811615806139c757507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316145b156139ea57816007604051630149876960e41b8152600401610d91929190614c2b565b886001600160a01b0316816001600160a01b031603613a215781600c604051630149876960e41b8152600401610d91929190614c2b565b6000600d6000836001600160a01b03166001600160a01b0316815260200190815260200160002060030160008c6001600160a01b03166001600160a01b031681526020019081526020016000205490506000600d60008c6001600160a01b03166001600160a01b0316815260200190815260200160002060030160008d6001600160a01b03166001600160a01b0316815260200190815260200160002054905089821215613ae757836008604051630149876960e41b8152600401610d91929190614c2b565b8982039150898101905081600d6000856001600160a01b03166001600160a01b0316815260200190815260200160002060030160008e6001600160a01b03166001600160a01b031681526020019081526020016000208190555080600d60008d6001600160a01b03166001600160a01b0316815260200190815260200160002060030160008e6001600160a01b03166001600160a01b031681526020019081526020016000208190555060016007600086815260200190815260200160002060006101000a81548160ff0219169083151502179055508b6001600160a01b03168b6001600160a01b0316846001600160a01b03167fecf06959d9b6442f0e285add9df1ac6dd0730c4c7b7c33aa8bff6583cb3090e38d88604051613c15929190918252602082015260400190565b60405180910390a4919b909a509098509650505050505050565b60006001600160a01b03821615613cf3576001600160a01b0382166000908152600d602052604090205460ff1615613c6957506001919050565b6001600160a01b0382166000818152600d6020526040808220805460ff19166001908117909155600c8054918201815583527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c70180546001600160a01b03191684179055517fef4ab4f35cd2027fcc6364f430a86765b6bbd24462cd31f5a6d09bb74241aaf19190a25b919050565b6000546001600160a01b03163314613d22576040516282b42960e81b815260040160405180910390fd5b565b6001600160a01b038116613d4b5760405163d92e233d60e01b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b038316908117825560405190917f4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b91a250565b61110b613cf8565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615613dd357613dce836145a1565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613e2d575060408051601f3d908101601f19168201909252613e2a91810190615169565b60015b613e905760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610d91565b6000805160206152978339815191528114613eff5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610d91565b50613dce83838361463d565b60006001600160a01b038416613f345760405163d92e233d60e01b815260040160405180910390fd5b81600003613f5557604051631f2a200560e01b815260040160405180910390fd5b6001600160a01b03831660009081526009602052604090205460ff16613f8e576040516368f7a67560e11b815260040160405180910390fd5b6001600160a01b038085166000908152600d6020908152604080832093871683526003909301905290812054613fc59084906150bd565b6001600160a01b038087166000908152600d60209081526040808320938916835260039093018152828220849055600f90529081208054929350859290919061400f9084906150bd565b909155505060408051848152602081018390526001600160a01b0380871692908816917f8cd158929f48bc4a1b49c8c3873d0561de29c5b993767bbe115997adeb1113cf910160405180910390a390505b9392505050565b60006040516323b872dd60e01b81528460048201528360248201528260448201526020600060648360008a5af13d15601f3d11600160005114161716915050806140ea5760405162461bcd60e51b81526020600482015260146024820152731514905394d1915497d19493d357d1905253115160621b6044820152606401610d91565b5050505050565b600080600084600003614111575050835460018501549091506000614270565b855461411d86826150bd565b935080600003614136576001870185905584925061426b565b8360000361415b5761414d81886001015487614662565b60006001890155915061426b565b60008112801561416b5750600084135b8061418157506000811380156141815750600084125b156141a65761419581886001015487614662565b60018801869055859350915061426b565b6000811380156141b557508084135b156141f657836141c58688615182565b60018901546141d49084615182565b6141de91906151a1565b6141e891906151cf565b60018801819055925061426b565b60008112801561420557508084125b1561423f57614213846150fe565b8561421d886150fe565b6142279190615182565b6001890154614235846150fe565b6141d49190615182565b60018701549250620f42406142548685614fe6565b61425e90886151e3565b6142689190615268565b91505b508286555b93509350939050565b60008060006143137f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008d8d8d8c8b80519060200120604051602001610d20969594939291909586526001600160a01b0394851660208701529290931660408501526060840152608083019190915260a082015260c00190565b3360009081526004602052604090205490915060ff1661434b5780600160405163c802b7f360e01b8152600401610d91929190614c2b565b881580614356575087155b156143795780600260405163c802b7f360e01b8152600401610d91929190614c2b565b851580159061439157506011546001600160a01b0316155b156143b45780600d60405163c802b7f360e01b8152600401610d91929190614c2b565b6001600160a01b038a166143e05780600360405163c802b7f360e01b8152600401610d91929190614c2b565b6001600160a01b038b1660009081526009602052604090205460ff1661441e5780600460405163c802b7f360e01b8152600401610d91929190614c2b565b60008181526007602052604090205460ff16156144535780600660405163c802b7f360e01b8152600401610d91929190614c2b565b87891415801561446557506000198914155b156144885780600960405163c802b7f360e01b8152600401610d91929190614c2b565b600061449482866136fd565b90506001600160a01b03811615806144dd57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316145b156145005781600760405163c802b7f360e01b8152600401610d91929190614c2b565b80614510828e858f8f8f8e61468e565b9350935050509850989650505050505050565b600060405163a9059cbb60e01b8152836004820152826024820152602060006044836000895af13d15601f3d116001600051141617169150508061459b5760405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b6044820152606401610d91565b50505050565b6001600160a01b0381163b61460e5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610d91565b60008051602061529783398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b614646836147b9565b6000825111806146535750805b15613dce5761459b83836147f9565b6000620f42406146728484614fe6565b61467c90866151e3565b6146869190615268565b949350505050565b6001600160a01b038087166000818152600f6020908152604080832054948c168352600d8252808320938352600390930190529081205491906146d184866151a1565b9050808212806146e057508083125b156147035787600860405163c802b7f360e01b8152600401610d91929190614c2b565b6001600160a01b038981166000818152600f6020908152604080832096869003968790558e8516808452600d8352818420858552600301835281842098879003988990558d84526007835292819020805460ff191660011790558051948c1685529084018a9052830188905260608301879052608083018690528a927f9eaa8b17b9a084ea901c6a072e094ed435b7a049ca673db6b17e98203886944e9060a00160405180910390a45050979650505050505050565b6147c2816145a1565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b606061406083836040518060600160405280602781526020016152b76027913960606001600160a01b0384163b6148815760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610d91565b600080856001600160a01b03168560405161489c919061514d565b600060405180830381855af49150503d80600081146148d7576040519150601f19603f3d011682016040523d82523d6000602084013e6148dc565b606091505b50915091506148ec8282866148f6565b9695505050505050565b60608315614905575081614060565b8251156149155782518084602001fd5b8160405162461bcd60e51b8152600401610d919190614987565b60005b8381101561494a578181015183820152602001614932565b8381111561459b5750506000910152565b6000815180845261497381602086016020860161492f565b601f01601f19169290920160200192915050565b602081526000614060602083018461495b565b80356001600160a01b0381168114613cf357600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f8301126149d857600080fd5b813567ffffffffffffffff808211156149f3576149f36149b1565b604051601f8301601f19908116603f01168101908282118183101715614a1b57614a1b6149b1565b81604052838152866020858801011115614a3457600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215614a6757600080fd5b614a708361499a565b9150602083013567ffffffffffffffff811115614a8c57600080fd5b614a98858286016149c7565b9150509250929050565b600080600080600060a08688031215614aba57600080fd5b614ac38661499a565b9450614ad16020870161499a565b93506040860135925060608601359150608086013567ffffffffffffffff811115614afb57600080fd5b614b07888289016149c7565b9150509295509295909350565b600060208284031215614b2657600080fd5b6140608261499a565b6020808252825182820181905260009190848201906040850190845b81811015614b6757835183529284019291840191600101614b4b565b50909695505050505050565b600060208284031215614b8557600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015614b675783516001600160a01b031683529284019291840191600101614ba8565b600080600060608486031215614be257600080fd5b614beb8461499a565b9250614bf96020850161499a565b9150604084013590509250925092565b600e8110614c2757634e487b7160e01b600052602160045260246000fd5b9052565b828152604081016140606020830184614c09565b60008060008060808587031215614c5557600080fd5b614c5e8561499a565b935060208501359250604085013567ffffffffffffffff80821115614c8257600080fd5b614c8e888389016149c7565b93506060870135915080821115614ca457600080fd5b50614cb1878288016149c7565b91505092959194509250565b60008060408385031215614cd057600080fd5b614cd98361499a565b946020939093013593505050565b60008060408385031215614cfa57600080fd5b614d038361499a565b9150614d116020840161499a565b90509250929050565b60008060408385031215614d2d57600080fd5b614d368361499a565b915060208301358015158114614d4b57600080fd5b809150509250929050565b60008060008060808587031215614d6c57600080fd5b614d758561499a565b966020860135965060408601359560600135945092505050565b83815260208101839052606081016146866040830184614c09565b60208101611a288284614c09565b600080600080600080600080610100898b031215614dd557600080fd5b614dde8961499a565b9750614dec60208a0161499a565b965060408901359550606089013594506080890135935060a0890135925060c089013567ffffffffffffffff80821115614e2557600080fd5b614e318c838d016149c7565b935060e08b0135915080821115614e4757600080fd5b50614e548b828c016149c7565b9150509295985092959890939650565b600081518084526020808501945080840160005b83811015614eaa5781518051885283810151848901526040908101519088015260609096019590820190600101614e78565b509495945050505050565b600081518084526020808501945080840160005b83811015614eaa57815180516001600160a01b031688528301518388015260409096019590820190600101614ec9565b600060a082018715158352602060a08185015281885180845260c086019150828a01935060005b81811015614f5057845180516001600160a01b031684528401518484015293830193604090920191600101614f20565b50508481036040860152614f648189614e64565b925050508281036060840152614f7a8186614eb5565b9150508260808301529695505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082821015614fcb57614fcb614fa3565b500390565b634e487b7160e01b600052603160045260246000fd5b60008083128015600160ff1b85018412161561500457615004614fa3565b6001600160ff1b038401831381161561501f5761501f614fa3565b50500390565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b600080821280156001600160ff1b03849003851316156150df576150df614fa3565b600160ff1b83900384128116156150f8576150f8614fa3565b50500190565b6000600160ff1b820161511357615113614fa3565b5060000390565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906148ec9083018461495b565b6000825161515f81846020870161492f565b9190910192915050565b60006020828403121561517b57600080fd5b5051919050565b600081600019048311821515161561519c5761519c614fa3565b500290565b600082198211156151b4576151b4614fa3565b500190565b634e487b7160e01b600052601260045260246000fd5b6000826151de576151de6151b9565b500490565b60006001600160ff1b038184138284138082168684048611161561520957615209614fa3565b600160ff1b600087128281168783058912161561522857615228614fa3565b6000871292508782058712848416161561524457615244614fa3565b8785058712818416161561525a5761525a614fa3565b505050929093029392505050565b600082615277576152776151b9565b600160ff1b82146000198414161561529157615291614fa3565b50059056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65647d551b4e977cb76b47865bfb6cdc77af264be2349f98f5d2cb58a453757caae7a164736f6c634300080f000a",
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

// Transfer is a paid mutator transaction binding the contract method 0x127217d6.
//
// Solidity: function transfer(address collateral, address to, uint256 amount, uint256 salt, bytes signature) returns(int256, int256)
func (_Accounts *AccountsTransactor) Transfer(opts *bind.TransactOpts, collateral common.Address, to common.Address, amount *big.Int, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "transfer", collateral, to, amount, salt, signature)
}

// Transfer is a paid mutator transaction binding the contract method 0x127217d6.
//
// Solidity: function transfer(address collateral, address to, uint256 amount, uint256 salt, bytes signature) returns(int256, int256)
func (_Accounts *AccountsSession) Transfer(collateral common.Address, to common.Address, amount *big.Int, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Transfer(&_Accounts.TransactOpts, collateral, to, amount, salt, signature)
}

// Transfer is a paid mutator transaction binding the contract method 0x127217d6.
//
// Solidity: function transfer(address collateral, address to, uint256 amount, uint256 salt, bytes signature) returns(int256, int256)
func (_Accounts *AccountsTransactorSession) Transfer(collateral common.Address, to common.Address, amount *big.Int, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Transfer(&_Accounts.TransactOpts, collateral, to, amount, salt, signature)
}

// TransferInsuranceFund is a paid mutator transaction binding the contract method 0x298dc9d5.
//
// Solidity: function transferInsuranceFund(address collateral, address to, uint256 amount) returns(int256 fromBalance, int256 toBalance)
func (_Accounts *AccountsTransactor) TransferInsuranceFund(opts *bind.TransactOpts, collateral common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "transferInsuranceFund", collateral, to, amount)
}

// TransferInsuranceFund is a paid mutator transaction binding the contract method 0x298dc9d5.
//
// Solidity: function transferInsuranceFund(address collateral, address to, uint256 amount) returns(int256 fromBalance, int256 toBalance)
func (_Accounts *AccountsSession) TransferInsuranceFund(collateral common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.TransferInsuranceFund(&_Accounts.TransactOpts, collateral, to, amount)
}

// TransferInsuranceFund is a paid mutator transaction binding the contract method 0x298dc9d5.
//
// Solidity: function transferInsuranceFund(address collateral, address to, uint256 amount) returns(int256 fromBalance, int256 toBalance)
func (_Accounts *AccountsTransactorSession) TransferInsuranceFund(collateral common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.TransferInsuranceFund(&_Accounts.TransactOpts, collateral, to, amount)
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

// UpdateWithdrawalFeeRecipient is a paid mutator transaction binding the contract method 0x27eba4ae.
//
// Solidity: function updateWithdrawalFeeRecipient(address newWithdrawalFeeRecipient) returns()
func (_Accounts *AccountsTransactor) UpdateWithdrawalFeeRecipient(opts *bind.TransactOpts, newWithdrawalFeeRecipient common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "updateWithdrawalFeeRecipient", newWithdrawalFeeRecipient)
}

// UpdateWithdrawalFeeRecipient is a paid mutator transaction binding the contract method 0x27eba4ae.
//
// Solidity: function updateWithdrawalFeeRecipient(address newWithdrawalFeeRecipient) returns()
func (_Accounts *AccountsSession) UpdateWithdrawalFeeRecipient(newWithdrawalFeeRecipient common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateWithdrawalFeeRecipient(&_Accounts.TransactOpts, newWithdrawalFeeRecipient)
}

// UpdateWithdrawalFeeRecipient is a paid mutator transaction binding the contract method 0x27eba4ae.
//
// Solidity: function updateWithdrawalFeeRecipient(address newWithdrawalFeeRecipient) returns()
func (_Accounts *AccountsTransactorSession) UpdateWithdrawalFeeRecipient(newWithdrawalFeeRecipient common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.UpdateWithdrawalFeeRecipient(&_Accounts.TransactOpts, newWithdrawalFeeRecipient)
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

// Withdraw is a paid mutator transaction binding the contract method 0xd8583484.
//
// Solidity: function withdraw(address collateral, address to, uint256 amount, uint256 withdrawAmount, uint256 salt, uint256 fee, bytes data, bytes signature) returns(int256)
func (_Accounts *AccountsTransactor) Withdraw(opts *bind.TransactOpts, collateral common.Address, to common.Address, amount *big.Int, withdrawAmount *big.Int, salt *big.Int, fee *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "withdraw", collateral, to, amount, withdrawAmount, salt, fee, data, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd8583484.
//
// Solidity: function withdraw(address collateral, address to, uint256 amount, uint256 withdrawAmount, uint256 salt, uint256 fee, bytes data, bytes signature) returns(int256)
func (_Accounts *AccountsSession) Withdraw(collateral common.Address, to common.Address, amount *big.Int, withdrawAmount *big.Int, salt *big.Int, fee *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Withdraw(&_Accounts.TransactOpts, collateral, to, amount, withdrawAmount, salt, fee, data, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd8583484.
//
// Solidity: function withdraw(address collateral, address to, uint256 amount, uint256 withdrawAmount, uint256 salt, uint256 fee, bytes data, bytes signature) returns(int256)
func (_Accounts *AccountsTransactorSession) Withdraw(collateral common.Address, to common.Address, amount *big.Int, withdrawAmount *big.Int, salt *big.Int, fee *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _Accounts.Contract.Withdraw(&_Accounts.TransactOpts, collateral, to, amount, withdrawAmount, salt, fee, data, signature)
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

// AccountsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Accounts contract.
type AccountsTransferIterator struct {
	Event *AccountsTransfer // Event containing the contract specifics and raw log

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
func (it *AccountsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsTransfer)
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
		it.Event = new(AccountsTransfer)
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
func (it *AccountsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsTransfer represents a Transfer event raised by the Accounts contract.
type AccountsTransfer struct {
	From         common.Address
	To           common.Address
	Collateral   common.Address
	Amount       *big.Int
	TransferHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xecf06959d9b6442f0e285add9df1ac6dd0730c4c7b7c33aa8bff6583cb3090e3.
//
// Solidity: event Transfer(address indexed from, address indexed to, address indexed collateral, uint256 amount, bytes32 transferHash)
func (_Accounts *AccountsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, collateral []common.Address) (*AccountsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "Transfer", fromRule, toRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return &AccountsTransferIterator{contract: _Accounts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xecf06959d9b6442f0e285add9df1ac6dd0730c4c7b7c33aa8bff6583cb3090e3.
//
// Solidity: event Transfer(address indexed from, address indexed to, address indexed collateral, uint256 amount, bytes32 transferHash)
func (_Accounts *AccountsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AccountsTransfer, from []common.Address, to []common.Address, collateral []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "Transfer", fromRule, toRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsTransfer)
				if err := _Accounts.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xecf06959d9b6442f0e285add9df1ac6dd0730c4c7b7c33aa8bff6583cb3090e3.
//
// Solidity: event Transfer(address indexed from, address indexed to, address indexed collateral, uint256 amount, bytes32 transferHash)
func (_Accounts *AccountsFilterer) ParseTransfer(log types.Log) (*AccountsTransfer, error) {
	event := new(AccountsTransfer)
	if err := _Accounts.contract.UnpackLog(event, "Transfer", log); err != nil {
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
	Fee            *big.Int
	Balance        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9eaa8b17b9a084ea901c6a072e094ed435b7a049ca673db6b17e98203886944e.
//
// Solidity: event Withdraw(address indexed account, address indexed collateral, bytes32 indexed withdrawHash, address to, uint256 amount, uint256 withdrawAmount, uint256 fee, int256 balance)
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

// WatchWithdraw is a free log subscription operation binding the contract event 0x9eaa8b17b9a084ea901c6a072e094ed435b7a049ca673db6b17e98203886944e.
//
// Solidity: event Withdraw(address indexed account, address indexed collateral, bytes32 indexed withdrawHash, address to, uint256 amount, uint256 withdrawAmount, uint256 fee, int256 balance)
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

// ParseWithdraw is a log parse operation binding the contract event 0x9eaa8b17b9a084ea901c6a072e094ed435b7a049ca673db6b17e98203886944e.
//
// Solidity: event Withdraw(address indexed account, address indexed collateral, bytes32 indexed withdrawHash, address to, uint256 amount, uint256 withdrawAmount, uint256 fee, int256 balance)
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

// AccountsWithdrawalFeeRecipientUpdatedIterator is returned from FilterWithdrawalFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for WithdrawalFeeRecipientUpdated events raised by the Accounts contract.
type AccountsWithdrawalFeeRecipientUpdatedIterator struct {
	Event *AccountsWithdrawalFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *AccountsWithdrawalFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsWithdrawalFeeRecipientUpdated)
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
		it.Event = new(AccountsWithdrawalFeeRecipientUpdated)
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
func (it *AccountsWithdrawalFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsWithdrawalFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsWithdrawalFeeRecipientUpdated represents a WithdrawalFeeRecipientUpdated event raised by the Accounts contract.
type AccountsWithdrawalFeeRecipientUpdated struct {
	WithdrawlFeeRecipient common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFeeRecipientUpdated is a free log retrieval operation binding the contract event 0xcbc6a0c385161960535b52c682db3a205870921d9a6cefbd50ada57faadb6976.
//
// Solidity: event WithdrawalFeeRecipientUpdated(address indexed withdrawlFeeRecipient)
func (_Accounts *AccountsFilterer) FilterWithdrawalFeeRecipientUpdated(opts *bind.FilterOpts, withdrawlFeeRecipient []common.Address) (*AccountsWithdrawalFeeRecipientUpdatedIterator, error) {

	var withdrawlFeeRecipientRule []interface{}
	for _, withdrawlFeeRecipientItem := range withdrawlFeeRecipient {
		withdrawlFeeRecipientRule = append(withdrawlFeeRecipientRule, withdrawlFeeRecipientItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "WithdrawalFeeRecipientUpdated", withdrawlFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &AccountsWithdrawalFeeRecipientUpdatedIterator{contract: _Accounts.contract, event: "WithdrawalFeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFeeRecipientUpdated is a free log subscription operation binding the contract event 0xcbc6a0c385161960535b52c682db3a205870921d9a6cefbd50ada57faadb6976.
//
// Solidity: event WithdrawalFeeRecipientUpdated(address indexed withdrawlFeeRecipient)
func (_Accounts *AccountsFilterer) WatchWithdrawalFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *AccountsWithdrawalFeeRecipientUpdated, withdrawlFeeRecipient []common.Address) (event.Subscription, error) {

	var withdrawlFeeRecipientRule []interface{}
	for _, withdrawlFeeRecipientItem := range withdrawlFeeRecipient {
		withdrawlFeeRecipientRule = append(withdrawlFeeRecipientRule, withdrawlFeeRecipientItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "WithdrawalFeeRecipientUpdated", withdrawlFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsWithdrawalFeeRecipientUpdated)
				if err := _Accounts.contract.UnpackLog(event, "WithdrawalFeeRecipientUpdated", log); err != nil {
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

// ParseWithdrawalFeeRecipientUpdated is a log parse operation binding the contract event 0xcbc6a0c385161960535b52c682db3a205870921d9a6cefbd50ada57faadb6976.
//
// Solidity: event WithdrawalFeeRecipientUpdated(address indexed withdrawlFeeRecipient)
func (_Accounts *AccountsFilterer) ParseWithdrawalFeeRecipientUpdated(log types.Log) (*AccountsWithdrawalFeeRecipientUpdated, error) {
	event := new(AccountsWithdrawalFeeRecipientUpdated)
	if err := _Accounts.contract.UnpackLog(event, "WithdrawalFeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
