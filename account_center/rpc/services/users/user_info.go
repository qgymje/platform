package users

import (
	"time"

	"platform/account_center/rpc/models"
	"platform/commons/codes"
)

type UserInfo struct {
	ID        string    `json:"user_id"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"-"`
	Token     string    `json:"token"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`

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
	u.Phone = user.Phone
	u.Email = user.Email
	u.Token = user.Token
	u.Nickname = user.Nickname
	u.CreatedAt = user.CreatedAt
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
