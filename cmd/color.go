package cmd

import (
	"io"

	"github.com/fatih/color"
)

type coloredErrWriter struct {
	io.Writer
}

func (c *coloredErrWriter) Write(b []byte) (n int, err error) {
	return color.New(color.FgRed).Fprint(c.Writer, string(b))
}
