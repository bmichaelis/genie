package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"{{{.Package}}}/api"
	"{{{.Package}}}/internal"
	"{{{.Package}}}/internal/grpc"
	"{{{.Package}}}/internal/healthcheck"
	"{{{.Package}}}/internal/http"
)

func main() {
	shutdown := make(chan bool)

	flag.Parse()

	greeterServer := internal.NewServer()

	grpcServer := grpc.NewServer()
	go func() {
		fmt.Printf("Listening to gRPC on %s\n", *api.GrpcAddr)

		if err := grpcServer.Serve(greeterServer); err != nil {
			glog.Fatalln(err)
			<-shutdown
		}
	}()

	healthCheckServer := healthcheck.NewServer()
	go func() {
		fmt.Printf("Listening to health check\n")
		if err := healthCheckServer.Serve(grpcServer.Server()); err != nil {
			glog.Fatalln(err)
			<-shutdown
		}
	}()

	httpServer := http.NewServer()
	go func() {
		fmt.Printf("Listening to HTTP on %s\n", *api.HttpAddr)
		if err := httpServer.Serve(); err != nil {
			glog.Fatalln(err)
			<-shutdown
		}
	}()

	<-shutdown
}
