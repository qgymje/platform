package controllers

import (
	"net/http"
	"platform/commons/codes"
	"platform/commons/typeids"

	"platform/commons/grpc_clients/gift"
	"platform/commons/grpc_clients/profile"
	"platform/commons/grpc_clients/room"
	pbgift "platform/commons/protos/gift"
	pbprofile "platform/commons/protos/profile"
	pbroom "platform/commons/protos/room"

	"github.com/gin-gonic/gin"
)

// Gift gift controller
type Gift struct {
	Base
}

// Info single gift info query
func (g *Gift) Info(c *gin.Context) {
	var errorCode codes.ErrorCode
	g.userInfo, errorCode = g.validUserInfo(c)
	if g.userInfo == nil {
		respformat := g.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	giftID := &pbgift.GiftID{GiftID: g.getGiftID(c)}
	gc := giftClient.NewGift(g.getGiftPCAddress())
	reply, err := gc.Info(giftID)

	if err != nil {
		respformat := g.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := g.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return

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

func isRoomValid(room *pbroom.RoomInfo) bool {
	if room == nil {
		return false
	}

	if room.IsPlaying && room.Broadcast != nil {
		return true
	}

	return false
}

// Send send a gift
func (g *Gift) Send(c *gin.Context) {
	var errorCode codes.ErrorCode
	g.userInfo, errorCode = g.validUserInfo(c)
	if g.userInfo == nil {
		respformat := g.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	rc := roomClient.NewRoom(g.getRoomRPCAddress())
	userRoom := &pbroom.UserRoom{
		UserID: g.userInfo.UserID,
		RoomID: g.getRoomID(c),
	}
	roomInfo, err := rc.Info(userRoom)
	if err != nil {
		respformat := g.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	if !isRoomValid(roomInfo) {
		respformat := g.Response(c, rpcErrorFormat(codes.ErrorCodeBroadcastClosed.String()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	giftID := &pbgift.GiftID{GiftID: g.getGiftID(c)}
	gc := giftClient.NewGift(g.getGiftPCAddress())
	giftInfo, err := gc.Info(giftID)

	if err != nil {
		respformat := g.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	amm := &pbprofile.Ammount{
		UserID:    g.userInfo.UserID,
		SnowFlake: giftInfo.SnowFlake,
		SnowBall:  giftInfo.SnowBall,
		TypeID:    uint32(typeids.PaymentSendGift),
		TargetID:  giftInfo.GiftID,
	}
	pc := profileClient.NewProfile(g.getProfileRPCAddress())
	wd, err := pc.Withdraw(amm)
	if err != nil {
		respformat := g.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	msgID := wd.MsgID
	// gift send_gift
	sendGift := &pbgift.SendGift{
		GiftID:      giftInfo.GiftID,
		UserID:      g.userInfo.UserID,
		ToUserID:    roomInfo.UserID,
		BroadcastID: roomInfo.Broadcast.BroadcastID,
		MsgID:       msgID,
		Number:      1,
	}
	// maybe can use gc
	gc2 := giftClient.NewGift(g.getGiftPCAddress())
	sendReply, err := gc2.Send(sendGift)
	if err != nil {
		pc2 := profileClient.NewProfile(g.getProfileRPCAddress())
		msg := &pbprofile.Message{
			MsgID:  msgID,
			UserID: g.userInfo.UserID,
		}
		_, err2 := pc2.WithdrawRollback(msg)
		if err2 != nil {
			respformat := g.Response(c, rpcErrorFormat(err2.Error()), nil)
			c.JSON(http.StatusOK, respformat)
			return
		}

		respformat := g.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	// profile Withdraw commit
	msg := &pbprofile.Message{
		MsgID:  sendReply.MsgID,
		UserID: g.userInfo.UserID,
	}
	pc3 := profileClient.NewProfile(g.getProfileRPCAddress())
	_, err = pc3.WithdrawCommit(msg)
	if err != nil {
		respformat := g.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	// gift broadcast
	gc3 := giftClient.NewGift(g.getGiftPCAddress())
	sendGiftID := &pbgift.SendGiftID{
		SendGiftID: sendReply.SendGiftID,
		Username:   g.userInfo.Nickname,
	}
	broadcastReply, err := gc3.Broadcast(sendGiftID)
	if err != nil {
		respformat := g.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := g.Response(c, codes.ErrorCodeSuccess, broadcastReply)
	c.JSON(http.StatusOK, respformat)
	return
}
