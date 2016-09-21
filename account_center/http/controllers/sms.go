package controllers

import (
	"net/http"
	"platform/commons/codes"

	"platform/commons/grpc_clients/sms"
	"platform/commons/grpc_clients/user"
	pbsms "platform/commons/protos/sms"
	pbuser "platform/commons/protos/user"

	"github.com/gin-gonic/gin"
)

// SMS sms controlelr
type SMS struct {
	Base
}

// RegisterCode create a register code
func (s *SMS) RegisterCode(c *gin.Context) {
	user := userClient.NewUser(s.getUserRPCAddress())

	phone := &pbuser.Phone{
		Phone:   s.getPhone(c),
		Country: s.getCountry(c),
	}

	reply, err := user.ValidCode(phone)
	if err != nil {
		respformat := s.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := s.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}

// VerifyRegisterCode verify register code
func (s *SMS) VerifyRegisterCode(c *gin.Context) {
	client := smsClient.NewSMS(s.getSMSRPCAddress())

	phoneCode := &pbsms.PhoneCode{
		Country: s.getCountry(c),
		Phone:   s.getPhone(c),
		Code:    s.getCode(c),
	}

	reply, err := client.Verify(phoneCode)
	if err != nil {
		respformat := s.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	respformat := s.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}
