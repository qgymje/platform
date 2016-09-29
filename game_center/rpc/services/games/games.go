package games

import (
	"math"
	"platform/game_center/rpc/models"
	"time"
)

// GamesConfig config of game list object
type GamesConfig struct {
	GameTypeID int
	Query      string
	PageNum    int
	PageSize   int
	Order      string
}

// Game is a service level Game object, include related informaction like company, player payment
type Game struct {
	GameID      string    `json:"_id"`
	GameTypeID  int       `json:"game_type_id"`
	Name        string    `json:"name"`
	Cover       string    `json:"cover"`
	Screenshots []string  `json:"screenshots"`
	Description string    `json:"description"`
	PlayTimes   int       `json:"play_times"`
	PlayerNum   int       `json:"player_num"`
	IsFree      bool      `json:"is_free"`
	CompanyName string    `json:"company_name"`
	PublishedAt time.Time `json":"published_at"`
}

// Games represents a object of  game list
type Games struct {
	modelGames []*models.Game
	gameTypeID int
	query      string
	offset     int
	limit      int
	order      string
}

// NewGames returns the Games object
func NewGames(config GamesConfig) *Games {
	g := new(Games)
	g.modelGames = []*models.Game{}
	g.gameTypeID = config.GameTypeID
	g.offset = int(math.Max(0, float64(config.PageNum-1)))
	g.limit = config.PageSize
	return g
}

// Do do the NewGames query
func (g *Games) Do() (err error) {
	return
}

// Games fetch the game list object
func (g *Games) Games() []*Game {
	return nil
}

func (g *Games) companyInfo() map[string]*models.Company {
	return nil
}
