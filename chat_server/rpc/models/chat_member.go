package models

import "time"

// ChatMember in the chat
type ChatMember struct {
	ID        int64  `orm:"column(id)"`
	ChatID    int64  `orm:"column(chat_id)"`
	UserID    string `orm:"column(user_id)"`
	CreatedAt time.Time
	DeletedAt time.Time `orm:"null"`
}
