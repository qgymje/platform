package controllers

import (
	"net/http"
	"platform/commons/codes"
	"platform/live_broadcast/http/services/broadcasts"
	"platform/utils"

	"platform/commons/grpc_clients/room"
	pbroom "platform/commons/protos/room"

	"github.com/gin-gonic/gin"
)

// Live broadcasting controller
type Live struct {
	Base
}

// Start create a broadcast
func (l *Live) Start(c *gin.Context) {
	var errorCode codes.ErrorCode
	l.userInfo, errorCode = l.validUserInfo(c)
	if l.userInfo == nil {
		respformat := l.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewRoom(l.getRoomRPCAddress())
	userInfo := pbroom.User{UserID: l.userInfo.UserID}
	reply, err := roomClient.Start(&userInfo)
	if err != nil {
		respformat := l.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := l.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// End  end a broadcast
func (l *Live) End(c *gin.Context) {
	var errorCode codes.ErrorCode
	l.userInfo, errorCode = l.validUserInfo(c)
	if l.userInfo == nil {
		respformat := l.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewRoom(l.getRoomRPCAddress())
	userInfo := pbroom.User{UserID: l.userInfo.UserID}
	reply, err := roomClient.End(&userInfo)
	if err != nil {
		respformat := l.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := l.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Enter enter a broadcast
func (l *Live) Enter(c *gin.Context) {
	var errorCode codes.ErrorCode
	l.userInfo, errorCode = l.validUserInfo(c)
	if l.userInfo == nil {
		respformat := l.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewRoom(l.getRoomRPCAddress())
	userRoom := &pbroom.UserRoom{UserID: l.userInfo.UserID, BroadcastID: l.getBroadcastID(c)}
	reply, err := roomClient.Enter(userRoom)
	if err != nil {
		respformat := l.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := l.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Leave leave live room
func (l *Live) Leave(c *gin.Context) {
	var errorCode codes.ErrorCode
	l.userInfo, errorCode = l.validUserInfo(c)
	if l.userInfo == nil {
		respformat := l.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewRoom(l.getRoomRPCAddress())
	userRoom := &pbroom.UserRoom{UserID: l.userInfo.UserID, BroadcastID: l.getBroadcastID(c)}
	reply, err := roomClient.Leave(userRoom)
	if err != nil {
		respformat := l.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := l.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Join 表示一个用户进入了直播间
func (l *Live) Join(c *gin.Context) {
	utils.Dump(c.Request)
	broadcasts.ServeWS(c)
}
