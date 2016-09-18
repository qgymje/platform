package roomClient

import (
	"context"

	pb "platform/commons/protos/room"
	"platform/utils"

	"google.golang.org/grpc"
)

type Room struct {
	conn   *grpc.ClientConn
	client pb.RoomClient
}

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

func (r *Room) Close() error {
	return r.conn.Close()
}

func (r *Room) Start(in *pb.User) (*pb.Status, error) {
	defer r.Close()
	rpb, err := r.client.Start(context.Background(), in)
	return rpb, err
}

func (r *Room) Create(in *pb.RoomRequest) (*pb.RoomResponse, error) {
	defer r.Close()
	rpb, err := r.client.Create(context.Background(), in)
	return rpb, err
}

func (r *Room) End(in *pb.User) (*pb.EndResponse, error) {
	defer r.Close()
	rpb, err := r.client.End(context.Background(), in)
	return rpb, err
}
