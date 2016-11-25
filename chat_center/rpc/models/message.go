package models

import "time"

// Message messages
type Message struct {
	ID        int64  `orm:"column(id)"`
	ChatID    int64  `orm:"column(chat_id)"`
	UserID    string `orm:"column(user_id)"` // sender
	Content   string
	CreatedAt time.Time
	DeletedAt time.Time `rom:"null"` // revocation
}

// TableName table name
func (Message) TableName() string {
	return TableNameMessage
}

// Create create a message
func (m *Message) Create() (err error) {
	m.CreatedAt = time.Now()
	_, err = GetDB().Insert(m)
	return err
}
