package controllers

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"platform/commons/codes"
	"platform/utils"

	"platform/commons/grpc_clients/user"
	pbuser "platform/commons/protos/user"

	"github.com/gin-gonic/gin"
)

const (
	headerTokenKey  = "Authorization"
	defaultPageSize = 20
)

// Base controller do common things
type Base struct {
	userInfo *pbuser.UserInfo
}

// ResponseFormat  response format object
type ResponseFormat struct {
	Code codes.ErrorCode        `json:"code"`
	Data interface{}            `json:"data"`
	Msg  string                 `json:"msg"`
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
	return map[string]interface{}{
		"url":       "http://" + c.Request.Host + c.Request.URL.String(),
		"method":    c.Request.Method,
		"timestamp": time.Now(),
	}
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

func (b *Base) getPageNum(c *gin.Context) (page int) {
	page, _ = strconv.Atoi(c.Query("page"))
	return int(math.Max(float64(page-1), 0.0))
}

func (b *Base) getPageSize(c *gin.Context) (num int) {
	num, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		num = defaultPageSize
	}
	return
}

func (b *Base) getBroadcastID(c *gin.Context) string {
	key := "broadcast_id"
	id := c.Query(key)
	if id == "" {
		id = c.PostForm(key)
	}
	return id
}

func (b *Base) getText(c *gin.Context) string {
	return c.PostForm("text")
}

func (b *Base) getTypeID(c *gin.Context) int {
	id, _ := strconv.Atoi(c.PostForm("type_id"))
	return id
}

func (b *Base) getStartTime(c *gin.Context) int64 {
	t, _ := utils.FormatTime(c.Query("start_time"))
	return t.Unix()
}

func (b *Base) getEndTime(c *gin.Context) int64 {
	t, _ := utils.FormatTime(c.Query("end_time"))
	return t.Unix()
}

func (b *Base) getBarrageRPCAddress() string {
	return "localhost:4007"
}

func (b *Base) getUserRPCAddress() string {
	return "localhost:4000"
}
