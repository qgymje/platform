package controllers

import (
	"net/http"

	"platform/commons/codes"
	pbroom "platform/commons/protos/room"
	pbuser "platform/commons/protos/user"
	"platform/utils"

	"platform/commons/grpc_clients/room"
	"platform/commons/grpc_clients/user"

	"github.com/gin-gonic/gin"
)

type Room struct {
	Base

	userInfo *pbuser.UserInfo
}

func (r *Room) validUserInfo(c *gin.Context) (*pbuser.UserInfo, codes.ErrorCode) {
	token, errorCode := r.getToken(c)
	if errorCode != codes.ErrorCodeSuccess {
		return nil, errorCode
	}
	pbToken := pbuser.Token{Token: token}
	auth := userClient.NewUser(r.getUserRPCAddress())

	var err error
	var userInfo *pbuser.UserInfo
	if userInfo, err = auth.Auth(&pbToken); err != nil {
		return nil, rpcErrorFormat(err.Error())
	}
	return userInfo, codes.ErrorCodeSuccess
}

func (r *Room) Create(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	utils.Dump("userinfo:", r.userInfo)
	var err error
	form, err := NewRoomBinding(c)
	if err != nil {
		respformat := r.Response(c, form.ErrorCode(), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewRoom(r.getRoomRPCAddress())
	var reply *pbroom.RoomResponse
	roomRequest := form.Config()
	roomRequest.UserID = r.userInfo.UserID
	if reply, err = roomClient.Create(roomRequest); err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Start 表示一个主播开始了直播
func (r *Room) Start(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewRoom(r.getRoomRPCAddress())
	userInfo := pbroom.User{UserID: r.userInfo.UserID}
	var reply *pbroom.Status
	var err error
	if reply, err = roomClient.Start(&userInfo); err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// End 表示主播结束直播
func (r *Room) End(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewRoom(r.getRoomRPCAddress())
	userInfo := pbroom.User{UserID: r.userInfo.UserID}
	var reply *pbroom.EndResponse
	var err error
	if reply, err = roomClient.End(&userInfo); err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Send 表示发送弹幕
func (r *Room) Barrage(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
}
