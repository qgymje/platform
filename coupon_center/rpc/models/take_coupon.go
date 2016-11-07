package models

import (
	"errors"
	"platform/utils"
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

// ErrSendCouponClosed send coupon is closed
var ErrSendCouponClosed = errors.New("sendCoupon is closed")

// Create create a take coupon
func (t *TakeCoupon) Create() (err error) {
	err = GetDB().Begin()

	sc := &SendCoupon{ID: t.SendCouponID}
	if err = sc.Find(); err != nil {
		GetDB().Rollback()
		if err == orm.ErrNoRows {
			return ErrNotFound
		}
		return
	}

	if sc.IsClosed() || sc.Number <= 0 {
		GetDB().Rollback()
		return ErrSendCouponClosed
	}

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

// Find find a take coupon
func (t *TakeCoupon) Find() (err error) {
	err = GetDB().QueryTable(TableNameTakeCoupon).Filter("user_id", t.UserID).Filter("sendcoupon_id", t.SendCouponID).One(t)
	if err == orm.ErrNoRows {
		return ErrNotFound
	}
	return err
}

// HasTaken has taken?
func (t *TakeCoupon) HasTaken() bool {
	utils.Dump(t)
	err := GetDB().QueryTable(TableNameTakeCoupon).Filter("sendcoupon_id", t.SendCouponID).Filter("user_id", t.UserID).One(t)
	if err == orm.ErrNoRows {
		return false
	}
	return true
}
