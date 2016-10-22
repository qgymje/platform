package services

import (
	"errors"
	"platform/barrage_center/rpc/services/barrages"
	pb "platform/commons/protos/barrage"
	"platform/utils"

	"golang.org/x/net/context"
)

// Server barrage server
type Server struct {
}

func srvBarrageToBarrage(b *barrages.Barrage) *pb.Content {
	return &pb.Content{}
}

// Send send a barrage
func (s *Server) Send(ctx context.Context, in *pb.Content) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.barrages.Send error: %+v", err)
		}
	}()

	config := &barrages.SenderConfig{
		TypeID:      in.TypeID,
		BroadcastID: in.BroadcastID,
		UserID:      in.UserID,
		Username:    in.Username,
		Level:       in.Level,
		Text:        in.Text,
		CreatedAt:   in.CreatedAt,
	}
	sender := barrages.NewSender(config)
	if err = sender.Do(); err != nil {
		return nil, errors.New(sender.ErrorCode().String())
	}

	return &pb.Status{Success: true}, nil
}

// List current barrage list
func (s *Server) List(ctx context.Context, in *pb.Broadcast) (*pb.Barrages, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.barrages.Listerror: %+v", err)
		}
	}()
	utils.Dump(in)
	config := &barrages.Config{
		BroadcastID: in.BroadcastID,
		StartTime:   in.StartTime,
		EndTime:     in.EndTime,
		PageNum:     int(in.Num),
		PageSize:    int(in.Size),
	}

	barrageList := barrages.NewBarrages(config)
	if err = barrageList.Do(); err != nil {
		return nil, errors.New(barrageList.ErrorCode().String())
	}

	list := barrageList.Barrages()
	utils.Dump(list)

	return nil, nil
}
