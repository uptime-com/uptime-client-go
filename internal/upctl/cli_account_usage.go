package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var accountUsageCmd = &cobra.Command{
	Use:     "account-usage",
	Aliases: []string{"usage"},
	Short:   "Show account usage and plan limits",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(accountUsageList(cmd.Context()))
	},
}

func init() {
	cmd.AddCommand(accountUsageCmd)
}

func accountUsageList(ctx context.Context) ([]upapi.AccountUsageItem, error) {
	result, err := api.AccountUsage().List(ctx)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}
