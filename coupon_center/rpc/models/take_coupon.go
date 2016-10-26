package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// TakeCoupon user take sendcoupon record
//go:generate gen_columns -tag=bson -path=./take_coupon.go
type TakeCoupon struct {
	SendCouponID bson.ObjectId `bson:"send_coupon_id"`
	UserID       bson.ObjectId `bson:"user_id"`
	CreatedAt    time.Time     `bson:"created_at"`
}
