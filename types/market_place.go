package types

import (
	"github.com/AutonomyNetwork/nft/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ exported.MarketPlace = MarketPlace{}

func NewMarketPlace(id, denomID, price string, seller sdk.AccAddress) MarketPlace {
	return MarketPlace{
		NftId:   id,
		DenomID: denomID,
		Price:   price,
		Seller:  seller.String(),
	}
}

func (m MarketPlace) GetNFTID() string {
	return m.NftId
}

func (m MarketPlace) GetDenomID() string {
	return m.DenomID
}

func (m MarketPlace) GetPrice() string {
	return m.Price
}

func (m MarketPlace) GetSeller() sdk.AccAddress {
	seller, _ := sdk.AccAddressFromBech32(m.Seller)
	return seller
}

func (m MarketPlace) GetBuyer() sdk.AccAddress {
	buyer, _ := sdk.AccAddressFromBech32(m.Buyer)
	return buyer
}
