package sms

import (
	pb "platform/commons/protos/sms"

	"golang.org/x/net/context"
)

// Server represent a sms service implement
type Server struct {
}

// Verify a sms code
func (s *Server) Verify(ctx context.Context, in *pb.PhoneCode) (*pb.Status, error) {
	//code := NewRegisterCode()
	//code.Verify(in.Country, in.Phone, in.Code)
	return &pb.Status{Success: true}, nil
}
