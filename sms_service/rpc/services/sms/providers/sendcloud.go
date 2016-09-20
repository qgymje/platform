package providers

// SendCloudSMSSender sendcloud sender object
type SendCloudSMSSender struct {
}

// NewSendCloudEmailSender use sendcloud to send sms
func NewSendCloudEmailSender() *SendCloudSMSSender {
	return new(SendCloudSMSSender)
}

// Send the sms
func (s *SendCloudSMSSender) Send(provider Provider) (err error) {
	return
}
