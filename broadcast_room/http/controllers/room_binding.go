package controllers

import (
	"platform/commons/codes"
	pb "platform/commons/protos/room"

	"github.com/gin-gonic/gin"
)

type RoomBinding struct {
	Name       string `form:"name" binding:"required"`
	Channel    string `form:"channel" binding:"required"`
	SubChannel string `form:"subChannel" binding:"required"`
	Cover      string `form:"cover" binding:"required"`
	Agreement  bool   `form:"agreement" binding:"required"`

	errorCode codes.ErrorCode
	config    *pb.RoomRequest
}

func NewRoomBinding(c *gin.Context) (form *RoomBinding, err error) {
	form = &RoomBinding{
		config: &pb.RoomRequest{},
	}
	if err = c.Bind(form); err != nil {
		form.errorCode = codes.ErrorCodeMissParameters
		return
	}

	if !form.Agreement {
		form.errorCode = codes.ErrorCodeAgreement
		return
	}
	return
}

func (b *RoomBinding) ErrorCode() codes.ErrorCode {
	return b.errorCode
}

func (b *RoomBinding) Config() *pb.RoomRequest {
	b.config.Name = b.Name
	b.config.Channel = b.Channel
	b.config.SubChannel = b.SubChannel
	b.config.Cover = b.Cover
	return b.config
}
