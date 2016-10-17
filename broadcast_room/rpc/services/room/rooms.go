package rooms

import (
	"platform/broadcast_room/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

// Config for Rooms
type Config struct {
	PageNum  int
	PageSize int
	Search   string
}

// Rooms rooms finder
type Rooms struct {
	config     *Config
	roomFinder *models.RoomFinder

	errorCode codes.ErrorCode
}

// NewRooms create a new Rooms object
func NewRooms(c *Config) *Rooms {
	r := new(Rooms)
	r.config = c
	r.roomFinder = models.NewRoomFinder().Limit(c.PageNum, c.PageSize).Search(c.Search)
	return r
}

// ErrorCode implement ErrorCoder
func (r *Rooms) ErrorCode() codes.ErrorCode {
	return r.errorCode
}

// Do do the dirty job
func (r *Rooms) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("rooms.Rooms.Do error: %+v", err)
		}
	}()

	if err = r.find(); err != nil {
		if err == models.ErrNotFound {
			r.errorCode = codes.ErrorCodeRoomNotFound
		} else {
			r.errorCode = codes.ErrorCodeRoomFinder
		}
		return
	}
	return
}

func (r *Rooms) find() error {
	return r.roomFinder.Do()
}

// Rooms fetch the game list object
func (r *Rooms) Rooms() []*Room {
	modelRooms := r.roomFinder.Result()
	srvRooms := []*Room{}
	for _, mRoom := range modelRooms {
		srvRoom := &Room{
			RoomID:    mRoom.GetID(),
			UserName:  mRoom.UserName,
			Name:      mRoom.Name,
			Cover:     mRoom.Cover,
			IsPlaying: mRoom.IsPlaying,
			FollowNum: mRoom.FollowNum,
		}
		srvRooms = append(srvRooms, srvRoom)
	}
	return srvRooms
}

// Count totoal result count
func (r *Rooms) Count() int64 {
	return r.roomFinder.Count()
}
