package models

//go:generate stringer -type=GameStatus
type GameStatus int

const (
	Unvalid GameStatus = iota
	Uploaded
	ValidFailed
	Published
	Down
	Available
)
