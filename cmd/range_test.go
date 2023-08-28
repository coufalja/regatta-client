package cmd

import (
	"bytes"
	"net"
	"strings"
	"testing"

	"github.com/jamf/regatta/proto"
	"github.com/jamf/regatta/regattaserver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Test_Range_All(t *testing.T) {
	resetRangeFlags()
	
	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(generateTLSConfig())))

	storage := new(mockKVService)
	storage.On("Range", mock.Anything, &proto.RangeRequest{Table: []byte("table"), Key: zero, RangeEnd: zero, Limit: 1}).
		Return(&proto.RangeResponse{Kvs: []*proto.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, nil)

	proto.RegisterKVServer(s, &regattaserver.KVServer{Storage: storage})
	go s.Serve(lis)
	defer s.Stop()

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", lis.Addr().String(), "--cert", "test.crt", "--limit", "1", "range", "table"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(buf.String()))
}

func Test_Range_All_Star(t *testing.T) {
	resetRangeFlags()

	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(generateTLSConfig())))

	storage := new(mockKVService)
	storage.On("Range", mock.Anything, &proto.RangeRequest{Table: []byte("table"), Key: zero, RangeEnd: zero, Limit: 1}).
		Return(&proto.RangeResponse{Kvs: []*proto.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, nil)

	proto.RegisterKVServer(s, &regattaserver.KVServer{Storage: storage})
	go s.Serve(lis)
	defer s.Stop()

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", lis.Addr().String(), "--cert", "test.crt", "--limit", "1", "range", "table", "*"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(buf.String()))
}

func Test_Range_Single(t *testing.T) {
	resetRangeFlags()

	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(generateTLSConfig())))

	storage := new(mockKVService)
	storage.On("Range", mock.Anything, &proto.RangeRequest{Table: []byte("table"), Key: []byte("test-key")}).
		Return(&proto.RangeResponse{Kvs: []*proto.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, nil)

	proto.RegisterKVServer(s, &regattaserver.KVServer{Storage: storage})
	go s.Serve(lis)
	defer s.Stop()

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", lis.Addr().String(), "--cert", "test.crt", "range", "table", "test-key"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(buf.String()))
}

func Test_Range_Prefix(t *testing.T) {
	resetRangeFlags()

	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(generateTLSConfig())))

	storage := new(mockKVService)
	storage.On("Range", mock.Anything, &proto.RangeRequest{Table: []byte("table"), Key: []byte("test-key"), RangeEnd: []byte("test-kez")}).
		Return(&proto.RangeResponse{Kvs: []*proto.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, nil)

	proto.RegisterKVServer(s, &regattaserver.KVServer{Storage: storage})
	go s.Serve(lis)
	defer s.Stop()

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", lis.Addr().String(), "--cert", "test.crt", "range", "table", "test-key*"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(buf.String()))
}

func resetRangeFlags() {
	rangeLimit = 0
	rangeBinary = false
}
