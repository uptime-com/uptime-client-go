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
	result, err := api.Integrations().List(ctx, integrationsListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var (
	integrationsGetCmd = &cobra.Command{
		Use:     "get <pk>",
		Aliases: []string{"show"},
		Short:   "Get an integration by ID",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(integrationsGet(cmd.Context(), args[0]))
		},
	}
)

func init() {
	integrationsCmd.AddCommand(integrationsGetCmd)
}

func integrationsGet(ctx context.Context, pkstr string) (*upapi.Integration, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Integrations().Get(ctx, upapi.PrimaryKey(pk))
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

func integrationsUpdateSubcommand(name string, flags any, fn func(context.Context, int64) (*upapi.Integration, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name + " <pk>",
		Short: "Update " + name + " integration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pk, err := parsePK(args[0])
			if err != nil {
				return err
			}
			return output(fn(cmd.Context(), pk))
		},
	}
	err := Bind(cmd.Flags(), flags)
	if err != nil {
		panic(err)
	}
	return cmd
}

var (
	integrationsUpdateCmd = &cobra.Command{
		Use:     "update",
		Aliases: []string{"up"},
		Short:   "Update an integration",
	}
)

var (
	integrationsUpdateCachetFlags          upapi.IntegrationCachet
	integrationsUpdateDatadogFlags         upapi.IntegrationDatadog
	integrationsUpdateGeckoboardFlags      upapi.IntegrationGeckoboard
	integrationsUpdateJiraServicedeskFlags upapi.IntegrationJiraServicedesk
	integrationsUpdateKlipfolioFlags       upapi.IntegrationKlipfolio
	integrationsUpdateLibratoFlags         upapi.IntegrationLibrato
	integrationsUpdateMicrosoftTeamsFlags  upapi.IntegrationMicrosoftTeams
	integrationsUpdateOpsgenieFlags        upapi.IntegrationOpsgenie
	integrationsUpdatePagerdutyFlags       upapi.IntegrationPagerduty
	integrationsUpdatePushbulletFlags      upapi.IntegrationPushbullet
	integrationsUpdatePushoverFlags        upapi.IntegrationPushover
	integrationsUpdateSlackFlags           upapi.IntegrationSlack
	integrationsUpdateStatusFlags          upapi.IntegrationStatus
	integrationsUpdateStatuspageFlags      upapi.IntegrationStatuspage
	integrationsUpdateTwitterFlags         upapi.IntegrationTwitter
	integrationsUpdateVictoropsFlags       upapi.IntegrationVictorops
	integrationsUpdateWavefrontFlags       upapi.IntegrationWavefront
	integrationsUpdateWebhookFlags         upapi.IntegrationWebhook
	integrationsUpdateZapierFlags          upapi.IntegrationZapier
)

func init() {
	integrationsUpdateCmd.AddCommand(
		integrationsUpdateSubcommand("cachet", &integrationsUpdateCachetFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateCachet(ctx, upapi.PrimaryKey(pk), integrationsUpdateCachetFlags)
		}),
		integrationsUpdateSubcommand("datadog", &integrationsUpdateDatadogFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateDatadog(ctx, upapi.PrimaryKey(pk), integrationsUpdateDatadogFlags)
		}),
		integrationsUpdateSubcommand("geckoboard", &integrationsUpdateGeckoboardFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateGeckoboard(ctx, upapi.PrimaryKey(pk), integrationsUpdateGeckoboardFlags)
		}),
		integrationsUpdateSubcommand("jira-servicedesk", &integrationsUpdateJiraServicedeskFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateJiraServiceDesk(ctx, upapi.PrimaryKey(pk), integrationsUpdateJiraServicedeskFlags)
		}),
		integrationsUpdateSubcommand("klipfolio", &integrationsUpdateKlipfolioFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateKlipfolio(ctx, upapi.PrimaryKey(pk), integrationsUpdateKlipfolioFlags)
		}),
		integrationsUpdateSubcommand("librato", &integrationsUpdateLibratoFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateLibrato(ctx, upapi.PrimaryKey(pk), integrationsUpdateLibratoFlags)
		}),
		integrationsUpdateSubcommand("microsoft-teams", &integrationsUpdateMicrosoftTeamsFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateMicrosoftTeams(ctx, upapi.PrimaryKey(pk), integrationsUpdateMicrosoftTeamsFlags)
		}),
		integrationsUpdateSubcommand("opsgenie", &integrationsUpdateOpsgenieFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateOpsgenie(ctx, upapi.PrimaryKey(pk), integrationsUpdateOpsgenieFlags)
		}),
		integrationsUpdateSubcommand("pagerduty", &integrationsUpdatePagerdutyFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdatePagerduty(ctx, upapi.PrimaryKey(pk), integrationsUpdatePagerdutyFlags)
		}),
		integrationsUpdateSubcommand("pushbullet", &integrationsUpdatePushbulletFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdatePushbullet(ctx, upapi.PrimaryKey(pk), integrationsUpdatePushbulletFlags)
		}),
		integrationsUpdateSubcommand("pushover", &integrationsUpdatePushoverFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdatePushover(ctx, upapi.PrimaryKey(pk), integrationsUpdatePushoverFlags)
		}),
		integrationsUpdateSubcommand("slack", &integrationsUpdateSlackFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateSlack(ctx, upapi.PrimaryKey(pk), integrationsUpdateSlackFlags)
		}),
		integrationsUpdateSubcommand("status", &integrationsUpdateStatusFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateStatus(ctx, upapi.PrimaryKey(pk), integrationsUpdateStatusFlags)
		}),
		integrationsUpdateSubcommand("statuspage", &integrationsUpdateStatuspageFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateStatuspage(ctx, upapi.PrimaryKey(pk), integrationsUpdateStatuspageFlags)
		}),
		integrationsUpdateSubcommand("twitter", &integrationsUpdateTwitterFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateTwitter(ctx, upapi.PrimaryKey(pk), integrationsUpdateTwitterFlags)
		}),
		integrationsUpdateSubcommand("victorops", &integrationsUpdateVictoropsFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateVictorops(ctx, upapi.PrimaryKey(pk), integrationsUpdateVictoropsFlags)
		}),
		integrationsUpdateSubcommand("wavefront", &integrationsUpdateWavefrontFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateWavefront(ctx, upapi.PrimaryKey(pk), integrationsUpdateWavefrontFlags)
		}),
		integrationsUpdateSubcommand("webhook", &integrationsUpdateWebhookFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateWebhook(ctx, upapi.PrimaryKey(pk), integrationsUpdateWebhookFlags)
		}),
		integrationsUpdateSubcommand("zapier", &integrationsUpdateZapierFlags, func(ctx context.Context, pk int64) (*upapi.Integration, error) {
			return api.Integrations().UpdateZapier(ctx, upapi.PrimaryKey(pk), integrationsUpdateZapierFlags)
		}),
	)
	integrationsCmd.AddCommand(integrationsUpdateCmd)
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
