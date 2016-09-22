package emailClient

import (
	"context"

	pb "platform/commons/protos/email"
	"platform/utils"

	"google.golang.org/grpc"
)

// Email grpc client email object
type Email struct {
	conn   *grpc.ClientConn
	client pb.EmailClient
}

// NewEmail new grpc client sms object
func NewEmail(address string) *Email {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("email grpc server cant not connect: %v", err)
	}

	s := new(Email)
	s.conn = conn
	s.client = pb.NewEmailClient(s.conn)

	return s
}

// Close close sms grpc client
func (s *Email) Close() error {
	return s.conn.Close()
}

// Verify Email object verify
func (s *Email) Verify(in *pb.EmailCode) (*pb.Status, error) {
	defer s.Close()
	return s.client.Verify(context.Background(), in)
}
