package sms

import (
	"platform/commons/codes"
	"platform/sms_service/rpc/models"
	"platform/sms_service/rpc/services/sms/providers"
	"platform/utils"
	"time"
)

// RegisterCode used when a new user before registering
type RegisterCode struct {
	smsModel *models.SMS

	errorCode codes.ErrorCode
}

// RegisterCodeConfig register code config
type RegisterCodeConfig struct {
	Phone     string
	Country   string
	Code      string
	CreatedAt time.Time
}

// NewRegisterCode create and verify sms register code
func NewRegisterCode(config *RegisterCodeConfig) *RegisterCode {
	c := new(RegisterCode)
	c.smsModel = &models.SMS{
		Phone:     config.Phone,
		Country:   config.Country,
		Content:   config.Code,
		Type:      int(models.RegisterCode),
		CreatedAt: config.CreatedAt,
	}
	return c
}

// ErrorCode implement ErrorCoder interface
func (r *RegisterCode) ErrorCode() codes.ErrorCode {
	return r.errorCode
}

// Create do the main bussiness logic
func (r *RegisterCode) Create() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("sms rpc Create error: %v", err)
		}
	}()

	if err = r.save(); err != nil {
		r.errorCode = codes.ErrorCodeSMSCreate
		return
	}

	return
}

func (r *RegisterCode) save() (err error) {
	return r.smsModel.Create()
}

func (r *RegisterCode) determinProvider() models.SMSProvider {
	return models.SendCloud
}

// Verify verify code is correct
func (r *RegisterCode) Verify(country, phone, code string) (err error) {
	//
	return
}

func (r *RegisterCode) send() error {
	sender := providers.NewSendCloudEmailSender()
	return sender.Send(r)
}

// Phone implement the Provider interface
func (r *RegisterCode) Phone() string {
	return r.smsModel.Phone
}

// Content implement the Provider interface
func (r *RegisterCode) Content() string {
	return r.smsModel.Content
}
