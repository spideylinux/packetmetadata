package packetmd

import (
	"context"
	"crypto/x509"
	"fmt"

	"github.com/packethost/hegel-client/hegel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var hegelAddr = "metadata.packet.net:50060"

// WatchResult represents a change in metadata
type WatchResult struct {
	JSON []byte
}

// WatchIterator is a struct for iterating over watch results
type WatchIterator struct {
	Next  func() WatchResult
	Close func() error
}

func getHegelClient() (hegel.HegelClient, func() error, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, nil, err
	}
	cred := credentials.NewClientTLSFromCert(certPool, "")
	conn, err := grpc.Dial(hegelAddr, grpc.WithTransportCredentials(cred))
	if err != nil {
		return nil, nil, err
	}
	return hegel.NewHegelClient(conn), conn.Close, nil
}

// Get returns instance metadata from hegel
func Get() ([]byte, error) {
	hegelClient, close, err := getHegelClient()
	if err != nil {
		return nil, err
	}
	defer close()
	res, err := hegelClient.Get(context.Background(), &hegel.GetRequest{})
	if err != nil {
		return nil, err
	}
	return []byte(res.JSON), nil
}

// Watch returns a channel that outputs JSON
func Watch() (*WatchIterator, error) {
	var currentState string

	hegelClient, close, err := getHegelClient()
	if err != nil {
		return nil, err
	}

	res, err := hegelClient.Get(context.Background(), &hegel.GetRequest{})
	if err != nil {
		return nil, err
	}

	currentState = res.JSON

	iterator := &WatchIterator{
		Next: func() WatchResult {
			return WatchResult{
				JSON: []byte(currentState),
			}
		},
		Close: close,
	}

	client, err := hegelClient.Subscribe(context.Background(), &hegel.SubscribeRequest{})
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			newResponse, err := client.Recv()
			if err != nil {
				fmt.Println(err)
			}
			newState := newResponse.JSON
			currentState = newState
		}
	}()
	return iterator, nil
}
