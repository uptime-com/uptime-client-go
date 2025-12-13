package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var credentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Manage credentials",
}

func init() {
	cmd.AddCommand(credentialsCmd)
}

var (
	credentialsListFlags = upapi.CredentialListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	credentialsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List credentials",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(credentialsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(credentialsListCmd.Flags(), &credentialsListFlags)
	if err != nil {
		panic(err)
	}
	credentialsCmd.AddCommand(credentialsListCmd)
}

func credentialsList(ctx context.Context) ([]upapi.Credential, error) {
	result, err := api.Credentials().List(ctx, credentialsListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var (
	credentialsCreateFlags = upapi.Credential{}
	credentialsCreateCmd   = &cobra.Command{
		Use:     "create <type>",
		Aliases: []string{"add"},
		Short:   "Create a new check",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(credentialsCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(credentialsCreateCmd.Flags(), &credentialsCreateFlags)
	if err != nil {
		panic(err)
	}
	credentialsCmd.AddCommand(credentialsCreateCmd)
}

func credentialsCreate(ctx context.Context) (*upapi.Credential, error) {
	return api.Credentials().Create(ctx, credentialsCreateFlags)
}

var (
	credentialsUpdateFlags = upapi.Credential{}
	credentialsUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up"},
		Short:   "Update existing credential",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(credentialsUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(credentialsUpdateCmd.Flags(), &credentialsUpdateFlags)
	if err != nil {
		panic(err)
	}
	credentialsCmd.AddCommand(credentialsUpdateCmd)
}

func credentialsUpdate(ctx context.Context, arg string) (*upapi.Credential, error) {
	pk, err := parsePK(arg)
	if err != nil {
		return nil, err
	}
	return api.Credentials().Update(ctx, upapi.PrimaryKey(pk), credentialsUpdateFlags)
}

var credentialsDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm"},
	Short:   "Delete a tag",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(credentialsDelete(cmd.Context(), args[0]))
	},
}

func init() {
	credentialsCmd.AddCommand(credentialsDeleteCmd)
}

func credentialsDelete(ctx context.Context, pkstr string) (*upapi.Credential, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	obj, err := api.Credentials().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return obj, api.Credentials().Delete(ctx, upapi.PrimaryKey(pk))
}
