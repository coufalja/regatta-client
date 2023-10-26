package cmd

import (
	"bytes"
	"testing"

	client "github.com/jamf/regatta-go"
	"github.com/stretchr/testify/mock"
)

func Test_Delete(t *testing.T) {
	mtbl := &mockTable{}
	mtbl.On("Delete", mock.Anything, "key", mock.Anything).Return(&client.DeleteResponse{}, error(nil))

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", "localhost:8443", "--cert", "test.crt", "delete", "table", "key"})
	RootCmd.Execute()

	mclient.AssertExpectations(t)
}
