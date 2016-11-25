package chats

import (
	"platform/commons/codes"
	"platform/utils"
)

// SenderConfig sender config
type SenderConfig struct {
	ChatID  string
	UserID  string
	Content string
}

// Sender sender
type Sender struct {
	config    *SenderConfig
	errorCode codes.ErrorCode
}

// NewSender create a new sender
func NewSender(c *SenderConfig) *Sender {
	s := new(Sender)
	s.config = c
	return s
}

// ErrorCode error code
func (s *Sender) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do do the main job
func (s *Sender) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("chats.Sender.Do error:%+v", err)
		}
	}()
	return
}
