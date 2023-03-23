package upctl

import (
	"context"
	"strconv"

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

func checksUpdateSubcommand(name string, flags any, fn func(context.Context, int) (*upapi.Check, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name + " { pk }",
		Short: "Update " + name + " check",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pk, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return output(fn(cmd.Context(), pk))
		},
	}
	err := Bind(cmd.Flags(), flags)
	if err != nil {
		panic(err)
	}
	return cmd
}

var (
	checksUpdateCmd = &cobra.Command{
		Use:     "update",
		Aliases: []string{"up"},
		Short:   "Update a check",
	}

	updateCreateAPIFlags         upapi.CheckAPI
	updateCreateBlacklistFlags   upapi.CheckBlacklist
	updateCreateDNSFlags         upapi.CheckDNS
	updateCreateGroupFlags       upapi.CheckGroup
	updateCreateHeartbeatFlags   upapi.CheckHeartbeat
	updateCreateHTTPFlags        upapi.CheckHTTP
	updateCreateICMPFlags        upapi.CheckICMP
	updateCreateIMAPFlags        upapi.CheckIMAP
	updateCreateMalwareFlags     upapi.CheckMalware
	updateCreateNTPFlags         upapi.CheckNTP
	updateCreatePOPFlags         upapi.CheckPOP
	updateCreateRUMFlags         upapi.CheckRUM
	updateCreateRUM2Flags        upapi.CheckRUM2
	updateCreateSMTPFlags        upapi.CheckSMTP
	updateCreateSSHFlags         upapi.CheckSSH
	updateCreateSSLCertFlags     upapi.CheckSSLCert
	updateCreateTCPFlags         upapi.CheckTCP
	updateCreateTransactionFlags upapi.CheckTransaction
	updateCreateUDPFlags         upapi.CheckUDP
	updateCreateWebhookFlags     upapi.CheckWebhook
	updateCreateWHOISFlags       upapi.CheckWHOIS
)

func init() {
	checksUpdateCmd.AddCommand(
		checksUpdateSubcommand("api", &updateCreateAPIFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateAPIFlags.PK = pk
			return api.Checks().UpdateAPI(ctx, updateCreateAPIFlags)
		}),
		checksUpdateSubcommand("blacklist", &updateCreateBlacklistFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateBlacklistFlags.PK = pk
			return api.Checks().UpdateBlacklist(ctx, updateCreateBlacklistFlags)
		}),
		checksUpdateSubcommand("dns", &updateCreateDNSFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateDNSFlags.PK = pk
			return api.Checks().UpdateDNS(ctx, updateCreateDNSFlags)
		}),
		checksUpdateSubcommand("group", &updateCreateGroupFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateGroupFlags.PK = pk
			return api.Checks().UpdateGroup(ctx, updateCreateGroupFlags)
		}),
		checksUpdateSubcommand("heartbeat", &updateCreateHeartbeatFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateHeartbeatFlags.PK = pk
			return api.Checks().UpdateHeartbeat(ctx, updateCreateHeartbeatFlags)
		}),
		checksUpdateSubcommand("http", &updateCreateHTTPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateHTTPFlags.PK = pk
			return api.Checks().UpdateHTTP(ctx, updateCreateHTTPFlags)
		}),
		checksUpdateSubcommand("icmp", &updateCreateICMPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateICMPFlags.PK = pk
			return api.Checks().UpdateICMP(ctx, updateCreateICMPFlags)
		}),
		checksUpdateSubcommand("imap", &updateCreateIMAPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateIMAPFlags.PK = pk
			return api.Checks().UpdateIMAP(ctx, updateCreateIMAPFlags)
		}),
		checksUpdateSubcommand("malware", &updateCreateMalwareFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateMalwareFlags.PK = pk
			return api.Checks().UpdateMalware(ctx, updateCreateMalwareFlags)
		}),
		checksUpdateSubcommand("ntp", &updateCreateNTPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateNTPFlags.PK = pk
			return api.Checks().UpdateNTP(ctx, updateCreateNTPFlags)
		}),
		checksUpdateSubcommand("pop", &updateCreatePOPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreatePOPFlags.PK = pk
			return api.Checks().UpdatePOP(ctx, updateCreatePOPFlags)
		}),
		checksUpdateSubcommand("rum", &updateCreateRUMFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateRUMFlags.PK = pk
			return api.Checks().UpdateRUM(ctx, updateCreateRUMFlags)
		}),
		checksUpdateSubcommand("rum2", &updateCreateRUM2Flags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateRUM2Flags.PK = pk
			return api.Checks().UpdateRUM2(ctx, updateCreateRUM2Flags)
		}),
		checksUpdateSubcommand("smtp", &updateCreateSMTPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateSMTPFlags.PK = pk
			return api.Checks().UpdateSMTP(ctx, updateCreateSMTPFlags)
		}),
		checksUpdateSubcommand("ssh", &updateCreateSSHFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateSSHFlags.PK = pk
			return api.Checks().UpdateSSH(ctx, updateCreateSSHFlags)
		}),
		checksUpdateSubcommand("sslcert", &updateCreateSSLCertFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateSSLCertFlags.PK = pk
			return api.Checks().UpdateSSLCert(ctx, updateCreateSSLCertFlags)
		}),
		checksUpdateSubcommand("tcp", &updateCreateTCPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateTCPFlags.PK = pk
			return api.Checks().UpdateTCP(ctx, updateCreateTCPFlags)
		}),
		checksUpdateSubcommand("transaction", &updateCreateTransactionFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateTransactionFlags.PK = pk
			return api.Checks().UpdateTransaction(ctx, updateCreateTransactionFlags)
		}),
		checksUpdateSubcommand("udp", &updateCreateUDPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateUDPFlags.PK = pk
			return api.Checks().UpdateUDP(ctx, updateCreateUDPFlags)
		}),
		checksUpdateSubcommand("webhook", &updateCreateWebhookFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateWebhookFlags.PK = pk
			return api.Checks().UpdateWebhook(ctx, updateCreateWebhookFlags)
		}),
		checksUpdateSubcommand("whois", &updateCreateWHOISFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			updateCreateWHOISFlags.PK = pk
			return api.Checks().UpdateWHOIS(ctx, updateCreateWHOISFlags)
		}),
	)
	checksCmd.AddCommand(checksUpdateCmd)
}

var (
	checksDeleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "Delete a check",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(checksDelete(cmd.Context(), args[0]))
		},
	}
)

func init() {
	checksCmd.AddCommand(checksDeleteCmd)
}

func checksDelete(ctx context.Context, pkstr string) (*upapi.Check, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	obj, err := api.Checks().Get(ctx, upapi.PrimaryKey(pk))
	if err != nil {
		return nil, err
	}
	return obj, api.Checks().Delete(ctx, upapi.PrimaryKey(pk))
}
