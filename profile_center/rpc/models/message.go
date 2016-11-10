package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Message message model
type Message struct {
	ID           int64  `orm:"column(id)"`
	UserID       string `orm:"column(user_id)"`
	TypeID       uint   `orm:"column(type_id)"`
	TargetID     string `orm:"column(target_id)"`
	SnowBall     uint
	SnowFlake    uint
	Status       int
	CreatedAt    time.Time
	RollbackedAt time.Time `orm:"null"`
	CommittedAt  time.Time `orm:"null"`
}

// TableName table name
func (Message) TableName() string {
	return TableNameMessage
}

// Create a message
func (m *Message) Create() (err error) {
	_, err = GetDB().Insert(m)
	return
}

// Find find the full message info
func (m *Message) Find() (err error) {
	return GetDB().QueryTable(TableNameMessage).Filter("id", m.ID).One(m)
}

// Rollback rollback
func (m *Message) Rollback() (err error) {
	m.RollbackedAt = time.Now()
	m.Status = int(Rollbacked)
	params := orm.Params{"rollbacked_at": m.RollbackedAt, "status": m.Status}
	_, err = GetDB().QueryTable(TableNameMessage).Filter("id", m.ID).Update(params)
	return
}

// Commit commit
func (m *Message) Commit() (err error) {
	m.CommittedAt = time.Now()
	m.Status = int(Committed)
	params := orm.Params{"committed_at": m.CommittedAt, "status": m.Status}
	_, err = GetDB().QueryTable(TableNameMessage).Filter("id", m.ID).Update(params)
	return
}
