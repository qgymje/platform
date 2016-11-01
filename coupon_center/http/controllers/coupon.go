package controllers

import (
	"math"
	"net/http"
	"platform/commons/codes"

	"platform/commons/grpc_clients/coupon"
	pbcoupon "platform/commons/protos/coupon"

	"github.com/gin-gonic/gin"
)

// Coupon coupon controller
type Coupon struct {
	Base
}

// List list stores
func (p *Coupon) List(c *gin.Context) {
	var errorCode codes.ErrorCode
	p.userInfo, errorCode = p.validUserInfo(c)
	if p.userInfo == nil {
		respformat := p.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	pageNum := p.getPageNum(c)
	pageSize := p.getPageSize(c)

	page := &pbcoupon.Page{
		Num:    int32(pageNum),
		Size:   int32(pageSize),
		UserID: p.userInfo.UserID,
	}

	cc := couponClient.NewCoupon(p.getCouponRPCAddress())
	couponList, err := cc.List(page)
	if err != nil {
		respformat := p.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	totalPage := int(math.Floor(float64(couponList.TotalNum) / float64(pageSize)))
	data := map[string]interface{}{
		"list":      couponList.Coupons,
		"page":      pageNum,
		"pageSize":  pageSize,
		"totalPage": totalPage,
	}

	respformat := p.Response(c, codes.ErrorCodeSuccess, data)
	c.JSON(http.StatusOK, respformat)
	return
}

// Send send to broadcast
func (p *Coupon) Send(c *gin.Context) {
	var errorCode codes.ErrorCode
	p.userInfo, errorCode = p.validUserInfo(c)
	if p.userInfo == nil {
		respformat := p.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	p.roomInfo, errorCode = p.validRoomInfo(c)
	if !p.isValidRoom() {
		respformat := p.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	sendCoupon := &pbcoupon.SendCoupon{
		CouponID:    p.getCouponID(c),
		BroadcastID: p.roomInfo.Broadcast.BroadcastID,
		UserID:      p.userInfo.UserID,
		Number:      int32(p.getNumber(c)),
		Duration:    int64(p.getDuration(c)),
		TypeID:      int32(p.getTypeID(c)),
	}

	cc := couponClient.NewCoupon(p.getCouponRPCAddress())
	reply, err := cc.Send(sendCoupon)
	if err != nil {
		respformat := p.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := p.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Take take a broadacst coupon
func (p *Coupon) Take(c *gin.Context) {
	var errorCode codes.ErrorCode
	p.userInfo, errorCode = p.validUserInfo(c)
	if p.userInfo == nil {
		respformat := p.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	takeCoupon := &pbcoupon.TakeCoupon{
		SendCouponID: p.getSendCouponID(c),
		UserID:       p.userInfo.UserID,
	}

	cc := couponClient.NewCoupon(p.getCouponRPCAddress())
	reply, err := cc.Take(takeCoupon)
	if err != nil {
		respformat := p.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := p.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// Stop stop sendcoupon
func (p *Coupon) Stop(c *gin.Context) {
	var errorCode codes.ErrorCode
	p.userInfo, errorCode = p.validUserInfo(c)
	if p.userInfo == nil {
		respformat := p.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	p.roomInfo, errorCode = p.validRoomInfo(c)
	if !p.isValidRoom() {
		respformat := p.Response(c, errorCode, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	takeCoupon := &pbcoupon.TakeCoupon{
		SendCouponID: p.getSendCouponID(c),
		UserID:       p.userInfo.UserID,
	}

	cc := couponClient.NewCoupon(p.getCouponRPCAddress())
	reply, err := cc.Stop(takeCoupon)
	if err != nil {
		respformat := p.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := p.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// ListByStore list by store
func (p *Coupon) ListByStore(c *gin.Context) {

}

// Show show a store
func (p *Coupon) Show(c *gin.Context) {

}

// Create create a store
func (p *Coupon) Create(c *gin.Context) {

}

// Update update a store
func (p *Coupon) Update(c *gin.Context) {

}

// Delete delete a store
func (p *Coupon) Delete(c *gin.Context) {

}
