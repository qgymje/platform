package controllers

import (
	"log"
	"net/http"
	"platform/commons/codes"
	"platform/commons/grpc_clients/user"
	pb "platform/commons/protos/user"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

type User struct {
	Base
	userRPCAddress string
}

func (u *User) getUserRPCAddress() string {
	if u.userRPCAddress != "" {
		return u.userRPCAddress
	}

	host := utils.GetConf().GetString("app.rpc_host")
	port := utils.GetConf().GetString("app.rpc_port")
	u.userRPCAddress = host + port
	return u.userRPCAddress
}

func (u *User) Auth(c *gin.Context) {
	token, errorCode := u.getToken(c)
	if errorCode != codes.ErrorCodeSuccess {
		respformat := u.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	pbToken := pb.Token{Token: token}
	auth := userClient.NewUser(u.getUserRPCAddress())
	if reply, err := auth.Auth(&pbToken); err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	} else {
		respformat := u.Response(c, codes.ErrorCodeSuccess, reply)
		c.JSON(http.StatusOK, respformat)
		return
	}
}

// Register 表示用户注册操作
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

// Login 表示用户登录操作
func (u *User) Login(c *gin.Context) {
	form, err := NewLoginBinding(c)
	if err != nil {
		respformat := u.Response(c, form.ErrorCode(), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	ul := userClient.NewUser(u.getUserRPCAddress())
	if reply, err := ul.Login(form.Config()); err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	} else {
		respformat := u.Response(c, codes.ErrorCodeSuccess, reply)
		c.JSON(http.StatusOK, respformat)
		return
	}
}

// Logout 表示用户登出操作
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
	} else {
		respformat := u.Response(c, codes.ErrorCodeSuccess, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
}

func (u *User) Info(c *gin.Context) {
	log.Println(u.apiVersion(c))
	userID := u.GetUserID(c)
	info := userClient.NewUser(u.getUserRPCAddress())
	if reply, err := info.Info(&pb.UserID{UserID: userID}); err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	} else {
		u.RemovePBUserInfoToken(reply)
		respformat := u.Response(c, codes.ErrorCodeSuccess, reply)
		c.JSON(http.StatusOK, respformat)
		return
	}
}
