package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CampaignStatus status of campaign
type CampaignStatus int

const (
	// CampaignStatusUnvalid  unvalid
	CampaignStatusUnvalid CampaignStatus = iota + 1
)

// Campaign  advertising campaign
type Campaign struct {
	CampaignID  bson.ObjectId `bson:"_id"`
	MerchantID  bson.ObjectId `bson:"merchant_id"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	StartTime   time.Time     `bson:"start_time"`
	EndTime     time.Time     `bson:"end_time"`
	Status      int           `bson:"status"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
	DeletedAt   time.Time     `bson:"deleted_at"`
}
