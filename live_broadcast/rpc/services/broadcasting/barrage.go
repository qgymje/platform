package broadcastings

import (
	"encoding/json"
	"fmt"

	"platform/commons/queues"
	"platform/live_broadcast/rpc/models"
	"platform/utils"
)

const MaxListLenToSave = 100

// 用于处理具体一个room的弹幕消息
type Barrage struct {
	roomID string
	list   []*models.Barrage
}

func NewBarrageReceive(roomID string) *Barrage {
	return &Barrage{
		roomID: roomID,
		list:   []*models.Barrage{},
	}
}

// 将收到的消息保存起来
func (b *Barrage) Save() (err error) {
	return models.Save(b.list)
}

func (b *Barrage) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastRoomFormat.String(), b.roomID)
}

func (b *Barrage) Channel() string {
	return queues.ChannelBroadcastBarrage.String()
}

func (b *Barrage) handleMessage(msg []byte) (*models.Barrage, error) {
	barrageMsg := queues.MessageBarrage{}
	if err := json.Unmarshal(msg, &barrageMsg); err != nil {
		return nil, err
	}
	roomIDObjectID, _ := models.StringToObjectID(b.roomID)
	userIDObjectID, _ := models.StringToObjectID(barrageMsg.UserID)
	barrage := models.Barrage{
		RoomID:    roomIDObjectID,
		UserID:    userIDObjectID,
		UserName:  barrageMsg.UserName,
		Message:   barrageMsg.Message,
		CreatedAt: barrageMsg.PubTime,
	}
	return &barrage, nil
}

func (b *Barrage) Handler(msgs <-chan []byte) {
	for msg := range msgs {
		barrage, err := b.handleMessage(msg)
		if err != nil {
			utils.GetLog().Error("receive msg parse error", err)
		}
		b.list = append(b.list, barrage)
		if len(b.list) > MaxListLenToSave {
			b.Save()
		}
	}
}

func (b *Barrage) Receive() error {
	return NewReceive(b).Do()
}
