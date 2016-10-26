package models

type _UserCouponColumn struct {
	CouponID  string
	CreatedAt string
	Number    string
	UserID    string
}

// UserCouponColumns usercoupon columns name
var UserCouponColumns _UserCouponColumn

func init() {
	UserCouponColumns.CouponID = "coupon_id"
	UserCouponColumns.CreatedAt = "created_at"
	UserCouponColumns.Number = "number"
	UserCouponColumns.UserID = "user_id"

}
