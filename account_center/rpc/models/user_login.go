package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type UserLogin struct {
	UserID    bson.ObjectId `bson:"user_id"`
	CreatedAt time.Time     `bson:"created_at"`
}

func (ul *UserLogin) Create() error {
	session := GetMongo()
	defer session.Close()

	ul.CreatedAt = time.Now()

	return session.DB(DBName).C(ColNameUserLogin).Insert(&ul)
}
