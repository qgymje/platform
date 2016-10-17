package rooms

import (
	"errors"
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

var (
	// ErrAlreadyCreated room already created error
	ErrAlreadyCreated = errors.New("room already created")
)

// CreatorConfig config of Creator
type CreatorConfig struct {
	UserID   string
	UserName string
	Name     string
	Cover    string // already uploaded?
}

// Creator create a room
type Creator struct {
	config    *CreatorConfig
	roomModel *models.Room
	userID    string

	errorCode codes.ErrorCode
}

// NewCreator create a Creator object
func NewCreator(c *CreatorConfig) *Creator {
	return &Creator{
		config: c,
		roomModel: &models.Room{
			UserName: c.UserName,
			Name:     c.Name,
			Cover:    c.Cover,
		},
		userID: c.UserID,
	}
}

// ErrorCode implement ErrorCoder
func (c *Creator) ErrorCode() codes.ErrorCode {
	return c.errorCode
}

// Do do the dirty work
func (c *Creator) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("rooms.Creator.Do error: %_v", err)
		}
	}()

	foundRoom := false
	if err = c.findRoomByUser(); err != nil {
		c.errorCode = codes.ErrorCodeRoomAlreadyCreated
		foundRoom = true
	}

	if foundRoom {
		if err = c.update(); err != nil {
			c.errorCode = codes.ErrorCodeRoomUpdate
			return
		}
	} else {
		if err = c.save(); err != nil {
			c.errorCode = codes.ErrorCodeRoomCreate
			return
		}
	}

	return
}

// GetRoomID room id
func (c *Creator) GetRoomID() string {
	return c.roomModel.GetID()
}

func (c *Creator) findRoomByUser() (err error) {
	c.roomModel, err = models.FindRoomByUserID(c.userID)
	if err != models.ErrNotFound {
		return ErrAlreadyCreated
	}
	return nil
}

func (c *Creator) save() (err error) {
	c.roomModel.UserID, err = models.StringToObjectID(c.userID)
	if err != nil {
		return err
	}
	if err = c.roomModel.Create(); err != nil {
		return
	}
	return
}

func (c *Creator) update() (err error) {
	c.roomModel.UserID, err = models.StringToObjectID(c.userID)
	if err != nil {
		return err
	}
	if err = c.roomModel.Update(c.config.Name, c.config.Cover); err != nil {
		return
	}
	return

}
