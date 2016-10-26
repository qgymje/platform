package users

import (
	"time"

	"platform/account_center/rpc/models"
	"platform/commons/codes"
)

// UserInfo service level user info object
type UserInfo struct {
	ID         string
	Phone      string
	Email      string
	Nickname   string
	Password   string
	Token      string
	Avatar     string
	Level      int64
	FollowNum  int64
	Popularity int64
	CreatedAt  time.Time

	errorCode codes.ErrorCode
}

// NewUserInfo create a new user
func NewUserInfo() *UserInfo {
	return new(UserInfo)
}

// ErrorCode implements ErrorCoder
func (u *UserInfo) ErrorCode() codes.ErrorCode {
	return u.errorCode
}

func (u *UserInfo) formatUserInfo(user *models.User) {
	u.ID = user.GetID()
	u.Phone = user.Phone
	u.Email = user.Email
	u.Token = user.Token
	u.Nickname = user.Nickname
	u.Avatar = user.Avatar
	u.Level = user.Level
	u.FollowNum = user.FollowNum
	u.Popularity = user.Popularity
	u.CreatedAt = user.CreatedAt
	u.errorCode = codes.ErrorCodeSuccess
}

// GetByToken get user by token
func (u *UserInfo) GetByToken(token string) error {
	user, err := models.FindUserByToken(token)
	if err != nil {
		u.errorCode = codes.ErrorCodeInvalidToken
		return err
	}
	u.formatUserInfo(user)
	return nil
}

// GetByID get user by id
func (u *UserInfo) GetByID(userID string) error {
	user, err := models.FindUserByID(userID)
	if err != nil {
		u.errorCode = codes.ErrorCodeUserNotFound
		return err
	}
	u.formatUserInfo(user)
	return nil
}
