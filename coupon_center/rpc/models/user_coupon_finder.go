package models

import (
	"platform/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// UserCouponFinder user coupon finder
type UserCouponFinder struct {
	offset, limit      int
	userID, couponID   string
	startTime, endTime time.Time

	userCoupons []*UserCoupon

	query orm.QuerySeter
}

// NewUserCouponFinder create a user coupon finder
func NewUserCouponFinder() *UserCouponFinder {
	f := new(UserCouponFinder)
	f.userCoupons = []*UserCoupon{}
	f.query = GetDB().QueryTable(TableNameUserCoupon)
	return f
}

// Limit limit
func (f *UserCouponFinder) Limit(offset, limit int) *UserCouponFinder {
	f.offset = offset
	f.limit = limit

	f.query = f.query.Offset(int64(f.offset))
	f.query = f.query.Limit(f.limit)

	return f
}

// Duration duration
func (f *UserCouponFinder) Duration(st, et int64) *UserCouponFinder {
	f.startTime = time.Unix(st, 0)
	f.endTime = time.Unix(et, 0)
	return f
}

// UserID user_id
func (f *UserCouponFinder) UserID(userID string) *UserCouponFinder {
	f.userID = userID

	f.query = f.query.Filter("user_id", userID)

	return f
}

// CouponID find by coupon id
func (f *UserCouponFinder) CouponID(couponID string) *UserCouponFinder {
	f.couponID = couponID

	id, _ := strconv.Atoi(couponID)
	f.query = f.query.Filter("coupon_id", id)

	return f
}

// Do the query
func (f *UserCouponFinder) Do() (err error) {
	defer func() {
		if err != nil {
			utils.Dump("models.UserCouponFinder.Do error: %+v", err)
		}
	}()

	n, err := f.query.RelatedSel("Coupon").All(&f.userCoupons)
	if err != nil {
		return
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

// Result result
func (f *UserCouponFinder) Result() []*UserCoupon {
	return f.userCoupons
}

// One one
func (f *UserCouponFinder) One() *UserCoupon {
	return f.userCoupons[0]
}

// Count count
func (f *UserCouponFinder) Count() int64 {
	n, _ := f.query.Limit(-1).Count()
	return n
}
