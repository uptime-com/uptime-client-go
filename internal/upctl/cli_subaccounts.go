package upctl

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var subaccountsCmd = &cobra.Command{
	Use:     "subaccounts",
	Aliases: []string{"subaccount", "sub"},
	Short:   "Manage subaccounts",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(subaccountsCmd)
}

var subaccountsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all subaccounts",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(subaccountsList(cmd.Context()))
	},
}

func init() {
	subaccountsCmd.AddCommand(subaccountsListCmd)
}

func subaccountsList(ctx context.Context) ([]upapi.Subaccount, error) {
	return api.Subaccounts().List(ctx)
}

var subaccountsGetCmd = &cobra.Command{
	Use:     "get <pk>",
	Aliases: []string{"show"},
	Short:   "Get a subaccount by primary key",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return output(subaccountsGet(cmd.Context(), args[0]))
	},
}

func init() {
	subaccountsCmd.AddCommand(subaccountsGetCmd)
}

func subaccountsGet(ctx context.Context, pkstr string) (*upapi.Subaccount, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Subaccounts().Get(ctx, upapi.PrimaryKey(pk))
}

var (
	subaccountsCreateFlags upapi.SubaccountCreateRequest
	subaccountsCreateCmd   = &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Short:   "Create a new subaccount (deducts one pack from main account)",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(subaccountsCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(subaccountsCreateCmd.Flags(), &subaccountsCreateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	subaccountsCmd.AddCommand(subaccountsCreateCmd)
}

func subaccountsCreate(ctx context.Context) (*upapi.Subaccount, error) {
	return api.Subaccounts().Create(ctx, subaccountsCreateFlags)
}

var (
	subaccountsUpdateFlags upapi.SubaccountUpdateRequest
	subaccountsUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up", "rename"},
		Short:   "Update a subaccount (rename)",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(subaccountsUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(subaccountsUpdateCmd.Flags(), &subaccountsUpdateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	subaccountsCmd.AddCommand(subaccountsUpdateCmd)
}

func subaccountsUpdate(ctx context.Context, pkstr string) (*upapi.Subaccount, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.Subaccounts().Update(ctx, upapi.PrimaryKey(pk), subaccountsUpdateFlags)
}

var (
	subaccountsTransferPacksFlags upapi.SubaccountPacks
	subaccountsTransferPacksCmd   = &cobra.Command{
		Use:   "transfer-packs <pk>",
		Short: "Transfer packs between main account and subaccount (positive: main->sub, negative: sub->main)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return subaccountsTransferPacks(cmd.Context(), args[0])
		},
	}
)

func init() {
	err := Bind(subaccountsTransferPacksCmd.Flags(), &subaccountsTransferPacksFlags)
	if err != nil {
		log.Fatalln(err)
	}
	subaccountsCmd.AddCommand(subaccountsTransferPacksCmd)
}

func subaccountsTransferPacks(ctx context.Context, pkstr string) error {
	pk, err := parsePK(pkstr)
	if err != nil {
		return err
	}
	return api.Subaccounts().TransferPacks(ctx, upapi.PrimaryKey(pk), subaccountsTransferPacksFlags)
}
