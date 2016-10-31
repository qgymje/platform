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

// Sync sync the sendCoupon
func Sync() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			allsendCoupons, err := fetchAllSendCoupons()
			if err != nil {
				utils.GetLog().Error("coupons.Sync.fetchAllsendCoupons error: %+v", err)
			}
			for _, sc := range allsendCoupons {
				copSync := NewSendCouponSync(sc)
				if err := copSync.Do(); err != nil {
					utils.GetLog().Error("coupons.sendCouponSync.Do error: %+v", err)
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
	sendCoupon *models.SendCoupon
}

// NewSendCouponSync new send coupon sync
func NewSendCouponSync(sc *models.SendCoupon) *SendCouponSync {
	s := new(SendCouponSync)
	s.sendCoupon = sc
	return s
}

// Do do the dirty work
func (s *SendCouponSync) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("coupons.sendCouponSync.Do error: %+v", err)
		}
	}()

	if err = s.tryClose(); err != nil {
		return err
	}

	if err = s.notify(); err != nil {
		return err
	}

	return
}

func (s *SendCouponSync) notify() error {
	return notifier.Publish(s)
}

func (s *SendCouponSync) tryClose() error {
	return s.sendCoupon.TryClose()
}

// Topic publish topic
func (s *SendCouponSync) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), s.sendCoupon.BroadcastID)
}

// Message publish message
func (s *SendCouponSync) Message() []byte {
	var msg []byte
	sendCouponMsg := queues.MessageSendCouponUpdate{
		SendCouponID: s.sendCoupon.GetID(),
		BroadcastID:  s.sendCoupon.BroadcastID,
		RemainAmount: s.sendCoupon.Number,
		RemainTime:   s.sendCoupon.RemainTime(),
		CouponID:     s.sendCoupon.Coupon.GetID(),
		Description:  s.sendCoupon.Coupon.Description,
		Image:        s.sendCoupon.Coupon.Image,
		Name:         s.sendCoupon.Coupon.Name,
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
