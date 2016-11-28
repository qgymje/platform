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

// CreateMembers insert members
func CreateMembers(members []*Member) error {
	l := len(members)
	_, err := GetDB().InsertMulti(l, members)
	return err
}

// TableName table name
func (Member) TableName() string {
	return TableNameMember
}

// GetUserID get user id
func (m *Member) GetUserID() string {
	return m.UserID
}

// Update update
func (m *Member) update() (err error) {
	return
}

// Delete leave or be removed the chat
func (m *Member) Delete() (err error) {
	return
}
