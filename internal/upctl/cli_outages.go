package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var outagesCmd = &cobra.Command{
	Use:     "outages",
	Aliases: []string{"outage", "o"},
	Short:   "Browse outages",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(outagesCmd)
}

var (
	outagesListFlags = upapi.OutageListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	outagesListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List outages",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(outagesList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(outagesListCmd.Flags(), &outagesListFlags)
	if err != nil {
		panic(err)
	}
	outagesCmd.AddCommand(outagesListCmd)
}

func outagesList(ctx context.Context) ([]upapi.Outage, error) {
	return api.Outages().List(ctx, upapi.OutageListOptions{
		PageSize: 100,
		Page:     outagesListFlags.Page,
	})
}
