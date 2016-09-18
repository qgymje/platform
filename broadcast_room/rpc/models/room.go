package models

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrObjectID = errors.New("not a valid objectID")
)

type BroadcastRoom struct {
	ID          bson.ObjectId `bson:"_id"`
	UserID      bson.ObjectId `json:"userID" bson:"userID"`           //用户ID
	Name        string        `json:"name" bson:"name"`               //标题
	Cover       string        `json:"cover" bson:"cover"`             //封面图片地址
	Channel     string        `json:"channel" bson:"channel"`         //领域
	SubChannel  string        `json:"subChannel" bson:"subChannel"`   // 子领域
	IsPlaying   bool          `json:"isPlaying" bson:"isPlaying"`     //是否正在直播
	Orientation int8          `json:"orientation" bson:"orientation"` //横竖屏 0 未设置 1 横屏 2 竖屏

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"` //创建时间
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"` //更新时间
}

func (r *BroadcastRoom) GetID() string {
	return r.ID.Hex()
}

func (r *BroadcastRoom) Create() error {
	session := GetMongo()
	defer session.Close()

	r.ID = bson.NewObjectId()
	r.CreatedAt = time.Now()

	return session.DB(DBName).C(ColNameRoom).Insert(&r)
}

func (r *BroadcastRoom) update(m bson.M) error {
	session := GetMongo()
	defer session.Close()

	m[BroadcastRoomColumns.UpdatedAt] = time.Now()
	change := bson.M{"$set": m}
	return session.DB(DBName).C(ColNameRoom).Update(bson.M{BroadcastRoomColumns.ID: r.ID}, change)
}

func (r *BroadcastRoom) playing(flag bool) error {
	change := bson.M{BroadcastRoomColumns.IsPlaying: flag}
	return r.update(change)
}

func (r *BroadcastRoom) Start() error {
	return r.playing(true)
}

func (r *BroadcastRoom) End() error {
	return r.playing(false)
}

func (r *BroadcastRoom) UpdateName(name string) error {
	change := bson.M{BroadcastRoomColumns.Name: name}
	return r.update(change)
}

func (r *BroadcastRoom) UpdateCover(cover string) error {
	change := bson.M{BroadcastRoomColumns.Cover: cover}
	return r.update(change)
}

func (r *BroadcastRoom) UpdateChannel(channel, subChannel string) error {
	change := bson.M{BroadcastRoomColumns.Channel: channel, BroadcastRoomColumns.SubChannel: subChannel}
	return r.update(change)
}

func StringToObjectID(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(string(id)) {
		return bson.ObjectId(""), ErrObjectID
	}
	return bson.ObjectIdHex(id), nil
}

func StringToObjectIDs(ids string) ([]bson.ObjectId, error) {
	IDHexs := []bson.ObjectId{}
	for _, id := range ids {
		if !bson.IsObjectIdHex(string(id)) {
			return nil, ErrObjectID
		}
		IDHexs = append(IDHexs, bson.ObjectIdHex(string(id)))
	}
	return IDHexs, nil
}

func findBroadcastRoom(m bson.M) (*BroadcastRoom, error) {
	session := GetMongo()
	defer session.Close()

	var room BroadcastRoom
	err := session.DB(DBName).C(ColNameRoom).Find(m).One(&room)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &room, nil

}

func findBroadcastRooms(m bson.M) (*BroadcastRoom, error) {
	session := GetMongo()
	defer session.Close()

	var room BroadcastRoom
	err := session.DB(DBName).C(ColNameRoom).Find(m).All(&room)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &room, nil
}

func FindBroadcastRoomByUserID(userID string) (*BroadcastRoom, error) {
	if !bson.IsObjectIdHex(userID) {
		return nil, ErrObjectID
	}
	m := bson.M{BroadcastRoomColumns.UserID: bson.ObjectIdHex(userID)}
	return findBroadcastRoom(m)
}
