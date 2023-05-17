package cmd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jamf/regatta/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type result struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var GetPrefix = cobra.Command{
	Use:     "range-all <table>",
	Example: "range-all example",
	RunE: func(cmd *cobra.Command, args []string) error {

		connOpts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: insecureOption})),
		}

		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(timeout, endpointOption, connOpts...)
		if err != nil {
			fmt.Println("There was error, while connecting to Regatta")
			return err
		}

		req := &proto.RangeRequest{
			Table:    []byte(args[0]),
			Key:      []byte{0},
			RangeEnd: []byte{0},
		}
		client := proto.NewKVClient(conn)
		response, err := client.Range(timeout, req)
		if err != nil {
			fmt.Println("There was error, while querying Regatta")
			return err
		}

		var results []result
		for _, kv := range response.Kvs {
			results = append(results, result{Key: string(kv.Key), Value: string(kv.Value)})
		}
		marshal, _ := json.Marshal(results)
		fmt.Println(string(marshal))
		return nil
	},
}

func init() {
	RootCmd.AddCommand(&GetPrefix)
}
