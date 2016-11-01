package models

import "time"

// Store store infomation
type Store struct {
	ID         int64 `orm:"column(id)"`
	MerchantID int64 `orm:"column(merchant_id)"`
	Name       string
	Country    string
	Province   string
	City       string
	Address    string
	Longitude  float64
	Latitude   float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

// TableName table name
func (Store) TableName() string {
	return TableNameStore
}
