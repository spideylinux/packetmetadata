package cmd

import (
	"context"

	"github.com/packethost/hegel-client/hegel"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdWatch)
}

var cmdWatch = &cobra.Command{
	Use:   "watch",
	Short: "Watch for metadata changes",
	Run: func(cmd *cobra.Command, args []string) {
		runHegelClient(func(hegelClient hegel.HegelClient) {
			client, err := hegelClient.Subscribe(context.Background(), &hegel.SubscribeRequest{})
			if err != nil {
				cmd.Println("error: ", err)
				return
			}
			for {
				hw, err := client.Recv()
				if err != nil {
					cmd.Println("error: ", err)
					return
				}
				cmd.Println(hw.JSON)
			}
		})
	},
}
