package profiles

import (
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/profile_center/rpc/models"
	"platform/utils"
)

// WithdrawConfig withdraw config
type WithdrawConfig struct {
	UserID    string
	MsgID     string
	SnowBall  uint
	SnowFlake uint
	TypeID    uint
	TargetID  string
}

// Withdraw withdraw
type Withdraw struct {
	config *WithdrawConfig

	profileModel *models.Profile

	errorCode codes.ErrorCode
}

// NewWithdraw new withdraw
func NewWithdraw(c *WithdrawConfig) *Withdraw {
	w := new(Withdraw)
	w.config = c
	w.profileModel = &models.Profile{}
	return w
}

// ErrorCode error code
func (w *Withdraw) ErrorCode() codes.ErrorCode {
	return w.errorCode
}

// Create create a new withdraw
func (w *Withdraw) Create() (msgID int64, err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("profiles.Withdraw.Create error: %+v", err)
		}
	}()

	if err = w.findProfile(); err != nil {
		w.errorCode = codes.ErrorCodeProfileNotFound
		return
	}
	utils.Dump(w.config)
	if msgID, err = w.profileModel.Withdraw(w.config.SnowBall, w.config.SnowFlake, w.config.TypeID, w.config.TargetID); err != nil {
		switch err {
		case models.ErrNotEnoughSnowBall:
			w.errorCode = codes.ErrorCodeNotEnoughSnowBall
		case models.ErrNotEnoughSnowFlake:
			w.errorCode = codes.ErrorCodeNotEnoughSnowFlake
		default:
			w.errorCode = codes.ErrorCodeWithdraw
		}
		return
	}

	return
}

// Rollback withdraw rollback
func (w *Withdraw) Rollback() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("profiles.Withdraw.Rollback error: %+v", err)
		}
	}()

	if err = w.findProfile(); err != nil {
		w.errorCode = codes.ErrorCodeProfileNotFound
		return
	}

	if err = w.profileModel.WithdrawRollback(w.config.MsgID); err != nil {
		w.errorCode = codes.ErrorCodeWithdrawRollback
		return
	}

	return
}

// Commit withdraw commit
func (w *Withdraw) Commit() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("profiles.Withdraw.Commit error: %+v", err)
		}
	}()

	if err = w.findProfile(); err != nil {
		w.errorCode = codes.ErrorCodeProfileNotFound
		return
	}

	if err = w.profileModel.WithdrawCommit(w.config.MsgID); err != nil {
		w.errorCode = codes.ErrorCodeWithdrawCommit
		return
	}

	return
}

func (w *Withdraw) findProfile() (err error) {
	w.profileModel.UserID = w.config.UserID
	if err = w.profileModel.Find(); err != nil {
		return
	}
	return
}

// Topic topic
func (w *Withdraw) Topic() string {
	return queues.TopicSendGiftSuccess.String()
}

// Message send after Commit success
func (w *Withdraw) Message() []byte {
	return nil
}
