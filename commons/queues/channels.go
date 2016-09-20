package queues

// Channel nsq channel type
type Channel string

func (c Channel) String() string {
	return string(c)
}

const (
	// ChannelDefault default channel
	ChannelDefault Channel = "default"
	// ChannelBroadcastStart broadcast start
	ChannelBroadcastStart Channel = "broadcast_start"
	// ChannelBroadcastEnd broadcast end
	ChannelBroadcastEnd Channel = "broadcast_end"
	// ChannelBroadcastBarrage barrage channel name
	ChannelBroadcastBarrage Channel = "default"
)
