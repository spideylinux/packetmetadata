package cmd

import (
	"github.com/packethost/hegel-client/packetmd"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdGet)
}

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "Get metadata",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := packetmd.Get()
		if err != nil {
			cmd.Println("error: ", err)
			return
		}
		cmd.Println(string(res))
	},
}
