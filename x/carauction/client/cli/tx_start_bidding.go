package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nasa70/car-auction/x/carauction/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStartBidding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-bidding [lot-id] [finish-time] [bid]",
		Short: "Broadcast message start-bidding",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argLotId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argFinishTime, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argBid, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStartBidding(
				clientCtx.GetFromAddress().String(),
				argLotId,
				argFinishTime,
				argBid,
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
