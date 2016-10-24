package models

import (
	"math"
	"platform/utils"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// BarrageFinder barrage finder
type BarrageFinder struct {
	startTime   time.Time
	endTime     time.Time
	skip, limit int
	where       bson.M
	broadcastID bson.ObjectId
	err         error
	barrages    []*Barrage
}

// NewBarrageFinder create a BarrageFinder
func NewBarrageFinder() *BarrageFinder {
	f := new(BarrageFinder)
	f.where = bson.M{}
	f.barrages = []*Barrage{}
	return f
}

// Err err
func (b *BarrageFinder) Err() error {
	return b.err
}

// BroadcastID find by broadcast id
func (b *BarrageFinder) BroadcastID(broadcastID string) *BarrageFinder {
	b.broadcastID, b.err = StringToObjectID(broadcastID)
	if b.err == nil {
		b.where[BarrageColumns.BroadcastID] = b.broadcastID
	}
	return b
}

// Duration duration
func (b *BarrageFinder) Duration(st, et int64) *BarrageFinder {
	startTime := time.Unix(st, 0)
	endTime := time.Unix(et, 0)
	b.where[BarrageColumns.CreatedAt] = bson.M{"$gt": startTime, "$lt": endTime}
	return b
}

// Limit limit
func (b *BarrageFinder) Limit(offset, limit int) *BarrageFinder {
	b.skip = int(math.Max(0, float64(offset-1))) * limit
	b.limit = limit
	return b
}

// Do do the search job
func (b *BarrageFinder) Do() (err error) {
	session := GetMongo()
	defer session.Close()
	utils.Dump(b.where)
	return session.DB(DBName).C(ColNameBarrage).Find(b.where).Limit(b.limit).All(&b.barrages)
}

// Result the result
func (b *BarrageFinder) Result() []*Barrage {
	return b.barrages
}

// Count return the total num of the query
func (b *BarrageFinder) Count() int64 {
	session := GetMongo()
	defer session.Close()

	n, err := session.DB(DBName).C(ColNameBarrage).Find(b.where).Count()
	if err != nil {
		return 0
	}
	return int64(n)
}
