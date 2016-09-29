package models

type _PlayerGameColumn struct {
	EndTime   string
	GameID    string
	StartTime string
	UserID    string
}

// PlayerGameColumns playergame columns name
var PlayerGameColumns _PlayerGameColumn

func init() {
	PlayerGameColumns.EndTime = "end_time"
	PlayerGameColumns.GameID = "game_id"
	PlayerGameColumns.StartTime = "start_time"
	PlayerGameColumns.UserID = "user_id"

}
