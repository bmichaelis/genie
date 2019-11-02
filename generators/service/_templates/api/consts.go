package api

import "flag"

var (
	GrpcAddr = flag.String("{{.PackageUpper}}_GRPC_ADDR", ":{{.GrpcPort}}", "gRPC address")
{{if .EnableHttp}}	HttpAddr = flag.String("{{.PackageUpper}}_HTTP_ADDR", ":{{.HttpPort}}", "HTTP address"){{end}}
)
