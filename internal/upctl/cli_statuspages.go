package upctl

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var statusPagesCmd = &cobra.Command{
	Use:     "statuspages",
	Aliases: []string{"statuspage", "sp"},
	Short:   "Manage status pages",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(statusPagesCmd)
}

var (
	statusPagesListFlags = upapi.StatusPageListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	statusPagesListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List statusPages",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(statusPagesList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(statusPagesListCmd.Flags(), &statusPagesListFlags)
	if err != nil {
		log.Fatalln(err)
	}
	statusPagesCmd.AddCommand(statusPagesListCmd)
}

func statusPagesList(ctx context.Context) ([]upapi.StatusPage, error) {
	result, err := api.StatusPages().List(ctx, upapi.StatusPageListOptions{
		PageSize: 100,
		Page:     statusPagesListFlags.Page,
	})
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var (
	statusPagesCreateFlags = upapi.StatusPage{
		UptimeCalculationType: "BY_INCIDENTS",
		PageType:              "PUBLIC",
		VisibilityLevel:       "UPTIME_USERS",
	}
	statusPagesCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Short:   "Create a new status page",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(statusPagesCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(statusPagesCreateCmd.Flags(), &statusPagesCreateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	statusPagesCmd.AddCommand(statusPagesCreateCmd)
}

func statusPagesCreate(ctx context.Context) (*upapi.StatusPage, error) {
	return api.StatusPages().Create(ctx, statusPagesCreateFlags)
}

var (
	statusPagesUpdateFlags upapi.StatusPage
	statusPagesUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up"},
		Short:   "Update a status page",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(statusPagesUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(statusPagesUpdateCmd.Flags(), &statusPagesUpdateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	statusPagesCmd.AddCommand(statusPagesUpdateCmd)
}

func statusPagesUpdate(ctx context.Context, pkstr string) (*upapi.StatusPage, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Update(ctx, upapi.PrimaryKey(pk), statusPagesUpdateFlags)
}

var (
	statusPagesDeleteCmd = &cobra.Command{
		Use:     "delete <pk>",
		Aliases: []string{"del", "rm", "remove"},
		Short:   "Delete a status page",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(statusPagesDelete(cmd.Context(), args[0]))
		},
	}
)

func init() {
	statusPagesCmd.AddCommand(statusPagesDeleteCmd)
}

func statusPagesDelete(ctx context.Context, pkstr string) (*upapi.StatusPage, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	obj, err := api.StatusPages().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return obj, api.StatusPages().Delete(ctx, upapi.PrimaryKey(pk))
}

// Current Status Commands

var (
	statusPagesCurrentStatusCmd = &cobra.Command{
		Use:     "current-status <status-page-pk>",
		Aliases: []string{"status", "cs"},
		Short:   "Get current status of a status page",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(statusPagesCurrentStatus(cmd.Context(), args[0]))
		},
	}
)

func init() {
	statusPagesCmd.AddCommand(statusPagesCurrentStatusCmd)
}

func statusPagesCurrentStatus(ctx context.Context, pkstr string) (*upapi.StatusPageCurrentStatus, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().CurrentStatus(upapi.PrimaryKey(pk)).Get(ctx)
}

// Status History Commands

var (
	statusPagesStatusHistoryListFlags = upapi.StatusPageStatusHistoryListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "-created_at",
	}
	statusPagesStatusHistoryListCmd = &cobra.Command{
		Use:     "status-history <status-page-pk>",
		Aliases: []string{"history", "sh"},
		Short:   "List status history for a status page",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(statusPagesStatusHistoryList(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(statusPagesStatusHistoryListCmd.Flags(), &statusPagesStatusHistoryListFlags)
	if err != nil {
		log.Fatalln(err)
	}
	statusPagesCmd.AddCommand(statusPagesStatusHistoryListCmd)
}

func statusPagesStatusHistoryList(ctx context.Context, pkstr string) ([]upapi.StatusPageStatusHistory, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	result, err := api.StatusPages().StatusHistory(upapi.PrimaryKey(pk)).List(ctx, statusPagesStatusHistoryListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var (
	statusPagesStatusHistoryGetCmd = &cobra.Command{
		Use:   "status-history-get <status-page-pk> <history-pk>",
		Short: "Get a specific status history entry",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(statusPagesStatusHistoryGet(cmd.Context(), args[0], args[1]))
		},
	}
)

func init() {
	statusPagesCmd.AddCommand(statusPagesStatusHistoryGetCmd)
}

func statusPagesStatusHistoryGet(ctx context.Context, statusPagePKStr, historyPKStr string) (*upapi.StatusPageStatusHistory, error) {
	statusPagePK, err := parsePK(statusPagePKStr)
	if err != nil {
		return nil, err
	}
	historyPK, err := parsePK(historyPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().StatusHistory(upapi.PrimaryKey(statusPagePK)).Get(ctx, upapi.PrimaryKey(historyPK))
}
