package controllers

import (
	"net/http"
	"platform/commons/codes"

	pbroom "platform/commons/protos/room"
	pbuser "platform/commons/protos/user"

	"github.com/gin-gonic/gin"
)

// Broadcast broadcasting
type Broadcast struct {
	Base
	userInfo *pbuser.UserInfo
}

// Start create a broadcast
func (r *Broadcast) Start(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewBroadcast(r.getBroadcastRPCAddress())
	userInfo := pbroom.User{UserID: r.userInfo.UserID}
	reply, err := roomClient.Start(&userInfo)
	if err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// End  end a broadcast
func (r *Broadcast) End(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewBroadcast(r.getBroadcastRPCAddress())
	userInfo := pbroom.User{UserID: r.userInfo.UserID}
	reply, err := roomClient.End(&userInfo)
	if err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// CurrentAudienceNum  current audience number
func (r *Broadcast) CurrentAudienceNum(c *gin.Context) {

}
