package users

import (
	"platform/account_center/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

// Config the config of users
type Config struct {
	PageNum  int
	PageSize int
	Search   string
	IDs      []string
}

// Users represents an object of a list of user
type Users struct {
	userFinder *models.UserFinder

	errorCode codes.ErrorCode
}

// NewUsers create a new Users object
func NewUsers(config *Config) *Users {
	u := new(Users)
	u.userFinder = models.NewUserFinder().Limit(config.PageNum, config.PageSize).Search(config.Search).ByIDs(config.IDs)
	return u
}

// ErrorCode ErrorCoder implement
func (u *Users) ErrorCode() codes.ErrorCode {
	return u.errorCode
}

// Do do the NewGames query
func (u *Users) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("users.Users.Do error: %+v", err)
		}
	}()

	if err = u.find(); err != nil {
		if err == models.ErrNotFound {
			u.errorCode = codes.ErrorCodeUsersNotFound
		} else {
			u.errorCode = codes.ErrorCodeUserFinder
		}
		return
	}
	return
}

func (u *Users) find() error {
	return u.userFinder.Do()
}

// Users maps models user to service user
func (u *Users) Users() []*UserInfo {
	modelUsers := u.userFinder.Result()
	srvUsers := []*UserInfo{}
	for _, mUser := range modelUsers {
		srvUser := &UserInfo{}
		srvUser.formatUserInfo(mUser)
		srvUsers = append(srvUsers, srvUser)
	}
	return srvUsers
}

// Count return the total number
func (u *Users) Count() int64 {
	return u.userFinder.Count()
}
