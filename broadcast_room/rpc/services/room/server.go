package rooms

import (
	"errors"

	"golang.org/x/net/context"

	pb "platform/commons/protos/room"
	"platform/utils"
)

type RoomServer struct {
}

func (r *RoomServer) Create(ctx context.Context, in *pb.RoomRequest) (*pb.RoomResponse, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Create error: %v", err)
		}
	}()

	config := RoomCreatorConfig{
		Name:       in.Name,
		UserID:     in.UserID,
		Channel:    in.Channel,
		SubChannel: in.SubChannel,
		Cover:      in.Cover,
	}
	rc := NewRoomCreator(&config)
	if err := rc.Do(); err != nil {
		return nil, errors.New(rc.ErrorCode().String())
	}

	response := pb.RoomResponse{
		RoomID:     rc.GetRoomID(),
		Name:       rc.GetName(),
		Channel:    rc.GetChannel(),
		SubChannel: rc.GetSubChannel(),
		Cover:      rc.GetCover(),
	}
	return &response, nil
}

func (r *RoomServer) Start(ctx context.Context, in *pb.User) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Start error: %v", err)
		}
	}()

	b := NewBroadcastStarter(in.UserID)
	if err := b.Start(); err != nil {
		return nil, errors.New(b.ErrorCode().String())
	}

	roomID, _ := b.GetRoomID()
	status := pb.Status{
		RoomID:  roomID,
		Success: true,
	}
	return &status, nil
}

func (r *RoomServer) End(ctx context.Context, in *pb.User) (*pb.EndResponse, error) {
	return nil, errors.New("not_implement")
}
