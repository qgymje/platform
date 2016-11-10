package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

// ErrMsgApplyAlreadyExists msg already exists
var ErrMsgApplyAlreadyExists = errors.New("msg already exists")

// SendGift send gift
type SendGift struct {
	ID          int64  `orm:"column(id)"`
	UserID      string `orm:"column(user_id)"`
	ToUserID    string `orm:"column(to_user_id)"`
	BroadcastID string `orm:"column(broadcast_id)"`
	Gift        *Gift  `orm:"rel(fk)"`
	CreatedAt   time.Time
}

// TableName tablename
func (SendGift) TableName() string {
	return TableNameSendGift
}

// Create create a record
func (sg *SendGift) Create(msgID int64) (err error) {
	if err = GetDB().Begin(); err != nil {
		return
	}

	msgModel := &MessageApply{
		MsgID: msgID,
	}

	err = msgModel.Find()
	if err != orm.ErrNoRows {
		return ErrMsgApplyAlreadyExists
	}

	sg.CreatedAt = time.Now()
	_, err = GetDB().Insert(sg)
	if err != nil {
		GetDB().Rollback()
		return
	}

	msgModel.SendGift = sg
	if err = msgModel.Create(); err != nil {
		GetDB().Rollback()
		return
	}

	if err = GetDB().Commit(); err != nil {
		return
	}

	return
}
