package models

type userLoginColumn struct {
	UserID    string
	CreatedAt string
}

var UserLoginColumns userLoginColumn

func init() {
	UserLoginColumns = userLoginColumn{
		UserID:    "userID",
		CreatedAt: "createdAt",
	}
}
