package internal

import (
	"context"
	{{ .service.Package }} "{{ .service.Package }}/generated"
	"log"
)

type Server struct {
	{{ .service.Package }}.AppServer
}

func (s *Server) SayHello(ctx context.Context, req *{{ .service.Package }}.HelloRequest) (*{{ .service.Package }}.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	name := req.GetName()
	if name == "" {
		name = {{ .service.Package }}.HelloRequest_world.String()
	}
	return &{{ .service.Package }}.HelloReply{Message: "Hello, " + name + "!"}, nil
}

func NewServer() {{ .service.Package }}.AppServer {
	return &Server{}
}
