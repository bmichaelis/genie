package main

import (
	"context"
	"fmt"
	"{{ .service.Package }}/api"
	{{ .service.Package }} "{{ .service.Package }}/generated"
)

func main() {
	client, err := api.Connect()
	if err != nil {
		panic(err)
	}

	result, err := client.SayHello(context.Background(), &{{ .service.Package }}.HelloRequest{
		Name: {{ .service.Package }}.HelloRequest_world.String(),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Greeting: ", result.Message)
}
