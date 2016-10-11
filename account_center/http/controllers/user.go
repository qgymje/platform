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
	regInfo := &pb.RegisterInfo{
		Account:         u.getAccount(c),
		Password:        u.getPassword(c),
		PasswordConfirm: u.getPasswordConfirm(c),
		Nickname:        u.getNickname(c),
	}

	uc := userClient.NewUser(u.getUserRPCAddress())
	reply, err := uc.Register(regInfo)
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
	loginInfo := &pb.LoginInfo{
		Account:  u.getAccount(c),
		Password: u.getPassword(c),
	}

	ul := userClient.NewUser(u.getUserRPCAddress())

	reply, err := ul.Login(loginInfo)
	if err != nil {
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
	_, err := lo.Logout(&pbToken)
	if err != nil {
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
	userID := u.getUserID(c)
	info := userClient.NewUser(u.getUserRPCAddress())

	reply, err := info.Info(&pb.UserID{UserID: userID})
	if err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	u.removePBUserInfoToken(reply)
	u.removePBUserInfoPhone(reply)
	u.removePBUserInfoEmail(reply)
	respformat := u.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// List user list
func (u *User) List(c *gin.Context) {
	client := userClient.NewUser(u.getUserRPCAddress())
	query := &pb.UserQuery{
		IDs: u.getIDs(c),
	}
	reply, err := client.List(query)
	if err != nil {
		respformat := u.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	for i := range reply.Users {
		u.removePBUserInfoToken(reply.Users[i])
	}

	data := map[string]interface{}{
		"list": reply.Users,
	}
	respformat := u.Response(c, codes.ErrorCodeSuccess, data)
	c.JSON(http.StatusOK, respformat)
	return
}
