package models

import (
	"platform/utils"
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

// NewAudience create new Audience model
func NewAudience(broadcastID, userID string) (*Audience, error) {
	broadcastObjID, err := StringToObjectID(broadcastID)
	if err != nil {
		return nil, err
	}
	userObjID, err := StringToObjectID(userID)
	if err != nil {
		return nil, err
	}

	return &Audience{
		BroadcastID: broadcastObjID,
		UserID:      userObjID,
	}, nil
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
	}
	a.LeaveTime = time.Time{}
	return true
}

// Enter enter a broadcast
func (a *Audience) Enter() error {
	session := GetMongo()
	defer session.Close()

	a.EnterTime = time.Now()

	return session.DB(DBName).C(ColNameAudience).Insert(&a)
}

// Leave leave a broadcast, update the least record
func (a *Audience) Leave() error {
	session := GetMongo()
	defer session.Close()

	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{AudienceColumns.LeaveTime: time.Now()}},
		ReturnNew: true,
	}
	where := bson.M{
		AudienceColumns.BroadcastID: a.BroadcastID,
		AudienceColumns.UserID:      a.UserID,
	}
	info, err := session.DB(DBName).C(ColNameAudience).Find(where).Sort("$natural").Apply(change, &a)
	utils.Dump(info)
	return err
}
