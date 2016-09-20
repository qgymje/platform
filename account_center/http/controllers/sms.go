package controllers

import (
	"net/http"
	"platform/commons/codes"

	"platform/commons/grpc_clients/user"
	pb "platform/commons/protos/user"

	"github.com/gin-gonic/gin"
)

// SMS sms controlelr
type SMS struct {
	Base
}

// RegisterCode create a register code
func (s *SMS) RegisterCode(c *gin.Context) {
	user := userClient.NewUser(s.getUserRPCAddress())
	var reply *pb.Code
	var err error
	if reply, err = user.ValidCode(&pb.Phone{Phone: s.GetPhone(c)}); err != nil {
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

}
