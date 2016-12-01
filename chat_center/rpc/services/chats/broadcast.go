package chats

import (
	"encoding/json"
	"fmt"
	"platform/chat_center/rpc/models"
	"platform/chat_center/rpc/services/chats/notifier"
	"platform/commons/queues"
	"platform/commons/typeids"
	"platform/utils"
)

// BroadcastConfig broadcast config
type BroadcastConfig struct {
	ToUserID string
	Message  *models.Message
}

// Broadcast object
type Broadcast struct {
	config *BroadcastConfig
}

// NewBroadcast new broadcast
func NewBroadcast(c *BroadcastConfig) *Broadcast {
	b := new(Broadcast)
	b.config = c
	return b
}

// Do the notify work
func (b *Broadcast) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("chats.Broadcast.Do error:%+v", err)
		}
	}()

	if err = b.notify(); err != nil {
		return
	}
	return
}

func (b *Broadcast) notify() error {
	return notifier.Publish(b)
}

// Topic topic
func (b *Broadcast) Topic() string {
	return fmt.Sprintf(queues.TopicUserFormat.String(), b.config.ToUserID)
}

// Message message send to the nsq
func (b *Broadcast) Message() []byte {
	msg := queues.MessageChatMessage{
		MessageID: b.config.Message.GetID(),
		ChatID:    b.config.Message.Chat.GetID(),
		UserID:    b.config.Message.UserID,
		Content:   b.config.Message.Content,
		SendTime:  b.config.Message.CreatedAt.Unix(),
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
