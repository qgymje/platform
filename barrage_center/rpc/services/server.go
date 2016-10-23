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

func srvBarrageToPbBarrage(b *barrages.Barrage) *pb.Content {
	return &pb.Content{
		TypeID:      int32(b.TypeID),
		BroadcastID: b.BroadcastID,
		UserID:      b.UserID,
		Text:        b.Text,
		CreatedAt:   b.CreatedAt.Unix(),
		Username:    b.Username,
		Level:       b.Level,
	}
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
	pbContents := []*pb.Content{}
	for i := range list {
		pbContent := srvBarrageToPbBarrage(list[i])
		pbContents = append(pbContents, pbContent)
	}

	return &pb.Barrages{List: pbContents}, nil
}
