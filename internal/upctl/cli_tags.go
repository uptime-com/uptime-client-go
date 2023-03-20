package upctl

import (
	"context"
	"errors"
	"log"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var tagsCmd = &cobra.Command{
	Use:     "tags",
	Aliases: []string{"tag", "t"},
	Short:   "Manage tags",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(tagsCmd)
}

var (
	tagsListFlags = upapi.TagListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	tagsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List tags",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(tagsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(tagsListCmd.Flags(), &tagsListFlags)
	if err != nil {
		log.Fatalln(err)
	}
	tagsCmd.AddCommand(tagsListCmd)
}

func tagsList(ctx context.Context) ([]upapi.Tag, error) {
	return api.Tags().List(ctx, upapi.TagListOptions{
		PageSize: 100,
		Page:     tagsListFlags.Page,
	})
}

var (
	tagsCreateFlags = upapi.Tag{
		ColorHex: "#000000",
	}
	tagsCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Short:   "Create a new tag",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(tagsCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(tagsCreateCmd.Flags(), &tagsCreateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	tagsCmd.AddCommand(tagsCreateCmd)
}

func tagsCreate(ctx context.Context) (*upapi.Tag, error) {
	return api.Tags().Create(ctx, tagsCreateFlags)
}

var (
	tagsUpdateFlags = struct {
		PK  int
		Tag upapi.Tag
	}{}
	tagsUpdateCmd = &cobra.Command{
		Use:     "update",
		Aliases: []string{"set"},
		Short:   "Update a tag",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(tagsUpdate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(tagsUpdateCmd.Flags(), &tagsUpdateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	tagsCmd.AddCommand(tagsUpdateCmd)
}

func tagsUpdate(ctx context.Context) (*upapi.Tag, error) {
	if tagsUpdateFlags.PK == 0 {
		return nil, errors.New("please provide tag primary key")
	}
	tagsUpdateFlags.Tag.PK = tagsUpdateFlags.PK
	return api.Tags().Update(ctx, tagsUpdateFlags.Tag)
}

var (
	tagsDeleteCmd = &cobra.Command{
		Use:     "delete <pk>",
		Aliases: []string{"del", "rm", "remove"},
		Short:   "Delete a tag",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(tagsDelete(cmd.Context(), args[0]))
		},
	}
)

func init() {
	tagsCmd.AddCommand(tagsDeleteCmd)
}

func tagsDelete(ctx context.Context, pkstr string) (*upapi.Tag, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	tag, err := api.Tags().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return tag, api.Tags().Delete(ctx, upapi.PrimaryKey(pk))
}
