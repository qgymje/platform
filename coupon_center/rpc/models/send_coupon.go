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
}
