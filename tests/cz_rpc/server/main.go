package main

import (
	"tech.cloudzen/libs/rpc"
	"github.com/micro/go-micro"
	proto "tech.cloudzen/tests/cz_rpc/server/proto"
	"golang.org/x/net/context"
	"fmt"
)

var cc int = 0

func main() {
	service := micro.NewService()
	service.Init(rpc.GenShardingOptions(&service, 150,
		micro.Name("cloudzen.rpc_test.server"),
		micro.Version("last"),
	)...)
	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	fmt.Println("Initialized RPC")
	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	cc++
	rsp.Greeting = "Hello " + req.Name
	rsp.Ltime = req.Ltime
	return nil
}

func (g *Greeter) Count(ctx context.Context, req *proto.NonParamRequest, rsp *proto.CountResponse) error {
	println(cc)
	rsp.Num = int64(cc)
	return nil
}

func (g *Greeter) Clear(ctx context.Context, req *proto.NonParamRequest, rsp *proto.NonParamResponse) error {
	println(cc)
	cc = 0
	return nil
}