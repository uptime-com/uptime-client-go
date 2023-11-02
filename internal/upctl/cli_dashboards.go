package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var dashboardsCmd = &cobra.Command{
	Use:   "dashboards",
	Short: "Manage dashboards",
}

func init() {
	cmd.AddCommand(dashboardsCmd)
}

var (
	dashboardsListFlags = upapi.DashboardListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	dashboardsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List dashboards",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(dashboardsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(dashboardsListCmd.Flags(), &dashboardsListFlags)
	if err != nil {
		panic(err)
	}
	dashboardsCmd.AddCommand(dashboardsListCmd)
}

func dashboardsList(ctx context.Context) ([]upapi.Dashboard, error) {
	return api.Dashboards().List(ctx, dashboardsListFlags)
}

var (
	dashboardsCreateFlags = upapi.Dashboard{}
	dashboardsCreateCmd   = &cobra.Command{
		Use:     "create <type>",
		Aliases: []string{"add"},
		Short:   "Create a new check",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(dashboardsCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(dashboardsCreateCmd.Flags(), &dashboardsCreateFlags)
	if err != nil {
		panic(err)
	}
	dashboardsCmd.AddCommand(dashboardsCreateCmd)
}

func dashboardsCreate(ctx context.Context) (*upapi.Dashboard, error) {
	return api.Dashboards().Create(ctx, dashboardsCreateFlags)
}

var (
	dashboardsUpdateFlags = upapi.Dashboard{}
	dashboardsUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up"},
		Short:   "Update existing dashboard",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(dashboardsUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(dashboardsUpdateCmd.Flags(), &dashboardsUpdateFlags)
	if err != nil {
		panic(err)
	}
	dashboardsCmd.AddCommand(dashboardsUpdateCmd)
}

func dashboardsUpdate(ctx context.Context, arg string) (*upapi.Dashboard, error) {
	pk, err := parsePK(arg)
	if err != nil {
		return nil, err
	}
	return api.Dashboards().Update(ctx, upapi.PrimaryKey(pk), dashboardsUpdateFlags)
}

var dashboardsDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Short:   "Delete a tag",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(dashboardsDelete(cmd.Context(), args[0]))
	},
}

func init() {
	dashboardsCmd.AddCommand(dashboardsDeleteCmd)
}

func dashboardsDelete(ctx context.Context, pkstr string) (*upapi.Dashboard, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	obj, err := api.Dashboards().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return obj, api.Dashboards().Delete(ctx, upapi.PrimaryKey(pk))
}
