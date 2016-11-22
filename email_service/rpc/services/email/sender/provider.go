package senders

// Provider provide data that sender need
type Provider interface {
	Address() string
	Content() string
	Title() string
}
