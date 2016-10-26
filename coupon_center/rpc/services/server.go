package services

import (
	pb "platform/commons/protos/coupon"

	"golang.org/x/net/context"
)

// Server represents the coupon service grpc server
type Server struct {
}

// List list
func (s *Server) List(context.Context, *pb.Page) (*pb.Coupons, error) {
	return nil, nil
}

// Send send
func (s *Server) Send(context.Context, *pb.SendCoupon) (*pb.Status, error) {
	return &pb.Status{Success: true}, nil
}

// Take take
func (s *Server) Take(context.Context, *pb.TakeCoupon) (*pb.Status, error) {
	return &pb.Status{Success: true}, nil
}
