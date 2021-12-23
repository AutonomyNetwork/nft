package exported

import sdk "github.com/cosmos/cosmos-sdk/types"

type MarketPlace interface {
	GetNFTID() string
	GetDenomID() string
	GetPrice() string
	GetSeller() sdk.AccAddress
	GetBuyer() sdk.AccAddress
}
