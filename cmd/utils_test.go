package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/jamf/regatta/regattapb"
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

func (m *mockKVService) Range(ctx context.Context, req *regattapb.RangeRequest) (*regattapb.RangeResponse, error) {
	called := m.Called(ctx, req)
	return called.Get(0).(*regattapb.RangeResponse), called.Error(1)
}

func (m *mockKVService) Put(ctx context.Context, req *regattapb.PutRequest) (*regattapb.PutResponse, error) {
	called := m.Called(ctx, req)
	return called.Get(0).(*regattapb.PutResponse), called.Error(1)
}

func (m *mockKVService) Delete(ctx context.Context, req *regattapb.DeleteRangeRequest) (*regattapb.DeleteRangeResponse, error) {
	called := m.Called(ctx, req)
	return called.Get(0).(*regattapb.DeleteRangeResponse), called.Error(1)
}

func (m *mockKVService) Txn(ctx context.Context, req *regattapb.TxnRequest) (*regattapb.TxnResponse, error) {
	called := m.Called(ctx, req)
	return called.Get(0).(*regattapb.TxnResponse), called.Error(1)
}
