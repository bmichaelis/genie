package main

import (
	"fmt"
	"github.com/golang/glog"
	"{{ .service.Package }}/api"
	service "{{ .service.Package }}/generated"
)

func main() {
	client := api.NewClient()
	if err := client.Connect(); err != nil {
		glog.Fatalf("Error connecting: %s", err.Error())
	}

	msg, err := client.SayHello(service.HelloRequest_world.String())
	if err != nil {
		glog.Fatalf("Error calling SayHello: %s", err.Error())
	}

	fmt.Println("Greeting: ", msg)
}
