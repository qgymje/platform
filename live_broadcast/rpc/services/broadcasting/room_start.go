package broadcastings

import (
	"encoding/json"

	"tech.cloudzen/commons"
	"tech.cloudzen/utils"
)

func StartToReceive() {
	room := &RoomStart{}
	room.receive()
}

// RoomStart 弹幕服务
type RoomStart struct{}

func (s *RoomStart) Topic() string {
	return commons.TopicBroadcastStart.String()
}

func (s *RoomStart) Channel() string {
	return commons.ChannelBroadcastStart.String()
}

func (s *RoomStart) handleMessage(msg []byte) (*commons.MessageBroadcastStart, error) {
	var msgBroStart commons.MessageBroadcastStart
	if err := json.Unmarshal(msg, &msgBroStart); err != nil {
		return nil, err
	}
	return &msgBroStart, nil
}

func (s *RoomStart) Handler(msgs <-chan []byte) {
	for msg := range msgs {
		msgBroStart, err := s.handleMessage(msg)
		if err != nil {
			utils.GetLog().Error("parse broadcast start msg error: %v", err)
		} else {
			utils.GetLog().Info("new broadcasting start: %v", msgBroStart)
			barrage := NewBarrageReceive(msgBroStart.RoomID)
			if err := barrage.Receive(); err != nil {
				utils.GetLog().Error("receive broadcast room error: %v", err)
			}
		}
	}
}

func (s *RoomStart) receive() error {
	return NewReceive(s).Do()
}
