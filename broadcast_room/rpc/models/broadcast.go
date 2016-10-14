package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Broadcast represent
//go:generate gen_columns -tag=bson -path=./broadcast.go
type Broadcast struct {
	BroadcastID   bson.ObjectId `bson:"_id"`
	RoomID        bson.ObjectId `bson:"room_id"`
	TotalAudience int64         `bson:"total_audience"` // update by user entering the room
	StartTime     time.Time     `bson:"start_time"`
	EndTime       time.Time     `bson:"end_time"`
}

// GetID get room id
func (b *Broadcast) GetID() string {
	return b.BroadcastID.Hex()
}

// Create create a room
func (b *Broadcast) Create() error {
	session := GetMongo()
	defer session.Close()

	b.BroadcastID = bson.NewObjectId()
	b.StartTime = time.Now()

	return session.DB(DBName).C(ColNameBroadcast).Insert(&b)
}
