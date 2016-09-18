package models

//go:generate stringer -type=GameStatus
type GameStatus int

const (
	Unvalid GameStatus = iota
	ValidFailed
	Published
	Down
	Available
)
