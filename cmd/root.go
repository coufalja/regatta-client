package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use:   "regatta-client",
	Short: "Client for Regatta store",
	Long: "Command-line tool wrapping API calls to Regatta (https://engineering.jamf.com/regatta/).\n" +
		"Simplifies querying for data in Regatta store and other operations.",
}

var endpointOption string
var insecureOption bool
var binaryData bool

func init() {
	RootCmd.PersistentFlags().StringVar(&endpointOption, "endpoint", "localhost:8443", "regatta API endpoint")
	RootCmd.PersistentFlags().BoolVar(&insecureOption, "insecure", false, "allow insecure connection")
	RootCmd.PersistentFlags().BoolVar(&binaryData, "binary", false, "avoid decoding keys and values into UTF-8 strings, but rather encode them as Base64 strings")

	RootCmd.AddCommand(&Range)
	RootCmd.AddCommand(&Delete)
}

func Execute() {
	_ = RootCmd.Execute()
}
