package coupons

import (
	"platform/commons/codes"
	"platform/coupon_center/rpc/models"
	"platform/utils"
	"strconv"
)

// TakerConfig config
type TakerConfig struct {
	SendCouponID string
	UserID       string
}

// Taker user take a coupon
type Taker struct {
	config          *TakerConfig
	takeCouponModel *models.TakeCoupon

	errorCode codes.ErrorCode
}

// NewTaker new taker
func NewTaker(c *TakerConfig) *Taker {
	t := new(Taker)
	t.config = c
	t.takeCouponModel = &models.TakeCoupon{}
	return t
}

// ErrorCode error code
func (t *Taker) ErrorCode() codes.ErrorCode {
	return t.errorCode
}

// Do do the dirty job
func (t *Taker) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("couopns.Taker.Do error: %+v", err)
		}
	}()

	if err = t.save(); err != nil {
		t.errorCode = codes.ErrorCodeTakeCouponCreate
		return err
	}
	return
}

func (t *Taker) save() (err error) {
	scID, _ := strconv.ParseInt(t.config.SendCouponID, 10, 0)
	t.takeCouponModel.SendCouponID = scID
	t.takeCouponModel.UserID = t.config.UserID

	return t.takeCouponModel.Create()
}
