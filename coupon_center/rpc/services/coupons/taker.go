package coupons

import "platform/commons/codes"

// TakerConfig config
type TakerConfig struct {
}

// Taker user take a coupon
type Taker struct {
	errorCode codes.ErrorCode
}

// NewTaker new taker
func NewTaker(c *TakerConfig) *Taker {
	t := new(Taker)
	return t
}

// ErrorCode error code
func (t *Taker) ErrorCode() codes.ErrorCode {
	return t.errorCode
}
