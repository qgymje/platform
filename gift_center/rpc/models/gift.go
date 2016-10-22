package models

import "gopkg.in/mgo.v2/bson"

// Gift gift model
type Gift struct {
	GiftID bson.ObjectId `bson:"_id"`
}
