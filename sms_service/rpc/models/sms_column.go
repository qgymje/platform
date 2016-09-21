package models

type _SMSColumn struct {
	Content   string
	Country   string
	CreatedAt string
	Phone     string
	Provider  string
	Type      string
	UsedAt    string
}

// SMSColumns sms columns name
var SMSColumns _SMSColumn

func init() {
	SMSColumns.Content = "content"
	SMSColumns.Country = "country"
	SMSColumns.CreatedAt = "created_at"
	SMSColumns.Phone = "phone"
	SMSColumns.Provider = "provider"
	SMSColumns.Type = "type"
	SMSColumns.UsedAt = "used_at"

}
