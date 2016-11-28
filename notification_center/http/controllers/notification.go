package controllers

import (
	"log"
	"net/http"
	"platform/commons/codes"
	"platform/live_broadcast/http/services/broadcasts"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

// Notification controller
type Notification struct {
	Base
}

// Notify notify
func (n *Notification) Notify(c *gin.Context) {
	var errorCode codes.ErrorCode
	n.userInfo, errorCode = n.validUserInfo(c)
	if n.userInfo == nil {
		respformat := n.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	utils.Dump(n.userInfo)

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

	nsqd := utils.GetConf().GetString("nsq.nsqd")
	nsqlookupds := utils.GetConf().GetStringSlice("nsq.nsqlookupd")
	client := broadcasts.NewClient(conn, nsqlookupds, nsqd, topic, channel)
	go client.WritePump()
	client.ReadPump()
}
