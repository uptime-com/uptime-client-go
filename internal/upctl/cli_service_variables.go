package upctl

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var serviceVariablesCmd = &cobra.Command{
	Use:     "servicevariables",
	Aliases: []string{"servicevariable", "sv", "svar"},
	Short:   "Manage service variables",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(serviceVariablesCmd)
}

var (
	serviceVariablesListFlags = upapi.ServiceVariableListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	serviceVariablesListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List service variables",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(serviceVariablesList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(serviceVariablesListCmd.Flags(), &serviceVariablesListFlags)
	if err != nil {
		log.Fatalln(err)
	}
	serviceVariablesCmd.AddCommand(serviceVariablesListCmd)
}

func serviceVariablesList(ctx context.Context) ([]upapi.ServiceVariable, error) {
	return api.ServiceVariables().List(ctx, serviceVariablesListFlags)
}

var (
	serviceVariablesGetCmd = &cobra.Command{
		Use:     "get <pk>",
		Aliases: []string{"show"},
		Short:   "Get a service variable",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(serviceVariablesGet(cmd.Context(), args[0]))
		},
	}
)

func init() {
	serviceVariablesCmd.AddCommand(serviceVariablesGetCmd)
}

func serviceVariablesGet(ctx context.Context, pkstr string) (*upapi.ServiceVariable, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.ServiceVariables().Get(ctx, upapi.PrimaryKey(pk))
}

var (
	serviceVariablesCreateFlags upapi.ServiceVariableCreateRequest
	serviceVariablesCreateCmd   = &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Short:   "Create a new service variable",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(serviceVariablesCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(serviceVariablesCreateCmd.Flags(), &serviceVariablesCreateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	serviceVariablesCmd.AddCommand(serviceVariablesCreateCmd)
}

func serviceVariablesCreate(ctx context.Context) (*upapi.ServiceVariable, error) {
	return api.ServiceVariables().Create(ctx, serviceVariablesCreateFlags)
}

var (
	serviceVariablesUpdateFlags upapi.ServiceVariableUpdateRequest
	serviceVariablesUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up"},
		Short:   "Update a service variable",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(serviceVariablesUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(serviceVariablesUpdateCmd.Flags(), &serviceVariablesUpdateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	serviceVariablesCmd.AddCommand(serviceVariablesUpdateCmd)
}

func serviceVariablesUpdate(ctx context.Context, pkstr string) (*upapi.ServiceVariable, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.ServiceVariables().Update(ctx, upapi.PrimaryKey(pk), serviceVariablesUpdateFlags)
}

var (
	serviceVariablesDeleteCmd = &cobra.Command{
		Use:     "delete <pk>",
		Aliases: []string{"del", "rm", "remove"},
		Short:   "Delete a service variable",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(serviceVariablesDelete(cmd.Context(), args[0]))
		},
	}
)

func init() {
	serviceVariablesCmd.AddCommand(serviceVariablesDeleteCmd)
}

func serviceVariablesDelete(ctx context.Context, pkstr string) (*upapi.ServiceVariable, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	sv, err := api.ServiceVariables().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return sv, api.ServiceVariables().Delete(ctx, upapi.PrimaryKey(pk))
}

