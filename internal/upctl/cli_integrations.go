package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var (
	integrationsCmd = &cobra.Command{
		Use:     "integrations",
		Aliases: []string{"integration", "int"},
		Short:   "Manage integrations",
	}
)

func init() {
	cmd.AddCommand(integrationsCmd)
}

var (
	integrationsListFlags = upapi.IntegrationListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	integrationsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List integrations",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(integrationsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(integrationsListCmd.Flags(), &integrationsListFlags)
	if err != nil {
		panic(err)
	}
	integrationsCmd.AddCommand(integrationsListCmd)
}

func integrationsList(ctx context.Context) ([]upapi.Integration, error) {
	return api.Integrations().List(ctx, integrationsListFlags)
}

var (
	integrationsCreateFlags = upapi.Integration{}
	integrationsCreateCmd   = &cobra.Command{
		Use:     "create <type>",
		Aliases: []string{"add"},
		Short:   "Create a new check",
		Args:    cobra.ExactArgs(1),
		ValidArgs: []string{
			"cachet",
			"webhook",
			"datadog",
			"geckoboard",
			"jiraservicedesk",
			"klipfolio",
			"librato",
			"microsoft_teams",
			"opsgenie",
			"pagerduty",
			"pushbullet",
			"pushover",
			"slack",
			"victorops",
			"status",
			"statuspage",
			"twitter",
			"wavefront",
			"zapier",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			integrationsCreateFlags.Module = args[0]
			return output(integrationsCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(integrationsCreateCmd.Flags(), &integrationsCreateFlags)
	if err != nil {
		panic(err)
	}
	integrationsCmd.AddCommand(integrationsCreateCmd)
}

func integrationsCreate(ctx context.Context) (*upapi.Integration, error) {
	return api.Integrations().Create(ctx, integrationsCreateFlags)
}

var (
	integrationsUpdateFlags = upapi.Integration{}
	integrationsUpdateCmd   = &cobra.Command{
		Use:   "update <pk>",
		Short: "Update existing check",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(integrationsUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(integrationsUpdateCmd.Flags(), &integrationsUpdateFlags)
	if err != nil {
		panic(err)
	}
	integrationsCmd.AddCommand(integrationsUpdateCmd)
}

func integrationsUpdate(ctx context.Context, pkStr string) (*upapi.Integration, error) {
	pk, err := parsePK(pkStr)
	if err != nil {
		return nil, err
	}
	integration, err := api.Integrations().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	err = ShallowCopy(&integrationsUpdateFlags, integration)
	if err != nil {
		return nil, err
	}
	return api.Integrations().Update(ctx, integrationsUpdateFlags)
}

var (
	integrationsDeleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "Delete a tag",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(integrationsDelete(cmd.Context(), args[0]))
		},
	}
)

func init() {
	integrationsCmd.AddCommand(integrationsDeleteCmd)
}

func integrationsDelete(ctx context.Context, pkstr string) (*upapi.Tag, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	tag, err := api.Tags().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return tag, api.Tags().Delete(ctx, upapi.PrimaryKey(pk))
}
