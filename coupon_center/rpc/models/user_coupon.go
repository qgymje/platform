package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// UserCoupon both host and audience's coupon
//go:generate gen_columns -tag=bson -path=./user_coupon.go
type UserCoupon struct {
	UserID    bson.ObjectId `bson:"user_id"`
	CouponID  bson.ObjectId `bson:"coupon_id"`
	Number    int           `bson:"number"` // number will be added or minused
	CreatedAt time.Time     `bson:"created_at"`
}
