package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidCollection  = sdkerrors.Register(ModuleName, 2, "invalid NFT collection")
	ErrUnknownCollection  = sdkerrors.Register(ModuleName, 3, "unknown NFT collection")
	ErrInvalidNFT         = sdkerrors.Register(ModuleName, 4, "invalid NFT")
	ErrNFTAlreadyExists   = sdkerrors.Register(ModuleName, 5, "NFT already exists")
	ErrUnknownNFT         = sdkerrors.Register(ModuleName, 6, "unknown NFT")
	ErrEmptyTokenData     = sdkerrors.Register(ModuleName, 7, "NFT tokenData can't be empty")
	ErrUnauthorized       = sdkerrors.Register(ModuleName, 8, "unauthorized address")
	ErrInvalidDenom       = sdkerrors.Register(ModuleName, 9, "invalid denom")
	ErrInvalidDenomSymbol = sdkerrors.Register(ModuleName, 10, "invalid denom symbol")
	ErrInvalidTokenID     = sdkerrors.Register(ModuleName, 11, "invalid tokenID")
	ErrInvalidTokenURI    = sdkerrors.Register(ModuleName, 12, "invalid tokenURI")
	ErrTransfer           = sdkerrors.Register(ModuleName, 13, "unauthorized to transfer")
	ErrUnknownMarketPlace = sdkerrors.Register(ModuleName, 14, "unknown MarketPlace collection")
	ErrFilledNFT          = sdkerrors.Register(ModuleName, 15, "nft is filled")
)
