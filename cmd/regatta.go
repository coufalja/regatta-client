package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/jamf/regatta/regattapb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func createClient() (regattapb.KVClient, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	if len(certOption) != 0 {
		certs, err := os.ReadFile(certOption)
		if err != nil {
			return nil, err
		}
		pool.AppendCertsFromPEM(certs)
	}

	connOpts := []grpc.DialOption{
		// nolint:gosec
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{RootCAs: pool, InsecureSkipVerify: insecureOption})),
	}

	conn, err := grpc.Dial(endpointOption, connOpts...)
	if err != nil {
		return nil, err
	}

	return regattapb.NewKVClient(conn), nil
}

func handleRegattaError(cmd *cobra.Command, err error) {
	if st := status.Convert(err); st != nil {
		switch st.Code() {
		case codes.NotFound:
			cmd.PrintErrln("The requested resource was not found:", st.Message())
		case codes.Unavailable:
			cmd.PrintErrln("Regatta is not reachable:", st.Message())
		default:
			cmd.PrintErrf("Received RPC error from Regatta, code '%s' with message '%s'\n", st.Code(), st.Message())
		}
	} else {
		cmd.PrintErrln("There was an error, while querying Regatta.", err)
	}
}
