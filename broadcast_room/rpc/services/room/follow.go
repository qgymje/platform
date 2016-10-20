package rooms

import (
	"errors"
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

// FollowConfig follow config
type FollowConfig struct {
	UserID string
	RoomID string
}

// Follow service level follow
type Follow struct {
	config      *FollowConfig
	modelFollow *models.RoomFollow

	errorCode codes.ErrorCode
}

// ErrorCode implement ErrorCoder
func (f *Follow) ErrorCode() codes.ErrorCode {
	return f.errorCode
}

// NewFollow create a new service level follow object
func NewFollow(c *FollowConfig) *Follow {
	f := new(Follow)
	f.config = c
	userID, _ := models.StringToObjectID(c.UserID)
	roomID, _ := models.StringToObjectID(c.RoomID)
	f.modelFollow = &models.RoomFollow{
		UserID: userID,
		RoomID: roomID,
	}
	return f
}

func (f *Follow) isRoomExists() bool {
	mRoom, err := models.FindRoomByID(f.config.RoomID)
	if mRoom != nil && err == nil {
		return true
	}
	return false
}

// Do follow a room
func (f *Follow) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("rooms.Follow.Do error: %_v", err)
		}
	}()

	if yes := f.isRoomExists(); !yes {
		f.errorCode = codes.ErrorCodeRoomNotFound
		return errors.New("room not exists")
	}

	if err = f.modelFollow.Follow(); err != nil {
		f.errorCode = codes.ErrorCodeFollow
		return
	}
	return
}

// Undo unfollow a room
func (f *Follow) Undo() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("rooms.Unfollow.Do error: %_v", err)
		}
	}()

	if yes := f.isRoomExists(); !yes {
		f.errorCode = codes.ErrorCodeRoomNotFound
		return errors.New("room not exists")
	}

	if err = f.modelFollow.Unfollow(); err != nil {
		f.errorCode = codes.ErrorCodeUnfollow
		return
	}
	return
}
