package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidCollection  = sdkerrors.Register(ModuleName, 112, "invalid NFT collection")
	ErrUnknownCollection  = sdkerrors.Register(ModuleName, 113, "unknown NFT collection")
	ErrInvalidNFT         = sdkerrors.Register(ModuleName, 114, "invalid NFT")
	ErrNFTAlreadyExists   = sdkerrors.Register(ModuleName, 115, "NFT already exists")
	ErrUnknownNFT         = sdkerrors.Register(ModuleName, 116, "unknown NFT")
	ErrEmptyTokenData     = sdkerrors.Register(ModuleName, 117, "NFT tokenData can't be empty")
	ErrUnauthorized       = sdkerrors.Register(ModuleName, 118, "unauthorized address")
	ErrInvalidDenom       = sdkerrors.Register(ModuleName, 119, "invalid denom")
	ErrInvalidDenomSymbol = sdkerrors.Register(ModuleName, 120, "invalid denom symbol")
	ErrInvalidTokenID     = sdkerrors.Register(ModuleName, 121, "invalid tokenID")
	ErrInvalidTokenURI    = sdkerrors.Register(ModuleName, 122, "invalid tokenURI")
	ErrTransfer           = sdkerrors.Register(ModuleName, 123, "unauthorized to transfer")
	ErrUnknownMarketPlace = sdkerrors.Register(ModuleName, 124, "unknown MarketPlace collection")
	ErrFilledNFT          = sdkerrors.Register(ModuleName, 125, "nft is filled")
	ErrCommunityNotFound  = sdkerrors.Register(ModuleName, 126, "community  ")
)
