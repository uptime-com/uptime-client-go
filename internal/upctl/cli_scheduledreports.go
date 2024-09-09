package upctl

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var scheduledReportsCmd = &cobra.Command{
	Use:     "scheduledreports",
	Aliases: []string{"scheduledreports"},
	Short:   "Manage Scheduled reports",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(scheduledReportsCmd)
}

var (
	scheduledReportsListFlags = upapi.ScheduledReportListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	scheduledReportsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List scheduled reports",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(scheduledReportsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(scheduledReportsListCmd.Flags(), &scheduledReportsListFlags)
	if err != nil {
		log.Fatalln(err)
	}
	scheduledReportsCmd.AddCommand(scheduledReportsListCmd)
}

func scheduledReportsList(ctx context.Context) ([]upapi.ScheduledReport, error) {
	return api.ScheduledReports().List(ctx, upapi.ScheduledReportListOptions{
		PageSize: 100,
		Page:     scheduledReportsListFlags.Page,
	})
}

var (
	scheduledReportsCreateFlags = upapi.ScheduledReport{}
	scheduledReportsCreateCmd   = &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Short:   "Create a new Scheduled report",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(scheduledReportsCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(scheduledReportsCreateCmd.Flags(), &scheduledReportsCreateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	scheduledReportsCmd.AddCommand(scheduledReportsCreateCmd)
}

func scheduledReportsCreate(ctx context.Context) (*upapi.ScheduledReport, error) {
	return api.ScheduledReports().Create(ctx, scheduledReportsCreateFlags)
}

var (
	scheduledReportsUpdateFlags upapi.ScheduledReport
	scheduledReportsUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up"},
		Short:   "Update a Scheduled report",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(scheduledReportsUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(scheduledReportsUpdateCmd.Flags(), &scheduledReportsUpdateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	scheduledReportsCmd.AddCommand(scheduledReportsUpdateCmd)
}

func scheduledReportsUpdate(ctx context.Context, pkstr string) (*upapi.ScheduledReport, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.ScheduledReports().Update(ctx, upapi.PrimaryKey(pk), scheduledReportsUpdateFlags)
}

var scheduledReportsDeleteCmd = &cobra.Command{
	Use:     "delete <pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Delete a Scheduled report",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(scheduledReportsDelete(cmd.Context(), args[0]))
	},
}

func init() {
	scheduledReportsCmd.AddCommand(scheduledReportsDeleteCmd)
}

func scheduledReportsDelete(ctx context.Context, pkstr string) (*upapi.ScheduledReport, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	obj, err := api.ScheduledReports().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return obj, api.ScheduledReports().Delete(ctx, upapi.PrimaryKey(pk))
}
