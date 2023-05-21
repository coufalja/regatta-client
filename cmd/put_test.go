package cmd

import (
	"bytes"
	"net"
	"testing"

	"github.com/jamf/regatta/proto"
	"github.com/jamf/regatta/regattaserver"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Test_Put(t *testing.T) {
	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(generateTLSConfig())))

	storage := new(mockKVService)
	storage.On("Put", mock.Anything, mock.Anything).Return(&proto.PutResponse{}, nil)

	proto.RegisterKVServer(s, &regattaserver.KVServer{Storage: storage})
	go s.Serve(lis)
	defer s.Stop()

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.PersistentFlags().Set("endpoint", lis.Addr().String())
	RootCmd.PersistentFlags().Set("cert", "test.crt")
	RootCmd.SetArgs([]string{"put", "table", "key", "data"})
	RootCmd.Execute()

	storage.AssertExpectations(t)
}
