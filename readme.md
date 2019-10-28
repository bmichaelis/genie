# Genny                                                     

Generates scaffolding for gRPC and HTTP services for Go

## Features

Genny allows you to:

* Choose your namespace
* Choose the name of your package
* Choose to enable HTTP endpoints
* Choose gRPC and HTTP ports
* Genny can be easily modified for your project needs

The generated project contains the following:

* Protofile for gRPC / HTTP service generation
* Generator using Docker (no locally installed proto compiler required)
* Best practice project layout for Go
* gRPC client for service-to-service communication
* Example unit tests
* Ready to in local docker environment using docker compose

## Roadmap

* (Optional) Ready to deploy Helm charts for Kubernetes
* (Optional) Example CRUD using Mongo database


## Usage

To use, clone this repo and run the following command in your terminal

```shell script
go install
genie
```
# Contribute

Feel free to fork this project to contribute other features. Here are some things to remember...

1. KISS - Keep it simple
2. Features must be optional (CRUD using different database)
3. Features should be generic in use. 