package sms

import (
	pb "platform/commons/protos/sms"

	"golang.org/x/net/context"
)

// Server represent a sms service implement
type Server struct {
}

// Verify a sms code
func (s *Server) Verify(context.Context, *pb.PhoneCode) (*pb.Status, error) {
	return &pb.Status{Success: true}, nil
}
