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

func Test_Range(t *testing.T) {
	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(generateTLSConfig())))

	storage := new(mockKVService)
	storage.On("Range", mock.Anything, mock.Anything).
		Return(&proto.RangeResponse{Kvs: []*proto.KeyValue{{Key: []byte("test-key"), Value: []byte("test-value")}}}, nil)

	proto.RegisterKVServer(s, &regattaserver.KVServer{Storage: storage})
	go s.Serve(lis)
	defer s.Stop()

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.PersistentFlags().Set("endpoint", lis.Addr().String())
	RootCmd.PersistentFlags().Set("insecure", "true")
	RootCmd.SetArgs([]string{"range", "table"})
	RootCmd.Execute()

	assert.Equal(t, `[{"key":"test-key","value":"test-value"}]`, strings.TrimSpace(buf.String()))
}
