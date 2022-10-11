package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/Sandeep-Narahari/nft/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey      sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc           codec.BinaryCodec
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

// NewKeeper creates new instances of the nft Keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, ak types.AccountKeeper, bk types.BankKeeper) Keeper {
	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		accountKeeper: ak,
		bankKeeper:    bk,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("AutonomyNetwork/%s", types.ModuleName))
}

func (k Keeper) CreateDenom(ctx sdk.Context, id, name, symbol, description, previewURI string,
	creator string, community_id string, depedent_collections []string) error {
	return k.SetDenom(ctx, types.NewDenom(id, name, symbol, description, previewURI, creator, community_id, depedent_collections))
}

// MintNFT mints an NFT and manages that NFTs existence within Collections and Owners
func (k Keeper) MintNFT(ctx sdk.Context,
	denomID, nftID, royalties string, transferable bool,
	owner, creator sdk.AccAddress, metadata types.Metadata, attributes string) error {
	if !k.HasDenomID(ctx, denomID) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not exists", denomID)
	}

	if k.HasNFT(ctx, denomID, nftID) {
		return sdkerrors.Wrapf(types.ErrNFTAlreadyExists, "NFT %s already exists in collection %s", nftID, denomID)
	}

	decValue, err := sdk.NewDecFromStr(royalties)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidNFT, "unable to parse the royalities %s", err.Error())
	}

	if !(decValue.GTE(sdk.ZeroDec()) && decValue.LTE(sdk.OneDec())) {
		return sdkerrors.Wrapf(types.ErrInvalidNFT, "royalities in between 0 to 1; given :%s", decValue.String())
	}
	k.SetNFT(ctx, denomID, types.NewBaseNFT(
		nftID,
		metadata,
		owner,
		transferable,
		royalties,
		creator,
		ctx.BlockTime(),
		attributes,
	))
	k.setOwner(ctx, denomID, nftID, owner)
	k.increaseSupply(ctx, denomID)
	return nil
}

// EditNFT updates an already existing NFTs
func (k Keeper) UpdateNFT(ctx sdk.Context,
	denomID, tokenID, name, description, royalties string,
	owner sdk.AccAddress) error {
	if !k.HasDenomID(ctx, denomID) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not exists", denomID)
	}

	nft, err := k.Authorize(ctx, denomID, tokenID, owner)
	if err != nil {
		return err
	}

	if name != "[do-not-modify]" {
		nft.Metadata.Name = name
	}

	if description != "[do-not-modify]" {
		nft.Metadata.Description = description
	}

	if royalties != "[do-not-modify]" {
		nft.Royalties = royalties
	}

	k.SetNFT(ctx, denomID, nft)
	return nil
}

// TransferOwner gets all the ID Collections owned by an address
func (k Keeper) TransferOwner(ctx sdk.Context,
	denomID, tokenID string,
	srcOwner, dstOwner sdk.AccAddress) error {
	if !k.HasDenomID(ctx, denomID) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not exists", denomID)
	}

	nft, err := k.Authorize(ctx, denomID, tokenID, srcOwner)
	if err != nil {
		return err
	}

	if !nft.Transferable {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "nft %s is not transferable", nft.Id)
	}

	nft.Owner = dstOwner.String()

	k.SetNFT(ctx, denomID, nft)
	k.swapOwner(ctx, denomID, tokenID, srcOwner, dstOwner)
	return nil
}

func (k Keeper) SellNFT(ctx sdk.Context, id, denomId string, price string, seller sdk.AccAddress) error {

	if !k.HasDenomID(ctx, denomId) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denomId %s does not exist", denomId)
	}

	if !k.HasNFT(ctx, denomId, id) {
		return sdkerrors.Wrapf(types.ErrInvalidNFT, "nft %s does not exist in collection %s", id, denomId)
	}

	nft, err := k.Authorize(ctx, denomId, id, seller)
	if err != nil {
		return err
	}

	if !nft.IsTransferable() {
		return sdkerrors.Wrapf(types.ErrTransfer, "nft %s is not transferable", id)
	}

	nft.Listed = true

	k.SetNFT(ctx, denomId, nft)
	k.SetNFTMarketPlace(ctx, types.NewMarketPlace(
		id,
		denomId,
		price,
		seller,
	))
	return nil
}

func (k Keeper) BuyNFT(ctx sdk.Context, id, denom_id string, buyer sdk.AccAddress) error {
	if !k.HasDenomID(ctx, denom_id) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom %s does not exist", denom_id)
	}

	if !k.HasNFT(ctx, denom_id, id) {
		return sdkerrors.Wrapf(types.ErrInvalidNFT, "nft %s does not exist in collection %s", id, denom_id)
	}

	orderNFT, err := k.GetMarketPlaceNFT(ctx, denom_id, id)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidNFT, "unable to find nft in market place %s", err.Error())
	}

	if orderNFT.GetFilled() == true {
		return sdkerrors.Wrapf(types.ErrFilledNFT, "%s is already filled", orderNFT.GetNFTID())
	}

	priceStr := orderNFT.GetPrice()
	price, err := sdk.ParseDecCoin(priceStr)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidNFT, "unable to parse the  nft price in market place %s", err.Error())
	}

	buyerAccount := k.bankKeeper.GetBalance(ctx, buyer, price.Denom)
	priceCoin, err := sdk.ParseCoinNormalized(priceStr)
	if buyerAccount.IsLT(priceCoin) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "buyer does not have balance")
	}

	nft, err := k.GetNFT(ctx, denom_id, id)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidNFT, "unable to get the nft  %s", err.Error())
	}

	royaltyDec, err := sdk.NewDecFromStr(nft.GetRoyalties())
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidNFT, "unable to convert royalities  %s", err.Error())
	}

	royalty := royaltyDec.Quo(sdk.OneDec())

	creatorRoyalty := royalty.Mul(price.Amount)
	creatorToken := sdk.NewCoin(price.Denom, creatorRoyalty.TruncateInt())

	if creatorToken.IsPositive() {
		err = k.bankKeeper.SendCoins(ctx, buyer, nft.GetCreator(), sdk.Coins{creatorToken})
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "unable to transfer tokens  %s", err.Error())
		}
	}

	sellerAmount := price.Amount.Sub(creatorRoyalty)
	sellerTokens := sdk.NewCoin(price.Denom, sellerAmount.TruncateInt())

	if sellerTokens.IsPositive() {
		err = k.bankKeeper.SendCoins(ctx, buyer, nft.GetOwner(), sdk.Coins{sellerTokens})
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "unable to transfer tokens  %s", err.Error())
		}
	}

	nft1 := nft.(types.NFT)
	nft1.Owner = buyer.String()
	nft1.Listed = false
	k.SetNFT(ctx, denom_id, nft1)

	orderNFT1 := orderNFT.(types.MarketPlace)
	orderNFT1.Buyer = buyer.String()
	orderNFT1.Filled = true
	k.SetNFTMarketPlace(ctx, orderNFT1)
	k.swapOwner(ctx, denom_id, id, nft.GetOwner(), buyer)
	return nil
}
