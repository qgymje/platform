package models

type _BarrageColumn struct {
	BroadcastID string
	CreatedAt   string
	Level       string
	Text        string
	UserID      string
	UserName    string
}

// BarrageColumns barrage columns name
var BarrageColumns _BarrageColumn

func init() {
	BarrageColumns.BroadcastID = "broadcast_id"
	BarrageColumns.CreatedAt = "created_at"
	BarrageColumns.Level = "level"
	BarrageColumns.Text = "text"
	BarrageColumns.UserID = "user_id"
	BarrageColumns.UserName = "username"

}
