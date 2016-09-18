package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type GamePreference struct {
	UserID     bson.ObjectId
	GameID     bson.ObjectId
	Preference map[string]interface{}
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
