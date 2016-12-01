package servers

import (
	"errors"
	pb "platform/commons/protos/profile"
	"platform/profile_center/rpc/services/profiles"
	"platform/utils"
	"strconv"

	"golang.org/x/net/context"
)

// ProfileServer server
type ProfileServer struct{}

// Withdraw withdraw
func (s *ProfileServer) Withdraw(ctx context.Context, in *pb.Ammount) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.ProfileServer.Withdraw error: %+v", err)
		}
	}()

	config := &profiles.WithdrawConfig{
		UserID:    in.UserID,
		SnowBall:  uint(in.SnowBall),
		SnowFlake: uint(in.SnowFlake),
		TypeID:    uint(in.TypeID),
		TargetID:  in.TargetID,
	}
	wd := profiles.NewWithdraw(config)
	msgID, err := wd.Create()
	if err != nil {
		return nil, errors.New(wd.ErrorCode().String())
	}
	smsgID := strconv.FormatInt(msgID, 10)
	return &pb.Status{Success: true, MsgID: smsgID}, nil
}

// WithdrawRollback withdraw rollback
func (s *ProfileServer) WithdrawRollback(ctx context.Context, in *pb.Message) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.ProfileServer.WithdrawRollback error: %+v", err)
		}
	}()

	config := &profiles.WithdrawConfig{
		UserID: in.UserID,
		MsgID:  in.MsgID,
	}
	wd := profiles.NewWithdraw(config)
	if err = wd.Rollback(); err != nil {
		return nil, errors.New(wd.ErrorCode().String())
	}

	return &pb.Status{Success: true, MsgID: in.MsgID}, nil
}

// WithdrawCommit withdraw commit
func (s *ProfileServer) WithdrawCommit(ctx context.Context, in *pb.Message) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.ProfileServer.WithdrawCommit error: %+v", err)
		}
	}()

	config := &profiles.WithdrawConfig{
		UserID: in.UserID,
		MsgID:  in.MsgID,
	}
	wd := profiles.NewWithdraw(config)
	if err = wd.Commit(); err != nil {
		return nil, errors.New(wd.ErrorCode().String())
	}

	return &pb.Status{Success: true, MsgID: in.MsgID}, nil
}
