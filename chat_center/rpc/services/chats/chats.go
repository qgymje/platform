package chats

import (
	"platform/commons/codes"
	"platform/utils"
)

// Config chats config
type Config struct {
	UserID string
}

// Chats chat list
type Chats struct {
	config    *Config
	errorCode codes.ErrorCode
}

// NewChats create a new chats
func NewChats(conf *Config) *Chats {
	c := new(Chats)
	c.config = conf
	return c
}

// ErrorCode error code
func (c *Chats) ErrorCode() codes.ErrorCode {
	return c.errorCode
}

// Do do the main job
func (c *Chats) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("chats.Chats.Do error:%+v", err)
		}
	}()
	return
}
