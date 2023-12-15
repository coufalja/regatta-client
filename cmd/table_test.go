package cmd

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	client "github.com/jamf/regatta-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Table(t *testing.T) {
	tests := []struct {
		name           string
		statusResponse *client.StatusResponse
		statusErr      error
		wantStdOut     string
		wantStdErr     string
	}{
		{
			name:           "print tables",
			statusResponse: &client.StatusResponse{Tables: map[string]*client.TableStatus{"example2": {DbSize: 1}, "example1": {DbSize: 2}}},
			wantStdOut:     "example1\nexample2",
		},
		{
			name:       "error while getting tables",
			statusErr:  errors.New("some error"),
			wantStdErr: "Received RPC error from Regatta, code 'Unknown' with message 'some error'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpoint := "localhost:8443"

			mclient := &mockClient{}
			regatta = mclient
			mclient.On("Status", mock.Anything, endpoint).Return(tt.statusResponse, tt.statusErr)

			stdoutBuf := new(bytes.Buffer)
			stderrBuf := new(bytes.Buffer)
			RootCmd.SetOut(stdoutBuf)
			RootCmd.SetErr(stderrBuf)

			RootCmd.SetArgs([]string{"table", "--endpoint", endpoint, "--no-color"})
			RootCmd.Execute()

			assert.Equal(t, tt.wantStdOut, strings.TrimSpace(stdoutBuf.String()))
			assert.Equal(t, tt.wantStdErr, strings.TrimSpace(stderrBuf.String()))
		})
	}
}
