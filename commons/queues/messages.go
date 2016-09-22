package queues

import "time"

// MessageRegisterSMS nsq message for ssending sms service
type MessageRegisterSMS struct {
	Phone     string    `json:"phone"`
	Country   string    `json:"country"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

// MessageRegisterEmail nsq message for ssending email service
type MessageRegisterEmail struct {
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

// MessageUserLogin nsq message for user login notification
type MessageUserLogin struct {
	UserID    string    `json:"user_id"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

// MessageBroadcastStart  nsq message for host start to broadcast
type MessageBroadcastStart struct {
	RoomID    string    `json:"room_id"`
	StartTime time.Time `json:"start_time"`
}

// MessageBroadcastEnd  nsq message for host end broadcast
type MessageBroadcastEnd struct {
	RoomID  string    `json:"room_id"`
	EndTime time.Time `json:"end_time"`
}

// BroadcastMessageType message type, like Gift Vote Coupon
type BroadcastMessageType int

// MessageBarrage  nsq message for barrage
type MessageBarrage struct {
	UserID   string    `json:"user_id"`
	UserName string    `json:"nickname"`
	Message  string    `json:"message"`
	PubTime  time.Time `json:"pub_time"`
}

// MessageGift nsq message for gift
type MessageGift struct {
	GiftID     string    `json:"gift_id"`
	GiftName   string    `json:"gift_name"`
	GifImage   string    `json:"gift_image"`
	GiftAmount int       `json:"git_amount"`
	UserID     string    `json:"user_id"`
	Nickname   string    `json:"nickname"`
	RoomID     string    `json:"room_id"`
	PubTime    time.Time `json:"pub_time"`
}
