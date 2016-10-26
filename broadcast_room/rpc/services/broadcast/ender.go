package broadcasts

import (
	"encoding/json"
	"errors"
	"time"

	"platform/broadcast_room/rpc/models"
	"platform/broadcast_room/rpc/services/notifier"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/utils"
)

const broadcastAtLeastTime = 2 * time.Minute

// EnderConfig ender config
type EnderConfig struct {
	UserID string
	TypeID int
}

// Ender end a broadcast
type Ender struct {
	config         *EnderConfig
	roomModel      *models.Room
	broadcastModel *models.Broadcast
	valid          bool

	errorCode codes.ErrorCode
}

// NewEnder end a new broadcast
func NewEnder(c *EnderConfig) *Ender {
	return &Ender{
		config:         c,
		roomModel:      &models.Room{},
		broadcastModel: &models.Broadcast{},
	}
}

// ErrorCode inplemnt ErrorCoder
func (e *Ender) ErrorCode() codes.ErrorCode {
	return e.errorCode
}

// Do  end the broadcast
func (e *Ender) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("broadcasts.Ender.Do error: %+v", err)
		}
	}()

	if err = e.validUser(); err != nil {
		e.errorCode = codes.ErrorCodeInvalidBroadcastringUser
		return
	}
	if err = e.validBroadcast(); err != nil {
		e.errorCode = codes.ErrorCodeInvalidBroadcastringUser
		return
	}

	if yes := e.isPlayedLognerThanLeastTime(); !yes {
		e.errorCode = codes.ErrorCodeBroadcastTooShort
		return errors.New("broadcast too short")
	}

	if err = e.stopPlay(); err != nil {
		e.errorCode = codes.ErrorCodeRoomUpdate
		return
	}

	if err = e.update(); err != nil {
		e.errorCode = codes.ErrorCodeBroadcastUpdate
	}

	if err = e.notify(); err != nil {
		e.errorCode = codes.ErrorCodeBroadcastNotify
		return
	}

	return
}

// GetBroadcast get broadcast info
func (e *Ender) GetBroadcast() (*Broadcast, error) {
	if !e.valid {
		return nil, errors.New("ender: unvalid process")
	}

	srvBro := modelBroadcastToSrvBroadcast(e.broadcastModel)
	return srvBro, nil
}

func (e *Ender) validUser() error {
	var err error
	e.roomModel, err = models.FindRoomByUserID(e.config.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (e *Ender) validBroadcast() error {
	var err error
	e.broadcastModel, err = models.FindBroadcastByRoomID(e.roomModel.GetID())
	if err != nil {
		return err
	}
	e.valid = true
	return nil
}

func (e *Ender) removeBroadcast() error {
	// delete broadcast topic
	return nil
}

func (e *Ender) isPlayedLognerThanLeastTime() bool {
	return time.Since(e.broadcastModel.StartTime) > broadcastAtLeastTime
}

func (e *Ender) stopPlay() error {
	if err := e.roomModel.EndPlaying(e.broadcastModel); err != nil {
		return err
	}
	return nil
}

func (e *Ender) update() error {
	return e.broadcastModel.End()
}

// Topic implement Notifier interface
func (e *Ender) Topic() string {
	return queues.TopicBroadcastEnd.String()
}

// Message implement Notifier interface
func (e *Ender) Message() []byte {
	var msg []byte
	broadcastEndMsg := queues.MessageBroadcastEnd{
		RoomID:      e.roomModel.GetID(),
		BroadcastID: e.broadcastModel.GetID(),
		EndTime:     time.Now().Unix(),
	}

	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		e.config.TypeID,
		broadcastEndMsg,
	}

	msg, _ = json.Marshal(data)
	return msg
}

func (e *Ender) notify() error {
	return notifier.Publish(e)
}
