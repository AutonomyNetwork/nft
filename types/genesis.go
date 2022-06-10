package types

// NewGenesisState creates a new genesis state.
func NewGenesisState(collections []Collection, orders []MarketPlace, communitites []Community) *GenesisState {
	return &GenesisState{
		Collections: collections,
		Orders:      orders,
		Communities: communitites,
	}
}
