package cmd

import (
	"github.com/packethost/packetmetadata/packetmetadata"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdGet)
}

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "Get metadata",
	Run: func(cmd *cobra.Command, args []string) {
		res, _, err := packetmetadata.Get()
		if err != nil {
			cmd.Println("error: ", err)
			return
		}
		cmd.Println(string(res))
	},
}
