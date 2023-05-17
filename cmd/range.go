package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jamf/regatta/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var Range = cobra.Command{
	Use:   "range <table> [key]",
	Short: "Retrieve data from Regatta store",
	Long: "Retrieves data from Regatta store using Range query as defined in API (https://engineering.jamf.com/regatta/api/#range).\n" +
		"You can either retrieve all items from the Regatta by providing no key.\n" +
		"Or you can query for a single item in Regatta by providing item's key.\n" +
		"Or you can query for all items with given prefix, by providing the given prefix and adding the asterisk (*) to the prefix.",
	Example: "regatta-client range table",
	Args:    cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
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
		req := createRangeRequest(args)

		response, err := client.Range(timeoutCtx, req)
		if err != nil && status.Code(err) == codes.NotFound {
			fmt.Println("The requested resource was not found.", err)
			return
		}
		if err != nil {
			fmt.Println("There was an error, while querying Regatta.", err)
			return
		}

		var results = make([]result, 0)
		for _, kv := range response.Kvs {
			results = append(results, result{Key: getValue(kv.Key), Value: getValue(kv.Value)})
		}
		marshal, _ := json.Marshal(results)
		fmt.Println(string(marshal))
		return
	},
}

func createRangeRequest(args []string) *proto.RangeRequest {
	table := args[0]
	if len(args) == 2 {
		key := args[1]
		if strings.HasSuffix(key, "*") {
			key = strings.TrimSuffix(key, "*")
			if len(key) == 0 {
				// get all
				return &proto.RangeRequest{
					Table:    []byte(table),
					Key:      []byte{0},
					RangeEnd: []byte{0},
				}
			}
			// prefix search
			return &proto.RangeRequest{
				Table:    []byte(table),
				Key:      []byte(key),
				RangeEnd: []byte(findNextString(key)),
			}
		} else {
			// get by ID
			return &proto.RangeRequest{
				Table: []byte(table),
				Key:   []byte(key),
			}
		}
	}
	// get all
	return &proto.RangeRequest{
		Table:    []byte(table),
		Key:      []byte{0},
		RangeEnd: []byte{0},
	}
}

func getValue(data []byte) string {
	if binaryData {
		return base64.StdEncoding.EncodeToString(data)
	}
	return string(data)
}

func init() {
	RootCmd.AddCommand(&Range)
}
