package chats

import (
	"fmt"
	"platform/chat_center/rpc/models"
	"platform/commons/codes"
	"platform/commons/queues"
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
	return
}

func (s *Sender) save() (err error) {
	return
}

func (s *Sender) notify() error {
	return nil
}

// Topic topic
func (s *Sender) Topic() string {
	return fmt.Sprintf(queues.TopicChatFormat.String(), s.config.ChatID)
}

// Message message send to the nsq
func (s *Sender) Message() []byte {
	return nil
}
