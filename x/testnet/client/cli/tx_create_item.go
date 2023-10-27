package cli

import (
	"strconv"

	"github.com/ashishkhuraishy/test-net/x/testnet/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdCreateItem() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-item [question] [options]",
		Short: "Broadcast message createItem",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argQuestion := args[0]
			argOptions := strings.Split(args[1], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateItem(
				clientCtx.GetFromAddress().String(),
				argQuestion,
				argOptions,
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
