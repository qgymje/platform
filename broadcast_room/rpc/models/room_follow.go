package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// RoomFollow room follow
//go:generate gen_columns -tag=bson -path=./room_follow.go
type RoomFollow struct {
	RoomID    bson.ObjectId `bson:"room_id"`
	UserID    bson.ObjectId `bson:"user_id"`
	CreatedAt time.Time     `bson:"created_at"`
	DeletedAt time.Time     `bson:"deleted_at"`
}

// Follow create a record
func (r *RoomFollow) Follow() error {
	session := GetMongo()
	defer session.Close()

	m := bson.M{}
	m[RoomFollowColumns.CreatedAt] = time.Now()
	m[RoomFollowColumns.DeletedAt] = time.Time{}
	change := bson.M{"$set": m}

	_, err := session.DB(DBName).C(ColNameRoomFollow).Upsert(bson.M{RoomFollowColumns.UserID: r.UserID, RoomFollowColumns.RoomID: r.RoomID}, change)

	room := &Room{RoomID: r.RoomID}
	room.AddFollowNum(1)
	return err
}

// Unfollow unfollow
func (r *RoomFollow) Unfollow() error {
	session := GetMongo()
	defer session.Close()

	m := bson.M{}
	m[RoomFollowColumns.DeletedAt] = time.Now()
	change := bson.M{"$set": m}
	_, err := session.DB(DBName).C(ColNameRoomFollow).Upsert(bson.M{RoomFollowColumns.UserID: r.UserID, RoomFollowColumns.RoomID: r.RoomID}, change)

	room := &Room{RoomID: r.RoomID}
	room.AddFollowNum(-1)
	return err
}
