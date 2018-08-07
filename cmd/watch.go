package cmd

import (
	"github.com/packethost/hegel-client/packetmd"
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
		iterator, err := packetmd.Watch()
		if err != nil {
			cmd.Println("error: ", err)
			return
		}

		for {
			newState, err := iterator.Next()
			cmd.Println("called next")
			if err != nil {
				cmd.Println("error: ", err)
				return
			}
			cmd.Println(string(newState.JSON))
		}
	},
}
