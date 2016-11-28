package friends

import (
	"errors"
	"platform/commons/codes"
	"platform/profile_center/rpc/models"
	"platform/utils"
	"strconv"
)

var (
	// ErrInvalidRequestID invalid request id
	ErrInvalidRequestID = errors.New("request friend invalid id")
)

// ResponseConfig response config
type ResponseConfig struct {
	RequestFriendID string
}

// Response request object
type Response struct {
	config    *ResponseConfig
	errorCode codes.ErrorCode

	requestFriendModel *models.RequestFriend
	friendModel        *models.Friend
}

// NewResponse new request object
func NewResponse(c *ResponseConfig) *Response {
	r := new(Response)
	r.config = c
	r.requestFriendModel = &models.RequestFriend{}
	return r
}

// ErrorCode may be use a base service object?
func (r *Response) ErrorCode() codes.ErrorCode {
	return r.errorCode
}

// Agree the request
func (r *Response) Agree() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("friends.Response.Agree error %+v", err)
		}
	}()

	if yes := r.isRequested(); !yes {
		r.errorCode = codes.ErrorCodeRequestFriendInvalidID
		return ErrInvalidRequestID
	}

	if err = r.saveFriend(); err != nil {
		r.errorCode = codes.ErrorCodeFriendSave
		return
	}

	// may be notified?

	return
}

// Refuse the request
func (r *Response) Refuse() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("friends.Response.Deny error %+v", err)
		}
	}()

	if yes := r.isRequested(); !yes {
		r.errorCode = codes.ErrorCodeRequestFriendInvalidID
		return ErrInvalidRequestID
	}

	if err = r.updateRefuse(); err != nil {
		r.errorCode = codes.ErrorCodeRequestFriendRefuse
		return
	}

	return
}

func (r *Response) updateRefuse() error {
	return r.requestFriendModel.Refuse()
}

func (r *Response) saveFriend() (err error) {
	r.friendModel = &models.Friend{}
	r.friendModel.RequestFriend = r.requestFriendModel
	r.friendModel.FromUserID = r.requestFriendModel.FromUserID
	r.friendModel.ToUserID = r.requestFriendModel.ToUserID
	return r.friendModel.Create()
}

func (r *Response) getRequestFriendID() int64 {
	i, _ := strconv.ParseInt(r.config.RequestFriendID, 10, 0)
	return i
}

func (r *Response) isRequested() bool {
	err := ErrInvalidRequestID
	if r.config.RequestFriendID != "" {
		r.requestFriendModel.ID = r.getRequestFriendID()
		err = r.requestFriendModel.FindByID()
	}

	return err == nil
}
