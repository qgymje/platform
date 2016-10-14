package rooms

import (
	"encoding/json"
	"errors"
	"time"

	"platform/account_center/rpc/services/notifier"
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/utils"
)

type Ender struct {
	userID        string
	broadcastRoom *models.Room

	valid     bool
	errorCode codes.ErrorCode
}

func NewEnder(userID string) *Ender {
	return &Ender{
		userID:        userID,
		broadcastRoom: &models.Room{},
	}
}

func (b *Ender) ErrorCode() codes.ErrorCode {
	return b.errorCode
}

// End 表示结束直播
func (b *Ender) End() (err error) {
	// 删除topic
	return
}

func (b *Ender) GetRoomID() (string, error) {
	if b.valid {
		return b.broadcastRoom.GetID(), nil
	}
	return "", errors.New("unvalid process")
}

func (b *Ender) validUser() error {
	var err error
	b.broadcastRoom, err = models.FindRoomByUserID(b.userID)
	if err != nil {
		return err
	}
	b.valid = true
	return nil
}

// 判断用户是否在权限
func (b *Ender) validRoomAuth() error {
	return nil
}

func (b *Ender) stopPlay() error {
	return b.broadcastRoom.EndPlaying()
}

func (b *Ender) startPlay() error {
	return b.broadcastRoom.StartPlaying()
}

func (b *Ender) Topic() string {
	return queues.TopicBroadcastStart.String()
}

func (b *Ender) Message() []byte {
	var data []byte
	msg := queues.MessageBroadcastEnd{
		RoomID:  b.broadcastRoom.GetID(),
		EndTime: time.Now(),
	}
	data, _ = json.Marshal(msg)
	utils.Dump("start msg:", data)
	return data
}

func (b *Ender) notify() error {
	return notifier.Publish(b)
}
