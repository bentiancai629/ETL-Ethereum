// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchange

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExchangeCoreOrder is an auto generated low-level Go binding around an user-defined struct.
type ExchangeCoreOrder struct {
	Exchange           common.Address
	Maker              common.Address
	Taker              common.Address
	MakerArtistFee     *big.Int
	TakerArtistFee     *big.Int
	MakerProtocolFee   *big.Int
	TakerProtocolFee   *big.Int
	FeeRecipient       common.Address
	FeeMethod          uint8
	Side               uint8
	SaleKind           uint8
	Target             common.Address
	HowToCall          uint8
	Data               []byte
	ReplacementPattern []byte
	Brand              common.Address
	PaymentToken       common.Address
	BasePrice          *big.Int
	Extra              *big.Int
	ListingTime        *big.Int
	ExpirationTime     *big.Int
	Salt               *big.Int
}

// ExchangeCoreSig is an auto generated low-level Go binding around an user-defined struct.
type ExchangeCoreSig struct {
	V uint8
	R [32]byte
	S [32]byte
}

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"inputs\":[{\"internalType\":\"contractProxyRegistry\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"contractTokenTransferProxy\",\"name\":\"tokenTransferProxyAddress\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"protocolFeeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerRelayerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerRelayerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"OrderApprovedPartOne\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staticTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"orderbookInclusionDesired\",\"type\":\"bool\"}],\"name\":\"OrderApprovedPartTwo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"buyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INVERSE_BASIS_POINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"orderbookInclusionDesired\",\"type\":\"bool\"}],\"name\":\"approveOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedOrders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"buy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structExchangeCore.Sig\",\"name\":\"buySig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structExchangeCore.Sig\",\"name\":\"sellSig\",\"type\":\"tuple\"}],\"name\":\"atomicMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"calculateCurrentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"buy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"sell\",\"type\":\"tuple\"}],\"name\":\"calculateMatchPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structExchangeCore.Sig\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelledOrFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinimumMakerProtocolFee\",\"type\":\"uint256\"}],\"name\":\"changeMinimumMakerProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinimumTakerProtocolFee\",\"type\":\"uint256\"}],\"name\":\"changeMinimumTakerProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newProtocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"changeProtocolFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRoyaltyFee\",\"type\":\"uint256\"}],\"name\":\"changeRoyaltyFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sellData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"buyData\",\"type\":\"bytes\"}],\"name\":\"dataCanMatched\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashToSign\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sellData\",\"type\":\"bytes\"}],\"name\":\"isPackage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumMakerProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumTakerProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"buy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"sell\",\"type\":\"tuple\"}],\"name\":\"ordersCanMatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractProxyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extradata\",\"type\":\"bytes\"}],\"name\":\"staticCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenTransferProxy\",\"outputs\":[{\"internalType\":\"contractTokenTransferProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structExchangeCore.Sig\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"validateOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerArtistFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"enumExchangeCore.FeeMethod\",\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.Side\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSaleKindInterface.SaleKind\",\"name\":\"saleKind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"enumAuthenticatedProxy.HowToCall\",\"name\":\"howToCall\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"brand\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"basePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeCore.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrderParameters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ExchangeBin is the compiled bytecode used for deploying new contracts.
var ExchangeBin = "0x60806040523480156200001157600080fd5b5060405162004ebe38038062004ebe833981810160405281019062000037919062000243565b600033905080600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506200039b565b600081519050620001f88162000333565b92915050565b6000815190506200020f816200034d565b92915050565b600081519050620002268162000367565b92915050565b6000815190506200023d8162000381565b92915050565b600080600080608085870312156200025a57600080fd5b60006200026a8782880162000215565b94505060206200027d878288016200022c565b93505060406200029087828801620001fe565b9250506060620002a387828801620001e7565b91505092959194509250565b6000620002bc8262000313565b9050919050565b6000620002d08262000313565b9050919050565b6000620002e482620002af565b9050919050565b6000620002f882620002af565b9050919050565b60006200030c82620002af565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200033e81620002c3565b81146200034a57600080fd5b50565b6200035881620002d7565b81146200036457600080fd5b50565b6200037281620002eb565b81146200037e57600080fd5b50565b6200038c81620002ff565b81146200039857600080fd5b50565b614b1380620003ab6000396000f3fe6080604052600436106101bb5760003560e01c806364df049e116100ec578063a25eb5d91161008a578063cae6047f11610064578063cae6047f14610663578063e57d4adb1461068e578063e9147ade146106cb578063f2fde38b14610708576101c2565b8063a25eb5d9146105d2578063a51d9341146105fd578063a79101b61461063a576101c2565b80637771bc5f116100c65780637771bc5f146105025780637b1039991461053f5780637ccefc521461056a5780638076f00514610595576101c2565b806364df049e14610483578063715018a6146104ae57806371af1e32146104c5576101c2565b80631a6b13e2116101595780634222fd1a116101335780634222fd1a146103b5578063514f0330146103f2578063570ca7351461041b5780635fda3f2214610446576101c2565b80631a6b13e21461033857806328a8ee6814610361578063393a4bfd1461038c576101c2565b80630eefdbad116101955780630eefdbad1461026a57806310796a4714610295578063107a9e1f146102d257806314350c241461030f576101c2565b80630990a6f7146101c75780630a992d54146102045780630e7a8a9914610241576101c2565b366101c257005b600080fd5b3480156101d357600080fd5b506101ee60048036038101906101e99190613bfa565b610731565b6040516101fb919061418d565b60405180910390f35b34801561021057600080fd5b5061022b60048036038101906102269190613d37565b61084e565b60405161023891906143ed565b60405180910390f35b34801561024d57600080fd5b5061026860048036038101906102639190613b41565b610883565b005b34801561027657600080fd5b5061027f61092c565b60405161028c9190614275565b60405180910390f35b3480156102a157600080fd5b506102bc60048036038101906102b79190613a5b565b610952565b6040516102c9919061418d565b60405180910390f35b3480156102de57600080fd5b506102f960048036038101906102f49190613d37565b610a0f565b604051610306919061418d565b60405180910390f35b34801561031b57600080fd5b5061033660048036038101906103319190613e78565b610b23565b005b34801561034457600080fd5b5061035f600480360381019061035a9190613e78565b610bbd565b005b34801561036d57600080fd5b50610376610c57565b60405161038391906143ed565b60405180910390f35b34801561039857600080fd5b506103b360048036038101906103ae9190613ada565b610c5d565b005b3480156103c157600080fd5b506103dc60048036038101906103d79190613d78565b610e34565b6040516103e991906143ed565b60405180910390f35b3480156103fe57600080fd5b5061041960048036038101906104149190613a32565b610eed565b005b34801561042757600080fd5b50610430610fc1565b60405161043d9190614053565b60405180910390f35b34801561045257600080fd5b5061046d60048036038101906104689190613d78565b610fe7565b60405161047a919061418d565b60405180910390f35b34801561048f57600080fd5b50610498611492565b6040516104a59190614053565b60405180910390f35b3480156104ba57600080fd5b506104c36114b8565b005b3480156104d157600080fd5b506104ec60048036038101906104e79190613ca2565b611608565b6040516104f9919061418d565b60405180910390f35b34801561050e57600080fd5b5061052960048036038101906105249190613d37565b61180f565b60405161053691906141a8565b60405180910390f35b34801561054b57600080fd5b50610554611b0b565b604051610561919061425a565b60405180910390f35b34801561057657600080fd5b5061057f611b31565b60405161058c91906143ed565b60405180910390f35b3480156105a157600080fd5b506105bc60048036038101906105b79190613bd1565b611b37565b6040516105c9919061418d565b60405180910390f35b3480156105de57600080fd5b506105e7611b57565b6040516105f4919061423f565b60405180910390f35b34801561060957600080fd5b50610624600480360381019061061f9190613c61565b611b7d565b604051610631919061418d565b60405180910390f35b34801561064657600080fd5b50610661600480360381019061065c9190613de4565b611bf2565b005b34801561066f57600080fd5b5061067861229b565b60405161068591906143ed565b60405180910390f35b34801561069a57600080fd5b506106b560048036038101906106b09190613bd1565b6122a1565b6040516106c2919061418d565b60405180910390f35b3480156106d757600080fd5b506106f260048036038101906106ed9190613d37565b6122c1565b6040516106ff91906141a8565b60405180910390f35b34801561071457600080fd5b5061072f600480360381019061072a91906139e0565b6122f9565b005b600061073c83610a0f565b6107495760009050610847565b6005600085815260200190815260200160002060009054906101000a900460ff16156107785760009050610847565b6006600085815260200190815260200160002060009054906101000a900460ff16156107a75760019050610847565b826020015173ffffffffffffffffffffffffffffffffffffffff16600185846000015185602001518660400151604051600081526020016040526040516107f194939291906141fa565b6020604051602081039080840390855afa158015610813573d6000803e3d6000fd5b5050506020604051035173ffffffffffffffffffffffffffffffffffffffff1614156108425760019050610847565b600090505b9392505050565b600061087c826101200151836101400151846102400151856102600151866102800151876102a001516124b9565b9050919050565b600061088f838361268f565b9050826020015173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16146108cd57600080fd5b60016005600083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d60405160405180910390a250505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008082518451610963919061448b565b67ffffffffffffffff8111156109a2577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156109d45781602001600182028036833780820191505090505b50905060006020820190506109e981856126bb565b90506109f581866126bb565b506000604051835160208501895afa925050509392505050565b60003073ffffffffffffffffffffffffffffffffffffffff16826000015173ffffffffffffffffffffffffffffffffffffffff1614610a515760009050610b1e565b610a65826101400151836102a00151612707565b610a725760009050610b1e565b600180811115610aab577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8261010001516001811115610ae9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b148015610b0b57506007548260a001511080610b0a57506008548260c00151105b5b15610b195760009050610b1e565b600190505b919050565b3373ffffffffffffffffffffffffffffffffffffffff16600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610bb3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610baa906143ad565b60405180910390fd5b8060078190555050565b3373ffffffffffffffffffffffffffffffffffffffff16600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610c4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c44906143ad565b60405180910390fd5b8060088190555050565b60085481565b816020015173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610c9957600080fd5b6000610ca4836122c1565b90506006600082815260200190815260200160002060009054906101000a900460ff1615610cd157600080fd5b60016006600083815260200190815260200160002060006101000a81548160ff0219169083151502179055508260e0015173ffffffffffffffffffffffffffffffffffffffff16836020015173ffffffffffffffffffffffffffffffffffffffff16827f90c7f9f5b58c15f0f635bfb99f55d3d78fdbef3559e7d8abf5c81052a527662286600001518760400151886060015189608001518a60a001518b60c001518c61010001518d61012001518e61014001518f6101600151604051610da19a999897969594939291906140b3565b60405180910390a4807fe55393c778364e440d958b39ac1debd99dcfae3775a8a04d1e79124adf6a2d08846101800151856101a00151866101c00151876101e001518861020001518961022001518a61024001518b61026001518c61028001518d6102a001518e6102c001518e604051610e269c9b9a99989796959493929190614290565b60405180910390a250505050565b600080610e63836101200151846101400151856102400151866102600151876102800151886102a001516124b9565b90506000610e938561012001518661014001518761024001518861026001518961028001518a6102a001516124b9565b905081811015610ea257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168460e0015173ffffffffffffffffffffffffffffffffffffffff161415610ee15780610ee3565b815b9250505092915050565b3373ffffffffffffffffffffffffffffffffffffffff16600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610f7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f74906143ad565b60405180910390fd5b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806001811115611022577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8361012001516001811115611060577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1480156110e157506001808111156110a1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82610120015160018111156110df577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b145b801561116657508161010001516001811115611126577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8361010001516001811115611164577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b145b80156111a7575081610220015173ffffffffffffffffffffffffffffffffffffffff1683610220015173ffffffffffffffffffffffffffffffffffffffff16145b80156112215750600073ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff1614806112205750826020015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff16145b5b801561129b5750600073ffffffffffffffffffffffffffffffffffffffff16836040015173ffffffffffffffffffffffffffffffffffffffff16148061129a5750816020015173ffffffffffffffffffffffffffffffffffffffff16836040015173ffffffffffffffffffffffffffffffffffffffff16145b5b801561138c5750600073ffffffffffffffffffffffffffffffffffffffff168260e0015173ffffffffffffffffffffffffffffffffffffffff161480156113135750600073ffffffffffffffffffffffffffffffffffffffff168360e0015173ffffffffffffffffffffffffffffffffffffffff1614155b8061138b5750600073ffffffffffffffffffffffffffffffffffffffff168260e0015173ffffffffffffffffffffffffffffffffffffffff161415801561138a5750600073ffffffffffffffffffffffffffffffffffffffff168360e0015173ffffffffffffffffffffffffffffffffffffffff16145b5b5b80156113cd575081610160015173ffffffffffffffffffffffffffffffffffffffff1683610160015173ffffffffffffffffffffffffffffffffffffffff16145b801561145257508161018001516001811115611412577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8361018001516001811115611450577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b145b801561146e575061146d836102800151846102a0015161278f565b5b801561148a5750611489826102800151836102a0015161278f565b5b905092915050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff16600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611548576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161153f906143ad565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600080600090505b600481101561170757828181518110611652577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168482815181106116b8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146116f4576000915050611809565b80806116ff906147cc565b915050611610565b602490505b60408110156118035782818151811061174e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168482815181106117b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146117f0576000915050611809565b80806117fb906147cc565b91505061170c565b60019150505b92915050565b60008061181b836127b2565b905060008167ffffffffffffffff81111561185f577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156118915781602001600182028036833780820191505090505b50905060006020820190506118aa8186600001516127f1565b90506118ba8186602001516127f1565b90506118ca8186604001516127f1565b90506118da818660600151612824565b90506118ea818660800151612824565b90506118fa818660a00151612824565b905061190a818660c00151612824565b905061191a818660e001516127f1565b905061196381866101000151600181111561195e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b612838565b90506119ac8186610120015160018111156119a7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b612838565b90506119f58186610140015160018111156119f0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b612838565b9050611a06818661016001516127f1565b9050611a4f818661018001516001811115611a4a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b612838565b9050611a6081866101a001516126bb565b9050611a7181866101c001516126bb565b9050611a8281866101e001516127f1565b9050611a93818661020001516126bb565b9050611aa4818661022001516127f1565b9050611ab581866102400151612824565b9050611ac681866102600151612824565b9050611ad781866102800151612824565b9050611ae881866102a00151612824565b9050611af981866102c00151612824565b90508260208301209350505050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b60056020528060005260406000206000915054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060f81b82604381518110611bbd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60008054906101000a900460ff1615611c40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c37906143cd565b60405180910390fd5b60016000806101000a81548160ff0219169083151502179055506000611c66858561268f565b90506000611c74848461268f565b9050611c808685610fe7565b611c8957600080fd5b6000808561016001519050803b915060008211611ca557600080fd5b600080896101c001515114611d2a576000896101c00151511115611cde57611cdd896101a00151886101a001518b6101c0015161284c565b5b6000876101c00151511115611d0857611d07876101a001518a6101a00151896101c0015161284c565b5b611d1c896101a00151886101a00151612ae4565b611d2557600080fd5b611d67565b611d3e876101a001518a6101a00151611608565b611d4757600080fd5b611d55876101a00151611b7d565b9050886101a00151876101a001819052505b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c455279189602001516040518263ffffffff1660e01b8152600401611dc89190614038565b60206040518083038186803b158015611de057600080fd5b505afa158015611df4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e189190613d0e565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611e5457600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166397204d8e6040518163ffffffff1660e01b815260040160206040518083038186803b158015611ebc57600080fd5b505afa158015611ed0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ef49190613a09565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b815260040160206040518083038186803b158015611f5057600080fd5b505afa158015611f64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f889190613a09565b73ffffffffffffffffffffffffffffffffffffffff1614611fa857600080fd5b600081905060016005600089815260200190815260200160002060006101000a81548160ff0219169083151502179055508261200b5760016005600088815260200190815260200160002060006101000a81548160ff0219169083151502179055505b60006120178c8b612b51565b90508173ffffffffffffffffffffffffffffffffffffffff16631b0f7ba98b61016001518c61018001518d6101a001516040518463ffffffff1660e01b81526004016120659392919061414f565b602060405180830381600087803b15801561207f57600080fd5b505af1158015612093573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120b79190613ba8565b6120c057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168c6101e0015173ffffffffffffffffffffffffffffffffffffffff161461211d576121138c6101e001518b6101a001518e6102000151610952565b61211c57600080fd5b5b600073ffffffffffffffffffffffffffffffffffffffff168a6101e0015173ffffffffffffffffffffffffffffffffffffffff161461217a576121708a6101e001518b6101a001518c6102000151610952565b61217957600080fd5b5b600073ffffffffffffffffffffffffffffffffffffffff168a60e0015173ffffffffffffffffffffffffffffffffffffffff1614156121bd5789602001516121c3565b8b602001515b73ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff168b60e0015173ffffffffffffffffffffffffffffffffffffffff16141561221c578c60200151612222565b8a602001515b73ffffffffffffffffffffffffffffffffffffffff167f9a4fbbef9b6121513b138cfc80c3a8c19a221e65d5ef1c2b6a3e12cbeca9d6108a8a8560405161226b939291906141c3565b60405180910390a3505050505050505060008060006101000a81548160ff02191690831515021790555050505050565b61271081565b60066020528060005260406000206000915054906101000a900460ff1681565b60006122cc8261180f565b6040516020016122dc919061435f565b604051602081830303815290604052805190602001209050919050565b3373ffffffffffffffffffffffffffffffffffffffff16600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612389576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612380906143ad565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156123f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123f09061438d565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060018111156124f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b86600181111561252d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141561253b57849050612685565b600180811115612574577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8660018111156125ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141561268457600083836125c1919061456c565b84426125cd919061456c565b866125d89190614512565b6125e291906144e1565b905060018081111561261d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b886001811115612656577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415612670578086612668919061456c565b915050612685565b808661267c919061448b565b915050612685565b5b9695505050505050565b60008061269b846122c1565b90506126a8818585610731565b6126b157600080fd5b8091505092915050565b600080825111156126fe57815180602001830160208401855b600183831014156126f457815181526020820191506020810190506126d4565b8387019650505050505b82905092915050565b6000806001811115612742577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b83600181111561277b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14806127875750600082115b905092915050565b600042831080156127aa575060008214806127a957508142105b5b905092915050565b600081610200015151826101c0015151836101a00151516101b06127d6919061448b565b6127e0919061448b565b6127ea919061448b565b9050919050565b60008060608373ffffffffffffffffffffffffffffffffffffffff16901b90508084526014840193508391505092915050565b600081835260208301925082905092915050565b600081835360018301925082905092915050565b815183511461285a57600080fd5b805183511461286857600080fd5b60006020845161287891906144e1565b905060006020826128899190614512565b90508160208261289991906144e1565b146128cd577f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60005b8281101561290a57806001016020028085015181870151811682890151821916178289015250508080612902906147cc565b9150506128d0565b600083111561293a5782905080600101602002808501518187015181168289015182191617828901525050612adc565b8190505b8551811015612adb57848181518110612980577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b8482815181106129c4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b16868281518110612a09577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b60ff60f81b868481518110612a52577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b181617868281518110612a99577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080612ad3906147cc565b91505061293e565b5b505050505050565b6000806001905083518351811460018114612b025760009250612b45565b600160208701838101602088015b600284838510011415612b40578051835114612b2f5760009650600093505b602083019250602081019050612b10565b505050505b50508091505092915050565b600080479050600073ffffffffffffffffffffffffffffffffffffffff1683610220015173ffffffffffffffffffffffffffffffffffffffff1614612b9e5760008114612b9d57600080fd5b5b6000612baa8585610e34565b9050600081118015612bee5750600073ffffffffffffffffffffffffffffffffffffffff1684610220015173ffffffffffffffffffffffffffffffffffffffff1614155b15612c0d57612c0c84610220015186602001518660200151846134b0565b5b60008190506000829050600073ffffffffffffffffffffffffffffffffffffffff168660e0015173ffffffffffffffffffffffffffffffffffffffff161461312f57866080015186608001511115612c6457600080fd5b600180811115612c9d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8661010001516001811115612cdb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156130fb578660c001518660c001511115612cf657600080fd5b600086606001511115612dd8576000612710848860600151612d189190614512565b612d2291906144e1565b9050600073ffffffffffffffffffffffffffffffffffffffff1687610220015173ffffffffffffffffffffffffffffffffffffffff161415612dbc578083612d6a919061456c565b92508660e0015173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015612db6573d6000803e3d6000fd5b50612dd6565b612dd587610220015188602001518960e00151846134b0565b5b505b600086608001511115612eba576000612710848860800151612dfa9190614512565b612e0491906144e1565b9050600073ffffffffffffffffffffffffffffffffffffffff1687610220015173ffffffffffffffffffffffffffffffffffffffff161415612e9e578082612e4c919061448b565b91508660e0015173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015612e98573d6000803e3d6000fd5b50612eb8565b612eb787610220015189602001518960e00151846134b0565b5b505b60008660a001511115612fd8576000612710848860a00151612edc9190614512565b612ee691906144e1565b9050600073ffffffffffffffffffffffffffffffffffffffff1687610220015173ffffffffffffffffffffffffffffffffffffffff161415612f9e578083612f2e919061456c565b9250600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015612f98573d6000803e3d6000fd5b50612fd6565b612fd58761022001518860200151600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846134b0565b5b505b60008660c0015111156130f6576000612710848860c00151612ffa9190614512565b61300491906144e1565b9050600073ffffffffffffffffffffffffffffffffffffffff1687610220015173ffffffffffffffffffffffffffffffffffffffff1614156130bc57808261304c919061448b565b9150600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156130b6573d6000803e3d6000fd5b506130f4565b6130f38761022001518960200151600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846134b0565b5b505b61312a565b61311286602001518760e00151886060015161357c565b61312987602001518760e00151886080015161357c565b5b6133aa565b85608001518760800151111561314457600080fd5b60018081111561317d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b86610100015160018111156131bb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141561337a57600073ffffffffffffffffffffffffffffffffffffffff1686610220015173ffffffffffffffffffffffffffffffffffffffff16141561320057600080fd5b8560c001518760c00151111561321557600080fd5b60008760600151111561325e5760006127108489606001516132379190614512565b61324191906144e1565b905061325c87610220015189602001518a60e00151846134b0565b505b6000876080015111156132a75760006127108489608001516132809190614512565b61328a91906144e1565b90506132a587610220015188602001518a60e00151846134b0565b505b60008760a00151111561330e576000612710848960a001516132c99190614512565b6132d391906144e1565b905061330c8761022001518960200151600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846134b0565b505b60008760c001511115613375576000612710848960c001516133309190614512565b61333a91906144e1565b90506133738761022001518860200151600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846134b0565b505b6133a9565b61339187602001518860e00151896060015161357c565b6133a886602001518860e00151896080015161357c565b5b5b600073ffffffffffffffffffffffffffffffffffffffff1686610220015173ffffffffffffffffffffffffffffffffffffffff1614156134a357808410156133f157600080fd5b856020015173ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f1935050505015801561343b573d6000803e3d6000fd5b506000818561344a919061456c565b905060008111156134a157876020015173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561349f573d6000803e3d6000fd5b505b505b8294505050505092915050565b600081111561357657600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315dacbea858585856040518563ffffffff1660e01b815260040161351a949392919061406e565b602060405180830381600087803b15801561353457600080fd5b505af1158015613548573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061356c9190613ba8565b61357557600080fd5b5b50505050565b6135aa600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168484846134b0565b505050565b60006135c26135bd8461442d565b614408565b9050828152602081018484840111156135da57600080fd5b6135e5848285614759565b509392505050565b6000813590506135fc816149fc565b92915050565b600081519050613611816149fc565b92915050565b60008135905061362681614a13565b92915050565b60008135905061363b81614a2a565b92915050565b60008151905061365081614a2a565b92915050565b60008135905061366581614a41565b92915050565b600082601f83011261367c57600080fd5b813561368c8482602086016135af565b91505092915050565b6000815190506136a481614a58565b92915050565b6000813590506136b981614a6f565b92915050565b6000813590506136ce81614a7f565b92915050565b6000813590506136e381614a8f565b92915050565b6000813590506136f881614a9f565b92915050565b60006102e0828403121561371157600080fd5b61371c6102e0614408565b9050600061372c848285016135ed565b600083015250602061374084828501613617565b6020830152506040613754848285016135ed565b6040830152506060613768848285016139b6565b606083015250608061377c848285016139b6565b60808301525060a0613790848285016139b6565b60a08301525060c06137a4848285016139b6565b60c08301525060e06137b884828501613617565b60e0830152506101006137cd848285016136aa565b610100830152506101206137e3848285016136e9565b610120830152506101406137f9848285016136d4565b6101408301525061016061380f848285016135ed565b61016083015250610180613825848285016136bf565b610180830152506101a082013567ffffffffffffffff81111561384757600080fd5b6138538482850161366b565b6101a0830152506101c082013567ffffffffffffffff81111561387557600080fd5b6138818482850161366b565b6101c0830152506101e0613897848285016135ed565b6101e08301525061020082013567ffffffffffffffff8111156138b957600080fd5b6138c58482850161366b565b610200830152506102206138db848285016135ed565b610220830152506102406138f1848285016139b6565b61024083015250610260613907848285016139b6565b6102608301525061028061391d848285016139b6565b610280830152506102a0613933848285016139b6565b6102a0830152506102c0613949848285016139b6565b6102c08301525092915050565b60006060828403121561396857600080fd5b6139726060614408565b90506000613982848285016139cb565b600083015250602061399684828501613656565b60208301525060406139aa84828501613656565b60408301525092915050565b6000813590506139c581614aaf565b92915050565b6000813590506139da81614ac6565b92915050565b6000602082840312156139f257600080fd5b6000613a00848285016135ed565b91505092915050565b600060208284031215613a1b57600080fd5b6000613a2984828501613602565b91505092915050565b600060208284031215613a4457600080fd5b6000613a5284828501613617565b91505092915050565b600080600060608486031215613a7057600080fd5b6000613a7e868287016135ed565b935050602084013567ffffffffffffffff811115613a9b57600080fd5b613aa78682870161366b565b925050604084013567ffffffffffffffff811115613ac457600080fd5b613ad08682870161366b565b9150509250925092565b600080600060608486031215613aef57600080fd5b6000613afd868287016135ed565b935050602084013567ffffffffffffffff811115613b1a57600080fd5b613b26868287016136fe565b9250506040613b378682870161362c565b9150509250925092565b600080600060a08486031215613b5657600080fd5b6000613b64868287016135ed565b935050602084013567ffffffffffffffff811115613b8157600080fd5b613b8d868287016136fe565b9250506040613b9e86828701613956565b9150509250925092565b600060208284031215613bba57600080fd5b6000613bc884828501613641565b91505092915050565b600060208284031215613be357600080fd5b6000613bf184828501613656565b91505092915050565b600080600060a08486031215613c0f57600080fd5b6000613c1d86828701613656565b935050602084013567ffffffffffffffff811115613c3a57600080fd5b613c46868287016136fe565b9250506040613c5786828701613956565b9150509250925092565b600060208284031215613c7357600080fd5b600082013567ffffffffffffffff811115613c8d57600080fd5b613c998482850161366b565b91505092915050565b60008060408385031215613cb557600080fd5b600083013567ffffffffffffffff811115613ccf57600080fd5b613cdb8582860161366b565b925050602083013567ffffffffffffffff811115613cf857600080fd5b613d048582860161366b565b9150509250929050565b600060208284031215613d2057600080fd5b6000613d2e84828501613695565b91505092915050565b600060208284031215613d4957600080fd5b600082013567ffffffffffffffff811115613d6357600080fd5b613d6f848285016136fe565b91505092915050565b60008060408385031215613d8b57600080fd5b600083013567ffffffffffffffff811115613da557600080fd5b613db1858286016136fe565b925050602083013567ffffffffffffffff811115613dce57600080fd5b613dda858286016136fe565b9150509250929050565b6000806000806101008587031215613dfb57600080fd5b600085013567ffffffffffffffff811115613e1557600080fd5b613e21878288016136fe565b9450506020613e3287828801613956565b935050608085013567ffffffffffffffff811115613e4f57600080fd5b613e5b878288016136fe565b92505060a0613e6c87828801613956565b91505092959194509250565b600060208284031215613e8a57600080fd5b6000613e98848285016139b6565b91505092915050565b613eaa8161466f565b82525050565b613eb9816145b2565b82525050565b613ec8816145a0565b82525050565b613ed7816145c4565b82525050565b613ee6816145d0565b82525050565b6000613ef78261445e565b613f018185614469565b9350613f11818560208601614768565b613f1a816148d1565b840191505092915050565b613f2e81614681565b82525050565b613f3d816146a5565b82525050565b613f4c816146c9565b82525050565b613f5b816146ed565b82525050565b613f6a816146ff565b82525050565b613f7981614711565b82525050565b613f8881614723565b82525050565b6000613f9b601c8361447a565b9150613fa6826148e2565b602082019050919050565b6000613fbe60268361447a565b9150613fc98261490b565b604082019050919050565b6000613fe160208361447a565b9150613fec8261495a565b602082019050919050565b6000614004601f8361447a565b915061400f82614983565b602082019050919050565b61402381614658565b82525050565b61403281614662565b82525050565b600060208201905061404d6000830184613ea1565b92915050565b60006020820190506140686000830184613eb0565b92915050565b60006080820190506140836000830187613ebf565b6140906020830186613ebf565b61409d6040830185613ebf565b6140aa606083018461401a565b95945050505050565b6000610140820190506140c9600083018d613ebf565b6140d6602083018c613ebf565b6140e3604083018b61401a565b6140f0606083018a61401a565b6140fd608083018961401a565b61410a60a083018861401a565b61411760c0830187613f52565b61412460e0830186613f7f565b614132610100830185613f70565b614140610120830184613ebf565b9b9a5050505050505050505050565b60006060820190506141646000830186613ebf565b6141716020830185613f61565b81810360408301526141838184613eec565b9050949350505050565b60006020820190506141a26000830184613ece565b92915050565b60006020820190506141bd6000830184613edd565b92915050565b60006060820190506141d86000830186613edd565b6141e56020830185613edd565b6141f2604083018461401a565b949350505050565b600060808201905061420f6000830187613edd565b61421c6020830186614029565b6142296040830185613edd565b6142366060830184613edd565b95945050505050565b60006020820190506142546000830184613f25565b92915050565b600060208201905061426f6000830184613f34565b92915050565b600060208201905061428a6000830184613f43565b92915050565b6000610180820190506142a6600083018f613f61565b81810360208301526142b8818e613eec565b905081810360408301526142cc818d613eec565b90506142db606083018c613ebf565b81810360808301526142ed818b613eec565b90506142fc60a083018a613ebf565b61430960c083018961401a565b61431660e083018861401a565b61432461010083018761401a565b61433261012083018661401a565b61434061014083018561401a565b61434e610160830184613ece565b9d9c50505050505050505050505050565b6000604082019050818103600083015261437881613f8e565b90506143876020830184613edd565b92915050565b600060208201905081810360008301526143a681613fb1565b9050919050565b600060208201905081810360008301526143c681613fd4565b9050919050565b600060208201905081810360008301526143e681613ff7565b9050919050565b6000602082019050614402600083018461401a565b92915050565b6000614412614423565b905061441e828261479b565b919050565b6000604051905090565b600067ffffffffffffffff821115614448576144476148a2565b5b614451826148d1565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600061449682614658565b91506144a183614658565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156144d6576144d5614815565b5b828201905092915050565b60006144ec82614658565b91506144f783614658565b92508261450757614506614844565b5b828204905092915050565b600061451d82614658565b915061452883614658565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561456157614560614815565b5b828202905092915050565b600061457782614658565b915061458283614658565b92508282101561459557614594614815565b5b828203905092915050565b60006145ab82614638565b9050919050565b60006145bd82614638565b9050919050565b60008115159050919050565b6000819050919050565b60006145e5826145a0565b9050919050565b60008190506145fa826149ac565b919050565b600081905061460d826149c0565b919050565b6000819050614620826149d4565b919050565b6000819050614633826149e8565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b600061467a82614735565b9050919050565b600061468c82614693565b9050919050565b600061469e82614638565b9050919050565b60006146b0826146b7565b9050919050565b60006146c282614638565b9050919050565b60006146d4826146db565b9050919050565b60006146e682614638565b9050919050565b60006146f8826145ec565b9050919050565b600061470a826145ff565b9050919050565b600061471c82614612565b9050919050565b600061472e82614625565b9050919050565b600061474082614747565b9050919050565b600061475282614638565b9050919050565b82818337600083830152505050565b60005b8381101561478657808201518184015260208101905061476b565b83811115614795576000848401525b50505050565b6147a4826148d1565b810181811067ffffffffffffffff821117156147c3576147c26148a2565b5b80604052505050565b60006147d782614658565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561480a57614809614815565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b600281106149bd576149bc614873565b5b50565b600281106149d1576149d0614873565b5b50565b600281106149e5576149e4614873565b5b50565b600281106149f9576149f8614873565b5b50565b614a05816145a0565b8114614a1057600080fd5b50565b614a1c816145b2565b8114614a2757600080fd5b50565b614a33816145c4565b8114614a3e57600080fd5b50565b614a4a816145d0565b8114614a5557600080fd5b50565b614a61816145da565b8114614a6c57600080fd5b50565b60028110614a7c57600080fd5b50565b60028110614a8c57600080fd5b50565b60028110614a9c57600080fd5b50565b60028110614aac57600080fd5b50565b614ab881614658565b8114614ac357600080fd5b50565b614acf81614662565b8114614ada57600080fd5b5056fea264697066735822122030f64bd4196fa23aaa46d12cdd4d40722af375c52d81d3922f80e2e5f23f314f64736f6c63430008040033"

