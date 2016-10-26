// Package models  data access layer
package models

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// ErrObjectID error object id
	ErrObjectID = errors.New("not a valid objectID")
)

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

// ColNameUserCoupon user_coupons
const ColNameUserCoupon = "user_coupons"

// ColNameSendCoupon send_coupons
const ColNameSendCoupon = "send_coupons"

// ColNameTakeCoupon take_coupons
const ColNameTakeCoupon = "take_coupons"

// InitMongodb init mongodb
func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

// GetMongo generate  a session copy
func GetMongo() *mgo.Session {
	return mongoSession.Copy()
}

// StringToObjectID string to bson objectId
func StringToObjectID(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(string(id)) {
		return bson.ObjectId(""), ErrObjectID
	}
	return bson.ObjectIdHex(id), nil
}

// StringsToObjectIDs strings to bson objectIds
func StringsToObjectIDs(ids []string) ([]bson.ObjectId, error) {
	IDHexs := []bson.ObjectId{}
	for _, id := range ids {
		if !bson.IsObjectIdHex(string(id)) {
			return nil, ErrObjectID
		}
		IDHexs = append(IDHexs, bson.ObjectIdHex(string(id)))
	}
	return IDHexs, nil
}
