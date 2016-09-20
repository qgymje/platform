package sms

type Provider interface {
	Phone() string
	Content() string
}
