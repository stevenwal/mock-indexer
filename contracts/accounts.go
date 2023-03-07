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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_domain\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_quote\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadLoadAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadRegisterSigningKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadRevokeSigningKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BadWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotKeeper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Credit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"}],\"name\":\"CreditInstrument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Debit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"}],\"name\":\"DebitInstrument\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"InstrumentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"InstrumentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"KeeperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NewAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"RegisteredSigningKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"revokeHash\",\"type\":\"bytes32\"}],\"name\":\"RevokedSigningKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"VersionInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"WithdrawProxyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountsList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"addInstrument\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorities\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collaterals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"creditInstrument\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"debitInstrument\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Balance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Position[]\",\"name\":\"positions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structAccountsInterface.SigningKey[]\",\"name\":\"signingKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccountsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"getCollateralBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollaterals\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"getInstrumentEntryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"getInstrumentPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInstruments\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInsuranceFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getQuoteBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getSigningKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"}],\"name\":\"hasSigningKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"instruments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keepers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Balance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"}],\"internalType\":\"structAccountsInterface.Position[]\",\"name\":\"positions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structAccountsInterface.SigningKey[]\",\"name\":\"signingKeys\",\"type\":\"tuple[]\"}],\"name\":\"loadAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signingKeySig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountSig\",\"type\":\"bytes\"}],\"name\":\"registerSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"removeCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"removeInstrument\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instrument\",\"type\":\"uint256\"}],\"name\":\"resetInstrumentPosition\",\"outputs\":[{\"internalType\":\"enumAccountsInterface.RevertReason\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signingKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"revokeSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signingKeyExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalPositions\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawProxy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateWithdrawProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawInsuranceFund\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawProxies\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101a0604052306080523480156200001657600080fd5b5060405162005dd438038062005dd4833981016040819052620000399162000433565b6001600160a01b038116620000615760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811660a08190526040805163313ce56760e01b8152905163313ce567916004808201926020929091908290030181865afa158015620000ac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000d2919062000509565b60ff1660c052620000f08383620002c3602090811b62003c2417901c565b60e052505060408051680aed2e8d0c8e4c2ee560bb1b6020808301919091527f6164647265737320636f6c6c61746572616c2c0000000000000000000000000060298301526a1859191c995cdcc81d1bcb60aa1b603c8301526e1d5a5b9d0c8d4d88185b5bdd5b9d0b608a1b60478301526c1d5a5b9d0c8d4d881cd85b1d0b609a1b60568301526b62797465733332206461746160a01b6063830152602960f81b606f83018190528351605081850301815260708401855280519083012061010052670a6d2cedc96caf2560c31b60908401526e1859191c995cdcc81858d8dbdd5b9d608a1b609884015260a783018190528351608881850301815260a88401855280519083012061012052680a4caced2e6e8cae4560bb1b60c88401526b1859191c995cdcc81ad95e4b60a21b60d18401526d75696e743235362065787069727960901b60dd84015260eb8301819052835160cc81850301815260ec8401855280519083012061014052660a4caecded6ca560cb1b61010c8401526a61646472657373206b657960a81b61011384015261011e830152825160ff81840301815261011f90920190925280519101206101605250736cdc77af264be2349f98f5d2cb58a453757caae76101805262000553565b6040516c08a92a06e626488dedac2d2dc5609b1b60208201526b1cdd1c9a5b99c81b985b594b60a21b602d8201526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398201526e1d5a5b9d0c8d4d8818da185a5b9259608a1b6048820152602960f81b605782015260009060580160405160208183030381529060405280519060200120836040516020016200035a919062000535565b60408051601f198184030181528282528051602091820120908301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201526080810183905260a00160405160208183030381529060405280519060200120905092915050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101562000400578181015183820152602001620003e6565b8381111562000410576000848401525b50505050565b80516001600160a01b03811681146200042e57600080fd5b919050565b6000806000606084860312156200044957600080fd5b83516001600160401b03808211156200046157600080fd5b818601915086601f8301126200047657600080fd5b8151818111156200048b576200048b620003cd565b604051601f8201601f19908116603f01168101908382118183101715620004b657620004b6620003cd565b81604052828152896020848701011115620004d057600080fd5b620004e3836020830160208801620003e3565b809750505050505060208401519150620005006040850162000416565b90509250925092565b6000602082840312156200051c57600080fd5b815160ff811681146200052e57600080fd5b9392505050565b6000825162000549818460208701620003e3565b9190910192915050565b60805160a05160c05160e05161010051610120516101405161016051610180516157a062000634600039600081816103d901528181613487015281816135050152818161369d015261434901526000610bec015260006112e3015260006114820152600061417f015260008181610b6901528181610bcb0152818161132801528181611461015261415e01526000610530015260008181610828015281816108e301528181612d5d01528181612dd60152612e9401526000818161120301528181611243015281816115ee0152818161162e01526116c101526157a06000f3fe60806040526004361061031a5760003560e01c80638340f549116101ab578063c4d66de8116100f7578063ea547fb811610095578063f0d2d5a81161006f578063f0d2d5a814610b17578063f160d36914610b37578063f698da2514610b57578063fbcbc0f114610b8b57600080fd5b8063ea547fb814610a7d578063eeb97d3b14610ac7578063efd1952314610af757600080fd5b8063d3057877116100d1578063d3057877146109ed578063d5c2073014610a0d578063d658d2e914610a2d578063e30e838414610a5d57600080fd5b8063c4d66de81461098d578063c99d3a06146109ad578063cf9b3c8d146109cd57600080fd5b8063999b93af11610164578063ac759c711161013e578063ac759c711461088a578063af1aa3b1146108aa578063b7d5820b14610913578063b8eb94e91461095d57600080fd5b8063999b93af146108165780639aeddeff1461084a5780639d56921f1461086a57600080fd5b80638340f549146107355780638869a3f2146107555780638a48ac03146107845780638da5cb5b146107995780638e991606146107b957806391223d69146107e657600080fd5b8063469048401161026a5780636205ed5e116102235780636eacd398116101fd5780636eacd398146106b357806372d39ed6146106e057806378b92636146107005780637de182c51461071557600080fd5b80636205ed5e1461061957806362946d3b146106465780636cd22eaf1461069357600080fd5b8063469048401461056457806348e4ccbe1461058457806349ad755b146105a45780634f1ef286146105c457806352d1902d146105d757806361d21920146105ec57600080fd5b806322bbad0b116102d75780633659cfe6116102b15780633659cfe6146104ce5780633bbd64bc146104ee5780633fd1e2bd1461051e5780634614be9f146103ca57600080fd5b806322bbad0b14610433578063241d400d14610473578063262709e6146104a057600080fd5b806306fdde031461031f57806307b1b8b21461036957806313af40351461038b57806314f326a1146103ab578063158626f7146103ca5780631635717c14610411575b600080fd5b34801561032b57600080fd5b50610353604051806040016040528060088152602001674163636f756e747360c01b81525081565b6040516103609190614bdb565b60405180910390f35b34801561037557600080fd5b50610389610384366004614d08565b610bbc565b005b34801561039757600080fd5b506103896103a6366004614d56565b610fbe565b3480156103b757600080fd5b50600c545b604051908152602001610360565b3480156103d657600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b039091168152602001610360565b34801561041d57600080fd5b50610426610fd2565b6040516103609190614d71565b34801561043f57600080fd5b5061046361044e366004614db5565b600b6020526000908152604090205460ff1681565b6040519015158152602001610360565b34801561047f57600080fd5b5061049361048e366004614d56565b61102a565b6040516103609190614dce565b3480156104ac57600080fd5b506104c06104bb366004614e0f565b6110a3565b604051610360929190614e6d565b3480156104da57600080fd5b506103896104e9366004614d56565b6111f9565b3480156104fa57600080fd5b50610463610509366004614d56565b60046020526000908152604090205460ff1681565b34801561052a57600080fd5b506105527f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff9091168152602001610360565b34801561057057600080fd5b506005546103f9906001600160a01b031681565b34801561059057600080fd5b5061038961059f366004614e81565b6112d5565b3480156105b057600080fd5b506103bc6105bf366004614eff565b6115b6565b6103896105d2366004614d08565b6115e4565b3480156105e357600080fd5b506103bc6116b4565b3480156105f857600080fd5b506103bc610607366004614db5565b60106020526000908152604090205481565b34801561062557600080fd5b506103bc610634366004614d56565b600e6020526000908152604090205481565b34801561065257600080fd5b50610463610661366004614f29565b6001600160a01b039182166000908152600d602090815260408083209390941682526002909201909152205460ff1690565b34801561069f57600080fd5b506103896106ae366004614f5c565b611767565b3480156106bf57600080fd5b506103bc6106ce366004614d56565b600f6020526000908152604090205481565b3480156106ec57600080fd5b506103bc6106fb366004614f98565b6117ea565b34801561070c57600080fd5b506104936118c0565b34801561072157600080fd5b506104c0610730366004614e0f565b611921565b34801561074157600080fd5b506103bc610750366004614e0f565b611a5d565b34801561076157600080fd5b5061077561077036600461503a565b611aa2565b60405161036093929190615073565b34801561079057600080fd5b50610493611c06565b3480156107a557600080fd5b506000546103f9906001600160a01b031681565b3480156107c557600080fd5b506107d96107d4366004614db5565b611c66565b604051610360919061508e565b3480156107f257600080fd5b50610463610801366004614d56565b60036020526000908152604090205460ff1681565b34801561082257600080fd5b506103f97f000000000000000000000000000000000000000000000000000000000000000081565b34801561085657600080fd5b506107d9610865366004614eff565b611d4e565b34801561087657600080fd5b506103896108853660046151ea565b611f0d565b34801561089657600080fd5b506107756108a536600461503a565b612b7d565b3480156108b657600080fd5b506103bc6108c5366004614d56565b6001600160a01b039081166000908152600d602090815260408083207f0000000000000000000000000000000000000000000000000000000000000000909416835260039093019052205490565b34801561091f57600080fd5b506103bc61092e366004614f29565b6001600160a01b039182166000908152600d602090815260408083209390941682526003909201909152205490565b34801561096957600080fd5b50610463610978366004614d56565b60066020526000908152604090205460ff1681565b34801561099957600080fd5b506103896109a8366004614d56565b612cc7565b3480156109b957600080fd5b506103896109c8366004614d56565b612e51565b3480156109d957600080fd5b506103f96109e8366004614db5565b613096565b3480156109f957600080fd5b50610389610a08366004614f5c565b6130c0565b348015610a1957600080fd5b50610389610a28366004614f5c565b613143565b348015610a3957600080fd5b50610463610a48366004614db5565b60076020526000908152604090205460ff1681565b348015610a6957600080fd5b506107d9610a78366004614db5565b6131c6565b348015610a8957600080fd5b506103bc610a98366004614eff565b6001600160a01b03919091166000908152600d6020908152604080832093835260049093019052206001015490565b348015610ad357600080fd5b50610463610ae2366004614d56565b60096020526000908152604090205460ff1681565b348015610b0357600080fd5b506103bc610b12366004614e0f565b613376565b348015610b2357600080fd5b50610389610b32366004614d56565b6136fa565b348015610b4357600080fd5b50610389610b52366004614d56565b6137f2565b348015610b6357600080fd5b506103bc7f000000000000000000000000000000000000000000000000000000000000000081565b348015610b9757600080fd5b50610bab610ba6366004614d56565b61386b565b60405161036095949392919061538f565b610bc4613d2c565b6000610c657f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000085604051602001610c2f9291909182526001600160a01b0316602082015260400190565b6040516020818303038152906040528051906020012061190160f01b600090815260029290925260229081526042822091905290565b3360009081526004602052604090205490915060ff16610ca9578060016040516001622b6d3960e21b03198152600401610ca0929190614e6d565b60405180910390fd5b6001600160a01b038316610cd8578060036040516001622b6d3960e21b03198152600401610ca0929190614e6d565b6001600160a01b0383166000908152600e60205260408120549003610d185780600a6040516001622b6d3960e21b03198152600401610ca0929190614e6d565b6000610d248284613d55565b90506001600160a01b038116610d55578160076040516001622b6d3960e21b03198152600401610ca0929190614e6d565b6001600160a01b038082166000908152600d60209081526040808320938816835260029093019052205460ff16610da75781600a6040516001622b6d3960e21b03198152600401610ca0929190614e6d565b6001600160a01b0381166000908152600d6020908152604080832060010180548251818502810185019093528083528493830182828015610e1157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610df3575b505050505090505b8051821015610e6057856001600160a01b0316818381518110610e3e57610e3e615423565b60200260200101516001600160a01b03160315610e6057816001019150610e19565b80518210610e895783600a6040516001622b6d3960e21b03198152600401610ca0929190614e6d565b8060018251610e98919061544f565b81518110610ea857610ea8615423565b6020026020010151600d6000856001600160a01b03166001600160a01b031681526020019081526020016000206001018381548110610ee957610ee9615423565b600091825260208083209190910180546001600160a01b0319166001600160a01b039485161790559185168152600d90915260409020600101805480610f3157610f31615466565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03858116808452600d83526040808520928b168086526002909301909352828420805460ff19169055915187939192917ff0d72ac27c7920d0a69f77bbf64f8509a0307c49e54eee377861c46101ef25cb91a45050600160025550505050565b610fc6613e78565b610fcf81613ea4565b50565b6060600a80548060200260200160405190810160405280929190818152602001828054801561102057602002820191906000526020600020905b81548152602001906001019080831161100c575b5050505050905090565b6001600160a01b0381166000908152600d602090815260409182902060010180548351818402810184019094528084526060939283018282801561109757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611079575b50505050509050919050565b33600090815260036020526040812054819060ff166110d557604051631890934360e11b815260040160405180910390fd5b6110dd613d2c565b6001600160a01b03841660009081526009602052604090205460ff1661110957506000905060046111ea565b61111285613f13565b506001600160a01b038086166000908152600d602090815260408083209388168352600390930190529081205461114a90859061547c565b6001600160a01b038088166000908152600d60209081526040808320938a16835260039093018152828220849055600f90529081208054929350869290919061119490849061547c565b909155505060408051858152602081018390526001600160a01b0380881692908916917fcb9aff7bd66be0077d57437ece119a9f166aad37de4491ac4f92fdb2c1ad8d3491015b60405180910390a39150600090505b60016002559094909350915050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036112415760405162461bcd60e51b8152600401610ca0906154bb565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661128a60008051602061572d833981519152546001600160a01b031690565b6001600160a01b0316146112b05760405162461bcd60e51b8152600401610ca090615507565b6112b981613fdc565b60408051600080825260208201909252610fcf91839190613fe4565b6112dd613d2c565b604080517f000000000000000000000000000000000000000000000000000000000000000060208201526001600160a01b0386169181019190915260608101849052600090611350907f000000000000000000000000000000000000000000000000000000000000000090608001610c2f565b3360009081526004602052604090205490915060ff1661138857806001604051634e7391e160e11b8152600401610ca0929190614e6d565b6001600160a01b0385166113b457806003604051634e7391e160e11b8152600401610ca0929190614e6d565b836000036113da57806002604051634e7391e160e11b8152600401610ca0929190614e6d565b6001600160a01b0385166000908152600e6020526040902054156114165780600a604051634e7391e160e11b8152600401610ca0929190614e6d565b60006114228284613d55565b90506001600160a01b03811661145057816007604051634e7391e160e11b8152600401610ca0929190614e6d565b61145981613f13565b5060006114c57f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000084604051602001610c2f9291909182526001600160a01b0316602082015260400190565b905060006114d38287613d55565b9050876001600160a01b0316816001600160a01b03161461150c5783600b604051634e7391e160e11b8152600401610ca0929190614e6d565b6001600160a01b038881166000818152600e602090815260408083208c9055938716808352600d82528483206001808201805480830182559086528486200180546001600160a01b031916871790558585526002909101835292859020805460ff191690931790925592518a815287937f1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac910160405180910390a450506001600255505050505050565b6001600160a01b0382166000908152600d602090815260408083208484526004019091529020545b92915050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361162c5760405162461bcd60e51b8152600401610ca0906154bb565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661167560008051602061572d833981519152546001600160a01b031690565b6001600160a01b03161461169b5760405162461bcd60e51b8152600401610ca090615507565b6116a482613fdc565b6116b082826001613fe4565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146117545760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610ca0565b5060008051602061572d83398151915290565b61176f613e78565b6001600160a01b0382166117965760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260036020526040808220805460ff191685151590811790915590519092917fc5238a63fa205f95e98807f6cbc91d2ad8555a2250075e47491f11d9c69db3bc91a35050565b60006117f4613d2c565b6000806118068a8a8a8a8a8a8a614154565b9150915061181382613f13565b506118286001600160a01b038b168a8961453b565b6001600160a01b03891660009081526006602052604090205460ff16156118ae57604051631e6fd1a760e31b81526001600160a01b038a169063f37e8d389061187b9085908e908c908b90600401615553565b600060405180830381600087803b15801561189557600080fd5b505af11580156118a9573d6000803e3d6000fd5b505050505b60016002559998505050505050505050565b6060600880548060200260200160405190810160405280929190818152602001828054801561102057602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116118fa575050505050905090565b33600090815260036020526040812054819060ff1661195357604051631890934360e11b815260040160405180910390fd5b61195b613d2c565b6001600160a01b03841660009081526009602052604090205460ff1661198757506000905060046111ea565b61199085613f13565b506001600160a01b038086166000908152600d60209081526040808320938816835260039093019052908120546119c8908590615586565b6001600160a01b038088166000908152600d60209081526040808320938a16835260039093018152828220849055600f905290812080549293508692909190611a12908490615586565b909155505060408051858152602081018390526001600160a01b0380881692908916917fc9a1b9e3b11def04d81baa28220fa980997040aef1abec95d1cd2bcfd4ba8a6591016111db565b6000611a67613d2c565b611a7084613f13565b506000611a7e8585856145b9565b9050611a956001600160a01b038516333086614715565b6001600255949350505050565b336000908152600360205260408120548190819060ff16611ad657604051631890934360e11b815260040160405180910390fd5b611ade613d2c565b6000868152600b602052604090205460ff16611b035750600091508190506005611bf4565b6001600160ff1b03851115611b215750600091508190506002611bf4565b611b2a87613f13565b506001600160a01b0387166000908152600d60209081526040808320898452600401909152812081908190611b60908989614798565b60008c815260106020526040812080549497509295509093508a92611b86908490615586565b90915550506040805189815260208101899052908101849052606081018390526080810182905289906001600160a01b038c16907f27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae653639060a0015b60405180910390a391945090925060009150505b60016002819055509450945094915050565b6060600c805480602002602001604051908101604052809291908181526020018280548015611020576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116118fa575050505050905090565b3360009081526003602052604081205460ff16611c9657604051631890934360e11b815260040160405180910390fd5b611c9e613d2c565b81600003611cae57506002611d44565b6000828152600b602052604090205460ff1615611ccd57506005611d44565b600a805460018181019092557fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8018390556000838152600b6020526040808220805460ff1916909317909255905183917fb8597a358785eb73ad0a81079eb5e50ba0ac43dd023f47c0e0ff47e58a9495c891a25060005b6001600255919050565b3360009081526003602052604081205460ff16611d7e57604051631890934360e11b815260040160405180910390fd5b611d86613d2c565b6000828152600b602052604090205460ff16611da457506005611f02565b611dad83613f13565b506001600160a01b0383166000908152600d6020908152604080832085845260040190915281208054828255600190910182905590811315611e6c5760008381526010602052604081208054839290611e0790849061547c565b909155505060408051828152600060208201819052818301819052606082018190526080820152905184916001600160a01b038716917fc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d19181900360a00190a3611efc565b6000811215611efc5760008381526010602052604081208054839290611e9390849061547c565b909155508390506001600160a01b0385167f27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae65363611ece846155c7565b6040805191825260006020830181905290820181905260608201819052608082015260a00160405180910390a35b60009150505b600160025592915050565b611f15613e78565b611f1e84613f13565b15611f3f57600760405163eb1b129160e01b8152600401610ca0919061508e565b60005b83518110156123ed5760096000858381518110611f6157611f61615423565b602090810291909101810151516001600160a01b031682528101919091526040016000205460ff16611fa957600460405163eb1b129160e01b8152600401610ca0919061508e565b838181518110611fbb57611fbb615423565b6020026020010151602001516000148061203e5750600d6000866001600160a01b03166001600160a01b03168152602001908152602001600020600301600085838151811061200c5761200c615423565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054600014155b1561205f57600260405163eb1b129160e01b8152600401610ca0919061508e565b83818151811061207157612071615423565b602002602001015160200151600d6000876001600160a01b03166001600160a01b0316815260200190815260200160002060030160008684815181106120b9576120b9615423565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000208190555060008482815181106120fd576120fd615423565b602002602001015160200151600f600087858151811061211f5761211f615423565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020546121569190615586565b905080600f600087858151811061216f5761216f615423565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000208190555060008113801561223457508482815181106121bc576121bc615423565b6020908102919091010151516040516370a0823160e01b81523060048201526001600160a01b03909116906370a0823190602401602060405180830381865afa15801561220d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061223191906155e3565b81115b1561225557600860405163eb1b129160e01b8152600401610ca0919061508e565b600085838151811061226957612269615423565b602002602001015160200151131561232d5784828151811061228d5761228d615423565b6020026020010151600001516001600160a01b0316866001600160a01b03167fc9a1b9e3b11def04d81baa28220fa980997040aef1abec95d1cd2bcfd4ba8a658785815181106122df576122df615423565b6020026020010151602001518886815181106122fd576122fd615423565b602002602001015160200151604051612320929190918252602082015260400190565b60405180910390a36123e4565b84828151811061233f5761233f615423565b6020026020010151600001516001600160a01b0316866001600160a01b03167fcb9aff7bd66be0077d57437ece119a9f166aad37de4491ac4f92fdb2c1ad8d3487858151811061239157612391615423565b6020026020010151602001516123a6906155c7565b8886815181106123b8576123b8615423565b6020026020010151602001516040516123db929190918252602082015260400190565b60405180910390a35b50600101611f42565b5060005b825181101561282b57600b600084838151811061241057612410615423565b6020908102919091018101515182528101919091526040016000205460ff1661244f57600560405163eb1b129160e01b8152600401610ca0919061508e565b82818151811061246157612461615423565b602002602001015160400151600014806124d55750600d6000866001600160a01b03166001600160a01b0316815260200190815260200160002060040160008483815181106124b2576124b2615423565b602002602001015160000151815260200190815260200160002060000154600014155b156124f657600260405163eb1b129160e01b8152600401610ca0919061508e565b604051806040016040528084838151811061251357612513615423565b602002602001015160400151815260200184838151811061253657612536615423565b602002602001015160200151815250600d6000876001600160a01b03166001600160a01b03168152602001908152602001600020600401600085848151811061258157612581615423565b602002602001015160000151815260200190815260200160002060008201518160000155602082015181600101559050508281815181106125c4576125c4615423565b602002602001015160400151601060008584815181106125e6576125e6615423565b6020026020010151600001518152602001908152602001600020600082825461260f9190615586565b92505081905550600083828151811061262a5761262a615423565b602002602001015160400151131561272d5782818151811061264e5761264e615423565b602002602001015160000151856001600160a01b03167f27dc6d35b2e946addacf727cf0484c1cf2c5e5b23c9d3379b610c8d20ae6536385848151811061269757612697615423565b6020026020010151604001518685815181106126b5576126b5615423565b6020026020010151602001518786815181106126d3576126d3615423565b6020026020010151604001518887815181106126f1576126f1615423565b602090810291909101810151810151604080519586529185019390935283015260608201526000608082015260a00160405180910390a3612823565b82818151811061273f5761273f615423565b602002602001015160000151856001600160a01b03167fc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d185848151811061278857612788615423565b60200260200101516040015161279d906155c7565b8685815181106127af576127af615423565b6020026020010151602001518786815181106127cd576127cd615423565b6020026020010151604001518887815181106127eb576127eb615423565b602090810291909101810151810151604080519586529185019390935283015260608201526000608082015260a00160405180910390a35b6001016123f1565b5060005b8151811015612b7657600e600083838151811061284e5761284e615423565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000205460001461289e57600b60405163eb1b129160e01b8152600401610ca0919061508e565b60006001600160a01b03168282815181106128bb576128bb615423565b6020026020010151600001516001600160a01b031614806128f957508181815181106128e9576128e9615423565b6020026020010151602001516000145b1561291a57600260405163eb1b129160e01b8152600401610ca0919061508e565b600d6000866001600160a01b03166001600160a01b03168152602001908152602001600020600201600083838151811061295657612956615423565b602090810291909101810151516001600160a01b031682528101919091526040016000205460ff161561299f57600760405163eb1b129160e01b8152600401610ca0919061508e565b8181815181106129b1576129b1615423565b602002602001015160200151600e60008484815181106129d3576129d3615423565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002081905550600d6000866001600160a01b03166001600160a01b03168152602001908152602001600020600101828281518110612a3d57612a3d615423565b60209081029190910181015151825460018082018555600094855283852090910180546001600160a01b0319166001600160a01b039384161790559088168352600d9091526040822084519192600290910191859085908110612aa257612aa2615423565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff0219169083151502179055506000828281518110612af957612af9615423565b6020026020010151600001516001600160a01b0316866001600160a01b03167f1223e2b0106e2b0d555d15671ece0909fbf091c37058aaee0e092b9e77eb54ac858581518110612b4b57612b4b615423565b602002602001015160200151604051612b6691815260200190565b60405180910390a460010161282f565b5050505050565b336000908152600360205260408120548190819060ff16612bb157604051631890934360e11b815260040160405180910390fd5b612bb9613d2c565b6000868152600b602052604090205460ff16612bde5750600091508190506005611bf4565b6001600160ff1b03851115612bfc5750600091508190506002611bf4565b612c0587613f13565b5060008080612c43612c16896155c7565b6001600160a01b038c166000908152600d602090815260408083208e845260040190915290209089614798565b60008c815260106020526040812080549497509295509093508a92612c6990849061547c565b90915550506040805189815260208101899052908101849052606081018390526080810182905289906001600160a01b038c16907fc56bfa56b404728134fe0ff4d49dae3033c8500daf1d798cff3fe28b2645b2d19060a001611be0565b6001612cd3818061544f565b60015414612cf4576040516302ed543d60e51b815260040160405180910390fd5b80600103612d025760016002555b612d0b82613ea4565b600580546001600160a01b038085166001600160a01b031992831681179093556008805460018181019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180547f0000000000000000000000000000000000000000000000000000000000000000909316929093168217909255600090815260096020526040808220805460ff191690931790925590517f7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc29190a26040516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016907f7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f90600090a2600181905560405181907f7a621ac638ec0ed2c353a3800daf6854f8682f565af567ad99fc910f4755938d90600090a25050565b612e59613e78565b6001600160a01b03811660009081526009602052604090205460ff16612e92576040516368f7a67560e11b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603612ee4576040516368f7a67560e11b815260040160405180910390fd5b6000806008805480602002602001604051908101604052809291908181526020018280548015612f3d57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612f1f575b505050505090505b8051821015612f8c57826001600160a01b0316818381518110612f6a57612f6a615423565b60200260200101516001600160a01b03160315612f8c57816001019150612f45565b80518210612fad576040516368f7a67560e11b815260040160405180910390fd5b8060018251612fbc919061544f565b81518110612fcc57612fcc615423565b602002602001015160088381548110612fe757612fe7615423565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600880548061302657613026615466565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03851680835260099091526040808320805460ff191690555190917fd89d2ee68ab04dca0193f48a4aff55e20fa5ec0429a8a8c1c51b8dad6178a59391a2505050565b600c81815481106130a657600080fd5b6000918252602090912001546001600160a01b0316905081565b6130c8613e78565b6001600160a01b0382166130ef5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260046020526040808220805460ff191685151590811790915590519092917f786c9db967bf0c6b16c7c91adae8a8c554b15a57d373fa2059607300f4616c0091a35050565b61314b613e78565b6001600160a01b0382166131725760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038216600081815260066020526040808220805460ff191685151590811790915590519092917f097359b7e2dca2adc06bfc6c3e4e8fcc69a2e916105499e8e4c60a569598cbc091a35050565b3360009081526003602052604081205460ff166131f657604051631890934360e11b815260040160405180910390fd5b6131fe613d2c565b6000828152600b602052604090205460ff1661321c57506005611d44565b600080600a80548060200260200160405190810160405280929190818152602001828054801561326b57602002820191906000526020600020905b815481526020019060010190808311613257575b505050505090505b80518210156132a8578381838151811061328f5761328f615423565b602002602001015103156132a857816001019150613273565b805182106132bb57600592505050611d44565b80600182516132ca919061544f565b815181106132da576132da615423565b6020026020010151600a83815481106132f5576132f5615423565b600091825260209091200155600a80548061331257613312615466565b600082815260208082208301600019908101839055909201909255858252600b90526040808220805460ff191690555185917f08b9963dddd96e14a775b44c82852b5a482f5060ae86706c04fc5878c3fe3d8d91a250506001600255506000919050565b6000613380613d2c565b3360009081526004602052604090205460ff166133c357600080516020615774833981519152600160405163c802b7f360e01b8152600401610ca0929190614e6d565b816000036133f757600080516020615774833981519152600260405163c802b7f360e01b8152600401610ca0929190614e6d565b6001600160a01b03831661343157600080516020615774833981519152600360405163c802b7f360e01b8152600401610ca0929190614e6d565b6001600160a01b03841660009081526009602052604090205460ff1661347d57600080516020615774833981519152600460405163c802b7f360e01b8152600401610ca0929190614e6d565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166000908152600d602090815260408083209388168352600390930190522054828112156134fb57600080516020615774833981519152600860405163c802b7f360e01b8152600401610ca0929190614e6d565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166000908152600d602090815260408083209389168352600390930181528282209386900393849055600f9052908120805485929061356590849061547c565b9091555061357f90506001600160a01b038616858561453b565b6001600160a01b03841660009081526006602052604090205460ff161561362157600554604051631e6fd1a760e31b81526001600160a01b03918216600482015286821660248201526044810185905260806064820152600060848201529085169063f37e8d389060a401600060405180830381600087803b15801561360457600080fd5b505af1158015613618573d6000803e3d6000fd5b50505050613662565b6005546001600160a01b0385811691161461366257600080516020615774833981519152600c60405163c802b7f360e01b8152600401610ca0929190614e6d565b604080516001600160a01b038681168252602082018690529181018590526060810183905260008051602061577483398151915291808816917f0000000000000000000000000000000000000000000000000000000000000000909116907f66efaee033f9c338e1e0c145c8653403e5f05eddfa1e35f67931c6ba3e231ea09060800160405180910390a46001600255949350505050565b613702613e78565b6001600160a01b0381166137295760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03811660009081526009602052604090205460ff1615613763576040516368f7a67560e11b815260040160405180910390fd5b6008805460018082019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180546001600160a01b0319166001600160a01b038416908117909155600081815260096020526040808220805460ff1916909417909355915190917f7db05e63d635a68c62fd7fd8f3107ae8ab584a383e102d1bd8a40f4c977e465f91a250565b6137fa613e78565b6001600160a01b0381166138215760405163d92e233d60e01b815260040160405180910390fd5b600580546001600160a01b0319166001600160a01b0383169081179091556040517f7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc290600090a250565b6001600160a01b0381166000908152600d602052604081205460085460ff90911691606091829182918067ffffffffffffffff8111156138ad576138ad614c05565b6040519080825280602002602001820160405280156138f257816020015b60408051808201909152600080825260208201528152602001906001900390816138cb5790505b50945060005b818110156139835760006008828154811061391557613915615423565b60009182526020808320909101546040805180820182526001600160a01b03928316808252928e168552600d8452818520838652600301845293205491830191909152885190925088908490811061396f5761396f615423565b6020908102919091010152506001016138f8565b50600a548067ffffffffffffffff8111156139a0576139a0614c05565b6040519080825280602002602001820160405280156139f557816020015b6139e260405180606001604052806000815260200160008152602001600081525090565b8152602001906001900390816139be5790505b50945060005b81811015613a97576000600a8281548110613a1857613a18615423565b6000918252602080832090910154604080516060810182528281526001600160a01b038f168552600d845281852083865260040180855282862060018101548387015295849052909352925492820192909252885191925090889084908110613a8357613a83615423565b6020908102919091010152506001016139fb565b506001600160a01b0388166000908152600d6020908152604080832060010180548251818502810185019093528083529192909190830182828015613b0557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613ae7575b50505050509050805167ffffffffffffffff811115613b2657613b26614c05565b604051908082528060200260200182016040528015613b6b57816020015b6040805180820190915260008082526020820152815260200190600190039081613b445790505b50945060005b8151811015613c14576040518060400160405280838381518110613b9757613b97615423565b60200260200101516001600160a01b03168152602001600e6000858581518110613bc357613bc3615423565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054815250868281518110613c0157613c01615423565b6020908102919091010152600101613b71565b5043935050505091939590929450565b6040516c08a92a06e626488dedac2d2dc5609b1b60208201526b1cdd1c9a5b99c81b985b594b60a21b602d8201526e1cdd1c9a5b99c81d995c9cda5bdb8b608a1b60398201526e1d5a5b9d0c8d4d8818da185a5b9259608a1b6048820152602960f81b60578201526000906058016040516020818303038152906040528051906020012083604051602001613cb991906155fc565b60408051601f198184030181528282528051602091820120908301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201526080810183905260a00160405160208183030381529060405280519060200120905092915050565b600254600114613d4f5760405163558a1e0360e11b815260040160405180910390fd5b60028055565b6000806000808451604003613d8757505050602082015160408301516001600160ff1b0381169060ff1c601b01613ddc565b8451604103613dd05750505060208201516040830151606084015160001a601b8114801590613dba57508060ff16601c14155b15613dcb57600093505050506115de565b613ddc565b600093505050506115de565b7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115613e1057600093505050506115de565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015613e63573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6000546001600160a01b03163314613ea2576040516282b42960e81b815260040160405180910390fd5b565b6001600160a01b038116613ecb5760405163d92e233d60e01b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b038316908117825560405190917f4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b91a250565b60006001600160a01b03821615613fd7576001600160a01b0382166000908152600d602052604090205460ff1615613f4d57506001919050565b6001600160a01b0382166000818152600d6020526040808220805460ff19166001908117909155600c8054918201815583527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c70180546001600160a01b03191684179055517fef4ab4f35cd2027fcc6364f430a86765b6bbd24462cd31f5a6d09bb74241aaf19190a25b919050565b610fcf613e78565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561401c5761401783614920565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015614076575060408051601f3d908101601f19168201909252614073918101906155e3565b60015b6140d95760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610ca0565b60008051602061572d83398151915281146141485760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610ca0565b506140178383836149bc565b60008060006141ee7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008c8c8c8b8b80519060200120604051602001610c2f969594939291909586526001600160a01b0394851660208701529290931660408501526060840152608083019190915260a082015260c00190565b3360009081526004602052604090205490915060ff166142265780600160405163c802b7f360e01b8152600401610ca0929190614e6d565b871580614231575086155b156142545780600260405163c802b7f360e01b8152600401610ca0929190614e6d565b6001600160a01b0389166142805780600360405163c802b7f360e01b8152600401610ca0929190614e6d565b6001600160a01b038a1660009081526009602052604090205460ff166142be5780600460405163c802b7f360e01b8152600401610ca0929190614e6d565b60008181526007602052604090205460ff16156142f35780600660405163c802b7f360e01b8152600401610ca0929190614e6d565b86881415801561430557506000198814155b156143285780600960405163c802b7f360e01b8152600401610ca0929190614e6d565b60006143348286613d55565b90506001600160a01b038116158061437d57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316145b156143a05781600760405163c802b7f360e01b8152600401610ca0929190614e6d565b6001600160a01b03808c166000818152600f60209081526040808320549486168352600d8252808320938352600390930190522054898212806143e257508981125b156144055783600860405163c802b7f360e01b8152600401610ca0929190614e6d565b8982039150898103905081600f60008f6001600160a01b03166001600160a01b031681526020019081526020016000208190555080600d6000856001600160a01b03166001600160a01b0316815260200190815260200160002060030160008f6001600160a01b03166001600160a01b031681526020019081526020016000208190555060016007600086815260200190815260200160002060006101000a81548160ff021916908315150217905550838d6001600160a01b0316846001600160a01b03167f66efaee033f9c338e1e0c145c8653403e5f05eddfa1e35f67931c6ba3e231ea08f8f8f8760405161452094939291906001600160a01b0394909416845260208401929092526040830152606082015260800190565b60405180910390a4919c919b50909950505050505050505050565b600060405163a9059cbb60e01b8152836004820152826024820152602060006044836000895af13d15601f3d11600160005114161716915050806145b35760405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b6044820152606401610ca0565b50505050565b60006001600160a01b0384166145e25760405163d92e233d60e01b815260040160405180910390fd5b8160000361460357604051631f2a200560e01b815260040160405180910390fd5b6001600160a01b03831660009081526009602052604090205460ff1661463c576040516368f7a67560e11b815260040160405180910390fd5b6001600160a01b038085166000908152600d6020908152604080832093871683526003909301905290812054614673908490615586565b6001600160a01b038087166000908152600d60209081526040808320938916835260039093018152828220849055600f9052908120805492935085929091906146bd908490615586565b909155505060408051848152602081018390526001600160a01b0380871692908816917f8cd158929f48bc4a1b49c8c3873d0561de29c5b993767bbe115997adeb1113cf910160405180910390a390505b9392505050565b60006040516323b872dd60e01b81528460048201528360248201528260448201526020600060648360008a5af13d15601f3d1160016000511416171691505080612b765760405162461bcd60e51b81526020600482015260146024820152731514905394d1915497d19493d357d1905253115160621b6044820152606401610ca0565b6000806000846000036147b8575050835460018501549091506000614917565b85546147c48682615586565b9350806000036147dd5760018701859055849250614912565b83600003614802576147f4818860010154876149e1565b600060018901559150614912565b6000811280156148125750600084135b8061482857506000811380156148285750600084125b1561484d5761483c818860010154876149e1565b600188018690558593509150614912565b60008113801561485c57508084135b1561489d578361486c8688615618565b600189015461487b9084615618565b6148859190615637565b61488f9190615665565b600188018190559250614912565b6000811280156148ac57508084125b156148e6576148ba846155c7565b856148c4886155c7565b6148ce9190615618565b60018901546148dc846155c7565b61487b9190615618565b60018701549250620f42406148fb868561547c565b6149059088615679565b61490f91906156fe565b91505b508286555b93509350939050565b6001600160a01b0381163b61498d5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610ca0565b60008051602061572d83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6149c583614a0d565b6000825111806149d25750805b15614017576145b38383614a4d565b6000620f42406149f1848461547c565b6149fb9086615679565b614a0591906156fe565b949350505050565b614a1681614920565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b606061470e838360405180606001604052806027815260200161574d6027913960606001600160a01b0384163b614ad55760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610ca0565b600080856001600160a01b031685604051614af091906155fc565b600060405180830381855af49150503d8060008114614b2b576040519150601f19603f3d011682016040523d82523d6000602084013e614b30565b606091505b5091509150614b40828286614b4a565b9695505050505050565b60608315614b5957508161470e565b825115614b695782518084602001fd5b8160405162461bcd60e51b8152600401610ca09190614bdb565b60005b83811015614b9e578181015183820152602001614b86565b838111156145b35750506000910152565b60008151808452614bc7816020860160208601614b83565b601f01601f19169290920160200192915050565b60208152600061470e6020830184614baf565b80356001600160a01b0381168114613fd757600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715614c3e57614c3e614c05565b60405290565b6040805190810167ffffffffffffffff81118282101715614c3e57614c3e614c05565b604051601f8201601f1916810167ffffffffffffffff81118282101715614c9057614c90614c05565b604052919050565b600082601f830112614ca957600080fd5b813567ffffffffffffffff811115614cc357614cc3614c05565b614cd6601f8201601f1916602001614c67565b818152846020838601011115614ceb57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060408385031215614d1b57600080fd5b614d2483614bee565b9150602083013567ffffffffffffffff811115614d4057600080fd5b614d4c85828601614c98565b9150509250929050565b600060208284031215614d6857600080fd5b61470e82614bee565b6020808252825182820181905260009190848201906040850190845b81811015614da957835183529284019291840191600101614d8d565b50909695505050505050565b600060208284031215614dc757600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015614da95783516001600160a01b031683529284019291840191600101614dea565b600080600060608486031215614e2457600080fd5b614e2d84614bee565b9250614e3b60208501614bee565b9150604084013590509250925092565b600d8110614e6957634e487b7160e01b600052602160045260246000fd5b9052565b8281526040810161470e6020830184614e4b565b60008060008060808587031215614e9757600080fd5b614ea085614bee565b935060208501359250604085013567ffffffffffffffff80821115614ec457600080fd5b614ed088838901614c98565b93506060870135915080821115614ee657600080fd5b50614ef387828801614c98565b91505092959194509250565b60008060408385031215614f1257600080fd5b614f1b83614bee565b946020939093013593505050565b60008060408385031215614f3c57600080fd5b614f4583614bee565b9150614f5360208401614bee565b90509250929050565b60008060408385031215614f6f57600080fd5b614f7883614bee565b915060208301358015158114614f8d57600080fd5b809150509250929050565b600080600080600080600060e0888a031215614fb357600080fd5b614fbc88614bee565b9650614fca60208901614bee565b955060408801359450606088013593506080880135925060a088013567ffffffffffffffff80821115614ffc57600080fd5b6150088b838c01614c98565b935060c08a013591508082111561501e57600080fd5b5061502b8a828b01614c98565b91505092959891949750929550565b6000806000806080858703121561505057600080fd5b61505985614bee565b966020860135965060408601359560600135945092505050565b8381526020810183905260608101614a056040830184614e4b565b602081016115de8284614e4b565b600067ffffffffffffffff8211156150b6576150b6614c05565b5060051b60200190565b600082601f8301126150d157600080fd5b813560206150e66150e18361509c565b614c67565b8281526060928302850182019282820191908785111561510557600080fd5b8387015b8581101561514c5781818a0312156151215760008081fd5b615129614c1b565b813581528582013586820152604080830135908201528452928401928101615109565b5090979650505050505050565b600082601f83011261516a57600080fd5b8135602061517a6150e18361509c565b82815260069290921b8401810191818101908684111561519957600080fd5b8286015b848110156151df57604081890312156151b65760008081fd5b6151be614c44565b6151c782614bee565b8152818501358582015283529183019160400161519d565b509695505050505050565b6000806000806080858703121561520057600080fd5b61520985614bee565b935060208086013567ffffffffffffffff8082111561522757600080fd5b818801915088601f83011261523b57600080fd5b81356152496150e18261509c565b81815260069190911b8301840190848101908b83111561526857600080fd5b938501935b828510156152b2576040858d0312156152865760008081fd5b61528e614c44565b61529786614bee565b8152858701358782015282526040909401939085019061526d565b9750505060408801359250808311156152ca57600080fd5b6152d689848a016150c0565b945060608801359250808311156152ec57600080fd5b5050614ef387828801615159565b600081518084526020808501945080840160005b83811015615340578151805188528381015184890152604090810151908801526060909601959082019060010161530e565b509495945050505050565b600081518084526020808501945080840160005b8381101561534057815180516001600160a01b03168852830151838801526040909601959082019060010161535f565b600060a082018715158352602060a08185015281885180845260c086019150828a01935060005b818110156153e657845180516001600160a01b0316845284015184840152938301936040909201916001016153b6565b505084810360408601526153fa81896152fa565b925050508281036060840152615410818661534b565b9150508260808301529695505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008282101561546157615461615439565b500390565b634e487b7160e01b600052603160045260246000fd5b60008083128015600160ff1b85018412161561549a5761549a615439565b6001600160ff1b03840183138116156154b5576154b5615439565b50500390565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090614b4090830184614baf565b600080821280156001600160ff1b03849003851316156155a8576155a8615439565b600160ff1b83900384128116156155c1576155c1615439565b50500190565b6000600160ff1b82016155dc576155dc615439565b5060000390565b6000602082840312156155f557600080fd5b5051919050565b6000825161560e818460208701614b83565b9190910192915050565b600081600019048311821515161561563257615632615439565b500290565b6000821982111561564a5761564a615439565b500190565b634e487b7160e01b600052601260045260246000fd5b6000826156745761567461564f565b500490565b60006001600160ff1b038184138284138082168684048611161561569f5761569f615439565b600160ff1b60008712828116878305891216156156be576156be615439565b600087129250878205871284841616156156da576156da615439565b878505871281841616156156f0576156f0615439565b505050929093029392505050565b60008261570d5761570d61564f565b600160ff1b82146000198414161561572757615727615439565b50059056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65647d551b4e977cb76b47865bfb6cdc77af264be2349f98f5d2cb58a453757caae7a164736f6c634300080f000a",
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

