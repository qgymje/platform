package models

type _SendCouponColumn struct {
	BroadcastID  string
	CouponID     string
	CreatedAt    string
	Duration     string
	Number       string
	SendCouponID string
	UserID       string
}

// SendCouponColumns sendcoupon columns name
var SendCouponColumns _SendCouponColumn

func init() {
	SendCouponColumns.BroadcastID = "broadcast_id"
	SendCouponColumns.CouponID = "coupon_id"
	SendCouponColumns.CreatedAt = "created_at"
	SendCouponColumns.Duration = "duration"
	SendCouponColumns.Number = "number"
	SendCouponColumns.SendCouponID = "_id"
	SendCouponColumns.UserID = "user_id"

}
