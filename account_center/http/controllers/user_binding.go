package controllers

import (
	"platform/commons/codes"
	pb "platform/commons/protos/user"

	"github.com/gin-gonic/gin"
)

type LoginBinding struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`

	errorCode codes.ErrorCode
	config    *pb.LoginInfo
}

func NewLoginBinding(c *gin.Context) (form *LoginBinding, err error) {
	form = &LoginBinding{
		config: &pb.LoginInfo{},
	}
	if err = c.Bind(form); err != nil {
		form.errorCode = codes.ErrorCodeMissParameters
		return form, err
	}
	return
}

func (b *LoginBinding) ErrorCode() codes.ErrorCode {
	return b.errorCode
}

func (b *LoginBinding) Config() *pb.LoginInfo {
	b.config.Name = b.Name
	b.config.Password = b.Password
	return b.config
}

type RegisterBinding struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"passwd" binding:"required"`
	Nickname string `form:"nickname"`

	errorCode codes.ErrorCode
	config    *pb.RegisterInfo
}

func NewRegisterBinding(c *gin.Context) (form *RegisterBinding, err error) {
	form = &RegisterBinding{
		config: &pb.RegisterInfo{},
	}
	if err = c.Bind(form); err != nil {
		form.errorCode = codes.ErrorCodeMissParameters
		return form, err
	}
	return
}

func (b *RegisterBinding) ErrorCode() codes.ErrorCode {
	return b.errorCode
}

func (b *RegisterBinding) Config() *pb.RegisterInfo {
	b.config.Name = b.Name
	b.config.Password = b.Password
	b.config.Nickname = b.Nickname
	return b.config
}
