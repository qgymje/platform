package users

import (
	"errors"
	"net/http"
	"strings"
)

// HEADER_TOKEN_KEY header的name
const HEADER_TOKEN_KEY = "Authorization"

// TokenParser 实现token解析
type TokenParser interface {
	Parse() error
}

var (
	ErrHeaderToken   = errors.New("header token 解析错误")
	ErrTokenNotFound = errors.New("header token 没有设置")
)

type HeaderTokenParser struct {
	req   *http.Request
	token string
}

func NewHeaderTokenParser(req *http.Request) *HeaderTokenParser {
	return &HeaderTokenParser{
		req: req,
	}
}

func (r *HeaderTokenParser) Parse() error {
	token := r.req.Header.Get(HEADER_TOKEN_KEY)
	if token == "" {
		return ErrTokenNotFound
	}

	authHeaderParts := strings.Split(token, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return ErrHeaderToken
	}
	r.token = authHeaderParts[1]
	return nil
}

func (r *HeaderTokenParser) Token() string {
	return r.token
}
