package barrages

import (
	"platform/barrage_center/rpc/models"
	"time"
)

// Barrage service level barrage object
type Barrage struct {
	TypeID      int
	BroadcastID string
	UserID      string
	Text        string
	Username    string
	Level       int64
	CreatedAt   time.Time
}

func modelBarrageTosrvBarrage(b *models.Barrage) *Barrage {
	return &Barrage{
		BroadcastID: b.GetBroadcastID(),
		UserID:      b.GetUserID(),
		Text:        b.Text,
		Username:    b.Username,
		Level:       b.Level,
		CreatedAt:   b.CreatedAt,
	}
}
