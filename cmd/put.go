package cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/jamf/regatta/proto"
	"github.com/spf13/cobra"
)

var putBinary bool

func init() {
	Put.Flags().BoolVar(&putBinary, "binary", false, "provided <value> is binary data encoded using Base64")
}

var Put = cobra.Command{
	Use:     "put <table> <key> <value>",
	Short:   "Put data into Regatta store",
	Long:    "Put data into Regatta store using Put query as defined in API (https://engineering.jamf.com/regatta/api/#put).",
	Example: "regatta-client put table key value",
	Args:    cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		connectTimeoutCtx, connectCancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer connectCancel()

		client, err := createClient(connectTimeoutCtx)
		if err != nil {
			fmt.Println("There was an error, while creating client.", err)
			return
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		req, err := createPutRequest(args)
		if err != nil {
			fmt.Println("There was an error while decoding parameters.", err)
			return
		}

		_, err = client.Put(timeoutCtx, req)
		if err != nil {
			fmt.Println("There was an error, while querying Regatta.", err)
			return
		}
	},
}

func createPutRequest(args []string) (*proto.PutRequest, error) {
	table := []byte(args[0])
	key := []byte(args[1])
	var value []byte
	if putBinary {
		var err error
		value, err = base64.StdEncoding.DecodeString(args[2])
		if err != nil {
			return nil, err
		}
	} else {
		value = []byte(args[2])
	}

	return &proto.PutRequest{Table: table, Key: key, Value: value}, nil
}
