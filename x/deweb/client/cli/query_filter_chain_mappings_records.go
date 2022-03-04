package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/deweb-services/deweb/x/deweb/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdFilterChainMappingsRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "filter-chain-mappings-records [owner] [address] [chain] [deleted] [limit] [offset]",
		Short: "Query filter_chain_mappings_records",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwner := args[0]

			var reqAddress string
			if len(args) > 1 {
				reqAddress = args[1]
			}

			var reqChain string
			if len(args) > 2 {
				reqChain = args[2]
			}

			var reqDeleted bool
			if len(args) > 3 {
				reqDeleted, err = strconv.ParseBool(args[3])
				if err != nil {
					return fmt.Errorf("cannot parse parameter deleted: %w", err)
				}
			}

			var reqLimit int64
			if len(args) > 4 {
				reqLimit, err = strconv.ParseInt(args[4], 10, 32)
				if err != nil {
					return fmt.Errorf("cannot parse parameter limit: %w", err)
				}
			} else {
				reqLimit = 10
			}

			var reqOffset int64
			if len(args) > 5 {
				reqOffset, err = strconv.ParseInt(args[5], 10, 32)
				if err != nil {
					return fmt.Errorf("cannot parse parameter offset: %w", err)
				}
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryFilterChainMappingsRecordsRequest{
				Owner:   reqOwner,
				Address: reqAddress,
				Chain:   reqChain,
				Deleted: reqDeleted,
				Limit:   int32(reqLimit),
				Offset:  int32(reqOffset),
			}

			res, err := queryClient.FilterChainMappingsRecords(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
