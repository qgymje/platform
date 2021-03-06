package controllers

import (
	"math"
	"net/http"

	"platform/commons/codes"
	"platform/commons/grpc_clients/room"
	pbroom "platform/commons/protos/room"

	"github.com/gin-gonic/gin"
)

// Room room controller
type Room struct {
	Base
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

	rc := roomClient.NewRoom(r.getRoomRPCAddress())
	roomRequest := &pbroom.CreateRequest{
		UserID:   r.userInfo.UserID,
		UserName: r.userInfo.Nickname,
		Name:     r.getName(c),
		Cover:    r.getCover(c),
	}
	reply, err := rc.Create(roomRequest)
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
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	pageNum := r.getPageNum(c)
	pageSize := r.getPageSize(c)

	rc := roomClient.NewRoom(r.getRoomRPCAddress())
	listRequest := &pbroom.ListRequest{
		Num:    int32(pageNum),
		Size:   int32(pageSize),
		Search: r.getSearch(c),
	}

	roomList, err := rc.List(listRequest)
	if err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	totalPage := int(math.Floor(float64(roomList.TotalNum) / float64(pageSize)))
	data := map[string]interface{}{
		"list":      roomList.Rooms,
		"page":      pageNum,
		"pageSize":  pageSize,
		"totalPage": totalPage,
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, data)
	c.JSON(http.StatusOK, respformat)
	return
}

// Info show user's room info
func (r *Room) Info(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	roomID := r.getRoomID(c)
	rc := roomClient.NewRoom(r.getRoomRPCAddress())
	userRoom := &pbroom.UserRoom{
		UserID: r.userInfo.UserID,
		RoomID: roomID,
	}
	info, err := rc.Info(userRoom)
	if err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, info)
	c.JSON(http.StatusOK, respformat)
	return
}

// Follow follow the room
func (r *Room) Follow(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	rc := roomClient.NewRoom(r.getRoomRPCAddress())
	userRoom := &pbroom.UserRoom{
		UserID: r.userInfo.UserID,
		RoomID: r.getRoomID(c),
	}
	reply, err := rc.Follow(userRoom)
	if err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Unfollow unfollow the room
func (r *Room) Unfollow(c *gin.Context) {
	var errorCode codes.ErrorCode
	r.userInfo, errorCode = r.validUserInfo(c)
	if r.userInfo == nil {
		respformat := r.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	rc := roomClient.NewRoom(r.getRoomRPCAddress())
	userRoom := &pbroom.UserRoom{
		UserID: r.userInfo.UserID,
		RoomID: r.getRoomID(c),
	}
	reply, err := rc.Unfollow(userRoom)
	if err != nil {
		respformat := r.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := r.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

type roomType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var roomTypes []roomType

func init() {
	roomTypes = []roomType{
		{1, "英雄联盟"},
		{2, "守望先锋"},
		{3, "炉石传说"},
		{4, "DOTA2"},
		{5, "魔兽世界"},
	}
}

// Types game types
func (r *Room) Types(c *gin.Context) {
	data := map[string]interface{}{
		"list": roomTypes,
	}
	respformat := r.Response(c, codes.ErrorCodeSuccess, data)
	c.JSON(http.StatusOK, respformat)
	return
}
