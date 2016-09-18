package controllers

import (
	"log"
	"strings"
	"time"

	"platform/commons/codes"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

const HEADER_TOKEN_KEY = "Authorization"

type Base struct {
	roomRPCAddress string
}

type responseFormat struct {
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

func (b *Base) codeWithMsg(code codes.ErrorCode) *responseFormat {
	msg := codes.GetErrorMsgByCode(code)
	return &responseFormat{
		Code: code,
		Msg:  msg,
		Meta: make(map[string]interface{}),
	}
}

func (b *Base) Response(c *gin.Context, code codes.ErrorCode, data interface{}) *responseFormat {
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
	// TODO: how to get rpc services address???????
	return "localhost:4000"
}

func (b *Base) getToken(c *gin.Context) (string, codes.ErrorCode) {
	if c.Param("token") != "" {
		return c.Param("token"), codes.ErrorCodeSuccess
	}

	token := c.Request.Header.Get(HEADER_TOKEN_KEY)
	if token == "" {
		return "", codes.ErrorCodeTokenNotFound
	}

	authHeaderParts := strings.Split(token, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", codes.ErrorCodeInvalidToken
	}
	return authHeaderParts[1], codes.ErrorCodeSuccess
}
