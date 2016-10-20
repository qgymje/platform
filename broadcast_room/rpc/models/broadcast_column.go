package models

type _BroadcastColumn struct {
	BroadcastID     string
	CurrentAudience string
	EndTime         string
	RoomID          string
	StartTime       string
	TotalAudience   string
}

// BroadcastColumns broadcast columns name
var BroadcastColumns _BroadcastColumn

func init() {
	BroadcastColumns.BroadcastID = "_id"
	BroadcastColumns.CurrentAudience = "current_audience"
	BroadcastColumns.EndTime = "end_time"
	BroadcastColumns.RoomID = "room_id"
	BroadcastColumns.StartTime = "start_time"
	BroadcastColumns.TotalAudience = "total_audience"

}
