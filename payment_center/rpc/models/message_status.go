package models

// MessageStatus message status
type MessageStatus int

const (
	// Created created
	Created MessageStatus = iota + 1
	// Confirmed confirmed
	Confirmed
	// Notified notified
	Notified
	// Acked  acknowledged
	Acked
)
