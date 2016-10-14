package models

type _BarrageColumn struct {
	CreatedAt string
	Message   string
	RoomID    string
	UserID    string
	UserName  string
}

// BarrageColumns barrage columns name
var BarrageColumns _BarrageColumn

func init() {
	BarrageColumns.CreatedAt = "createdAt"
	BarrageColumns.Message = "message"
	BarrageColumns.RoomID = "roomID"
	BarrageColumns.UserID = "userID"
	BarrageColumns.UserName = "userName"

}
