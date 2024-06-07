package cmd

import (
	"fmt"

	client "github.com/jamf/regatta-go"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func connect(_ *cobra.Command, _ []string) error {
	// Allow for mocking in tests.
	if regatta == nil {
		cl, err := client.New(
			client.WithEndpoints(endpoint),
			client.WithBlock(),
			client.WithReturnConnectionError(),
			client.WithDialTimeout(dialTimeout),
			client.WithSecureConfig(&client.SecureConfig{Cacert: cert, InsecureSkipVerify: insecure}),
		)
		if err != nil {
			return fmt.Errorf("unable to establish connection to the Regatta: %w", err)
		}
		regatta = cl
	}
	return nil
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
