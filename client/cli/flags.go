package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTokenName = "name"
	FlagTokenURI  = "uri"
	FlagTokenData = "data"
	FlagRecipient = "recipient"
	FlagOwner     = "owner"

	FlagDenomName        = "name"
	FlagDenom            = "denom"
	FlagData             = "data"
	FlagSymbol           = "symbol"
	FlagDenomDescription = "description"
	FlagPreviewURI       = "preview_uri"
	FlagRoyalties        = "royalties"
	FlagMediaURI         = "media_uri"
	FlagTransferable     = "transferable"
)

var (
	FsCreateDenom = flag.NewFlagSet("", flag.ContinueOnError)
	FsMintNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferNFT = flag.NewFlagSet("", flag.ContinueOnError)
	FsQuerySupply = flag.NewFlagSet("", flag.ContinueOnError)
	FsQueryOwner  = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreateDenom.String(FlagData, "", "Denom data")
	FsCreateDenom.String(FlagDenomName, "", "The name of the denom")
	FsCreateDenom.String(FlagSymbol, "", "The symbol of the denom")
	FsCreateDenom.String(FlagDenomDescription, "", "Description of the denom")
	FsCreateDenom.String(FlagPreviewURI, "", "preview_uri of the denom")
	FsCreateDenom.String(FlagRoyalties, "", "royalties")

	FsMintNFT.String(FlagTokenURI, "", "URI for supplemental off-chain tokenData (should return a JSON object)")
	FsMintNFT.String(FlagRecipient, "", "Receiver of the nft, if not filled, the default is the sender of the transaction")
	FsMintNFT.String(FlagTokenData, "", "The origin data of nft")
	FsMintNFT.String(FlagTokenName, "", "Name of the nft")
	FsMintNFT.String(FlagTransferable, "", "transferable")

	FsEditNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain tokenData (should return a JSON object)")
	FsEditNFT.String(FlagTokenData, "[do-not-modify]", "The tokenData of nft")
	FsEditNFT.String(FlagTokenName, "[do-not-modify]", "The name of nft")

	FsTransferNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain tokenData (should return a JSON object)")
	FsTransferNFT.String(FlagTokenData, "[do-not-modify]", "The tokenData of nft")
	FsTransferNFT.String(FlagTokenName, "[do-not-modify]", "The name of nft")

	FsQuerySupply.String(FlagOwner, "", "The owner of a nft")

	FsQueryOwner.String(FlagDenom, "", "The name of a collection")
}
