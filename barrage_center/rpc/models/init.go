package models

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoSession *mgo.Session

// ErrNotFound not found error
var ErrNotFound = mgo.ErrNotFound

// DBName db name
const DBName = "barrage_center"

// ColNameBarrage collection name
const ColNameBarrage = "barrages"

// InitMongodb init mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

// GetMongo generate  a session copy
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}

var (
	// ErrObjectID error object id
	ErrObjectID = errors.New("not a valid objectID")
)

// StringToObjectID string to bson objectId
func StringToObjectID(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(string(id)) {
		return bson.ObjectId(""), ErrObjectID
	}
	return bson.ObjectIdHex(id), nil
}

// StringsToObjectIDs strings to bson objectIds
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
