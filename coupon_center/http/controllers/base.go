package controllers

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"platform/commons/codes"
	"platform/commons/grpc_clients/user"
	pbuser "platform/commons/protos/user"

	"platform/utils"

	"github.com/gin-gonic/gin"
)

var defaultPageSize = 20

const headerTokenKey = "Authorization"
const versionKey = "version"

// Base controller do common things
type Base struct {
	userInfo *pbuser.UserInfo
}

var uploadPath string

// SetUploadPath upload path
func SetUploadPath(p string) {
	uploadPath = strings.TrimRight(p, "/") + "/"
}

func getUploadPath() string {
	return uploadPath
}

func (b *Base) apiVersion(c *gin.Context) int {
	v, _ := c.Get(versionKey)
	vi, _ := strconv.Atoi(v.(string))
	return vi
}

// ResponseFormat  response format object
type ResponseFormat struct {
	Code codes.ErrorCode        `json:"code"`
	Msg  string                 `json:"msg"`
	Data interface{}            `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

func rpcErrorFormat(code string) codes.ErrorCode {
	log.Println("rpc error: ", code)
	codePart := strings.Split(code, " ")
	return codes.ErrorCode(codePart[len(codePart)-1])
}

func (b *Base) codeWithMsg(code codes.ErrorCode) *ResponseFormat {
	msg := codes.GetErrorMsgByCode(code)
	return &ResponseFormat{
		Code: code,
		Msg:  msg,
		Meta: make(map[string]interface{}),
	}
}

// Response response formatted json
func (b *Base) Response(c *gin.Context, code codes.ErrorCode, data interface{}) *ResponseFormat {
	respformat := b.codeWithMsg(code)
	respformat.Data = data
	if !utils.IsProd() {
		respformat.Meta = b.Meta(c)
	}
	return respformat
}

// Meta 在返回错误时候, 带上额外的信息
func (b *Base) Meta(c *gin.Context) map[string]interface{} {
	meta := map[string]interface{}{
		"url":       "http://" + c.Request.Host + c.Request.URL.String(),
		"method":    c.Request.Method,
		"timestamp": time.Now(),
		"header":    c.Request.Header,
	}

	requestBegin, _ := c.Get("request_begin")
	responseTime := fmt.Sprintf("%.2fms", time.Since(requestBegin.(time.Time)).Seconds()*1000)
	meta["response_time"] = responseTime

	return meta
}

func (b *Base) getToken(c *gin.Context) (string, codes.ErrorCode) {
	if c.Param("token") != "" {
		return c.Param("token"), codes.ErrorCodeSuccess
	}

	token := c.Request.Header.Get(headerTokenKey)
	if token == "" {
		return "", codes.ErrorCodeTokenNotFound
	}

	authHeaderParts := strings.Split(token, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", codes.ErrorCodeInvalidToken
	}
	return authHeaderParts[1], codes.ErrorCodeSuccess
}

func (b *Base) validUserInfo(c *gin.Context) (*pbuser.UserInfo, codes.ErrorCode) {
	token, errorCode := b.getToken(c)
	if errorCode != codes.ErrorCodeSuccess {
		return nil, errorCode
	}
	pbToken := pbuser.Token{Token: token}
	auth := userClient.NewUser(b.getUserRPCAddress())

	var err error
	var userInfo *pbuser.UserInfo
	if userInfo, err = auth.Auth(&pbToken); err != nil {
		return nil, rpcErrorFormat(err.Error())
	}
	return userInfo, codes.ErrorCodeSuccess
}

func (b *Base) getUserID(c *gin.Context) string {
	return c.Param("user_id")
}

func (b *Base) getPageNum(c *gin.Context) (page int) {
	page, _ = strconv.Atoi(c.Param("page"))
	return int(math.Max(float64(page-1), 0.0))
}

func (b *Base) getPageSize(c *gin.Context) (num int) {
	num, err := strconv.Atoi(c.Param("page_num"))
	if err != nil {
		num = defaultPageSize
	}
	return
}

func (b *Base) getCouponID(c *gin.Context) string {
	return c.PostForm("coupon_id")
}

func (b *Base) getSendCouponID(c *gin.Context) string {
	return c.PostForm("sendcoupon_id")
}

func (b *Base) getBroadcastID(c *gin.Context) string {
	return c.PostForm("broadcastID")
}

func (b *Base) getNumber(c *gin.Context) int {
	num, _ := strconv.Atoi(c.PostForm("number"))
	return num
}

func (b *Base) getDuration(c *gin.Context) int {
	dur, _ := strconv.Atoi(c.PostForm("duration"))
	return dur
}

func (b *Base) getTypeID(c *gin.Context) int {
	id, _ := strconv.Atoi(c.PostForm("type_id"))
	return id
}

func (b *Base) getUserRPCAddress() string {
	return "127.0.0.1:4000"
}

func (b *Base) getCouponRPCAddress() string {
	return "127.0.0.1:4004"
}
