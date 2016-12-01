package services

import (
	"fmt"
	"platform/commons/queues"
	"platform/notification_center/http/services/receiver"
)

// Notification chat consumer
type Notification struct {
	userID   string
	receiver *receiver.Receive
	msg      chan Message
}

// NewNotification new chat consumer
func NewNotification(userID string) *Notification {
	n := new(Notification)
	n.userID = userID
	n.receiver = receiver.NewReceive(n)
	n.msg = make(chan Message)
	return n
}

// Topic topic name
func (n *Notification) Topic() string {
	return fmt.Sprintf(queues.TopicUserFormat.String(), n.userID)
}

// Channel channel name
func (n *Notification) Channel() string {
	return queues.ChannelUserNotification.String()
}

// Handler handler
func (n *Notification) Handler(msg <-chan []byte) {
	for v := range msg {
		//log.Println("[notification] got a message: ", string(v))
		n.msg <- Message(v)
	}
}

// Consume comsume msg
func (n *Notification) Consume() <-chan Message {
	go n.receiver.Do()
	return n.msg
}

// Stop stop consume
func (n *Notification) Stop() {
	n.receiver.Stop()
}
