package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/AutonomyNetwork/nft/exported"
)

// RegisterLegacyAminoCodec concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "AutonomyNetwork/nft/MsgCreateDenom", nil)
	cdc.RegisterConcrete(&MsgMintNFT{}, "AutonomyNetwork/nft/MsgMintNFT", nil)
	cdc.RegisterConcrete(&MsgUpdateNFT{}, "AutonomyNetwork/nft/MsgUpdateNFT", nil)
	cdc.RegisterConcrete(&MsgTransferNFT{}, "AutonomyNetwork/nft/MsgTransferNFT", nil)
	cdc.RegisterConcrete(&MsgSellNFT{}, "AutonomyNetwork/nft/MsgSellNFT", nil)

	cdc.RegisterInterface((*exported.NFT)(nil), nil)
	cdc.RegisterInterface((*exported.MarketPlace)(nil), nil)
	cdc.RegisterConcrete(&NFT{}, "AutonomyNetwork/nft/NFT", nil)
	cdc.RegisterConcrete(&MarketPlace{}, "AutonomyNetwork/nft/MarketPlace", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgMintNFT{},
		&MsgUpdateNFT{},
		&MsgTransferNFT{},
		&MsgSellNFT{},
	)

	registry.RegisterImplementations((*exported.NFT)(nil), &NFT{})
	registry.RegisterImplementations((*exported.MarketPlace)(nil), &MarketPlace{})

}

var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// return supply protobuf code
func MustMarshalSupply(cdc codec.BinaryCodec, supply uint64) []byte {
	supplyWrap := gogotypes.UInt64Value{Value: supply}
	return cdc.MustMarshal(&supplyWrap)
}

// return th supply
func MustUnMarshalSupply(cdc codec.BinaryCodec, value []byte) uint64 {
	var supplyWrap gogotypes.UInt64Value
	cdc.MustUnmarshal(value, &supplyWrap)
	return supplyWrap.Value
}

// return the tokenID protobuf code
func MustMarshalTokenID(cdc codec.BinaryCodec, tokenID string) []byte {
	tokenIDWrap := gogotypes.StringValue{Value: tokenID}
	return cdc.MustMarshal(&tokenIDWrap)
}

// return th tokenID
func MustUnMarshalTokenID(cdc codec.BinaryCodec, value []byte) string {
	var tokenIDWrap gogotypes.StringValue
	cdc.MustUnmarshal(value, &tokenIDWrap)
	return tokenIDWrap.Value
}
