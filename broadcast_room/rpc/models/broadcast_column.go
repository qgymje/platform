package models

type _BroadcastColumn struct {
	BroadCastID   string
	EndTime       string
	RoomID        string
	StartTime     string
	TotalAudience string
}

// BroadcastColumns broadcast columns name
var BroadcastColumns _BroadcastColumn

func init() {
	BroadcastColumns.BroadCastID = "_id"
	BroadcastColumns.EndTime = "end_time"
	BroadcastColumns.RoomID = "room_id"
	BroadcastColumns.StartTime = "start_time"
	BroadcastColumns.TotalAudience = "total_audience"

}
