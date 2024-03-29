module {{ .service.Package }}

go 1.13

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/grpc-ecosystem/grpc-health-probe v0.3.0 // indirect
	golang.org/x/net v0.0.0-20191011234655-491137f69257
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.24.0
)
