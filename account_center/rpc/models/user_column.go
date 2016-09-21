package models

type _UserColumn struct {
	Avatar    string
	CreatedAt string
	Email     string
	ID        string
	Nickname  string
	Password  string
	Phone     string
	Salt      string
	Token     string
}

// UserColumns user columns name
var UserColumns _UserColumn

func init() {
	UserColumns.Avatar = "avatar"
	UserColumns.CreatedAt = "created_at"
	UserColumns.Email = "email"
	UserColumns.ID = "_id"
	UserColumns.Nickname = "nickname"
	UserColumns.Password = "password"
	UserColumns.Phone = "phone"
	UserColumns.Salt = "salt"
	UserColumns.Token = "token"

}
