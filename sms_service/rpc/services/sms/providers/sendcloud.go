package providers

type SendCloudSMSSender struct {
}

func NewSendCloudEmailSender() SendCloudSMSSender {
	return new(SendCloudSMSSender)
}
