package services

import (
	"fmt"
	"platform/commons/queues"
)

// Notification chat consumer
type Notification struct {
	userID string
}

// NewNotification new chat consumer
func NewNotification() *Notification {
	n := new(Notification)
	return n
}

// Topic topic name
func (n *Notification) Topic() string {
	return fmt.Sprintf(queues.TopicNotificationFormat.String(), n.userID)
}

// Channel channel name
func (n *Notification) Channel() string {
	return n.userID
}

// Handler handler
func (n *Notification) Handler(<-chan []byte) {

}

// Stop stop consume
func (n *Notification) Stop() error {
	return nil
}
