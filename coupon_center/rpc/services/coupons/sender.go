package coupons

import (
	"encoding/json"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/coupon_center/rpc/models"
	"platform/coupon_center/rpc/services/notifier"
	"platform/utils"
)

// SendCoupon service level send coupon
type SendCoupon struct {
	SendCouponID string `json:"send_coupon_id"`
	BroadcastID  string `json:"broadcast_id"`
	RemainAmount int    `json:"remain_amount"`
	RemainTime   int64  `json:"remain_time"`
	Coupon       struct {
		CouponID string `json:"coupon_id"`
		Name     string `json:"name"`
		Image    string `json:"image"`
	}
}

// SenderConfig sender config
type SenderConfig struct {
	TypeID int
}

// Sender sender
type Sender struct {
	config          *SenderConfig
	sendCouponModel *models.SendCoupon

	errorCode codes.ErrorCode
}

// NewSender new sender
func NewSender(c *SenderConfig) *Sender {
	s := new(Sender)
	s.config = c
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

	return
}

func (s *Sender) findCoupon() (err error) {
	return
}

func (s *Sender) findSendCoupon() (err error) {
	return
}

func (s *Sender) save() (err error) {
	return
}

func (s *Sender) notify() (err error) {
	return notifier.Publish(s)
}

// Topic NSQ topic
func (s *Sender) Topic() string {
	return queues.TopicSendCouponUpdate.String()
}

// Message publish message
func (s *Sender) Message() []byte {
	_ = s.findSendCoupon()

	var msg []byte
	sendCouponMsg := queues.MessageSendCouponUpdate{
		SendCouponID: s.sendCouponModel.GetID(),
		BroadcastID:  s.sendCouponModel.GetBroadcastID(),
		Number:       s.sendCouponModel.Number,
		RemainTime:   s.sendCouponModel.RemainTime(),
	}

	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		int(s.config.TypeID),
		barrageMsg,
	}
	msg, _ = json.Marshal(data)
	return msg
}
