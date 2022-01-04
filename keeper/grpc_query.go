package keeper

import (
	"context"
	"fmt"
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
	denomId := strings.ToLower(strings.TrimSpace(request.DenomId))

	ctx := sdk.UnwrapSDKContext(c)

	nfts := k.GetMarketPlaceNFTs(ctx, denomId)
	//nfts1 := nfts.(types.NFTs)

	fmt.Println(nfts)
	return nil, nil
}
