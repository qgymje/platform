package games

import "platform/game_center/rpc/models"

type EnderConfig struct {
	GameID string
	UserID string
}

// GameEnder represents a services of game end
type Ender struct {
	playerGameModel *models.PlayerGame
}
