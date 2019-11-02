package internal

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	service "{{ .service.Package }}/generated"
	"log"
)

type Server struct {
	service.{{ .service.Service }}Server
}

func (s *Server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	name := req.GetName()
	if name == "" {
		name = service.HelloRequest_world.String()
	}
	return &service.HelloReply{Message: "Hello, " + name}, nil
}

func NewServer() service.{{ .service.Service }}Server {
	return &Server{}
}
