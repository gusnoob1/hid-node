package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hypersign-protocol/hid-node/x/did/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSchemaCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema-count",
		Short: "Query schemaCount",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySchemaCountRequest{}

			res, err := queryClient.SchemaCount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
