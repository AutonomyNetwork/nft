package keeper

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AutonomyNetwork/nft/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      codec.BinaryCodec
}

// NewKeeper creates new instances of the nft Keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("AutonomyNetwork/%s", types.ModuleName))
}

func (k Keeper) CreateDenom(ctx sdk.Context, id, name, symbol, description, previewURI string,
	creator sdk.AccAddress) error {
	return k.SetDenom(ctx, types.NewDenom(id, name, symbol, description, previewURI, creator))
}

// MintNFT mints an NFT and manages that NFTs existence within Collections and Owners
func (k Keeper) MintNFT(ctx sdk.Context,
	denomID, nftID, data, royalties string, transferable bool,
	owner, creator sdk.AccAddress, metadata types.Metadata) error {
	if !k.HasDenomID(ctx, denomID) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not exists", denomID)
	}

	if k.HasNFT(ctx, denomID, nftID) {
		return sdkerrors.Wrapf(types.ErrNFTAlreadyExists, "NFT %s already exists in collection %s", nftID, denomID)
	}

	k.SetNFT(ctx, denomID, types.NewBaseNFT(
		nftID,
		metadata,
		owner,
		data,
		transferable,
		royalties,
		creator,
		ctx.BlockTime(),
	))
	k.setOwner(ctx, denomID, nftID, owner)
	k.increaseSupply(ctx, denomID)
	return nil
}

//// EditNFT updates an already existing NFTs
//func (k Keeper) EditNFT(ctx sdk.Context,
//	denomID, tokenID, tokenNm, tokenURI, tokenData string,
//	owner sdk.AccAddress) error {
//	if !k.HasDenomID(ctx, denomID) {
//		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not exists", denomID)
//	}
//
//	nft, err := k.Authorize(ctx, denomID, tokenID, owner)
//	if err != nil {
//		return err
//	}
//
//	if tokenNm != types.DoNotModify {
//		nft.Name = tokenNm
//	}
//
//	if tokenURI != types.DoNotModify {
//		nft.URI = tokenURI
//	}
//
//	if tokenData != types.DoNotModify {
//		nft.Data = tokenData
//	}
//
//	k.setNFT(ctx, denomID, nft)
//	return nil
//}
//
//// TransferOwner gets all the ID Collections owned by an address
//func (k Keeper) TransferOwner(ctx sdk.Context,
//	denomID, tokenID, tokenNm, tokenURI, tokenData string,
//	srcOwner, dstOwner sdk.AccAddress) error {
//	if !k.HasDenomID(ctx, denomID) {
//		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not exists", denomID)
//	}
//
//	nft, err := k.Authorize(ctx, denomID, tokenID, srcOwner)
//	if err != nil {
//		return err
//	}
//
//	nft.Owner = dstOwner
//
//	if tokenNm != types.DoNotModify {
//		nft.Name = tokenNm
//	}
//
//	if tokenURI != types.DoNotModify {
//		nft.URI = tokenURI
//	}
//
//	if tokenData != types.DoNotModify {
//		nft.Data = tokenData
//	}
//
//	k.setNFT(ctx, denomID, nft)
//	k.swapOwner(ctx, denomID, tokenID, srcOwner, dstOwner)
//	return nil
//}
//
//// BurnNFT delete a specified nft
//func (k Keeper) BurnNFT(ctx sdk.Context,
//	denomID, tokenID string,
//	owner sdk.AccAddress) error {
//	if !k.HasDenomID(ctx, denomID) {
//		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not exists", denomID)
//	}
//
//	nft, err := k.Authorize(ctx, denomID, tokenID, owner)
//	if err != nil {
//		return err
//	}
//
//	k.deleteNFT(ctx, denomID, nft)
//	k.deleteOwner(ctx, denomID, tokenID, owner)
//	k.decreaseSupply(ctx, denomID)
//	return nil
//}
