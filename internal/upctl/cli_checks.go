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
	result, err := api.Checks().List(ctx, checksListFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
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
	checksCreateRDAPFlags        upapi.CheckRDAP
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
		checksCreateSubcommand("rdap", &checksCreateRDAPFlags, func(ctx context.Context) (*upapi.Check, error) {
			return api.Checks().CreateRDAP(ctx, checksCreateRDAPFlags)
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
	updateCreateRDAPFlags        upapi.CheckRDAP
)

func init() {
	checksUpdateCmd.AddCommand(
		checksUpdateSubcommand("api", &updateCreateAPIFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateAPI(ctx, upapi.PrimaryKey(pk), updateCreateAPIFlags)
		}),
		checksUpdateSubcommand("blacklist", &updateCreateBlacklistFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateBlacklist(ctx, upapi.PrimaryKey(pk), updateCreateBlacklistFlags)
		}),
		checksUpdateSubcommand("dns", &updateCreateDNSFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateDNS(ctx, upapi.PrimaryKey(pk), updateCreateDNSFlags)
		}),
		checksUpdateSubcommand("group", &updateCreateGroupFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateGroup(ctx, upapi.PrimaryKey(pk), updateCreateGroupFlags)
		}),
		checksUpdateSubcommand("heartbeat", &updateCreateHeartbeatFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateHeartbeat(ctx, upapi.PrimaryKey(pk), updateCreateHeartbeatFlags)
		}),
		checksUpdateSubcommand("http", &updateCreateHTTPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateHTTP(ctx, upapi.PrimaryKey(pk), updateCreateHTTPFlags)
		}),
		checksUpdateSubcommand("icmp", &updateCreateICMPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateICMP(ctx, upapi.PrimaryKey(pk), updateCreateICMPFlags)
		}),
		checksUpdateSubcommand("imap", &updateCreateIMAPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateIMAP(ctx, upapi.PrimaryKey(pk), updateCreateIMAPFlags)
		}),
		checksUpdateSubcommand("malware", &updateCreateMalwareFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateMalware(ctx, upapi.PrimaryKey(pk), updateCreateMalwareFlags)
		}),
		checksUpdateSubcommand("ntp", &updateCreateNTPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateNTP(ctx, upapi.PrimaryKey(pk), updateCreateNTPFlags)
		}),
		checksUpdateSubcommand("pop", &updateCreatePOPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdatePOP(ctx, upapi.PrimaryKey(pk), updateCreatePOPFlags)
		}),
		checksUpdateSubcommand("rum", &updateCreateRUMFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateRUM(ctx, upapi.PrimaryKey(pk), updateCreateRUMFlags)
		}),
		checksUpdateSubcommand("rum2", &updateCreateRUM2Flags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateRUM2(ctx, upapi.PrimaryKey(pk), updateCreateRUM2Flags)
		}),
		checksUpdateSubcommand("smtp", &updateCreateSMTPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateSMTP(ctx, upapi.PrimaryKey(pk), updateCreateSMTPFlags)
		}),
		checksUpdateSubcommand("ssh", &updateCreateSSHFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateSSH(ctx, upapi.PrimaryKey(pk), updateCreateSSHFlags)
		}),
		checksUpdateSubcommand("sslcert", &updateCreateSSLCertFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateSSLCert(ctx, upapi.PrimaryKey(pk), updateCreateSSLCertFlags)
		}),
		checksUpdateSubcommand("tcp", &updateCreateTCPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateTCP(ctx, upapi.PrimaryKey(pk), updateCreateTCPFlags)
		}),
		checksUpdateSubcommand("transaction", &updateCreateTransactionFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateTransaction(ctx, upapi.PrimaryKey(pk), updateCreateTransactionFlags)
		}),
		checksUpdateSubcommand("udp", &updateCreateUDPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateUDP(ctx, upapi.PrimaryKey(pk), updateCreateUDPFlags)
		}),
		checksUpdateSubcommand("webhook", &updateCreateWebhookFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateWebhook(ctx, upapi.PrimaryKey(pk), updateCreateWebhookFlags)
		}),
		checksUpdateSubcommand("whois", &updateCreateWHOISFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateWHOIS(ctx, upapi.PrimaryKey(pk), updateCreateWHOISFlags)
		}),
		checksUpdateSubcommand("rdap", &updateCreateRDAPFlags, func(ctx context.Context, pk int) (*upapi.Check, error) {
			return api.Checks().UpdateRDAP(ctx, upapi.PrimaryKey(pk), updateCreateRDAPFlags)
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

var (
	checksStatsFlags = upapi.CheckStatsOptions{}
	checksStatsCmd   = &cobra.Command{
		Use:   "stats",
		Short: "Get check statistics",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return output(checksStats(cmd.Context(), args[0]))
		},
	}
)

func init() {
	err := Bind(checksStatsCmd.Flags(), &checksStatsFlags)
	if err != nil {
		panic(err)
	}
	checksCmd.AddCommand(checksStatsCmd)
}

func checksStats(ctx context.Context, pkstr string) ([]upapi.CheckStats, error) {
	pk, err := parsePK(pkstr)
	if err != nil {
		return nil, err
	}
	result, err := api.Checks().Stats(ctx, upapi.PrimaryKey(pk), checksStatsFlags)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}
