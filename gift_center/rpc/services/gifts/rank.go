package gifts

import (
	"encoding/json"
	"fmt"
	"platform/commons/queues"
	"platform/commons/typeids"
	"platform/coupon_center/rpc/services/notifier"
	"platform/utils"
	"time"
)

const rankDuration = 10 * time.Second

// RankConfig rank config
type RankConfig struct {
	BroadcastID string
	RankList    []*queues.MessageSendGiftBroadcast
}

// Rank rank
type Rank struct {
	config   *RankConfig
	rankList []*queues.MessageSendGiftBroadcast
}

// NewRank new rank
func NewRank(c *RankConfig) *Rank {
	r := new(Rank)
	r.config = c
	r.rankList = []*queues.MessageSendGiftBroadcast{}
	return r
}

// Do the rank work
func (r *Rank) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("gifts.Rank.Do error: %+v", err)
		}
	}()

	r.filterRankList()

	if err = r.notify(); err != nil {
		return
	}

	return
}

func (r *Rank) isMissingRank(msg *queues.MessageSendGiftBroadcast) bool {
	return time.Since(time.Unix(msg.LastSendTime, 0)) >= rankDuration
}

func (r *Rank) filterRankList() {
	for _, rank := range r.config.RankList {
		if !r.isMissingRank(rank) {
			r.rankList = append(r.rankList, rank)
		}
	}
}

func (r *Rank) notify() error {
	return notifier.Publish(r)
}

// Topic topic
func (r *Rank) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), r.config.BroadcastID)
}

// Message publish message
func (r *Rank) Message() []byte {
	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		int(typeids.GiftSenderRank),
		r.rankList,
	}
	msg, _ := json.Marshal(data)
	return msg
}
