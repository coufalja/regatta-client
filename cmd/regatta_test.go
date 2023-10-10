package cmd

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	client "github.com/jamf/regatta-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_handleRegattaError(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantMsg string
	}{
		{
			name:    "generic error",
			err:     errors.New("some error"),
			wantMsg: `Received RPC error from Regatta, code 'Unknown' with message 'some error'`,
		},
		{
			name:    "internal Regatta error",
			err:     status.Error(codes.Internal, "internal Regatta error"),
			wantMsg: `Received RPC error from Regatta, code 'Internal' with message 'internal Regatta error'`,
		},
		{
			name:    "not found Regatta error",
			err:     status.Error(codes.NotFound, "resource not found"),
			wantMsg: `The requested resource was not found: resource not found`,
		},
		{
			name:    "unavailable Regatta error",
			err:     status.Error(codes.Unavailable, "resource unavailable"),
			wantMsg: `Regatta is not reachable: resource unavailable`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			RootCmd.SetErr(buf)

			handleRegattaError(&RootCmd, tt.err)

			assert.Equal(t, tt.wantMsg, strings.TrimSpace(buf.String()))
		})
	}
}

type mockClient struct {
	client.Client
	mock.Mock
}

type mockTable struct {
	mock.Mock
}

func (m *mockTable) Put(ctx context.Context, key, val string, opts ...client.OpOption) (*client.PutResponse, error) {
	args := m.Called(ctx, key, val, opts)
	return args.Get(0).(*client.PutResponse), args.Error(1)
}

func (m *mockTable) Get(ctx context.Context, key string, opts ...client.OpOption) (*client.GetResponse, error) {
	args := m.Called(ctx, key, opts)
	return args.Get(0).(*client.GetResponse), args.Error(1)
}

func (m *mockTable) Delete(ctx context.Context, key string, opts ...client.OpOption) (*client.DeleteResponse, error) {
	args := m.Called(ctx, key, opts)
	return args.Get(0).(*client.DeleteResponse), args.Error(1)
}

func (m *mockTable) Txn(ctx context.Context) client.Txn {
	args := m.Called(ctx)
	return args.Get(0).(client.Txn)
}

func (m *mockClient) Table(table string) client.Table {
	called := m.Called(table)
	return called.Get(0).(*mockTable)
}
