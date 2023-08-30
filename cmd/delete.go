package cmd

import (
	"context"
	"strings"
	"time"

	"github.com/jamf/regatta/regattapb"
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
	Args: cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := createClient()
		if err != nil {
			cmd.PrintErrln("There was an error, while establishing connection to Regatta.", err)
			return
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		req := createDeleteRangeRequest(args)

		_, err = client.DeleteRange(timeoutCtx, req)
		if err != nil {
			handleRegattaError(cmd, err)
		}
	},
}

func createDeleteRangeRequest(args []string) *regattapb.DeleteRangeRequest {
	table := args[0]
	key := args[1]
	if strings.HasSuffix(key, "*") {
		key = strings.TrimSuffix(key, "*")
		if len(key) == 0 {
			// delete all
			return &regattapb.DeleteRangeRequest{
				Table:    []byte(table),
				Key:      []byte{0},
				RangeEnd: []byte{0},
				PrevKv:   true,
			}
		}
		// delete by prefix
		return &regattapb.DeleteRangeRequest{
			Table:    []byte(table),
			Key:      []byte(key),
			RangeEnd: []byte(findNextString(key)),
			PrevKv:   true,
		}
	}
	// delete single
	return &regattapb.DeleteRangeRequest{
		Table:  []byte(table),
		Key:    []byte(key),
		PrevKv: true,
	}
}
