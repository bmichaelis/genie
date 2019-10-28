# HelloWorld gRPC Generation

## Getting Started

### Running from project

```shell script
go mod download
go run cmd/{{{.Package}}}/main.go
```

### Running as docker comtainer

```shell script
docker build -t {{{.Package}}} .
docker run {{{.Package}}}
```

### Running with docker-compose

```shell script
docker-compose up --build
```

### Testing service with gRPC client

Be sure the service is running

```shell script
go run test/main.go
```
{{{if .EnableHttp}}}
### Testing with curl

```shell script
curl http://localhost:{{{.HttpPort}}}/v1/grpcService/hello?name=world
```
{{{end}}}

### References

#### Go layout standards

* https://github.com/golang-standards/project-layout

#### gRPC 

* https://github.com/znly/docker-protobuf
* https://github.com/grpc-ecosystem/grpc-gateway
* https://medium.com/@lchenn/generate-grpc-and-protobuf-libraries-with-containers-c15ba4e4f3ad
* https://github.com/grpc/grpc-go/blob/master/grpcServices/{{{.Package}}}/greeter_client/main.go
* https://grpc.io/docs/quickstart/go/
* https://grpc.io/docs/tutorials/basic/go/
* https://github.com/gogo/protobuf

#### Heath Check (for Kubernetes)

* https://developpaper.com/k8s-and-health-examination-best-practice-of-grpc-service-health-examination/
* https://github.com/grpc/grpc/blob/master/doc/health-checking.md
* https://github.com/grpc-ecosystem/grpc-health-probe/
* https://github.com/grpc/grpc-go/tree/master/health
