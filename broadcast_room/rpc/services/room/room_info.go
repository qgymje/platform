package rooms

import (
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"time"
)

// Room  service level room info
type Room struct {
	RoomID    string
	UserName  string
	Name      string
	Cover     string
	IsPlaying bool
	IsFollow  bool
	FollowNum int64
	Broadcast *Broadcast

	errorCode codes.ErrorCode
}

// Broadcast service level broadcast
type Broadcast struct {
	BroadcastID     string
	RoomID          string
	TotalAudience   int64
	CurrentAudience int64
	StartTime       time.Time
	Duration        int64
}

// NewRoom create a new room
func NewRoom() *Room {
	return new(Room)
}

// ErrorCode implement ErrorCoder
func (r *Room) ErrorCode() codes.ErrorCode {
	return r.errorCode
}

func modelRoomToSrvRoom(r *models.Room, b *models.Broadcast) *Room {
	srvRoom := &Room{
		RoomID:    r.GetID(),
		UserName:  r.UserName,
		Name:      r.Name,
		Cover:     r.Cover,
		IsPlaying: r.IsPlaying,
		FollowNum: r.FollowNum,
	}
	if b != nil {
		srvBro := &Broadcast{
			BroadcastID:     b.GetID(),
			RoomID:          r.GetID(),
			TotalAudience:   b.TotalAudience,
			CurrentAudience: b.CurrentAudience,
			StartTime:       b.StartTime,
			Duration:        b.Duration(),
		}
		srvRoom.Broadcast = srvBro
	}
	return srvRoom
}

// GetByID get by room id
func (r *Room) GetByID(roomID string) (*Room, error) {
	var err error
	mRoom, err := models.FindRoomByID(roomID)
	if err != nil {
		r.errorCode = codes.ErrorCodeRoomNotFound
		return nil, err
	}

	mBroadcast, err := r.getBroadcast(mRoom)
	if err != nil {
		return nil, err
	}

	srvRoom := modelRoomToSrvRoom(mRoom, mBroadcast)
	return srvRoom, nil
}

func (r *Room) getBroadcast(mRoom *models.Room) (mBroadcast *models.Broadcast, err error) {
	if mRoom.IsPlaying && mRoom.BroadcastID != nil {
		mBroadcast, err = models.FindBroadcastByID(mRoom.GetBroadcastID())
		if err != nil {
			r.errorCode = codes.ErrorCodeBroadcastNotFound
			return nil, err
		}
	}
	return
}

// GetByUserID room by  user_id
func (r *Room) GetByUserID(userID string) (*Room, error) {
	mRoom, err := models.FindRoomByUserID(userID)
	if err != nil {
		return &Room{}, nil
	}

	mBroadcast, err := r.getBroadcast(mRoom)
	if err != nil {
		return nil, err
	}

	srvRoom := modelRoomToSrvRoom(mRoom, mBroadcast)
	return srvRoom, nil
}
