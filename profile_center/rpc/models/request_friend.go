package models

import "time"

// RequestFriend request friend
type RequestFriend struct {
	ID         int64  `orm:"column(id)"`
	FromUserID string `orm:"column(from_user_id)"`
	ToUserID   string `orm:"column(to_user_id)"`
	Message    string
	Status     int8 // 0 unprocessed 1 success 2 deny
	CreatedAt  time.Time
}
