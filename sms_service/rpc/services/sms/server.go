package sms

import "golang.org/x/net/context"

// Server represent a sms service implement
type Server struct {
}

// Verify a sms code
func (s *Server) Verify(context.Context, *PhoneCode) (*Status, error) {

}
