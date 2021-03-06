package gameClient

import (
	"golang.org/x/net/context"

	pb "platform/commons/protos/game"
	"platform/utils"

	"google.golang.org/grpc"
)

// Game grpc game client
type Game struct {
	conn   *grpc.ClientConn
	client pb.GameClient
}

// NewGame new grpc game client object
func NewGame(address string) *Game {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("game grpc server cant not connect: %v", err)
	}

	g := new(Game)
	g.conn = conn
	g.client = pb.NewGameClient(g.conn)

	return g
}

// Close close the connection
func (g *Game) Close() error {
	return g.conn.Close()
}

// Create create a game
func (g *Game) Create(in *pb.GameInfo) (*pb.Status, error) {
	defer g.Close()
	return g.client.Create(context.Background(), in)
}

// Start start a game
func (g *Game) Start(in *pb.UserGame) (*pb.GameVM, error) {
	defer g.Close()
	return g.client.Start(context.Background(), in)
}

// List list available games
func (g *Game) List(in *pb.Page) (*pb.Games, error) {
	defer g.Close()
	return g.client.List(context.Background(), in)
}

// Preference fetch a game preference
func (g *Game) Preference(in *pb.UserGame) (*pb.PreferenceConfig, error) {
	defer g.Close()
	return g.client.Preference(context.Background(), in)
}

// UpdatePreference update a game preference
func (g *Game) UpdatePreference(in *pb.UserGame) (*pb.Status, error) {
	defer g.Close()
	return g.client.UpdatePreference(context.Background(), in)
}
