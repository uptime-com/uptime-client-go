package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var spSubscribersCmd = &cobra.Command{
	Use:     "subscribers",
	Aliases: []string{"subscriber", "sub"},
	Short:   "Manage status page subscribers",
	Args:    cobra.NoArgs,
}

func init() {
	statusPagesCmd.AddCommand(spSubscribersCmd)
}

var (
	spSubscribersListFlags = upapi.StatusPageSubscriberListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	spSubscribersListCmd = &cobra.Command{
		Use:     "list <status-page-pk>",
		Aliases: []string{"ls"},
		Short:   "List status page subscribers",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spSubscribersList(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spSubscribersListCmd.Flags(), &spSubscribersListFlags)
	if err != nil {
		panic(err)
	}
	spSubscribersCmd.AddCommand(spSubscribersListCmd)
}

func spSubscribersList(ctx context.Context, pkstr string) ([]upapi.StatusPageSubscriber, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	result, err := api.StatusPages().Subscribers(upapi.PrimaryKey(pk)).List(ctx, spSubscribersListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var spSubscribersGetCmd = &cobra.Command{
	Use:     "get <status-page-pk> <subscriber-pk>",
	Aliases: []string{"show"},
	Short:   "Get a status page subscriber",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spSubscribersGet(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spSubscribersCmd.AddCommand(spSubscribersGetCmd)
}

func spSubscribersGet(ctx context.Context, spPKStr, subPKStr string) (*upapi.StatusPageSubscriber, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	subPK, err := parsePK(subPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Subscribers(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(subPK))
}

var (
	spSubscribersCreateFlags upapi.StatusPageSubscriber
	spSubscribersCreateCmd   = &cobra.Command{
		Use:     "create <status-page-pk>",
		Aliases: []string{"new"},
		Short:   "Create a status page subscriber",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spSubscribersCreate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spSubscribersCreateCmd.Flags(), &spSubscribersCreateFlags)
	if err != nil {
		panic(err)
	}
	spSubscribersCmd.AddCommand(spSubscribersCreateCmd)
}

func spSubscribersCreate(ctx context.Context, pkstr string) (*upapi.StatusPageSubscriber, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Subscribers(upapi.PrimaryKey(pk)).Create(ctx, spSubscribersCreateFlags)
}

var spSubscribersDeleteCmd = &cobra.Command{
	Use:     "delete <status-page-pk> <subscriber-pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Delete a status page subscriber",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spSubscribersDelete(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spSubscribersCmd.AddCommand(spSubscribersDeleteCmd)
}

func spSubscribersDelete(ctx context.Context, spPKStr, subPKStr string) (*upapi.StatusPageSubscriber, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	subPK, err := parsePK(subPKStr)
	if err != nil {
		return nil, err
	}
	obj, err := api.StatusPages().Subscribers(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(subPK))
	if err != nil {
		return nil, err
	}
	return obj, api.StatusPages().Subscribers(upapi.PrimaryKey(spPK)).Delete(ctx, upapi.PrimaryKey(subPK))
}
