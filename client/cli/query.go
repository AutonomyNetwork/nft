package cli

import (
	"context"
	"fmt"
	"github.com/AutonomyNetwork/nft/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"strings"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                types.ModuleName,
		Short:              "Querying commands for the NFT module",
		DisableFlagParsing: true,
	}

	queryCmd.AddCommand(
		GetCmdQueryDenom(),
		GetCmdQueryDenoms(),
		GetCmdQueryMarketNFT(),
		GetCmdQueryOwnerCollections(),
		GetCmdQueryNFT(),
		GetCmdQueryMarketPlace(),
	)

	return queryCmd
}

// GetCmdQueryDenoms queries all denoms
func GetCmdQueryDenoms() *cobra.Command {
	cmd := &cobra.Command{
		Use: "denoms",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all denominations of all collections of NFTs
Example:
$ %s query nft denoms`, version.AppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Denoms(context.Background(), &types.QueryDenomsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryDenoms queries the specified denoms
func GetCmdQueryDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use: "denom [denomID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the denominations by the specified denmo name
Example:
$ %s query nft denom <denom>`, version.AppName)),
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

			denom := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denom); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Denom(context.Background(), &types.QueryDenomRequest{
				DenomId: denom,
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp.Denom)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryNFT queries a single NFTs from a collection
func GetCmdQueryNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use: "token [denomID] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query a single NFT from a collection
Example:
$ %s query nft token <denom> <tokenID>`, version.AppName)),
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

			denom := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denom); err != nil {
				return err
			}

			tokenID := strings.TrimSpace(args[1])
			if err := types.ValidateNFTID(tokenID); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.NFT(context.Background(), &types.QueryNFTRequest{
				DenomId: denom,
				Id:      tokenID,
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp.NFT)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryMarketNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use: "marketNFT [denomID] [NFTID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query a single NFT from a market place.
Example:
$ %s query nft marketNFT [denomID] [NFTID]`,
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

			if err := types.ValidateDenomID(args[0]); err != nil {
				return err
			}
			if err := types.ValidateNFTID(args[1]); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.MarketPlaceNFT(context.Background(), &types.QueryMarketPlaceNFTRequest{
				DenomId: args[0],
				Id:      args[1],
			})

			if err != nil {
				return err
			}

			return cliCtx.PrintProto(res.MarketPlace)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdQueryOwnerCollections() *cobra.Command {
	cmd := &cobra.Command{
		Use: "collections [owner]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all NFTs for a owner.
Example:
$ %s query nft collections [owner]`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cliCtx, err = client.ReadPersistentCommandFlags(cliCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.OwnerNFTs(context.Background(), &types.QueryOwnerNFTsRequest{
				Owner: args[0],
			})

			if err != nil {
				return err
			}

			return cliCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdQueryMarketPlace() *cobra.Command {
	cmd := &cobra.Command{
		Use: "marketPlace",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query marketPlace.
Example:
$ %s query nft marketPlace`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cliCtx, err = client.ReadPersistentCommandFlags(cliCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.MarketPlace(context.Background(), &types.QueryMarketPlaceRequest{})
			if err != nil {
				return err
			}

			return cliCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
