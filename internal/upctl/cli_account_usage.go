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
		return output(accountUsageGet(cmd.Context()))
	},
}

func init() {
	cmd.AddCommand(accountUsageCmd)
}

func accountUsageGet(ctx context.Context) (*upapi.AccountUsage, error) {
	return api.AccountUsage().Get(ctx)
}
