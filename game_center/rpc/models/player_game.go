package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// PlayerGame represents a player with games
//go:generate gen_columns -tag=bson -path=./player_game.go
type PlayerGame struct {
	UserID    bson.ObjectId `bson:"user_id"`
	GameID    bson.ObjectId `bson:"game_id"`
	StartTime time.Time     `bson:"start_time`
	EndTime   time.Time     `bson:"end_time"`
}
