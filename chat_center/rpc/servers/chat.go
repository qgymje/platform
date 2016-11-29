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

func srvChatToPbChat(s *chats.ChatInfo) *pb.ChatInfo {
	return &pb.ChatInfo{
		ChatID:  s.ChatID,
		UserID:  s.UserID,
		Name:    s.Name,
		Members: s.Members,
	}
}

// List get the chat list
func (s *ChatServer) List(ctx context.Context, in *pb.Page) (*pb.ChatList, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.ChatServer.Create error: %+v", err)
		}
	}()

	config := &chats.Config{
		PageNum:  int(in.Num),
		PageSize: int(in.Size),
		UserID:   in.UserID,
	}
	cc := chats.NewChats(config)
	err = cc.Do()
	if err != nil {
		return nil, errors.New(cc.ErrorCode().String())
	}
	chatList := cc.Result()
	var pbChatList []*pb.ChatInfo
	for i := range chatList {
		pbChatInfo := srvChatToPbChat(chatList[i])
		pbChatList = append(pbChatList, pbChatInfo)
	}

	return &pb.ChatList{List: pbChatList}, nil
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
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.ChatServer.Send error: %+v", err)
		}
	}()

	config := &chats.SenderConfig{
		ChatID:  in.ChatID,
		UserID:  in.UserID,
		Content: in.Content,
	}

	sender := chats.NewSender(config)
	if err = sender.Do(); err != nil {
		return nil, errors.New(sender.ErrorCode().String())
	}
	msgID := sender.GetMessageID()
	return &pb.Status{Success: true, MessageID: msgID}, nil
}
