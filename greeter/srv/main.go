// Package main
package main

import (
	"context"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"log"
	"time"
	//hello "github.com/go-micro/examples/greeter/srv/proto/hello"
	"go-micro.dev/v4"
	pb "github.com/Henry-jk/go-microservice-study/greeter/srv/proto"

)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	//log.Log("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {

	// 初始化 Consul 注册中心
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"), // 替换为您的 Consul 地址
	)

	service := micro.NewService(
		micro.Name("my.micro.srv.greeter"),
		micro.Address("192.168.0.107:9090"),
		micro.Registry(consulReg), // 使用 Consul 注册中心
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	pb.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}