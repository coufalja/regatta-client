package cmd

import (
	"context"
	"strings"

	"github.com/fatih/color"
	client "github.com/jamf/regatta-go"
	"github.com/spf13/cobra"
)

// DeleteCmd is a subcommand used for deleting records in a table.
var DeleteCmd = cobra.Command{
	Use:   "delete <table> <key>",
	Short: "Delete data from Regatta store",
	Long: "Deletes data from Regatta store using DeleteRange query as defined in API (https://engineering.jamf.com/regatta/api/#deleterange).\n" +
		"You can delete single item in Regatta by providing item's key.\n" +
		"Or you can delete items with given prefix, by providing the given prefix and adding the asterisk (*) to the prefix.\n" +
		"Or you can delete all items in the table by specifying only asterisk (*).\n" +
		"When key or prefix is provided, it needs to be valid UTF-8 string.",
	Example: "regatta-client delete table key\n" +
		"regatta-client delete table 'prefix*'\n" +
		"regatta-client delete table '*'",
	Args: cobra.ExactArgs(2),
	ValidArgsFunction: func(_ *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, "You must specify a Regatta table name")
		} else if len(args) == 1 {
			comps = cobra.AppendActiveHelp(comps, "You must specify a key or prefix to delete")
		}
		return comps, cobra.ShellCompDirectiveNoFileComp
	},
	PreRunE: connect,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
		defer cancel()
		key, opts := keyAndOptsForDelete(args)
		resp, err := regatta.Table(args[0]).Delete(ctx, key, opts...)
		if err != nil {
			handleRegattaError(cmd, err)
			return
		}
		count := color.New(color.FgBlue).Sprint(resp.Deleted)
		cmd.Println(count)
	},
}

func keyAndOptsForDelete(args []string) (string, []client.OpOption) {
	key := args[1]
	if strings.HasSuffix(key, "*") {
		key = strings.TrimSuffix(key, "*")
		if len(key) == 0 {
			// delete all
			return zero, []client.OpOption{client.WithRange(zero), client.WithCount()}
		}
		// delete by prefix
		return key, []client.OpOption{client.WithPrefix(), client.WithCount()}
	}
	// delete single
	return key, []client.OpOption{client.WithCount()}
}
