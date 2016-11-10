package gifts

import (
	"fmt"
	"platform/commons/queues"
)

// Broadcast represents a broadcast
type Broadcast struct {
	broadcastID string
}

// Topic topic
func (b *Broadcast) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), b.broadcastID)
}

// Message publish message
func (b *Broadcast) Message() []byte {
	// fetch from redis
	return nil
}
