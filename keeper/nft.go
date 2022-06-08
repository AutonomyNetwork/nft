package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
	"github.com/AutonomyNetwork/nft/exported"
	"github.com/AutonomyNetwork/nft/types"
)

// GetNFT gets the entire NFT tokenData struct
func (k Keeper) GetNFT(ctx sdk.Context, denomID, tokenID string) (nft exported.NFT, err error) {
	store := ctx.KVStore(k.storeKey)
	
	bz := store.Get(types.KeyNFT(denomID, tokenID))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownCollection, "not found NFT: %s", denomID)
	}
	
	var baseNFT types.NFT
	k.cdc.MustUnmarshal(bz, &baseNFT)
	return baseNFT, nil
}

// GetNFTs return the all NFT by the specified denomID
func (k Keeper) GetNFTs(ctx sdk.Context, denom string) (nfts []exported.NFT) {
	store := ctx.KVStore(k.storeKey)
	
	iterator := sdk.KVStorePrefixIterator(store, types.KeyNFT(denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var baseNFT types.NFT
		k.cdc.MustUnmarshal(iterator.Value(), &baseNFT)
		nfts = append(nfts, baseNFT)
	}
	return nfts
}

// Get All NFTs
func(k Keeper) GetAllNFTs(ctx sdk.Context)(nfts []types.NFT){
	store := ctx.KVStore(k.storeKey)
	
	iterator := sdk.KVStorePrefixIterator(store, types.PrefixNFT)
	defer iterator.Close()
	
	
	for ; iterator.Valid(); iterator.Next() {
		var baseNFT types.NFT
		k.cdc.MustUnmarshal(iterator.Value(), &baseNFT)
		nfts = append(nfts, baseNFT)
	}
	return nfts
}

// Authorize check if the sender is the issuer of nft, if it returns nft, if not, return an error
func (k Keeper) Authorize(ctx sdk.Context,
	denomID, tokenID string,
	owner sdk.AccAddress) (types.NFT, error) {
	nft, err := k.GetNFT(ctx, denomID, tokenID)
	if err != nil {
		return types.NFT{}, err
	}
	
	if !owner.Equals(nft.GetOwner()) {
		return types.NFT{}, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}
	return nft.(types.NFT), nil
}

// HasNFT determine if nft exists
func (k Keeper) HasNFT(ctx sdk.Context, denomID, tokenID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyNFT(denomID, tokenID))
}

func (k Keeper) SetNFT(ctx sdk.Context, denomID string, nft types.NFT) {
	store := ctx.KVStore(k.storeKey)
	
	bz := k.cdc.MustMarshal(&nft)
	store.Set(types.KeyNFT(denomID, nft.GetID()), bz)
}

// deleteNFT deletes an existing NFT from store
func (k Keeper) deleteNFT(ctx sdk.Context, denomID string, nft exported.NFT) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNFT(denomID, nft.GetID()))
}

