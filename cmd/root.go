package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use: "regatta",
}

var endpointOption string
var insecureOption bool

func init() {
	RootCmd.PersistentFlags().StringVar(&endpointOption, "endpoint", "http://localhost:8443", "Regatta API endpoint")
	RootCmd.PersistentFlags().BoolVar(&insecureOption, "insecure", false, "Allow insecure connection")
}

func Execute() {
	RootCmd.Execute()
}
