package services

import (
	"fmt"
	"platform/commons/queues"
	"platform/notification_center/http/services/receiver"
)

// Chat chat consumer
type Chat struct {
	userID   string
	receiver *receiver.Receive
}

// NewChat new chat consumer
func NewChat() *Chat {
	c := new(Chat)
	return c
}

// Topic topic name
func (c *Chat) Topic() string {
	return fmt.Sprintf(queues.TopicChatFormat.String(), c.userID)
}

// Channel channel name
func (c *Chat) Channel() string {
	return c.userID
}

// Handler handler
func (c *Chat) Handler(<-chan []byte) {

}

// Stop stop consume
func (c *Chat) Stop() error {
	return nil
}
