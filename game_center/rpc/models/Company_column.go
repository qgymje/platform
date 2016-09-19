package models

type _CompanyColumn struct {
	CreatedAt string
	ID        string
	Loactions string
	Name      string
	UpdatedAt string
	UserID    string
	Valid     string
}

// CompanyColumns company columns name
var CompanyColumns _CompanyColumn

func init() {
	CompanyColumns.CreatedAt = "created_at"
	CompanyColumns.ID = "_id"
	CompanyColumns.Loactions = "locations"
	CompanyColumns.Name = "name"
	CompanyColumns.UpdatedAt = "updated_at"
	CompanyColumns.UserID = "userID"
	CompanyColumns.Valid = "valid"

}
