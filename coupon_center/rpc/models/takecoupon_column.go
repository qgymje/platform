package models

type _TakeCouponColumn struct {
	CreatedAt    string
	SendCouponID string
	UserID       string
}

// TakeCouponColumns takecoupon columns name
var TakeCouponColumns _TakeCouponColumn

func init() {
	TakeCouponColumns.CreatedAt = "created_at"
	TakeCouponColumns.SendCouponID = "send_coupon_id"
	TakeCouponColumns.UserID = "user_id"

}
