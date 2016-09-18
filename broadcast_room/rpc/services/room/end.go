package rooms

import (
	"encoding/json"
	"errors"
	"time"

	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/utils"
)

type BroadcastEnder struct {
	userID        string
	broadcastRoom *models.BroadcastRoom

	valid     bool
	errorCode codes.ErrorCode
}

func NewBroadcastEnder(userID string) *BroadcastEnder {
	return &BroadcastEnder{
		userID:        userID,
		broadcastRoom: &models.BroadcastRoom{},
	}
}

func (b *BroadcastEnder) ErrorCode() codes.ErrorCode {
	return b.errorCode
}

// End 表示结束直播
func (b *BroadcastEnder) End() (err error) {
	// 删除topic
	return
}

func (b *BroadcastEnder) GetRoomID() (string, error) {
	if b.valid {
		return b.broadcastRoom.ID.Hex(), nil
	}
	return "", errors.New("unvalid process")
}

func (b *BroadcastEnder) validUser() error {
	var err error
	b.broadcastRoom, err = models.FindBroadcastRoomByUserID(b.userID)
	if err != nil {
		return err
	}
	b.valid = true
	return nil
}

// 判断用户是否在权限
func (b *BroadcastEnder) validRoomAuth() error {
	return nil
}

func (b *BroadcastEnder) stopPlay() error {
	return b.broadcastRoom.End()
}

func (b *BroadcastEnder) startPlay() error {
	return b.broadcastRoom.Start()
}

func (b *BroadcastEnder) Topic() string {
	return queues.TopicBroadcastStart.String()
}

func (b *BroadcastEnder) Message() []byte {
	var data []byte
	msg := queues.MessageBroadcastEnd{
		RoomID:  b.broadcastRoom.GetID(),
		EndTime: time.Now(),
	}
	data, _ = json.Marshal(msg)
	utils.Dump("start msg:", data)
	return data
}

func (b *BroadcastEnder) notify() error {
	return Publish(b)
}

// 启动一个进程去收消息
func (b *BroadcastEnder) startConsumer() (err error) {
	// b.GetRoomID()
	// new 一个对象
	return
}
