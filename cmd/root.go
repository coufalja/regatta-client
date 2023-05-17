package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use: "regatta",
}

var endpointOption string
var insecureOption bool
var binaryData bool

func init() {
	RootCmd.PersistentFlags().StringVar(&endpointOption, "endpoint", "localhost:8443", "Regatta API endpoint")
	RootCmd.PersistentFlags().BoolVar(&insecureOption, "insecure", false, "Allow insecure connection")
	RootCmd.PersistentFlags().BoolVar(&binaryData, "binary", false, "This option will avoid decoding keys and values into UTF-8 strings, but rather encode them as BASE64 strings")
}

func Execute() {
	RootCmd.Execute()
}
