package controllers

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"platform/commons/codes"
	pb "platform/commons/protos/user"

	"platform/utils"

	"github.com/gin-gonic/gin"
)

var defaultPageSize = 20

const HEADER_TOKEN_KEY = "Authorization"
const VERSION = "version"

type Base struct {
}

func (b *Base) apiVersion(c *gin.Context) int {
	v, _ := c.Get(VERSION)
	vi, _ := strconv.Atoi(v.(string))
	return vi
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

type responseFormat struct {
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
	meta := map[string]interface{}{
		"url":       "http://" + c.Request.Host + c.Request.URL.String(),
		"method":    c.Request.Method,
		"timestamp": time.Now(),
		"header":    c.Request.Header,
	}

	request_begin, _ := c.Get("request_begin")
	response_time := fmt.Sprintf("%.2fms", time.Since(request_begin.(time.Time)).Seconds()*1000)
	meta["response_time"] = response_time

	return meta
}

func (b *Base) GetUserID(c *gin.Context) string {
	return c.Param("user_id")
}

func (b *Base) GetPageNum(c *gin.Context) (page int) {
	page, _ = strconv.Atoi(c.Param("page"))
	return int(math.Max(float64(page-1), 0.0))
}

func (b *Base) GetPageSize(c *gin.Context) (num int) {
	num, err := strconv.Atoi(c.Param("page_num"))
	if err != nil {
		num = defaultPageSize
	}
	return
}

func (b *Base) RemovePBUserInfoToken(u *pb.UserInfo) {
	u.Token = ""
}
