package models

type BarrageColumn struct {
	RoomID    string
	UserName  string
	CreatedAt string
	Message   string
	UserID    string
}

var BarrageColumns BarrageColumn

func init() {
	BarrageColumns.CreatedAt = "createdAt"
	BarrageColumns.Message = "message"
	BarrageColumns.RoomID = "roomID"
	BarrageColumns.UserID = "userID"
	BarrageColumns.UserName = "userName"
}
