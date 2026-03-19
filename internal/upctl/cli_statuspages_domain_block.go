package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var spDomainBlockCmd = &cobra.Command{
	Use:     "domain-block",
	Aliases: []string{"block"},
	Short:   "Manage status page subscription domain block list",
	Args:    cobra.NoArgs,
}

func init() {
	statusPagesCmd.AddCommand(spDomainBlockCmd)
}

var (
	spDomainBlockListFlags = upapi.StatusPageSubsDomainBlockListListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	spDomainBlockListCmd = &cobra.Command{
		Use:     "list <status-page-pk>",
		Aliases: []string{"ls"},
		Short:   "List blocked subscription domains",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spDomainBlockList(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spDomainBlockListCmd.Flags(), &spDomainBlockListFlags)
	if err != nil {
		panic(err)
	}
	spDomainBlockCmd.AddCommand(spDomainBlockListCmd)
}

func spDomainBlockList(ctx context.Context, pkstr string) ([]upapi.StatusPageSubsDomainBlockList, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	result, err := api.StatusPages().SubscriptionDomainBlockList(upapi.PrimaryKey(pk)).List(ctx, spDomainBlockListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var spDomainBlockGetCmd = &cobra.Command{
	Use:     "get <status-page-pk> <domain-pk>",
	Aliases: []string{"show"},
	Short:   "Get a blocked subscription domain",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spDomainBlockGet(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spDomainBlockCmd.AddCommand(spDomainBlockGetCmd)
}

func spDomainBlockGet(ctx context.Context, spPKStr, domPKStr string) (*upapi.StatusPageSubsDomainBlockList, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	domPK, err := parsePK(domPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().SubscriptionDomainBlockList(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(domPK))
}

var (
	spDomainBlockCreateFlags upapi.StatusPageSubsDomainBlockList
	spDomainBlockCreateCmd   = &cobra.Command{
		Use:     "create <status-page-pk>",
		Aliases: []string{"new"},
		Short:   "Add a domain to the block list",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spDomainBlockCreate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spDomainBlockCreateCmd.Flags(), &spDomainBlockCreateFlags)
	if err != nil {
		panic(err)
	}
	spDomainBlockCmd.AddCommand(spDomainBlockCreateCmd)
}

func spDomainBlockCreate(ctx context.Context, pkstr string) (*upapi.StatusPageSubsDomainBlockList, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().SubscriptionDomainBlockList(upapi.PrimaryKey(pk)).Create(ctx, spDomainBlockCreateFlags)
}

var (
	spDomainBlockUpdateFlags upapi.StatusPageSubsDomainBlockList
	spDomainBlockUpdateCmd   = &cobra.Command{
		Use:     "update <status-page-pk> <domain-pk>",
		Aliases: []string{"up"},
		Short:   "Update a blocked subscription domain",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spDomainBlockUpdate(cmd.Context(), args[0], args[1]))
		},
	}
)

func init() {
	err := Bind(spDomainBlockUpdateCmd.Flags(), &spDomainBlockUpdateFlags)
	if err != nil {
		panic(err)
	}
	spDomainBlockCmd.AddCommand(spDomainBlockUpdateCmd)
}

func spDomainBlockUpdate(ctx context.Context, spPKStr, domPKStr string) (*upapi.StatusPageSubsDomainBlockList, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	domPK, err := parsePK(domPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().SubscriptionDomainBlockList(upapi.PrimaryKey(spPK)).Update(ctx, upapi.PrimaryKey(domPK), spDomainBlockUpdateFlags)
}

var spDomainBlockDeleteCmd = &cobra.Command{
	Use:     "delete <status-page-pk> <domain-pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Remove a domain from the block list",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spDomainBlockDelete(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spDomainBlockCmd.AddCommand(spDomainBlockDeleteCmd)
}

func spDomainBlockDelete(ctx context.Context, spPKStr, domPKStr string) (*upapi.StatusPageSubsDomainBlockList, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	domPK, err := parsePK(domPKStr)
	if err != nil {
		return nil, err
	}
	obj, err := api.StatusPages().SubscriptionDomainBlockList(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(domPK))
	if err != nil {
		return nil, err
	}
	return obj, api.StatusPages().SubscriptionDomainBlockList(upapi.PrimaryKey(spPK)).Delete(ctx, upapi.PrimaryKey(domPK))
}
