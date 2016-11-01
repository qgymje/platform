package models

// Merchant merchant model object
type Merchant struct {
	ID int64 `orm:"column(id)"`
}

// TableName table name
func (Merchant) TableName() string {
	return TableNameMerchant
}
