package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var spIncidentsCmd = &cobra.Command{
	Use:     "incidents",
	Aliases: []string{"incident", "inc"},
	Short:   "Manage status page incidents",
	Args:    cobra.NoArgs,
}

func init() {
	statusPagesCmd.AddCommand(spIncidentsCmd)
}

var (
	spIncidentsListFlags = upapi.StatusPageIncidentListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	spIncidentsListCmd = &cobra.Command{
		Use:     "list <status-page-pk>",
		Aliases: []string{"ls"},
		Short:   "List status page incidents",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spIncidentsList(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spIncidentsListCmd.Flags(), &spIncidentsListFlags)
	if err != nil {
		panic(err)
	}
	spIncidentsCmd.AddCommand(spIncidentsListCmd)
}

func spIncidentsList(ctx context.Context, pkstr string) ([]upapi.StatusPageIncident, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	result, err := api.StatusPages().Incidents(upapi.PrimaryKey(pk)).List(ctx, spIncidentsListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var spIncidentsGetCmd = &cobra.Command{
	Use:     "get <status-page-pk> <incident-pk>",
	Aliases: []string{"show"},
	Short:   "Get a status page incident",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spIncidentsGet(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spIncidentsCmd.AddCommand(spIncidentsGetCmd)
}

func spIncidentsGet(ctx context.Context, spPKStr, incPKStr string) (*upapi.StatusPageIncident, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	incPK, err := parsePK(incPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Incidents(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(incPK))
}

// spIncidentFlags is a CLI-bindable subset of StatusPageIncident.
// Complex nested fields (Updates, AffectedComponents) are not bindable
// to CLI flags and must be managed via the API directly.
type spIncidentFlags struct {
	Name                             string `json:"name"`
	IncidentType                     string `json:"incident_type,omitempty" flag:"incident-type"`
	StartsAt                         string `json:"starts_at" flag:"starts-at"`
	EndsAt                           string `json:"ends_at,omitempty" flag:"ends-at"`
	IncludeInGlobalMetrics           bool   `json:"include_in_global_metrics,omitempty" flag:"include-in-global-metrics"`
	UpdateComponentStatus            bool   `json:"update_component_status,omitempty" flag:"update-component-status"`
	NotifySubscribers                bool   `json:"notify_subscribers,omitempty" flag:"notify-subscribers"`
	SendMaintenanceStartNotification bool   `json:"send_maintenance_start_notification,omitempty" flag:"send-maintenance-start-notification"`
}

func (f spIncidentFlags) toIncident() upapi.StatusPageIncident {
	return upapi.StatusPageIncident{
		Name:                             f.Name,
		IncidentType:                     f.IncidentType,
		StartsAt:                         f.StartsAt,
		EndsAt:                           f.EndsAt,
		IncludeInGlobalMetrics:           f.IncludeInGlobalMetrics,
		UpdateComponentStatus:            f.UpdateComponentStatus,
		NotifySubscribers:                f.NotifySubscribers,
		SendMaintenanceStartNotification: f.SendMaintenanceStartNotification,
	}
}

var (
	spIncidentsCreateFlags spIncidentFlags
	spIncidentsCreateCmd   = &cobra.Command{
		Use:     "create <status-page-pk>",
		Aliases: []string{"new"},
		Short:   "Create a status page incident",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spIncidentsCreate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spIncidentsCreateCmd.Flags(), &spIncidentsCreateFlags)
	if err != nil {
		panic(err)
	}
	spIncidentsCmd.AddCommand(spIncidentsCreateCmd)
}

func spIncidentsCreate(ctx context.Context, pkstr string) (*upapi.StatusPageIncident, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Incidents(upapi.PrimaryKey(pk)).Create(ctx, spIncidentsCreateFlags.toIncident())
}

var (
	spIncidentsUpdateFlags spIncidentFlags
	spIncidentsUpdateCmd   = &cobra.Command{
		Use:     "update <status-page-pk> <incident-pk>",
		Aliases: []string{"up"},
		Short:   "Update a status page incident",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spIncidentsUpdate(cmd.Context(), args[0], args[1]))
		},
	}
)

func init() {
	err := Bind(spIncidentsUpdateCmd.Flags(), &spIncidentsUpdateFlags)
	if err != nil {
		panic(err)
	}
	spIncidentsCmd.AddCommand(spIncidentsUpdateCmd)
}

func spIncidentsUpdate(ctx context.Context, spPKStr, incPKStr string) (*upapi.StatusPageIncident, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	incPK, err := parsePK(incPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Incidents(upapi.PrimaryKey(spPK)).Update(ctx, upapi.PrimaryKey(incPK), spIncidentsUpdateFlags.toIncident())
}

var spIncidentsDeleteCmd = &cobra.Command{
	Use:     "delete <status-page-pk> <incident-pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Delete a status page incident",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spIncidentsDelete(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spIncidentsCmd.AddCommand(spIncidentsDeleteCmd)
}

func spIncidentsDelete(ctx context.Context, spPKStr, incPKStr string) (*upapi.StatusPageIncident, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	incPK, err := parsePK(incPKStr)
	if err != nil {
		return nil, err
	}
	obj, err := api.StatusPages().Incidents(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(incPK))
	if err != nil {
		return nil, err
	}
	return obj, api.StatusPages().Incidents(upapi.PrimaryKey(spPK)).Delete(ctx, upapi.PrimaryKey(incPK))
}
