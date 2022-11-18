package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/nasa70/car-auction/x/carauction/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdCreateLot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-lot [assets]",
		Short: "Broadcast message create-lot",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCastAssets := strings.Split(args[0], listSeparator)
			argAssets := make([]uint64, len(argCastAssets))
			for i, arg := range argCastAssets {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				argAssets[i] = value
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateLot(
				clientCtx.GetFromAddress().String(),
				argAssets,
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
