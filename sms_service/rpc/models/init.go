// Package models  data access layer
package models

import mgo "gopkg.in/mgo.v2"

var mongoSession *mgo.Session

// ErrNotFound not found error
var ErrNotFound = mgo.ErrNotFound

// DBName db name
const DBName = "sms_services"

// ColNameSMS sms collection name
const ColNameSMS = "sms"

// InitMongodb init mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

// GetMongo generate  a session copy
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}
