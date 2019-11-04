module {{ .service.Package }}

go 1.13

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/fatih/color v1.7.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.1.2
	golang.org/x/net v0.0.0-20191011234655-491137f69257 // indirect
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.24.0
)
