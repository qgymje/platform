package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Audience provide a audience enter/leave time log
//go:generate gen_columns -tag=bson -path=./audience.go
type Audience struct {
	BroadcastID bson.ObjectId `bson:"broadcast_id"`
	UserID      bson.ObjectId `bson:"user_id"`
	EnterTime   time.Time     `bson:"enter_time"`
	LeaveTime   time.Time     `bson:"leave_time"` // maybe null
}

// HasEntered did audience entere the broadcast room before?
func (a *Audience) HasEntered() bool {
	session := GetMongo()
	defer session.Close()

	c := bson.M{AudienceColumns.BroadcastID: a.BroadcastID, AudienceColumns.UserID: a.UserID}
	err := session.DB(DBName).C(ColNameAudience).Find(c).One(&a)
	if err != nil {
		if err == mgo.ErrNotFound {
			return false
		}
		return false
	}
	return false
}

// Enter enter a broadcast
func (a *Audience) Enter() error {
	session := GetMongo()
	defer session.Close()

	a.EnterTime = time.Now()

	return session.DB(DBName).C(ColNameAudience).Insert(&a)
}

// Leave leave a broadcast
func (a *Audience) Leave() error {
	session := GetMongo()
	defer session.Close()

	m := bson.M{}
	m[AudienceColumns.LeaveTime] = time.Now()
	change := bson.M{"$set": m}

	return session.DB(DBName).C(ColNameAudience).Update(bson.M{AudienceColumns.BroadcastID: a.BroadcastID, AudienceColumns.UserID: a.UserID}, change)
}
