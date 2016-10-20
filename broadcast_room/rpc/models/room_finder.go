package models

import (
	"math"

	"gopkg.in/mgo.v2/bson"
)

// RoomFinder is a room finder
type RoomFinder struct {
	skip, limit int
	order       string
	where       bson.M
	rooms       []*Room
	err         error
}

// NewRoomFinder create a new RoomFinder
func NewRoomFinder() *RoomFinder {
	f := new(RoomFinder)
	f.where = bson.M{}
	f.rooms = []*Room{}

	return f
}

// RoomID find by room id
func (r *RoomFinder) RoomID(roomID string) *RoomFinder {
	var roomObjID bson.ObjectId
	roomObjID, r.err = StringToObjectID(roomID)
	if r.err == nil {
		r.where[RoomColumns.RoomID] = roomObjID
	}
	return r
}

// UserID find by user id
func (r *RoomFinder) UserID(userID string) *RoomFinder {
	var userObjID bson.ObjectId
	userObjID, r.err = StringToObjectID(userID)
	if userObjID == "" {
		r.where[RoomColumns.UserID] = userObjID
	}
	return r
}

// Limit limit
func (r *RoomFinder) Limit(offset, limit int) *RoomFinder {
	r.skip = int(math.Max(0, float64(offset-1))) * limit
	r.limit = limit
	return r
}

// Order order
func (r *RoomFinder) Order(o string) *RoomFinder {
	r.order = o
	return r
}

// Search search
func (r *RoomFinder) Search(search string) *RoomFinder {
	if search != "" {
		r.where["$text"] = bson.M{"$search": search}
	}
	return r
}

// Do do the search job
func (r *RoomFinder) Do() (err error) {
	session := GetMongo()
	defer session.Close()

	return session.DB(DBName).C(ColNameRoom).Find(r.where).Skip(r.skip).Limit(r.limit).All(&r.rooms)
}

// One get only one result
func (r *RoomFinder) One() *Room {
	if len(r.rooms) > 0 {
		return r.rooms[0]
	}
	return nil
}

// Result return the games that found
func (r *RoomFinder) Result() []*Room {
	return r.rooms
}

// Count return the total num of the query
func (r *RoomFinder) Count() int64 {
	session := GetMongo()
	defer session.Close()

	n, err := session.DB(DBName).C(ColNameRoom).Find(r.where).Count()
	if err != nil {
		return 0
	}
	return int64(n)
}
