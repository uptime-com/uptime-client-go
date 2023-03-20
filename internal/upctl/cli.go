package upctl

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

// cmd represents the base command when called without any subcommands
var (
	api upapi.API

	cmdArgs = struct {
		Color  bool   `flag:"color"  usage:"Enable color for json output"`
		Output string `flag:"output" short:"o" usage:"Output format (json|spew)"`
		Token  string `flag:"token"  usage:"Uptime.com API token"`
		Trace  bool   `flag:"trace"  usage:"Trace HTTP requests"`
	}{
		Color:  true,
		Output: "json",
	}

	cmd = &cobra.Command{
		Use:           "upctl",
		Short:         "Uptime.com command line API client",
		Long:          "", // TODO: add long description
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			opts := []upapi.Option{
				upapi.WithToken(viper.GetString("token")),
			}
			if cmdArgs.Trace {
				opts = append(opts, upapi.WithTrace(os.Stderr))
			}
			api, err = upapi.New(opts...)
			return err
		},
	}
)

func init() {
	err := Bind(cmd.PersistentFlags(), &cmdArgs)
	if err != nil {
		panic(err)
	}
	err = cmd.PersistentFlags().MarkHidden("token")
	if err != nil {
		return
	}
	err = viper.BindPFlags(cmd.PersistentFlags())
	if err != nil {
		panic(err)
	}
	err = viper.BindEnv("token", "UPCTL_TOKEN", "UPTIME_TOKEN")
	if err != nil {
		panic(err)
	}
}

func Execute(version string) {
	cmd.Version = version
	err := cmd.Execute()
	if err != nil {
		var uperr = new(upapi.Error)
		if errors.As(err, &uperr) {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), "\nError: %v\n\n", uperr)
		} else {
			_, _ = fmt.Fprintf(cmd.OutOrStderr(), "\nError: %v\n\n", err)
		}
		os.Exit(1)
	}
}
