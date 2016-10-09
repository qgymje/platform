package games

import "platform/game_center/rpc/models"

// EnderConfig config of Ender
type EnderConfig struct {
	GameID string
	UserID string
}

// Ender represents a services of game end
type Ender struct {
	playerGameModel *models.PlayerGame
}
