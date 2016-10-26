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
	BroadcastID string    `json:"broadcast_id"`
	UserID      string    `json:"user_id"`
	Username    string    `json:"username"`
	Level       int64     `json:"level"`
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"created_at"`
}

// MessageSendCouponUpdate message send coupon update
type MessageSendCouponUpdate struct {
	SendCouponID string `json:"send_coupon_id"`
	BroadcastID  string `json:"broadcast_id"`
	Number       int    `json:"number"`
	Duration     int64  `json:"duration"`
}
