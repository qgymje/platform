package games

import (
	"platform/commons/codes"
	"platform/game_center/rpc/models"
	"platform/utils"
	"time"
)

// Config config of game list object
type Config struct {
	GameTypeID int
	Search     string
	PageNum    int
	PageSize   int
	Order      string
}

// Game is a service level Game object, include related informaction like company, player payment
type Game struct {
	GameID       string
	CompanyID    string
	CompanyName  string
	GameTypeID   int
	GameTypeName string
	Name         string
	Cover        string
	Screenshots  []string
	Description  string
	PlayTimes    int64
	PlayerNum    int64
	IsFree       bool
	Charge       float64
	PayStatus    bool
	PublishedAt  time.Time
}

// Games represents a object of  game list
type Games struct {
	gameFinder *models.GameFinder

	gameids []string

	errorCode codes.ErrorCode
}

// ErrorCode return errorCode
func (g *Games) ErrorCode() codes.ErrorCode {
	return g.errorCode
}

// NewGames returns the Games object
func NewGames(config *Config) *Games {
	g := new(Games)
	g.gameids = []string{}
	g.gameFinder = models.NewGameFinder().Limit(config.PageNum, config.PageSize).GameTypeID(config.GameTypeID).Status(models.Published).Search(config.Search)
	return g
}

// Do do the NewGames query
func (g *Games) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("games.Games.Do error: %+v", err)
		}
	}()

	if err = g.find(); err != nil {
		if err == models.ErrNotFound {
			g.errorCode = codes.ErrorCodeGameNotFound
		} else {
			g.errorCode = codes.ErrorCodeGameFinder
		}
		return
	}
	return
}

func (g *Games) find() error {
	return g.gameFinder.Do()
}

// Games fetch the game list object
func (g *Games) Games() []*Game {
	modelGames := g.gameFinder.Result()
	srvGames := []*Game{}
	for _, mGame := range modelGames {
		g.gameids = append(g.gameids, mGame.GetID())
		srvGame := &Game{
			GameID:       mGame.GetID(),
			CompanyID:    mGame.GetCompanyID(),
			CompanyName:  g.getCompanyNameByID(mGame.GetID()),
			GameTypeID:   mGame.GameTypeID,
			GameTypeName: g.getGameTypeNameByID(mGame.GameTypeID),
			Name:         mGame.Name,
			Cover:        mGame.Cover,
			Screenshots:  mGame.Screenshots,
			Description:  mGame.Description,
			PlayerNum:    mGame.PlayerNum,
			PlayTimes:    mGame.PlayTimes,
			IsFree:       mGame.IsFree,
			Charge:       mGame.Charge,
			PayStatus:    false,
			PublishedAt:  mGame.PublishedAt,
		}
		srvGames = append(srvGames, srvGame)
	}
	return srvGames
}

// Count total result count
func (g *Games) Count() int64 {
	return g.gameFinder.Count()
}

func (g *Games) companyInfo() map[string]*models.Company {
	return nil
}

func (g *Games) getCompanyNameByID(id string) string {
	return "Blizzard Entertainment"
}

func (g *Games) getGameTypeNameByID(id int) string {
	return gameTypeByID(id)
}
