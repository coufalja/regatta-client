package cmd

import (
	"github.com/spf13/cobra"
)

var Version = "unknown"

var RootCmd = cobra.Command{
	Use:   "regatta-client",
	Short: "Client for Regatta store",
	Long: "Command-line tool wrapping API calls to Regatta (https://engineering.jamf.com/regatta/).\n" +
		"Simplifies querying for data in Regatta store and other operations.",
	Version: Version,
}

var endpointOption string
var insecureOption bool

func init() {
	RootCmd.PersistentFlags().StringVar(&endpointOption, "endpoint", "localhost:8443", "regatta API endpoint")
	RootCmd.PersistentFlags().BoolVar(&insecureOption, "insecure", false, "allow insecure connection")

	RootCmd.AddCommand(&Range)
	RootCmd.AddCommand(&Delete)
	RootCmd.AddCommand(&Put)
}

func Execute() {
	_ = RootCmd.Execute()
}
