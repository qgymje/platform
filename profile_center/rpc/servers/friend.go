package servers

import (
	"errors"
	pb "platform/commons/protos/profile"
	"platform/profile_center/rpc/services/friends"
	"platform/utils"

	"golang.org/x/net/context"
)

// FriendServer friend server
type FriendServer struct{}

// FriendRequest friend request
func (s *FriendServer) FriendRequest(ctx context.Context, in *pb.Request) (*pb.RequestID, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("profiles.Server.WithdrawCommit error: %+v", err)
		}
	}()
	config := &friends.RequestConfig{
		FromUserID: in.FromUserID,
		ToUserID:   in.ToUserID,
		Message:    in.Message,
	}
	req := friends.NewRequest(config)
	if err = req.Do(); err != nil {
		return nil, errors.New(req.ErrorCode().String())
	}
	reqID := req.GetRequestID()
	return &pb.RequestID{RequestID: reqID}, nil
}

// FriendAgree agree
func (s *FriendServer) FriendAgree(ctx context.Context, in *pb.RequestID) (*pb.Status, error) {

	return &pb.Status{}, nil
}

// FriendRefuse refuse
func (s *FriendServer) FriendRefuse(ctx context.Context, in *pb.RequestID) (*pb.Status, error) {

	return &pb.Status{}, nil
}

// FriendList friends list
func (s *FriendServer) FriendList(ctx context.Context, in *pb.Message) (*pb.Friends, error) {
	return &pb.Friends{}, nil
}
