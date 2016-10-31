package models

import "time"

// Account account model object
type Account struct {
	ID         int64  `orm:"column(id)"` // userid
	MerchantID string `rom:"column(merchant_id)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

// TableName table name
func (Account) TableName() string {
	return TableNameAccount
}

// Permission permission controller
type Permission struct {
	ID       int64
	Module   string
	Readable bool
	Writable bool
}

// TableName table name
func (Permission) TableName() string {
	return TableNamePermission
}
