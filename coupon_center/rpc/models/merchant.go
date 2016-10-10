package models

import "gopkg.in/mgo.v2/bson"

// Merchant merchant model object
type Merchant struct {
	MerchantID bson.ObjectId `bson:"_id"`
}
