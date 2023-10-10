package cmd

import (
	"bytes"
	"strings"
	"testing"

	client "github.com/jamf/regatta-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Range_All(t *testing.T) {
	resetRangeFlags()

	mtbl := &mockTable{}
	mtbl.On("Get", mock.Anything, string([]byte{0}), mock.Anything).Return(&client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, error(nil))

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", "localhost:8443", "--cert", "test.crt", "--limit", "1", "range", "table"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(buf.String()))
}

func Test_Range_All_Star(t *testing.T) {
	resetRangeFlags()

	mtbl := &mockTable{}
	mtbl.On("Get", mock.Anything, string([]byte{0}), mock.Anything).Return(&client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, error(nil))

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", "localhost:8443", "--cert", "test.crt", "--limit", "1", "range", "table", "*"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(buf.String()))
	mclient.AssertExpectations(t)
}

func Test_Range_Single(t *testing.T) {
	resetRangeFlags()

	mtbl := &mockTable{}
	mtbl.On("Get", mock.Anything, "test-key", mock.Anything).Return(&client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, error(nil))

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", "localhost:8443", "--cert", "test.crt", "range", "table", "test-key"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(buf.String()))
	mclient.AssertExpectations(t)
}

func Test_Range_Prefix(t *testing.T) {
	resetRangeFlags()

	mtbl := &mockTable{}
	mtbl.On("Get", mock.Anything, "test-key", mock.Anything).Return(&client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, error(nil))

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", "localhost:8443", "--cert", "test.crt", "range", "table", "test-key*"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(buf.String()))
	mclient.AssertExpectations(t)
}

func resetRangeFlags() {
	rangeLimit = 0
	rangeBinary = false
}
