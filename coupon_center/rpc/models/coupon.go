package models

import (
	"strconv"
	"time"
)

// Coupon coupon model object
//go:generate gen_columns -tag=db -path=./coupon.go
type Coupon struct {
	ID            int64 `orm:"column(id)"`
	CampaignID    int64 `orm:"column(campaign_id)"`
	CouponType    int
	TargetType    int
	Name          string
	Image         string
	Description   string
	Price         float64
	InitialNumber int
	TakenNumber   int
	UsedNumber    int
	UnitPrice     float64
	IsRunning     bool
	IsImmediate   bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

// TableName table name
func (Coupon) TableName() string {
	return TableNameCoupon
}

// GetID get coupon id
func (c *Coupon) GetID() string {
	return strconv.FormatInt(c.ID, 10)
}

// GetCampaignID get campaign id
func (c *Coupon) GetCampaignID() string {
	return strconv.FormatInt(c.CampaignID, 10)
}

// Create create a coupon
func (c *Coupon) Create() (err error) {
	now := time.Now()

	c.CreatedAt = now
	c.UpdatedAt = now

	return
}
