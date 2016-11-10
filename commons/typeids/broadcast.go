package typeids

// BroadcastTypeID broadcat type id
type BroadcastTypeID int

const (
	// BroadcastInfo broadcast info every 30s
	BroadcastInfo BroadcastTypeID = 10000
	// CouponSender cuopon sender info every 5s
	CouponSender = 10001
	// CouponSenderStop auto stop if timeout
	CouponSenderStop = 10002
	// GiftInfo gift info
	GiftInfo = 10003
)
