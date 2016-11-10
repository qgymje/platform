package services

import (
	"errors"

	"golang.org/x/net/context"

	"platform/broadcast_room/rpc/services/broadcast"
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
	room := &pb.RoomInfo{
		RoomID:    r.RoomID,
		Name:      r.Name,
		UserID:    r.UserID,
		UserName:  r.UserName,
		Cover:     r.Cover,
		IsPlaying: r.IsPlaying,
		FollowNum: r.FollowNum,
	}
	if room.IsPlaying && r.Broadcast != nil {
		bro := r.Broadcast
		broadcastInfo := &pb.BroadcastInfo{
			BroadcastID:     bro.BroadcastID,
			RoomID:          bro.RoomID,
			StartTime:       bro.StartTime.Unix(),
			TotalAudience:   bro.TotalAudience,
			CurrentAudience: bro.CurrentAudience,
			Duration:        bro.Duration,
		}
		room.Broadcast = broadcastInfo
	}

	return room
}

func srvBroadcastToPbBroadcast(bro *broadcasts.Broadcast) *pb.BroadcastInfo {
	b := &pb.BroadcastInfo{
		BroadcastID:     bro.BroadcastID,
		RoomID:          bro.RoomID,
		StartTime:       bro.StartTime.Unix(),
		CurrentAudience: bro.CurrentAudience,
		TotalAudience:   bro.TotalAudience,
		Duration:        bro.Duration,
	}
	return b
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
func (s *Server) Info(ctx context.Context, in *pb.UserRoom) (*pb.RoomInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Info error: %+v", err)
		}
	}()

	r := rooms.NewRoom()
	var info *rooms.Room
	if in.RoomID != "" {
		info, err = r.GetByID(in.RoomID)
	} else {
		info, err = r.GetByUserID(in.UserID)
	}
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

	starterConfig := &broadcasts.StarterConfig{
		UserID: in.UserID,
	}
	starter := broadcasts.NewStarter(starterConfig)
	if err := starter.Do(); err != nil {
		return nil, errors.New(starter.ErrorCode().String())
	}

	bro, _ := starter.GetBroadcast()
	return srvBroadcastToPbBroadcast(bro), nil
}

// End end broadcastring
func (s *Server) End(ctx context.Context, in *pb.User) (*pb.BroadcastInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.End error: %+v", err)
		}
	}()

	enderConfig := &broadcasts.EnderConfig{
		UserID: in.UserID,
		TypeID: int(in.TypeID),
	}
	ender := broadcasts.NewEnder(enderConfig)
	if err := ender.Do(); err != nil {
		return nil, errors.New(ender.ErrorCode().String())
	}

	broInfo, _ := ender.GetBroadcast()
	return srvBroadcastToPbBroadcast(broInfo), nil
}

// Enter when a user enter a broadcast
func (s *Server) Enter(ctx context.Context, in *pb.UserRoom) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Enter error: %+v", err)
		}
	}()

	enterConfig := &broadcasts.EnterConfig{
		BroadcastID: in.BroadcastID,
		UserID:      in.UserID,
		TypeID:      int(in.TypeID),
		Level:       in.Level,
		Username:    in.Username,
	}
	enter := broadcasts.NewEnter(enterConfig)
	if err := enter.Do(); err != nil {
		return nil, errors.New(enter.ErrorCode().String())
	}
	return &pb.Status{Success: true, BroadcastID: in.BroadcastID}, nil
}

// Leave when a user leave a broadcast
func (s *Server) Leave(ctx context.Context, in *pb.UserRoom) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Leave error: %+v", err)
		}
	}()

	leaverConfig := &broadcasts.LeaverConfig{
		BroadcastID: in.BroadcastID,
		UserID:      in.UserID,
		TypeID:      int(in.TypeID),
		Level:       in.Level,
		Username:    in.Username,
	}
	leaver := broadcasts.NewLeaver(leaverConfig)
	if err := leaver.Do(); err != nil {
		return nil, errors.New(leaver.ErrorCode().String())
	}
	return &pb.Status{Success: true, BroadcastID: in.BroadcastID}, nil
}

// Follow when a user follow a room
func (s *Server) Follow(ctx context.Context, in *pb.UserRoom) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Follow error: %+v", err)
		}
	}()

	config := &rooms.FollowConfig{
		UserID: in.UserID,
		RoomID: in.RoomID,
	}
	follow := rooms.NewFollow(config)
	if err = follow.Do(); err != nil {
		return nil, errors.New(follow.ErrorCode().String())
	}
	return &pb.Status{Success: true, RoomID: in.RoomID}, nil
}

// Unfollow when a user unfollow a room
func (s *Server) Unfollow(ctx context.Context, in *pb.UserRoom) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.rooms.Unfollow error: %+v", err)
		}
	}()

	config := &rooms.FollowConfig{
		UserID: in.UserID,
		RoomID: in.RoomID,
	}
	follow := rooms.NewFollow(config)
	if err = follow.Undo(); err != nil {
		return nil, errors.New(follow.ErrorCode().String())
	}
	return &pb.Status{Success: true, RoomID: in.RoomID}, nil
}
