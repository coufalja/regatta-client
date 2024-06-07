package cmd

import (
	"context"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const unknownVersion = "unknown"

// VersionCmd is a subcommand used for printing client and server version.
var VersionCmd = cobra.Command{
	Use:     "version",
	Short:   "Get current version of regatta-client and a Regatta server",
	Long:    "Get current version of regatta-client and a Regatta server using Status API (https://engineering.jamf.com/regatta/api/#status).",
	Example: "regatta-client version",
	Run: func(cmd *cobra.Command, args []string) {
		printClientVersion(cmd, RootCmd.Version)
		err := connect(cmd, args)
		if err != nil {
			printServerVersion(cmd, unknownVersion, color.FgRed)
			cmd.PrintErrln(err)
			return
		}
		ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
		defer cancel()
		resp, err := regatta.Status(ctx, endpoint)
		if err != nil {
			printServerVersion(cmd, unknownVersion, color.FgRed)
			cmd.PrintErrln(err)
			return
		}
		printServerVersion(cmd, resp.Version, color.FgGreen)
	},
}

func printClientVersion(cmd *cobra.Command, version string) {
	col := color.New(color.FgBlue).Sprint
	ver := color.New(color.FgGreen).Sprint(version)

	cmd.Println(col("client version") + ": " + ver)
}

func printServerVersion(cmd *cobra.Command, version string, vercol color.Attribute) {
	col := color.New(color.FgBlue).Sprint
	ver := color.New(vercol).Sprint(version)

	cmd.Println(col("server version") + ": " + ver)
}
