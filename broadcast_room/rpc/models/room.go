package models

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// ErrObjectID error object id
	ErrObjectID = errors.New("not a valid objectID")
)

// Room room model object
//go:generate gen_columns -tag=bson -path=./room.go
type Room struct {
	RoomID    bson.ObjectId `bson:"_id"`
	UserID    bson.ObjectId `bson:"user_id"`
	UserName  string        `bson:"user_name"`
	Name      string        `bson:"name"`
	Cover     string        `bson:"cover"`
	IsPlaying bool          `bson:"is_playing"`
	FollowNum int64         `bson:"follow_num"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
	DeletedAt time.Time     `bson:"deleted_at"`
}

// GetID get room id
func (r *Room) GetID() string {
	return r.RoomID.Hex()
}

// Create create a room
func (r *Room) Create() error {
	session := GetMongo()
	defer session.Close()

	r.RoomID = bson.NewObjectId()
	r.CreatedAt = time.Now()

	return session.DB(DBName).C(ColNameRoom).Insert(&r)
}

func (r *Room) update(m bson.M) error {
	session := GetMongo()
	defer session.Close()

	m[RoomColumns.UpdatedAt] = time.Now()
	change := bson.M{"$set": m}
	return session.DB(DBName).C(ColNameRoom).Update(bson.M{RoomColumns.RoomID: r.RoomID}, change)
}

func (r *Room) playing(flag bool) error {
	change := bson.M{RoomColumns.IsPlaying: flag}
	return r.update(change)
}

// StartPlaying update IsPlaying to true
func (r *Room) StartPlaying() error {
	return r.playing(true)
}

// EndPlaying update IsPlaying to false
func (r *Room) EndPlaying() error {
	return r.playing(false)
}

// Update udpate a room info
func (r *Room) Update(name string, cover string) error {
	change := bson.M{RoomColumns.Name: name, RoomColumns.Cover: cover}
	return r.update(change)
}

func findRoom(m bson.M) (*Room, error) {
	session := GetMongo()
	defer session.Close()

	var room Room
	err := session.DB(DBName).C(ColNameRoom).Find(m).One(&room)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &room, nil
}

func findRooms(m bson.M) ([]*Room, error) {
	session := GetMongo()
	defer session.Close()

	var rooms []*Room
	err := session.DB(DBName).C(ColNameRoom).Find(m).All(&rooms)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return rooms, nil
}

// FindRoomByUserID find room by user_id
func FindRoomByUserID(userID string) (*Room, error) {
	if !bson.IsObjectIdHex(userID) {
		return nil, ErrObjectID

	}
	m := bson.M{RoomColumns.UserID: bson.ObjectIdHex(userID)}
	return findRoom(m)

}
