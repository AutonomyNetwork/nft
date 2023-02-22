package types

import (
	"regexp"
	"strings"
	"unicode/utf8"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// constant used to indicate that some field should not be updated
const (
	DoNotModify  = "[do-not-modify]"
	MinDenomLen  = 3
	MaxDenomLen  = 64
	MinSymbolLen = 3
	MaxSymbolLen = 12
	MaxURILen    = 256
)

const (
	TypeCreateDenom     = "create_denom"
	TypeMintNFT         = "mint_nft"
	TypeUpdateNFT       = "update_nft"
	TypeTransferNFT     = "transfer_nft"
	TypeSellNFT         = "sell_nft"
	TypeBuyNFT          = "buy_nft"
	TypeCreateCommunity = "create_community"
	TypeJoinCommunity   = "join_community"
	TypeUpdtaeCommunity = "update_community"
	TypeUpdateDenom     = "update_denom"
)

var (
	// IsAlphaNumeric only accepts alphanumeric characters
	IsAlphaNumeric   = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	IsBeginWithAlpha = regexp.MustCompile(`^[a-zA-Z].*`).MatchString
	IsAlpha          = regexp.MustCompile(`^[a-zA-Z].*`).MatchString
)

var (
	_ sdk.Msg = &MsgCreateDenom{}
	_ sdk.Msg = &MsgMintNFT{}
	_ sdk.Msg = &MsgUpdateNFT{}
	_ sdk.Msg = &MsgTransferNFT{}
	_ sdk.Msg = &MsgSellNFT{}
	_ sdk.Msg = &MsgBuyNFT{}
	_ sdk.Msg = &MsgCreateCommunity{}
	_ sdk.Msg = &MsgJoinCommunity{}
	_ sdk.Msg = &MsgUpdateCommunity{}
	_ sdk.Msg = &MsgUpdateDenom{}
)

func NewMsgCreateDenom(name, symbol, description, preview_uri, creator, community_id string, dependecy_collection []string) *MsgCreateDenom {
	return &MsgCreateDenom{
		Id:                 GenUniqueID(DenomPrefix),
		Name:               name,
		Symbol:             symbol,
		Description:        description,
		PreviewURI:         preview_uri,
		Creator:            creator,
		CommunityId:        community_id,
		DepedentCollection: dependecy_collection,
	}
}

func (msg MsgCreateDenom) Route() string { return RouterKey }

func (msg MsgCreateDenom) Type() string { return TypeCreateDenom }

func (msg MsgCreateDenom) ValidateBasic() error {
	if err := ValidateDenomID(msg.Id); err != nil {
		return err
	}

	if err := ValidateDenomSymbol(msg.Symbol); err != nil {
		return err
	}

	name := strings.TrimSpace(msg.Name)
	if len(name) > 0 && !utf8.ValidString(name) {
		return sdkerrors.Wrap(ErrInvalidDenom, "denom name is invalid")
	}

	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg MsgCreateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateDenom) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgMintNFT(denomId, data, creator, royalties string, metadata Metadata, transferable bool) *MsgMintNFT {
	return &MsgMintNFT{
		Id:           GenUniqueID(NFTPrefix),
		DenomId:      denomId,
		Data:         data,
		Creator:      creator,
		Royalties:    royalties,
		Metadata:     metadata,
		Transferable: transferable,
	}
}
func (msg MsgMintNFT) Route() string { return RouterKey }

func (msg MsgMintNFT) Type() string { return TypeMintNFT }

func (msg MsgMintNFT) ValidateBasic() error {
	if err := ValidateNFTID(msg.Id); err != nil {
		return err
	}

	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	name := strings.TrimSpace(msg.Metadata.Name)
	if len(name) > 0 && !utf8.ValidString(name) {
		return sdkerrors.Wrap(ErrInvalidDenom, "name is invalid")
	}

	description := strings.TrimSpace(msg.Metadata.Description)
	if len(description) > 0 && !utf8.ValidString(description) {
		return sdkerrors.Wrap(ErrInvalidDenom, "description is invalid")
	}

	if err := ValidateMediaURI(msg.Metadata.MediaURI); err != nil {
		return err
	}

	if err := ValidatePreviewURI(msg.Metadata.PreviewURI); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateNFT(id, denomId, royalties, description, name, owner string) *MsgUpdateNFT {
	return &MsgUpdateNFT{
		Id:          id,
		DenomID:     denomId,
		Royalties:   royalties,
		Description: description,
		Name:        name,
		Owner:       owner,
	}
}

func (msg MsgUpdateNFT) Route() string { return RouterKey }

func (msg MsgUpdateNFT) Type() string { return TypeUpdateNFT }

func (msg MsgUpdateNFT) ValidateBasic() error {
	// if _, err := sdk.AccAddressFromBech32(msg.Owner); err != nil {
	//	return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address %s", err)
	// }
	return nil
}

func (msg MsgUpdateNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgUpdateNFT) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{from}
}

func NewMsgTransferNFT(id, denom, sender, recipient string) *MsgTransferNFT {
	return &MsgTransferNFT{
		Id:        id,
		DenomId:   denom,
		Sender:    sender,
		Recipient: recipient,
	}
}

func (msg MsgTransferNFT) Route() string { return RouterKey }

func (msg MsgTransferNFT) Type() string { return TypeTransferNFT }

func (msg MsgTransferNFT) ValidateBasic() error {
	if err := ValidateNFTID(msg.Id); err != nil {
		return err
	}
	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address %s", err)
	}
	return nil
}

