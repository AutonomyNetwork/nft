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

func (m msgServer) UpdateNFT(ctx context.Context, nft *types.MsgUpdateNFT) (*types.MsgUpdateNFTResponse, error) {
	panic("implement me")
}

func (m msgServer) TransferNFT(ctx context.Context, nft *types.MsgTransferNFT) (*types.MsgTransferNFTResponse, error) {
	panic("implement me")
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
		msg.Data,
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
			Creator:   msg.Creator,
		},
	)

	return &types.MsgMintNFTResponse{}, nil
}

//func (m msgServer) EditONFT(goCtx context.Context,
//	msg *types.MsgEditONFT) (*types.MsgEditONFTResponse, error) {
//
//	sender, err := sdk.AccAddressFromBech32(msg.Sender)
//	if err != nil {
//		return nil, err
//	}
//
//	ctx := sdk.UnwrapSDKContext(goCtx)
//	if err := m.Keeper.EditONFT(ctx, msg.DenomId, msg.Id,
//		msg.Metadata,
//		msg.Data,
//		msg.Transferable,
//		msg.Extensible,
//		sender,
//	); err != nil {
//		return nil, err
//	}
//
//	ctx.EventManager().EmitTypedEvent(
//		&types.EventEditONFT{
//			Id:      msg.Id,
//			DenomId: msg.DenomId,
//			Owner:   msg.Sender,
//		},
//	)
//	return &types.MsgEditONFTResponse{}, nil
//}
//
//func (m msgServer) TransferONFT(goCtx context.Context,
//	msg *types.MsgTransferONFT) (*types.MsgTransferONFTResponse, error) {
//
//	sender, err := sdk.AccAddressFromBech32(msg.Sender)
//	if err != nil {
//		return nil, err
//	}
//
//	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
//	if err != nil {
//		return nil, err
//	}
//
//	ctx := sdk.UnwrapSDKContext(goCtx)
//	if err := m.Keeper.TransferOwnership(ctx, msg.DenomId, msg.Id,
//		sender,
//		recipient,
//	); err != nil {
//		return nil, err
//	}
//
//	ctx.EventManager().EmitTypedEvent(
//		&types.EventTransferONFT{
//			Id:        msg.Id,
//			DenomId:   msg.DenomId,
//			Sender:    msg.Sender,
//			Recipient: msg.Recipient,
//		},
//	)
//
//	return &types.MsgTransferONFTResponse{}, nil
//}
//
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