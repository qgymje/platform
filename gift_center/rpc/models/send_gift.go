package models

import "time"

// SendGift send gift
type SendGift struct {
	UserID      string
	ToUserID    string
	BroadcastID string
	GiftID      int
	Number      int
	CreatedAt   time.Time
}
