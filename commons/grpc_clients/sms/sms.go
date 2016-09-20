package smsClient

import (
	"context"

	pb "platform/commons/protos/sms"
	"platform/utils"

	"google.golang.org/grpc"
)

// SMS grpc client sms object
type SMS struct {
	conn   *grpc.ClientConn
	client pb.SMSClient
}

// NewSMS new grpc client sms object
func NewSMS(address string) *SMS {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("sms grpc server cant not connect: %v", err)
	}

	s := new(SMS)
	s.conn = conn
	s.client = pb.NewSMSClient(s.conn)

	return s
}

// Close close sms grpc client
func (s *SMS) Close() error {
	return s.conn.Close()
}

// Verify SMS object verify
func (s *SMS) Verify(ctx context.Context, in *pb.PhoneCode, opts ...grpc.CallOption) (*pb.Status, error) {
	defer s.Close()
	return s.client.Verify(context.Background(), in, opts...)
}
