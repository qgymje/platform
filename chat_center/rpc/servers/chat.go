package servers

import (
	pb "platform/commons/protos/chat"

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
	return nil, nil
}

// Send send a message to a chat
func (s *ChatServer) Send(ctx context.Context, in *pb.SendMessage) (*pb.Status, error) {
	return nil, nil
}
