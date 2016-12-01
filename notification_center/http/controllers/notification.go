package controllers

import (
	"log"
	"net/http"
	"platform/commons/codes"
	"platform/notification_center/http/services"
	"platform/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Notification controller
type Notification struct {
	Base
}

// Notify notify
func (n *Notification) Notify(c *gin.Context) {
	var errorCode codes.ErrorCode
	n.userInfo, errorCode = n.validUserInfo(c)
	utils.Dump(n.userInfo)
	if n.userInfo == nil {
		respformat := n.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	//utils.Dump(n.userInfo)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		// 正式环境下需根据配置文件读取url来做判断
		return true
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println("new client:", conn.RemoteAddr())

	chatConsumer := services.NewChat(n.userInfo.UserID)
	//notificationConsumer := services.NewNotification(n.userInfo.UserID)
	consumers := []services.Consumer{chatConsumer}
	nsqsession := services.NewNSQSession(consumers)
	client := services.NewClient(conn, nsqsession)
	go client.WritePump()
	client.ReadPump()
}
