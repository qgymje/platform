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

var uploadPath string

// SetUploadPath upload path
func SetUploadPath(p string) {
	uploadPath = strings.TrimRight(p, "/") + "/"
}

func getUploadPath() string {
	return uploadPath
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

func (b *Base) getName(c *gin.Context) string {
	return c.PostForm("name")
}

func (b *Base) getMembers(c *gin.Context) []string {
	m := c.Request.PostForm["members"]
	utils.Dump(m)
	return []string{"57e226dac86ab45af3d14807", "57e3a9eec86ab40cf7f5247c"}
}

func (b *Base) getChatID(c *gin.Context) string {
	return c.PostForm("chat_id")
}

func (b *Base) getContent(c *gin.Context) string {
	return c.PostForm("content")
}

func (b *Base) getRoomID(c *gin.Context) string {
	return c.PostForm("room_id")
}

func (b *Base) getChatPCAddress() string {
	return "localhost:4011"
}

func (b *Base) getUserRPCAddress() string {
	return "localhost:4000"
}

func (b *Base) getProfileRPCAddress() string {
	return "localhost:4010"
}
