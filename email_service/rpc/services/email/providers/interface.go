package email

// Provider  is an service interface
type Provider interface {
	Address() string
	Content() string
	Title() string
}
