package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/nasa70/car-auction/x/carauction/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-asset [desription]",
		Short: "Broadcast message create_asset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDesription := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAsset(
				clientCtx.GetFromAddress().String(),
				argDesription,
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
