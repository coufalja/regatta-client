package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jamf/regatta/proto"
	"github.com/spf13/cobra"
)

var RangeOne = cobra.Command{
	Use:     "range-one <table> <key>",
	Example: "range-all example",
	RunE: func(cmd *cobra.Command, args []string) error {
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := createClient(timeout)
		if err != nil {
			fmt.Println("There was an error, while creating client")
			return err
		}

		req := &proto.RangeRequest{
			Table: []byte(args[0]),
			Key:   []byte(args[1]),
		}
		response, err := client.Range(timeout, req)
		if err != nil {
			fmt.Println("There was an error, while querying Regatta")
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
	RootCmd.AddCommand(&RangeOne)
}
