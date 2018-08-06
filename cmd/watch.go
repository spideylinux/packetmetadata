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
		// runHegelClient(func(hegelClient hegel.HegelClient) {
		// 	ctx := context.Background()
		// 	currentHW, err := hegelClient.Get(ctx, &hegel.GetRequest{})
		// 	if err != nil {
		// 		cmd.Println("error: ", err)
		// 		return
		// 	}

		// 	val := gjson.Get(currentHW.JSON, args[0])
		// 	if err != nil {
		// 		cmd.Println("error: ", err)
		// 		return
		// 	}

		// 	parsedCurrentHw := val.Value()

		// 	client, err := hegelClient.Subscribe(ctx, &hegel.SubscribeRequest{})
		// 	if err != nil {
		// 		cmd.Println("error: ", err)
		// 		return
		// 	}

		// 	cmd.Println(parsedCurrentHw)
		// 	for {
		// 		newHw, err := client.Recv()
		// 		if err != nil {
		// 			cmd.Println("error: ", err)
		// 			return
		// 		}

		// 		val := gjson.Get(newHw.JSON, args[0])
		// 		if err != nil {
		// 			cmd.Println("error: ", err)
		// 			return
		// 		}

		// 		parsedNewHw := val.Value()

		// 		if reflect.DeepEqual(parsedNewHw, parsedCurrentHw) {
		// 			continue
		// 		}

		// 		patch, err := jsonpatch.CreateMergePatch([]byte(currentHW.JSON), []byte(newHw.JSON))
		// 		if err != nil {
		// 			cmd.Println("error: ", err)
		// 			return
		// 		}

		// 		parsedCurrentHw = parsedNewHw

		// 		cmd.Println(parsedNewHw, string(patch))
		// 	}
		// })
		res, errs, err := packetmd.Watch()
		if err != nil {
			cmd.Println("error: ", err)
			return
		}

		go func() {
			for {
				newError := <-errs
				cmd.Println("error: ", newError)
			}
		}()

		for {
			newState := <-res
			cmd.Println(string(newState.JSON))
		}
	},
}
