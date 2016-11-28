package servers

import (
	"errors"
	"platform/chat_center/rpc/services/chats"
	pb "platform/commons/protos/chat"
	"platform/utils"

	"golang.org/x/net/context"
)

// ChatServer is a chat grpc server
type ChatServer struct {
}

// List get the chat list
func (s *ChatServer) List(ctx context.Context, in *pb.Page) (*pb.ChatList, error) {
	return nil, nil
}

// Create create a chat
func (s *ChatServer) Create(ctx context.Context, in *pb.Creator) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.ChatServer.Create error: %+v", err)
		}
	}()

	config := &chats.CreatorConfig{
		Name:    in.Name,
		UserID:  in.UserID,
		Members: in.Members,
	}
	creator := chats.NewCreator(config)
	err = creator.Do()
	if err != nil && err != chats.ErrChatExists {
		return nil, errors.New(creator.ErrorCode().String())
	}
	chatID := creator.GetChatID()
	return &pb.Status{Success: true, ChatID: chatID}, nil
}

// Send send a message to a chat
func (s *ChatServer) Send(ctx context.Context, in *pb.SendMessage) (*pb.Status, error) {
	return nil, nil
}
