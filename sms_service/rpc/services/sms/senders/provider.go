package senders

// Provider objecter provider
// service object implement the interface to support data that sender need
type Provider interface {
	Phone() string
	Content() string
}
