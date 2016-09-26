package games

import (
	"golang.org/x/net/context"

	pb "platform/commons/protos/game"
	"platform/utils"
)

var _ pb.GameServer = (*GameServer)(nil)

// GameServer implement the gRPC Game protocol
type GameServer struct {
}

// Start receive a user_id and game_id to request a game vms through protobuf protocol
func (g *GameServer) Start(context.Context, *pb.UserGame) (*pb.GameVM, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.Start error: %v", err)
		}
	}()
	return &pb.GameVM{}, nil
}

// List all available games
func (g *GameServer) List(context.Context, *pb.Page) (*pb.Games, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.List error: %v", err)
		}
	}()
	return &pb.Games{Games: []*pb.GameInfo{}}, nil
}

// Preference fetch a user's preference of a game
func (g *GameServer) Preference(context.Context, *pb.UserGame) (*pb.PreferenceConfig, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.Preference error: %v", err)
		}
	}()
	return &pb.PreferenceConfig{Json: ""}, nil
}

// UpdatePreference update a preference of a game for one user
func (g *GameServer) UpdatePreference(context.Context, *pb.UserGame) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.UpdatePreference error: %v", err)
		}
	}()
	return &pb.Status{Success: true}, nil
}
