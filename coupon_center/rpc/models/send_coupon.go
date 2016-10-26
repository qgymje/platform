package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SendCoupon a send coupon object
//go:generate gen_columns -tag=bson -path=./send_coupon.go
type SendCoupon struct {
	SendCouponID bson.ObjectId `bson:"_id"`
	CouponID     bson.ObjectId `bson:"coupon_id"`
	BroadcastID  bson.ObjectId `bson:"broadcast_id"`
	UserID       bson.ObjectId `bson:"user_id"`
	Number       int           `bson:"number"`
	Duration     int           `bson:"duration"`
	CreatedAt    time.Time     `bson:"created_at"`
	ClosedAt     time.Time     `bson:"closed_at"`
}

// GetID get id
func (s *SendCoupon) GetID() string {
	return s.SendCouponID.Hex()
}

// GetCouponID get coupon id
func (s *SendCoupon) GetCouponID() string {
	return s.CouponID.Hex()
}

// GetBroadcastID get broadcast id
func (s *SendCoupon) GetBroadcastID() string {
	return s.BroadcastID.Hex()
}

// GetUserID get user id
func (s *SendCoupon) GetUserID() string {
	return s.UserID.Hex()
}

// RemainTime remain time
func (s *SendCoupon) RemainTime() int64 {
	return int64(time.Since(s.CreatedAt).Seconds())
}

// Create create a send coupon
func (s *SendCoupon) Create() error {
	session := GetMongo()
	defer session.Close()

	s.CreatedAt = time.Now()

	return session.DB(DBName).C(ColNameSendCoupon).Insert(&s)
}
