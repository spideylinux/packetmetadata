package cmd

import (
	"context"

	"github.com/packethost/hegel-client/hegel"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdGet)
}

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "Get metadata",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := hegelClient.Get(context.Background(), &hegel.GetRequest{})
		if err != nil {
			cmd.Println("error: ", err)
		} else {
			cmd.Println(res.JSON)
		}
	},
}
