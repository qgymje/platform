package models

type _UserLoginColumn struct {
	CreatedAt string
	UserID    string
}

// UserLoginColumns userlogin columns name
var UserLoginColumns _UserLoginColumn

func init() {
	UserLoginColumns.CreatedAt = "created_at"
	UserLoginColumns.UserID = "user_id"

}
