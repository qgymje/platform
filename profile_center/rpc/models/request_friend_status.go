package models

// RequestFriendStatus status
//go:generate stringer -type RequestFriendStatus
type RequestFriendStatus int

const (
	// Unresponsed unresponsed
	Unresponsed RequestFriendStatus = iota
	// Agreed aggreed
	Agreed
	// Refused resfused
	Refused
)
