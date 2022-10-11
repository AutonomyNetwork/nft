package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Sandeep-Narahari/nft/types"
)

func (k Keeper) SetCommunity(ctx sdk.Context, community types.Community) error {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&community)
	store.Set(types.KeyCommunityID(community.Id), bz)
	return nil
}

func (k Keeper) HasCommunity(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyCommunityID(id))
}

func (k Keeper) GetCommunityByID(ctx sdk.Context, id string) (types.Community, bool) {
	store := ctx.KVStore(k.storeKey)

	found := k.HasCommunity(ctx, id)
	if !found {
		return types.Community{}, false
	}

	bz := store.Get(types.KeyCommunityID(id))
	if bz == nil {
		return types.Community{}, false
	}

	var community types.Community
	k.cdc.MustUnmarshal(bz, &community)
	return community, true
}

func (k Keeper) GetCommunities(ctx sdk.Context) (communities []types.Community) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PrefixCommunity)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var community types.Community
		k.cdc.MustUnmarshal(iterator.Value(), &community)
		communities = append(communities, community)
	}
	return communities
}

// Check collection creator is a community creator
func (k Keeper) AuthorizedCommunityMember(ctx sdk.Context, community_id string, creator sdk.AccAddress) bool {
	community, found := k.GetCommunityByID(ctx, community_id)
	if !found {
		return false
	}

	addr, err := sdk.AccAddressFromBech32(community.Creator)
	if err != nil {
		return false
	}

	if !addr.Equals(creator) {
		return false
	}

	return true
}

func (k Keeper) GetCommunityMembers(ctx sdk.Context, community_id string) (types.CommunityMembers, error) {
	store := ctx.KVStore(k.storeKey)

	if !k.HasCommunity(ctx, community_id) {
		return types.CommunityMembers{}, sdkerrors.Wrapf(types.ErrCommunityNotFound, "community doesn't exist :%s", community_id)
	}

	data := store.Get(types.KeyCommunityMembers(community_id))
	if data == nil {
		return types.CommunityMembers{}, nil
	}

	var cm types.CommunityMembers
	k.cdc.MustUnmarshal(data, &cm)
	return cm, nil
}

func (k Keeper) SetCommunityMembers(ctx sdk.Context, community_member types.CommunityMembers) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyCommunityMembers(community_member.CommunityId), k.cdc.MustMarshal(&community_member))

}
