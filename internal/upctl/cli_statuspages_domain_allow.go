package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var spDomainAllowCmd = &cobra.Command{
	Use:     "domain-allow",
	Aliases: []string{"allow"},
	Short:   "Manage status page subscription domain allow list",
	Args:    cobra.NoArgs,
}

func init() {
	statusPagesCmd.AddCommand(spDomainAllowCmd)
}

var (
	spDomainAllowListFlags = upapi.StatusPageSubsDomainAllowListListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	spDomainAllowListCmd = &cobra.Command{
		Use:     "list <status-page-pk>",
		Aliases: []string{"ls"},
		Short:   "List allowed subscription domains",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spDomainAllowList(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spDomainAllowListCmd.Flags(), &spDomainAllowListFlags)
	if err != nil {
		panic(err)
	}
	spDomainAllowCmd.AddCommand(spDomainAllowListCmd)
}

func spDomainAllowList(ctx context.Context, pkstr string) ([]upapi.StatusPageSubsDomainAllowList, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	result, err := api.StatusPages().SubscriptionDomainAllowList(upapi.PrimaryKey(pk)).List(ctx, spDomainAllowListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var spDomainAllowGetCmd = &cobra.Command{
	Use:     "get <status-page-pk> <domain-pk>",
	Aliases: []string{"show"},
	Short:   "Get an allowed subscription domain",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spDomainAllowGet(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spDomainAllowCmd.AddCommand(spDomainAllowGetCmd)
}

func spDomainAllowGet(ctx context.Context, spPKStr, domPKStr string) (*upapi.StatusPageSubsDomainAllowList, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	domPK, err := parsePK(domPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().SubscriptionDomainAllowList(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(domPK))
}

var (
	spDomainAllowCreateFlags upapi.StatusPageSubsDomainAllowList
	spDomainAllowCreateCmd   = &cobra.Command{
		Use:     "create <status-page-pk>",
		Aliases: []string{"new"},
		Short:   "Add a domain to the allow list",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spDomainAllowCreate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spDomainAllowCreateCmd.Flags(), &spDomainAllowCreateFlags)
	if err != nil {
		panic(err)
	}
	spDomainAllowCmd.AddCommand(spDomainAllowCreateCmd)
}

func spDomainAllowCreate(ctx context.Context, pkstr string) (*upapi.StatusPageSubsDomainAllowList, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().SubscriptionDomainAllowList(upapi.PrimaryKey(pk)).Create(ctx, spDomainAllowCreateFlags)
}

var (
	spDomainAllowUpdateFlags upapi.StatusPageSubsDomainAllowList
	spDomainAllowUpdateCmd   = &cobra.Command{
		Use:     "update <status-page-pk> <domain-pk>",
		Aliases: []string{"up"},
		Short:   "Update an allowed subscription domain",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spDomainAllowUpdate(cmd.Context(), args[0], args[1]))
		},
	}
)

func init() {
	err := Bind(spDomainAllowUpdateCmd.Flags(), &spDomainAllowUpdateFlags)
	if err != nil {
		panic(err)
	}
	spDomainAllowCmd.AddCommand(spDomainAllowUpdateCmd)
}

func spDomainAllowUpdate(ctx context.Context, spPKStr, domPKStr string) (*upapi.StatusPageSubsDomainAllowList, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	domPK, err := parsePK(domPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().SubscriptionDomainAllowList(upapi.PrimaryKey(spPK)).Update(ctx, upapi.PrimaryKey(domPK), spDomainAllowUpdateFlags)
}

var spDomainAllowDeleteCmd = &cobra.Command{
	Use:     "delete <status-page-pk> <domain-pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Remove a domain from the allow list",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spDomainAllowDelete(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spDomainAllowCmd.AddCommand(spDomainAllowDeleteCmd)
}

func spDomainAllowDelete(ctx context.Context, spPKStr, domPKStr string) (*upapi.StatusPageSubsDomainAllowList, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	domPK, err := parsePK(domPKStr)
	if err != nil {
		return nil, err
	}
	obj, err := api.StatusPages().SubscriptionDomainAllowList(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(domPK))
	if err != nil {
		return nil, err
	}
	return obj, api.StatusPages().SubscriptionDomainAllowList(upapi.PrimaryKey(spPK)).Delete(ctx, upapi.PrimaryKey(domPK))
}
