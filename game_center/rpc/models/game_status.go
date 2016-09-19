package models

// GameStatus game status
//go:generate stringer -type=GameStatus
type GameStatus int

const (
	// Unvalid just apply
	Unvalid GameStatus = iota
	// Uploaded uploaded games
	Uploaded
	// ValidFailed valid failed
	ValidFailed
	// Published Valid success
	Published
	// Down game is not available
	Down
	// Available game is available
	Available
)
