package games

import "platform/game_center/rpc/models"

// CreaterConfig config of Creater
type CreaterConfig struct {
	CompanyID    string
	Name         string
	GameTypeID   int
	Cover        string
	Screenshoots []string
	Description  string
	PlayerNum    int
	IsFree       bool
	Charge       float64
}

// Creater represent a process of creating a game
type Creater struct {
	gameModel *models.Game
}

// NewCreater create a new Creater object
func NewCreater(config *CreaterConfig) *Creater {
	return new(Creater)
}
