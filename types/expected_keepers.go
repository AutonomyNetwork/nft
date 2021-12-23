package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// AccountKeeper defines the expected account keeper for query account
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}
type (
	// BankKeeper defines the expected interface needed to retrieve account balances.
	BankKeeper interface {
		keeper.SendKeeper
		GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	}
)
