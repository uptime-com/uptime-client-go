package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var spUsersCmd = &cobra.Command{
	Use:     "users",
	Aliases: []string{"user"},
	Short:   "Manage status page users",
	Args:    cobra.NoArgs,
}

func init() {
	statusPagesCmd.AddCommand(spUsersCmd)
}

var (
	spUsersListFlags = upapi.StatusPageUserListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	spUsersListCmd = &cobra.Command{
		Use:     "list <status-page-pk>",
		Aliases: []string{"ls"},
		Short:   "List status page users",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spUsersList(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spUsersListCmd.Flags(), &spUsersListFlags)
	if err != nil {
		panic(err)
	}
	spUsersCmd.AddCommand(spUsersListCmd)
}

func spUsersList(ctx context.Context, pkstr string) ([]upapi.StatusPageUser, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	result, err := api.StatusPages().Users(upapi.PrimaryKey(pk)).List(ctx, spUsersListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var spUsersGetCmd = &cobra.Command{
	Use:     "get <status-page-pk> <user-pk>",
	Aliases: []string{"show"},
	Short:   "Get a status page user",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spUsersGet(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spUsersCmd.AddCommand(spUsersGetCmd)
}

func spUsersGet(ctx context.Context, spPKStr, userPKStr string) (*upapi.StatusPageUser, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	userPK, err := parsePK(userPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Users(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(userPK))
}

var (
	spUsersCreateFlags upapi.StatusPageUser
	spUsersCreateCmd   = &cobra.Command{
		Use:     "create <status-page-pk>",
		Aliases: []string{"new"},
		Short:   "Create a status page user",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spUsersCreate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(spUsersCreateCmd.Flags(), &spUsersCreateFlags)
	if err != nil {
		panic(err)
	}
	spUsersCmd.AddCommand(spUsersCreateCmd)
}

func spUsersCreate(ctx context.Context, pkstr string) (*upapi.StatusPageUser, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Users(upapi.PrimaryKey(pk)).Create(ctx, spUsersCreateFlags)
}

var (
	spUsersUpdateFlags upapi.StatusPageUser
	spUsersUpdateCmd   = &cobra.Command{
		Use:     "update <status-page-pk> <user-pk>",
		Aliases: []string{"up"},
		Short:   "Update a status page user",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(spUsersUpdate(cmd.Context(), args[0], args[1]))
		},
	}
)

func init() {
	err := Bind(spUsersUpdateCmd.Flags(), &spUsersUpdateFlags)
	if err != nil {
		panic(err)
	}
	spUsersCmd.AddCommand(spUsersUpdateCmd)
}

func spUsersUpdate(ctx context.Context, spPKStr, userPKStr string) (*upapi.StatusPageUser, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	userPK, err := parsePK(userPKStr)
	if err != nil {
		return nil, err
	}
	return api.StatusPages().Users(upapi.PrimaryKey(spPK)).Update(ctx, upapi.PrimaryKey(userPK), spUsersUpdateFlags)
}

var spUsersDeleteCmd = &cobra.Command{
	Use:     "delete <status-page-pk> <user-pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Delete a status page user",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(spUsersDelete(cmd.Context(), args[0], args[1]))
	},
}

func init() {
	spUsersCmd.AddCommand(spUsersDeleteCmd)
}

func spUsersDelete(ctx context.Context, spPKStr, userPKStr string) (*upapi.StatusPageUser, error) {
	spPK, err := parsePK(spPKStr)
	if err != nil {
		return nil, err
	}
	userPK, err := parsePK(userPKStr)
	if err != nil {
		return nil, err
	}
	obj, err := api.StatusPages().Users(upapi.PrimaryKey(spPK)).Get(ctx, upapi.PrimaryKey(userPK))
	if err != nil {
		return nil, err
	}
	return obj, api.StatusPages().Users(upapi.PrimaryKey(spPK)).Delete(ctx, upapi.PrimaryKey(userPK))
}
