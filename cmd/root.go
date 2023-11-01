package cmd

import (
	"os"
	"time"

	client "github.com/jamf/regatta-go"
	"github.com/spf13/cobra"
)

// Version is set during release of project.
var Version = "unknown"

var regatta tableClient

type tableClient interface {
	Table(string) client.Table
}

// RootCmd is a root command for all the subcommands of regatta-client.
var RootCmd = cobra.Command{
	Use:   "regatta-client",
	Short: "Client for Regatta store",
	Long: "Command-line tool wrapping API calls to Regatta (https://engineering.jamf.com/regatta/).\n" +
		"Simplifies querying for data in Regatta store and other operations.",
	Version:      Version,
	SilenceUsage: true,
}

var (
	endpointOption string
	insecureOption bool
	certOption     string
	timeout        time.Duration
	dialTimeout    time.Duration
)

func init() {
	// register common flags directly to the subcommands
	for _, c := range []*cobra.Command{&Range, &Delete, &Put} {
		c.Flags().StringVar(&endpointOption, "endpoint", "localhost:8443", "Regatta API endpoint")
		c.Flags().BoolVar(&insecureOption, "insecure", false, "allow insecure connection, controls whether certificates are validated")
		c.Flags().StringVar(&certOption, "cert", "", "Regatta CA cert")
		c.Flags().DurationVar(&timeout, "timeout", 10*time.Second, "timeout for the Regatta operation")
		c.Flags().DurationVar(&dialTimeout, "dial-timeout", 2*time.Second, "timeout for establishing the connection to the Regatta")
	}

	RootCmd.AddCommand(&Range)
	RootCmd.AddCommand(&Delete)
	RootCmd.AddCommand(&Put)
	RootCmd.AddCommand(&Man)

	RootCmd.SetOut(os.Stdout)
}

// Execute executes root command of regatta-client.
func Execute() {
	_ = RootCmd.Execute()
}
