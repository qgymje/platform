# CloudZen RPC

RPC middleware based on go-micro with sharding and automatic load-balancing features
 
## How to use

For complete and most recent example, please refer: https://github.com/micro/go-micro

### 1. Install code generator

`go get github.com/golang/protobuf`

or 

`go get github.com/micro/protobuf/{proto,protoc-gen-go}`

### 2. Create the Protocol Buffer

One of the key requirements of microservices is strongly defined interfaces so we utilised protobuf to define the handler and request/response. Here's a definition for the Greeter handler with the method Hello which takes a HelloRequest and HelloResponse both with one string arguments.

For more information on protocol buffer, please refer: https://developers.google.com/protocol-buffers/docs/proto

`tech.cloudzen/tests/cz_rpc/server/proto/greeter.proto`

    syntax = "proto3";
    
    service Greeter {
        rpc Hello(HelloRequest) returns (HelloResponse) {}
        rpc Count(NonParamRequest) returns (CountResponse) {}
        rpc Clear(NonParamRequest) returns (NonParamResponse) {}
    }
    
    message HelloRequest {
        string name = 1;
        int64 ltime = 3;
    }
    
    message HelloResponse {
        string greeting = 2;
        int64 ltime = 4;
    }
    
    message NonParamRequest {}
    message NonParamResponse {}
    
    message CountResponse {
        int64 num = 5;
    }
    
The proto file above also contains additional `Count` and `Clear` rpc with their message types for testing purpose later.
 
### 3. Compile the protobuf file

You need to recompile the proto file after every time you modified the file. 

`protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src $GOPATH/tech.cloudzen/tests/cz_rpc/server/proto/greeter.proto`

You shall *NOT* modify the *.pb.go file by hand

### 4. Define the service

Below is the code sample for the Greeter service. It basically implements the interface defined above for the Greeter handler, initialises the service, registers the handler and then runs itself. Simple as that.

`tech.cloudzen/tests/cz_rpc/server/main.go`

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

You may define how much resources (150 in this test) you have for this node in `GenShardingOptions` function. Please note, the number usually above 100 but not higher than 500. Higher number gives you fine grain sharding but lower performance in clients. If you have a number larger than this (for example, free memory in bytes), you need to reduce it. In contrast, you need to multiply the number if it is small, or resources may unable to sharding on nodes evenly.

### 5. Run Service

`go run tech.cloudzen/tests/cz_rpc/server/main.go`

    Initialized RPC
    2016/07/19 16:04:25 Listening on [::]:34880
    2016/07/19 16:04:25 Broker Listening on [::]:36430
    2016/07/19 16:04:25 Registering node: cloudzen.rpc_test.server-68ac793a-4d87-11e6-b816-2c56dcd3f678
    
### 6. Define a client

#### Traditional selector

Below is the client code to query the greeter service. Notice we're using the code generated client interface proto.NewGreeterClient. This reduces the amount of boiler plate code we need to write. The greeter client can be reused throughout the code if need be.

`client.go`

    package main
    
    import (
        "fmt"
        micro "github.com/micro/go-micro"
        proto "github.com/micro/go-micro/examples/service/proto"
        "golang.org/x/net/context"
    )
    
    
    func main() {
        // Create a new service. Optionally include some options here.
        service := micro.NewService(micro.Name("greeter.client"))
    
        // Create new greeter client
        greeter := proto.NewGreeterClient("greeter", service.Client())
    
        // Call the greeter
        rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
        if err != nil {
            fmt.Println(err)
        }
    
        // Print response
        fmt.Println(rsp.Greeting)
    }

#### Sticky selector

The RPC framework keeps the features from go-micro for round-robin and random selector. But some times, we need to stick some resources (users, for example) to one server until the server left the cluster. This is what `CallForID` for.

`client.go`

    package main
    
    import (
        "tech.cloudzen/libs/rpc"
        "github.com/micro/go-micro"
        "strconv"
        "tech.cloudzen/tests/cz_rpc/server/proto"
        "time"
        "fmt"
        "github.com/micro/go-micro/client"
        "golang.org/x/net/context"
    )
    
    func main() {
        service := micro.NewService()
        service.Init(
            micro.Name("cloudzen.rpc_test.client"),
            micro.Version("last"),
        )
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
    
In this version, `strconv.Itoa(int(i))` is the key you want to sharding. You need to convert other types of data into meaningful strings to ensure it works properly. You also need to provide service name and full method name as parameters. If you don't know what exactly that is, you can refer the *.pb.go file generated from *.proto file before.
The code is more complex than the traditional version, this is because the go-micro did some jobs for us in the *.pb.go. If you want the same convenience, you can write this by hand.

## Sticky selector benchmark

Performance did not optimized but response time is about 2000 ns for each request and requests is able to distributed to cluster nodes by id in almost even manner.

Here is the 4-node example

192.168.0.234 951

192.168.0.234 1032

192.168.0.234 957

192.168.0.234 1043

192.168.0.234 1017
