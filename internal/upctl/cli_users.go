package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var usersCmd = &cobra.Command{
	Use:     "users",
	Aliases: []string{"user", "u"},
	Short:   "Manage users",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(usersCmd)
}

var (
	usersListFlags = upapi.UserListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	usersListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List users",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(usersList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(usersListCmd.Flags(), &usersListFlags)
	if err != nil {
		panic(err)
	}
	usersCmd.AddCommand(usersListCmd)
}

func usersList(ctx context.Context) ([]upapi.User, error) {
	result, err := api.Users().List(ctx, usersListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var usersGetCmd = &cobra.Command{
	Use:     "get <pk>",
	Aliases: []string{"show"},
	Short:   "Get a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(usersGet(cmd.Context(), args[0]))
	},
}

func init() {
	usersCmd.AddCommand(usersGetCmd)
}

func usersGet(ctx context.Context, pkstr string) (*upapi.User, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Users().Get(ctx, upapi.PrimaryKey(pk))
}

var (
	usersCreateFlags upapi.UserCreateRequest
	usersCreateCmd   = &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Short:   "Create a new user",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(usersCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(usersCreateCmd.Flags(), &usersCreateFlags)
	if err != nil {
		panic(err)
	}
	usersCmd.AddCommand(usersCreateCmd)
}

func usersCreate(ctx context.Context) (*upapi.User, error) {
	return api.Users().Create(ctx, usersCreateFlags)
}

var (
	usersUpdateFlags upapi.UserUpdateRequest
	usersUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up"},
		Short:   "Update a user",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(usersUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(usersUpdateCmd.Flags(), &usersUpdateFlags)
	if err != nil {
		panic(err)
	}
	usersCmd.AddCommand(usersUpdateCmd)
}

func usersUpdate(ctx context.Context, pkstr string) (*upapi.User, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Users().Update(ctx, upapi.PrimaryKey(pk), usersUpdateFlags)
}

var usersDeleteCmd = &cobra.Command{
	Use:     "delete <pk>",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Delete a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(usersDelete(cmd.Context(), args[0]))
	},
}

func init() {
	usersCmd.AddCommand(usersDeleteCmd)
}

func usersDelete(ctx context.Context, pkstr string) (*upapi.User, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	obj, err := api.Users().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return obj, api.Users().Delete(ctx, upapi.PrimaryKey(pk))
}

var usersDeactivateCmd = &cobra.Command{
	Use:   "deactivate <pk>",
	Short: "Deactivate a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(usersDeactivate(cmd.Context(), args[0]))
	},
}

func init() {
	usersCmd.AddCommand(usersDeactivateCmd)
}

func usersDeactivate(ctx context.Context, pkstr string) (*upapi.User, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Users().Deactivate(ctx, upapi.PrimaryKey(pk))
}

var usersReactivateCmd = &cobra.Command{
	Use:   "reactivate <pk>",
	Short: "Reactivate a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(usersReactivate(cmd.Context(), args[0]))
	},
}

func init() {
	usersCmd.AddCommand(usersReactivateCmd)
}

func usersReactivate(ctx context.Context, pkstr string) (*upapi.User, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Users().Reactivate(ctx, upapi.PrimaryKey(pk))
}
