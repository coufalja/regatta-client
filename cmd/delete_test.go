package cmd

import (
	"bytes"
	"net"
	"testing"

	"github.com/jamf/regatta/regattapb"
	"github.com/jamf/regatta/regattaserver"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func Test_Delete(t *testing.T) {
	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(generateTLSConfig())))

	storage := new(mockKVService)
	storage.On("Delete", mock.Anything, mock.Anything).Return(&regattapb.DeleteRangeResponse{}, nil)

	regattapb.RegisterKVServer(s, &regattaserver.KVServer{Storage: storage})
	go s.Serve(lis)
	defer s.Stop()

	buf := new(bytes.Buffer)
	RootCmd.SetOut(buf)
	RootCmd.SetArgs([]string{"--endpoint", lis.Addr().String(), "--cert", "test.crt", "delete", "table", "key"})
	RootCmd.Execute()

	storage.AssertExpectations(t)
}
