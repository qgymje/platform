package controllers

import (
	"log"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"

	"platform/commons/codes"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

const headerTokenKey = "Authorization"
const defaultPageSize = 20

// Base controller do common things
type Base struct {
	gameRPCAddress string
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
		"header":    c.Request.Header,
	}
}

func (b *Base) getGameRPCAddress() string {
	return "localhost:4002"
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

func (b *Base) getSearch(c *gin.Context) string {
	q, _ := url.QueryUnescape(c.Param("search"))
	return q
}

func (b *Base) getGameType(c *gin.Context) int {
	param := "game_type"
	id := c.Query(param)
	if id == "" {
		id = c.PostForm(param)
	}
	intid, _ := strconv.Atoi(id)
	return intid
}

func (b *Base) getGameID(c *gin.Context) string {
	return c.Param("game_id")
}

func (b *Base) getCompayID(c *gin.Context) string {
	return c.PostForm("company_id")
}

func (b *Base) getName(c *gin.Context) string {
	return c.PostForm("name")
}

func (b *Base) getCover(c *gin.Context) string {
	return c.PostForm("cover")
}

func (b *Base) getScreenshots(c *gin.Context) []string {
	s := c.Request.MultipartForm.Value["screenshots[]"]
	return s
}

func (b *Base) getDescription(c *gin.Context) string {
	return c.PostForm("description")
}

func (b *Base) getPlayerNum(c *gin.Context) int {
	p := c.PostForm("player_num")
	intp, _ := strconv.Atoi(p)
	return intp
}

func (b *Base) getIsFree(c *gin.Context) bool {
	f := c.PostForm("is_free")
	boolf, _ := strconv.ParseBool(f)
	return boolf
}

func (b *Base) getCharge(c *gin.Context) float64 {
	charge := c.PostForm("charge")
	floatCharge, _ := strconv.ParseFloat(charge, 64)
	return floatCharge
}
