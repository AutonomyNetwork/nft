package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

// constant used to indicate that some field should not be updated
const (
	DoNotModify  = "[do-not-modify]"
	MinDenomLen  = 3
	MaxDenomLen  = 64
	MinSymbolLen = 3
	MaxSymbolLen = 12

	MaxURILen = 256
)

const (
	TypeCreateDenom = "create_denom"
	TypeMintNFT     = "mint_nft"
	TypeUpdateNFT   = "update_nft"
	TypeTransferNFT = "transfer_nft"
	TypeSellNFT     = "sell_nft"
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
)

func NewMsgCreateDenom(name, symbol, description, preview_uri, creator string) *MsgCreateDenom {
	return &MsgCreateDenom{
		Id:          GenUniqueID(DenomPrefix),
		Name:        name,
		Symbol:      symbol,
		Description: description,
		PreviewURI:  preview_uri,
		Creator:     creator,
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

func NewMsgUpdateNFT(id, denomId, royalties, data, description, name, owner string) *MsgUpdateNFT {
	return &MsgUpdateNFT{
		Id:          id,
		DenomID:     denomId,
		Royalties:   royalties,
		Data:        data,
		Description: description,
		Name:        name,
		Owner:       owner,
	}
}

func (msg MsgUpdateNFT) Route() string { return RouterKey }

func (msg MsgUpdateNFT) Type() string { return TypeUpdateNFT }

func (msg MsgUpdateNFT) ValidateBasic() error {
	//if _, err := sdk.AccAddressFromBech32(msg.Owner); err != nil {
	//	return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address %s", err)
	//}
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
