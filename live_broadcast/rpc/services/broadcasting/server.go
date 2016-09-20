package broadcastings

import (
	pb "platform/commons/protos/broadcasting"
	"platform/utils"

	"golang.org/x/net/context"
)

type BroadcastingServer struct{}

func (b *BroadcastingServer) RecentBarrageList(ctx context.Context, room *pb.Room) (*pb.Barrages, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("broadcasting.RecentBarrageList error: %v", err)
		}
	}()

	return nil, nil
}

func (b *BroadcastingServer) CurrentAudienceNum(ctx context.Context, in *pb.Room) (*pb.Num, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("broadcasting.CurrentAudienceNum error: %v", err)
		}
	}()

	return nil, nil
}
