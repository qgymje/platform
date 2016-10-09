package models

import (
	"platform/utils"

	mgo "gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session

// ErrNotFound not found error
var ErrNotFound = mgo.ErrNotFound

// DBName db name
const DBName = "game_center"

// ColNameGame collection name
const ColNameGame = "games"

// ColNameGamePreference collection name
const ColNameGamePreference = "game_preferences"

// ColNameCompany collection name
const ColNameCompany = "companies"

// ColNamePlayerGame collection name
const ColNamePlayerGame = "player_games"

func ensureIndex() {
	c := mongoSession.DB(DBName).C(ColNameGame)
	index := mgo.Index{
		Key: []string{"$text:" + GameColumns.Name, "$text:" + GameColumns.Description},
	}
	err := c.EnsureIndex(index)
	if err != nil {
		utils.GetLog().Error("ensureIndex error: %v", err)
	}
}

// InitMongodb init mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
	ensureIndex()
}

// GetMongo generate  a session copy
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}
