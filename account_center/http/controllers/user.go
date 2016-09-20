package controllers

import (
	"net/http"
	"platform/commons/codes"
	"platform/commons/grpc_clients/user"
	pb "platform/commons/protos/user"

	"github.com/gin-gonic/gin"
)

// User represent user actions
type User struct {
	Base
}

// Auth verify token
func (u *User) Auth(c *gin.Context) {
	token, errorCode := u.getToken(c)
	if errorCode != codes.ErrorCodeSuccess {
		respformat := u.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	pbToken := pb.Token{Token: token}
	auth := userClient.NewUser(u.getUserRPCAddress())
	var reply *pb.UserInfo
	var err error
	if reply, err = auth.Auth(&pbToken); err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	respformat := u.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Register  register action
func (u *User) Register(c *gin.Context) {
	form, err := NewRegisterBinding(c)
	if err != nil {
		respformat := u.Response(c, form.ErrorCode(), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	uc := userClient.NewUser(u.getUserRPCAddress())
	reply, err := uc.Register(form.Config())
	if err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := u.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Login user login
func (u *User) Login(c *gin.Context) {
	form, err := NewLoginBinding(c)
	if err != nil {
		respformat := u.Response(c, form.ErrorCode(), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	ul := userClient.NewUser(u.getUserRPCAddress())
	var reply *pb.UserInfo
	if reply, err = ul.Login(form.Config()); err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	respformat := u.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Logout  user logout
func (u *User) Logout(c *gin.Context) {
	token, errorCode := u.getToken(c)
	if errorCode != codes.ErrorCodeSuccess {
		respformat := u.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	pbToken := pb.Token{Token: token}
	lo := userClient.NewUser(u.getUserRPCAddress())
	if _, err := lo.Logout(&pbToken); err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := u.Response(c, codes.ErrorCodeSuccess, nil)
	c.JSON(http.StatusOK, respformat)
	return
}

// Info user info
func (u *User) Info(c *gin.Context) {
	userID := u.GetUserID(c)
	info := userClient.NewUser(u.getUserRPCAddress())
	var reply *pb.UserInfo
	var err error
	if reply, err = info.Info(&pb.UserID{UserID: userID}); err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	u.RemovePBUserInfoToken(reply)
	respformat := u.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}
