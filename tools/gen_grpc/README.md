# gen_grpc

## sketch 
auto generate grpc client code and server code based on protobuf.

## usage

```
gen_grpc -proto=./protos/user.proto -type=client -dest=../grpc_clients/user/user.go
gen_grpc -proto=./protos/user.proto -type=server -dest=../../broadcast_room/rpc/services/server.go
```
