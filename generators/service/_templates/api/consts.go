package api

import "flag"

var (
	GrpcAddr = flag.String("{{ .service.PackageUpper }}_GRPC_ADDR", ":{{ .service.GrpcPort }}", "gRPC address")
{{if .service.EnableHttp}}	HttpAddr = flag.String("{{ .service.PackageUpper }}_HTTP_ADDR", ":{{ .service.HttpPort }}", "HTTP address"){{end}}
)
