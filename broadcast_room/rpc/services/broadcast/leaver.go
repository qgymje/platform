package broadcasts

import (
	"errors"
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

// LeaverConfig leaver config
type LeaverConfig struct {
	UserID, BroadcastID string
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
		return err
	}

	total, current := 0, -1
	if err = l.updateBroadcast(total, current); err != nil {
		return err
	}

	return
}

func (l *Leaver) updateBroadcast(total, current int) (err error) {
	return l.broadcastModel.AddAudience(total, current)
}
