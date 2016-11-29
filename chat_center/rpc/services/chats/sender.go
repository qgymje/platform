package chats

import (
	"encoding/json"
	"errors"
	"fmt"
	"platform/chat_center/rpc/models"
	"platform/chat_center/rpc/services/chats/notifier"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/commons/typeids"
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
	return notifier.Publish(s)
}

// Topic topic
func (s *Sender) Topic() string {
	return fmt.Sprintf(queues.TopicChatFormat.String(), s.config.ChatID)
}

// Message message send to the nsq
func (s *Sender) Message() []byte {
	msg := queues.MessageChatMessage{
		MessageID: s.messageModel.GetID(),
		ChatID:    s.messageModel.Chat.GetID(),
		UserID:    s.config.UserID,
		Content:   s.messageModel.Content,
		SendTime:  s.messageModel.CreatedAt.Unix(),
	}

	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		Type: int(typeids.ChatMessage),
		Data: msg,
	}

	m, _ := json.Marshal(&data)
	return m
}
