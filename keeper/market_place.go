package keeper

import (
	"github.com/AutonomyNetwork/nft/exported"
	"github.com/AutonomyNetwork/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetMarketPlaceNFT(ctx sdk.Context, denomID, id string) (marketplace exported.MarketPlace, err error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyMarketPlaceNFT(denomID, id))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownMarketPlace, "nft %s not found in market place", id)
	}

	var order types.MarketPlace
	k.cdc.MustUnmarshal(bz, &order)
	return order, nil
}

func (k Keeper) GetMarketPlaceNFTs(ctx sdk.Context, denomID string) (marketNFTs []types.NFT) {

	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyMarketPlaceNFT(denomID, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var nft types.NFT
		k.cdc.MustUnmarshal(iterator.Value(), &nft)
		marketNFTs = append(marketNFTs, nft)
	}

	return marketNFTs
}
