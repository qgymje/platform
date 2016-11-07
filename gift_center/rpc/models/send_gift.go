package models

import "time"

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
// TODO: it's just a simple implement
func (sg *SendGift) Create() (err error) {
	sg.CreatedAt = time.Now()
	_, err = GetDB().Insert(sg)
	return
}
