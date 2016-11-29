package chats

import (
	"platform/chat_center/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

// Config chats config
type Config struct {
	UserID   string
	PageNum  int
	PageSize int
}

// Chats chat list
type Chats struct {
	config     *Config
	chatFinder *models.ChatFinder
	chatList   []*ChatInfo
	errorCode  codes.ErrorCode
}

// NewChats create a new chats
func NewChats(conf *Config) *Chats {
	c := new(Chats)
	c.config = conf
	c.chatList = []*ChatInfo{}
	c.chatFinder = models.NewChatFinder().Limit(c.config.PageNum, c.config.PageSize)
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

	if err = c.findChats(); err != nil {
		c.errorCode = codes.ErrorCodeChatNotExists
		return
	}

	return
}

// Result get chat list
func (c *Chats) Result() []*ChatInfo {
	modelChats := c.chatFinder.Result()
	for i := range modelChats {
		chatInfo := modelChatToSrvChat(modelChats[i])
		c.chatList = append(c.chatList, chatInfo)
	}
	return c.chatList
}

// Count total count
func (c *Chats) Count() int64 {
	return c.chatFinder.Count()
}

func (c *Chats) findChats() (err error) {
	return c.chatFinder.Do()
}
