package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

//
import (
	//	"fmt"
	//	"strings"
	//
	//	"github.com/cosmos/cosmos-sdk/client"
	//	"github.com/cosmos/cosmos-sdk/client/flags"
	//	"github.com/cosmos/cosmos-sdk/client/tx"
	//	sdk "github.com/cosmos/cosmos-sdk/types"
	//	"github.com/cosmos/cosmos-sdk/version"
	//	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	//
	"github.com/AutonomyNetwork/nft/types"
)

//
// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "NFT transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdCreateDenom(),
		GetCmdMintNFT(),
		//GetCmdEditNFT(),
		//GetCmdTransferNFT(),
		//GetCmdBurnNFT(),
	)

	return txCmd
}

//
// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdCreateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create [denom]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new denom.
Example:
$ %s tx nft create --from=<key-name> --name=<name> --data=<data> --symbol=<symbol> --description=<description> --preview_uri=<preview_uri> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, _ := client.GetClientTxContext(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			fmt.Println("Creator address: ", clientCtx.FromAddress)

			msg := types.NewMsgCreateDenom(
				args[0],
				viper.GetString(FlagSymbol),
				viper.GetString(FlagDenomDescription),
				viper.GetString(FlagPreviewURI),
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsCreateDenom)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

//
// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use: "mint [denomID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Mint an NFT and set the owner to the recipient.
Example:
$ %s tx nft mint [denomID] --media_uri=<media_uri> --preview_uri=<preview_uri> --name=<name> --description=<description> --data=<data> --transferable=<transferable> --creator=<creator> --royalties=<royalties> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			name := viper.GetString(FlagTokenName)
			description := viper.GetString(FlagDenomDescription)
			media_uri := viper.GetString(FlagMediaURI)
			previewURI := viper.GetString(FlagPreviewURI)

			metaData := types.Metadata{
				Name:        name,
				Description: description,
				MediaURI:    media_uri,
				PreviewURI:  previewURI,
			}

			msg := types.NewMsgMintNFT(
				args[0],
				viper.GetString(FlagTokenData),
				clientCtx.GetFromAddress().String(),
				viper.GetString(FlagRoyalties),
				metaData,
				viper.GetBool(FlagTransferable),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsMintNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

//
//// GetCmdEditNFT is the CLI command for sending an MsgEditNFT transaction
//func GetCmdEditNFT() *cobra.Command {
//	cmd := &cobra.Command{
//		Use: "edit [denomID] [tokenID]",
//		Long: strings.TrimSpace(
//			fmt.Sprintf(`Edit the tokenData of an NFT.
//Example:
//$ %s tx nft edit [denomID] [tokenID] --uri=<uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
//				version.AppName,
//			),
//		),
//		Args: cobra.ExactArgs(2),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			clientCtx := client.GetClientContextFromCmd(cmd)
//			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
//			if err != nil {
//				return err
//			}
//
//			msg := types.NewMsgEditNFT(
//				args[1],
//				args[0],
//				viper.GetString(FlagTokenName),
//				viper.GetString(FlagTokenURI),
//				viper.GetString(FlagTokenData),
//				clientCtx.GetFromAddress(),
//			)
//			if err := msg.ValidateBasic(); err != nil {
//				return err
//			}
//			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
//		},
//	}
//	cmd.Flags().AddFlagSet(FsEditNFT)
//	flags.AddTxFlagsToCmd(cmd)
//
//	return cmd
//}
//
//// GetCmdTransferNFT is the CLI command for sending a TransferNFT transaction
//func GetCmdTransferNFT() *cobra.Command {
//	cmd := &cobra.Command{
//		Use: "transfer [recipient] [denomID] [tokenID]",
//		Long: strings.TrimSpace(
//			fmt.Sprintf(`Transfer a NFT to a recipient.
//Example:
//$ %s tx nft transfer [recipient] [denomID] [tokenID] --uri=<uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
//				version.AppName,
//			),
//		),
//		Args: cobra.ExactArgs(3),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			clientCtx := client.GetClientContextFromCmd(cmd)
//			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
//			if err != nil {
//				return err
//			}
//
//			recipient, err := sdk.AccAddressFromBech32(args[0])
//			if err != nil {
//				return err
//			}
//
//			msg := types.NewMsgTransferNFT(
//				args[2],
//				args[1],
//				viper.GetString(FlagTokenName),
//				viper.GetString(FlagTokenURI),
//				viper.GetString(FlagTokenData),
//				clientCtx.GetFromAddress(),
//				recipient,
//			)
//			if err := msg.ValidateBasic(); err != nil {
//				return err
//			}
//			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
//		},
//	}
//	cmd.Flags().AddFlagSet(FsTransferNFT)
//	flags.AddTxFlagsToCmd(cmd)
//
//	return cmd
//}
//
//// GetCmdBurnNFT is the CLI command for sending a BurnNFT transaction
//func GetCmdBurnNFT() *cobra.Command {
//	cmd := &cobra.Command{
//		Use: "burn [denomID] [tokenID]",
//		Long: strings.TrimSpace(
//			fmt.Sprintf(`Burn an NFT.
//Example:
//$ %s tx nft burn [denomID] [tokenID] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
//				version.AppName,
//			),
//		),
//		Args: cobra.ExactArgs(2),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			clientCtx := client.GetClientContextFromCmd(cmd)
//			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
//			if err != nil {
//				return err
//			}
//
//			msg := types.NewMsgBurnNFT(clientCtx.GetFromAddress(), args[1], args[0])
//			if err := msg.ValidateBasic(); err != nil {
//				return err
//			}
//			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
//		},
//	}
//	flags.AddTxFlagsToCmd(cmd)
//
//	return cmd
//}
