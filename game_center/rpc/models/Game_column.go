package models

type _GameColumn struct {
	Cover       string
	CreatedAt   string
	Description string
	GameID      string
	Name        string
	PublishedAt string
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
	GameColumns.Name = "name"
	GameColumns.PublishedAt = "published_at"
	GameColumns.Status = "status"
	GameColumns.UpdatedAt = "updated_at"

}
