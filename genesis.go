package nft

import (
	"unicode/utf8"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AutonomyNetwork/nft/keeper"
	"github.com/AutonomyNetwork/nft/types"
)

// InitGenesis sets nft information for genesis.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	if err := ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	for _, c := range data.Collections {
		if err := k.SetDenom(ctx, c.Denom); err != nil {
			panic(err)
		}
		if err := k.SetCollection(ctx, c); err != nil {
			panic(err)
		}
	}
	
	for _, o := range data.Orders {
		k.SetNFTMarketPlace(ctx, o)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(k.GetCollections(ctx), k.GetMarketPlace(ctx))
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.Collection{}, []types.MarketPlace{})
}

// ValidateGenesis performs basic validation of nfts genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data types.GenesisState) error {
	for _, c := range data.Collections {
		if err := types.ValidateDenomID(c.Denom.Id); err != nil {
			return err
		}
		if !utf8.ValidString(c.Denom.Name) {
			return sdkerrors.Wrap(types.ErrInvalidDenom, "denom name is invalid")
		}

		for _, nft := range c.NFTs {
			if nft.GetOwner().Empty() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner")
			}

			if err := types.ValidateNFTID(nft.GetID()); err != nil {
				return err
			}

			if err := types.ValidateMediaURI(nft.GetMediaURI()); err != nil {
				return err
			}
		}
	}
	return nil
}
