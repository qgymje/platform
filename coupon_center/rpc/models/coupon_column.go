package models

type _CouponColumn struct {
	CampaignID    string
	CouponID      string
	CouponType    string
	CreatedAt     string
	DeletedAt     string
	Description   string
	Image         string
	InitialNumber string
	IsImmediate   string
	IsRunning     string
	Name          string
	Price         string
	TakenNumber   string
	TargetType    string
	UnitPrice     string
	UpdatedAt     string
	UsedNumber    string
}

// CouponColumns coupon columns name
var CouponColumns _CouponColumn

func init() {
	CouponColumns.CampaignID = "campaign_id"
	CouponColumns.CouponID = "_id"
	CouponColumns.CouponType = "type"
	CouponColumns.CreatedAt = "created_at"
	CouponColumns.DeletedAt = "deleted_at"
	CouponColumns.Description = "description"
	CouponColumns.Image = "image"
	CouponColumns.InitialNumber = "initial_number"
	CouponColumns.IsImmediate = "is_immediate"
	CouponColumns.IsRunning = "is_running"
	CouponColumns.Name = "name"
	CouponColumns.Price = "price"
	CouponColumns.TakenNumber = ""
	CouponColumns.TargetType = "target_type"
	CouponColumns.UnitPrice = "unit_price"
	CouponColumns.UpdatedAt = "updated_at"
	CouponColumns.UsedNumber = "used_number"

}
