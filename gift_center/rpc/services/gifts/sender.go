package gifts

import (
	"platform/commons/codes"
	"platform/utils"
)

// SenderConfig send config
type SenderConfig struct {
	UserID string
	GiftID string
}

// Sender send gift
type Sender struct {
	UserID string
	GiftID string

	errorCode codes.ErrorCode
}

// NewSender createa a new sender
func NewSender(c *SenderConfig) *Sender {
	s := new(Sender)
	return s
}

// ErrorCode error code
func (s *Sender) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do do the dirty world
func (s *Sender) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("gifts.Sender.Do error: %+v", err)
		}
	}()

	return
}
