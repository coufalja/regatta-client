package cmd

import (
	"context"
	"encoding/base64"

	"github.com/spf13/cobra"
)

var putBinary bool

func init() {
	PutCmd.Flags().BoolVar(&putBinary, "binary", false, "provided <value> is binary data encoded using Base64")
}

// PutCmd is a subcommand used for creating/updating records in a table.
var PutCmd = cobra.Command{
	Use:     "put <table> <key> <value>",
	Short:   "Put data into Regatta store",
	Long:    "Put data into Regatta store using Put query as defined in API (https://engineering.jamf.com/regatta/api/#put).",
	Example: "regatta-client put table key value",
	Args:    cobra.ExactArgs(3),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, "You must specify a Regatta table name")
		} else if len(args) == 1 {
			comps = cobra.AppendActiveHelp(comps, "You must specify a key for the value to be inserted")
		} else if len(args) == 2 {
			comps = cobra.AppendActiveHelp(comps, "You must specify a value to be inserted")
		}
		return comps, cobra.ShellCompDirectiveNoFileComp
	},
	PreRunE: connect,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
		defer cancel()
		key, value, err := keyAndValueForPut(args)
		if err != nil {
			cmd.PrintErrln("There was an error while decoding parameters.", err)
			return
		}
		_, err = regatta.Table(args[0]).Put(ctx, key, value)
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
