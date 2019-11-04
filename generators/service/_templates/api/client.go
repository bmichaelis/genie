package api

import (
	{{ .service.Package }} "{{ .service.Package }}/generated"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

func Connect() ({{ .service.Package }}.AppClient, error) {
	var err error
	conn, err = grpc.Dial(*GrpcAddr, grpc.WithInsecure())
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
