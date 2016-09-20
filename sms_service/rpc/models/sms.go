package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SMSProvider sms service provider
type SMSProvider int

const (
	// SendCloud used in China mainland
	SendCloud SMSProvider = iota + 1
	// Twilio out China mainland
	Twilio
)

// SMSType sms sender type
type SMSType int

const (
	// RegisterCode register code
	RegisterCode SMSType = iota + 1
)

// SMS represents sms sender data struct
// it's will be a standalone services when etcd is ready
type SMS struct {
	Phone     string    `bson:"phone"`
	Content   string    `bson:"content"`
	Type      int       `bson:"type"`
	Provider  int       `bson:"provider"`
	UsedAt    time.Time `bson:"used_at"`
	CreatedAt time.Time `bson:"created_at"`
}

// Create create a sms record
func (s *SMS) Create() error {
	session := GetMongo()
	defer session.Close()

	return session.DB(DBName).C(ColNameSMS).Insert(&s)
}

func (s *SMS) update(condition, m bson.M) error {
	session := GetMongo()
	defer session.Close()

	change := bson.M{"$set": m}
	return session.DB(DBName).C(ColNameSMS).Update(condition, change)
}

func (s *SMS) condition() bson.M {
	return bson.M{SMSColumns.Phone: s.Phone, SMSColumns.Content: s.Content}
}

// Use use register code
func (s *SMS) Use() error {
	change := bson.M{SMSColumns.UsedAt: time.Now()}
	condition := s.condition()
	condition[SMSColumns.Type] = RegisterCode
	return s.update(condition, change)
}

func findSMS(m bson.M) (*SMS, error) {
	session := GetMongo()
	defer session.Close()

	var sms SMS
	err := session.DB(DBName).C(ColNameSMS).Find(m).One(&sms)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &sms, nil
}

// FindSMSByPhone  find sms code by phone
func FindSMSByPhone(phone string) (*SMS, error) {
	m := bson.M{SMSColumns.Phone: phone}
	return findSMS(m)
}
