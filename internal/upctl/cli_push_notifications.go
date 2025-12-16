package upctl

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var pushNotificationsCmd = &cobra.Command{
	Use:     "push-notifications",
	Aliases: []string{"push-notification", "pn", "push"},
	Short:   "Manage push notification profiles",
	Args:    cobra.NoArgs,
}

func init() {
	cmd.AddCommand(pushNotificationsCmd)
}

var (
	pushNotificationsListFlags = upapi.PushNotificationProfileListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	pushNotificationsListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List push notification profiles",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(pushNotificationsList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(pushNotificationsListCmd.Flags(), &pushNotificationsListFlags)
	if err != nil {
		log.Fatalln(err)
	}
	pushNotificationsCmd.AddCommand(pushNotificationsListCmd)
}

func pushNotificationsList(ctx context.Context) ([]upapi.PushNotificationProfile, error) {
	result, err := api.PushNotifications().List(ctx, pushNotificationsListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

var (
	pushNotificationsGetCmd = &cobra.Command{
		Use:     "get <pk>",
		Aliases: []string{"show"},
		Short:   "Get a push notification profile",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(pushNotificationsGet(cmd.Context(), args[0]))
		},
	}
)

func init() {
	pushNotificationsCmd.AddCommand(pushNotificationsGetCmd)
}

func pushNotificationsGet(ctx context.Context, pkstr string) (*upapi.PushNotificationProfile, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.PushNotifications().Get(ctx, upapi.PrimaryKey(pk))
}

var (
	pushNotificationsCreateFlags upapi.PushNotificationProfileCreateRequest
	pushNotificationsCreateCmd   = &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Short:   "Create a new push notification profile",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(pushNotificationsCreate(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(pushNotificationsCreateCmd.Flags(), &pushNotificationsCreateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	pushNotificationsCmd.AddCommand(pushNotificationsCreateCmd)
}

func pushNotificationsCreate(ctx context.Context) (*upapi.PushNotificationProfile, error) {
	return api.PushNotifications().Create(ctx, pushNotificationsCreateFlags)
}

var (
	pushNotificationsUpdateFlags upapi.PushNotificationProfileUpdateRequest
	pushNotificationsUpdateCmd   = &cobra.Command{
		Use:     "update <pk>",
		Aliases: []string{"up"},
		Short:   "Update a push notification profile",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(pushNotificationsUpdate(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(pushNotificationsUpdateCmd.Flags(), &pushNotificationsUpdateFlags)
	if err != nil {
		log.Fatalln(err)
	}
	pushNotificationsCmd.AddCommand(pushNotificationsUpdateCmd)
}

func pushNotificationsUpdate(ctx context.Context, pkstr string) (*upapi.PushNotificationProfile, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	return api.PushNotifications().Update(ctx, upapi.PrimaryKey(pk), pushNotificationsUpdateFlags)
}

var (
	pushNotificationsDeleteCmd = &cobra.Command{
		Use:     "delete <pk>",
		Aliases: []string{"del", "rm", "remove"},
		Short:   "Delete a push notification profile",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(pushNotificationsDelete(cmd.Context(), args[0]))
		},
	}
)

func init() {
	pushNotificationsCmd.AddCommand(pushNotificationsDeleteCmd)
}

func pushNotificationsDelete(ctx context.Context, pkstr string) (*upapi.PushNotificationProfile, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	profile, err := api.PushNotifications().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return profile, api.PushNotifications().Delete(ctx, upapi.PrimaryKey(pk))
}
