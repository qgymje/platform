package queues

// Topic nsq topic name
type Topic string

func (t Topic) String() string {
	return string(t)
}

const (
	// TopicRegisterSMS sms code topic name
	TopicRegisterSMS Topic = "register_sms"
	// TopicUserRegister user register topic name
	TopicUserRegister Topic = "user_register"
	// TopicUserLogin user login topic name
	TopicUserLogin Topic = "user_login"
	// TopicBroadcastStart broadcast start topic name
	TopicBroadcastStart Topic = "broadcast_start"
	// TopicBroadcastEnd broadcast end topic name
	TopicBroadcastEnd Topic = "broadcast_end"
	// TopicBroadcastRoomFormat broadcast room format
	TopicBroadcastRoomFormat Topic = "room_%s"
)
