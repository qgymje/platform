package main

import (
	"context"
	"log"

	pb "platform/tests/gamevm_mock_server/grpc_client/gamevm"

	"google.golang.org/grpc"
)

// GameVM game vm object
type GameVM struct {
	conn   *grpc.ClientConn
	client pb.HeartBeatClient
}

// NewGameVM new grpc client sms object
func NewGameVM(address string) *GameVM {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("game vm grpc server cant not connect: %v", err)
	} else {
		log.Println("game vm server connect ok")
	}

	g := new(GameVM)
	g.conn = conn
	g.client = pb.NewHeartBeatClient(g.conn)

	return g
}

// Close close sms grpc client
func (g *GameVM) Close() error {
	return g.conn.Close()
}

// NotifyAlive Email object verify
func (g *GameVM) NotifyAlive(in *pb.AgentRequest) (*pb.CzsReply, error) {
	defer g.Close()
	return g.client.NotifyAlive(context.Background(), in)
}

const address = "10.211.55.12:6060"

func main() {
	client := NewGameVM("192.168.0.104:6060")
	req := pb.AgentRequest{
		Key: "agent",
	}
	reply, err := client.NotifyAlive(&req)
	if err != nil {
		log.Println(err)
	}
	log.Println(reply)
}
