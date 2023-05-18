package cmd

import (
	"crypto/tls"
	"fmt"

	"github.com/jamf/regatta/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func createClient() (proto.KVClient, error) {
	connOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: insecureOption})),
	}

	conn, err := grpc.Dial(endpointOption, connOpts...)
	if err != nil {
		return nil, err
	}

	return proto.NewKVClient(conn), nil
}

func handleRegattaError(err error) {
	if st := status.Convert(err); st != nil {
		switch st.Code() {
		case codes.NotFound:
			fmt.Println("The requested resource was not found:", st.Message())
		case codes.Unavailable:
			fmt.Println("Regatta is not reachable:", st.Message())
		default:
			fmt.Printf("Received RPC error from Regatta, code '%s' with message '%s'\n", st.Code(), st.Message())
		}
	} else {
		fmt.Println("There was an error, while querying Regatta.", err)
	}
}
