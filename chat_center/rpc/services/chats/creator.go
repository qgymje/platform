package chats

import (
	"platform/chat_center/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

// CreatorConfig  config
type CreatorConfig struct {
	Name    string
	UserID  string
	Members []string
}

// Creator create a chat
type Creator struct {
	config   *CreatorConfig
	chatMode *models.Chat

	errorCode codes.ErrorCode
}

// NewCreator create a new Creator
func NewCreator(conf *CreatorConfig) *Creator {
	c := new(Creator)
	c.config = conf
	c.chatModel = &models.Chat{}
	return c
}

// ErrorCode error code
func (c *Creator) ErrorCode() codes.ErrorCode {
	return c.errorCode
}

// Do do the main job
func (c *Creator) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("chats.Creator.Do error:%+v", err)
		}
	}()
	return
}

func (c *Creator) findChat() (err error) {
	return
}
