package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var spMetricsCmd = &cobra.Command{
	Use:     "metrics",
	Aliases: []string{"metric"},
	Short:   "Manage status page metrics",
	Args:    cobra.NoArgs,
}

func init() {
	statusPagesCmd.AddCommand(spMetricsCmd)
}

var (
	spMetricsListFlags = upapi.StatusPageMetricListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	spMetricsListCmd = &cobra.Command{
		Use:     "list <status-page-pk>",
		Aliases: []string{"ls"},
		Short:   "List status page metrics",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spMetricsList(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spMetricsListCmd.Flags(), &spMetricsListFlags)
	if err != nil {
		panic(err)
	}
	spMetricsCmd.AddCommand(spMetricsListCmd)
}

func spMetricsList(ctx context.Context, pkstr string) ([]upapi.StatusPageMetric, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	result, err := api.StatusPages().Metrics(upapi.PrimaryKey(pk)).List(ctx, spMetricsListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var spMetricsGetCmd = &cobra.Command{
	Use:     "get <status-page-pk> <metric-pk>",
	Aliases: []string{"show"},
	Short:   "Get a status page metric",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spMetricsGet(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spMetricsCmd.AddCommand(spMetricsGetCmd)
}

func spMetricsGet(ctx context.Context, spPKStr, metricPKStr string) (*upapi.StatusPageMetric, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	metricPK, err := parsePK(metricPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Metrics(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(metricPK))
}

var (
	spMetricsCreateFlags upapi.StatusPageMetric
	spMetricsCreateCmd   = &cobra.Command{
		Use:     "create <status-page-pk>",
		Aliases: []string{"new"},
		Short:   "Create a status page metric",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spMetricsCreate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spMetricsCreateCmd.Flags(), &spMetricsCreateFlags)
	if err != nil {
		panic(err)
	}
	spMetricsCmd.AddCommand(spMetricsCreateCmd)
}

func spMetricsCreate(ctx context.Context, pkstr string) (*upapi.StatusPageMetric, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Metrics(upapi.PrimaryKey(pk)).Create(ctx, spMetricsCreateFlags)
}

var (
	spMetricsUpdateFlags upapi.StatusPageMetric
	spMetricsUpdateCmd   = &cobra.Command{
		Use:     "update <status-page-pk> <metric-pk>",
		Aliases: []string{"up"},
		Short:   "Update a status page metric",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spMetricsUpdate(cmd.Context(), args[0], args[1]))
		},
	}
)

func init() {
	err := Bind(spMetricsUpdateCmd.Flags(), &spMetricsUpdateFlags)
	if err != nil {
		panic(err)
	}
	spMetricsCmd.AddCommand(spMetricsUpdateCmd)
}

func spMetricsUpdate(ctx context.Context, spPKStr, metricPKStr string) (*upapi.StatusPageMetric, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	metricPK, err := parsePK(metricPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Metrics(upapi.PrimaryKey(spPK)).Update(ctx, upapi.PrimaryKey(metricPK), spMetricsUpdateFlags)
}

var spMetricsDeleteCmd = &cobra.Command{
	Use:     "delete <status-page-pk> <metric-pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Delete a status page metric",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spMetricsDelete(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spMetricsCmd.AddCommand(spMetricsDeleteCmd)
}

func spMetricsDelete(ctx context.Context, spPKStr, metricPKStr string) (*upapi.StatusPageMetric, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	metricPK, err := parsePK(metricPKStr)
	if err != nil {
		return nil, err
	}
	obj, err := api.StatusPages().Metrics(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(metricPK))
	if err != nil {
		return nil, err
	}
	return obj, api.StatusPages().Metrics(upapi.PrimaryKey(spPK)).Delete(ctx, upapi.PrimaryKey(metricPK))
}
