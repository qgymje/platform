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
		Name:     in.Name,
		UserName: in.UserName,
		UserID:   in.UserID,
		Cover:    in.Cover,
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

func srvRoomToPbRoom(r *rooms.Room) *pb.RoomInfo {
	return &pb.RoomInfo{
		RoomID:    r.RoomID,
		Name:      r.Name,
		UserName:  r.UserName,
		Cover:     r.Cover,
		IsPlaying: r.IsPlaying,
		FollowNum: r.FollowNum,
	}
}

// List the rooms
func (s *Server) List(ctx context.Context, in *pb.ListRequest) (*pb.Rooms, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.List error: %+v", err)
		}
	}()

	config := &rooms.Config{
		PageNum:  int(in.Num),
		PageSize: int(in.Size),
		Search:   in.Search,
	}
	srvRooms := rooms.NewRooms(config)
	if err = srvRooms.Do(); err != nil {
		return nil, errors.New(srvRooms.ErrorCode().String())
	}

	srvRoomList := srvRooms.Rooms()
	count := srvRooms.Count()

	var pbRooms []*pb.RoomInfo
	for _, srvRoom := range srvRoomList {
		pbRoom := srvRoomToPbRoom(srvRoom)
		pbRooms = append(pbRooms, pbRoom)
	}

	return &pb.Rooms{Rooms: pbRooms, TotalNum: count}, nil
}

// Info user's room
func (s *Server) Info(ctx context.Context, in *pb.User) (*pb.RoomInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Info error: %+v", err)
		}
	}()

	r := rooms.NewRoom()
	info, err := r.GetByUserID(in.UserID)
	if err != nil {
		return nil, errors.New(r.ErrorCode().String())
	}

	pbRoom := srvRoomToPbRoom(info)
	return pbRoom, nil
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

// Enter when a user enter a broadcast
func (s *Server) Enter(ctx context.Context, in *pb.UserRoom) (*pb.Status, error) {
	return nil, nil
}

// Leave when a user leave a broadcast
func (s *Server) Leave(ctx context.Context, in *pb.UserRoom) (*pb.Status, error) {
	return nil, nil
}

// Follow when a user follow a room
func (s *Server) Follow(ctx context.Context, in *pb.UserRoom) (*pb.Status, error) {
	return nil, nil
}

// Unfollow when a user unfollow a room
func (s *Server) Unfollow(ctx context.Context, in *pb.UserRoom) (*pb.Status, error) {
	return nil, nil
}
