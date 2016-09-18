package main

import (
	"tech.cloudzen/libs/rpc"
	"github.com/micro/go-micro"
	"strconv"
	"tech.cloudzen/tests/cz_rpc/server/proto"
	"time"
	"fmt"
	//"math"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService()
	service.Init(
		micro.Name("cloudzen.rpc_test.client"),
		micro.Version("last"),
	)
	services, _ := registry.GetService("cloudzen.rpc_test.server")
	clearClients(services)
	for i := int64(0); i < 5000; i ++ {
		var res greeter.HelloResponse
		stime := time.Now().UnixNano()
		rpc.CallForID(
			strconv.Itoa(int(i)),
			"cloudzen.rpc_test.server",
			"Greeter.Hello",
			&greeter.HelloRequest{
				Name: "John",
				Ltime: stime,
			},
			&res,
		)
		tnow := time.Now().UnixNano()
		fmt.Println(res.Greeting, stime, "-", tnow, (tnow - stime) / 1000, "ns")
	}
	printDistributions(services)
}


func clearClients (services []*registry.Service) {
	for _, svr := range services {
		for _, node := range svr.Nodes {
			cr := greeter.NonParamResponse{}
			ccerr := client.CallRemote(
				context.TODO(),
				node.Address + ":" + strconv.Itoa(node.Port),
				client.NewRequest(
					"cloudzen.rpc_test.server",
					"Greeter.Clear",
					&greeter.NonParamRequest{},
				),
				&cr,
			)
			if (ccerr != nil) {
				panic(ccerr)
			}
		}
	}
}

func printDistributions (services []*registry.Service) {
	for _, svr := range services {
		for _, node := range svr.Nodes {
			cr := greeter.CountResponse{}
			ccerr := client.CallRemote(
				context.TODO(),
				node.Address + ":" + strconv.Itoa(node.Port),
				client.NewRequest(
					"cloudzen.rpc_test.server",
					"Greeter.Count",
					&greeter.NonParamRequest{},
				),
				&cr,
			)
			if (ccerr != nil) {
				panic(ccerr)
			} else {
				fmt.Println(node.Address, cr.Num)
			}
		}
	}
}
