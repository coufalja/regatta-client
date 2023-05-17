package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jamf/regatta/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var Delete = cobra.Command{
	Use:   "delete <table> <key>",
	Short: "Delete data from Regatta store",
	Long: "Deletes data from Regatta store using DeleteRange query as defined in API (https://engineering.jamf.com/regatta/api/#deleterange).\n" +
		"You can delete single item in Regatta by providing item's key.\n" +
		"Or you can delete items with given prefix, by providing the given prefix and adding the asterisk (*) to the prefix.\n" +
		"When key or prefix is provided, it needs to be valid UTF-8 string.",
	Example: "regatta-client delete table key\n" +
		"regatta-client delete table 'prefix*'",
	Args: cobra.MatchAll(cobra.ExactArgs(2)),
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
		req := createDeleteRangeRequest(args)

		_, err = client.DeleteRange(timeoutCtx, req)
		if err != nil && status.Code(err) == codes.NotFound {
			fmt.Println("The requested resource was not found.", err)
			return
		}
		if err != nil {
			fmt.Println("There was an error, while querying Regatta.", err)
			return
		}
	},
}

type deleteCommandResult struct {
	Deleted int    `json:"deleted"`
	Keys    string `json:"keys"`
}

func createDeleteRangeRequest(args []string) *proto.DeleteRangeRequest {
	table := args[0]
	key := args[1]
	if strings.HasSuffix(key, "*") {
		key = strings.TrimSuffix(key, "*")
		if len(key) == 0 {
			// delete all
			return &proto.DeleteRangeRequest{
				Table:    []byte(table),
				Key:      []byte{0},
				RangeEnd: []byte{0},
				PrevKv:   true,
			}
		}
		// delete by prefix
		return &proto.DeleteRangeRequest{
			Table:    []byte(table),
			Key:      []byte(key),
			RangeEnd: []byte(findNextString(key)),
			PrevKv:   true,
		}
	} else {
		// delete single
		return &proto.DeleteRangeRequest{
			Table:  []byte(table),
			Key:    []byte(key),
			PrevKv: true,
		}
	}
}
