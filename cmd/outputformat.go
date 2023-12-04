package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type outputFormat string

const (
	plainFormat outputFormat = "plain"
	jsonFormat  outputFormat = "json"
)

func (o *outputFormat) String() string {
	return string(*o)
}

func (o *outputFormat) Set(s string) error {
	f := outputFormat(s)
	switch f {
	case plainFormat, jsonFormat:
		*o = f
		return nil
	default:
		return fmt.Errorf(`must be one of: %s, %s`, plainFormat, jsonFormat)
	}
}

func (o outputFormat) Type() string {
	return "outputFormat"
}

func outputFormatCompletion(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return []string{
		string(plainFormat) + "\t values are printed in a way to make them as human readable as possible",
		string(jsonFormat) + "\t values are printed as a JSON array of JSON objects, where each object represents single key-value pair",
	}, cobra.ShellCompDirectiveDefault
}
