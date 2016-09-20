package sms

import "github.com/astaxie/beego/validation"

// Phone phone object
type Phone struct {
	phone string
	valid *validation.Validation
}

// NewPhone a new phone object
func NewPhone(phone string) *Phone {
	return &Phone{
		phone: phone,
		valid: &validation.Validation{},
	}
}

func (p *Phone) String() string {
	return p.phone
}

// IsValid is phone number valid
func (p *Phone) IsValid() bool {
	if v := p.valid.Mobile(p.phone, "phone"); v.Ok {
		return true
	}
	return false
}
