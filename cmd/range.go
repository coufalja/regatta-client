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

var RangeAll = cobra.Command{
	Use:     "range <table> [key]",
	Example: "regatta-client range table",
	Args:    cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := createClient(timeout)
		if err != nil {
			fmt.Println("There was an error, while creating client", err)
			return
		}

		req := createRangeRequest(args)

		response, err := client.Range(timeout, req)
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
	if len(args) == 2 {
		key := args[1]
		if strings.HasSuffix(key, "*") {
			key = strings.TrimSuffix(key, "*")
			return &proto.RangeRequest{
				Table:    []byte(args[0]),
				Key:      []byte(key),
				RangeEnd: []byte(findNextString(key)),
			}
		} else {
			return &proto.RangeRequest{
				Table: []byte(args[0]),
				Key:   []byte(args[1]),
			}
		}
	}
	return &proto.RangeRequest{
		Table:    []byte(args[0]),
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

func findNextString(str string) string {
	// Convert string to byte slice for mutation
	bytes := []byte(str)

	// Start from the last character and increment its byte value
	i := len(bytes) - 1
	for i >= 0 {
		if bytes[i] < 255 {
			bytes[i]++
			break
		}
		bytes[i] = 0
		i--
	}

	return string(bytes)
}

func init() {
	RootCmd.AddCommand(&RangeAll)
}
