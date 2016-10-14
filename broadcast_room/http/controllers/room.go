package controllers

import (
	"net/http"

	"platform/commons/codes"
	pbroom "platform/commons/protos/room"
	pbuser "platform/commons/protos/user"

	"platform/commons/grpc_clients/room"
	"platform/commons/grpc_clients/user"

	"github.com/gin-gonic/gin"
)

// Room room controller
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

// Create create a broadcast room
func (r *Room) Create(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	roomClient := roomClient.NewRoom(r.getRoomRPCAddress())
	roomRequest := &pbroom.RoomRequest{
		UserID:     r.userInfo.UserID,
		Name:       r.getName(c),
		Channel:    r.getChannel(c),
		SubChannel: r.getSubChannel(c),
		Cover:      r.getCover(c),
	}
	reply, err := roomClient.Create(roomRequest)
	if err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// List list the rooms
func (r *Room) List(c *gin.Context) {

}

// Update update a room info
func (r *Room) Update(c *gin.Context) {

}
