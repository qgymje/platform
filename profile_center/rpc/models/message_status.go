package models

// MessageStatus message status
//go:generate stringer -type MessageStatus
type MessageStatus int

const (
	// Created created
	Created MessageStatus = iota + 1
	// Rollbacked rollbacked
	Rollbacked
	// Committed committed
	Committed
)
