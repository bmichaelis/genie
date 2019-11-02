package api

import "flag"

var (
	GrpcAddr = flag.String("{{ .service.EnvVar }}_GRPC_ADDR", ":{{ .service.GrpcPort }}", "gRPC address")
{{if .service.EnableHttp}}	HttpAddr = flag.String("{{ .service.EnvVar }}_HTTP_ADDR", ":{{ .service.HttpPort }}", "HTTP address"){{end}}
)