// LoadAccount is a paid mutator transaction binding the contract method 0x9d56921f.
//
// Solidity: function loadAccount(address account, (address,int256)[] balances, (uint256,uint256,int256)[] positions, (address,uint256)[] signingKeys) returns()
func (_Accounts *AccountsTransactor) LoadAccount(opts *bind.TransactOpts, account common.Address, balances []AccountsInterfaceBalance, positions []AccountsInterfacePosition, signingKeys []AccountsInterfaceSigningKey) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "loadAccount", account, balances, positions, signingKeys)
}

// LoadAccount is a paid mutator transaction binding the contract method 0x9d56921f.
//
// Solidity: function loadAccount(address account, (address,int256)[] balances, (uint256,uint256,int256)[] positions, (address,uint256)[] signingKeys) returns()
func (_Accounts *AccountsSession) LoadAccount(account common.Address, balances []AccountsInterfaceBalance, positions []AccountsInterfacePosition, signingKeys []AccountsInterfaceSigningKey) (*types.Transaction, error) {
	return _Accounts.Contract.LoadAccount(&_Accounts.TransactOpts, account, balances, positions, signingKeys)
}

// LoadAccount is a paid mutator transaction binding the contract method 0x9d56921f.
//
// Solidity: function loadAccount(address account, (address,int256)[] balances, (uint256,uint256,int256)[] positions, (address,uint256)[] signingKeys) returns()
func (_Accounts *AccountsTransactorSession) LoadAccount(account common.Address, balances []AccountsInterfaceBalance, positions []AccountsInterfacePosition, signingKeys []AccountsInterfaceSigningKey) (*types.Transaction, error) {
	return _Accounts.Contract.LoadAccount(&_Accounts.TransactOpts, account, balances, positions, signingKeys)
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
