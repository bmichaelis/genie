package http

import (
	"context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
	"{{ .service.Package }}/api"
	service "{{ .service.Package }}/generated"
)

type Serverer interface {
	Serve() error
}

type Server struct {
}

func (*Server) Serve() error {
	defer glog.Flush()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := service.Register{{ .service.Service }}HandlerFromEndpoint(ctx, mux, *api.GrpcAddr, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(*api.HttpAddr, mux)
}

func NewServer() Serverer {
	return &Server{}
}
