package models

import "time"

// Friend friend relation
type Friend struct {
	ID         int64  `orm:"column(id)"`
	FromUserID string `orm:"column(from_user_id)"`
	ToUserID   string `orm:"column(to_user_id)"`
	CreatedAt  time.Time
}

// Create a friend record
func (f *Friend) Create() (err error) {
	return
}

// IsFriend is friend?
func (f *Friend) IsFriend() bool {
	return false
}
