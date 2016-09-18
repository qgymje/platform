package rooms

import (
	"errors"
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

var (
	ErrAlreadyCreated = errors.New("broadcast room already created")
)

type RoomCreator struct {
	room   *models.BroadcastRoom
	userID string

	errorCode codes.ErrorCode
}

type RoomCreatorConfig struct {
	UserID     string
	Name       string
	Channel    string
	SubChannel string
	Cover      string // already uploaded?
}

func NewRoomCreator(config *RoomCreatorConfig) *RoomCreator {
	return &RoomCreator{
		room: &models.BroadcastRoom{
			Name:       config.Name,
			Channel:    config.Channel,
			SubChannel: config.SubChannel,
			Cover:      config.Cover,
		},
		userID: config.UserID,
	}
}

func (c *RoomCreator) ErrorCode() codes.ErrorCode {
	return c.errorCode
}

func (c *RoomCreator) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("rooms.RoomCreator.Do error: %v", err)
		}
	}()

	if err = c.findRoomByUser(); err != nil {
		c.errorCode = codes.ErrorCodeRoomAlreadyCreated
		return
	}

	if err = c.create(); err != nil {
		c.errorCode = codes.ErrorCodeRoomCreate
		return
	}

	return
}

func (c *RoomCreator) GetRoomID() string {
	return c.room.GetID()
}

func (c *RoomCreator) GetName() string {
	return c.room.Name
}

func (c *RoomCreator) GetChannel() string {
	return c.room.Channel
}

func (c *RoomCreator) GetSubChannel() string {
	return c.room.SubChannel
}

func (c *RoomCreator) GetCover() string {
	return c.room.Cover
}

func (c *RoomCreator) findRoomByUser() (err error) {
	_, err = models.FindBroadcastRoomByUserID(c.userID)
	if err != models.ErrNotFound {
		return ErrAlreadyCreated
	}
	return nil
}

func (c *RoomCreator) create() (err error) {
	c.room.UserID, err = models.StringToObjectID(c.userID)
	if err != nil {
		return err
	}
	if err = c.room.Create(); err != nil {
		return
	}
	return
}
