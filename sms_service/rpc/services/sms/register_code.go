package sms

// RegisterCode used when a new user before registering
type RegisterCode struct {
	phone *Phone
}

// NewRegisterCode create and verify sms register code
func NewRegisterCode(phone string) *RegisterCode {
	c := new(RegisterCode)
	c.Phone = NewPhone(phone)
	return c
}

// Create do the main bussiness logic
func (r *RegisterCode) Create() (err error) {
	return
}

// Verify verify code is correct
func (r *RegisterCode) Verify() (err error) {
	return
}

func (r *RegisterCode) send() error {
	sender := provider.NewSendCloudSMSSender()
	return sender.Send(r)
}

func (r *RegisterCode) Phone() string {
	return r.phone.String()
}

func (r *RegisterCode) Content() string {
	return "4000"
}
