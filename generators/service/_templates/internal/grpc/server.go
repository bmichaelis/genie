package grpc

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net"
	"{{ .service.Package }}/api"
	service "{{ .service.Package }}/generated"
)

type Serverer interface {
	Serve(server service.{{ .service.Service }}Server) error
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

func (s *Server) Serve(server service.{{ .service.Service }}Server) error {
	defer glog.Flush()

	lis, err := net.Listen("tcp", *api.GrpcAddr)
	if err != nil {
		return err
	}

	service.Register{{ .service.Service }}Server(s.Server(), server)

	return s.Server().Serve(lis)
}

func NewServer() Serverer {
	return &Server{}
}
