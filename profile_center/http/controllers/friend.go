package controllers

import (
	"net/http"
	"platform/commons/codes"
	"platform/commons/grpc_clients/profile"
	"platform/commons/grpc_clients/user"
	pb "platform/commons/protos/profile"
	pbuser "platform/commons/protos/user"

	"github.com/gin-gonic/gin"
)

// Friend friend controller
type Friend struct {
	Base
}

// Recommend recommend friends
func (f *Friend) Recommend(c *gin.Context) {

}

// List friend list
func (f *Friend) List(c *gin.Context) {
	var errorCode codes.ErrorCode
	f.userInfo, errorCode = f.validUserInfo(c)
	if f.userInfo == nil {
		respformat := f.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	p := profileClient.NewProfile(f.getProfileRPCAddress())
	config := &pb.Message{
		UserID: f.userInfo.UserID,
	}
	reply, err := p.FriendList(config)
	if err != nil {
		respformat := f.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	uc := userClient.NewUser(f.getUserRPCAddress())
	userQuery := &pbuser.UserQuery{
		Num:  1,
		Size: int32(len(reply.FriendIDs)),
		IDs:  reply.FriendIDs,
	}
	users, err := uc.List(userQuery)
	if err != nil {
		respformat := f.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	for i := range users.Users {
		f.removePBUserInfoToken(users.Users[i])
		f.removePBUserInfoPhone(users.Users[i])
		f.removePBUserInfoEmail(users.Users[i])
	}

	respformat := f.Response(c, codes.ErrorCodeSuccess, users)
	c.JSON(http.StatusOK, respformat)
	return
}

// Request add friend
func (f *Friend) Request(c *gin.Context) {
	var errorCode codes.ErrorCode
	f.userInfo, errorCode = f.validUserInfo(c)
	if f.userInfo == nil {
		respformat := f.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	p := profileClient.NewProfile(f.getProfileRPCAddress())
	config := &pb.Request{
		FromUserID: f.userInfo.UserID,
		ToUserID:   f.getUserID(c),
		Message:    f.getMessage(c),
	}
	reply, err := p.FriendRequest(config)
	if err != nil {
		respformat := f.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := f.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Agree agree
func (f *Friend) Agree(c *gin.Context) {

}

// Refuse refuse
func (f *Friend) Refuse(c *gin.Context) {

}
