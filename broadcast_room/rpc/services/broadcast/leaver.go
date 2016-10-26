package broadcasts

import (
	"encoding/json"
	"errors"
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/coupon_center/rpc/services/notifier"
	"platform/utils"
)

// LeaverConfig leaver config
type LeaverConfig struct {
	TypeID      int
	BroadcastID string
	UserID      string
	Username    string
	Level       int64
}

// Leaver when a user enter a broadcast
type Leaver struct {
	config         *LeaverConfig
	audiecntModel  *models.Audience
	broadcastModel *models.Broadcast

	errorCode codes.ErrorCode
}

// NewLeaver create a new leaver
func NewLeaver(c *LeaverConfig) *Leaver {
	l := new(Leaver)
	l.config = c
	l.broadcastModel = &models.Broadcast{}
	return l
}

// ErrorCode error code
func (l *Leaver) ErrorCode() codes.ErrorCode {
	return l.errorCode
}

// Do do the dirty work
func (l *Leaver) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("broadcast.Leaver.Do error: %+v", err)
		}
	}()

	if yes := l.isBroadcastClosed(); yes {
		l.errorCode = codes.ErrorCodeBroadcastClosed
		return errors.New("broadcast is closed")
	}

	if err = l.update(); err != nil {
		return
	}

	if err = l.notify(); err != nil {
		l.errorCode = codes.ErrorCodeBroadcastNotify
		return
	}
	return nil
}

func (l *Leaver) isBroadcastClosed() bool {
	var err error
	l.broadcastModel, err = models.FindBroadcastByID(l.config.BroadcastID)
	if l.broadcastModel != nil && err == nil {
		return !l.broadcastModel.IsPlaying()
	}
	return true
}

func (l *Leaver) update() (err error) {
	if l.audiecntModel, err = models.NewAudience(l.config.BroadcastID, l.config.UserID); err != nil {
		return err
	}

	if err = l.audiecntModel.Leave(); err != nil {
		l.errorCode = codes.ErrorCodeAudienceUpdate
		return err
	}

	total, current := 0, -1
	if err = l.updateBroadcast(total, current); err != nil {
		l.errorCode = codes.ErrorCodeBroadcastUpdate
		return err
	}

	return
}

func (l *Leaver) updateBroadcast(total, current int) (err error) {
	return l.broadcastModel.AddAudience(total, current)
}

func (l *Leaver) notify() (err error) {
	return notifier.Publish(l)
}

// Topic topic
func (l *Leaver) Topic() string {
	return queues.TopicBroadcastLeave.String()
}

// Message message
func (l *Leaver) Message() []byte {
	var msg []byte
	broadcastLeaveMsg := queues.MessageBroadcastLeave{
		BroadcastID: l.broadcastModel.GetID(),
		UserID:      l.config.UserID,
		Username:    l.config.Username,
		Level:       l.config.Level,
		LeaveTime:   l.audiecntModel.LeaveTime.Unix(),
	}

	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		l.config.TypeID,
		broadcastLeaveMsg,
	}

	msg, _ = json.Marshal(data)
	return msg
}
