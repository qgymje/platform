package models

import "time"

// Message message model
type Message struct {
	ID          int64  `orm:"column(id)"`
	UserID      string `orm:"column(user_id)"`
	TypeID      int    `orm:"column(type_id)"`
	TargetID    string `orm:"column(target_id)"`
	Status      int
	CreatedAt   time.Time
	ConfirmedAt time.Time
	NotifiedAt  time.Time
	AckAt       time.Time
}
