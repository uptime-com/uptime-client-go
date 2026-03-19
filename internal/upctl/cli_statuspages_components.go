package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var spComponentsCmd = &cobra.Command{
	Use:     "components <status-page-pk>",
	Aliases: []string{"component", "comp"},
	Short:   "Manage status page components",
	Args:    cobra.NoArgs,
}

func init() {
	statusPagesCmd.AddCommand(spComponentsCmd)
}

var (
	spComponentsListFlags = upapi.StatusPageComponentListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	spComponentsListCmd = &cobra.Command{
		Use:     "list <status-page-pk>",
		Aliases: []string{"ls"},
		Short:   "List status page components",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spComponentsList(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spComponentsListCmd.Flags(), &spComponentsListFlags)
	if err != nil {
		panic(err)
	}
	spComponentsCmd.AddCommand(spComponentsListCmd)
}

func spComponentsList(ctx context.Context, pkstr string) ([]upapi.StatusPageComponent, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	result, err := api.StatusPages().Components(upapi.PrimaryKey(pk)).List(ctx, spComponentsListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var spComponentsGetCmd = &cobra.Command{
	Use:     "get <status-page-pk> <component-pk>",
	Aliases: []string{"show"},
	Short:   "Get a status page component",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spComponentsGet(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spComponentsCmd.AddCommand(spComponentsGetCmd)
}

func spComponentsGet(ctx context.Context, spPKStr, compPKStr string) (*upapi.StatusPageComponent, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	compPK, err := parsePK(compPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Components(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(compPK))
}

var (
	spComponentsCreateFlags upapi.StatusPageComponent
	spComponentsCreateCmd   = &cobra.Command{
		Use:     "create <status-page-pk>",
		Aliases: []string{"new"},
		Short:   "Create a status page component",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spComponentsCreate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spComponentsCreateCmd.Flags(), &spComponentsCreateFlags)
	if err != nil {
		panic(err)
	}
	spComponentsCmd.AddCommand(spComponentsCreateCmd)
}

func spComponentsCreate(ctx context.Context, pkstr string) (*upapi.StatusPageComponent, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Components(upapi.PrimaryKey(pk)).Create(ctx, spComponentsCreateFlags)
}

var (
	spComponentsUpdateFlags upapi.StatusPageComponent
	spComponentsUpdateCmd   = &cobra.Command{
		Use:     "update <status-page-pk> <component-pk>",
		Aliases: []string{"up"},
		Short:   "Update a status page component",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spComponentsUpdate(cmd.Context(), args[0], args[1]))
		},
	}
)

func init() {
	err := Bind(spComponentsUpdateCmd.Flags(), &spComponentsUpdateFlags)
	if err != nil {
		panic(err)
	}
	spComponentsCmd.AddCommand(spComponentsUpdateCmd)
}

func spComponentsUpdate(ctx context.Context, spPKStr, compPKStr string) (*upapi.StatusPageComponent, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	compPK, err := parsePK(compPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Components(upapi.PrimaryKey(spPK)).Update(ctx, upapi.PrimaryKey(compPK), spComponentsUpdateFlags)
}

var spComponentsDeleteCmd = &cobra.Command{
	Use:     "delete <status-page-pk> <component-pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Delete a status page component",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spComponentsDelete(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spComponentsCmd.AddCommand(spComponentsDeleteCmd)
}

func spComponentsDelete(ctx context.Context, spPKStr, compPKStr string) (*upapi.StatusPageComponent, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	compPK, err := parsePK(compPKStr)
	if err != nil {
		return nil, err
	}
	obj, err := api.StatusPages().Components(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(compPK))
	if err != nil {
		return nil, err
	}
	return obj, api.StatusPages().Components(upapi.PrimaryKey(spPK)).Delete(ctx, upapi.PrimaryKey(compPK))
}
