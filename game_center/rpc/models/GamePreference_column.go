package models

type _GamePreferenceColumn struct {
	CreatedAt  string
	GameID     string
	Preference string
	UpdatedAt  string
	UserID     string
}

var GamePreferenceColumns _GamePreferenceColumn

func init() {
	GamePreferenceColumns.CreatedAt = "created_at"
	GamePreferenceColumns.GameID = "gameID"
	GamePreferenceColumns.Preference = "preference"
	GamePreferenceColumns.UpdatedAt = "updated_at"
	GamePreferenceColumns.UserID = "userID"

}
