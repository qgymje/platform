package email

import (
	"platform/commons/codes"
	"platform/email_service/rpc/models"
	"platform/utils"
	"time"
)

// RegisterCode used when a new user before registering
type RegisterCode struct {
	emailModel *models.Email

	errorCode codes.ErrorCode
}

// RegisterCodeConfig email register config
type RegisterCodeConfig struct {
	Email     string
	Code      string
	CreatedAt time.Time
}

// NewRegisterCode create and verify sms register code
func NewRegisterCode(config *RegisterCodeConfig) *RegisterCode {
	c := new(RegisterCode)
	c.emailModel = &models.Email{
		Email:     config.Email,
		Content:   config.Code,
		Type:      int(models.RegisterCode),
		CreatedAt: config.CreatedAt,
	}
	return c
}

// ErrorCode implement the ErrorCoder interface
func (r *RegisterCode) ErrorCode() codes.ErrorCode {
	return r.errorCode
}

// Create do the main bussiness logic
func (r *RegisterCode) Create() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("email rpc Create error: %v", err)
		}
	}()

	if err = r.save(); err != nil {
		r.errorCode = codes.ErrorCodeEmailCreate
		return
	}
	return
}

// Verify verify code is correct
func (r *RegisterCode) Verify() (err error) {
	return
}

func (r *RegisterCode) save() error {
	return r.emailModel.Create()
}
