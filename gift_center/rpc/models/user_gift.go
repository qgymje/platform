package models

import "time"

// UserGift user gift
type UserGift struct {
	ID        int64  `orm:"column(id)"`
	UserID    string `orm:"column(user_id)"`
	Gift      *Gift
	Number    int
	Price     float64
	CreatedAt time.Time
}
