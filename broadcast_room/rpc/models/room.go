package models

import (
	"platform/utils"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Room room model object
//go:generate gen_columns -tag=bson -path=./room.go
type Room struct {
	RoomID      bson.ObjectId  `bson:"_id"`
	UserID      bson.ObjectId  `bson:"user_id"`
	UserName    string         `bson:"user_name"`
	Name        string         `bson:"name"`
	Cover       string         `bson:"cover"`
	IsPlaying   bool           `bson:"is_playing"`
	FollowNum   int64          `bson:"follow_num"`
	BroadcastID *bson.ObjectId `bson:"broadcast_id"`
	CreatedAt   time.Time      `bson:"created_at"`
	UpdatedAt   time.Time      `bson:"updated_at"`
	DeletedAt   time.Time      `bson:"deleted_at"`
}

// GetID get room id
func (r *Room) GetID() string {
	return r.RoomID.Hex()
}

// GetUserID get user id
func (r *Room) GetUserID() string {
	return r.UserID.Hex()
}

// GetBroadcastID get room id
func (r *Room) GetBroadcastID() string {
	if r.BroadcastID != nil {
		return r.BroadcastID.Hex()
	}
	return ""
}

// Create create a room
func (r *Room) Create() error {
	session := GetMongo()
	defer session.Close()

	r.RoomID = bson.NewObjectId()
	r.FollowNum = int64(utils.RandomInt(1, 100))
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	return session.DB(DBName).C(ColNameRoom).Insert(&r)
}

func (r *Room) update(m bson.M) error {
	session := GetMongo()
	defer session.Close()

	m[RoomColumns.UpdatedAt] = time.Now()
	change := bson.M{"$set": m}
	return session.DB(DBName).C(ColNameRoom).Update(bson.M{RoomColumns.RoomID: r.RoomID}, change)
}

// AddFollowNum update follow num
func (r *Room) AddFollowNum(n int) error {
	session := GetMongo()
	defer session.Close()

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{RoomColumns.FollowNum: n}},
		ReturnNew: true,
	}
	_, err := session.DB(DBName).C(ColNameRoom).Find(bson.M{RoomColumns.RoomID: r.RoomID}).Apply(change, &r)
	return err
}

// StartPlaying update IsPlaying to true
func (r *Room) StartPlaying(broadcast *Broadcast) error {
	change := bson.M{RoomColumns.BroadcastID: broadcast.BroadcastID, RoomColumns.IsPlaying: true}
	return r.update(change)
}

// EndPlaying update IsPlaying to false
func (r *Room) EndPlaying(broadcast *Broadcast) error {
	change := bson.M{RoomColumns.BroadcastID: nil, RoomColumns.IsPlaying: false}
	return r.update(change)
}

// Update udpate a room info
func (r *Room) Update(name string, cover string) error {
	change := bson.M{RoomColumns.Name: name, RoomColumns.Cover: cover}
	return r.update(change)
}

// FindRoomByUserID find room by user_id
func FindRoomByUserID(userID string) (*Room, error) {
	finder := NewRoomFinder().UserID(userID)
	if err := finder.Do(); err != nil {
		if err == mgo.ErrNotFound {
			return nil, err
		}
	}
	return finder.One(), nil
}

// FindRoomByID find room by room_id
func FindRoomByID(roomID string) (*Room, error) {
	finder := NewRoomFinder().RoomID(roomID)
	if err := finder.Do(); err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return finder.One(), nil
}
