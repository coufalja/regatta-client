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

func Test_Delete(t *testing.T) {
	mtbl := &mockTable{}
	mtbl.On("Delete", mock.Anything, "key", mock.Anything).Return(&client.DeleteResponse{Deleted: 1}, nil)

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"delete", "table", "key", "--no-color"})
	RootCmd.Execute()

	assert.Equal(t, `1`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)
	mclient.AssertExpectations(t)
}

func Test_Delete_Prefix(t *testing.T) {
	mtbl := &mockTable{}
	mtbl.On("Delete", mock.Anything, "key", mock.Anything).Return(&client.DeleteResponse{Deleted: 2}, nil)

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"delete", "table", "key*", "--no-color"})
	RootCmd.Execute()

	assert.Equal(t, `2`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)
	mclient.AssertExpectations(t)
}

func Test_Delete_All(t *testing.T) {
	mtbl := &mockTable{}
	mtbl.On("Delete", mock.Anything, string([]byte{0}), mock.Anything).Return(&client.DeleteResponse{Deleted: 3}, nil)

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"delete", "table", "*", "--no-color"})
	RootCmd.Execute()

	assert.Equal(t, `3`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)
	mclient.AssertExpectations(t)
}

func Test_Delete_Error(t *testing.T) {
	mtbl := &mockTable{}
	mtbl.On("Delete", mock.Anything, "key", mock.Anything).Return(&client.DeleteResponse{}, status.Error(codes.NotFound, "table not found"))

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"delete", "table", "key"})
	RootCmd.Execute()

	assert.Empty(t, stdoutBuf)
	assert.Equal(t, `The requested resource was not found: table not found`, strings.TrimSpace(stderrBuf.String()))
	mclient.AssertExpectations(t)
}
