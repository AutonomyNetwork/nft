package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/AutonomyNetwork/nft/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Denom(c context.Context, request *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	denom := strings.ToLower(strings.TrimSpace(request.DenomId))
	ctx := sdk.UnwrapSDKContext(c)

	denomObject, err := k.GetDenom(ctx, denom)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCollection, "invalid collection %s", request.DenomId)
	}

	return &types.QueryDenomResponse{
		Denom: &denomObject,
	}, nil
}

func (k Keeper) Denoms(c context.Context, request *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	denoms := k.GetDenoms(ctx)
	return &types.QueryDenomsResponse{
		Denoms: denoms,
	}, nil
}

func (k Keeper) DenomIDsByOwner(c context.Context, request *types.QueryDenomIDsByOwnerRequest) (*types.QueryDenomIDsByOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if len(request.Address) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCollection, "invalid user request address %s", request.Address)
	}
	denoms, err := k.GetDenomsByOwner(ctx, request.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCollection, "invalid collection %s", request.Address)
	}
	var ids []string

	for _, denom := range denoms {
		ids = append(ids, denom.Id)
	}

	return &types.QueryDenomIDsByOwnerResponse{
		Ids: ids,
	}, nil
}

func (k Keeper) NFT(c context.Context, request *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {
	denom := strings.ToLower(strings.TrimSpace(request.DenomId))
	tokenID := strings.ToLower(strings.TrimSpace(request.Id))
	ctx := sdk.UnwrapSDKContext(c)

	denomObject, err := k.GetDenom(ctx, denom)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCollection, "invalid collection %s", request.DenomId)
	}

	nft, err := k.GetNFT(ctx, denom, tokenID)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid NFT %s from collection %s", request.Id, request.DenomId)
	}

	NFT, ok := nft.(types.NFT)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid type NFT %s from collection %s", request.Id, request.DenomId)
	}

	return &types.QueryNFTResponse{
		NFT:   &NFT,
		Denom: &denomObject,
	}, nil
}

func (k Keeper) MarketPlaceNFT(c context.Context, request *types.QueryMarketPlaceNFTRequest) (*types.QueryMarketPlaceNFTResponse, error) {
	denomID := strings.ToLower(strings.TrimSpace(request.DenomId))
	nftID := strings.ToLower(strings.TrimSpace(request.Id))
	ctx := sdk.UnwrapSDKContext(c)

	marketplace, err := k.GetMarketPlaceNFT(ctx, denomID, nftID)
	if err != nil {
		return nil, err
	}

	MarketPlace, ok := marketplace.(types.MarketPlace)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid type MarketPlace %s from collection %s", request.Id, request.DenomId)
	}

	nft, err := k.GetNFT(ctx, denomID, nftID)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid NFT %s from collection %s", request.Id, request.DenomId)
	}

	NFT, ok := nft.(types.NFT)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid type NFT %s from collection %s", request.Id, request.DenomId)
	}

	return &types.QueryMarketPlaceNFTResponse{
		MarketPlace: &MarketPlace,
		NFT:         &NFT,
	}, nil

}

func (k Keeper) MarketPlace(c context.Context, request *types.QueryMarketPlaceRequest) (*types.QueryMarketPlaceResponse, error) {
	var nfts []types.MarketPlace
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)

	marketPlaceStore := prefix.NewStore(store, types.PrefixMarketPlace)

	pageRes, err := query.FilteredPaginate(marketPlaceStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var nft types.MarketPlace
		k.cdc.MustUnmarshal(value, &nft)

		if accumulate {
			nfts = append(nfts, nft)
		}

		return true, nil
	})

	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid market place query %s", err.Error())
	}

	return &types.QueryMarketPlaceResponse{
		MarketPlace: nfts,
		Pagination:  pageRes,
	}, nil
}

