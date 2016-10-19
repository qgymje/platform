package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Broadcast represent
//go:generate gen_columns -tag=bson -path=./broadcast.go
type Broadcast struct {
	BroadcastID   bson.ObjectId `bson:"_id"`
	RoomID        bson.ObjectId `bson:"room_id"`
	TotalAudience int64         `bson:"total_audience"` // update by user entering the room
	StartTime     time.Time     `bson:"start_time"`
	EndTime       time.Time     `bson:"end_time"`
}

// GetID get room id
func (b *Broadcast) GetID() string {
	return b.BroadcastID.Hex()
}

// GetRoomID get room id
func (b *Broadcast) GetRoomID() string {
	return b.RoomID.Hex()
}

// Create create a room
func (b *Broadcast) Create() error {
	session := GetMongo()
	defer session.Close()

	b.BroadcastID = bson.NewObjectId()
	b.StartTime = time.Now()

	return session.DB(DBName).C(ColNameBroadcast).Insert(&b)
}

func (b *Broadcast) update(m bson.M) error {
	session := GetMongo()
	defer session.Close()

	change := bson.M{"$set": m}
	return session.DB(DBName).C(ColNameBroadcast).Update(bson.M{BroadcastColumns.BroadcastID: b.BroadcastID}, change)
}

// AddAudience add total audience number
func (b *Broadcast) AddAudience(n int) error {
	session := GetMongo()
	defer session.Close()

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{BroadcastColumns.TotalAudience: n}},
		ReturnNew: true,
	}
	_, err := session.DB(DBName).C(ColNameRoom).Find(bson.M{BroadcastColumns.BroadcastID: b.BroadcastID}).Apply(change, &b)
	return err

}

// End end the broadcast
func (b *Broadcast) End() error {
	m := bson.M{BroadcastColumns.EndTime: time.Now()}
	return b.update(m)
}

// FindBroadcastByID find by broadcast id
func FindBroadcastByID(broadcastID string) (*Broadcast, error) {
	session := GetMongo()
	defer session.Close()

	finder := NewBroadcastFinder().BroadcastID(broadcastID)
	if err := finder.Do(); err != nil {
		return nil, err
	}

	broadcast := finder.One()
	if broadcast != nil {
		return broadcast, nil
	}
	return nil, ErrNotFound
}

// FindBroadcastByRoomID find broadcast by room id
func FindBroadcastByRoomID(roomID string) (*Broadcast, error) {
	session := GetMongo()
	defer session.Close()
	var broadcast Broadcast
	roomObjID, _ := StringToObjectID(roomID)
	err := session.DB(DBName).C(ColNameBroadcast).Find(bson.M{BroadcastColumns.RoomID: roomObjID}).Sort(DESC + BroadcastColumns.StartTime).One(&broadcast)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
	}
	return &broadcast, nil
}
