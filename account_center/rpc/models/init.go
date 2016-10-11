// Package models  data access layer
package models

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ErrObjectID invalid object id
var ErrObjectID = errors.New("not a valid objectID")

var mongoSession *mgo.Session

// ErrNotFound not found error
var ErrNotFound = mgo.ErrNotFound

// DBName db name
const DBName = "account_center"

// ColNameUser collection name
const ColNameUser = "users"

// ColNameUserLogin collection name
const ColNameUserLogin = "user_logins"

// InitMongodb init mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

// GetMongo generate  a session copy
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}

// StringToObjectID string to mongo objectid
func StringToObjectID(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(string(id)) {
		return bson.ObjectId(""), ErrObjectID
	}
	return bson.ObjectIdHex(id), nil
}

// StringsToObjectIDs strings to mongo objectids
func StringsToObjectIDs(ids []string) ([]bson.ObjectId, error) {
	IDHexs := []bson.ObjectId{}
	for _, id := range ids {
		if !bson.IsObjectIdHex(string(id)) {
			return nil, ErrObjectID
		}
		IDHexs = append(IDHexs, bson.ObjectIdHex(string(id)))
	}
	return IDHexs, nil
}
