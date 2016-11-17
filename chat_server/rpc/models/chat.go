package models

import "time"

// Chat chat model
type Chat struct {
	ID        int64 `orm:"column(id)"`
	Name      string
	UserID    string        `orm:"user_id"` // creater
	Members   []*ChatMember `orm:"reverse(many)"`
	CreatedAt time.Time
}
