# Genie                                                     

Genie is an extensible service generator used to create gRPC micro-services for Go

<img src="https://repository-images.githubusercontent.com/217768929/edad3580-f8ed-11e9-8730-333c308c5f3f" width="500">

## Features

Genie provides the following:

* Simple setup for
    * gRPC service
    * Helm charts for Kubernetes deployment
    * Mongo database CRUD service
* Easily add new generators for your own custom needs
* File overrides
* Uses Go Templates

The generated project contains the following:

* Proto file for gRPC with optional HTTP service generation
* Generator using Docker (no locally installed proto compiler required)
* Common project layout for Go
* gRPC client for service-to-service communication
* Swagger file for HTTP endpoints
* Ready-to-use local docker environment using docker-compose
* (Optional) Helm charts for Kubernetes
* (Optional) Example CRUD using Mongo database

## Roadmap

* Example unit tests

## Usage

To use, clone this repo and run the following command in your terminal

```shell script
go install
genie
```
## Contribute

Feel free to fork this project to contribute other generators.

## Tools

* https://github.com/TheThingsIndustries/docker-protobuf
* https://github.com/envoyproxy/protoc-gen-validate
* https://github.com/gogo/protobuf
* https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/string-value
    * https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto

## References

* [Gogo Options](https://github.com/gogo/protobuf/blob/master/extensions.md)
* [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [Beating JSON performance with Protobuf](https://auth0.com/blog/beating-json-performance-with-protobuf/)
* [Part 1 -- How to develop Go gRPC microservice with HTTP/REST endpoint, middleware, Kubernetes deployment, etc.](https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-daebb36a97e9)
* [Part 2 -- How to develop Go gRPC microservice with HTTP/REST endpoint, middleware, Kubernetes deployment, etc.](https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-af1fff81aeb2)
* [gRPC Go Mongo Example](https://github.com/kyriediculous/go-grpc-mongodb)

## Examples

* https://github.com/gogo/grpc-example/blob/master/proto/example.proto
* https://github.com/dropbox/gogoprotobuf/blob/master/test/example/example.proto
* https://github.com/TheThingsIndustries/protoc-gen-fieldmask/blob/master/testdata/testdata.proto
