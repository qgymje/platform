package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// TakeCoupon take coupon model
type TakeCoupon struct {
	ID           int64  `orm:"column(id)"`
	SendCouponID int64  `orm:"column(sendcoupon_id)"`
	UserID       string `orm:"column(user_id)"`
	CreatedAt    time.Time
}

// TableName table name
func (TakeCoupon) TableName() string {
	return TableNameTakeCoupon
}

// Create create a take coupon
func (t *TakeCoupon) Create() (err error) {
	err = GetDB().Begin()

	sc := &SendCoupon{ID: t.SendCouponID}
	if err = sc.UpdateNumber(-1); err != nil {
		GetDB().Rollback()
		return
	}

	t.CreatedAt = time.Now()
	if _, err = GetDB().Insert(t); err != nil {
		GetDB().Rollback()
		return err
	}

	if err = GetDB().Commit(); err != nil {
		return
	}
	return
}

// HasTaken has taken?
func (t *TakeCoupon) HasTaken() bool {
	err := GetDB().QueryTable(TableNameTakeCoupon).Filter("sendcoupon_id", t.SendCouponID).Filter("user_id", t.UserID).One(t)
	if err == orm.ErrNoRows {
		return false
	}
	return true
}
