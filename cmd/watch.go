package cmd

import (
	jsonpatch "github.com/evanphx/json-patch"
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

		currentState, err := packetmd.Get()
		if err != nil {
			cmd.Println("error: ", err)
			return
		}

		for {
			next, err := iterator.Next()
			newState := next.JSON
			if err != nil {
				cmd.Println("error: ", err)
				return
			}
			patch, err := jsonpatch.CreateMergePatch(currentState, newState)
			if err != nil {
				cmd.Println("error: ", err)
				return
			}
			currentState = newState
			cmd.Println(string(patch))
		}
	},
}
