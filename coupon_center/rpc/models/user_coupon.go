package models

import "time"

// UserCoupon both host and audience's coupon
type UserCoupon struct {
	ID        int64   `orm:"column(id)"`
	UserID    string  `orm:"column(user_id)"`
	Coupon    *Coupon `orm:"rel(fk)"`
	Number    int
	CreatedAt time.Time
}

// TableName table name
func (UserCoupon) TableName() string {
	return TableNameUserCoupon
}

// Create a user coupon record
func (uc *UserCoupon) Create() (err error) {
	now := time.Now()
	uc.CreatedAt = now

	return nil
}

// Find find with coupon object
func (uc *UserCoupon) Find() (err error) {
	return GetDB().QueryTable(TableNameUserCoupon).RelatedSel("Coupon").Filter("user_id", uc.UserID).Filter("coupon_id", uc.Coupon.ID).One(uc)
}

// UpdateNumber update number
func (uc *UserCoupon) UpdateNumber(num int) (err error) {
	uc.Number += num
	if _, err := GetDB().Update(uc, "number"); err != nil {
		return err
	}
	return
}
