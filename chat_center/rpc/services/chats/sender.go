package chats

import (
	"errors"
	"platform/chat_center/rpc/models"
	"platform/commons/codes"
	"platform/utils"
	"strconv"
)

var (
	// ErrChatNotExists chat not exists
	ErrChatNotExists = errors.New("chat not exists")
)

// SenderConfig sender config
type SenderConfig struct {
	ChatID  string
	UserID  string
	Content string
}

// Sender sender
type Sender struct {
	config *SenderConfig

	chatModel    *models.Chat
	messageModel *models.Message

	errorCode codes.ErrorCode
}

// NewSender create a new sender
func NewSender(c *SenderConfig) *Sender {
	s := new(Sender)
	s.config = c
	s.chatModel = &models.Chat{}
	s.messageModel = &models.Message{}
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

	if yes := s.isChatExists(); !yes {
		s.errorCode = codes.ErrorCodeChatNotExists
		return ErrChatNotExists
	}

	if err = s.save(); err != nil {
		s.errorCode = codes.ErrorCodeChatMessageCreate
		return
	}

	if err = s.notify(); err != nil {
		s.errorCode = codes.ErrorCodeChatMessageNotify
		return
	}

	return
}

// GetMessageID get message id
func (s *Sender) GetMessageID() string {
	return s.messageModel.GetID()
}

func (s *Sender) save() (err error) {
	s.messageModel.Chat = s.chatModel
	s.messageModel.UserID = s.config.UserID
	s.messageModel.Content = s.config.Content
	return s.messageModel.Create()
}

func (s *Sender) getChatID() int64 {
	id, _ := strconv.ParseInt(s.config.ChatID, 10, 0)
	return id
}

func (s *Sender) findChat() (err error) {
	s.chatModel.ID = s.getChatID()
	return s.chatModel.FindByID()
}

func (s *Sender) isChatExists() bool {
	if err := s.findChat(); err != nil {
		return false
	}
	return true
}

func (s *Sender) notify() error {
	for _, member := range s.chatModel.Members {
		config := &BroadcastConfig{
			ToUserID: member.UserID,
			Message:  s.messageModel,
		}
		broadcast := NewBroadcast(config)
		if err := broadcast.Do(); err != nil {
			return err
		}
	}
	return nil
}
