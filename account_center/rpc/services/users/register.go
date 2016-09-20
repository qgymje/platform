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

type RegisterConfig struct {
	LoginConfig
	Nickname string
}

// NewRegister 生成一个注册用户对象
func NewRegister(config *RegisterConfig) *Register {
	s := new(Register)
	s.Login = newLoginByRawData(config.Name, config.Password)
	s.Nickname = config.Nickname

	return s
}

// Do 做具体注册的操作
func (s *Register) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("services.users.Register.Do error: %v", err)
		}
	}()

	if err = s.findUser(); err != nil {
		s.errorCode = codes.ErrorCodeNameAlreadyExist
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
	if models.IsNameUsed(s.name) {
		return fmt.Errorf("username is used: %s", s.name)
	}
	return nil
}

// save 将数据保存到db
func (s *Register) saveUser() (err error) {
	if err = s.password.Valid(); err != nil {
		return
	}

	s.userModel.Name = s.name
	s.userModel.Salt = s.password.GenSalt()
	s.userModel.Password = s.password.GenPwd()
	s.userModel.Nickname = s.Nickname

	return s.userModel.Create()
}

func (s *Register) Topic() string {
	return queues.TopicUserRegister.String()
}

func (s *Register) Message() []byte {
	message := queues.MessageUserLogin{
		UserID:   s.userModel.GetID(),
		Name:     s.userModel.Name,
		Nickname: s.userModel.Nickname,
		HeadImg:  s.userModel.HeadImg,
		RegTime:  s.userModel.RegTime,
	}
	msg, _ := json.Marshal(&message)
	return msg
}

func (s *Register) notify() error {
	return notifier.Publish(s)
}
