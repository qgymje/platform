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
	RoomID      string `json:"room_id"`
	BroadcastID string `json:"broadcast_id"`
	StartTime   int64  `json:"start_time"`
}

// MessageBroadcastEnd  nsq message for host end broadcast
type MessageBroadcastEnd struct {
	RoomID      string `json:"room_id"`
	BroadcastID string `json:"broadcast_id"`
	EndTime     int64  `json:"end_time"`
}

// MessageBroadcastEnter broadcast enter
type MessageBroadcastEnter struct {
	BroadcastID string `json:"broadcast_id"`
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	Level       int64  `json:"level"`
	EnterTime   int64  `json:"enter_time"`
}

// MessageBroadcastLeave broadcast leave
type MessageBroadcastLeave struct {
	BroadcastID string `json:"broadcast_id"`
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	Level       int64  `json:"level"`
	LeaveTime   int64  `json:"leave_time"`
}

// MessageBroadcastInfo nsq message for broadcast info
type MessageBroadcastInfo struct {
	BroadcastID     string `json:"broadcast_id"`
	RoomID          string `json:"room_id"`
	TotalAudience   int64  `json:"total_audience"`
	CurrentAudience int64  `json:"current_audience"`
	StartTime       int64  `json:"start_time"`
	Duration        int64  `json:"duration"`
}

// MessageBarrage  nsq message for barrage
type MessageBarrage struct {
	BroadcastID string `json:"broadcast_id"`
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	Level       int64  `json:"level"`
	Text        string `json:"text"`
	CreatedAt   int64  `json:"created_at"`
}

// MessageSendCouponUpdate message send coupon update
type MessageSendCouponUpdate struct {
	SendCouponID string `json:"send_coupon_id"`
	BroadcastID  string `json:"broadcast_id"`
	RemainAmount int    `json:"remain_amount"`
	RemainTime   int64  `json:"remain_time"`
	CouponID     string `json:"coupon_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Description  string `json:"description"`
}

// MessageSendCouponStop stop send coupon broadcast
type MessageSendCouponStop struct {
	SendCouponID string `json:"send_coupon_id"`
	BroadcastID  string `json:"broadcast_id"`
	StopTime     int64  `json:"stop_time"`
}

// MessageSendGiftSuccess send gift success
type MessageSendGiftSuccess struct {
	UserID      string `json:"user_id"`
	ToUserID    string `json:"to_user_id"`
	GiftID      string `json:"gift_id"`
	SnowBall    uint   `json:"snow_ball"`
	SnowFlake   uint   `json:"snow_flake"`
	SuccessTime int64  `json:"success_time"`
}

// MessageSendGiftBroadcast send gift broadcast
type MessageSendGiftBroadcast struct {
	BroadcastID  string `json:"broadcast_id"`
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	GiftID       string `json:"gift_id"`
	GiftName     string `json:"gift_name"`
	GiftImage    string `json:"gift_image"`
	Combo        int    `json:"combo"`
	Amount       uint   `json:"amount"`
	TotalPrice   uint   `json:"total_price"`
	LastSendTime int64  `json:"last_send_time"`
}

// MessageChatMessage chat message
type MessageChatMessage struct {
	MessageID string `json:"message_id"`
	ChatID    string `json:"chat_id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	SendTime  int64  `json:"send_time"`
}
