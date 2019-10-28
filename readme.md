# Genie                                                     

Genie is an extensible service generator used to create gRPC micro-services for Go

<img src="https://repository-images.githubusercontent.com/217768929/edad3580-f8ed-11e9-8730-333c308c5f3f" width="500">

## Features

Genie allows you to:

* Choose your namespace
* Choose the name of your package
* Choose to enable HTTP annotations for endpoint support
* Choose gRPC and HTTP ports
* Easily add new generators

The generated project contains the following:

* Proto file for gRPC / HTTP service generation
* Generator using Docker (no locally installed proto compiler required)
* Common project layout for Go
* gRPC client for service-to-service communication
* Swagger file for HTTP endpoints
* Ready to in local docker environment using docker compose
* (Optional) Helm charts for Kubernetes

## Roadmap

* Example unit tests
* (Optional) Example CRUD using Mongo database

## Usage

To use, clone this repo and run the following command in your terminal

```shell script
go install
genie
```
## Contribute

Feel free to fork this project to contribute other generators.