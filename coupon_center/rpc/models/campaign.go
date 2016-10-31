package models

import "time"

// CampaignStatus status of campaign
type CampaignStatus int

const (
	// CampaignStatusUnvalid  unvalid
	CampaignStatusUnvalid CampaignStatus = iota + 1
)

// Campaign  advertising campaign
type Campaign struct {
	ID          int64 `orm:"column(id)"`
	MerchantID  int64 `orm:"column(merchant_id)"`
	Name        string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// TableName table name
func (Campaign) TableName() string {
	return TableNameCampaign
}
