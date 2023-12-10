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

func Test_Version(t *testing.T) {
	tests := []struct {
		name           string
		statusResponse *client.StatusResponse
		statusErr      error
		wantStdOut     string
		wantStdErr     string
	}{
		{
			name:           "print client and server versions",
			statusResponse: &client.StatusResponse{Version: "v3.2.1"},
			wantStdOut:     "client version: v1.2.3\nserver version: v3.2.1",
		},
		{
			name:       "error getting server version",
			statusErr:  errors.New("some error"),
			wantStdOut: "client version: v1.2.3\nserver version: unknown",
			wantStdErr: "some error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RootCmd.Version = "v1.2.3"
			endpoint := "localhost:8443"

			mclient := &mockClient{}
			regatta = mclient
			mclient.On("Status", mock.Anything, endpoint).Return(tt.statusResponse, tt.statusErr)

			stdoutBuf := new(bytes.Buffer)
			stderrBuf := new(bytes.Buffer)
			RootCmd.SetOut(stdoutBuf)
			RootCmd.SetErr(stderrBuf)

			RootCmd.SetArgs([]string{"version", "--endpoint", endpoint, "--no-color"})
			RootCmd.Execute()

			assert.Equal(t, tt.wantStdOut, strings.TrimSpace(stdoutBuf.String()))
			assert.Equal(t, tt.wantStdErr, strings.TrimSpace(stderrBuf.String()))
		})
	}
}
