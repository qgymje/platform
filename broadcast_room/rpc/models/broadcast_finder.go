package models

import (
	"math"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// BroadcastFinder is a broadcast finder
type BroadcastFinder struct {
	skip, limit int
	where       bson.M
	ids         []bson.ObjectId
	broadcasts  []*Broadcast
	err         error
}

// NewBroadcastFinder create a BroadcastFinder
func NewBroadcastFinder() *BroadcastFinder {
	f := new(BroadcastFinder)
	f.where = bson.M{}
	f.ids = []bson.ObjectId{}
	f.broadcasts = []*Broadcast{}
	return f
}

// BroadcastID find by broadcast id
func (b *BroadcastFinder) BroadcastID(broadcastID string) *BroadcastFinder {
	var broadcastObjID bson.ObjectId
	broadcastObjID, b.err = StringToObjectID(broadcastID)
	if b.err == nil {
		b.where[BroadcastColumns.BroadcastID] = broadcastObjID
	}
	return b
}

// Limit limit
func (b *BroadcastFinder) Limit(offset, limit int) *BroadcastFinder {
	b.skip = int(math.Max(0, float64(offset-1))) * limit
	b.limit = limit
	return b
}

// ByIDs by user ids
func (b *BroadcastFinder) ByIDs(ids []string) *BroadcastFinder {
	b.ids, b.err = StringsToObjectIDs(ids)
	if b.err == nil {
		b.skip = 0
		b.limit = len(b.ids)
		b.where[BroadcastColumns.BroadcastID] = bson.M{"$in": b.ids}
	}
	return b
}

// IsPlaying fetch by the playing broadcast
func (b *BroadcastFinder) IsPlaying() *BroadcastFinder {
	b.where[BroadcastColumns.EndTime] = time.Time{}
	return b
}

// Do do the search job
func (b *BroadcastFinder) Do() (err error) {
	session := GetMongo()
	defer session.Close()

	return session.DB(DBName).C(ColNameBroadcast).Find(b.where).Skip(b.skip).Limit(b.limit).All(&b.broadcasts)
}

// One get only one result
func (b *BroadcastFinder) One() *Broadcast {
	if len(b.broadcasts) > 0 {
		return b.broadcasts[0]
	}
	return nil
}

// Result return the games that found
func (b *BroadcastFinder) Result() []*Broadcast {
	return b.broadcasts
}

// FetchByID fetch by broadcast id
func (b *BroadcastFinder) FetchByID(id string) *Broadcast {
	for i := range b.broadcasts {
		if b.broadcasts[i].GetID() == id {
			return b.broadcasts[i]
		}
	}
	return nil
}

// FetchByRoomID fetch broadcast by room id
func (b *BroadcastFinder) FetchByRoomID(id string) *Broadcast {
	for i := range b.broadcasts {
		if b.broadcasts[i].GetRoomID() == id {
			return b.broadcasts[i]
		}
	}
	return nil
}

// Count return the total num of the query
func (b *BroadcastFinder) Count() int64 {
	session := GetMongo()
	defer session.Close()

	n, err := session.DB(DBName).C(ColNameBroadcast).Find(b.where).Count()
	if err != nil {
		return 0
	}
	return int64(n)
}
