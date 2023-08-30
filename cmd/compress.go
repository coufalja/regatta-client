package cmd

import (
	"errors"

	"github.com/jamf/regatta/regattaserver/encoding/gzip"
	"github.com/jamf/regatta/regattaserver/encoding/snappy"
	"github.com/spf13/cobra"
)

var (
	noCompressName = "none"
	noCompress     = compressType(noCompressName)
	gzipCompress   = compressType(gzip.Name)
)

type compressType string

func (c *compressType) String() string {
	return string(*c)
}

func (c *compressType) Set(v string) error {
	switch v {
	case gzip.Name, snappy.Name, noCompressName:
		*c = compressType(v)
		return nil
	default:
		return errors.New(`must be one of "gzip", "snappy" or "none"`)
	}
}

func (c *compressType) Type() string {
	return "compressType"
}

func compressTypeCompletion(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"gzip\thelp text for gzip",
		"snappy\thelp text for snappy",
		"none\thelp text for none",
	}, cobra.ShellCompDirectiveDefault
}
