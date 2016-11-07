package controllers

import (
	"net/http"
	"platform/commons/codes"

	"platform/commons/grpc_clients/gift"
	pbgift "platform/commons/protos/gift"

	"github.com/gin-gonic/gin"
)

// Gift gift controller
type Gift struct {
	Base
}

// List gift list
func (g *Gift) List(c *gin.Context) {
	var errorCode codes.ErrorCode
	g.userInfo, errorCode = g.validUserInfo(c)
	if g.userInfo == nil {
		respformat := g.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	pageNum := g.getPageNum(c)
	pageSize := g.getPageSize(c)

	page := &pbgift.Page{
		Num:  int32(pageNum),
		Size: int32(pageSize),
	}

	gc := giftClient.NewGift(g.getGiftPCAddress())
	reply, err := gc.List(page)

	if err != nil {
		respformat := g.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := g.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Send send a gift
func (g *Gift) Send(c *gin.Context) {

}
