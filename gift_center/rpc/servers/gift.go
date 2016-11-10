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

// Info gift info
func (g *Gift) Info(ctx context.Context, in *pb.GiftID) (*pb.GiftInfo, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.Gift.Info error: %+v", err)
		}
	}()

	gs := gifts.NewGifts().SetGiftID(in.GiftID)
	if err = gs.Do(); err != nil {
		return nil, errors.New(gs.ErrorCode().String())
	}

	pbGift := srvGiftToPbGift(gs.One())
	return pbGift, nil
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
	srvGifts := gs.Result()
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

	senderConfig := &gifts.SenderConfig{
		UserID:      in.UserID,
		GiftID:      in.GiftID,
		ToUserID:    in.ToUserID,
		MsgID:       in.MsgID,
		BroadcastID: in.BroadcastID,
	}

	sender := gifts.NewSender(senderConfig)
	if err = sender.Do(); err != nil {
		return nil, errors.New(sender.ErrorCode().String())
	}
	return &pb.Status{Success: true, MsgID: in.MsgID}, nil
}

// Broadcast broadcast by send gift
func (g *Gift) Broadcast(ctx context.Context, in *pb.SendGift) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.Gift.Broadcast error: %+v", err)
		}
	}()
	utils.Dump(in)
	return nil, nil

}
