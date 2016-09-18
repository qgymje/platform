package users

import (
	"errors"
	"fmt"
	"time"

	"platform/commons/codes"
	"platform/utils"

	"github.com/dgrijalva/jwt-go"
)

var ErrGenerate = errors.New("生成token错误")
var ErrVerify = errors.New("验证token失败")

type Token struct {
	token       *jwt.Token
	secretKey   []byte
	claims      map[string]interface{}
	tokenString string

	errorCode codes.ErrorCode
}

var defaultSecretKey string
var defaultExpires int64

func NewToken() *Token {
	defaultSecretKey = utils.GetConf().GetString("auth.secret_key")
	expires := utils.GetConf().GetInt("auth.expires") //second
	defaultExpires = time.Now().Add(time.Second * time.Duration(expires)).Unix()

	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtToken.Claims.(jwt.MapClaims)["exp"] = defaultExpires

	return &Token{
		token:     jwtToken,
		secretKey: []byte(defaultSecretKey),
		claims:    make(map[string]interface{}),
	}
}

func (t *Token) Generate(claims map[string]interface{}) (tokenString string, err error) {
	for k, v := range claims {
		t.token.Claims.(jwt.MapClaims)[k] = v
	}

	tokenString, err = t.token.SignedString(t.secretKey)
	if err != nil {
		err = ErrGenerate
		t.errorCode = codes.ErrorCodeGenerateToekn
	}

	return
}

func (t *Token) Verify(tokenString string) (valid bool, err error) {
	//自带过期处理
	t.token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.errorCode = codes.ErrorCodeInvalidToken
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return t.secretKey, nil
	})
	if err != nil {
		err = ErrVerify
	}
	valid = t.token.Valid
	t.tokenString = tokenString

	return
}

// Claims 获取jwt 里的数据
func (t *Token) Claims() map[string]interface{} {
	return t.token.Claims.(jwt.MapClaims)
}

func (t *Token) String() string {
	return t.tokenString
}

func (t *Token) ErrorCode() codes.ErrorCode {
	return t.errorCode
}

func (t *Token) GetUserInfo() (*UserInfo, error) {
	ui := NewUserInfo()
	err := ui.GetByToken(t.tokenString)
	return ui, err
}
