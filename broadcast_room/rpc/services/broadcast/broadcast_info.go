package broadcasts

import (
	"platform/broadcast_room/rpc/models"
	"time"
)

// Broadcast service level broadcast
type Broadcast struct {
	BroadcastID     string
	RoomID          string
	TotalAudience   int64
	CurrentAudience int64
	StartTime       time.Time
	Duration        int64
}

func modelBroadcastToSrvBroadcast(b *models.Broadcast) *Broadcast {
	return &Broadcast{
		BroadcastID:     b.GetID(),
		RoomID:          b.GetRoomID(),
		TotalAudience:   b.TotalAudience,
		CurrentAudience: b.CurrentAudience,
		StartTime:       b.StartTime,
		Duration:        b.Duration(),
	}
}
