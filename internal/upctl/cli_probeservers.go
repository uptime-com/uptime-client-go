package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var (
	probeserversCmd = &cobra.Command{
		Use:     "servers",
		Aliases: []string{"locations"},
		Short:   "List check servers/locations",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(probeservers(cmd.Context()))
		},
	}
)

func init() {
	cmd.AddCommand(probeserversCmd)
}

func probeservers(ctx context.Context) ([]upapi.ProbeServer, error) {
	result, err := api.ProbeServers().List(ctx)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}
