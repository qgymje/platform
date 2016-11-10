package models

import "time"

// MessageApply message apply
type MessageApply struct {
	ID        int64     `orm:"column(id)"`
	MsgID     int64     `orm:"column(msg_id)"`
	SendGift  *SendGift `orm:"rel(fk)"`
	CreatedAt time.Time
}

// TableName table name
func (MessageApply) TableName() string {
	return TableNameMessageApply
}

// Create create a message apply
func (m *MessageApply) Create() (err error) {
	m.CreatedAt = time.Now()
	_, err = GetDB().Insert(m)
	return
}

// Find find the message
func (m *MessageApply) Find() (err error) {
	return GetDB().QueryTable(TableNameMessageApply).Filter("msg_id", m.MsgID).One(m)
}
