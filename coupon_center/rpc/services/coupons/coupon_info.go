package coupons

import (
	"platform/coupon_center/rpc/models"
	"time"
)

// Coupon service level coupon
type Coupon struct {
	CouponID    string
	Name        string
	Image       string
	Description string
	Price       float64
}

// UserCoupon service level coupon
type UserCoupon struct {
	CouponID    string
	Name        string
	Image       string
	Description string
	Price       float64

	Number    int
	CreatedAt time.Time
}

func modelUserCouponToSrvUserCoupon(m *models.UserCoupon) *UserCoupon {
	return &UserCoupon{
		CouponID:    m.Coupon.GetID(),
		Number:      m.Number,
		CreatedAt:   m.CreatedAt,
		Name:        m.Coupon.Name,
		Image:       m.Coupon.Image,
		Description: m.Coupon.Description,
		Price:       m.Coupon.Price,
	}
}
