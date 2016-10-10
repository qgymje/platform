package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Account account model object
type Account struct {
	AccountID   bson.ObjectId `bson:"_id"` // userid
	MerchantID  bson.ObjectId `bson:"merchant_id"`
	Permissions []*Permission `bson:"permissions"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
	DeletedAt   time.Time     `bson:"deleted_at"`
}

// Permission permission controller
type Permission struct {
	Module   string
	Readable bool
	Writable bool
}
