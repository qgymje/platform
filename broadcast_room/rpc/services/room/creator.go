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
	UserID string
	Name   string
	Cover  string // already uploaded?
}

// Creator create a room
type Creator struct {
	room   *models.Room
	userID string

	errorCode codes.ErrorCode
}

// NewCreator create a Creator object
func NewCreator(config *CreatorConfig) *Creator {
	return &Creator{
		room: &models.Room{
			Name:  config.Name,
			Cover: config.Cover,
		},
		userID: config.UserID,
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

	if err = c.findRoomByUser(); err != nil {
		c.errorCode = codes.ErrorCodeRoomAlreadyCreated
		return
	}

	if err = c.save(); err != nil {
		c.errorCode = codes.ErrorCodeRoomCreate
		return
	}
	return
}

// GetRoomID room id
func (c *Creator) GetRoomID() string {
	return c.room.GetID()
}

func (c *Creator) findRoomByUser() (err error) {
	_, err = models.FindRoomByUserID(c.userID)
	if err != models.ErrNotFound {
		return ErrAlreadyCreated
	}
	return nil
}

func (c *Creator) save() (err error) {
	c.room.UserID, err = models.StringToObjectID(c.userID)
	if err != nil {
		return err
	}
	if err = c.room.Create(); err != nil {
		return
	}
	return
}
