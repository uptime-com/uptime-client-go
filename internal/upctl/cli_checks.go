package upctl

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

var (
	checksCmd = &cobra.Command{
		Use:     "checks",
		Aliases: []string{"check", "c"},
		Short:   "Manage checks",
	}
)

func init() {
	cmd.AddCommand(checksCmd)
}

var (
	checksListFlags = upapi.CheckListOptions{
		Page:     1,
		PageSize: 100,
		Ordering: "pk",
	}
	checksListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List checks",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(checksList(cmd.Context()))
		},
	}
)

func init() {
	err := Bind(checksListCmd.Flags(), &checksListFlags)
	if err != nil {
		panic(err)
	}
	checksCmd.AddCommand(checksListCmd)
}

func checksList(ctx context.Context) ([]upapi.Check, error) {
	return api.Checks().List(ctx, checksListFlags)
}

var (
	checksCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"add"},
		Short:   "Create a new check",
	}
)

func checksCreateSubcommand(name string, flags any, fn func(context.Context) (*upapi.Check, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Create a new " + name + " check",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(fn(cmd.Context()))
		},
	}
	err := Bind(cmd.Flags(), flags)
	if err != nil {
		panic(err)
	}
	return cmd
}

var (
	checksCreateAPIFlags         upapi.CheckAPI
	checksCreateBlacklistFlags   upapi.CheckBlacklist
	checksCreateDNSFlags         upapi.CheckDNS
	checksCreateGroupFlags       upapi.CheckGroup
	checksCreateHeartbeatFlags   upapi.CheckHeartbeat
	checksCreateHTTPFlags        upapi.CheckHTTP
	checksCreateICMPFlags        upapi.CheckICMP
	checksCreateIMAPFlags        upapi.CheckIMAP
	checksCreateMalwareFlags     upapi.CheckMalware
	checksCreateNTPFlags         upapi.CheckNTP
	checksCreatePOPFlags         upapi.CheckPOP
	checksCreateRUMFlags         upapi.CheckRUM
	checksCreateRUM2Flags        upapi.CheckRUM2
	checksCreateSMTPFlags        upapi.CheckSMTP
	checksCreateSSHFlags         upapi.CheckSSH
	checksCreateSSLCertFlags     upapi.CheckSSLCert
	checksCreateTCPFlags         upapi.CheckTCP
	checksCreateTransactionFlags upapi.CheckTransaction
	checksCreateUDPFlags         upapi.CheckUDP
	checksCreateWebhookFlags     upapi.CheckWebhook
	checksCreateWHOISFlags       upapi.CheckWHOIS
)

func init() {
	checksCreateCmd.AddCommand(
		checksCreateSubcommand("api", &checksCreateAPIFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateAPI(ctx, checksCreateAPIFlags)
		}),
		checksCreateSubcommand("blacklist", &checksCreateBlacklistFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateBlacklist(ctx, checksCreateBlacklistFlags)
		}),
		checksCreateSubcommand("dns", &checksCreateDNSFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateDNS(ctx, checksCreateDNSFlags)
		}),
		checksCreateSubcommand("group", &checksCreateGroupFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateGroup(ctx, checksCreateGroupFlags)
		}),
		checksCreateSubcommand("heartbeat", &checksCreateHeartbeatFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateHeartbeat(ctx, checksCreateHeartbeatFlags)
		}),
		checksCreateSubcommand("http", &checksCreateHTTPFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateHTTP(ctx, checksCreateHTTPFlags)
		}),
		checksCreateSubcommand("icmp", &checksCreateICMPFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateICMP(ctx, checksCreateICMPFlags)
		}),
		checksCreateSubcommand("imap", &checksCreateIMAPFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateIMAP(ctx, checksCreateIMAPFlags)
		}),
		checksCreateSubcommand("malware", &checksCreateMalwareFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateMalware(ctx, checksCreateMalwareFlags)
		}),
		checksCreateSubcommand("ntp", &checksCreateNTPFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateNTP(ctx, checksCreateNTPFlags)
		}),
		checksCreateSubcommand("pop", &checksCreatePOPFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreatePOP(ctx, checksCreatePOPFlags)
		}),
		checksCreateSubcommand("rum", &checksCreateRUMFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateRUM(ctx, checksCreateRUMFlags)
		}),
		checksCreateSubcommand("rum2", &checksCreateRUM2Flags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateRUM2(ctx, checksCreateRUM2Flags)
		}),
		checksCreateSubcommand("smtp", &checksCreateSMTPFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateSMTP(ctx, checksCreateSMTPFlags)
		}),
		checksCreateSubcommand("ssh", &checksCreateSSHFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateSSH(ctx, checksCreateSSHFlags)
		}),
		checksCreateSubcommand("sslcert", &checksCreateSSLCertFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateSSLCert(ctx, checksCreateSSLCertFlags)
		}),
		checksCreateSubcommand("tcp", &checksCreateTCPFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateTCP(ctx, checksCreateTCPFlags)
		}),
		checksCreateSubcommand("transaction", &checksCreateTransactionFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateTransaction(ctx, checksCreateTransactionFlags)
		}),
		checksCreateSubcommand("udp", &checksCreateUDPFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateUDP(ctx, checksCreateUDPFlags)
		}),
		checksCreateSubcommand("webhook", &checksCreateWebhookFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateWebhook(ctx, checksCreateWebhookFlags)
		}),
		checksCreateSubcommand("whois", &checksCreateWHOISFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateWHOIS(ctx, checksCreateWHOISFlags)
		}),
	)
	checksCmd.AddCommand(checksCreateCmd)
}

var (
	checksDeleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "Delete a tag",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(checksDelete(cmd.Context(), args[0]))
		},
	}
)

func init() {
	checksCmd.AddCommand(checksDeleteCmd)
}

func checksDelete(ctx context.Context, pkstr string) (*upapi.Tag, error) {
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
