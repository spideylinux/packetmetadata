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

const hegelAddr = "metadata.packet.net:50060"

func runHegelClient(f func(hegel.HegelClient)) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	cred := credentials.NewClientTLSFromCert(certPool, "")
	conn, err := grpc.Dial(hegelAddr, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatal(err)
	}
	f(hegel.NewHegelClient(conn))
	conn.Close()
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
