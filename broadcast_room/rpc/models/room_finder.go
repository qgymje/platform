package models

import (
	"math"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RoomFinder is a room finder
type RoomFinder struct {
	skip, limit int
	search      string
	order       string

	rooms []*Room
}

// NewRoomFinder create a new RoomFinder
func NewRoomFinder() *RoomFinder {
	f := new(RoomFinder)
	f.rooms = []*Room{}

	return f
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
	r.search = search
	return r
}

func (r *RoomFinder) condition() bson.M {
	where := bson.M{}

	if r.search != "" {
		where["$text"] = bson.M{"$search": r.search}
	}
	return where
}

// Do do the search job
func (r *RoomFinder) Do() (err error) {
	session := GetMongo()
	defer session.Close()

	err = session.DB(DBName).C(ColNameRoom).Find(r.condition()).Skip(r.skip).Limit(r.limit).All(&r.rooms)
	if err != nil {
		if err == mgo.ErrNotFound {
			return ErrNotFound
		}
		return err
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

	n, err := session.DB(DBName).C(ColNameRoom).Find(r.condition()).Count()
	if err != nil {
		return 0
	}
	return int64(n)
}
