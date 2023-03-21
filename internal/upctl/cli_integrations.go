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
	integrationsCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"add"},
		Short:   "Create a new integration",
		Args:    cobra.NoArgs,
	}
)

func integrationsCreateSubcommand(name string, flags any, fn func(context.Context) (*upapi.Integration, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Create a new " + name + " integration",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(fn(cmd.Context()))
		},
	}
	err := Bind(cmd.Flags(), flags)
	if err != nil {
		panic(err)
	}
	return cmd
}

var (
	integrationsCreateCachetFlags          upapi.IntegrationCachet
	integrationsCreateDatadogFlags         upapi.IntegrationDatadog
	integrationsCreateGeckoboardFlags      upapi.IntegrationGeckoboard
	integrationsCreateJiraServicedeskFlags upapi.IntegrationJiraServicedesk
	integrationsCreateKlipfolioFlags       upapi.IntegrationKlipfolio
	integrationsCreateLibratoFlags         upapi.IntegrationLibrato
	integrationsCreateMicrosoftTeamsFlags  upapi.IntegrationMicrosoftTeams
	integrationsCreateOpsgenieFlags        upapi.IntegrationOpsgenie
	integrationsCreatePagerdutyFlags       upapi.IntegrationPagerduty
	integrationsCreatePushbulletFlags      upapi.IntegrationPushbullet
	integrationsCreatePushoverFlags        upapi.IntegrationPushover
	integrationsCreateSlackFlags           upapi.IntegrationSlack
	integrationsCreateStatusFlags          upapi.IntegrationStatus
	integrationsCreateStatuspageFlags      upapi.IntegrationStatuspage
	integrationsCreateTwitterFlags         upapi.IntegrationTwitter
	integrationsCreateVictoropsFlags       upapi.IntegrationVictorops
	integrationsCreateWavefrontFlags       upapi.IntegrationWavefront
	integrationsCreateWebhookFlags         upapi.IntegrationWebhook
	integrationsCreateZapierFlags          upapi.IntegrationZapier
)

func init() {
	integrationsCreateCmd.AddCommand(
		integrationsCreateSubcommand("cachet", &integrationsCreateCachetFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateCachet(ctx, integrationsCreateCachetFlags)
		}),
		integrationsCreateSubcommand("datadog", &integrationsCreateDatadogFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateDatadog(ctx, integrationsCreateDatadogFlags)
		}),
		integrationsCreateSubcommand("geckoboard", &integrationsCreateGeckoboardFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateGeckoboard(ctx, integrationsCreateGeckoboardFlags)
		}),
		integrationsCreateSubcommand("jira-servicedesk", &integrationsCreateJiraServicedeskFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateJiraServicedesk(ctx, integrationsCreateJiraServicedeskFlags)
		}),
		integrationsCreateSubcommand("klipfolio", &integrationsCreateKlipfolioFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateKlipfolio(ctx, integrationsCreateKlipfolioFlags)
		}),
		integrationsCreateSubcommand("librato", &integrationsCreateLibratoFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateLibrato(ctx, integrationsCreateLibratoFlags)
		}),
		integrationsCreateSubcommand("microsoft-teams", &integrationsCreateMicrosoftTeamsFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateMicrosoftTeams(ctx, integrationsCreateMicrosoftTeamsFlags)
		}),
		integrationsCreateSubcommand("opsgenie", &integrationsCreateOpsgenieFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateOpsgenie(ctx, integrationsCreateOpsgenieFlags)
		}),
		integrationsCreateSubcommand("pagerduty", &integrationsCreatePagerdutyFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreatePagerduty(ctx, integrationsCreatePagerdutyFlags)
		}),
		integrationsCreateSubcommand("pushbullet", &integrationsCreatePushbulletFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreatePushbullet(ctx, integrationsCreatePushbulletFlags)
		}),
		integrationsCreateSubcommand("pushover", &integrationsCreatePushoverFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreatePushover(ctx, integrationsCreatePushoverFlags)
		}),
		integrationsCreateSubcommand("slack", &integrationsCreateSlackFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateSlack(ctx, integrationsCreateSlackFlags)
		}),
		integrationsCreateSubcommand("status", &integrationsCreateStatusFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateStatus(ctx, integrationsCreateStatusFlags)
		}),
		integrationsCreateSubcommand("statuspage", &integrationsCreateStatuspageFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateStatuspage(ctx, integrationsCreateStatuspageFlags)
		}),
		integrationsCreateSubcommand("twitter", &integrationsCreateTwitterFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateTwitter(ctx, integrationsCreateTwitterFlags)
		}),
		integrationsCreateSubcommand("victorops", &integrationsCreateVictoropsFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateVictorops(ctx, integrationsCreateVictoropsFlags)
		}),
		integrationsCreateSubcommand("wavefront", &integrationsCreateWavefrontFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateWavefront(ctx, integrationsCreateWavefrontFlags)
		}),
		integrationsCreateSubcommand("webhook", &integrationsCreateWebhookFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateWebhook(ctx, integrationsCreateWebhookFlags)
		}),
		integrationsCreateSubcommand("zapier", &integrationsCreateZapierFlags, func(ctx context.Context) (*upapi.Integration, error) {
			return api.Integrations().CreateZapier(ctx, integrationsCreateZapierFlags)
		}),
	)
	integrationsCmd.AddCommand(integrationsCreateCmd)
}

var (
	integrationsDeleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "Delete a tag",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(integrationsDelete(cmd.Context(), args[0]))
		},
	}
)

func init() {
	integrationsCmd.AddCommand(integrationsDeleteCmd)
}

func integrationsDelete(ctx context.Context, pkstr string) (*upapi.Integration, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	integration, err := api.Integrations().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return integration, api.Integrations().Delete(ctx, upapi.PrimaryKey(pk))
}
