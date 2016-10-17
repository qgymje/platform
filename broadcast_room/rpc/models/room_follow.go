package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// RoomFollow room follow
//go:generate gen_columns -tag=bson -path=./room_follow.go
type RoomFollow struct {
	RoomID    bson.ObjectId `bson:"room_id"`
	UserID    bson.ObjectId `bson:"user_id"`
	CreatedAt time.Time     `bson:"created_at"`
}

// Create create a room
func (r *RoomFollow) Create() error {
	session := GetMongo()
	defer session.Close()

	r.CreatedAt = time.Now()

	return session.DB(DBName).C(ColNameRoom).Insert(&r)
}
