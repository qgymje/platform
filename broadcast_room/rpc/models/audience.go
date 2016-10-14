package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Audience provide a audience enter/leave time log
//go:generate gen_columns -tag=bson -path=./audience.go
type Audience struct {
	BroadcastID bson.ObjectId `bson:"broadcast_id"`
	UserID      bson.ObjectId `bson:"user_id"`
	EnterTime   time.Time     `bson:"enter_time"`
	LeaveTime   time.Time     `bson:"leave_time"`
}
