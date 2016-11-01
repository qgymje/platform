package services

import (
	"errors"
	pb "platform/commons/protos/coupon"
	"platform/coupon_center/rpc/services/coupons"
	"platform/utils"

	"golang.org/x/net/context"
)

// Server represents the coupon service grpc server
type Server struct {
}

func srvUserCouponToPbUserCoupon(c *coupons.UserCoupon) *pb.CouponInfo {
	return &pb.CouponInfo{
		CouponID:    c.CouponID,
		Name:        c.Name,
		Image:       c.Image,
		Number:      int64(c.Number),
		Description: c.Description,
		Price:       float32(c.Price),
	}
}

// List list
func (s *Server) List(ctx context.Context, in *pb.Page) (*pb.Coupons, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("coupons.Server.List error: %+v", err)
		}
	}()

	config := &coupons.UserCouponsConfig{
		UserID:   in.UserID,
		PageSize: int(in.Size),
		PageNum:  int(in.Num),
	}

	uc := coupons.NewUserCoupons(config)
	err = uc.Do()
	if err != nil {
		return nil, errors.New(uc.ErrorCode().String())
	}

	pbUserCoupons := []*pb.CouponInfo{}
	srvUserCoupons := uc.Result()
	for i := range srvUserCoupons {
		pbUserCoupon := srvUserCouponToPbUserCoupon(srvUserCoupons[i])
		pbUserCoupons = append(pbUserCoupons, pbUserCoupon)
	}

	totalNumber := uc.Count()
	return &pb.Coupons{Coupons: pbUserCoupons, TotalNum: totalNumber}, nil
}

// Send send
func (s *Server) Send(ctx context.Context, in *pb.SendCoupon) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("coupons.Server.Send error: %+v", err)
		}
	}()

	config := &coupons.SenderConfig{
		TypeID:      int(in.TypeID),
		UserID:      in.UserID,
		CouponID:    in.CouponID,
		BroadcastID: in.BroadcastID,
		Number:      int(in.Number),
		Duration:    in.Duration,
	}
	sender := coupons.NewSender(config)
	if err = sender.Do(); err != nil {
		return nil, errors.New(sender.ErrorCode().String())
	}

	return &pb.Status{Success: true, SendCouponID: sender.GetSendCouponID()}, nil
}

// Take take
func (s *Server) Take(ctx context.Context, in *pb.TakeCoupon) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("coupons.Server.Take error: %+v", err)
		}
	}()

	config := &coupons.TakerConfig{
		SendCouponID: in.SendCouponID,
		UserID:       in.UserID,
	}
	taker := coupons.NewTaker(config)
	if err = taker.Do(); err != nil {
		return nil, errors.New(taker.ErrorCode().String())
	}

	return &pb.Status{Success: true, SendCouponID: in.SendCouponID}, nil
}

// Stop stop
func (s *Server) Stop(ctx context.Context, in *pb.TakeCoupon) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("coupons.Server.Stoperror: %+v", err)
		}
	}()

	config := &coupons.StopperConfig{
		SendCouponID: in.SendCouponID,
		UserID:       in.UserID,
	}
	stopper := coupons.NewStopper(config)
	if err = stopper.Do(); err != nil {
		return nil, errors.New(stopper.ErrorCode().String())
	}

	return &pb.Status{Success: true, SendCouponID: in.SendCouponID}, nil
}
