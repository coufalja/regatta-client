package cmd

import (
	"context"
	"crypto/tls"

	"github.com/jamf/regatta/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func createClient(ctx context.Context) (proto.KVClient, error) {
	connOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: insecureOption})),
	}

	conn, err := grpc.DialContext(ctx, endpointOption, connOpts...)
	if err != nil {
		return nil, err
	}

	return proto.NewKVClient(conn), nil
}