// DeployExchange deploys a new Ethereum contract, binding an instance of Exchange to it.
func DeployExchange(auth *bind.TransactOpts, backend bind.ContractBackend, registryAddress common.Address, tokenTransferProxyAddress common.Address, tokenAddress common.Address, protocolFeeAddress common.Address) (common.Address, *types.Transaction, *Exchange, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExchangeBin), backend, registryAddress, tokenTransferProxyAddress, tokenAddress, protocolFeeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Exchange *ExchangeCaller) INVERSEBASISPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "INVERSE_BASIS_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Exchange *ExchangeSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Exchange.Contract.INVERSEBASISPOINT(&_Exchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Exchange *ExchangeCallerSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Exchange.Contract.INVERSEBASISPOINT(&_Exchange.CallOpts)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 ) view returns(bool)
func (_Exchange *ExchangeCaller) ApprovedOrders(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "approvedOrders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 ) view returns(bool)
func (_Exchange *ExchangeSession) ApprovedOrders(arg0 [32]byte) (bool, error) {
	return _Exchange.Contract.ApprovedOrders(&_Exchange.CallOpts, arg0)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 ) view returns(bool)
func (_Exchange *ExchangeCallerSession) ApprovedOrders(arg0 [32]byte) (bool, error) {
	return _Exchange.Contract.ApprovedOrders(&_Exchange.CallOpts, arg0)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x5700e569.
//
// Solidity: function calculateCurrentPrice((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) view returns(uint256)
func (_Exchange *ExchangeCaller) CalculateCurrentPrice(opts *bind.CallOpts, order ExchangeCoreOrder) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "calculateCurrentPrice", order)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x5700e569.
//
// Solidity: function calculateCurrentPrice((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) view returns(uint256)
func (_Exchange *ExchangeSession) CalculateCurrentPrice(order ExchangeCoreOrder) (*big.Int, error) {
	return _Exchange.Contract.CalculateCurrentPrice(&_Exchange.CallOpts, order)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x5700e569.
//
// Solidity: function calculateCurrentPrice((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) view returns(uint256)
func (_Exchange *ExchangeCallerSession) CalculateCurrentPrice(order ExchangeCoreOrder) (*big.Int, error) {
	return _Exchange.Contract.CalculateCurrentPrice(&_Exchange.CallOpts, order)
}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0xca238615.
//
// Solidity: function calculateMatchPrice((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) buy, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) sell) view returns(uint256)
func (_Exchange *ExchangeCaller) CalculateMatchPrice(opts *bind.CallOpts, buy ExchangeCoreOrder, sell ExchangeCoreOrder) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "calculateMatchPrice", buy, sell)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0xca238615.
//
// Solidity: function calculateMatchPrice((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) buy, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) sell) view returns(uint256)
func (_Exchange *ExchangeSession) CalculateMatchPrice(buy ExchangeCoreOrder, sell ExchangeCoreOrder) (*big.Int, error) {
	return _Exchange.Contract.CalculateMatchPrice(&_Exchange.CallOpts, buy, sell)
}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0xca238615.
//
// Solidity: function calculateMatchPrice((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) buy, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) sell) view returns(uint256)
func (_Exchange *ExchangeCallerSession) CalculateMatchPrice(buy ExchangeCoreOrder, sell ExchangeCoreOrder) (*big.Int, error) {
	return _Exchange.Contract.CalculateMatchPrice(&_Exchange.CallOpts, buy, sell)
}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) view returns(bool)
func (_Exchange *ExchangeCaller) CancelledOrFinalized(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "cancelledOrFinalized", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) view returns(bool)
func (_Exchange *ExchangeSession) CancelledOrFinalized(arg0 [32]byte) (bool, error) {
	return _Exchange.Contract.CancelledOrFinalized(&_Exchange.CallOpts, arg0)
}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) view returns(bool)
func (_Exchange *ExchangeCallerSession) CancelledOrFinalized(arg0 [32]byte) (bool, error) {
	return _Exchange.Contract.CancelledOrFinalized(&_Exchange.CallOpts, arg0)
}

