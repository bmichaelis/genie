package internal

import (
	"context"
	service "{{.Package}}/generated"
	"log"
)

type Server struct {
	service.{{.Service}}Server
}

func (s *Server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	name := req.GetName()
	if name == "" {
		name = service.HelloRequest_world.String()
	}
	return &service.HelloReply{Message: "Hello, " + name}, nil
}

func NewServer() service.{{ .Service }}Server {
	return &Server{}
}
