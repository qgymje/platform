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
			utils.GetLog().Error("servers.FriendServer.FriendRequest error: %+v", err)
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
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.FriendServer.FriendList error: %+v", err)
		}
	}()
	config := &friends.Config{
		UserID: in.UserID,
	}
	fri := friends.NewFriends(config)
	if err = fri.Do(); err != nil {
		return nil, errors.New(fri.ErrorCode().String())
	}

	var pbFriends []string
	friendIDs := fri.Result()
	for _, f := range friendIDs {
		pbFriends = append(pbFriends, f.UserID)
	}

	return &pb.Friends{FriendIDs: pbFriends}, nil
}
