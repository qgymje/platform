package coupons

import "platform/coupon_center/rpc/models"

// Config coupons config
type Config struct {
}

// Coupons service coupons
type Coupons struct {
	config       *Config
	couponFinder *models.CouponFinder
}
