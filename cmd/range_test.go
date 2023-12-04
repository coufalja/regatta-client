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

func Test_Range_All(t *testing.T) {
	resetRangeFlags()

	mtbl := &mockTable{}
	mtbl.On("Get", mock.Anything, string([]byte{0}), mock.Anything).
		Return(&client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, nil)

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"--limit", "1", "range", "table"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)
	mclient.AssertExpectations(t)
}

func Test_Range_All_Star(t *testing.T) {
	resetRangeFlags()

	mtbl := &mockTable{}
	mtbl.On("Get", mock.Anything, string([]byte{0}), mock.Anything).
		Return(&client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, nil)

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"--limit", "1", "range", "table", "*"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)
	mclient.AssertExpectations(t)
}

func Test_Range_All_Paging(t *testing.T) {
	resetRangeFlags()

	resp1 := client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}, More: true, Count: 1}
	resp2 := client.GetResponse{Kvs: []*client.KeyValue{
		{Key: []byte("test-key"), Value: []byte("test-value")}, {Key: []byte("test-key2"), Value: []byte("test-value2")}}, Count: 2}

	fake, cancel := client.NewFake(
		client.FakeResponse{Response: resp1.OpResponse(), Err: nil},
		client.FakeResponse{Response: resp2.OpResponse(), Err: nil},
	)
	defer cancel()
	regatta = fake.Client()

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"range", "table"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"},{"key":"test-key2","value":"test-value2"}]`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)
}

func Test_Range_Error(t *testing.T) {
	resetRangeFlags()

	fake, cancel := client.NewFake(
		client.FakeResponse{Err: status.Error(codes.NotFound, "table not found")},
	)
	defer cancel()
	regatta = fake.Client()

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"range", "table"})
	RootCmd.Execute()

	assert.Empty(t, stdoutBuf)
	assert.Equal(t, `The requested resource was not found: table not found`, strings.TrimSpace(stderrBuf.String()))
}

func Test_Range_All_Paging_Limit(t *testing.T) {
	resetRangeFlags()

	resp1 := client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}, More: true, Count: 1}
	resp2 := client.GetResponse{Kvs: []*client.KeyValue{
		{Key: []byte("test-key"), Value: []byte("test-value")}, {Key: []byte("test-key2"), Value: []byte("test-value2")}}, Count: 1}

	fake, cancel := client.NewFake(
		client.FakeResponse{Response: resp1.OpResponse(), Err: nil},
		client.FakeResponse{Response: resp2.OpResponse(), Err: nil},
	)
	defer cancel()
	regatta = fake.Client()

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"--limit", "1", "range", "table"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)
}

func Test_Range_Single(t *testing.T) {
	resetRangeFlags()

	mtbl := &mockTable{}
	mtbl.On("Get", mock.Anything, "test-key", mock.Anything).
		Return(&client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, nil)

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"range", "table", "test-key"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)
	mclient.AssertExpectations(t)
}

func Test_Range_Prefix(t *testing.T) {
	resetRangeFlags()

	mtbl := &mockTable{}
	mtbl.On("Get", mock.Anything, "test-key", mock.Anything).
		Return(&client.GetResponse{Kvs: []*client.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, nil)

	mclient := &mockClient{}
	regatta = mclient
	mclient.On("Table", "table").Return(mtbl)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	RootCmd.SetOut(stdoutBuf)
	RootCmd.SetErr(stderrBuf)

	RootCmd.SetArgs([]string{"range", "table", "test-key*"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)

	mclient.AssertExpectations(t)
}

func resetRangeFlags() {
	rangeLimit = 0
	rangeBinary = false
}
