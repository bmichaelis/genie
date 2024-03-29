package grpc

import (
	"{{ .service.Package }}/api"
	{{ .service.Package }} "{{ .service.Package }}/generated"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net"
)

type Serverer interface {
	Serve(server {{ .service.Package }}.AppServer) error
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

func (s *Server) Serve(server {{ .service.Package }}.AppServer) error {
	defer glog.Flush()

	lis, err := net.Listen("tcp", api.GrpcAddr)
	if err != nil {
		return err
	}

	{{ .service.Package }}.RegisterAppServer(s.Server(), server)

	return s.Server().Serve(lis)
}

func NewServer() Serverer {
	return &Server{}
}
