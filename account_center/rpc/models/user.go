package models

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User  a user model object
type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Phone     string        `bson:"phone"`
	Email     string        `bson:"email"`
	Nickname  string        `bson:"nickname"`
	Password  string        `bson:"password" json:"-"`
	Salt      string        `bson:"salt" json:"-"`
	Token     string        `bson:"token" json:"-"`
	Avatar    string        `bson:"avatar"`
	CreatedAt time.Time     `bson:"created_at"`
}

// GetID get hexed user_id
func (u *User) GetID() string {
	return u.ID.Hex()
}

// Create 插入一个用户数据
func (u *User) Create() error {
	session := GetMongo()
	defer session.Close()

	u.ID = bson.NewObjectId()
	u.CreatedAt = time.Now()

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

// FindUserByPhone find by phone
func FindUserByPhone(phone string) (*User, error) {
	config := bson.M{UserColumns.Phone: phone}
	return findUser(config)
}

// FindUserByEmail find by email
func FindUserByEmail(email string) (*User, error) {
	config := bson.M{UserColumns.Email: email}
	return findUser(config)
}

// FindUserByToken 根据Token查询用户信息
func FindUserByToken(token string) (*User, error) {
	config := bson.M{UserColumns.Token: token}
	return findUser(config)
}

// IsPhoneUsed is phone used
func IsPhoneUsed(phone string) bool {
	_, err := FindUserByPhone(phone)
	return err != ErrNotFound
}

// IsEmailUsed is email used
func IsEmailUsed(email string) bool {
	_, err := FindUserByEmail(email)
	return err != ErrNotFound
}
