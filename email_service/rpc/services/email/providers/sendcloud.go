package providers

// SendCloudEmailSender sendcloud email sender object
type SendCloudEmailSender struct {
}

// NewSendCloudEmailSender sendcloud
func NewSendCloudEmailSender() *SendCloudEmailSender {
	return new(SendCloudEmailSender)
}

// Send the email
func (s *SendCloudEmailSender) Send(provider Provider) {

}
