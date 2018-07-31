package cmd

import (
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/packethost/hegel-client/hegel"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var hegelClient hegel.HegelClient

func init() {
	cred := credentials.NewClientTLSFromCert(x509.NewCertPool(), "")
	conn, err := grpc.Dial("hegel.packet.net:50060", grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	hegelClient = hegel.NewHegelClient(conn)
}

var rootCmd = &cobra.Command{
	Use:   "hegel-client",
	Short: "Packet metadata client",
	Long:  `Hegel is Packet's gRPC metadata service, this is a client for it`,
}

// Execute is the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
