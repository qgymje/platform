package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// UserGift user gift
type UserGift struct {
	UserID    bson.ObjectId `bson:"user_id"`
	GiftID    bson.ObjectId `bson:"gift_id"`
	Number    int           `bson:"number"`
	Price     float64       `bson:"price"`
	CreatedAt time.Time     `bson:"created_at"`
}
