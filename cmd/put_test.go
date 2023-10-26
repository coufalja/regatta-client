package cmd

import (
	"bytes"
	"testing"

	client "github.com/jamf/regatta-go"
	"github.com/stretchr/testify/mock"
)

func Test_Put(t *testing.T) {
	resetPutFlags()
	mtbl := &mockTable{}
	mtbl.On("Put", mock.Anything, "key", "data", mock.Anything).Return(&client.PutResponse{}, error(nil))

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", "localhost:8443", "--cert", "test.crt", "put", "table", "key", "data"})
	RootCmd.Execute()

	mclient.AssertExpectations(t)
}

func resetPutFlags() {
	putBinary = false
}
