package roomClient

import (
	"context"

	pb "platform/commons/protos/room"
	"platform/utils"

	"google.golang.org/grpc"
)

// Room room grpc client
type Room struct {
	conn   *grpc.ClientConn
	client pb.RoomClient
}

// NewRoom create a new grpc room client
func NewRoom(address string) *Room {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("room grpc server cant not connect: %v", err)
	}

	r := new(Room)
	r.conn = conn
	r.client = pb.NewRoomClient(r.conn)

	return r
}

// Close close grpc client
func (r *Room) Close() error {
	return r.conn.Close()
}

// Create create a room
func (r *Room) Create(in *pb.CreateRequest) (*pb.Status, error) {
	defer r.Close()
	return r.client.Create(context.Background(), in)
}

// List list the rooms
func (r *Room) List(in *pb.ListRequest) (*pb.Rooms, error) {
	defer r.Close()
	return r.client.List(context.Background(), in)
}

// Info user's room info
func (r *Room) Info(in *pb.UserRoom) (*pb.RoomInfo, error) {
	defer r.Close()
	return r.client.Info(context.Background(), in)
}

// Start broadcast
func (r *Room) Start(in *pb.User) (*pb.BroadcastInfo, error) {
	defer r.Close()
	rpb, err := r.client.Start(context.Background(), in)
	return rpb, err
}

// End end broadcast
func (r *Room) End(in *pb.User) (*pb.BroadcastInfo, error) {
	defer r.Close()
	return r.client.End(context.Background(), in)
}

// Enter enter the broadcast
func (r *Room) Enter(in *pb.UserRoom) (*pb.Status, error) {
	defer r.Close()
	return r.client.Enter(context.Background(), in)
}

// Leave leave the broadcast
func (r *Room) Leave(in *pb.UserRoom) (*pb.Status, error) {
	defer r.Close()
	return r.client.Leave(context.Background(), in)
}

// Follow follow action
func (r *Room) Follow(in *pb.UserRoom) (*pb.Status, error) {
	defer r.Close()
	return r.client.Follow(context.Background(), in)
}

// Unfollow unfollow action
func (r *Room) Unfollow(in *pb.UserRoom) (*pb.Status, error) {
	defer r.Close()
	return r.client.Unfollow(context.Background(), in)
}
