package types

import (
	"bytes"
	"errors"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "nft"
	
	// StoreKey is the default store key for NFT
	StoreKey = ModuleName
	
	// QuerierRoute is the querier route for the NFT store.
	QuerierRoute = ModuleName
	
	// RouterKey is the message route for the NFT module
	RouterKey = ModuleName
)

const (
	DenomPrefix     = "nftdenom"
	NFTPrefix       = "nft"
	CommunityPrefix = "community"
)

var (
	PrefixNFT         = []byte{0x01}
	PrefixOwners      = []byte{0x02} // key for a owner
	PrefixCollection  = []byte{0x03} // key for balance of NFTs held by the denom
	PrefixDenom       = []byte{0x04} // key for denom of the nft
	PrefixDenomName   = []byte{0x05} // key for denom name of the nft
	PrefixMarketPlace = []byte{0x06} // key for market place
	
	PrefixCommunity = []byte{0x07}
	PrefixMembers   = []byte{0x08}
	
	delimiter = []byte("/")
)

// SplitKeyOwner return the address,denom,id from the key of stored owner
func SplitKeyOwner(key []byte) (address sdk.AccAddress, denom, id string, err error) {
	key = key[len(PrefixOwners)+len(delimiter):]
	keys := bytes.Split(key, delimiter)
	if len(keys) != 3 {
		return address, denom, id, errors.New("wrong KeyOwner")
	}
	
	address, _ = sdk.AccAddressFromBech32(string(keys[0]))
	denom = string(keys[1])
	id = string(keys[2])
	return
}

// KeyOwner gets the key of a collection owned by an account address
func KeyOwner(address sdk.AccAddress, denomID, tokenID string) []byte {
	key := append(PrefixOwners, delimiter...)
	if address != nil {
		key = append(key, []byte(address.String())...)
		key = append(key, delimiter...)
	}
	
	if address != nil && len(denomID) > 0 {
		key = append(key, []byte(denomID)...)
		key = append(key, delimiter...)
	}
	
	if address != nil && len(denomID) > 0 && len(tokenID) > 0 {
		
		key = append(key, []byte(tokenID)...)
	}
	return key
}

// KeyNFT gets the key of nft stored by an denom and id
func KeyNFT(denomID, tokenID string) []byte {
	key := append(PrefixNFT, delimiter...)
	if len(denomID) > 0 {
		key = append(key, []byte(denomID)...)
		key = append(key, delimiter...)
	}
	
	if len(denomID) > 0 && len(tokenID) > 0 {
		key = append(key, []byte(tokenID)...)
	}
	return key
}

// KeyCollection gets the storeKey by the collection
func KeyCollection(denomID string) []byte {
	key := append(PrefixCollection, delimiter...)
	return append(key, []byte(denomID)...)
}

// KeyDenomID gets the storeKey by the denom id
func KeyDenomID(id string) []byte {
	key := append(PrefixDenom, delimiter...)
	return append(key, []byte(id)...)
}

// KeyDenomID gets the storeKey by the denom name
func KeyDenomName(name string) []byte {
	key := append(PrefixDenomName, delimiter...)
	return append(key, []byte(name)...)
}

func KeyMarketPlace(id string) []byte {
	key := append(PrefixMarketPlace, delimiter...)
	return append(key, []byte(id)...)
}

func KeyMarketPlaceNFT(denomID, tokenID string) []byte {
	key := append(PrefixMarketPlace, delimiter...)
	if len(denomID) > 0 {
		key = append(key, []byte(denomID)...)
		key = append(key, delimiter...)
	}
	
	if len(denomID) > 0 && len(tokenID) > 0 {
		key = append(key, []byte(tokenID)...)
	}
	return key
}

func KeyCommunityID(id string) []byte {
	key := append(PrefixCommunity, delimiter...)
	return append(key, []byte(id)...)
}

func KeyCommunityMembers(id string) []byte {
	key := append(PrefixMembers, delimiter...)
	return append(key, []byte(id)...)
}
