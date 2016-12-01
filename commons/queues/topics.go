package queues

// Topic nsq topic name
type Topic string

func (t Topic) String() string {
	return string(t)
}

const (
	TopicRegisterSMS   Topic = "register_sms"
	TopicRegisterEmail Topic = "register_email"
	TopicUserRegister  Topic = "user_register"
	TopicUserLogin     Topic = "user_login"

	TopicGamePublish Topic = "game_publish"
	TopicGameUpdate  Topic = "game_update"

	TopicBroadcastStart  Topic = "broadcast_start"
	TopicBroadcastEnd    Topic = "broadcast_end"
	TopicBroadcastEnter  Topic = "broadcast_enter"
	TopicBroadcastLeave  Topic = "broadcast_leave"
	TopicBroadcastFormat Topic = "broadcast_%s"

	TopicSendCouponUpdate Topic = "sendcoupon_update"

	TopicSendGiftSuccess Topic = "sendgift_success"

	TopicUserFormat Topic = "user_%s"
)
