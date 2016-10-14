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

// Start broadcast
func (r *Room) Start(in *pb.User) (*pb.Status, error) {
	defer r.Close()
	rpb, err := r.client.Start(context.Background(), in)
	return rpb, err
}

// End end broadcast
func (r *Room) End(in *pb.User) (*pb.EndResponse, error) {
	defer r.Close()
	return r.client.End(context.Background(), in)
}

// CurrentAudienceNum current audience number
func (r *Room) CurrentAudienceNum(in *pb.Broadcast) (*pb.Num, error) {
	defer r.Close()
	return r.client.CurrentAudienceNum(context.Background(), in)
}
