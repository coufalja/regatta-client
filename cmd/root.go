package cmd

import (
	"context"
	"os"
	"runtime/debug"
	"time"

	"github.com/fatih/color"
	client "github.com/jamf/regatta-go"
	"github.com/spf13/cobra"
)

// Version is set during release of project.
var Version string

var regatta tableClient

type tableClient interface {
	Table(string) client.Table
	Status(ctx context.Context, endpoint string) (*client.StatusResponse, error)
}

// RootCmd is a root command for all the subcommands of regatta-client.
var RootCmd = cobra.Command{
	Use:   "regatta-client",
	Short: "Client for Regatta store",
	Long: "Command-line tool wrapping API calls to Regatta (https://engineering.jamf.com/regatta/).\n" +
		"Simplifies querying for data in Regatta store and other operations.",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		color.NoColor = noColor
		return nil
	},
	Version:      Version,
	SilenceUsage: true,
}

var (
	endpoint    string
	insecure    bool
	cert        string
	timeout     time.Duration
	dialTimeout time.Duration
	noColor     bool
)

func init() {
	RootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable color output")
	RootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", "localhost:8443", "Regatta API endpoint")
	RootCmd.PersistentFlags().BoolVar(&insecure, "insecure", false, "allow insecure connection, controls whether certificates are validated")
	RootCmd.PersistentFlags().StringVar(&cert, "cert", "", "Regatta CA cert")
	RootCmd.PersistentFlags().DurationVar(&timeout, "timeout", 10*time.Second, "timeout for the Regatta operation")
	RootCmd.PersistentFlags().DurationVar(&dialTimeout, "dial-timeout", 2*time.Second, "timeout for establishing the connection to the Regatta")

	RootCmd.AddCommand(&RangeCmd)
	RootCmd.AddCommand(&DeleteCmd)
	RootCmd.AddCommand(&PutCmd)
	RootCmd.AddCommand(&ManCmd)
	RootCmd.AddCommand(&VersionCmd)
	RootCmd.AddCommand(&TableCmd)

	RootCmd.SetOut(os.Stdout)
	RootCmd.SetErr(&coloredErrWriter{os.Stderr})

	info, ok := debug.ReadBuildInfo()
	if ok && len(Version) == 0 {
		RootCmd.Version = info.Main.Version
	}
}

// Execute executes root command of regatta-client.
func Execute() {
	_ = RootCmd.Execute()
}
