package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Sandeep-Narahari/nft/exported"
	"github.com/Sandeep-Narahari/nft/types"
)

func (k Keeper) SetNFTMarketPlace(ctx sdk.Context, order types.MarketPlace) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&order)
	store.Set(types.KeyMarketPlaceNFT(order.GetDenomID(), order.GetNFTID()), bz)
}

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

func (k Keeper) GetMarketPlace(ctx sdk.Context) (marketPlace []types.MarketPlace) {

	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PrefixMarketPlace)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var marketPlace1 types.MarketPlace
		k.cdc.MustUnmarshal(iterator.Value(), &marketPlace1)
		marketPlace = append(marketPlace, marketPlace1)
	}

	return marketPlace
}
