package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	GameID      bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Cover       string        `bson:"cover"`
	Description string        `bson:"description"`
	Status      int           `bson:"status"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
	PublishedAt time.Time     `bson:"published_at"`
}
