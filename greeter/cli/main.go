package main

import (
	"context"
	"fmt"
	_ "github.com/go-micro/examples/greeter/srv/proto/hello"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go-micro.dev/v4/registry"
	"log"

	//hello "github.com/go-micro/examples/greeter/srv/proto/hello"
	"go-micro.dev/v4"
	pb "github.com/Henry-jk/go-microservice-study/greeter/srv/proto"
	wrapperTrace "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
)

func main() {

	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
		},
	}

	closer, err := cfg.InitGlobalTracer(
		"my.micro.cli.greeter",
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	// 初始化 Consul 注册中心
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"), // 替换为您的 Consul 地址
	)

	// create a new service
	service := micro.NewService(
		micro.Name("my.micro.cli.greeter"),
		micro.Registry(consulReg), // 使用 Consul 注册中心
		micro.WrapClient(wrapperTrace.NewClientWrapper(opentracing.GlobalTracer())),
		)

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