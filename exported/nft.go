package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NFT non fungible token interface
type NFT interface {
	GetID() string
	GetName() string
	GetOwner() sdk.AccAddress
	IsTransferable() bool
	GetRoyalties() string
	GetCreator() sdk.AccAddress
}
