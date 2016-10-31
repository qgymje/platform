package models

import "time"

// UseCoupon use coupon model
type UseCoupon struct {
	ID        int64  `orm:"column(id)"`
	UserID    string `orm:"column(user_id)"`
	CouponID  int64  `orm:"column(cupon_id)"`
	Latitude  float64
	Longitude float64
	UsedAt    time.Time
}

// TableName table name
func (UseCoupon) TableName() string {
	return TableNameUseCoupon
}
