package games

import (
	"errors"

	"golang.org/x/net/context"

	pb "platform/commons/protos/game"
	"platform/utils"
)

// GameServer implement the gRPC Game protocol
type GameServer struct {
}

// Create create a game
func (g *GameServer) Create(ctx context.Context, in *pb.GameInfo) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.Create error: %v", err)
		}
	}()

	config := CreaterConfig{
		CompanyID:   in.CompanyID,
		Name:        in.Name,
		GameTypeID:  int(in.GameTypeID),
		Cover:       in.Cover,
		Screenshots: in.Screenshots,
		Description: in.Description,
		PlayerNum:   in.PlayerNum,
		IsFree:      in.IsFree,
		Charge:      in.Charge,
	}

	creater := NewCreater(&config)
	if err = creater.Do(); err != nil {
		return nil, errors.New(creater.ErrorCode().String())
	}
	return &pb.Status{Success: true}, nil
}

// Start receive a user_id and game_id to request a game vms through protobuf protocol
func (g *GameServer) Start(ctx context.Context, in *pb.UserGame) (*pb.GameVM, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.Start error: %v", err)
		}
	}()
	return &pb.GameVM{}, nil
}

// End end the game
func (g *GameServer) End(ctx context.Context, in *pb.UserGame) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.End error: %v", err)
		}
	}()
	return &pb.Status{Success: true}, nil
}

func srvGameToPbGame(g *Game) *pb.GameInfo {
	return &pb.GameInfo{
		GameID:       g.GameID,
		Name:         g.Name,
		GameTypeID:   int32(g.GameTypeID),
		GameTypeName: g.GameTypeName,
		Description:  g.Description,
		Cover:        g.Cover,
		Screenshots:  g.Screenshots,
		PlayTimes:    g.PlayTimes,
		PlayerNum:    g.PlayerNum,
		IsFree:       g.IsFree,
		PayStatus:    g.PayStatus,
	}

}

// List all available games
func (g *GameServer) List(ctx context.Context, in *pb.Page) (*pb.Games, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.List error: %v", err)
		}
	}()

	config := &Config{
		GameTypeID: int(in.GameTypeID),
		PageNum:    int(in.Num),
		PageSize:   int(in.Size),
		Search:     in.Search,
	}
	games := NewGames(config)
	if err = games.Do(); err != nil {
		return nil, errors.New(games.ErrorCode().String())
	}

	srvGames := games.Games()
	count := games.Count()

	var pbGames []*pb.GameInfo
	for _, srvGame := range srvGames {
		pbGame := srvGameToPbGame(srvGame)
		pbGames = append(pbGames, pbGame)
	}
	return &pb.Games{Games: pbGames, TotalNum: count}, nil
}

// Preference fetch a user's preference of a game
func (g *GameServer) Preference(ctx context.Context, in *pb.UserGame) (*pb.PreferenceConfig, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.Preference error: %v", err)
		}
	}()
	return &pb.PreferenceConfig{Json: ""}, nil
}

// UpdatePreference update a preference of a game for one user
func (g *GameServer) UpdatePreference(ctx context.Context, in *pb.UserGame) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.games.UpdatePreference error: %v", err)
		}
	}()
	return &pb.Status{Success: true}, nil
}
