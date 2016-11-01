package models

import (
	"platform/game_center/rpc/models"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// SendCoupon a send coupon object
type SendCoupon struct {
	ID          int64   `orm:"column(id)"`
	Coupon      *Coupon `orm:"rel(fk)"`
	BroadcastID string  `orm:"column(broadcast_id)"`
	UserID      string  `orm:"column(user_id)"`
	Number      int
	Duration    int64
	CreatedAt   time.Time
	ClosedAt    time.Time `orm:"null"`
}

// TableName table name
func (SendCoupon) TableName() string {
	return TableNameSendCoupon
}

// Create a sendcoupon
func (s *SendCoupon) Create() (err error) {
	s.CreatedAt = time.Now()
	_, err = GetDB().Insert(s)
	if err != nil {
		return err
	}
	return
}

// Find find with coupon object
func (s *SendCoupon) Find() (err error) {
	return GetDB().QueryTable(TableNameSendCoupon).RelatedSel("Coupon").Filter("id", s.ID).One(s)
}

// GetID get id
func (s *SendCoupon) GetID() string {
	return strconv.FormatInt(s.ID, 10)
}

// GetCouponID get coupon id
func (s *SendCoupon) GetCouponID() string {
	return strconv.FormatInt(s.Coupon.ID, 10)
}

// GetBroadcastID get broadcast id
func (s *SendCoupon) GetBroadcastID() string {
	return s.BroadcastID
}

// GetUserID get user id
func (s *SendCoupon) GetUserID() string {
	return s.UserID
}

// RemainTime remain time
func (s *SendCoupon) RemainTime() int64 {
	return s.Duration - int64(time.Since(s.CreatedAt).Seconds())
}

// IsClosed is closed?
func (s *SendCoupon) IsClosed() bool {
	if s.ClosedAt.IsZero() {
		return false
	}
	return true
}

// Close by stoper
func (s *SendCoupon) Close() error {
	s.ClosedAt = time.Now()
	if _, err := GetDB().Update(s, "closed_at"); err != nil {
		return err
	}
	return nil
}

// TryClose try to close
func (s *SendCoupon) TryClose() error {
	if s.RemainTime() > 0 {
		return nil
	}

	s.ClosedAt = time.Now()
	if _, err := GetDB().Update(s, "closed_at"); err != nil {
		return err
	}
	return nil
}

// UpdateNumber update number
func (s *SendCoupon) UpdateNumber(num int) (err error) {
	s.Number += num
	if _, err := GetDB().Update(s, "number"); err != nil {
		return err
	}
	return
}

// FindRunningSendCoupons find running sendCoupons
func FindRunningSendCoupons() (sc []*SendCoupon, err error) {
	_, err = GetDB().QueryTable(TableNameSendCoupon).RelatedSel("Coupon").Filter("closed_at__isnull", true).All(&sc)
	if err == orm.ErrNoRows {
		return nil, models.ErrNotFound
	}
	return
}
