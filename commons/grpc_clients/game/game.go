package gameClient

import (
	"context"

	pb "platform/commons/protos/game"
	"platform/utils"

	"google.golang.org/grpc"
)

type Game struct {
	conn   *grpc.ClientConn
	client pb.GameClient
}

func NewGame(address string) *Game {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("game grpc server cant not connect: %v", err)
	}

	g := new(Game)
	g.conn = conn
	g.client = pb.NewUserClient(g.conn)

	return g
}

func (g *Game) Close() error {
	return g.conn.Close()
}

func (g *Game) Start(ctx context.Context, in *pb.UserID) (*pb.GameVM, error) {
	defer u.Close()
	return g.client.Start(context.Background(), in)
}

func (g *Game) List(ctx context.Context, in *pb.Page) (*pb.Games, error) {
	defer u.Close()
	return g.client.List(context.Background(), in)
}
