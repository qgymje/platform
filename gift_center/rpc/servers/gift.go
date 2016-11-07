package servers

import (
	"errors"
	pb "platform/commons/protos/gift"
	"platform/gift_center/rpc/services/gifts"
	"platform/utils"

	"golang.org/x/net/context"
)

// Gift server
type Gift struct {
}

func srvGiftToPbGift(s *gifts.Gift) *pb.GiftInfo {
	return &pb.GiftInfo{
		GiftID:    s.GiftID,
		Name:      s.Name,
		Image:     s.Image,
		SnowBall:  uint32(s.SnowBall),
		SnowFlake: uint32(s.SnowFlake),
		Combo:     int32(s.Combo),
	}
}

// List gift list
func (g *Gift) List(ctx context.Context, in *pb.Page) (*pb.Gifts, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.Gift.List error: %+v", err)
		}
	}()

	gs := gifts.NewGifts()
	if err = gs.Do(); err != nil {
		return nil, errors.New(gs.ErrorCode().String())
	}

	pbGifts := []*pb.GiftInfo{}
	srvGifts := gs.Gifts()
	for i := range srvGifts {
		pbGift := srvGiftToPbGift(srvGifts[i])
		pbGifts = append(pbGifts, pbGift)
	}

	return &pb.Gifts{List: pbGifts}, nil
}

// Send send a gift
func (g *Gift) Send(ctx context.Context, in *pb.SendGift) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.Gift.Send error: %+v", err)
		}
	}()

	return nil, nil
}
