package models

type _VoteColumn struct {
	Count       string
	CreatedAt   string
	Duration    string
	Name        string
	Number      string
	OptionCount string
	UserID      string
	VoteID      string
}

// VoteColumns vote columns name
var VoteColumns _VoteColumn

func init() {
	VoteColumns.Count = "count"
	VoteColumns.CreatedAt = "created_at"
	VoteColumns.Duration = "duration"
	VoteColumns.Name = "name"
	VoteColumns.Number = "number"
	VoteColumns.OptionCount = "option_count"
	VoteColumns.UserID = "user_id"
	VoteColumns.VoteID = "vote_id"

}
