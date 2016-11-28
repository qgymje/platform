package controllers

import (
	"net/http"
	"platform/commons/codes"
	"platform/commons/grpc_clients/chat"
	pb "platform/commons/protos/chat"

	"github.com/gin-gonic/gin"
)

// Chat chat controller
type Chat struct {
	Base
}

// List chat list
func (ch *Chat) List(c *gin.Context) {

}

// Create create a chat
func (ch *Chat) Create(c *gin.Context) {
	var errorCode codes.ErrorCode
	ch.userInfo, errorCode = ch.validUserInfo(c)
	if ch.userInfo == nil {
		respformat := ch.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	cc := chatClient.NewChat(ch.getChatPCAddress())
	config := &pb.Creator{
		UserID:  ch.userInfo.UserID,
		Members: ch.getMembers(c),
		Name:    ch.getName(c),
	}
	reply, err := cc.Create(config)
	if err != nil {
		respformat := ch.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := ch.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Send send a message
func (ch *Chat) Send(c *gin.Context) {
	var errorCode codes.ErrorCode
	ch.userInfo, errorCode = ch.validUserInfo(c)
	if ch.userInfo == nil {
		respformat := ch.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

}
