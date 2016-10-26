package broadcasts

import (
	"encoding/json"
	"fmt"
	"platform/broadcast_room/rpc/models"
	"platform/broadcast_room/rpc/services/notifier"
	"platform/commons/queues"
	"platform/utils"
	"time"
)

// Sync broadcast all broadcast info every 30s
// will be called by the main
// must be caaaed as go Sync()
func Sync() {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			allBroadcasts, err := fetchAllPlayingBroadcasts()
			if err != nil {
				utils.GetLog().Error("broadcasts.Sync.fetchAllPlayingBroadcasts error: %+v", err)
			}
			for _, bro := range allBroadcasts {
				broSync := NewBroadcastSync(bro)
				if err := broSync.Do(); err != nil {
					utils.GetLog().Error("broadcasts.BroadcastSync.Do error: %+v", err)
				}
			}
		}
	}
}

// notice: the number of broadcast will cause a lot of memory
func fetchAllPlayingBroadcasts() ([]*Broadcast, error) {
	broadcastFinder := models.NewBroadcastFinder().IsPlaying()
	if err := broadcastFinder.Do(); err != nil {
		return nil, err
	}
	mBroadcasts := broadcastFinder.Result()
	srvBroadcasts := []*Broadcast{}
	for i := range mBroadcasts {
		srvBroadcast := modelBroadcastToSrvBroadcast(mBroadcasts[i])
		srvBroadcasts = append(srvBroadcasts, srvBroadcast)
	}
	return srvBroadcasts, nil
}

// BroadcastSync sync broadcast
type BroadcastSync struct {
	broadcast *Broadcast
}

// NewBroadcastSync create a broadcast sync object
func NewBroadcastSync(broadcast *Broadcast) *BroadcastSync {
	b := new(BroadcastSync)
	b.broadcast = broadcast
	return b
}

// Do do the dirty job
func (b *BroadcastSync) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("broadcasts.BroadcastSync.Do error: %+v", err)
		}
	}()

	if err = b.notify(); err != nil {
		return err
	}
	return
}

func (b *BroadcastSync) notify() error {
	return notifier.Publish(b)
}

// Topic publish topic
func (b *BroadcastSync) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), b.broadcast.BroadcastID)
}

// Message publish message
func (b *BroadcastSync) Message() []byte {
	var msg []byte
	broadcastMsg := queues.MessageBroadcastInfo{
		BroadcastID:     b.broadcast.BroadcastID,
		RoomID:          b.broadcast.RoomID,
		TotalAudience:   b.broadcast.TotalAudience,
		CurrentAudience: b.broadcast.CurrentAudience,
		StartTime:       b.broadcast.StartTime.Unix(),
		Duration:        b.broadcast.Duration,
	}

	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		10000,
		broadcastMsg,
	}
	msg, _ = json.Marshal(&data)
	return msg
}
