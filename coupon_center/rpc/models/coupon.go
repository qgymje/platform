package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Coupon coupon model object
type Coupon struct {
	CouponID      bson.ObjectId `bson:"_id"`
	CouponType    int           `bson:"type"`
	TargetType    int           `bson:"target_type"`
	Name          string        `bson:"name"`
	Image         string        `bson:"image"`
	Description   string        `bson:"description"`
	Value         float64       `bson:"price"`
	InitialNumber int           `bson:"initial_number"`
	TakenNumber   int           `bosn:"taken_number"`
	UsedNumber    int           `bson:"used_number"`
	UintPrice     float64       `bson:"uint_price"`
	IsRunning     bool          `bson:"is_running"`
	IsImmediate   bool          `bson:"is_immediate"`
	CreatedAt     time.Time     `bson:"created_at"`
	UpdatedAt     time.Time     `bson:"updated_at"`
	DeletedAt     time.Time     `bson:"deleted_at"`
}
