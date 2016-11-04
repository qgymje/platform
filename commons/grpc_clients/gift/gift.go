package giftClient

import (
	"golang.org/x/net/context"

	pb "platform/commons/protos/gift"
	"platform/utils"

	"google.golang.org/grpc"
)

// Gift grpc gift client
type Gift struct {
	conn   *grpc.ClientConn
	client pb.GiftClient
}

// NewGift new grpc gift client object
func NewGift(address string) *Gift {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("gift grpc server cant not connect: %v", err)
	}

	g := new(Gift)
	g.conn = conn
	g.client = pb.NewGiftClient(g.conn)

	return g
}

// Close close the connection
func (g *Gift) Close() error {
	return g.conn.Close()
}

// List gift list
func (g *Gift) List(in *pb.Page) (*pb.Status, error) {
	defer g.Close()
	return g.client.List(context.Background(), in)
}

// Send gift list
func (g *Gift) Send(in *pb.SendGift) (*pb.Status, error) {
	defer g.Close()
	return g.client.Send(context.Background(), in)
}
