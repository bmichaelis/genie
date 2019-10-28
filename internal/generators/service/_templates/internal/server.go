package internal

import (
	"context"
	grpcService "{{.Package}}/generated"
	"log"
)

type Serverer interface {
	SayHello(ctx context.Context, in *grpcService.HelloRequest) (*grpcService.HelloReply, error)
}

type Server struct {
	grpcService.GreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *grpcService.HelloRequest) (*grpcService.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	name := in.GetName()
	if name == "" {
		name = grpcService.HelloRequest_world.String()
	}
	return &grpcService.HelloReply{Message: "Hello, " + name}, nil
}

func NewServer() Serverer {
	return &Server{}
}
