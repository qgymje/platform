package controllers

import (
	"net/http"
	"platform/commons/codes"
	"platform/commons/grpc_clients/profile"
	pb "platform/commons/protos/profile"

	"github.com/gin-gonic/gin"
)

// Friend friend controller
type Friend struct {
	Base
}

// List friend list
func (f *Friend) List(c *gin.Context) {

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