// DataCanMatched is a free data retrieval call binding the contract method 0x71af1e32.
//
// Solidity: function dataCanMatched(bytes sellData, bytes buyData) pure returns(bool)
func (_Exchange *ExchangeCaller) DataCanMatched(opts *bind.CallOpts, sellData []byte, buyData []byte) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "dataCanMatched", sellData, buyData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DataCanMatched is a free data retrieval call binding the contract method 0x71af1e32.
//
// Solidity: function dataCanMatched(bytes sellData, bytes buyData) pure returns(bool)
func (_Exchange *ExchangeSession) DataCanMatched(sellData []byte, buyData []byte) (bool, error) {
	return _Exchange.Contract.DataCanMatched(&_Exchange.CallOpts, sellData, buyData)
}

// DataCanMatched is a free data retrieval call binding the contract method 0x71af1e32.
//
// Solidity: function dataCanMatched(bytes sellData, bytes buyData) pure returns(bool)
func (_Exchange *ExchangeCallerSession) DataCanMatched(sellData []byte, buyData []byte) (bool, error) {
	return _Exchange.Contract.DataCanMatched(&_Exchange.CallOpts, sellData, buyData)
}

// ExchangeToken is a free data retrieval call binding the contract method 0xa25eb5d9.
//
// Solidity: function exchangeToken() view returns(address)
func (_Exchange *ExchangeCaller) ExchangeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "exchangeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeToken is a free data retrieval call binding the contract method 0xa25eb5d9.
//
// Solidity: function exchangeToken() view returns(address)
func (_Exchange *ExchangeSession) ExchangeToken() (common.Address, error) {
	return _Exchange.Contract.ExchangeToken(&_Exchange.CallOpts)
}

