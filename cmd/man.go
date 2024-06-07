package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// ManCmd is a subcommand used for printing man pages.
var ManCmd = cobra.Command{
	Use:     "man",
	Short:   "Generates man pages",
	Example: "regatta-client man .",
	Args:    cobra.MatchAll(cobra.ExactArgs(1)),
	RunE: func(_ *cobra.Command, args []string) error {
		return doc.GenManTree(&RootCmd, nil, args[0])
	},
}
