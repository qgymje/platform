package broadcastings

import (
	"golang.org/x/net/context"
	pb "tech.cloudzen/protos/broadcasting"
	"tech.cloudzen/utils"
)

type BroadcastingServer struct{}

func (b *BroadcastingServer) RecentBarrageList(ctx context.Context, room *pb.Room) (*pb.Barrages, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("broadcasting.RecentBarrageList error: ", err.Error())
		}
	}()

	return nil, nil
}

func (b *BroadcastingServer) CurrentAudienceNum(ctx context.Context, in *pb.Room) (*pb.Num, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("broadcasting.CurrentAudienceNum error: ", err.Error())
		}
	}()

	return nil, nil
}
