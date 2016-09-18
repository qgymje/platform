package users

import (
	"time"

	"platform/account_center/rpc/models"
	"platform/commons/codes"
)

type UserInfo struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Password string    `json:"-"`
	Token    string    `json:"token"`
	HeadImg  string    `json:"headImg"`
	RegTime  time.Time `json:"regTime"`

	errorCode codes.ErrorCode
}

func NewUserInfo() *UserInfo {
	return new(UserInfo)
}

func (u *UserInfo) ErrorCode() codes.ErrorCode {
	return u.errorCode
}

func (u *UserInfo) formatUserInfo(user *models.User) {
	u.ID = user.GetID()
	u.Name = user.Name
	u.Token = user.Token
	u.Nickname = user.Nickname
	u.RegTime = user.RegTime
	u.errorCode = codes.ErrorCodeSuccess
}

func (u *UserInfo) GetByToken(token string) error {
	user, err := models.FindUserByToken(token)
	if err != nil {
		u.errorCode = codes.ErrorCodeUserNotFound
		return err
	}
	u.formatUserInfo(user)
	return nil
}

func (u *UserInfo) GetByID(userID string) error {
	user, err := models.FindUserByID(userID)
	if err != nil {
		u.errorCode = codes.ErrorCodeUserNotFound
		return err
	}
	u.formatUserInfo(user)
	return nil
}
