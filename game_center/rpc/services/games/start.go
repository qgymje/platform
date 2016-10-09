package games

import "platform/game_center/rpc/models"

// StarterConfig config of starter
type StarterConfig struct {
	GameID string
	UserID string
}

// Starter represents a services of game starting
type Starter struct {
	playerGameModel *models.PlayerGame
}
