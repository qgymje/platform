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

type BroadcastStarter struct {
	userID        string
	broadcastRoom *models.BroadcastRoom

	valid     bool
	errorCode codes.ErrorCode
}

func NewBroadcastStarter(userID string) *BroadcastStarter {
	return &BroadcastStarter{
		userID:        userID,
		broadcastRoom: &models.BroadcastRoom{},
	}
}

func (b *BroadcastStarter) ErrorCode() codes.ErrorCode {
	return b.errorCode
}
func (b *BroadcastStarter) Start() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("rooms.Start error: %v", err)
		}
	}()

	// 验证
	if err = b.validUser(); err != nil {
		b.errorCode = codes.ErrorCodeInvalidBroadcastringUser
		return
	}

	if err = b.startPlay(); err != nil {
		b.errorCode = codes.ErrorCodeBroadcastRoomUpdate
		return
	}

	if err = b.notify(); err != nil {
		b.errorCode = codes.ErrorCodeBroadcastNotify
		return
	}

	return
}

func (b *BroadcastStarter) GetRoomID() (string, error) {
	if b.valid {
		return b.broadcastRoom.ID.Hex(), nil
	}
	return "", errors.New("unvalid process")
}

func (b *BroadcastStarter) validUser() error {
	var err error
	b.broadcastRoom, err = models.FindBroadcastRoomByUserID(b.userID)
	if err != nil {
		return err
	}
	b.valid = true
	return nil
}

// 判断用户是否在权限
func (b *BroadcastStarter) validRoomAuth() error {
	return nil
}

func (b *BroadcastStarter) stopPlay() error {
	return b.broadcastRoom.End()
}

func (b *BroadcastStarter) startPlay() error {
	return b.broadcastRoom.Start()
}

func (b *BroadcastStarter) Topic() string {
	return queues.TopicBroadcastStart.String()
}

func (b *BroadcastStarter) Message() []byte {
	var data []byte
	msg := queues.MessageBroadcastStart{
		RoomID:    b.broadcastRoom.GetID(),
		StartTime: time.Now(),
	}
	data, _ = json.Marshal(msg)
	utils.Dump("start msg:", data)
	return data
}

func (b *BroadcastStarter) notify() error {
	return Publish(b)
}

// 启动一个进程去收消息
func (b *BroadcastStarter) startConsumer() (err error) {
	// b.GetRoomID()
	// new 一个对象
	return
}
