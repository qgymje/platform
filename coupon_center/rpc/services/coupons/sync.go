package coupons

import (
	"encoding/json"
	"fmt"
	"platform/commons/queues"
	"platform/coupon_center/rpc/models"
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
			utils.Dump(allSendCoupons)
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

func fetchAllSendCoupons() ([]*models.SendCoupon, error) {
	return models.FindRunningSendCoupons()
}

// SendCouponSync send cuopon sync
type SendCouponSync struct {
	sendcoupon *models.SendCoupon
}

// NewSendCouponSync new send coupon sync
func NewSendCouponSync(sc *models.SendCoupon) *SendCouponSync {
	s := new(SendCouponSync)
	s.sendcoupon = sc
	return s
}

// Do do the dirty work
func (s *SendCouponSync) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("coupons.SendCouponSync.Do error: %+v", err)
		}
	}()
	if err = s.notify(); err != nil {
		return err
	}
	return
}

func (s *SendCouponSync) notify() error {
	return notifier.Publish(s)
}

// Topic publish topic
func (s *SendCouponSync) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), s.sendcoupon.GetCouponID())
}

// Message publish message
func (s *SendCouponSync) Message() []byte {
	var msg []byte
	sendCouponMsg := queues.MessageSendCouponUpdate{
		SendCouponID: s.sendcoupon.GetCouponID(),
		BroadcastID:  s.sendcoupon.BroadcastID,
		RemainAmount: s.sendcoupon.Number,
		RemainTime:   s.sendcoupon.RemainTime(),
		CouponID:     s.sendCoupon.Coupon.GetID(),
		Description:  s.sendcoupon.Coupon.Description,
		Image:        s.sendcoupon.Coupon.Image,
		Name:         s.sendcoupon.Coupon.Name,
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
