package models

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoSession *mgo.Session
var ErrNotFound = mgo.ErrNotFound

const DBName = "liveBroadcast"
const ColNameBarrage = "barrages"

// InitMongodb 根据配置初始化mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

// GetMongo 封装一个session, 在使用后需要Close()
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}

var (
	ErrObjectID = errors.New("not a valid objectID")
)

func StringToObjectID(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(string(id)) {
		return bson.ObjectId(""), ErrObjectID
	}
	return bson.ObjectIdHex(id), nil
}
