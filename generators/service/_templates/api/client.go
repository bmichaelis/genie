package api

import (
	{{ .service.Package }} "{{ .service.Package }}/generated"
	"{{ .service.Package }}/tools/env"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

var (
	GrpcAddr = env.String("{{ .service.EnvVar }}_GRPC_ADDR", ":8080", "gRPC address")
	HttpAddr = env.String("{{ .service.EnvVar }}_HTTP_ADDR", ":3000", "HTTP address")
)

func Connect() ({{ .service.Package }}.AppClient, error) {
	var err error
	conn, err = grpc.Dial(GrpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := {{ .service.Package }}.NewAppClient(conn)
	return client, nil
}

func Disconnect() error {
	if conn != nil {
		if err := conn.Close(); err != nil {
			return err
		}
	}
	return nil
}
