package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// UserCoupon both host and audience's coupon
type UserCoupon struct {
	UserID    bson.ObjectId `bson:"user_id"`
	CouponID  bson.ObjectId `bson:"coupon_id"`
	Number    int           `bson:"number"` // number will be added or minused
	CreatedAt time.Time     `bson:"created_at"`
}
