package keeper

import (
	"context"
	"reflect"
	"strings"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
	"github.com/AutonomyNetwork/nft/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the NFT MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) CreateDenom(goCtx context.Context,
	msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	id := strings.ToLower(strings.TrimSpace(msg.Id))
	name := strings.ToLower(strings.TrimSpace(msg.Name))
	
	collectionCreator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "address decode failed")
	}
	
	if len(msg.CommunityId) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "invalid community id")
	}
	
	// check community exist
	if !m.HasCommunity(ctx, msg.CommunityId) {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "%s, community not exist", id)
	}
	
	access := m.AuthorizedCommunityMember(ctx, msg.CommunityId, collectionCreator)
	if !(access) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "%s, doesn't have access to create collection", msg.Creator)
	}
	
	// Check is there any dependent collection
	if len(msg.DepedentCollection) != 0 {
		for _, id := range msg.DepedentCollection {
			if !m.HasCommunity(ctx, id) {
				return nil, sdkerrors.Wrapf(types.ErrUnknownCollection, "%s, dependent collection not found", id)
			}
		}
	}
	
	if err := m.Keeper.CreateDenom(ctx,
		id,
		name,
		msg.Symbol,
		msg.Description,
		msg.PreviewURI,
		msg.Creator,
		msg.CommunityId,
		msg.DepedentCollection,
	); err != nil {
		return nil, err
	}
	
	ctx.EventManager().EmitTypedEvent(
		&types.EventCreateDenom{
			Id:      msg.Id,
			Symbol:  msg.Symbol,
			Name:    msg.Name,
			Creator: msg.Creator,
		},
	)
	
	return &types.MsgCreateDenomResponse{}, nil
}

func (m msgServer) MintNFT(goCtx context.Context,
	msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	
	owner := creator
	ctx := sdk.UnwrapSDKContext(goCtx)
	denom, err := m.GetDenom(ctx, msg.DenomId)
	if err!=nil{
		return nil, sdkerrors.Wrapf(types.ErrInvalidCollection, "%s collection", msg.DenomId)
	}
	
	if !strings.EqualFold(denom.Creator, msg.Creator){
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "%s don't have access to mint nft in %s collection", msg.Creator, denom.Id)
	}
	
	if err := m.Keeper.MintNFT(ctx,
		msg.DenomId,
		msg.Id,
		msg.Royalties,
		msg.Transferable,
		owner,
		creator,
		msg.Metadata,
		msg.Data,
	); err != nil {
		return nil, err
	}
	
	ctx.EventManager().EmitTypedEvent(
		&types.EventMintNFT{
			Id:      msg.Id,
			DenomId: msg.DenomId,
			Creator: msg.Creator,
		},
	)
	
	return &types.MsgMintNFTResponse{}, nil
}

func (m msgServer) UpdateNFT(goCtx context.Context,
	msg *types.MsgUpdateNFT) (*types.MsgUpdateNFTResponse, error) {
	
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}
	
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.UpdateNFT(ctx, msg.DenomID, msg.Id,
		msg.Name,
		msg.Description,
		msg.Royalties,
		owner,
	); err != nil {
		return nil, err
	}
	
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateNFT{
			Id:      msg.Id,
			DenomId: msg.DenomID,
			Owner:   msg.Owner,
		},
	)
	return &types.MsgUpdateNFTResponse{}, nil
}

func (m msgServer) TransferNFT(goCtx context.Context,
	msg *types.MsgTransferNFT) (*types.MsgTransferNFTResponse, error) {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}
	
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.TransferOwner(ctx, msg.DenomId, msg.Id,
		sender,
		recipient,
	); err != nil {
		return nil, err
	}
	
	ctx.EventManager().EmitTypedEvent(
		&types.EventTransferNFT{
			Id:        msg.Id,
			DenomID:   msg.DenomId,
			Sender:    msg.Sender,
			Recipient: msg.Recipient,
		},
	)
	
	return &types.MsgTransferNFTResponse{}, nil
}

func (m msgServer) SellNFT(goCtx context.Context,
	msg *types.MsgSellNFT) (*types.MsgSellNFTResponse, error) {
	
	seller, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return nil, err
	}
	
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.SellNFT(ctx, msg.Id, msg.DenomId, msg.Price, seller); err != nil {
		return nil, err
	}
	
	ctx.EventManager().EmitTypedEvent(
		&types.EventSellNFT{
			Id:      msg.Id,
			DenomId: msg.DenomId,
			Price:   msg.Price,
			Seller:  msg.Seller,
		})
	
	return &types.MsgSellNFTResponse{}, nil
}

func (m msgServer) BuyNFT(goCtx context.Context,
	msg *types.MsgBuyNFT) (*types.MsgBuyNFTResponse, error) {
	
	buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}
	
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.BuyNFT(ctx, msg.Id, msg.DenomId, buyer); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidNFT, err.Error())
	}
	
	ctx.EventManager().EmitTypedEvent(
		&types.EventBuyNFT{
			Id:      msg.Id,
			DenomId: msg.DenomId,
			Buyer:   msg.Buyer,
		})
	
	return &types.MsgBuyNFTResponse{}, nil
}

func (m msgServer) CreateCommunity(goCtx context.Context, msg *types.MsgCreateCommunity) (*types.MsgCreateCommunityResponse, error) {
	
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	community := types.Community{
		Name:        msg.Name,
		Id:          msg.Id,
		Creator:     msg.Creator,
		Description: msg.Description,
		PreviewURI:  msg.PreviewUri,
	}
	if err := m.Keeper.SetCommunity(ctx, community); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, err.Error())
	}
	
	ctx.EventManager().EmitTypedEvent(
		&types.EventCreateCommunity{
			Id:      msg.Id,
			Creator: msg.Creator,
			Name:    msg.Name,
		})
	
	return &types.MsgCreateCommunityResponse{
		Id: msg.Id,
	}, nil
}

func (m msgServer) JoinCommunity(goCtx context.Context, msg *types.MsgJoinCommunity) (*types.MsgJoinCommunityResponse, error) {
	
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	if !m.HasCommunity(ctx, msg.CommunityId) {
		return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "communit not exis: %s", msg.CommunityId)
	}
	
	cm, err := m.GetCommunityMembers(ctx, msg.CommunityId)
	if reflect.DeepEqual(cm, types.CommunityMembers{}) && err == nil {
		cm.CommunityId = msg.CommunityId
		cm.Addresses = append(cm.Addresses, msg.Address)
	} else {
		for _, address := range cm.Addresses {
			if strings.EqualFold(address, msg.Address) {
				return nil, sdkerrors.Wrapf(types.ErrCommunityNotFound, "address already exist")
			}
		}
		cm.Addresses = append(cm.Addresses, msg.Address)
	}
	
	m.SetCommunityMembers(ctx, cm)
	ctx.EventManager().EmitTypedEvent(
		&types.EventJoinCommunity{
			Id:      msg.CommunityId,
			Creator: msg.Address,
		})
	
	return &types.MsgJoinCommunityResponse{
	
	}, nil
}
