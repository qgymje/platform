package email

import "github.com/astaxie/beego/validation"

// Email email object
type Email struct {
	email string
	valid *validation.Validation
}

// NewEmail a new email object
func NewEmail(email string) *Email {
	return &Email{
		email: email,
		valid: &validation.Validation{},
	}
}

func (e *Email) String() string {
	return e.email
}

// IsValid is phone number valid
func (e *Email) IsValid() bool {
	if v := e.valid.Email(e.email, "email"); v.Ok {
		return true
	}
	return false
}
