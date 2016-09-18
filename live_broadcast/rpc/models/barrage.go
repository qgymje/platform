package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Barrage struct {
	RoomID    bson.ObjectId `bson:"roomID"`
	UserID    bson.ObjectId `bson:"userID"`
	UserName  string        `bson:"userName"`
	Message   string        `bson:"message"`
	CreatedAt time.Time     `bson:"createdAt"`
}

func Save(barrages []*Barrage) (err error) {
	session := GetMongo()
	defer session.Clone()

	return session.DB(DBName).C(ColNameBarrage).Insert(&barrages)
}

func findBarrages(m bson.M) ([]*Barrage, error) {
	session := GetMongo()
	defer session.Close()

	var barrages []*Barrage

	err := session.DB(DBName).C(ColNameBarrage).Find(m).All(&barrages)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return barrages, nil
}

func FindBarragesByRoom(roomID string) ([]*Barrage, error) {
	m := bson.M{BarrageColumns.RoomID: bson.ObjectIdHex(roomID)}
	return findBarrages(m)
}
