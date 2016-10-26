package coupons

import (
	"encoding/json"
	"fmt"
	"platform/commons/queues"
	"platform/coupon_center/rpc/services/notifier"
	"platform/utils"
	"time"
)

// Sync sync the sendcoupon
func Sync() {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			allSendCoupons, err := fetchAllSendCoupons()
			if err != nil {
				utils.GetLog().Error("coupons.Sync.fetchAllSendCoupons error: %+v", err)
			}
			for _, sc := range allSendCoupons {
				copSync := NewSendCouponSync(sc)
				if err := copSync.Do(); err != nil {
					utils.GetLog().Error("coupons.SendCouponSync.Do error: %+v", err)
				}
			}
		}
	}
}

func fetchAllSendCoupons() []*SendCoupon {
	return nil
}

// SendCouponSync send cuopon sync
type SendCouponSync struct {
	sendcoupon *SendCoupon
}

// NewSendCouponSync new send coupon sync
func NewSendCouponSync(sc) *SendCouponSync {
	s := new(SendCouponSync)
	return s
}

// Do do the dirty work
func (s *SendCouponSync) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("coupons.SendCouponSync.Do error: %+v", err)
		}
	}()

	if err = b.notify(); err != nil {
		return err
	}
	return
}

func (s *SendCouponSync) notify() error {
	return notifier.Publish(b)
}

// Topic publish topic
func (s *SendCouponSync) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), s.sendcoupon.GetID())
}

// Message publish message
func (s *SendCouponSync) Message() []byte {
	var msg []byte
	sendCouponMsg := queues.MessageSendCouponUpdate{
		SendCouponID: b.sendcoupon.GetID(),
		BroadcastID:  b.sendcoupon.GetBroadcastID(),
		RemainAmount: b.sendcoupon.Number,
		RemainTime:   b.sendcoupon.RemainTime(),
	}

	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		10001,
		sendCouponMsg,
	}
	msg, _ = json.Marshal(&data)
	return msg
}
