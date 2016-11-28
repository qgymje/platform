package chats

import (
	"errors"
	"platform/chat_center/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

var (
	// ErrChatExists chat already exists
	ErrChatExists = errors.New("chat already exists")
)

// CreatorConfig  config
type CreatorConfig struct {
	Name    string
	UserID  string
	Members []string
}

// Creator create a chat
type Creator struct {
	config    *CreatorConfig
	chatModel *models.Chat

	hasAppendCreatorID bool
	errorCode          codes.ErrorCode
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

	if yes := c.isChatExists(); yes {
		c.errorCode = codes.ErrorCodeChatAlreadyExists
		return ErrChatExists
	}

	if err = c.save(); err != nil {
		c.errorCode = codes.ErrorCodeChatCreate
		return
	}
	return
}

// GetChatID get chat id
func (c *Creator) GetChatID() string {
	return c.chatModel.GetID()
}

func (c *Creator) members() []string {
	if !c.hasAppendCreatorID {
		c.config.Members = append(c.config.Members, c.config.UserID)
		c.hasAppendCreatorID = true
	}
	return c.config.Members
}

func (c *Creator) findChat() (err error) {
	c.chatModel.Sign = c.chatModel.GenSign(c.members())
	return c.chatModel.FindBySign()
}

func (c *Creator) isChatExists() bool {
	if err := c.findChat(); err != nil {
		return false
	}
	return true
}

func (c *Creator) save() error {
	c.chatModel.Name = c.config.Name
	c.chatModel.UserID = c.config.UserID
	return c.chatModel.Create(c.members())
}
