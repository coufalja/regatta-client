package cmd

import (
	"strings"

	client "github.com/jamf/regatta-go"
	"github.com/spf13/cobra"
)

// Delete is a subcommand used for deleting records in a table.
var Delete = cobra.Command{
	Use:   "delete <table> <key>",
	Short: "Delete data from Regatta store",
	Long: "Deletes data from Regatta store using DeleteRange query as defined in API (https://engineering.jamf.com/regatta/api/#deleterange).\n" +
		"You can delete single item in Regatta by providing item's key.\n" +
		"Or you can delete items with given prefix, by providing the given prefix and adding the asterisk (*) to the prefix.\n" +
		"When key or prefix is provided, it needs to be valid UTF-8 string.",
	Example: "regatta-client delete table key\n" +
		"regatta-client delete table 'prefix*'",
	Args:    cobra.MatchAll(cobra.ExactArgs(2)),
	PreRunE: connect,
	Run: func(cmd *cobra.Command, args []string) {
		key, opts := keyAndOptsForDelete(args)
		_, err := regatta.Table(args[0]).Delete(cmd.Context(), key, opts...)
		if err != nil {
			handleRegattaError(cmd, err)
		}
	},
}

func keyAndOptsForDelete(args []string) (string, []client.OpOption) {
	key := args[1]
	if strings.HasSuffix(key, "*") {
		key = strings.TrimSuffix(key, "*")
		if len(key) == 0 {
			// delete all
			return zero, []client.OpOption{client.WithRange(zero), client.WithPrevKV()}
		}
		// delete by prefix
		return key, []client.OpOption{client.WithPrefix(), client.WithPrevKV()}
	}
	// delete single
	return key, []client.OpOption{client.WithPrevKV()}
}
