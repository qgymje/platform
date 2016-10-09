package games

import (
	"platform/commons/codes"
	"platform/game_center/rpc/models"
	"platform/utils"

	"gopkg.in/mgo.v2/bson"
)

// CreaterConfig config of Creater
type CreaterConfig struct {
	CompanyID   string
	Name        string
	GameTypeID  int
	Cover       string
	Screenshots []string
	Description string
	PlayerNum   int64
	IsFree      bool
	Charge      float64
}

// Creater represent a process of creating a game
type Creater struct {
	gameModel *models.Game

	errorCode codes.ErrorCode
}

// NewCreater create a new Creater object
func NewCreater(config *CreaterConfig) *Creater {
	c := new(Creater)
	c.gameModel = &models.Game{
		CompanyID:   bson.ObjectIdHex(config.CompanyID),
		GameTypeID:  config.GameTypeID,
		Name:        config.Name,
		Cover:       config.Cover,
		Screenshots: config.Screenshots,
		Description: config.Description,
		PlayerNum:   config.PlayerNum,
		IsFree:      config.IsFree,
		Charge:      config.Charge,
	}
	return c
}

// ErrorCode error code
func (c *Creater) ErrorCode() codes.ErrorCode {
	return c.errorCode
}

// Do the creation job
func (c *Creater) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("games.Creater.Do error: %+v", err)
		}
	}()

	if err = c.save(); err != nil {
		c.errorCode = codes.ErrorCodeGameCreate
		return
	}
	// pretend to pass the validation
	c.gameModel.Valid()
	return
}

func (c *Creater) save() (err error) {
	return c.gameModel.Create()
}
