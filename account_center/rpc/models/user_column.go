package models

type userColumn struct {
	ID       string
	Name     string
	Nickname string
	Password string
	Salt     string
	Token    string
	HeadImg  string
	RegTime  string
}

var UserColumns userColumn

func init() {
	UserColumns = userColumn{
		ID:       "_id",
		Name:     "name",
		Nickname: "nickname",
		Password: "password",
		Salt:     "salt",
		Token:    "token",
		HeadImg:  "headImg",
		RegTime:  "regTime",
	}
}
