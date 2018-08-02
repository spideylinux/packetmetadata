package cmd

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/buger/jsonparser"

	"github.com/packethost/hegel-client/hegel"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdWatch)
}

var cmdWatch = &cobra.Command{
	Use:   "watch",
	Short: "Watch for metadata changes",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runHegelClient(func(hegelClient hegel.HegelClient) {
			ctx := context.Background()
			currentHW, err := hegelClient.Get(ctx, &hegel.GetRequest{})
			if err != nil {
				cmd.Println("error: ", err)
				return
			}

			keyPath := strings.Split(args[0], ".")

			val, _, _, err := jsonparser.Get([]byte(currentHW.JSON), keyPath...)
			if err != nil {
				cmd.Println("error: ", err)
				return
			}

			var parsedCurrentHw interface{}
			err = json.Unmarshal(val, &parsedCurrentHw)
			if err != nil {
				cmd.Println("error: ", err)
				return
			}

			client, err := hegelClient.Subscribe(ctx, &hegel.SubscribeRequest{})
			if err != nil {
				cmd.Println("error: ", err)
				return
			}

			cmd.Println(parsedCurrentHw)
			for {
				newHw, err := client.Recv()
				if err != nil {
					cmd.Println("error: ", err)
					return
				}

				val, _, _, err := jsonparser.Get([]byte(newHw.JSON), keyPath...)
				if err != nil {
					cmd.Println("error: ", err)
					return
				}

				var parsedNewHw interface{}
				err = json.Unmarshal(val, &parsedNewHw)
				if err != nil {
					cmd.Println("error: ", err)
					return
				}

				if reflect.DeepEqual(parsedNewHw, parsedCurrentHw) {
					continue
				}

				parsedCurrentHw = parsedNewHw

				cmd.Println(parsedNewHw)
			}
		})
	},
}
