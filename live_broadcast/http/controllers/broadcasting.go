package controllers

import (
	"platform/live_broadcast/http/services/broadcasts2"

	"github.com/gin-gonic/gin"
)

type Broadcasting struct {
	Base
}

// Join 表示一个用户进入了直播间
func (r *Broadcasting) Join(c *gin.Context) {
	// 验证用户信息
	broadcasts2.ServeWS(c)
}

func (r *Broadcasting) Leave(c *gin.Context) {
	// 验证用户信息
}

// RecentBarrages 通过gRPC调用
func (r *Broadcasting) RecentBarrages(c *gin.Context) {
	//roomID := c.Param("room_id")
}

// CurrentAudienceNum 通过gRPC调用
func (r *Broadcasting) CurrentAudienceNum(c *gin.Context) {

}
