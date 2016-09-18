package queues

type Topic string

func (t Topic) String() string {
	return string(t)
}

const (
	TopicUserRegister Topic = "user_register"
	TopicUserLogin    Topic = "user_login"

	TopicBroadcastStart      Topic = "broadcast_start"
	TopicBroadcastEnd        Topic = "broadcast_end"
	TopicBroadcastRoomFormat Topic = "room_%s"
)
