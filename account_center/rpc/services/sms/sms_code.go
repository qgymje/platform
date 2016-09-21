package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"platform/account_center/rpc/services/notifier"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/utils"
	"time"
)

const (
	codeLength = 6
)

// Code sms code object
type Code struct {
	phone     *Phone
	country   string
	code      string
	errorCode codes.ErrorCode
}

// NewCode returns a Code object
func NewCode(phone, country string) *Code {
	s := new(Code)
	s.phone = NewPhone(phone)
	s.country = country
	return s
}

// ErrorCode implement the ErrorCoder interface
func (s *Code) ErrorCode() codes.ErrorCode {
	return s.errorCode
}

// Do do the notify work
func (s *Code) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("code do error: %v", err)
		}
	}()

	if !s.phone.IsValid() {
		s.errorCode = codes.ErrorCodeInvalidPhone
		return errors.New("invalid phone number")
	}
	if err = s.notify(); err != nil {
		s.errorCode = codes.ErrorCodeLoginNotify
		return
	}
	return
}

func genCode(length int) string {
	n := math.Pow(10, float64(length))
	return fmt.Sprintf("%d", utils.GetRand().Intn(int(n)))
}

// GetCode return the code that generate for the request
func (s *Code) GetCode() string {
	if s.code == "" {
		s.code = genCode(codeLength)
	}
	return s.code
}

// Topic the topic name
func (s *Code) Topic() string {
	return queues.TopicRegisterSMS.String()
}

// Message the message body
func (s *Code) Message() []byte {
	message := queues.MessageRegisterSMS{
		Phone:     s.phone.String(),
		Country:   s.country,
		Code:      s.GetCode(),
		CreatedAt: time.Now(),
	}
	msg, _ := json.Marshal(&message)
	return msg
}

func (s *Code) notify() error {
	return notifier.Publish(s)
}
