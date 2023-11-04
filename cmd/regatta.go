package cmd

import (
	client "github.com/jamf/regatta-go"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type logger struct{}

func (p logger) Infof(_ string, _ ...any)  {}
func (p logger) Debugf(_ string, _ ...any) {}
func (p logger) Warnf(_ string, _ ...any)  {}
func (p logger) Errorf(_ string, _ ...any) {}

func connect(cmd *cobra.Command, _ []string) error {
	// Allow for mocking in tests.
	if regatta == nil {
		cc, err := client.NewClientConfig(&client.ConfigSpec{
			Logger:    logger{},
			Endpoints: []string{endpointOption},
			Secure: &client.SecureConfig{
				Cacert:             certOption,
				InsecureSkipVerify: insecureOption,
			},
			DialTimeout: dialTimeout,
		})
		if err != nil {
			cmd.PrintErrln("There was an error, with config of connection to Regatta.")
			return err
		}
		cc.DialOptions = append(cc.DialOptions, grpc.WithBlock(), grpc.WithReturnConnectionError())
		cl, err := client.New(cc)
		if err != nil {
			cmd.PrintErrln("There was an error, while establishing connection to Regatta.")
			return err
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
