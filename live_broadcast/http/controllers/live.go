package controllers

import (
	"platform/live_broadcast/http/services/broadcasts"

	"github.com/gin-gonic/gin"
)

// Live broadcasting controller
type Live struct {
	Base
}

// Start start to broadcast
func (r *Live) Start(c *gin.Context) {
	broadcasts.ServeWS(c)
}

// End end broadcast
func (r *Live) End(c *gin.Context) {

}

// Join 表示一个用户进入了直播间
func (r *Live) Join(c *gin.Context) {
	broadcasts.ServeWS(c)
}

// Leave leave live room
func (r *Live) Leave(c *gin.Context) {
	// stop ws
	// remove consumer channel of room topic
}
