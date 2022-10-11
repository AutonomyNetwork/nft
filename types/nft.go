package types

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Sandeep-Narahari/nft/exported"
)

var _ exported.NFT = NFT{}

// NewBaseNFT creates a new NFT instance
func NewBaseNFT(id string, metadata Metadata, owner sdk.AccAddress,
	transferable bool, royalties string, creator sdk.AccAddress, createdTime time.Time, attributes string) NFT {
	return NFT{
		Id:           strings.ToLower(strings.TrimSpace(id)),
		Metadata:     metadata,
		Owner:        owner.String(),
		Transferable: transferable,
		Royalties:    royalties,
		Creator:      creator.String(),
		CreatedAt:    createdTime,
		Data:         attributes,
	}
}

func (nft NFT) GetID() string {
	return nft.Id
}

func (nft NFT) GetName() string {
	return nft.Metadata.Name
}

func (nft NFT) GetDescription() string {
	return nft.Metadata.Description
}

func (nft NFT) GetMetadata() Metadata {
	return nft.Metadata
}

func (nft NFT) GetMediaURI() string {
	return nft.Metadata.MediaURI
}

func (nft NFT) GetPreviewURI() string {
	return nft.Metadata.PreviewURI
}

func (nft NFT) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(nft.Owner)
	return owner
}

func (nft NFT) IsTransferable() bool {
	return nft.Transferable
}

func (nft NFT) GetRoyalties() string {
	return nft.Royalties
}

func (nft NFT) GetCreator() sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(nft.Creator)
	return creator
}

func (nft NFT) GetCreatedTime() time.Time {
	return nft.CreatedAt
}

func (nft NFT) GetAttributes() string {
	return nft.Data
}

// ----------------------------------------------------------------------------
// NFT

// NFTs define a list of NFT
type NFTs []exported.NFT

// NewNFTs creates a new set of NFTs
func NewNFTs(nfts ...exported.NFT) NFTs {
	if len(nfts) == 0 {
		return NFTs{}
	}
	return NFTs(nfts)
}

func ValidateNFTID(id string) error {
	nftID := strings.TrimSpace(id)
	if len(nftID) < MinDenomLen || len(nftID) > MaxDenomLen {
		return sdkerrors.Wrapf(ErrInvalidTokenID, "invalid tokenID %s, only accepts value [%d, %d]", nftID, MinDenomLen, MaxDenomLen)
	}
	if !IsBeginWithAlpha(nftID) || !IsAlphaNumeric(nftID) {
		return sdkerrors.Wrapf(ErrInvalidTokenID, "invalid tokenID %s, only accepts alphanumeric characters,and begin with an english letter", nftID)
	}
	return nil
}

func ValidateMediaURI(tokenURI string) error {
	if len(tokenURI) > MaxURILen {
		return sdkerrors.Wrapf(ErrInvalidTokenURI, "invalid media URI %s, only accepts value [0, %d]", tokenURI, MaxURILen)
	}
	return nil
}

func ValidatePreviewURI(previewURI string) error {
	if len(previewURI) > MaxURILen {
		return sdkerrors.Wrapf(ErrInvalidTokenURI, "invalid previewURI %s, only accepts value [0, %d]", previewURI, MaxURILen)
	}
	return nil
}
