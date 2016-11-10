package gifts

import (
	"platform/commons/codes"
	"platform/gift_center/rpc/models"
	"platform/utils"
	"strconv"
)

// SenderConfig send config
type SenderConfig struct {
	UserID      string
	GiftID      string
	ToUserID    string
	MsgID       string
	BroadcastID string
}

// Sender send gift
type Sender struct {
	UserID        string
	GiftID        string
	sendGiftModel *models.SendGift

	config    *SenderConfig
	errorCode codes.ErrorCode
}

// NewSender createa a new sender
func NewSender(c *SenderConfig) *Sender {
	s := new(Sender)
	s.config = c
	s.sendGiftModel = &models.SendGift{}
	s.sendGiftModel.Gift = &models.Gift{}
	return s
}

// ErrorCode error code
func (s *Sender) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do do the dirty world
func (s *Sender) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("gifts.Sender.Do error: %+v", err)
		}
	}()

	if err = s.save(); err != nil {
		switch err {
		case models.ErrMsgApplyAlreadyExists:
			s.errorCode = codes.ErrorCodeSendGiftMsgApply
			return
		}
	}

	return
}

func (s *Sender) getGiftID() int64 {
	id, _ := strconv.ParseInt(s.config.GiftID, 10, 0)
	return id
}

func (s *Sender) getMsgID() int64 {
	id, _ := strconv.ParseInt(s.config.MsgID, 10, 0)
	return id
}

func (s *Sender) save() (err error) {
	s.sendGiftModel.UserID = s.config.UserID
	s.sendGiftModel.ToUserID = s.config.ToUserID
	s.sendGiftModel.BroadcastID = s.config.BroadcastID
	s.sendGiftModel.Gift.ID = s.getGiftID()

	return s.sendGiftModel.Create(s.getMsgID())
}
