package grpc

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"{{{.Package}}}/api"
	grpcService "{{{.Package}}}/generated"
	"{{{.Package}}}/internal"
	"net"
)

type Serverer interface {
	Serve(server internal.Serverer) error
	Server() *grpc.Server
}

type Server struct {
	gs *grpc.Server
}

func (s *Server) Server() *grpc.Server {
	if s.gs == nil {
		s.gs = grpc.NewServer()
	}
	return s.gs
}

func (s *Server) Serve(server internal.Serverer) error {
	defer glog.Flush()

	lis, err := net.Listen("tcp", *api.GrpcAddr)
	if err != nil {
		return err
	}

	grpcService.RegisterGreeterServer(s.Server(), server)

	return s.Server().Serve(lis)
}

func NewServer() Serverer {
	return &Server{}
}
