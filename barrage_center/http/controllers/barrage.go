package controllers

import (
	"net/http"
	"platform/commons/codes"
	"platform/commons/grpc_clients/barrage"
	pb "platform/commons/protos/barrage"
	"platform/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Barrage barrage controller
type Barrage struct {
	Base
}

// Create a barrage
func (b *Barrage) Create(c *gin.Context) {
	var errorCode codes.ErrorCode
	b.userInfo, errorCode = b.validUserInfo(c)
	if b.userInfo == nil {
		respformat := b.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	bc := barrageClient.NewBarrage(b.getBarrageRPCAddress())
	content := &pb.Content{
		TypeID:      int32(b.getTypeID(c)),
		BroadcastID: b.getBroadcastID(c),
		UserID:      b.userInfo.UserID,
		Text:        b.getText(c),
		Username:    b.userInfo.Nickname,
		Level:       b.userInfo.Level,
		CreatedAt:   time.Now().Unix(),
	}

	reply, err := bc.Send(content)
	if err != nil {
		respformat := b.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := b.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// List current list defined by client
// duration = 10m, num = 100
func (b *Barrage) List(c *gin.Context) {
	var errorCode codes.ErrorCode
	b.userInfo, errorCode = b.validUserInfo(c)
	if b.userInfo == nil {
		respformat := b.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	bc := barrageClient.NewBarrage(b.getBarrageRPCAddress())
	broadcast := &pb.Broadcast{
		BroadcastID: b.getBroadcastID(c),
		StartTime:   b.getStartTime(c),
		EndTime:     b.getEndTime(c),
		Num:         int32(b.getPageNum(c)),
		Size:        int32(b.getPageSize(c)),
	}
	utils.Dump(broadcast)
	reply, err := bc.List(broadcast)
	if err != nil {
		respformat := b.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := b.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return

}
