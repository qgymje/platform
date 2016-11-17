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
