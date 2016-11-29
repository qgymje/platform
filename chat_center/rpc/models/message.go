package models

import (
	"strconv"
	"time"
)

// Message messages
type Message struct {
	ID        int64  `orm:"column(id)"`
	Chat      *Chat  `orm:"rel(fk)"`
	UserID    string `orm:"column(user_id)"` // sender
	Content   string
	CreatedAt time.Time
	DeletedAt time.Time `orm:"null"` // revocation
}

// TableName table name
func (Message) TableName() string {
	return TableNameMessage
}

// GetID get message id
func (m *Message) GetID() string {
	return strconv.FormatInt(m.ID, 10)
}

// Create create a message
func (m *Message) Create() (err error) {
	m.CreatedAt = time.Now()
	_, err = GetDB().Insert(m)
	return err
}
