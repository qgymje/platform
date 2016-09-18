package models

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User 表示一个用户对象
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Nickname string        `bson:"nickname"`
	Password string        `bson:"password" json:"-"`
	Salt     string        `bson:"salt" json:"-"`
	Token    string        `bson:"token" json:"-"`
	HeadImg  string        `bson:"headImg"`
	RegTime  time.Time     `bson:"regTime"`
}

func (u *User) GetID() string {
	return u.ID.Hex()
}

// Create 插入一个用户数据
func (u *User) Create() error {
	session := GetMongo()
	defer session.Close()

	u.ID = bson.NewObjectId()
	u.RegTime = time.Now()

	return session.DB(DBName).C(ColNameUser).Insert(&u)
}

// UpdateToken 更新用户token
func (u *User) UpdateToken(token string) error {
	session := GetMongo()
	defer session.Close()

	change := bson.M{"$set": bson.M{UserColumns.Token: token}}
	err := session.DB(DBName).C(ColNameUser).Update(bson.M{UserColumns.ID: u.ID}, change)
	if err != nil {
		return err
	}
	u.Token = token
	return nil
}

// RemoveToken 移除Token
func (u *User) RemoveToken() error {
	// maybe null?
	return u.UpdateToken("")
}

func findUser(m bson.M) (*User, error) {
	session := GetMongo()
	defer session.Close()

	var user User
	err := session.DB(DBName).C(ColNameUser).Find(m).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil

}

// FindUserByID 根据用户id查用户信息
func FindUserByID(id string) (*User, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("not a valid ObjectIdHex")
	}
	config := bson.M{UserColumns.ID: bson.ObjectIdHex(id)}
	return findUser(config)
}

// FindUserByName 根据用户名查用户信息
func FindUserByName(name string) (*User, error) {
	config := bson.M{UserColumns.Name: name}
	return findUser(config)
}

// FindUserByToken 根据Token查询用户信息
func FindUserByToken(token string) (*User, error) {
	config := bson.M{UserColumns.Token: token}
	return findUser(config)
}

// IsNameUsed 判断Name是否已经被使用
func IsNameUsed(name string) bool {
	_, err := FindUserByName(name)
	return err != ErrNotFound
}
