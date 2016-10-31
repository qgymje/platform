package main

// Server test server
type Server struct {
	Host   string
	Path   string
	Params url.Value
}

var devServer = Server{"http://localhost:8080/"}

// User login user
type User struct {
	Phone    string
	Password string
}

// Login login
func (u *User) Login() (token string, err error) {

}

func main() {

}
