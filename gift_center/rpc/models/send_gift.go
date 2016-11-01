package models

import "time"

// SendGift send gift
type SendGift struct {
	ID          int64  `orm:"column(id)"`
	UserID      string `orm:"column(user_id)"`
	ToUserID    string `orm:"column(to_user_id)"`
	BroadcastID string `orm:"column(broadcast_id)"`
	Gift        *Gift
	Number      int
	CreatedAt   time.Time
}
