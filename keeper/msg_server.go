package keeper

import (
	"context"
	"github.com/AutonomyNetwork/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
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

	id := strings.ToLower(strings.TrimSpace(msg.Id))
	name := strings.ToLower(strings.TrimSpace(msg.Name))

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.CreateDenom(ctx,
		id,
		name,
		msg.Symbol,
		msg.Description,
		msg.PreviewURI,
		creator,
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

//func (m msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//	sender, err := sdk.AccAddressFromBech32(msg.Sender)
//	if err != nil {
//		return nil, err
//	}
//
//	err = m.Keeper.UpdateDenom(ctx, msg.Id, msg.Name, msg.Description, msg.PreviewURI, sender)
//	if err != nil {
//		return nil, err
//	}
//
//	ctx.EventManager().EmitTypedEvent(
//		&types.EventUpdateDenom{
//			Id:      msg.Id,
//			Name:    msg.Name,
//			Creator: msg.Sender,
//		},
//	)
//
//	return &types.MsgUpdateDenomResponse{}, nil
//}

//func (m msgServer) TransferDenom(goCtx context.Context, msg *types.MsgTransferDenom) (*types.MsgTransferDenomResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//	sender, err := sdk.AccAddressFromBech32(msg.Sender)
//	if err != nil {
//		return nil, err
//	}
//	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
//	if err != nil {
//		return nil, err
//	}
//
//	err = m.Keeper.TransferDenomOwner(ctx, msg.Id, sender, recipient)
//	if err != nil {
//		return nil, err
//	}
//	ctx.EventManager().EmitTypedEvent(
//		&types.EventTransferDenom{
//			Id:        msg.Id,
//			Sender:    msg.Sender,
//			Recipient: msg.Recipient,
//		},
//	)
//
//	return &types.MsgTransferDenomResponse{}, nil
//}

func (m msgServer) MintNFT(goCtx context.Context,
	msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	owner := creator
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.MintNFT(ctx,
		msg.DenomId,
		msg.Id,
		msg.Royalties,
		msg.Transferable,
		owner,
		creator,
		msg.Metadata,
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
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.BuyNFT(ctx, msg.Id, msg.DenomId, buyer); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventBuyNFT{
			Id:      msg.Id,
			DenomId: msg.DenomId,
			Buyer:   msg.Buyer,
		})

	return &types.MsgBuyNFTResponse{}, nil
}

//func (m msgServer) BurnONFT(goCtx context.Context,
//	msg *types.MsgBurnONFT) (*types.MsgBurnONFTResponse, error) {
//
//	sender, err := sdk.AccAddressFromBech32(msg.Sender)
//	if err != nil {
//		return nil, err
//	}
//
//	ctx := sdk.UnwrapSDKContext(goCtx)
//	if err := m.Keeper.BurnONFT(ctx, msg.DenomId, msg.Id, sender); err != nil {
//		return nil, err
//	}
//
//	ctx.EventManager().EmitTypedEvent(
//		&types.EventBurnONFT{
//			Id:      msg.Id,
//			DenomId: msg.DenomId,
//			Owner:   msg.Sender,
//		},
//	)
//
//	return &types.MsgBurnONFTResponse{}, nil
//}
