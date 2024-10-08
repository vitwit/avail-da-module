package cli

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/vitwit/avail-da-module/x/cada/keeper"
	types "github.com/vitwit/avail-da-module/x/cada/types"
)

// NewTxCmd creates and returns a Cobra command for transaction subcommands related to the cada module.
func NewTxCmd(_ *keeper.Keeper) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      types.ModuleName + " transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(NewUpdateBlobStatusCmd())

	return txCmd
}

// NewUpdateBlobStatusCmd creates and returns a Cobra command to update the status of a blob within a specified range of blocks.
func NewUpdateBlobStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-blob [from_block] [to_block] [status] [avail_height]",
		Short: `update blob status by giving blocks range and status(success|failure)
		 and the avail height at which the blob is stored`,
		Example: "simd update-blob 11 15 success 120",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fromBlock, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			toBlock, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			isSuccess, err := ParseStatus(args[2])
			if err != nil {
				return err
			}

			availHeight, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}

			msg := types.MsgUpdateBlobStatusRequest{
				BlocksRange: &types.Range{
					From: uint64(fromBlock),
					To:   uint64(toBlock),
				},
				ValidatorAddress: clientCtx.GetFromAddress().String(),
				AvailHeight:      uint64(availHeight),
				IsSuccess:        isSuccess,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// ParseStatus converts a string status to a boolean value indicating success or failure.
func ParseStatus(status string) (bool, error) {
	status = strings.ToUpper(status)
	if status == "SUCCESS" {
		return true, nil
	}

	if status == "FAILURE" {
		return false, nil
	}

	return false, errors.New("invalid status")
}
