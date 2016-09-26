package models

type _EmailColumn struct {
	Email     string
	Content   string
	CreatedAt string
	Provider  string
	Type      string
	UsedAt    string
}

// EmailColumns email columns name
var EmailColumns _EmailColumn

func init() {
	EmailColumns.Email = "email"
	EmailColumns.Content = "content"
	EmailColumns.CreatedAt = "created_at"
	EmailColumns.Provider = "provider"
	EmailColumns.Type = "type"
	EmailColumns.UsedAt = "used_at"

}
