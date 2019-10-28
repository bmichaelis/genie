package healthcheck

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Serverer interface {
	Serve(grpcServer *grpc.Server) error
}

type Server struct {
}

func (*Server) Serve(grpcServer *grpc.Server) error {
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	return nil
}

func NewServer() Serverer {
	return &Server{}
}