func (k Keeper) OwnerNFTs(c context.Context, request *types.QueryOwnerNFTsRequest) (*types.QueryOwnerNFTsResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	owner, err := sdk.AccAddressFromBech32(request.Owner)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid owner address %s", owner.String())
	}

	ownerNFTCollections := k.GetOwnerNFTs(ctx, owner)
	return &types.QueryOwnerNFTsResponse{
		Owner:       request.Owner,
		Collections: ownerNFTCollections,
	}, nil
}

func (k Keeper) Community(c context.Context, request *types.QueryCommunityRequest) (*types.QueryCommunityResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)

	if len(request.CommunityId) < 5 {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "invalid community id: %s", request.CommunityId)
	}

	community, found := k.GetCommunityByID(ctx, request.CommunityId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "invalid community id: %s", request.CommunityId)
	}

	return &types.QueryCommunityResponse{
		Community: &community,
	}, nil
}

func (k Keeper) Communities(c context.Context, request *types.QueryCommunitiesRequest) (*types.QueryCommunitiesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var communities []types.Community

	store := ctx.KVStore(k.storeKey)

	communityStore := prefix.NewStore(store, types.PrefixCommunity)

	pageRes, err := query.FilteredPaginate(communityStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var comminity types.Community
		k.cdc.MustUnmarshal(value, &comminity)

		if accumulate {
			communities = append(communities, comminity)
		}

		return true, nil
	})

	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid communitu query %s", err.Error())
	}

	return &types.QueryCommunitiesResponse{
		Communities: communities,
		Pagination:  pageRes,
	}, nil

}

func (k Keeper) CommunityMembers(c context.Context, request *types.QueryCommunityMembersRequest) (*types.QueryCommunityMembersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if len(request.CommunityId) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "invalid community id: %s", request.CommunityId)
	}

	cm, err := k.GetCommunityMembers(ctx, request.CommunityId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "invalid community id: %s", request.CommunityId)
	}

	return &types.QueryCommunityMembersResponse{Members: &cm}, nil
}

func (k Keeper) Collection(c context.Context, request *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)

	collections, err := k.GetCollection(ctx, request.DenomId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCollection, err.Error())
	}
	return &types.QueryCollectionResponse{
		Collection: &collections,
	}, nil
}

func (k Keeper) CommunityCollections(c context.Context, request *types.QueryCommunityCollectionsRequest) (*types.QueryCommunityCollectionsResponse, error) {
	var denoms []*types.Denom
	ctx := sdk.UnwrapSDKContext(c)

	if len(request.CommunityId) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "invalid community id: %s", request.CommunityId)
	}
	collections := k.GetCollections(ctx)

	community, found := k.GetCommunityByID(ctx, request.CommunityId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "invalid community id: %s", request.CommunityId)
	}

	for _, collection := range collections {
		if strings.EqualFold(request.CommunityId, collection.Denom.CommunityId) {
			denoms = append(denoms, &collection.Denom)
		}
	}

	return &types.QueryCommunityCollectionsResponse{
		Community: &community,
		Denoms:    denoms,
	}, nil
}

func (k Keeper) AllNFTs(c context.Context, request *types.QueryAllNFTsRequest) (*types.QueryAllNFTsResponse, error) {

	var listNFTs []types.ALLNFT
	var allNFT types.ALLNFT

	var denomInfo types.DenomInfo
	var communityInfo types.CommunityInfo

	ctx := sdk.UnwrapSDKContext(c)
	denoms := k.GetDenoms(ctx)

	for _, denom := range denoms { //TODO: work on time complexity
		collection, _ := k.GetCollection(ctx, denom.Id)

		community, _ := k.GetCommunityByID(ctx, denom.CommunityId)
		for _, nft := range collection.NFTs {

			denomInfo.Name = collection.Denom.Name
			denomInfo.DenomId = collection.Denom.Id

			communityInfo.CommunityId = community.Id
			communityInfo.Name = community.Name

			allNFT.DenomInfo = denomInfo
			allNFT.CommunityInfo = communityInfo
			allNFT.Nft = nft

			listNFTs = append(listNFTs, allNFT)
		}

	}

	return &types.QueryAllNFTsResponse{
		All: listNFTs,
	}, nil
}
