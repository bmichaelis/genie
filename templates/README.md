# HelloWorld gRPC Generation

#### Running from project

```shell script
go mod download
go run cmd/{{.Package}}/main.go
```

#### Running as docker comtainer

```shell script
docker build -t {{.Package}} .
docker run {{.Package}}
```

#### Running with docker-compose

```shell script
docker-compose up --build
```

#### Testing service with gRPC client

Be sure the service is running

```shell script
go run test/main.go
```

#### Testing with curl

```shell script
curl http://localhost:21000/v1/grpcService/hello?name=world
```

## Go layout

* https://github.com/golang-standards/project-layout

### References

* https://github.com/znly/docker-protobuf
* https://github.com/grpc-ecosystem/grpc-gateway
* https://medium.com/@lchenn/generate-grpc-and-protobuf-libraries-with-containers-c15ba4e4f3ad
* https://github.com/grpc/grpc-go/blob/master/grpcServices/{{.Package}}/greeter_client/main.go
* https://grpc.io/docs/quickstart/go/
* https://grpc.io/docs/tutorials/basic/go/
* https://github.com/gogo/protobuf

### Heath Check (for Kubernetes)

* https://developpaper.com/k8s-and-health-examination-best-practice-of-grpc-service-health-examination/
* https://github.com/grpc/grpc/blob/master/doc/health-checking.md
* https://github.com/grpc-ecosystem/grpc-health-probe/
* https://github.com/grpc/grpc-go/tree/master/health

### Mongo

* https://github.com/amsokol/mongo-go-driver-protobuf
* https://gist.github.com/amsokol/d6f3495375c7a78e4d02fc76d1b0d3c4# {{.Package}}-grpc
