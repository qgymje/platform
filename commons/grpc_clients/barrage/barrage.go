package barrageClient

import (
	"context"

	pb "platform/commons/protos/barrage"
	"platform/utils"

	"google.golang.org/grpc"
)

// Barrage grpc client barrage object
type Barrage struct {
	conn   *grpc.ClientConn
	client pb.BarrageClient
}

// NewBarrage new grpc client barrage object
func NewBarrage(address string) *Barrage {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("barrage grpc server cant not connect: %v", err)
	}

	b := new(Barrage)
	b.conn = conn
	b.client = pb.NewBarrageClient(b.conn)

	return b
}

// Close close barrage grpc client
func (b *Barrage) Close() error {
	return b.conn.Close()
}

// Send send a barrage
func (b *Barrage) Send(in *pb.Content) (*pb.Status, error) {
	defer b.Close()
	return b.client.Send(context.Background(), in)
}

// List recent list
func (b *Barrage) List(in *pb.Broadcast) (*pb.Barrages, error) {
	defer b.Close()
	return b.client.List(context.Background(), in)
}
