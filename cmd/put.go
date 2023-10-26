package cmd

import (
	"encoding/base64"

	"github.com/spf13/cobra"
)

var putBinary bool

func init() {
	Put.Flags().BoolVar(&putBinary, "binary", false, "provided <value> is binary data encoded using Base64")
}

// Put is a subcommand used for creating/updating records in a table.
var Put = cobra.Command{
	Use:     "put <table> <key> <value>",
	Short:   "Put data into Regatta store",
	Long:    "Put data into Regatta store using Put query as defined in API (https://engineering.jamf.com/regatta/api/#put).",
	Example: "regatta-client put table key value",
	Args:    cobra.MatchAll(cobra.ExactArgs(3)),
	PreRunE: connect,
	Run: func(cmd *cobra.Command, args []string) {
		key, value, err := keyAndValueForPut(args)
		if err != nil {
			cmd.PrintErrln("There was an error while decoding parameters.", err)
			return
		}
		_, err = regatta.Table(args[0]).Put(cmd.Context(), key, value)
		if err != nil {
			handleRegattaError(cmd, err)
		}
	},
}

func keyAndValueForPut(args []string) (string, string, error) {
	if putBinary {
		value, err := base64.StdEncoding.DecodeString(args[2])
		if err != nil {
			return "", "", err
		}
		return args[1], string(value), nil
	}
	return args[1], args[2], nil
}
