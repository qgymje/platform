package coupons

import (
	"platform/commons/codes"
	"platform/coupon_center/rpc/models"
	"platform/utils"
	"time"
)

// UserCouponsConfig coupons config
type UserCouponsConfig struct {
	UserID    string
	PageSize  int
	PageNum   int
	StartTime time.Time
	EndTime   time.Time
}

// UserCoupons service coupons
type UserCoupons struct {
	config          *UserCouponsConfig
	userCoupnFinder *models.UserCouponFinder

	errorCode codes.ErrorCode
}

// NewUserCoupons create a service level coupons
func NewUserCoupons(c *UserCouponsConfig) *UserCoupons {
	cp := new(UserCoupons)
	cp.config = c
	cp.userCoupnFinder = models.NewUserCouponFinder().UserID(c.UserID)
	return cp
}

// ErrorCode error code
func (c *UserCoupons) ErrorCode() codes.ErrorCode {
	return c.errorCode
}

// Do find by user_id
func (c *UserCoupons) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("coupons.UserCoupons.Do error: %+v", err)
		}
	}()

	if err = c.userCoupnFinder.Do(); err != nil {
		if err == models.ErrNotFound {
			c.errorCode = codes.ErrorCodeUserCouponNotFound
		} else {
			c.errorCode = codes.ErrorCodeUserCouponFind
			return
		}
	}

	return
}

// Result return user coupons
func (c *UserCoupons) Result() (uc []*UserCoupon) {
	mUserCoupons := c.userCoupnFinder.Result()
	for i := range mUserCoupons {
		userCoupon := modelUserCouponToSrvUserCoupon(mUserCoupons[i])
		uc = append(uc, userCoupon)
	}
	return
}

// Count total number
func (c *UserCoupons) Count() int64 {
	return c.userCoupnFinder.Count()
}
