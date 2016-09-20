package providers

type Provider interface {
	Phone() string
	Content() string
}
