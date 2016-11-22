package senders

// SendCloudEmailSender sendcloud email sender object
type SendCloudEmailSender struct {
	provider Provider
}

// NewSendCloudEmailSender sendcloud
func NewSendCloudEmailSender(provider Provider) *SendCloudEmailSender {
	s := new(SendCloudEmailSender)
	s.provider = provider
	return s
}

// Send the email
func (s *SendCloudEmailSender) Send() {

}
