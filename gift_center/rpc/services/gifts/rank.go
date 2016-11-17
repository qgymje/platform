package gifts

import (
	"fmt"
	"platform/commons/queues"
)

// RankConfig rank config
type RankConfig struct {
	BroadcastID string
}

// Rank rank
type Rank struct {
	config *RankConfig
}

// NewRank new rank
func NewRank(c *RankConfig) *Rank {

}

// Topic topic
func (r *Rank) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), r.config.BroadcastID)
}
