package friends

import (
	"errors"
	"platform/commons/codes"
	"platform/profile_center/rpc/models"
	"platform/utils"
)

// RequestConfig config
type RequestConfig struct {
	FromUserID string
	ToUserID   string
	Message    string
}

// Request request object
type Request struct {
	config    *RequestConfig
	errorCode codes.ErrorCode

	requestFriendModel *models.RequestFriend
}

// NewRequest new request object
func NewRequest(c *RequestConfig) *Request {
	r := new(Request)
	r.config = c
	r.requestFriendModel = &models.RequestFriend{}
	return r
}

// ErrorCode may be use a base service object?
func (r *Request) ErrorCode() codes.ErrorCode {
	return r.errorCode
}

// Do the dirty work?
func (r *Request) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("friends.Request.Do error %+v", err)
		}
	}()

	if yes := r.isRequested(); yes {
		r.errorCode = codes.ErrorCodeRequestFriendAlreadySend
		return errors.New("request friend already send")
	}

	if err = r.save(); err != nil {
		r.errorCode = codes.ErrorCodeRequestFriendSave
		return
	}

	if err = r.autoAgree(); err != nil {
		return
	}

	return
}

// GetRequestID get request id
func (r *Request) GetRequestID() string {
	return r.requestFriendModel.GetID()
}

func (r *Request) autoAgree() error {
	config := &ResponseConfig{RequestFriendID: r.requestFriendModel.GetID()}
	resp := NewResponse(config)
	if err := resp.Agree(); err != nil {
		r.errorCode = resp.ErrorCode()
		return err
	}
	return nil
}

func (r *Request) isRequested() bool {
	err := r.requestFriendModel.Find()
	return err == nil
}

func (r *Request) save() error {
	r.requestFriendModel.FromUserID = r.config.FromUserID
	r.requestFriendModel.ToUserID = r.config.ToUserID
	r.requestFriendModel.Message = r.config.Message
	return r.requestFriendModel.Create()
}
