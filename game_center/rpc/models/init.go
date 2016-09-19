package models

import mgo "gopkg.in/mgo.v2"

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

// InitMongodb init mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

// GetMongo generate  a session copy
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}
