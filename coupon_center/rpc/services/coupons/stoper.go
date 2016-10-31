package coupons

import (
	"encoding/json"
	"fmt"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/coupon_center/rpc/services/notifier"
)

// StoperConfig stoper config
type StoperConfig struct {
	SendCouponID string
	BroadcastID  string
	TypeID       int // 10003
}

// Stoper stop the send coupon
type Stoper struct {
	config    *StoperConfig
	errorCode codes.ErrorCode
}

// NewStoper create a stoper
func NewStoper(c *StoperConfig) *Stoper {
	s := new(Stoper)
	s.config = c
	return s
}

// ErrorCode error code
func (s *Stoper) ErrorCode() codes.ErrorCode {
	return t.errorCode
}

// Do do the dirty job
func (s *Stoper) Do() (err error) {
	return
}

func (s *Stoper) notify() (err error) {
	return notifier.Publish(s)
}

// Topic NSQ topic
func (s *Stoper) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), s.config.BroadcastID)
}

// Message publish message
func (s *Stoper) Message() []byte {
	var msg []byte
	sendCouponMsg := queues.MessageSendCouponStop{
		SendCouponID: s.config.SendCouponID,
		BroadcastID:  s.config.BroadcastID,
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
