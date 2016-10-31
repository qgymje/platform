package coupons

import (
	"encoding/json"
	"errors"
	"fmt"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/coupon_center/rpc/models"
	"platform/coupon_center/rpc/services/notifier"
	"platform/utils"
	"strconv"
	"time"
)

// SenderConfig sender config
type SenderConfig struct {
	TypeID      int
	UserID      string
	CouponID    string
	BroadcastID string
	Number      int
	Duration    int64
}

// SendCoupon service level send coupon
type SendCoupon struct {
	SendCouponID string `json:"send_coupon_id"`
	BroadcastID  string `json:"broadcast_id"`
	RemainAmount int    `json:"remain_amount"`
	RemainTime   int64  `json:"remain_time"`

	Coupon
}

// Sender sender
type Sender struct {
	config          *SenderConfig
	sendCouponModel *models.SendCoupon
	userCouponModel *models.UserCoupon

	userCouponFinder *models.UserCouponFinder

	errorCode codes.ErrorCode
}

// NewSender new sender
func NewSender(c *SenderConfig) *Sender {
	s := new(Sender)
	s.config = c
	s.sendCouponModel = &models.SendCoupon{}
	s.userCouponModel = &models.UserCoupon{}
	s.userCouponFinder = models.NewUserCouponFinder()
	return s
}

// ErrorCode error coder
func (s *Sender) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do do the dirty work
func (s *Sender) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("coupons.Sender.Do error: %+v", err)
		}
	}()

	if err = s.validSendCoupon(); err != nil {
		return
	}

	if err = s.reduceNumber(); err != nil {
		s.errorCode = codes.ErrorCodeUserCouponUpdate
		return
	}

	if err = s.save(); err != nil {
		return
	}

	if err = s.notify(); err != nil {
		return
	}

	return
}

// GetSendCouponID get sendCoupon id
func (s *Sender) GetSendCouponID() string {
	return s.sendCouponModel.GetID()
}

func (s *Sender) validSendCoupon() (err error) {
	s.userCouponFinder.UserID(s.config.UserID).CouponID(s.config.CouponID)
	if err = s.userCouponFinder.Do(); err != nil {
		if err == models.ErrNotFound {
			s.errorCode = codes.ErrorCodeUserCouponNotFound
		} else {
			s.errorCode = codes.ErrorCodeUserCouponFind
		}
		return
	}

	s.userCouponModel = s.userCouponFinder.One()

	if s.userCouponModel.Number < s.config.Number {
		s.errorCode = codes.ErrorCodeSendCouponNumberNotEnough
		return errors.New("number not enough")
	}

	return
}

func (s *Sender) reduceNumber() (err error) {
	return s.userCouponModel.UpdateNumber(-s.config.Number)
}

func (s *Sender) save() (err error) {
	s.sendCouponModel.Coupon = &models.Coupon{}
	s.sendCouponModel.UserID = s.config.UserID
	id, _ := strconv.Atoi(s.config.CouponID)
	s.sendCouponModel.Coupon.ID = int64(id)
	s.sendCouponModel.BroadcastID = s.config.BroadcastID
	s.sendCouponModel.Number = s.config.Number
	s.sendCouponModel.Duration = s.config.Duration

	if err = s.sendCouponModel.Create(); err != nil {
		return
	}

	return s.sendCouponModel.Find()
}

func (s *Sender) notify() (err error) {
	return notifier.Publish(s)
}

func (s *Sender) autoStop() {
	timeout := time.After(s.config.Duration * time.Second)
	for {
		select {
		case <-timeout:
			config := &StoperConfig{
				SendCouponID: s.sendCouponModel.GetID(),
				BroadcastID:  s.config.BroadcastID,
				TypeID:       10003,
			}
			stoper := NewStoper(config)
			if err := stoper.Do(); err != nil {
				utils.GetLog().Error("coupons.Sender.autoStop error: %+v", err)
			}
		}
	}
}

// Topic NSQ topic
func (s *Sender) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), s.config.BroadcastID)
}

// Message publish message
func (s *Sender) Message() []byte {
	var msg []byte
	sendCouponMsg := queues.MessageSendCouponUpdate{
		SendCouponID: s.sendCouponModel.GetID(),
		BroadcastID:  s.sendCouponModel.GetBroadcastID(),
		RemainAmount: s.sendCouponModel.Number,
		RemainTime:   s.sendCouponModel.RemainTime(),

		CouponID:    s.sendCouponModel.Coupon.GetID(),
		Description: s.sendCouponModel.Coupon.Description,
		Image:       s.sendCouponModel.Coupon.Image,
		Name:        s.sendCouponModel.Coupon.Name,
	}

	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		int(s.config.TypeID),
		sendCouponMsg,
	}
	msg, _ = json.Marshal(data)
	return msg
}
