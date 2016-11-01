package controllers

import (
	"log"
	"net/http"
	"platform/commons/codes"
	"platform/live_broadcast/http/services/broadcasts"
	"platform/utils"

	"platform/commons/grpc_clients/room"
	pbroom "platform/commons/protos/room"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	userInfo := pbroom.User{
		UserID: l.userInfo.UserID,
		TypeID: int32(l.getTypeID(c)),
	}
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
	userRoom := &pbroom.UserRoom{
		UserID:      l.userInfo.UserID,
		BroadcastID: l.getBroadcastID(c),
		TypeID:      int32(l.getTypeID(c)),
		Username:    l.userInfo.Nickname,
		Level:       l.userInfo.Level,
	}
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
	userRoom := &pbroom.UserRoom{
		UserID:      l.userInfo.UserID,
		BroadcastID: l.getBroadcastID(c),
		TypeID:      int32(l.getTypeID(c)),
		Username:    l.userInfo.Nickname,
		Level:       l.userInfo.Level,
	}
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Join 表示一个用户进入了直播间
func (l *Live) Join(c *gin.Context) {
	var errorCode codes.ErrorCode
	l.userInfo, errorCode = l.validUserInfo(c)
	if l.userInfo == nil {
		respformat := l.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	utils.Dump(l.userInfo)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		// 正式环境下需根据配置文件读取url来做判断
		return true
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("new client:", conn.RemoteAddr())

	broadcastID := c.Param("broadcast_id")
	topic := "broadcast_" + broadcastID
	channel := l.userInfo.UserID

	nsqd := utils.GetConf().GetString("nsq.nsqd")
	nsqlookupds := utils.GetConf().GetStringSlice("nsq.nsqlookupd")
	client := broadcasts.NewClient(conn, nsqlookupds, nsqd, topic, channel)
	go client.WritePump()
	client.ReadPump()
}