func (msg MsgTransferNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgTransferNFT) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{from}
}

func NewMsgSellNFT(nftID, denomID, price, seller string) *MsgSellNFT {
	return &MsgSellNFT{
		Id:      nftID,
		DenomId: denomID,
		Price:   price,
		Seller:  seller,
	}
}

func (msg MsgSellNFT) Route() string { return RouterKey }

func (msg MsgSellNFT) Type() string { return TypeSellNFT }

func (msg MsgSellNFT) ValidateBasic() error {
	if err := ValidateNFTID(msg.Id); err != nil {
		return err
	}

	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Seller); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid seller address %s", err)
	}

	return nil
}

func (msg MsgSellNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSellNFT) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Seller)
	return []sdk.AccAddress{from}
}

func NewMsgBuyNFT(id, denomId, buyer string) *MsgBuyNFT {
	return &MsgBuyNFT{
		Id:      id,
		DenomId: denomId,
		Buyer:   buyer,
	}
}

func (msg MsgBuyNFT) Route() string { return RouterKey }

func (msg MsgBuyNFT) Type() string { return TypeBuyNFT }

func (msg MsgBuyNFT) ValidateBasic() error {
	if err := ValidateNFTID(msg.Id); err != nil {
		return err
	}

	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Buyer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid buyer address %s", msg.Buyer)
	}

	return nil
}

func (msg MsgBuyNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgBuyNFT) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Buyer)
	return []sdk.AccAddress{from}
}

func NewMsgCreateCommunity(name, desc, creator, uri string) *MsgCreateCommunity {
	return &MsgCreateCommunity{
		Name:        name,
		Description: desc,
		Creator:     creator,
		PreviewUri:  uri,
		Id:          GenUniqueID(CommunityPrefix),
	}
}

func (msg MsgCreateCommunity) Route() string { return RouterKey }

func (msg MsgCreateCommunity) Type() string { return TypeCreateCommunity }

func (m MsgCreateCommunity) ValidateBasic() error {
	return nil // TODO: Implement validate methods
}
func (msg MsgCreateCommunity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateCommunity) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgJoinCommunity(id, creator string) *MsgJoinCommunity {
	return &MsgJoinCommunity{
		CommunityId: id,
		Address:     creator,
	}
}

func (msg MsgJoinCommunity) Route() string { return RouterKey }

func (msg MsgJoinCommunity) Type() string { return TypeJoinCommunity }

func (msg MsgJoinCommunity) ValidateBasic() error {
	return nil // TODO: validate basic
}

func (msg MsgJoinCommunity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgJoinCommunity) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{from}
}

func NewMsgUpdateCommunity(id, description, data, address string, tags []string) *MsgUpdateCommunity {
	return &MsgUpdateCommunity{
		Id:          id,
		Description: description,
		Data:        data,
		Tags:        tags,
		Address:     address,
	}
}

func (msg MsgUpdateCommunity) Route() string { return RouterKey }

func (msg MsgUpdateCommunity) Type() string { return TypeUpdtaeCommunity }

func (msg MsgUpdateCommunity) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "owner address is invalid")
	}

	return nil
}

func (msg MsgUpdateCommunity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgUpdateCommunity) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{from}
}

func NewMsgUpdateDenom(id, description, symbol, address string) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Id:          id,
		Description: description,
		Symbol:      symbol,
		Address:     address,
	}
}

func (msg MsgUpdateDenom) Route() string { return RouterKey }

func (msg MsgUpdateDenom) Type() string { return TypeUpdateDenom }

func (msg MsgUpdateDenom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}
	return nil
}

func (msg MsgUpdateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgUpdateDenom) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{from}
}
