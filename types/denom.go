package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewDenom return a new denom
func NewDenom(id, name, symbol, description, preview_uri string, creator sdk.AccAddress ) Denom {
	return Denom{
		Id:      id,
		Name:    name,
		Symbol: symbol,
		Creator: creator.String(),
		Description: description,
		PreviewURI: preview_uri,
	}
}

func ValidateDenomID(denomID string) error {
	denomID = strings.TrimSpace(denomID)
	if len(denomID) < MinDenomLen || len(denomID) > MaxDenomLen {
		return sdkerrors.Wrapf(ErrInvalidDenom, "invalid denom %s, only accepts value [%d, %d]", denomID, MinDenomLen, MaxDenomLen)
	}
	if !IsBeginWithAlpha(denomID) || !IsAlphaNumeric(denomID) {
		return sdkerrors.Wrapf(ErrInvalidDenom, "invalid denom %s, only accepts alphanumeric characters,and begin with an english letter", denomID)
	}
	return nil
}

func ValidateDenomSymbol(denomSymbol string) error {
	denomSymbol = strings.TrimSpace(denomSymbol)
	if len(denomSymbol) < MinSymbolLen || len(denomSymbol) > MaxSymbolLen {
		return sdkerrors.Wrapf(ErrInvalidDenomSymbol, "invalid denom symbol %s, only accepts value [%d, %d]", MinSymbolLen, MaxSymbolLen)
	}

	if !IsBeginWithAlpha(denomSymbol) || !IsAlpha(denomSymbol) {
		return sdkerrors.Wrapf(ErrInvalidDenomSymbol, "invalid denom symbol %s, only accepts alphabetic characters", denomSymbol)
	}
	return nil
 }