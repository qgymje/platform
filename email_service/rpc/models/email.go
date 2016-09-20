package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// EmailProvider sms service provider
type EmailProvider int

const (
	// SendCloud used in China mainland
	SendCloud EmailProvider = iota + 1
)

// EmailType sms sender type
type EmailType int

const (
	// RegisterCode register code
	RegisterCode EmailType = iota + 1
)

// Email represents sms sender data struct
// it's will be a standalone services when etcd is ready
type Email struct {
	Address   string    `bson:"address"`
	Content   string    `bson:"content"`
	Type      int       `bson:"type"`
	Provider  int       `bson:"provider"`
	UsedAt    time.Time `bson:"used_at"`
	CreatedAt time.Time `bson:"created_at"`
}

// Create create a sms record
func (e *Email) Create() error {
	session := GetMongo()
	defer session.Close()

	return session.DB(DBName).C(ColNameEmail).Insert(&e)
}

func (e *Email) update(condition, m bson.M) error {
	session := GetMongo()
	defer session.Close()

	change := bson.M{"$set": m}
	return session.DB(DBName).C(ColNameEmail).Update(condition, change)
}

func (e *Email) condition() bson.M {
	return bson.M{EmailColumns.Address: e.Address}
}

// Use use register code
func (e *Email) Use() error {
	change := bson.M{EmailColumns.UsedAt: time.Now()}
	condition := e.condition()
	condition[EmailColumns.Type] = RegisterCode
	return e.update(condition, change)
}

func findEmail(m bson.M) (*Email, error) {
	session := GetMongo()
	defer session.Close()

	var sms Email
	err := session.DB(DBName).C(ColNameEmail).Find(m).One(&sms)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &sms, nil
}

// FindEmailByAddress find sms code by address
func FindEmailByAddress(address string) (*Email, error) {
	m := bson.M{EmailColumns.Address: address}
	return findEmail(m)
}
