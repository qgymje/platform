package models

type _GameColumn struct {
	Cover       string
	CreatedAt   string
	Description string
	GameID      string
	GameTypeID  string
	IsFree      string
	Name        string
	PlayTimes   string
	PlayerNum   string
	PublishedAt string
	Screenshots string
	Status      string
	UpdatedAt   string
}

// GameColumns game columns name
var GameColumns _GameColumn

func init() {
	GameColumns.Cover = "cover"
	GameColumns.CreatedAt = "created_at"
	GameColumns.Description = "description"
	GameColumns.GameID = "_id"
	GameColumns.GameTypeID = "game_type_id"
	GameColumns.IsFree = "is_free"
	GameColumns.Name = "name"
	GameColumns.PlayTimes = "play_times"
	GameColumns.PlayerNum = "player_num"
	GameColumns.PublishedAt = "published_at"
	GameColumns.Screenshots = "screenshots"
	GameColumns.Status = "status"
	GameColumns.UpdatedAt = "updated_at"

}
