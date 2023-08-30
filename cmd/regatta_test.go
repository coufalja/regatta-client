package cmd

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_handleRegattaError(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantMsg string
	}{
		{
			name:    "generic error",
			err:     errors.New("some error"),
			wantMsg: `Received RPC error from Regatta, code 'Unknown' with message 'some error'`,
		},
		{
			name:    "internal Regatta error",
			err:     status.Error(codes.Internal, "internal Regatta error"),
			wantMsg: `Received RPC error from Regatta, code 'Internal' with message 'internal Regatta error'`,
		},
		{
			name:    "not found Regatta error",
			err:     status.Error(codes.NotFound, "resource not found"),
			wantMsg: `The requested resource was not found: resource not found`,
		},
		{
			name:    "unavailable Regatta error",
			err:     status.Error(codes.Unavailable, "resource unavailable"),
			wantMsg: `Regatta is not reachable: resource unavailable`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			RootCmd.SetErr(buf)

			handleRegattaError(&RootCmd, tt.err)

			assert.Equal(t, tt.wantMsg, strings.TrimSpace(buf.String()))
		})
	}
}
