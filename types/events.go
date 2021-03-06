package types

// NFT module event types
var (
	EventTypeIssueDenom = "issue_denom"
	EventTypeTransfer   = "transfer_nft"
	EventTypeEditNFT    = "edit_nft"
	EventTypeMintNFT    = "mint_nft"
	EventTypeBurnNFT    = "burn_nft"
	
	AttributeValueCategory = ModuleName
	
	AttributeKeySender    = "sender"
	AttributeKeyRecipient = "recipient"
	AttributeKeyOwner     = "owner"
	AttributeKeyTokenID   = "token-id"
	AttributeKeyTokenURI  = "token-uri"
	AttributeKeyDenom     = "denom"
)
