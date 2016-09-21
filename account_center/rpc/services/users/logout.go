package users

import (
	"errors"
	"platform/account_center/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

var (
	// ErrTokenUnvalid unvalid tokne
	ErrTokenUnvalid = errors.New("token不正确")
	// ErrTokenUpdate update token failed
	ErrTokenUpdate = errors.New("token更新出错")
)

// Logout 登出操作
type Logout struct {
	token     string
	userModel *models.User

	errorCode         codes.ErrorCode
	ensureDidFindUser bool
}

// NewLogout varify token and compare the claims ?
func NewLogout(token string) *Logout {
	s := new(Logout)
	s.token = token
	return s
}

// ErrorCode implement ErrorCoder
func (s *Logout) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do the logout job
func (s *Logout) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.Logout.Do error: %v", err)
		}
	}()

	if err = s.findUserByToken(); err != nil {
		s.errorCode = codes.ErrorCodeInvalidToken
		return ErrTokenUnvalid
	}

	if err = s.removeToken(); err != nil {
		s.errorCode = codes.ErrorCodeUpdateTokenFail
		return ErrTokenUpdate
	}

	return
}

func (s *Logout) findUserByToken() (err error) {
	s.userModel, err = models.FindUserByToken(s.token)
	if err != nil {
		return
	}
	s.ensureDidFindUser = true
	return nil
}

func (s *Logout) removeToken() error {
	return s.userModel.RemoveToken()
}
