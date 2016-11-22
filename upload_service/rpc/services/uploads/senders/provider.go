package senders

// Provider provider interface
type Provider interface {
	Filename() string
	Content() []byte
}
