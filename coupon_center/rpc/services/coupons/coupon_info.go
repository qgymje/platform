package coupons

import "time"

// Coupon service level coupon
type Coupon struct {
	CouponID    string
	Name        string
	Image       string
	Number      int
	Description string
	Price       float64
	StartTime   time.Time
	EndTime     time.Time
}
