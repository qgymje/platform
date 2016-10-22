package queues

// Topic nsq topic name
type Topic string

func (t Topic) String() string {
	return string(t)
}

const (
	// TopicRegisterSMS sms code topic name
	TopicRegisterSMS Topic = "register_sms"
	// TopicRegisterEmail register code topic name
	TopicRegisterEmail Topic = "register_email"
	// TopicUserRegister user register topic name
	TopicUserRegister Topic = "user_register"
	// TopicUserLogin user login topic name
	TopicUserLogin Topic = "user_login"

	// TopicGamePublish game publish
	TopicGamePublish Topic = "game_publish"
	// TopicGameUpdate game update
	TopicGameUpdate Topic = "game_update"

	// TopicBroadcastStart broadcast start topic name
	TopicBroadcastStart Topic = "broadcast_start"
	// TopicBroadcastEnd broadcast end topic name
	TopicBroadcastEnd Topic = "broadcast_end"
	// TopicBroadcastFormat broadcast format
	TopicBroadcastFormat Topic = "broadcast_%s"
)
