package broadcasts

import (
	"encoding/json"
	"errors"
	"fmt"
	"platform/broadcast_room/rpc/models"
	"platform/broadcast_room/rpc/services/notifier"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/utils"
)

// StarterConfig starter config
type StarterConfig struct {
	UserID string
	TypeID int
}

// Starter start a broadcast process wrapper
type Starter struct {
	config         *StarterConfig
	roomModel      *models.Room
	broadcastModel *models.Broadcast

	valid     bool
	errorCode codes.ErrorCode
}

// NewStarter create a Starter
func NewStarter(c *StarterConfig) *Starter {
	return &Starter{
		config:         c,
		roomModel:      &models.Room{},
		broadcastModel: &models.Broadcast{},
	}
}

// ErrorCode implement ErrorCoder
func (s *Starter) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do do the dirty job
func (s *Starter) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("broadcasts.Starter.Do error: %+v", err)
		}
	}()

	if err = s.validUser(); err != nil {
		s.errorCode = codes.ErrorCodeInvalidBroadcastringUser
		return
	}

	if yes := s.isPlaying(); yes {
		s.errorCode = codes.ErrorCodeBroadcastIsOn
		return errors.New("boradcast is on")
	}

	if err = s.save(); err != nil {
		s.errorCode = codes.ErrorCodeBroadcastCreate
	}

	if err = s.startPlay(); err != nil {
		s.errorCode = codes.ErrorCodeRoomUpdate
		return
	}

	if err = s.notify(); err != nil {
		s.errorCode = codes.ErrorCodeBroadcastNotify
		return
	}

	return
}

// GetBroadcast get broadcast info
func (s *Starter) GetBroadcast() (*Broadcast, error) {
	if !s.valid {
		return nil, errors.New("starter: unvalid process")
	}
	srvBro := modelBroadcastToSrvBroadcast(s.broadcastModel)
	return srvBro, nil
}

func (s *Starter) validUser() error {
	var err error
	s.roomModel, err = models.FindRoomByUserID(s.config.UserID)
	if err != nil {
		return err
	}
	s.valid = true
	return nil
}

func (s *Starter) isPlaying() bool {
	return s.roomModel.IsPlaying && s.roomModel.BroadcastID != nil
}

func (s *Starter) save() error {
	s.broadcastModel.RoomID = s.roomModel.RoomID
	return s.broadcastModel.Create()
}

func (s *Starter) startPlay() error {
	if err := s.roomModel.StartPlaying(s.broadcastModel); err != nil {
		return err
	}
	return nil
}

// Topic topic
func (s *Starter) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), s.broadcastModel.GetID())
}

// Message publish message
func (s *Starter) Message() []byte {
	var data []byte
	msg := queues.MessageBroadcastStart{
		RoomID:      s.roomModel.GetID(),
		BroadcastID: s.broadcastModel.GetID(),
		StartTime:   s.broadcastModel.StartTime.Unix(),
	}
	data, _ = json.Marshal(msg)
	return data
}

func (s *Starter) notify() error {
	return notifier.Publish(s)
}
