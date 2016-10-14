package services

import (
	"errors"

	"golang.org/x/net/context"

	"platform/broadcast_room/rpc/services/room"
	pb "platform/commons/protos/room"
	"platform/utils"
)

// Server room grpc server
type Server struct {
}

// Create create or update a broadcast room
func (s *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Create error: %+v", err)
		}
	}()

	config := rooms.CreatorConfig{
		Name:   in.Name,
		UserID: in.UserID,
		Cover:  in.Cover,
	}
	rc := rooms.NewCreator(&config)
	if err := rc.Do(); err != nil {
		return nil, errors.New(rc.ErrorCode().String())
	}

	status := pb.Status{
		RoomID:  rc.GetRoomID(),
		Success: true,
	}
	return &status, nil
}

// List the rooms
func (s *Server) List(ctx context.Context, in *pb.ListRequest) (*pb.Rooms, error) {
	return nil, errors.New("not_impl")
}

// Start start to broadcast
func (s *Server) Start(ctx context.Context, in *pb.User) (*pb.BroadcastInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Start error: %+v", err)
		}
	}()

	b := rooms.NewStarter(in.UserID)
	if err := b.Do(); err != nil {
		return nil, errors.New(b.ErrorCode().String())
	}

	broadcastID, _ := b.GetBroadcastID()
	info := pb.BroadcastInfo{
		BroadcastID: broadcastID,
	}
	return &info, nil
}

// End end broadcastring
func (s *Server) End(ctx context.Context, in *pb.User) (*pb.BroadcastInfo, error) {
	return nil, errors.New("not_implement")
}

// CurrentAudienceNum current audience number
func (s *Server) CurrentAudienceNum(ctx context.Context, in *pb.Broadcast) (*pb.Num, error) {
	return nil, errors.New("not_implement")
}
