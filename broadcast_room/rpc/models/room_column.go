package models

type _RoomColumn struct {
	Cover     string
	CreatedAt string
	DeletedAt string
	FollowNum string
	IsPlaying string
	Name      string
	RoomID    string
	UpdatedAt string
	UserID    string
	UserName  string
}

// RoomColumns room columns name
var RoomColumns _RoomColumn

func init() {
	RoomColumns.Cover = "cover"
	RoomColumns.CreatedAt = "created_at"
	RoomColumns.DeletedAt = "deleted_at"
	RoomColumns.FollowNum = "follow_num"
	RoomColumns.IsPlaying = "is_playing"
	RoomColumns.Name = "name"
	RoomColumns.RoomID = "_id"
	RoomColumns.UpdatedAt = "updated_at"
	RoomColumns.UserID = "user_id"
	RoomColumns.UserName = "user_name"

}
