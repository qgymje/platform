package broadcasts

import (
	"errors"
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

// EnterConfig enter config
type EnterConfig struct {
	UserID, BroadcastID string
}

// Enter when a user enter a broadcast
type Enter struct {
	config         *EnterConfig
	audiecntModel  *models.Audience
	broadcastModel *models.Broadcast

	errorCode codes.ErrorCode
}

// NewEnter create a new enter object
func NewEnter(c *EnterConfig) *Enter {
	e := new(Enter)
	e.config = c
	e.broadcastModel = &models.Broadcast{}
	return e
}

// ErrorCode error code
func (e *Enter) ErrorCode() codes.ErrorCode {
	return e.errorCode
}

// Do do the dirty work
func (e *Enter) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("broadcast.Enter.Do error: %+v", err)
		}
	}()

	if yes := e.isBroadcastOn(); !yes {
		e.errorCode = codes.ErrorCodeBroadcastClosed
		return errors.New("broadcast is closed")
	}

	if err = e.save(); err != nil {
		return err
	}

	return
}

func (e *Enter) isBroadcastOn() bool {
	var err error
	e.broadcastModel, err = models.FindBroadcastByID(e.config.BroadcastID)
	if e.broadcastModel != nil && err == nil {
		return e.broadcastModel.IsPlaying()
	}
	return false
}

func (e *Enter) save() (err error) {
	if e.audiecntModel, err = models.NewAudience(e.config.BroadcastID, e.config.UserID); err != nil {
		return err
	}

	total, current := 1, 1
	if yes := e.audiecntModel.HasEntered(); yes {
		total, current = 0, 1
	}

	if err = e.audiecntModel.Enter(); err != nil {
		return err
	}

	if err = e.updateBroadcast(total, current); err != nil {
		return err
	}

	return
}

func (e *Enter) updateBroadcast(total, current int) (err error) {
	return e.broadcastModel.AddAudience(total, current)
}
