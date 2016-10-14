package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Vote vote model object
//go:generate gen_columns -tag=bson -path=./vote.go
type Vote struct {
	VoteID      bson.ObjectId `bson:"vote_id"`
	UserID      bson.ObjectId `bson:"user_id"`
	Number      int           `bson:"number"`
	OptionCount []struct {
		Name  string `bson:"name"`
		Count int64  `bson:"count"`
	} `bson:"option_count"`
	Duration  int64     `bson:"duration"`
	CreatedAt time.Time `bson:"created_at"`
}
