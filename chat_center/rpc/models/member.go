package models

import "time"

// Member in the chat
type Member struct {
	ID        int64  `orm:"column(id)"`
	Chat      *Chat  `orm:"column(chat_id);rel(fk)"`
	UserID    string `orm:"column(user_id)"`
	CreatedAt time.Time
	DeletedAt time.Time `orm:"null"`
}

// TableName table name
func (Member) TableName() string {
	return TableNameMember
}

// Create join the chat
func (m *Member) Create() (err error) {
	return
}

// Update update
func (m *Member) update() (err error) {
	return
}

// Delete leave or be removed the chat
func (m *Member) Delete() (err error) {
	return
}
