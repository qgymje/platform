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

// Starter start a broadcast process wrapper
type Starter struct {
	userID         string
	modelRoom      *models.Room
	modelBroadcast *models.Broadcast

	valid     bool
	errorCode codes.ErrorCode
}

// NewStarter create a Starter
func NewStarter(userID string) *Starter {
	return &Starter{
		userID:         userID,
		modelRoom:      &models.Room{},
		modelBroadcast: &models.Broadcast{},
	}
}

// ErrorCode implement ErrorCoder
func (b *Starter) ErrorCode() codes.ErrorCode {
	return b.errorCode
}

// Do do the dirty job
func (b *Starter) Do() (err error) {
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

// GetBroadcastID get broadcast id
func (b *Starter) GetBroadcastID() (string, error) {
	if b.valid {
		return b.modelBroadcast.GetID(), nil
	}
	return "", errors.New("unvalid process")
}

func (b *Starter) validUser() error {
	var err error
	b.modelRoom, err = models.FindRoomByUserID(b.userID)
	if err != nil {
		return err
	}
	b.valid = true
	return nil
}

func (b *Starter) validRoomAuth() error {
	return nil
}

func (b *Starter) stopPlay() error {
	return b.modelRoom.EndPlaying()
}

func (b *Starter) startPlay() error {
	return b.modelRoom.StartPlaying()
}

// Topic topic
func (b *Starter) Topic() string {
	return queues.TopicBroadcastStart.String()
}

// Message publish message
func (b *Starter) Message() []byte {
	var data []byte
	msg := queues.MessageBroadcastStart{
		RoomID:      b.modelRoom.GetID(),
		BroadcastID: b.modelBroadcast.GetID(),
		StartTime:   time.Now(),
	}
	data, _ = json.Marshal(msg)
	utils.Dump("start msg:", data)
	return data
}

func (b *Starter) notify() error {
	return notifier.Publish(b)
}
