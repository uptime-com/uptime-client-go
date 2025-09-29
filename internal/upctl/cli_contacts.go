package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var (
	contactsCmd = &cobra.Command{
		Use:   "contacts",
		Short: "Manage contacts",
	}
)

func init() {
	cmd.AddCommand(contactsCmd)
}

var (
	contactsListFlags = upapi.ContactListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	contactsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List contacts",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(contactsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(contactsListCmd.Flags(), &contactsListFlags)
	if err != nil {
		panic(err)
	}
	contactsCmd.AddCommand(contactsListCmd)
}

func contactsList(ctx context.Context) ([]upapi.Contact, error) {
	return api.Contacts().List(ctx, contactsListFlags)
}

var (
	contactsCreateFlags = upapi.Contact{}
	contactsCreateCmd   = &cobra.Command{
		Use:     "create <type>",
		Aliases: []string{"add"},
		Short:   "Create a new check",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(contactsCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(contactsCreateCmd.Flags(), &contactsCreateFlags)
	if err != nil {
		panic(err)
	}
	contactsCmd.AddCommand(contactsCreateCmd)
}

func contactsCreate(ctx context.Context) (*upapi.Contact, error) {
	return api.Contacts().Create(ctx, contactsCreateFlags)
}

var (
	contactsUpdateFlags = upapi.Contact{}
	contactsUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up"},
		Short:   "Update existing contact",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(contactsUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(contactsUpdateCmd.Flags(), &contactsUpdateFlags)
	if err != nil {
		panic(err)
	}
	contactsCmd.AddCommand(contactsUpdateCmd)
}

func contactsUpdate(ctx context.Context, arg string) (*upapi.Contact, error) {
	pk, err := parsePK(arg)
	if err != nil {
		return nil, err
	}
	return api.Contacts().Update(ctx, upapi.PrimaryKey(pk), contactsUpdateFlags)
}

var (
	contactsDeleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "Delete a tag",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(contactsDelete(cmd.Context(), args[0]))
		},
	}
)

func init() {
	contactsCmd.AddCommand(contactsDeleteCmd)
}

func contactsDelete(ctx context.Context, pkstr string) (*upapi.Contact, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	obj, err := api.Contacts().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return obj, api.Contacts().Delete(ctx, upapi.PrimaryKey(pk))
}
