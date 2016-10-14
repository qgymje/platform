package controllers

import (
	"log"
	"strings"
	"time"

	"platform/commons/codes"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

const headerTokenKey = "Authorization"

// Base controller do common things
type Base struct {
	roomRPCAddress string
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

func (b *Base) getRoomRPCAddress() string {
	if b.roomRPCAddress != "" {
		return b.roomRPCAddress
	}

	host := utils.GetConf().GetString("app.rpc_host")
	port := utils.GetConf().GetString("app.rpc_port")
	b.roomRPCAddress = host + port
	return b.roomRPCAddress
}

func (b *Base) getUserRPCAddress() string {
	return "localhost:4000"
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

func (b *Base) getName(c *gin.Context) string {
	return c.PostForm("name")
}

func (b *Base) getChannel(c *gin.Context) string {
	return c.PostForm("channel")
}

func (b *Base) getSubChannel(c *gin.Context) string {
	return c.PostForm("sub_channel")
}

func (b *Base) getCover(c *gin.Context) string {
	return c.PostForm("cover")
}

func (b *Base) getAgreement(c *gin.Context) string {
	return c.PostForm("agreement")
}

func (b *Base) getBarrage(c *gin.Context) string {
	return c.PostForm("barrage")
}
