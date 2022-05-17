package keeper

import (
	"context"
	"strings"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
	"github.com/AutonomyNetwork/nft/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Denom(c context.Context, request *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	denom := strings.ToLower(strings.TrimSpace(request.DenomId))
	ctx := sdk.UnwrapSDKContext(c)
	
	denomObject, err := k.GetDenom(ctx, denom)
	if err != nil {
		return nil, err
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

func (k Keeper) NFT(c context.Context, request *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {
	denom := strings.ToLower(strings.TrimSpace(request.DenomId))
	tokenID := strings.ToLower(strings.TrimSpace(request.Id))
	ctx := sdk.UnwrapSDKContext(c)
	
	nft, err := k.GetNFT(ctx, denom, tokenID)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid NFT %s from collection %s", request.Id, request.DenomId)
	}
	
	NFT, ok := nft.(types.NFT)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid type NFT %s from collection %s", request.Id, request.DenomId)
	}
	
	return &types.QueryNFTResponse{
		NFT: &NFT,
	}, nil
}

func (k Keeper) MarketPlaceNFT(c context.Context, request *types.QueryMarketPlaceNFTRequest) (*types.QueryMarketPlaceNFTResponse, error) {
	denom := strings.ToLower(strings.TrimSpace(request.DenomId))
	nftID := strings.ToLower(strings.TrimSpace(request.Id))
	ctx := sdk.UnwrapSDKContext(c)
	
	marketplace, err := k.GetMarketPlaceNFT(ctx, denom, nftID)
	if err != nil {
		return nil, err
	}
	
	MarketPlace, ok := marketplace.(types.MarketPlace)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid type MarketPlace %s from collection %s", request.Id, request.DenomId)
	}
	
	return &types.QueryMarketPlaceNFTResponse{
		MarketPlace: &MarketPlace,
	}, nil
	
}

func (k Keeper) MarketPlace(c context.Context, request *types.QueryMarketPlaceRequest) (*types.QueryMarketPlaceResponse, error) {
	
	ctx := sdk.UnwrapSDKContext(c)
	
	marketPlace := k.GetMarketPlace(ctx)
	return &types.QueryMarketPlaceResponse{
		MarketPlace: marketPlace,
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
	
	return &types.QueryCommunitiesResponse{Communities: k.GetCommunities(ctx)}, nil
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
