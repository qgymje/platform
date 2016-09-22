package controllers

import (
	"net/http"
	"platform/commons/codes"
	"platform/commons/grpc_clients/email"
	"platform/commons/grpc_clients/user"
	pbemail "platform/commons/protos/email"
	pbuser "platform/commons/protos/user"

	"github.com/gin-gonic/gin"
)

// Email sms controlelr
type Email struct {
	Base
}

// RegisterCode create a register code
func (e *Email) RegisterCode(c *gin.Context) {
	email := &pbuser.Email{
		Email: e.getEmail(c),
	}

	user := userClient.NewUser(e.getUserRPCAddress())
	reply, err := user.EmailCode(email)
	if err != nil {
		respformat := e.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := e.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return

}

// VerifyRegisterCode verify register code
func (e *Email) VerifyRegisterCode(c *gin.Context) {
	client := emailClient.NewEmail(e.getEmailRPCAddress())

	emailCode := &pbemail.EmailCode{
		Email: e.getEmail(c),
		Code:  e.getCode(c),
	}

	reply, err := client.Verify(emailCode)
	if err != nil {
		respformat := e.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := e.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return

}
