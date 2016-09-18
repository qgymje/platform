package users

import (
	"encoding/json"
	"errors"
	"platform/account_center/rpc/models"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/utils"
)

var (
	// ErrLogin 登录失败
	ErrLogin = errors.New("login fail, password error")
)

// Login 用于登录操作
type Login struct {
	name      string
	password  *Password
	userModel *models.User

	errorCode codes.ErrorCode
}

// LoginConfig 定义登录需要的数据
type LoginConfig struct {
	Name     string
	Password string
}

// NewLogin 用于创建Login的业务对象
func NewLogin(config *LoginConfig) *Login {
	l := new(Login)
	l.name = config.Name
	l.password = NewPassword(config.Password)
	l.userModel = &models.User{}
	l.errorCode = codes.ErrorCodeSuccess
	return l
}

// NewLoginByRawData 用于内部创建对象
func newLoginByRawData(name, password string) *Login {
	config := LoginConfig{
		Name:     name,
		Password: password,
	}
	return NewLogin(&config)
}

// ErrorCode 统一的错误码处理
func (s *Login) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do 用于处理登录操作业务流程
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
	s.userModel, err = models.FindUserByName(s.name)
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

// GetUserInfo 获取用户信息
func (s *Login) GetUserInfo() (*UserInfo, error) {
	u := &UserInfo{
		ID:        s.userModel.GetID(),
		Name:      s.userModel.Name,
		Nickname:  s.userModel.Nickname,
		HeadImg:   s.userModel.HeadImg,
		Token:     s.userModel.Token,
		RegTime:   s.userModel.RegTime,
		errorCode: s.ErrorCode(),
	}
	return u, nil
}

// Topic 发送给NSQ的话题
func (s *Login) Topic() string {
	return queues.TopicUserLogin.String()
}

// Message 发送给NSQ的消息内容
func (s *Login) Message() []byte {
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

func (s *Login) notify() error {
	return Publish(s)
}
