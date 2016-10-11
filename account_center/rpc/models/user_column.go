package models

type _UserColumn struct {
	Avatar     string
	CreatedAt  string
	DeletedAt  string
	Email      string
	FollowNum  string
	ID         string
	Level      string
	Nickname   string
	Password   string
	Phone      string
	Popularity string
	Salt       string
	Token      string
	UpdatedAt  string
}

// UserColumns user columns name
var UserColumns _UserColumn

func init() {
	UserColumns.Avatar = "avatar"
	UserColumns.CreatedAt = "created_at"
	UserColumns.DeletedAt = "deleted_at"
	UserColumns.Email = "email"
	UserColumns.FollowNum = "follow_num"
	UserColumns.ID = "_id"
	UserColumns.Level = "level"
	UserColumns.Nickname = "nickname"
	UserColumns.Password = "password"
	UserColumns.Phone = "phone"
	UserColumns.Popularity = "popularity"
	UserColumns.Salt = "salt"
	UserColumns.Token = "token"
	UserColumns.UpdatedAt = "updated_at"

}
