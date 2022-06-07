package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/AutonomyNetwork/nft/types"
)

//

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
		GetCmdMintBulkNFT(),
		GetCmdUpdateNFT(),
		GetCmdTransferNFT(),
		GetCmdSellNFT(),
		GetCmdSellBulkNFT(),
		GetCmdBuyNFT(),
		GetCmdCreateCommunity(),
		GetCmdJoinCommunity(),
	)

	return txCmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdCreateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [denom]",
		Short: "create collection with denom name",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new denom.
Example:
$ %s tx nft create [denom] [community-id] [, seperated dependent_collections] --symbol=<symbol> --description=<description> --preview_uri=<preview_uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, _ := client.GetClientTxContext(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			var collections []string
			if len(args[2]) > 0 {
				collections = strings.Split(args[2], ",")
			}

			msg := types.NewMsgCreateDenom(
				args[0],
				viper.GetString(FlagSymbol),
				viper.GetString(FlagDescription),
				viper.GetString(FlagPreviewURI),
				clientCtx.GetFromAddress().String(),
				args[1],
				collections,
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

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [denomID]",
		Short: "mint nft with repective denom id or collection id",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Mint an NFT and set the owner to the recipient.
Example:
$ %s tx nft mint [denomID] --media_uri=<media_uri> --preview_uri=<preview_uri> --name=<name> --description=<description> --transferable=<transferable> --royalties=<royalties> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			name := viper.GetString(FlagTokenName)
			description := viper.GetString(FlagDescription)
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
	cmd.Flags().AddFlagSet(FsCreateDenom)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

type File struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Attributes  interface{} `json:"attributes"`
	Data        interface{} `json:"data"`
}

//
// GetCmdMintBulkNFT is the CLI command for a MintBulkNFT transaction
func GetCmdMintBulkNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-bulk [denomID] [file-path]",
		Short: "mint-bulk nfts with repective denom id or collection id",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Mint bulk NFTs and set the owner to the recipient.
Example:
$ %s tx nft mint [denomID] [file-path] --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			bz, err := ioutil.ReadFile(args[1]) // just pass the file name
			if err != nil {
				fmt.Print(err)
			}

			var files []File
			var msgs []sdk.Msg

			_ = json.Unmarshal(bz, &files)

			for _, file := range files {

				metaData := types.Metadata{
					Name:        file.Name,
					Description: file.Description,
					MediaURI:    file.Image,
					PreviewURI:  file.Image,
				}
				bz, _ := json.Marshal(file.Data)
				msg := types.NewMsgMintNFT(
					args[0],
					string(bz),
					clientCtx.GetFromAddress().String(),
					"0.1",
					metaData,
					true,
				)

				if err := msg.ValidateBasic(); err != nil {
					return err
				}

				msgs = append(msgs, msg)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdUpdateNFT is the CLI command for sending an MsgEditNFT transaction
func GetCmdUpdateNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [denomID] [tokenID]",
		Short: "update the meta-data of nft based on token and denom id",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Edit the tokenData of an NFT.
Example:
$ %s tx nft update [denomID] [tokenID] --name=<name> --description=<description> --royalties=<royalties> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateNFT(
				args[1],
				args[0],
				viper.GetString(FlagRoyalties),
				viper.GetString(FlagDescription),
				viper.GetString(FlagTokenName),
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsEditNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdTransferNFT is the CLI command for sending a TransferNFT transaction
func GetCmdTransferNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [denomID] [nftID] [recipient]",
		Short: "transfer nft directly to user",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Transfer a NFT to a recipient.
Example:Command
$ %s tx nft transfer [denomID] [nftID] [recipient] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferNFT(
				args[1],
				args[0],
				clientCtx.GetFromAddress().String(),
				args[2],
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsTransferNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdSellNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell [denomID] [NFTID] [price]",
		Short: "Sell nft",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Add an NFT to market place.
Example:
$ %s tx nft sell [denomID] [NFTID] [price] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cliCtx, err = client.ReadPersistentCommandFlags(cliCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgSellNFT(
				args[1],
				args[0],
				args[2],
				cliCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetCmdSellBulkNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-bulk [denomID] [NFTID,price]",
		Short: "Sell bulk nfts",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Add an NFTs to market place.
Example:
$ %s tx nft sell-bulk [denomID] [NFTID:price,NFTID:price] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cliCtx, err = client.ReadPersistentCommandFlags(cliCtx, cmd.Flags())
			if err != nil {
				return err
			}

			var msgs []sdk.Msg
			values := strings.Split(args[1], ",")
			for _, v := range values {
				data := strings.Split(v, ":")
				msg := types.NewMsgSellNFT(
					data[0],
					args[0],
					data[1],
					cliCtx.GetFromAddress().String(),
				)
				if err := msg.ValidateBasic(); err != nil {
					return err
				}

				msgs = append(msgs, msg)
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetCmdBuyNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy [denomID] [NFTID]",
		Short: "buy nft",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Buy an NFT from market place.
Example:
$ %s tx nft buy [denomID] [NFTID] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyNFT(
				args[1],
				args[0],
				clientCtx.GetFromAddress().String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateCommunity is the CLI command for a CreateCommunity transaction
func GetCmdCreateCommunity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-community [name]",
		Short: "Create new community",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new community.
Example:
$ %s tx nft create-community [name]  --description=<description> --preview_uri=<preview_uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
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

			msg := types.NewMsgCreateCommunity(
				args[0],
				viper.GetString(FlagDescription),
				clientCtx.GetFromAddress().String(),
				viper.GetString(FlagPreviewURI),
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

// GetCmdJoinCommunity is the CLI command for a JoinCommunity transaction
func GetCmdJoinCommunity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "join-community [community-id]   ",
		Short: "Join in existing community",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new community.
Example:
$ %s tx nft join-community [community-id]   --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
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

			msg := types.NewMsgJoinCommunity(
				args[0],
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
