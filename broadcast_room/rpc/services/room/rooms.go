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
	config          *Config
	roomFinder      *models.RoomFinder
	broadcastFinder *models.BroadcastFinder

	errorCode codes.ErrorCode
}

// NewRooms create a new Rooms object
func NewRooms(c *Config) *Rooms {
	r := new(Rooms)
	r.config = c
	r.roomFinder = models.NewRoomFinder().Limit(c.PageNum, c.PageSize).Search(c.Search)
	r.broadcastFinder = models.NewBroadcastFinder()
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
	if err := r.roomFinder.Do(); err != nil {
		return err
	}

	if err := r.findBroadcasts(); err != nil {
		return err
	}

	return nil
}

func (r *Rooms) findBroadcasts() error {
	modelRooms := r.roomFinder.Result()
	broadcastIDs := []string{}
	for i := range modelRooms {
		if modelRooms[i].IsPlaying {
			broadcastIDs = append(broadcastIDs, modelRooms[i].BroadcastID.Hex())
		}
	}
	if err := r.broadcastFinder.ByIDs(broadcastIDs).Do(); err != nil {
		if err == models.ErrNotFound {
			return nil
		}
		return err
	}
	return nil
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
		if bro := r.broadcastFinder.FetchByRoomID(mRoom.GetID()); bro != nil {
			srvBroadcast := &Broadcast{
				BroadcastID:     bro.GetID(),
				RoomID:          mRoom.GetID(),
				CurrentAudience: bro.CurrentAudience,
				TotalAudience:   bro.TotalAudience,
				StartTime:       bro.StartTime,
				Duration:        bro.Duration(),
			}
			srvRoom.Broadcast = srvBroadcast
		}
		srvRooms = append(srvRooms, srvRoom)
	}
	return srvRooms
}

// Count totoal result count
func (r *Rooms) Count() int64 {
	return r.roomFinder.Count()
}
