package main

import (
	"context"
	"fmt"
	_ "github.com/go-micro/examples/greeter/srv/proto/hello"

	//hello "github.com/go-micro/examples/greeter/srv/proto/hello"
	"go-micro.dev/v4"
	pb "github.com/Henry-jk/go-microservice-study/greeter/srv/proto"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := pb.NewSayService("my.micro.srv.greeter", service.Client())

	// Make request
	rsp, err := cl.Hello(context.Background(), &pb.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}