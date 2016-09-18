package models

type BroadcastRoomColumn struct {
	Channel     string
	SubChannel  string
	Cover       string
	CreatedAt   string
	ID          string
	IsPlaying   string
	Name        string
	Orientation string
	Score       string
	UpdatedAt   string
	UserID      string
}

var BroadcastRoomColumns BroadcastRoomColumn

func init() {
	BroadcastRoomColumns.Channel = "channel"
	BroadcastRoomColumns.SubChannel = "subChannel"
	BroadcastRoomColumns.Cover = "cover"
	BroadcastRoomColumns.CreatedAt = "createdAt"
	BroadcastRoomColumns.ID = "_id"
	BroadcastRoomColumns.IsPlaying = "isPlaying"
	BroadcastRoomColumns.Name = "name"
	BroadcastRoomColumns.Orientation = "orientation"
	BroadcastRoomColumns.Score = "score"
	BroadcastRoomColumns.UpdatedAt = "updatedAt"
	BroadcastRoomColumns.UserID = "userID"
}
