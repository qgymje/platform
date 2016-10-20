// Package models  data access  layer
package models

import (
	"platform/utils"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoSession *mgo.Session

// ErrNotFound not found error
var ErrNotFound = mgo.ErrNotFound

// DBName db name
const DBName = "broadcast_room"

// ColNameRoom collection name
const ColNameRoom = "rooms"

// ColNameBroadcast collection name
const ColNameBroadcast = "broadcasts"

// ColNameAudience audience collection name
const ColNameAudience = "audiences"

// ColNameRoomFollow room_follows
const ColNameRoomFollow = "room_follows"

// ASC sort asc
const ASC = ""

// DESC sort desc
const DESC = "-"

func ensureIndex() {
	c := mongoSession.DB(DBName).C(ColNameRoom)
	index := mgo.Index{
		Key: []string{"$text:" + RoomColumns.Name, "$text:" + RoomColumns.UserName},
	}
	err := c.EnsureIndex(index)
	if err != nil {
		utils.GetLog().Error("endureIndex error: %v", err)
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
