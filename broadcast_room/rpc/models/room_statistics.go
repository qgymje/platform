package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type RoomStatistics struct {
	RoomID         bson.ObjectId `bson:"roomID"`
	BarrageNumber  uint64        `bson:"barrageNumber"`
	AudienceNumber uint64        `bson:"audienceNumber"`
	CreatedAt      time.Time     `bson:"createdAt"`
}
