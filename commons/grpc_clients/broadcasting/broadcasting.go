package broadcastingClient

import (
	"context"

	pb "platform/commons/protos/broadcasting"
	"platform/utils"

	"google.golang.org/grpc"
)

type Broadcasting struct {
	conn   *grpc.ClientConn
	client pb.BroadcastingClient
}

func NewBroadcasting(address string) *Broadcasting {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("broadcastingr grpc server cant not connect: %v", err)
	}

	b := new(Broadcasting)
	b.conn = conn
	b.client = pb.NewBroadcastingClient(b.conn)

	return b
}

func (b *Broadcasting) Close() error {
	return b.conn.Close()
}

func (b *Broadcasting) RecentBarrageList(in *pb.Room) (*pb.Barrages, error) {
	defer b.Close()
	bpb, err := b.client.RecentBarrageList(context.Background(), in)
	return bpb, err
}

func (b *Broadcasting) CurrnetAudienceNum(in *pb.Room) (*pb.Num, error) {
	defer b.Close()
	bpb, err := b.client.CurrentAudienceNum(context.Background(), in)
	return bpb, err
}
