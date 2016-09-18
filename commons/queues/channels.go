package queues

type Channel string

func (c Channel) String() string {
	return string(c)
}

const (
	ChannelBroadcastStart   Channel = "broadcast_start"
	ChannelBroadcastEnd     Channel = "broadcast_end"
	ChannelBroadcastBarrage Channel = "default"
)
