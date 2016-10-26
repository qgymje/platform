package couponClient

import (
	"context"
	pb "platform/commons/protos/coupon"
	"platform/utils"

	"google.golang.org/grpc"
)

// Coupon grpc client email object
type Coupon struct {
	conn   *grpc.ClientConn
	client pb.CouponClient
}

// NewCoupon new grpc client coupon object
func NewCoupon(address string) *Coupon {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("coupon grpc server cant not connect: %v", err)
	}

	c := new(Coupon)
	c.conn = conn
	c.client = pb.NewCouponClient(c.conn)

	return c
}

// Close close sms grpc client
func (c *Coupon) Close() error {
	return c.conn.Close()
}

// List list
func (c *Coupon) List(in *pb.Page) (*pb.Coupons, error) {
	defer c.Close()
	return c.client.List(context.Background(), in)
}

// Send send
func (c *Coupon) Send(in *pb.SendCoupon) (*pb.Status, error) {
	defer c.Close()
	return c.client.Send(context.Background(), in)
}

// Take take
func (c *Coupon) Take(in *pb.TakeCoupon) (*pb.Status, error) {
	defer c.Close()
	return c.client.Take(context.Background(), in)
}
