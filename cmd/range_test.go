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

	assert.Equal(t, `test-key: test-value`, strings.TrimSpace(stdoutBuf.String()))
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

	assert.Equal(t, `test-key: test-value`, strings.TrimSpace(stdoutBuf.String()))
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

	assert.Equal(t, "test-key: test-value\ntest-key2: test-value2", strings.TrimSpace(stdoutBuf.String()))
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

	assert.Equal(t, `test-key: test-value`, strings.TrimSpace(stdoutBuf.String()))
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

	assert.Equal(t, `test-key: test-value`, strings.TrimSpace(stdoutBuf.String()))
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

	assert.Equal(t, `test-key: test-value`, strings.TrimSpace(stdoutBuf.String()))
	assert.Empty(t, stderrBuf)
	mclient.AssertExpectations(t)
}

func Test_Range_Output(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantStdOut string
		wantStdErr string
	}{
		{
			name:       "json output",
			args:       []string{"--output", "json", "range", "table"},
			wantStdOut: `[{"key":"test-key","value":"test-value"}]`,
		},
		{
			name:       "plain output",
			args:       []string{"--output", "plain", "range", "table"},
			wantStdOut: `test-key: test-value`,
		},
		{
			name:       "invalid output",
			args:       []string{"--output", "invalid", "range", "table"},
			wantStdErr: `Error: invalid argument "invalid" for "--output" flag: must be one of: plain, json`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

			RootCmd.SetArgs(tt.args)
			RootCmd.Execute()

			assert.Equal(t, tt.wantStdOut, strings.TrimSpace(stdoutBuf.String()))
			assert.Equal(t, tt.wantStdErr, strings.TrimSpace(stderrBuf.String()))
		})
	}
}

func resetRangeFlags() {
	rangeLimit = 0
	rangeBinary = false
}
