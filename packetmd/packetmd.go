package packetmd

import (
	"context"
	"crypto/x509"

	"github.com/packethost/hegel-client/hegel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var hegelAddr = "metadata.packet.net:50060"

// WatchResult represents a change in metadata
type WatchResult struct {
	JSON []byte
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
func Watch() (chan WatchResult, chan error, error) {
	resultChanel := make(chan WatchResult)
	errorChannel := make(chan error)
	var currentState string

	hegelClient, close, err := getHegelClient()
	if err != nil {
		return nil, nil, err
	}
	defer close()

	res, err := hegelClient.Get(context.Background(), &hegel.GetRequest{})
	if err != nil {
		errorChannel <- err
	}

	currentState = res.JSON
	resultChanel <- WatchResult{JSON: []byte(currentState)}

	client, err := hegelClient.Subscribe(context.Background(), &hegel.SubscribeRequest{})
	if err != nil {
		return nil, nil, err
	}

	go func() {
		for {
			newResponse, err := client.Recv()
			if err != nil {
				errorChannel <- err
				continue
			}
			newState := newResponse.JSON
			resultChanel <- WatchResult{JSON: []byte(newState)}
			currentState = newState
		}
	}()
	return resultChanel, errorChannel, nil
}
