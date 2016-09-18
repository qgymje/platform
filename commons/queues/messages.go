package queues

import "time"

// MessageUserLogin nsq消息结构
type MessageUserLogin struct {
	UserID   string    `json:"userID"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	HeadImg  string    `json:"headImg"`
	RegTime  time.Time `json:"regTime"`
}

// MessageBroadcastStart 表示开始直播的消息结构
type MessageBroadcastStart struct {
	RoomID    string    `json:"roomID"`
	StartTime time.Time `json:"startTime"`
}

// MessageBroadcastEnd 表示结束直播的消息结构
type MessageBroadcastEnd struct {
	RoomID  string    `json:"roomID"`
	EndTime time.Time `json:"endTime"`
}

type BroadcastMessageType int

// MessageBarrage 直播过程的弹幕消息结构
type MessageBarrage struct {
	UserID   string    `json:"userID"`
	UserName string    `json:"nickname"`
	Message  string    `json:"message"`
	PubTime  time.Time `json:"pubTime"`
}

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
