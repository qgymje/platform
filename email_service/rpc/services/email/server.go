package email

import "golang.org/x/net/context"

// Server represent a email service implement
type Server struct {
}

// Create create a email record and send by provider
func (s *Server) Create(context.Context, *Phone) (*Code, error) {

}

// Verify a email code
func (s *Server) Verify(context.Context, *PhoneCode) (*Status, error) {

}
