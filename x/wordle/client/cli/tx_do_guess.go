package cli

import (
	"strconv"

	"github.com/TropicalDog17/wordle/x/wordle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDoGuess() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "do-guess [game-index] [letter]",
		Short: "Broadcast message doGuess",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGameIndex := args[0]
			argLetter := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDoGuess(
				clientCtx.GetFromAddress().String(),
				argGameIndex,
				argLetter,
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
