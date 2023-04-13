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

	paymentInfo := types.PaymentInfo{}
	if msg.PrimarySale == true {
		if msg.TotalNfts == 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidTotalNFTs, "total nfts should not be %s", msg.TotalNfts)
		} else {
			msg.AvailableNfts = msg.TotalNfts
		}
		paymentInfo.AccessType = msg.AccessType
		paymentInfo.Amount = msg.Amount
		paymentInfo.Currency = msg.Currency
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
		msg.Category,
		msg.PrimarySale,
		msg.TotalNfts,
		msg.AvailableNfts,
		msg.Data,
		paymentInfo,
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

	var owner sdk.AccAddress
	ctx := sdk.UnwrapSDKContext(goCtx)
	denom, err := m.GetDenom(ctx, msg.DenomId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCollection, "%s collection", msg.DenomId)
	}

	if denom.PrimarySale == false && strings.EqualFold(denom.Creator, msg.Creator) {
		owner = sdk.AccAddress(denom.Creator)
		// return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "%s don't have access to mint nft in %s collection", msg.Creator, denom.Id)
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
	}

	if denom.PrimarySale == true {
		if denom.AvailableNfts == 0 {
			return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "not enough nfts in %s collection to mint", denom.Id)
		}
		owner = sdk.AccAddress(msg.Creator)
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
		denom.AvailableNfts = denom.AvailableNfts - 1
	}

	m.Keeper.SetDenom(ctx, denom)

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
	denom, err := m.Keeper.GetDenom(ctx, msg.DenomID)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCollection, "%s collection", msg.DenomID)
	}

	if denom.PrimarySale == true {
		return nil, sdkerrors.Wrapf(types.ErrInvalidNFT, "cannot update nft which is in primary sale %s", msg.Id)
	}
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

	ctx := sdk.UnwrapSDKContext(goCtx)
	seller, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return nil, err
	}

	if msg.ListedType == types.Crypto {
		if err := m.Keeper.SellNFT(ctx, msg.Id, msg.DenomId, msg.Price, seller); err != nil {
			return nil, err
		}
	} else if msg.ListedType == types.Fiat {
		if err := m.Keeper.SellNFTWithFiat(ctx, msg.Id, msg.DenomId, msg.Currency, msg.FiatAmount, seller); err != nil {
			return nil, err
		}
	} else {
		return nil, sdkerrors.Wrapf(types.ErrFilledNFT, "listed type of  %s does not exist", msg.ListedType)
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

	if msg.ListedType == types.Fiat {
		if err := m.Keeper.BuyNFTWithFiat(ctx, msg.Id, msg.DenomId, msg.Currency, msg.FiatAmount, msg.OrderRefId, buyer); err != nil {
			return nil, err
		}
	} else { // TODO: update this for fiat

		if err := m.Keeper.BuyNFT(ctx, msg.Id, msg.DenomId, buyer); err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidNFT, err.Error())
		}
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
		Tags:        msg.Tags,
		Data:        msg.Data,
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

	return &types.MsgJoinCommunityResponse{}, nil
}

func (m msgServer) UpdateCommunity(goCtx context.Context, msg *types.MsgUpdateCommunity) (*types.MsgUpdateCommunityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address: %s", msg.Address)
	}

	if err := m.Keeper.UpdateCommunity(ctx, msg.Description,
		msg.Data,
		msg.Id,
		msg.Tags,
		owner,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateNFT{
			Id:    msg.Id,
			Owner: owner.String(),
		},
	)

	return &types.MsgUpdateCommunityResponse{
		Id: msg.Id,
	}, nil
}

func (m msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address: %s", msg.Address)
	}

	if err := m.Keeper.UpdateDenom(ctx, msg.Description, msg.Symbol, msg.Id, owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateDenom{
			Id:    msg.Id,
			Owner: owner.String(),
		},
	)
	return &types.MsgUpdateDenomResponse{}, nil
}

func (m msgServer) DeleteMarketPlaceNFT(goCtx context.Context, msg *types.MsgDeleteMarketPlaceNFT) (*types.MsgDeleteMarketPlaceNFTResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	nft, err := m.Keeper.GetNFT(ctx, msg.DenomId, msg.NftId)
	if err != nil {
		return nil, err
	}
	order, err := m.Keeper.GetMarketPlaceNFT(ctx, msg.DenomId, msg.NftId)
	if err != nil {
		return nil, err
	}

	if nft.GetOwner().String() != msg.Address || msg.Address != order.GetSeller().String() {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is unathorized to perform this operation", msg.Address)
	}

	m.Keeper.DeleteMarketPlaceNFT(ctx, msg.DenomId, msg.NftId)

	return &types.MsgDeleteMarketPlaceNFTResponse{}, nil
}
