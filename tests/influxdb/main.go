package main

const (
	MyDB     = "square_holes"
	username = "bubba"
	password = "bubba"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "htp://localhost:8086",
		Username: username,
		Password: password,
	})
}
