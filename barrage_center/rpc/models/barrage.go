package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Barrage message
//go:generate gen_columns -tag=bson -path=./barrage.go
type Barrage struct {
	BroadcastID bson.ObjectId `bson:"broadcast_id"`
	UserID      bson.ObjectId `bson:"user_id"`
	Text        string        `bson:"text"`
	Username    string        `bson:"username"`
	Level       int64         `bson:"level"`
	CreatedAt   time.Time     `bson:"created_at"`
}

// GetBroadcastID broadcast id
func (b *Barrage) GetBroadcastID() string {
	return b.BroadcastID.Hex()
}

// GetUserID get user id
func (b *Barrage) GetUserID() string {
	return b.UserID.Hex()
}

// Create new barrages
func (b *Barrage) Create() (err error) {
	session := GetMongo()
	defer session.Clone()

	return session.DB(DBName).C(ColNameBarrage).Insert(&b)
}
