package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var checksCloudStatusGroupsCmd = &cobra.Command{
	Use:   "cloudstatus-groups",
	Short: "Discover cloud status providers usable in cloudstatus checks",
}

func init() {
	checksCmd.AddCommand(checksCloudStatusGroupsCmd)
}

var (
	checksCloudStatusGroupsListFlags = upapi.CloudStatusGroupListOptions{
		Page:     1,
		PageSize: 100,
	}
	checksCloudStatusGroupsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List cloud status provider groups",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(checksCloudStatusGroupsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(checksCloudStatusGroupsListCmd.Flags(), &checksCloudStatusGroupsListFlags)
	if err != nil {
		panic(err)
	}
	checksCloudStatusGroupsCmd.AddCommand(checksCloudStatusGroupsListCmd)
}

func checksCloudStatusGroupsList(ctx context.Context) ([]upapi.CloudStatusGroupListItem, error) {
	result, err := api.Checks().ListCloudStatusGroups(ctx, checksCloudStatusGroupsListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var checksCloudStatusServicesCmd = &cobra.Command{
	Use:   "cloudstatus-services",
	Short: "Discover cloud status services usable in cloudstatus checks",
}

func init() {
	checksCmd.AddCommand(checksCloudStatusServicesCmd)
}

var (
	checksCloudStatusServicesListFlags = upapi.CloudStatusServiceListOptions{
		Page:     1,
		PageSize: 100,
	}
	checksCloudStatusServicesListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List cloud status services (filter via --group=<id|name substring>)",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(checksCloudStatusServicesList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(checksCloudStatusServicesListCmd.Flags(), &checksCloudStatusServicesListFlags)
	if err != nil {
		panic(err)
	}
	checksCloudStatusServicesCmd.AddCommand(checksCloudStatusServicesListCmd)
}

func checksCloudStatusServicesList(ctx context.Context) ([]upapi.CloudStatusService, error) {
	result, err := api.Checks().ListCloudStatusServices(ctx, checksCloudStatusServicesListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}
