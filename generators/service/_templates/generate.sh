#!/bin/bash

mkdir -p ./generated
docker pull znly/protoc
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
  --go_out=plugins=grpc:generated \
  --grpc-gateway_out=logtostderr=true:generated \
  --swagger_out=logtostderr=true:generated \
  -Iprotos ./protos/*.proto

go fmt ./generated/...
