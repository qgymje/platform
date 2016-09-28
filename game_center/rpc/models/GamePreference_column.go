package models

type _GamePreferenceColumn struct {
	CreatedAt  string
	GameID     string
	Preference string
	UpdatedAt  string
	UserID     string
}

// GamePreferenceColumns gamepreference columns name
var GamePreferenceColumns _GamePreferenceColumn

func init() {
	GamePreferenceColumns.CreatedAt = "created_at"
	GamePreferenceColumns.GameID = "game_id"
	GamePreferenceColumns.Preference = "preference"
	GamePreferenceColumns.UpdatedAt = "updated_at"
	GamePreferenceColumns.UserID = "user_id"

}
