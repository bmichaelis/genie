package api

import (
	"context"
	"google.golang.org/grpc"
	grpcService "{{{.Package}}}/generated"
	"time"
)

type Clienter interface {
	Connect() error
	Disconnect() error
	SayHello(name string) (string, error)
}

type Client struct {
	Connection    *grpc.ClientConn
	GreeterClient grpcService.GreeterClient
}

func (client *Client) Connect() error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*GrpcAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client.Connection = conn
	client.GreeterClient = grpcService.NewGreeterClient(conn)
	return nil
}

func (client *Client) Disconnect() error {
	if client.Connection != nil {
		if err := client.Connection.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (client *Client) SayHello(name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GreeterClient.SayHello(ctx, &grpcService.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}

func NewClient() Clienter {
	return &Client{}
}
