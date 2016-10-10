// Package models  data access layer
package models

import mgo "gopkg.in/mgo.v2"

var mongoSession *mgo.Session

// ErrNotFound not found error
var ErrNotFound = mgo.ErrNotFound

// DBName db name
const DBName = "coupon_center"

// ColNameMerchant collection name
const ColNameMerchant = "merchants"

// ColNameAccount collection name
const ColNameAccount = "accounts"

// ColNameStore collection name
const ColNameStore = "stores"

// ColNameCampaign collection name
const ColNameCampaign = "compaigns"

// ColNameCoupon collection name
const ColNameCoupon = "coupons"

// InitMongodb init mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

// GetMongo generate  a session copy
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}
