package barrages

import (
	"encoding/json"
	"fmt"
	"platform/barrage_center/rpc/models"
	"platform/barrage_center/rpc/services/notifier"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/utils"
	"time"
)

// SenderConfig sender config
type SenderConfig struct {
	TypeID      int32
	BroadcastID string
	UserID      string
	Username    string
	Level       int64
	Text        string
	CreatedAt   int64
}

// Sender send a barrage
type Sender struct {
	config       *SenderConfig
	barrageModel *models.Barrage

	errorCode codes.ErrorCode
}

// NewSender create a new Sender
func NewSender(c *SenderConfig) *Sender {
	s := new(Sender)
	s.config = c
	s.barrageModel = &models.Barrage{}
	return s
}

// ErrorCode errorCode
func (s *Sender) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do the dirty job
func (s *Sender) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("rpc.Sender.Do error: %+v", err)
		}
	}()

	if err = s.preSave(); err != nil {
		s.errorCode = codes.ErrorCodeBarrageCreate
		return
	}

	if err = s.save(); err != nil {
		s.errorCode = codes.ErrorCodeBarrageCreate
		return
	}

	if err = s.notify(); err != nil {
		s.errorCode = codes.ErrorCodeBarrageNotify
		return
	}

	return
}

func (s *Sender) preSave() error {
	broadcastObjID, err := models.StringToObjectID(s.config.BroadcastID)
	userObjID, err := models.StringToObjectID(s.config.UserID)
	if err != nil {
		return err
	}

	s.barrageModel.BroadcastID = broadcastObjID
	s.barrageModel.Username = s.config.Username
	s.barrageModel.Text = s.config.Text
	s.barrageModel.UserID = userObjID
	s.barrageModel.Level = s.config.Level
	s.barrageModel.CreatedAt = time.Unix(s.config.CreatedAt, 0)

	return nil
}

func (s *Sender) save() (err error) {
	return s.barrageModel.Create()
}

func (s *Sender) notify() (err error) {
	return notifier.DeferredPublish(s, 50*time.Microsecond)
}

// Topic publish topic
func (s *Sender) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), s.config.BroadcastID)
}

// Message publish message
func (s *Sender) Message() []byte {
	var msg []byte
	barrageMsg := queues.MessageBarrage{
		BroadcastID: s.config.BroadcastID,
		UserID:      s.config.UserID,
		Username:    s.config.Username,
		Level:       s.config.Level,
		Text:        s.config.Text,
		CreatedAt:   s.barrageModel.CreatedAt.Unix(),
	}

	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		int(s.config.TypeID),
		barrageMsg,
	}
	msg, _ = json.Marshal(data)
	return msg
}
