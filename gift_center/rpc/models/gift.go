package models

import "time"

// Gift gift model
type Gift struct {
	ID        int64 `orm:"column(id)"`
	Name      string
	Image     string
	Price     float64
	CreatedAt time.Time
}
