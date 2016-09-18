package broadcastings

import (
	"log"

	"tech.cloudzen/commons"
)

type RoomEnd struct{}

func (s *RoomEnd) Topic() string {
	return commons.TopicBroadcastEnd.String()
}

func (s *RoomEnd) Channel() string {
	return commons.ChannelBroadcastEnd.String()
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
