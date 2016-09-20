package queues

import "time"

// MessageRegisterSMS nsq message for ssending sms service
type MessageRegisterSMS struct {
	Phone     string    `json:"phone"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

// MessageUserLogin nsq message for user login notification
type MessageUserLogin struct {
	UserID   string    `json:"userID"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	HeadImg  string    `json:"headImg"`
	RegTime  time.Time `json:"regTime"`
}

// MessageBroadcastStart  nsq message for host start to broadcast
type MessageBroadcastStart struct {
	RoomID    string    `json:"roomID"`
	StartTime time.Time `json:"startTime"`
}

// MessageBroadcastEnd  nsq message for host end broadcast
type MessageBroadcastEnd struct {
	RoomID  string    `json:"roomID"`
	EndTime time.Time `json:"endTime"`
}

// BroadcastMessageType message type, like Gift Vote Coupon
type BroadcastMessageType int

// MessageBarrage  nsq message for barrage
type MessageBarrage struct {
	UserID   string    `json:"userID"`
	UserName string    `json:"nickname"`
	Message  string    `json:"message"`
	PubTime  time.Time `json:"pubTime"`
}

// MessageGift nsq message for gift
type MessageGift struct {
	GiftID     string    `json:"giftID"`
	GiftName   string    `json:"giftName"`
	GifImage   string    `json:"giftImage"`
	GiftAmount int       `json:"gitAmount"`
	UserID     string    `json:"userID"`
	Nickname   string    `json:"nickname"`
	RoomID     string    `json:"roomID"`
	PubTime    time.Time `json:"pubTime"`
}
