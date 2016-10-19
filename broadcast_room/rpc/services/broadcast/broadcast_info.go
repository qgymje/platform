package broadcasts

import "time"

// Broadcast service level broadcast
type Broadcast struct {
	BroadcastID   string
	RoomID        string
	TotalAudience int64
	StartTime     time.Time
}
