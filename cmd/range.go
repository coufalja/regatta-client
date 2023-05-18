package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/jamf/regatta/proto"
	"github.com/spf13/cobra"
)

var rangeBinary bool

func init() {
	Range.Flags().BoolVar(&rangeBinary, "binary", false, "avoid decoding keys and values into UTF-8 strings, but rather encode them as Base64 strings")
}

// Range is a subcommand used for retrieving records from a table.
var Range = cobra.Command{
	Use:   "range <table> [key]",
	Short: "Retrieve data from Regatta store",
	Long: "Retrieves data from Regatta store using Range query as defined in API (https://engineering.jamf.com/regatta/api/#range).\n" +
		"You can either retrieve all items from the Regatta by providing no key.\n" +
		"Or you can query for a single item in Regatta by providing item's key.\n" +
		"Or you can query for all items with given prefix, by providing the given prefix and adding the asterisk (*) to the prefix.\n" +
		"When key or prefix is provided, it needs to be valid UTF-8 string.",
	Example: "regatta-client range table\n" +
		"regatta-client range table key\n" +
		"regatta-client range table 'prefix*'",
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := createClient()
		if err != nil {
			cmd.PrintErrln("There was an error, while creating client.", err)
			return
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		req := createRangeRequest(args)

		response, err := client.Range(timeoutCtx, req)
		if err != nil {
			handleRegattaError(cmd, err)
			return
		}

		results := make([]rangeCommandResult, 0)
		for _, kv := range response.Kvs {
			results = append(results, rangeCommandResult{Key: getValue(kv.Key), Value: getValue(kv.Value)})
		}
		marshal, _ := json.Marshal(results)
		cmd.Println(string(marshal))
	},
}

type rangeCommandResult struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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
		}
		// get by ID
		return &proto.RangeRequest{
			Table: []byte(table),
			Key:   []byte(key),
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
	if rangeBinary {
		return base64.StdEncoding.EncodeToString(data)
	}
	return string(data)
}
