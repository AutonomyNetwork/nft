package types

import (
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewDenom return a new denom
func NewDenom(id, name, symbol, description, preview_uri, creator, community_id string, denoms []string, category string, onDemandMinting bool, totoalNfts, availableNfts int64, data string) Denom {
	return Denom{
		Id:              id,
		Name:            name,
		Symbol:          symbol,
		Creator:         creator,
		Description:     description,
		PreviewURI:      preview_uri,
		DependentDenoms: denoms,
		CommunityId:     community_id,
		Category:        category,
		OnDemandMinting: onDemandMinting,
		TotalNfts:       totoalNfts,
		AvailableNfts:   availableNfts,
		Data:            data,
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
