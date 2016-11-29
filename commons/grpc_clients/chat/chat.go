package chatClient

import (
	"context"
	pb "platform/commons/protos/chat"
	"platform/utils"

	"google.golang.org/grpc"
)

// Chat grpc client email object
type Chat struct {
	conn   *grpc.ClientConn
	client pb.ChatClient
}

// NewChat new grpc client chat object
func NewChat(address string) *Chat {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("chat grpc server cant not connect: %v", err)
	}

	c := new(Chat)
	c.conn = conn
	c.client = pb.NewChatClient(c.conn)

	return c
}

// Close close sms grpc client
func (c *Chat) Close() error {
	return c.conn.Close()
}

// List list
func (c *Chat) List(in *pb.Page) (*pb.ChatList, error) {
	defer c.Close()
	return c.client.List(context.Background(), in)
}

// Create create
func (c *Chat) Create(in *pb.Creator) (*pb.Status, error) {
	defer c.Close()
	return c.client.Create(context.Background(), in)
}

// Send send message
func (c *Chat) Send(in *pb.SendMessage) (*pb.Status, error) {
	defer c.Close()
	return c.client.Send(context.Background(), in)
}
