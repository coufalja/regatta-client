package cmd

import (
	"context"
	"sort"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// TableCmd is a subcommand used for printing tables.
var TableCmd = cobra.Command{
	Use:     "table",
	Short:   "Print available tables",
	Long:    "Print available tables using Status API (https://engineering.jamf.com/regatta/api/#status).",
	Example: "regatta-client table",
	PreRunE: connect,
	Run: func(cmd *cobra.Command, _ []string) {
		ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
		defer cancel()

		resp, err := regatta.Status(ctx, endpoint)
		if err != nil {
			handleRegattaError(cmd, err)
			return
		}

		var sortedTables []string
		for tableName := range resp.Tables {
			sortedTables = append(sortedTables, tableName)
		}

		sort.Strings(sortedTables)

		for _, tableName := range sortedTables {
			coloredTableName := color.New(color.FgGreen).Sprint(tableName)
			cmd.Println(coloredTableName)
		}
	},
}
