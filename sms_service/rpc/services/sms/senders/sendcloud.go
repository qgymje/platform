package senders

// SendCloudSMSSender sendcloud sender object
type SendCloudSMSSender struct {
	provider Provider
}

// NewSendCloudEmailSender use sendcloud to send sms
func NewSendCloudEmailSender(provider Provider) *SendCloudSMSSender {
	s := new(SendCloudSMSSender)
	s.provider = provider
	return s
}

// Send the sms
func (s *SendCloudSMSSender) Send() (err error) {
	return
}
