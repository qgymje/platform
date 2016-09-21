package users

import (
	"encoding/json"
	"fmt"

	"platform/account_center/rpc/models"
	"platform/account_center/rpc/services/notifier"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/utils"
)

// Register 注册对象
type Register struct {
	*Login
	Nickname string
}

// RegisterConfig register info
type RegisterConfig struct {
	Account         string
	Password        string
	PasswordConfirm string
	Nickname        string
}

// NewRegister new register object
func NewRegister(config *RegisterConfig) *Register {
	s := new(Register)
	s.Login = newLoginByRawData(config.Account, config.Password)
	s.Nickname = config.Nickname

	return s
}

// Do do the actual register job
func (s *Register) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.Register.Do error: %v", err)
		}
	}()

	if err = s.findUser(); err != nil {
		return
	}

	if err = s.saveUser(); err != nil {
		s.errorCode = codes.ErrorCodeCreateUserFail
		return
	}

	if err = s.updateToken(); err != nil {
		s.errorCode = codes.ErrorCodeUpdateTokenFail
		return
	}

	if err = s.notify(); err != nil {
		s.errorCode = codes.ErrorCodeRegisterNotify
		return
	}

	return
}

func (s *Register) findUser() (err error) {
	s.determinLoginType()
	switch s.loginType {
	case email:
		if models.IsEmailUsed(s.account) {
			s.errorCode = codes.ErrorCodeEmailAlreadyExist
			return fmt.Errorf("email is used: %s", s.account)
		}
	case phone:
		if models.IsPhoneUsed(s.account) {
			s.errorCode = codes.ErrorCodePhoneAlreadyExist
			return fmt.Errorf("phone is used: %s", s.account)
		}
	}
	return nil
}

func (s *Register) saveUser() (err error) {
	if err = s.password.Valid(); err != nil {
		return
	}
	s.determinLoginType()

	switch s.loginType {
	case email:
		s.userModel.Email = s.account
	case phone:
		s.userModel.Phone = s.account
	}

	s.userModel.Salt = s.password.GenSalt()
	s.userModel.Password = s.password.GenPwd()
	s.userModel.Nickname = s.Nickname

	return s.userModel.Create()
}

// Topic implement notifier
func (s *Register) Topic() string {
	return queues.TopicUserRegister.String()
}

// Message implement notifier
func (s *Register) Message() []byte {
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

func (s *Register) notify() error {
	return notifier.Publish(s)
}
