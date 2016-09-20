package providers

// SendCloudEmailSender sendcloud email sender object
type SendCloudEmailSender struct {
}

func NewSendCloudEmailSender() *SendCloudEmailSender {
	return new(SendCloudEmailSender)
}

func (s *SendCloudEmailSender) Send(provider Provider) {

}
