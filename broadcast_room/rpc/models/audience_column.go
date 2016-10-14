package models

type _AudienceColumn struct {
	BroadcastID string
	EnterTime   string
	LeaveTime   string
	UserID      string
}

// AudienceColumns audience columns name
var AudienceColumns _AudienceColumn

func init() {
	AudienceColumns.BroadcastID = "broadcast_id"
	AudienceColumns.EnterTime = "enter_time"
	AudienceColumns.LeaveTime = "leave_time"
	AudienceColumns.UserID = "user_id"

}
