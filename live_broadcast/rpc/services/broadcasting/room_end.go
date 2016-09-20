package broadcastings

import (
	"log"

	"platform/commons/queues"
)

type RoomEnd struct{}

func (s *RoomEnd) Topic() string {
	return queues.TopicBroadcastEnd.String()
}

func (s *RoomEnd) Channel() string {
	return queues.ChannelBroadcastEnd.String()
}

func (s *RoomEnd) Handler(msgs <-chan []byte) {
	for msg := range msgs {
		log.Println("receive new broadcast end")
		log.Println(string(msg))
	}
}

func (s *RoomEnd) receive() error {
	return NewReceive(s).Do()
}
