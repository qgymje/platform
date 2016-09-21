package users

import (
	"encoding/json"
	"errors"
	"platform/account_center/rpc/models"
	"platform/account_center/rpc/services/notifier"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/utils"

	"github.com/astaxie/beego/validation"
)

var (
	// ErrLogin  login failed
	ErrLogin = errors.New("login fail, password error")
)

//go:generate stringer -type=loginType
type loginType int

const (
	phone loginType = iota + 1
	email
)

// Login  login object
type Login struct {
	account   string
	password  *Password
	userModel *models.User

	loginType loginType

	valid     *validation.Validation
	errorCode codes.ErrorCode
}

// LoginConfig  login configuration
type LoginConfig struct {
	Account  string
	Password string
}

// NewLogin new a login object by config
func NewLogin(config *LoginConfig) *Login {
	l := new(Login)
	l.account = config.Account
	l.password = NewPassword(config.Password)
	l.userModel = &models.User{}
	l.valid = &validation.Validation{}
	l.errorCode = codes.ErrorCodeSuccess
	return l
}

// NewLoginByRawData new login object by raw data
func newLoginByRawData(account, password string) *Login {
	config := LoginConfig{
		Account:  account,
		Password: password,
	}
	return NewLogin(&config)
}

// ErrorCode  implement ErrorCoder
func (s *Login) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

func (s *Login) determinLoginType() {
	if s.isEmail() {
		s.loginType = email
	}

	if s.isPhone() {
		s.loginType = phone
	}
}

func (s *Login) isEmail() bool {
	if v := s.valid.Email(s.account, "email"); v.Ok {
		return true
	}
	return false
}

func (s *Login) isPhone() bool {
	if v := s.valid.Mobile(s.account, "phone"); v.Ok {
		return true
	}
	return false
}

// Do do the login logic
func (s *Login) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.Login.Do error: %v", err)
		}
	}()

	if err = s.findUser(); err != nil {
		s.errorCode = codes.ErrorCodeUserNotFound
		return
	}

	if err = s.validPassword(); err != nil {
		s.errorCode = codes.ErrorCodeLoginFailed
		return
	}

	if err = s.updateToken(); err != nil {
		s.errorCode = codes.ErrorCodeUpdateTokenFail
		return
	}

	if err = s.notify(); err != nil {
		s.errorCode = codes.ErrorCodeLoginNotify
		return
	}

	s.logLogin()

	return
}

func (s *Login) findUser() (err error) {
	s.determinLoginType()
	switch s.loginType {
	case email:
		s.userModel, err = models.FindUserByEmail(s.account)
	case phone:
		s.userModel, err = models.FindUserByPhone(s.account)
	}
	return
}

func (s *Login) validPassword() (err error) {
	s.password.SetSalt(s.userModel.Salt)
	if !s.password.IsEncryptedSame(s.userModel.Password) {
		return ErrLogin
	}
	return
}

func (s *Login) updateToken() (err error) {
	claims := map[string]interface{}{
		"id": s.userModel.ID,
	}
	token, err := NewToken().Generate(claims)
	if err != nil {
		return
	}
	err = s.userModel.UpdateToken(token)

	return
}

func (s *Login) logLogin() {
	userLogin := &models.UserLogin{
		UserID: s.userModel.ID,
	}
	_ = userLogin.Create()
}

// GetUserInfo  get user info
func (s *Login) GetUserInfo() (*UserInfo, error) {
	u := &UserInfo{
		ID:        s.userModel.GetID(),
		Phone:     s.userModel.Phone,
		Email:     s.userModel.Email,
		Nickname:  s.userModel.Nickname,
		Avatar:    s.userModel.Avatar,
		Token:     s.userModel.Token,
		CreatedAt: s.userModel.CreatedAt,
		errorCode: s.ErrorCode(),
	}
	return u, nil
}

// Topic implement notifier
func (s *Login) Topic() string {
	return queues.TopicUserLogin.String()
}

// Message implement notifier
func (s *Login) Message() []byte {
	message := queues.MessageUserLogin{
		UserID:    s.userModel.GetID(),
		Phone:     s.userModel.Phone,
		Email:     s.userModel.Email,
		Nickname:  s.userModel.Nickname,
		Avatar:    s.userModel.Avatar,
		CreatedAt: s.userModel.CreatedAt,
	}

	msg, _ := json.Marshal(&message)
	return msg
}

func (s *Login) notify() error {
	return notifier.Publish(s)
}
