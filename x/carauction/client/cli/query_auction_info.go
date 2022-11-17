package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/nasa70/car-auction/x/carauction/types"
	"github.com/spf13/cobra"
)

func CmdShowAuctionInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-auction-info",
		Short: "shows auctionInfo",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetAuctionInfoRequest{}

			res, err := queryClient.AuctionInfo(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
