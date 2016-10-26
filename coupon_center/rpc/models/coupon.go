package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Coupon coupon model object
//go:generate gen_columns -tag=bson -path=./coupon.go
type Coupon struct {
	CouponID      bson.ObjectId `bson:"_id"`
	CampaignID    bson.ObjectId `bson:"campaign_id"`
	CouponType    int           `bson:"type"`
	TargetType    int           `bson:"target_type"`
	Name          string        `bson:"name"`
	Image         string        `bson:"image"`
	Description   string        `bson:"description"`
	Price         float64       `bson:"price"`
	InitialNumber int           `bson:"initial_number"`
	TakenNumber   int           `bosn:"taken_number"`
	UsedNumber    int           `bson:"used_number"`
	UnitPrice     float64       `bson:"unit_price"`
	IsRunning     bool          `bson:"is_running"`
	IsImmediate   bool          `bson:"is_immediate"`
	CreatedAt     time.Time     `bson:"created_at"`
	UpdatedAt     time.Time     `bson:"updated_at"`
	DeletedAt     time.Time     `bson:"deleted_at"`
}

// GetID get coupon id
func (c *Coupon) GetID() string {
	return c.CouponID.Hex()
}

// GetCampaignID get campaign id
func (c *Coupon) GetCampaignID() string {
	return c.CampaignID.Hex()
}
