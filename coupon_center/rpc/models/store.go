package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Store store infomation
type Store struct {
	StoreID    bson.ObjectId `bson:"_id"`
	MerchantID bson.ObjectId `bson:"merchant_id"`
	Name       string        `bson:"name"`
	Country    string        `bson:"country"`
	Province   string        `bson:"province"`
	City       string        `bson:"city"`
	Address    string        `bson:"address"`
	Longitude  float64       `bson:"longitude"`
	Latitude   float64       `bson:"latitude"`
	CreatedAt  time.Time     `bson:"created_at"`
	UpdatedAt  time.Time     `bson:"updated_at"`
	DeletedAt  time.Time     `bson:"deleted_at"`
}
