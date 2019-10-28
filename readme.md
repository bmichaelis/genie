# Genie                                                     

Genie is an extensible service generator used to create gRPC microservices for GO

## Features

Genie allows you to:

* Choose your namespace
* Choose the name of your package
* Choose to enable HTTP annotations for endpoint support
* Choose gRPC and HTTP ports
* Easily add new generators

The generated project contains the following:

* Protofile for gRPC / HTTP service generation
* Generator using Docker (no locally installed proto compiler required)
* Best practice project layout for Go
* gRPC client for service-to-service communication
* Swagger file for HTTP endpoints
* Example unit tests
* Ready to in local docker environment using docker compose
* (Optional) Helm charts for Kubernetes

## Roadmap

* (Optional) Example CRUD using Mongo database

## Usage

To use, clone this repo and run the following command in your terminal

```shell script
go install
genie
```
## Contribute

Feel free to fork this project to contribute other generators.