// ExchangeToken is a free data retrieval call binding the contract method 0xa25eb5d9.
//
// Solidity: function exchangeToken() view returns(address)
func (_Exchange *ExchangeCallerSession) ExchangeToken() (common.Address, error) {
	return _Exchange.Contract.ExchangeToken(&_Exchange.CallOpts)
}

// HashOrder is a free data retrieval call binding the contract method 0x2398a426.
//
// Solidity: function hashOrder((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) pure returns(bytes32 hash)
func (_Exchange *ExchangeCaller) HashOrder(opts *bind.CallOpts, order ExchangeCoreOrder) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x2398a426.
//
// Solidity: function hashOrder((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) pure returns(bytes32 hash)
func (_Exchange *ExchangeSession) HashOrder(order ExchangeCoreOrder) ([32]byte, error) {
	return _Exchange.Contract.HashOrder(&_Exchange.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x2398a426.
//
// Solidity: function hashOrder((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) pure returns(bytes32 hash)
func (_Exchange *ExchangeCallerSession) HashOrder(order ExchangeCoreOrder) ([32]byte, error) {
	return _Exchange.Contract.HashOrder(&_Exchange.CallOpts, order)
}

// HashToSign is a free data retrieval call binding the contract method 0x929e7ea2.
//
// Solidity: function hashToSign((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) pure returns(bytes32)
func (_Exchange *ExchangeCaller) HashToSign(opts *bind.CallOpts, order ExchangeCoreOrder) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "hashToSign", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashToSign is a free data retrieval call binding the contract method 0x929e7ea2.
//
// Solidity: function hashToSign((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) pure returns(bytes32)
func (_Exchange *ExchangeSession) HashToSign(order ExchangeCoreOrder) ([32]byte, error) {
	return _Exchange.Contract.HashToSign(&_Exchange.CallOpts, order)
}

// HashToSign is a free data retrieval call binding the contract method 0x929e7ea2.
//
// Solidity: function hashToSign((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) pure returns(bytes32)
func (_Exchange *ExchangeCallerSession) HashToSign(order ExchangeCoreOrder) ([32]byte, error) {
	return _Exchange.Contract.HashToSign(&_Exchange.CallOpts, order)
}

// IsPackage is a free data retrieval call binding the contract method 0xa51d9341.
//
// Solidity: function isPackage(bytes sellData) pure returns(bool)
func (_Exchange *ExchangeCaller) IsPackage(opts *bind.CallOpts, sellData []byte) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isPackage", sellData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPackage is a free data retrieval call binding the contract method 0xa51d9341.
//
// Solidity: function isPackage(bytes sellData) pure returns(bool)
func (_Exchange *ExchangeSession) IsPackage(sellData []byte) (bool, error) {
	return _Exchange.Contract.IsPackage(&_Exchange.CallOpts, sellData)
}

// IsPackage is a free data retrieval call binding the contract method 0xa51d9341.
//
// Solidity: function isPackage(bytes sellData) pure returns(bool)
func (_Exchange *ExchangeCallerSession) IsPackage(sellData []byte) (bool, error) {
	return _Exchange.Contract.IsPackage(&_Exchange.CallOpts, sellData)
}

// MinimumMakerProtocolFee is a free data retrieval call binding the contract method 0x7ccefc52.
//
// Solidity: function minimumMakerProtocolFee() view returns(uint256)
func (_Exchange *ExchangeCaller) MinimumMakerProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "minimumMakerProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumMakerProtocolFee is a free data retrieval call binding the contract method 0x7ccefc52.
//
// Solidity: function minimumMakerProtocolFee() view returns(uint256)
func (_Exchange *ExchangeSession) MinimumMakerProtocolFee() (*big.Int, error) {
	return _Exchange.Contract.MinimumMakerProtocolFee(&_Exchange.CallOpts)
}

// MinimumMakerProtocolFee is a free data retrieval call binding the contract method 0x7ccefc52.
//
// Solidity: function minimumMakerProtocolFee() view returns(uint256)
func (_Exchange *ExchangeCallerSession) MinimumMakerProtocolFee() (*big.Int, error) {
	return _Exchange.Contract.MinimumMakerProtocolFee(&_Exchange.CallOpts)
}

// MinimumTakerProtocolFee is a free data retrieval call binding the contract method 0x28a8ee68.
//
// Solidity: function minimumTakerProtocolFee() view returns(uint256)
func (_Exchange *ExchangeCaller) MinimumTakerProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "minimumTakerProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumTakerProtocolFee is a free data retrieval call binding the contract method 0x28a8ee68.
//
// Solidity: function minimumTakerProtocolFee() view returns(uint256)
func (_Exchange *ExchangeSession) MinimumTakerProtocolFee() (*big.Int, error) {
	return _Exchange.Contract.MinimumTakerProtocolFee(&_Exchange.CallOpts)
}

// MinimumTakerProtocolFee is a free data retrieval call binding the contract method 0x28a8ee68.
//
// Solidity: function minimumTakerProtocolFee() view returns(uint256)
func (_Exchange *ExchangeCallerSession) MinimumTakerProtocolFee() (*big.Int, error) {
	return _Exchange.Contract.MinimumTakerProtocolFee(&_Exchange.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Exchange *ExchangeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Exchange *ExchangeSession) Operator() (common.Address, error) {
	return _Exchange.Contract.Operator(&_Exchange.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Exchange *ExchangeCallerSession) Operator() (common.Address, error) {
	return _Exchange.Contract.Operator(&_Exchange.CallOpts)
}

// OrdersCanMatch is a free data retrieval call binding the contract method 0xba435100.
//
// Solidity: function ordersCanMatch((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) buy, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) sell) view returns(bool)
func (_Exchange *ExchangeCaller) OrdersCanMatch(opts *bind.CallOpts, buy ExchangeCoreOrder, sell ExchangeCoreOrder) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "ordersCanMatch", buy, sell)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrdersCanMatch is a free data retrieval call binding the contract method 0xba435100.
//
// Solidity: function ordersCanMatch((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) buy, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) sell) view returns(bool)
func (_Exchange *ExchangeSession) OrdersCanMatch(buy ExchangeCoreOrder, sell ExchangeCoreOrder) (bool, error) {
	return _Exchange.Contract.OrdersCanMatch(&_Exchange.CallOpts, buy, sell)
}

// OrdersCanMatch is a free data retrieval call binding the contract method 0xba435100.
//
// Solidity: function ordersCanMatch((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) buy, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) sell) view returns(bool)
func (_Exchange *ExchangeCallerSession) OrdersCanMatch(buy ExchangeCoreOrder, sell ExchangeCoreOrder) (bool, error) {
	return _Exchange.Contract.OrdersCanMatch(&_Exchange.CallOpts, buy, sell)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCallerSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Exchange *ExchangeCaller) ProtocolFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "protocolFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Exchange *ExchangeSession) ProtocolFeeRecipient() (common.Address, error) {
	return _Exchange.Contract.ProtocolFeeRecipient(&_Exchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Exchange *ExchangeCallerSession) ProtocolFeeRecipient() (common.Address, error) {
	return _Exchange.Contract.ProtocolFeeRecipient(&_Exchange.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Exchange *ExchangeCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Exchange *ExchangeSession) Registry() (common.Address, error) {
	return _Exchange.Contract.Registry(&_Exchange.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Exchange *ExchangeCallerSession) Registry() (common.Address, error) {
	return _Exchange.Contract.Registry(&_Exchange.CallOpts)
}

// RoyaltyFee is a free data retrieval call binding the contract method 0xb8997a97.
//
// Solidity: function royaltyFee() view returns(uint256)
func (_Exchange *ExchangeCaller) RoyaltyFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "royaltyFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyFee is a free data retrieval call binding the contract method 0xb8997a97.
//
// Solidity: function royaltyFee() view returns(uint256)
func (_Exchange *ExchangeSession) RoyaltyFee() (*big.Int, error) {
	return _Exchange.Contract.RoyaltyFee(&_Exchange.CallOpts)
}

// RoyaltyFee is a free data retrieval call binding the contract method 0xb8997a97.
//
// Solidity: function royaltyFee() view returns(uint256)
func (_Exchange *ExchangeCallerSession) RoyaltyFee() (*big.Int, error) {
	return _Exchange.Contract.RoyaltyFee(&_Exchange.CallOpts)
}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes data, bytes extradata) view returns(bool result)
func (_Exchange *ExchangeCaller) StaticCall(opts *bind.CallOpts, target common.Address, data []byte, extradata []byte) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "staticCall", target, data, extradata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes data, bytes extradata) view returns(bool result)
func (_Exchange *ExchangeSession) StaticCall(target common.Address, data []byte, extradata []byte) (bool, error) {
	return _Exchange.Contract.StaticCall(&_Exchange.CallOpts, target, data, extradata)
}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes data, bytes extradata) view returns(bool result)
func (_Exchange *ExchangeCallerSession) StaticCall(target common.Address, data []byte, extradata []byte) (bool, error) {
	return _Exchange.Contract.StaticCall(&_Exchange.CallOpts, target, data, extradata)
}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() view returns(address)
func (_Exchange *ExchangeCaller) TokenTransferProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "tokenTransferProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() view returns(address)
func (_Exchange *ExchangeSession) TokenTransferProxy() (common.Address, error) {
	return _Exchange.Contract.TokenTransferProxy(&_Exchange.CallOpts)
}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() view returns(address)
func (_Exchange *ExchangeCallerSession) TokenTransferProxy() (common.Address, error) {
	return _Exchange.Contract.TokenTransferProxy(&_Exchange.CallOpts)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x44fccb7c.
//
// Solidity: function validateOrder(bytes32 hash, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) view returns(bool)
func (_Exchange *ExchangeCaller) ValidateOrder(opts *bind.CallOpts, hash [32]byte, order ExchangeCoreOrder, sig ExchangeCoreSig) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "validateOrder", hash, order, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOrder is a free data retrieval call binding the contract method 0x44fccb7c.
//
// Solidity: function validateOrder(bytes32 hash, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) view returns(bool)
func (_Exchange *ExchangeSession) ValidateOrder(hash [32]byte, order ExchangeCoreOrder, sig ExchangeCoreSig) (bool, error) {
	return _Exchange.Contract.ValidateOrder(&_Exchange.CallOpts, hash, order, sig)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x44fccb7c.
//
// Solidity: function validateOrder(bytes32 hash, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) view returns(bool)
func (_Exchange *ExchangeCallerSession) ValidateOrder(hash [32]byte, order ExchangeCoreOrder, sig ExchangeCoreSig) (bool, error) {
	return _Exchange.Contract.ValidateOrder(&_Exchange.CallOpts, hash, order, sig)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0x933e2fab.
//
// Solidity: function validateOrderParameters((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) view returns(bool)
func (_Exchange *ExchangeCaller) ValidateOrderParameters(opts *bind.CallOpts, order ExchangeCoreOrder) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "validateOrderParameters", order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0x933e2fab.
//
// Solidity: function validateOrderParameters((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) view returns(bool)
func (_Exchange *ExchangeSession) ValidateOrderParameters(order ExchangeCoreOrder) (bool, error) {
	return _Exchange.Contract.ValidateOrderParameters(&_Exchange.CallOpts, order)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0x933e2fab.
//
// Solidity: function validateOrderParameters((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order) view returns(bool)
func (_Exchange *ExchangeCallerSession) ValidateOrderParameters(order ExchangeCoreOrder) (bool, error) {
	return _Exchange.Contract.ValidateOrderParameters(&_Exchange.CallOpts, order)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x8cd0abac.
//
// Solidity: function approveOrder(address sender, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order, bool orderbookInclusionDesired) returns()
func (_Exchange *ExchangeTransactor) ApproveOrder(opts *bind.TransactOpts, sender common.Address, order ExchangeCoreOrder, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "approveOrder", sender, order, orderbookInclusionDesired)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x8cd0abac.
//
// Solidity: function approveOrder(address sender, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order, bool orderbookInclusionDesired) returns()
func (_Exchange *ExchangeSession) ApproveOrder(sender common.Address, order ExchangeCoreOrder, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _Exchange.Contract.ApproveOrder(&_Exchange.TransactOpts, sender, order, orderbookInclusionDesired)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x8cd0abac.
//
// Solidity: function approveOrder(address sender, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order, bool orderbookInclusionDesired) returns()
func (_Exchange *ExchangeTransactorSession) ApproveOrder(sender common.Address, order ExchangeCoreOrder, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _Exchange.Contract.ApproveOrder(&_Exchange.TransactOpts, sender, order, orderbookInclusionDesired)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0x5bd0dff3.
//
// Solidity: function atomicMatch((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) buy, (uint8,bytes32,bytes32) buySig, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) sell, (uint8,bytes32,bytes32) sellSig) returns()
func (_Exchange *ExchangeTransactor) AtomicMatch(opts *bind.TransactOpts, buy ExchangeCoreOrder, buySig ExchangeCoreSig, sell ExchangeCoreOrder, sellSig ExchangeCoreSig) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "atomicMatch", buy, buySig, sell, sellSig)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0x5bd0dff3.
//
// Solidity: function atomicMatch((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) buy, (uint8,bytes32,bytes32) buySig, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) sell, (uint8,bytes32,bytes32) sellSig) returns()
func (_Exchange *ExchangeSession) AtomicMatch(buy ExchangeCoreOrder, buySig ExchangeCoreSig, sell ExchangeCoreOrder, sellSig ExchangeCoreSig) (*types.Transaction, error) {
	return _Exchange.Contract.AtomicMatch(&_Exchange.TransactOpts, buy, buySig, sell, sellSig)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0x5bd0dff3.
//
// Solidity: function atomicMatch((address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) buy, (uint8,bytes32,bytes32) buySig, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) sell, (uint8,bytes32,bytes32) sellSig) returns()
func (_Exchange *ExchangeTransactorSession) AtomicMatch(buy ExchangeCoreOrder, buySig ExchangeCoreSig, sell ExchangeCoreOrder, sellSig ExchangeCoreSig) (*types.Transaction, error) {
	return _Exchange.Contract.AtomicMatch(&_Exchange.TransactOpts, buy, buySig, sell, sellSig)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2c1307ce.
//
// Solidity: function cancelOrder(address sender, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) returns()
func (_Exchange *ExchangeTransactor) CancelOrder(opts *bind.TransactOpts, sender common.Address, order ExchangeCoreOrder, sig ExchangeCoreSig) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancelOrder", sender, order, sig)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2c1307ce.
//
// Solidity: function cancelOrder(address sender, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) returns()
func (_Exchange *ExchangeSession) CancelOrder(sender common.Address, order ExchangeCoreOrder, sig ExchangeCoreSig) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, sender, order, sig)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2c1307ce.
//
// Solidity: function cancelOrder(address sender, (address,address,address,uint256,uint256,uint256,uint256,address,uint8,uint8,uint8,address,uint8,bytes,bytes,address,address,uint256,uint256,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) returns()
func (_Exchange *ExchangeTransactorSession) CancelOrder(sender common.Address, order ExchangeCoreOrder, sig ExchangeCoreSig) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, sender, order, sig)
}

// ChangeMinimumMakerProtocolFee is a paid mutator transaction binding the contract method 0x14350c24.
//
// Solidity: function changeMinimumMakerProtocolFee(uint256 newMinimumMakerProtocolFee) returns()
func (_Exchange *ExchangeTransactor) ChangeMinimumMakerProtocolFee(opts *bind.TransactOpts, newMinimumMakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "changeMinimumMakerProtocolFee", newMinimumMakerProtocolFee)
}

// ChangeMinimumMakerProtocolFee is a paid mutator transaction binding the contract method 0x14350c24.
//
// Solidity: function changeMinimumMakerProtocolFee(uint256 newMinimumMakerProtocolFee) returns()
func (_Exchange *ExchangeSession) ChangeMinimumMakerProtocolFee(newMinimumMakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeMinimumMakerProtocolFee(&_Exchange.TransactOpts, newMinimumMakerProtocolFee)
}

// ChangeMinimumMakerProtocolFee is a paid mutator transaction binding the contract method 0x14350c24.
//
// Solidity: function changeMinimumMakerProtocolFee(uint256 newMinimumMakerProtocolFee) returns()
func (_Exchange *ExchangeTransactorSession) ChangeMinimumMakerProtocolFee(newMinimumMakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeMinimumMakerProtocolFee(&_Exchange.TransactOpts, newMinimumMakerProtocolFee)
}

// ChangeMinimumTakerProtocolFee is a paid mutator transaction binding the contract method 0x1a6b13e2.
//
// Solidity: function changeMinimumTakerProtocolFee(uint256 newMinimumTakerProtocolFee) returns()
func (_Exchange *ExchangeTransactor) ChangeMinimumTakerProtocolFee(opts *bind.TransactOpts, newMinimumTakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "changeMinimumTakerProtocolFee", newMinimumTakerProtocolFee)
}

// ChangeMinimumTakerProtocolFee is a paid mutator transaction binding the contract method 0x1a6b13e2.
//
// Solidity: function changeMinimumTakerProtocolFee(uint256 newMinimumTakerProtocolFee) returns()
func (_Exchange *ExchangeSession) ChangeMinimumTakerProtocolFee(newMinimumTakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeMinimumTakerProtocolFee(&_Exchange.TransactOpts, newMinimumTakerProtocolFee)
}

// ChangeMinimumTakerProtocolFee is a paid mutator transaction binding the contract method 0x1a6b13e2.
//
// Solidity: function changeMinimumTakerProtocolFee(uint256 newMinimumTakerProtocolFee) returns()
func (_Exchange *ExchangeTransactorSession) ChangeMinimumTakerProtocolFee(newMinimumTakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeMinimumTakerProtocolFee(&_Exchange.TransactOpts, newMinimumTakerProtocolFee)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_Exchange *ExchangeTransactor) ChangeProtocolFeeRecipient(opts *bind.TransactOpts, newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "changeProtocolFeeRecipient", newProtocolFeeRecipient)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_Exchange *ExchangeSession) ChangeProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeProtocolFeeRecipient(&_Exchange.TransactOpts, newProtocolFeeRecipient)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_Exchange *ExchangeTransactorSession) ChangeProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeProtocolFeeRecipient(&_Exchange.TransactOpts, newProtocolFeeRecipient)
}

// ChangeRoyaltyFee is a paid mutator transaction binding the contract method 0x1172a704.
//
// Solidity: function changeRoyaltyFee(uint256 newRoyaltyFee) returns()
func (_Exchange *ExchangeTransactor) ChangeRoyaltyFee(opts *bind.TransactOpts, newRoyaltyFee *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "changeRoyaltyFee", newRoyaltyFee)
}

// ChangeRoyaltyFee is a paid mutator transaction binding the contract method 0x1172a704.
//
// Solidity: function changeRoyaltyFee(uint256 newRoyaltyFee) returns()
func (_Exchange *ExchangeSession) ChangeRoyaltyFee(newRoyaltyFee *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeRoyaltyFee(&_Exchange.TransactOpts, newRoyaltyFee)
}

// ChangeRoyaltyFee is a paid mutator transaction binding the contract method 0x1172a704.
//
// Solidity: function changeRoyaltyFee(uint256 newRoyaltyFee) returns()
func (_Exchange *ExchangeTransactorSession) ChangeRoyaltyFee(newRoyaltyFee *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeRoyaltyFee(&_Exchange.TransactOpts, newRoyaltyFee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeSession) Receive() (*types.Transaction, error) {
	return _Exchange.Contract.Receive(&_Exchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeTransactorSession) Receive() (*types.Transaction, error) {
	return _Exchange.Contract.Receive(&_Exchange.TransactOpts)
}

// ExchangeOrderApprovedPartOneIterator is returned from FilterOrderApprovedPartOne and is used to iterate over the raw logs and unpacked data for OrderApprovedPartOne events raised by the Exchange contract.
type ExchangeOrderApprovedPartOneIterator struct {
	Event *ExchangeOrderApprovedPartOne // Event containing the contract specifics and raw log

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
func (it *ExchangeOrderApprovedPartOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOrderApprovedPartOne)
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
		it.Event = new(ExchangeOrderApprovedPartOne)
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
func (it *ExchangeOrderApprovedPartOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOrderApprovedPartOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOrderApprovedPartOne represents a OrderApprovedPartOne event raised by the Exchange contract.
type ExchangeOrderApprovedPartOne struct {
	Hash             [32]byte
	Exchange         common.Address
	Maker            common.Address
	Taker            common.Address
	MakerRelayerFee  *big.Int
	TakerRelayerFee  *big.Int
	MakerProtocolFee *big.Int
	TakerProtocolFee *big.Int
	FeeRecipient     common.Address
	FeeMethod        uint8
	Side             uint8
	SaleKind         uint8
	Target           common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderApprovedPartOne is a free log retrieval operation binding the contract event 0x90c7f9f5b58c15f0f635bfb99f55d3d78fdbef3559e7d8abf5c81052a5276622.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, uint256 makerRelayerFee, uint256 takerRelayerFee, uint256 makerProtocolFee, uint256 takerProtocolFee, address indexed feeRecipient, uint8 feeMethod, uint8 side, uint8 saleKind, address target)
func (_Exchange *ExchangeFilterer) FilterOrderApprovedPartOne(opts *bind.FilterOpts, hash [][32]byte, maker []common.Address, feeRecipient []common.Address) (*ExchangeOrderApprovedPartOneIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OrderApprovedPartOne", hashRule, makerRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOrderApprovedPartOneIterator{contract: _Exchange.contract, event: "OrderApprovedPartOne", logs: logs, sub: sub}, nil
}

// WatchOrderApprovedPartOne is a free log subscription operation binding the contract event 0x90c7f9f5b58c15f0f635bfb99f55d3d78fdbef3559e7d8abf5c81052a5276622.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, uint256 makerRelayerFee, uint256 takerRelayerFee, uint256 makerProtocolFee, uint256 takerProtocolFee, address indexed feeRecipient, uint8 feeMethod, uint8 side, uint8 saleKind, address target)
func (_Exchange *ExchangeFilterer) WatchOrderApprovedPartOne(opts *bind.WatchOpts, sink chan<- *ExchangeOrderApprovedPartOne, hash [][32]byte, maker []common.Address, feeRecipient []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OrderApprovedPartOne", hashRule, makerRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOrderApprovedPartOne)
				if err := _Exchange.contract.UnpackLog(event, "OrderApprovedPartOne", log); err != nil {
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

// ParseOrderApprovedPartOne is a log parse operation binding the contract event 0x90c7f9f5b58c15f0f635bfb99f55d3d78fdbef3559e7d8abf5c81052a5276622.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, uint256 makerRelayerFee, uint256 takerRelayerFee, uint256 makerProtocolFee, uint256 takerProtocolFee, address indexed feeRecipient, uint8 feeMethod, uint8 side, uint8 saleKind, address target)
func (_Exchange *ExchangeFilterer) ParseOrderApprovedPartOne(log types.Log) (*ExchangeOrderApprovedPartOne, error) {
	event := new(ExchangeOrderApprovedPartOne)
	if err := _Exchange.contract.UnpackLog(event, "OrderApprovedPartOne", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeOrderApprovedPartTwoIterator is returned from FilterOrderApprovedPartTwo and is used to iterate over the raw logs and unpacked data for OrderApprovedPartTwo events raised by the Exchange contract.
type ExchangeOrderApprovedPartTwoIterator struct {
	Event *ExchangeOrderApprovedPartTwo // Event containing the contract specifics and raw log

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
func (it *ExchangeOrderApprovedPartTwoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOrderApprovedPartTwo)
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
		it.Event = new(ExchangeOrderApprovedPartTwo)
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
func (it *ExchangeOrderApprovedPartTwoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOrderApprovedPartTwoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOrderApprovedPartTwo represents a OrderApprovedPartTwo event raised by the Exchange contract.
type ExchangeOrderApprovedPartTwo struct {
	Hash                      [32]byte
	HowToCall                 uint8
	Data                      []byte
	ReplacementPattern        []byte
	StaticTarget              common.Address
	PaymentToken              common.Address
	BasePrice                 *big.Int
	Extra                     *big.Int
	ListingTime               *big.Int
	ExpirationTime            *big.Int
	Salt                      *big.Int
	OrderbookInclusionDesired bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterOrderApprovedPartTwo is a free log retrieval operation binding the contract event 0xa58be88a9401219b4c2e662028f7ea3dc31bb7ff3a8d5ad378d62f955de1a04d.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, uint8 howToCall, bytes data, bytes replacementPattern, address staticTarget, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_Exchange *ExchangeFilterer) FilterOrderApprovedPartTwo(opts *bind.FilterOpts, hash [][32]byte) (*ExchangeOrderApprovedPartTwoIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OrderApprovedPartTwo", hashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOrderApprovedPartTwoIterator{contract: _Exchange.contract, event: "OrderApprovedPartTwo", logs: logs, sub: sub}, nil
}

// WatchOrderApprovedPartTwo is a free log subscription operation binding the contract event 0xa58be88a9401219b4c2e662028f7ea3dc31bb7ff3a8d5ad378d62f955de1a04d.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, uint8 howToCall, bytes data, bytes replacementPattern, address staticTarget, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_Exchange *ExchangeFilterer) WatchOrderApprovedPartTwo(opts *bind.WatchOpts, sink chan<- *ExchangeOrderApprovedPartTwo, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OrderApprovedPartTwo", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOrderApprovedPartTwo)
				if err := _Exchange.contract.UnpackLog(event, "OrderApprovedPartTwo", log); err != nil {
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

// ParseOrderApprovedPartTwo is a log parse operation binding the contract event 0xa58be88a9401219b4c2e662028f7ea3dc31bb7ff3a8d5ad378d62f955de1a04d.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, uint8 howToCall, bytes data, bytes replacementPattern, address staticTarget, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_Exchange *ExchangeFilterer) ParseOrderApprovedPartTwo(log types.Log) (*ExchangeOrderApprovedPartTwo, error) {
	event := new(ExchangeOrderApprovedPartTwo)
	if err := _Exchange.contract.UnpackLog(event, "OrderApprovedPartTwo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Exchange contract.
type ExchangeOrderCancelledIterator struct {
	Event *ExchangeOrderCancelled // Event containing the contract specifics and raw log

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
func (it *ExchangeOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOrderCancelled)
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
		it.Event = new(ExchangeOrderCancelled)
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
func (it *ExchangeOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOrderCancelled represents a OrderCancelled event raised by the Exchange contract.
type ExchangeOrderCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_Exchange *ExchangeFilterer) FilterOrderCancelled(opts *bind.FilterOpts, hash [][32]byte) (*ExchangeOrderCancelledIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OrderCancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOrderCancelledIterator{contract: _Exchange.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_Exchange *ExchangeFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *ExchangeOrderCancelled, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OrderCancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOrderCancelled)
				if err := _Exchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_Exchange *ExchangeFilterer) ParseOrderCancelled(log types.Log) (*ExchangeOrderCancelled, error) {
	event := new(ExchangeOrderCancelled)
	if err := _Exchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the Exchange contract.
type ExchangeOrdersMatchedIterator struct {
	Event *ExchangeOrdersMatched // Event containing the contract specifics and raw log

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
func (it *ExchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOrdersMatched)
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
		it.Event = new(ExchangeOrdersMatched)
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
func (it *ExchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOrdersMatched represents a OrdersMatched event raised by the Exchange contract.
type ExchangeOrdersMatched struct {
	BuyHash  [32]byte
	SellHash [32]byte
	Maker    common.Address
	Taker    common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x9a4fbbef9b6121513b138cfc80c3a8c19a221e65d5ef1c2b6a3e12cbeca9d610.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, uint256 price)
func (_Exchange *ExchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address) (*ExchangeOrdersMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOrdersMatchedIterator{contract: _Exchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x9a4fbbef9b6121513b138cfc80c3a8c19a221e65d5ef1c2b6a3e12cbeca9d610.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, uint256 price)
func (_Exchange *ExchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *ExchangeOrdersMatched, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOrdersMatched)
				if err := _Exchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x9a4fbbef9b6121513b138cfc80c3a8c19a221e65d5ef1c2b6a3e12cbeca9d610.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, uint256 price)
func (_Exchange *ExchangeFilterer) ParseOrdersMatched(log types.Log) (*ExchangeOrdersMatched, error) {
	event := new(ExchangeOrdersMatched)
	if err := _Exchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Exchange contract.
type ExchangeOwnershipTransferredIterator struct {
	Event *ExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipTransferred)
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
		it.Event = new(ExchangeOwnershipTransferred)
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
func (it *ExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Exchange contract.
type ExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipTransferredIterator{contract: _Exchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipTransferred)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*ExchangeOwnershipTransferred, error) {
	event := new(ExchangeOwnershipTransferred)
	if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
