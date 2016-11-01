package coupons

import (
	"encoding/json"
	"fmt"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/commons/typeids"
	"platform/coupon_center/rpc/models"
	"platform/coupon_center/rpc/services/notifier"
	"platform/utils"
	"strconv"
)

// StopperConfig stopper config
type StopperConfig struct {
	SendCouponID string
	UserID       string
}

// Stopper stop the send coupon
type Stopper struct {
	config          *StopperConfig
	sendCouponModel *models.SendCoupon
	userCouponModel *models.UserCoupon

	errorCode codes.ErrorCode
}

// NewStopper create a stoper
func NewStopper(c *StopperConfig) *Stopper {
	s := new(Stopper)
	s.config = c
	s.sendCouponModel = &models.SendCoupon{}
	s.userCouponModel = &models.UserCoupon{}
	s.userCouponModel.Coupon = &models.Coupon{}
	return s
}

// ErrorCode error code
func (s *Stopper) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do do the dirty job
func (s *Stopper) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("couopns.Stopper.Do error: %+v", err)
		}
	}()

	if err = s.updateUserCoupon(); err != nil {
		return
	}

	if err = s.closeSendCoupon(); err != nil {
		return
	}

	if err = s.notify(); err != nil {
		return
	}

	return
}

func (s *Stopper) updateUserCoupon() (err error) {
	id, _ := strconv.ParseInt(s.config.SendCouponID, 10, 0)
	s.sendCouponModel.ID = id
	s.sendCouponModel.UserID = s.config.UserID

	if err = s.sendCouponModel.Find(); err != nil {
		return err
	}

	s.userCouponModel.UserID = s.config.UserID
	s.userCouponModel.Coupon.ID = s.sendCouponModel.Coupon.ID
	if err = s.userCouponModel.Find(); err != nil {
		return err
	}

	return s.userCouponModel.UpdateNumber(s.sendCouponModel.Number)
}

func (s *Stopper) closeSendCoupon() (err error) {
	return s.sendCouponModel.Close()
}

func (s *Stopper) notify() (err error) {
	return notifier.Publish(s)
}

// Topic NSQ topic
func (s *Stopper) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), s.sendCouponModel.BroadcastID)
}

// Message publish message
func (s *Stopper) Message() []byte {
	var msg []byte
	sendCouponMsg := queues.MessageSendCouponStop{
		SendCouponID: s.config.SendCouponID,
		BroadcastID:  s.sendCouponModel.BroadcastID,
		StopTime:     s.sendCouponModel.ClosedAt.Unix(),
	}

	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		int(typeids.CouponSenderStop),
		sendCouponMsg,
	}
	msg, _ = json.Marshal(data)
	return msg
}
