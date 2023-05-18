package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/jamf/regatta/proto"
	"github.com/stretchr/testify/mock"
)

func generateTLSConfig() *tls.Config {
	cert, err := tls.LoadX509KeyPair("test.crt", "test.key")
	if err != nil {
		panic(err)
	}

	certs, err := os.ReadFile("test.crt")
	if err != nil {
		panic(err)
	}

	pool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}
	pool.AppendCertsFromPEM(certs)

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"doq"},
		RootCAs:      pool,
	}
}

type mockKVService struct {
	mock.Mock
}

func (m *mockKVService) Range(ctx context.Context, req *proto.RangeRequest) (*proto.RangeResponse, error) {
	called := m.Called(ctx, req)
	return called.Get(0).(*proto.RangeResponse), called.Error(1)
}

func (m *mockKVService) Put(ctx context.Context, req *proto.PutRequest) (*proto.PutResponse, error) {
	called := m.Called(ctx, req)
	return called.Get(0).(*proto.PutResponse), called.Error(1)
}

func (m *mockKVService) Delete(ctx context.Context, req *proto.DeleteRangeRequest) (*proto.DeleteRangeResponse, error) {
	called := m.Called(ctx, req)
	return called.Get(0).(*proto.DeleteRangeResponse), called.Error(1)
}

func (m *mockKVService) Txn(ctx context.Context, req *proto.TxnRequest) (*proto.TxnResponse, error) {
	called := m.Called(ctx, req)
	return called.Get(0).(*proto.TxnResponse), called.Error(1)
}
