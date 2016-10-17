package rooms

import (
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
)

// Room  service level room info
type Room struct {
	RoomID    string
	UserName  string
	Name      string
	Cover     string
	IsPlaying bool
	FollowNum int64

	errorCode codes.ErrorCode
}

// NewRoom create a new room
func NewRoom() *Room {
	return new(Room)
}

// ErrorCode implement ErrorCoder
func (r *Room) ErrorCode() codes.ErrorCode {
	return r.errorCode
}

func modelRoomToSrvRoom(r *models.Room) *Room {
	srvRoom := &Room{
		RoomID:    r.GetID(),
		UserName:  r.UserName,
		Name:      r.Name,
		Cover:     r.Cover,
		IsPlaying: r.IsPlaying,
		FollowNum: r.FollowNum,
	}
	return srvRoom
}

// GetByUserID room by  user_id
func (r *Room) GetByUserID(userID string) (*Room, error) {
	mRoom, err := models.FindRoomByUserID(userID)
	if err != nil {
		r.errorCode = codes.ErrorCodeRoomNotFound
		return nil, err
	}
	srvRoom := modelRoomToSrvRoom(mRoom)
	return srvRoom, nil
}
