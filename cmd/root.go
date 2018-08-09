package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var hegelAddr = os.Getenv("HEGEL_ADDR")

func init() {
	if hegelAddr == "" {
		hegelAddr = "metadata.packet.net:50060"
	}
}

var rootCmd = &cobra.Command{
	Use:   "packet-metadata",
	Short: "Packet metadata client",
	Long:  `This is a client for Packet's gRPC metadata service`,
}

// Execute is the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
