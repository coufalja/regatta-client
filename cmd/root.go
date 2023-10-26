package cmd

import (
	"context"
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
)

func init() {
	RootCmd.PersistentFlags().StringVar(&endpointOption, "endpoint", "localhost:8443", "regatta API endpoint")
	RootCmd.PersistentFlags().BoolVar(&insecureOption, "insecure", false, "allow insecure connection, controls whether certificates are validated")
	RootCmd.PersistentFlags().StringVar(&certOption, "cert", "", "regatta CA cert")

	RootCmd.AddCommand(&Range)
	RootCmd.AddCommand(&Delete)
	RootCmd.AddCommand(&Put)
	RootCmd.AddCommand(&Man)

	RootCmd.SetOut(os.Stdout)
}

// Execute executes root command of regatta-client.
func Execute() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = RootCmd.ExecuteContext(ctx)
}
