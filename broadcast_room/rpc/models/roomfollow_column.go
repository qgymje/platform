package models

type _RoomFollowColumn struct {
	CreatedAt string
	DeletedAt string
	RoomID    string
	UserID    string
}

// RoomFollowColumns roomfollow columns name
var RoomFollowColumns _RoomFollowColumn

func init() {
	RoomFollowColumns.CreatedAt = "created_at"
	RoomFollowColumns.DeletedAt = "deleted_at"
	RoomFollowColumns.RoomID = "room_id"
	RoomFollowColumns.UserID = "user_id"

}
