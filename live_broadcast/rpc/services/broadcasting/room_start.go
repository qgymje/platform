package broadcastings

import (
	"encoding/json"

	"platform/commons/queues"
	"platform/utils"
)

func StartToReceive() {
	go func() {
		(&RoomStart{}).receive()
	}()
}

// RoomStart 弹幕服务
type RoomStart struct{}

func (s *RoomStart) Topic() string {
	return queues.TopicBroadcastStart.String()
}

func (s *RoomStart) Channel() string {
	return queues.ChannelBroadcastStart.String()
}

func (s *RoomStart) handleMessage(msg []byte) (*queues.MessageBroadcastStart, error) {
	var msgBroStart queues.MessageBroadcastStart
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
