package email

import (
	pb "platform/commons/protos/email"

	"golang.org/x/net/context"
)

// Server represent a email service implement
type Server struct {
}

// Verify a email code
func (s *Server) Verify(context.Context, *pb.EmailCode) (*pb.Status, error) {
	return &pb.Status{Success: true}, nil
}
