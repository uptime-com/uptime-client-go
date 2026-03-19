package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var alertsCmd = &cobra.Command{
	Use:     "alerts",
	Aliases: []string{"alert", "a"},
	Short:   "Browse alerts",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(alertsCmd)
}

var (
	alertsListFlags = upapi.AlertListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "-created_at",
	}
	alertsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List alerts",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(alertsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(alertsListCmd.Flags(), &alertsListFlags)
	if err != nil {
		panic(err)
	}
	alertsCmd.AddCommand(alertsListCmd)
}

func alertsList(ctx context.Context) ([]upapi.AlertItem, error) {
	result, err := api.Alerts().List(ctx, alertsListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var alertsGetCmd = &cobra.Command{
	Use:     "get <pk>",
	Aliases: []string{"show"},
	Short:   "Get an alert",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(alertsGet(cmd.Context(), args[0]))
	},
}

func init() {
	alertsCmd.AddCommand(alertsGetCmd)
}

func alertsGet(ctx context.Context, pkstr string) (*upapi.AlertItem, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Alerts().Get(ctx, upapi.PrimaryKey(pk))
}

var alertsRootCauseCmd = &cobra.Command{
	Use:     "root-cause <pk>",
	Aliases: []string{"rc"},
	Short:   "Get root cause analysis for an alert",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(alertsRootCause(cmd.Context(), args[0]))
	},
}

func init() {
	alertsCmd.AddCommand(alertsRootCauseCmd)
}

func alertsRootCause(ctx context.Context, pkstr string) (*upapi.AlertRootCause, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Alerts().RootCause(ctx, upapi.PrimaryKey(pk))
}

var alertsIgnoreCmd = &cobra.Command{
	Use:   "ignore <pk>",
	Short: "Toggle ignore state of an alert",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(alertsIgnore(cmd.Context(), args[0]))
	},
}

func init() {
	alertsCmd.AddCommand(alertsIgnoreCmd)
}

func alertsIgnore(ctx context.Context, pkstr string) (*upapi.AlertItem, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Alerts().Ignore(ctx, upapi.PrimaryKey(pk))
}
