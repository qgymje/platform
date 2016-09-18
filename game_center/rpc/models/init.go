package models

import mgo "gopkg.in/mgo.v2"

var mongoSession *mgo.Session
var ErrNotFound = mgo.ErrNotFound

const DBName = "game_center"
const ColNameRoom = "games"
const ColNameRoomAuth = "game_preferences"

// InitMongodb 根据配置初始化mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

// GetMongo 封装一个session, 在使用后需要Close()
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}
