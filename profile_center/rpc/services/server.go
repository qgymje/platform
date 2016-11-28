package services

import (
	"errors"
	pb "platform/commons/protos/profile"
	"platform/profile_center/rpc/services/profiles"
	"platform/utils"
	"strconv"

	"golang.org/x/net/context"
)

// Server server
type Server struct{}

// Withdraw withdraw
func (s *Server) Withdraw(ctx context.Context, in *pb.Ammount) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("profiles.Server.Withdraw error: %+v", err)
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
func (s *Server) WithdrawRollback(ctx context.Context, in *pb.Message) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("profiles.Server.WithdrawRollback error: %+v", err)
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
func (s *Server) WithdrawCommit(ctx context.Context, in *pb.Message) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("profiles.Server.WithdrawCommit error: %+v", err)
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

// FriendRequest friend request
func (s *Server) FriendRequest(ctx context.Context, in *pb.Request) (*pb.RequestID, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("profiles.Server.WithdrawCommit error: %+v", err)
		}
	}()

	return &pb.RequestID{}, nil
}

// FriendAgree agree
func (s *Server) FriendAgree(ctx context.Context, in *pb.RequestID) (*pb.Status, error) {

	return &pb.Status{}, nil
}

// FriendRefuse refuse
func (s *Server) FriendRefuse(ctx context.Context, in *pb.RequestID) (*pb.Status, error) {

	return &pb.Status{}, nil
}
