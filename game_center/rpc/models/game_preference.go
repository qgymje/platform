package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//go:generate gen_columns -tag=bson -path=game_preference.go
// 每个玩家保存的游戏配置信息
type GamePreference struct {
	UserID     bson.ObjectId          `bson:"userID"`
	GameID     bson.ObjectId          `bson:"gameID"`
	Preference map[string]interface{} `bson:"preference"`
	CreatedAt  time.Time              `bson:"created_at"`
	UpdatedAt  time.Time              `bson:"updated_at"`
}
