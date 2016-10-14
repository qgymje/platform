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
