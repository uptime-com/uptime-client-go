package upctl

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var slaReportsCmd = &cobra.Command{
	Use:     "slareports",
	Aliases: []string{"slareports"},
	Short:   "Manage SLA reports",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(slaReportsCmd)
}

var (
	slaReportsListFlags = upapi.SLAReportListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	slaReportsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List slaReports",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(slaReportsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(slaReportsListCmd.Flags(), &slaReportsListFlags)
	if err != nil {
		log.Fatalln(err)
	}
	slaReportsCmd.AddCommand(slaReportsListCmd)
}

func slaReportsList(ctx context.Context) ([]upapi.SLAReport, error) {
	result, err := api.SLAReports().List(ctx, upapi.SLAReportListOptions{
		PageSize: 100,
		Page:     slaReportsListFlags.Page,
	})
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var (
	slaReportsCreateFlags = upapi.SLAReport{}
	slaReportsCreateCmd   = &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Short:   "Create a new SLA report",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(slaReportsCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(slaReportsCreateCmd.Flags(), &slaReportsCreateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	slaReportsCmd.AddCommand(slaReportsCreateCmd)
}

func slaReportsCreate(ctx context.Context) (*upapi.SLAReport, error) {
	return api.SLAReports().Create(ctx, slaReportsCreateFlags)
}

var (
	slaReportsUpdateFlags upapi.SLAReport
	slaReportsUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up"},
		Short:   "Update a SLA report",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(slaReportsUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(slaReportsUpdateCmd.Flags(), &slaReportsUpdateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	slaReportsCmd.AddCommand(slaReportsUpdateCmd)
}

func slaReportsUpdate(ctx context.Context, pkstr string) (*upapi.SLAReport, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.SLAReports().Update(ctx, upapi.PrimaryKey(pk), slaReportsUpdateFlags)
}

var slaReportsDeleteCmd = &cobra.Command{
	Use:     "delete <pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Delete a SLA report",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(slaReportsDelete(cmd.Context(), args[0]))
	},
}

func init() {
	slaReportsCmd.AddCommand(slaReportsDeleteCmd)
}

func slaReportsDelete(ctx context.Context, pkstr string) (*upapi.SLAReport, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	obj, err := api.SLAReports().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return obj, api.SLAReports().Delete(ctx, upapi.PrimaryKey(pk))
}
