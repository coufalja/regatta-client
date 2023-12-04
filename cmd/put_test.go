package cmd

import (
	"bytes"
	"strings"
	"testing"

	client "github.com/jamf/regatta-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_Put(t *testing.T) {
	resetPutFlags()
	mtbl := &mockTable{}
	mtbl.On("Put", mock.Anything, "key", "data", mock.Anything).Return(&client.PutResponse{}, nil)

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"put", "table", "key", "data"})
	RootCmd.Execute()

	assert.Empty(t, stdoutBuf)
	assert.Empty(t, stderrBuf)
	mclient.AssertExpectations(t)
}

func Test_Put_Error(t *testing.T) {
	resetPutFlags()
	mtbl := &mockTable{}
	mtbl.On("Put", mock.Anything, "key", "data", mock.Anything).Return(&client.PutResponse{}, status.Error(codes.NotFound, "table not found"))

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"put", "table", "key", "data"})
	RootCmd.Execute()

	assert.Empty(t, stdoutBuf)
	assert.Equal(t, `The requested resource was not found: table not found`, strings.TrimSpace(stderrBuf.String()))
	mclient.AssertExpectations(t)
}

func resetPutFlags() {
	putBinary = false
}
