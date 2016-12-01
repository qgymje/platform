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
	msg      chan Message
}

// NewChat new chat consumer
func NewChat(userID string) *Chat {
	c := new(Chat)
	c.userID = userID
	c.msg = make(chan Message)
	c.receiver = receiver.NewReceive(c)
	return c
}

// Topic topic name
func (c *Chat) Topic() string {
	return fmt.Sprintf(queues.TopicUserFormat.String(), c.userID)
}

// Channel channel name
func (c *Chat) Channel() string {
	return queues.ChannelUserChat.String()
}

// Handler handler
func (c *Chat) Handler(msg <-chan []byte) {
	for v := range msg {
		//log.Println("[chat] got a message: ", string(v))
		c.msg <- Message(v)
	}
}

// Consume comsume msg
func (c *Chat) Consume() <-chan Message {
	go c.receiver.Do()

	return c.msg
}

// Stop stop consume
func (c *Chat) Stop() {
	c.receiver.Stop()
}
