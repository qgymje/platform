package games

import "platform/game_center/rpc/models"

type StarterConfig struct {
	GameID string
	UserID string
}

// GameStarter represents a services of game starting
type Starter struct {
	playerGameModel *models.PlayerGame
}
