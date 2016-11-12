package gifts

import (
	"encoding/json"
	"fmt"
	"log"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/coupon_center/rpc/services/notifier"
	"platform/gift_center/rpc/models"
	"platform/utils"
	"strconv"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

var redisConn redis.Conn
var once sync.Once

// InitRedis init redis conn
func InitRedis() {
	once.Do(func() {
		var err error
		redisConn, err = utils.GetRedis()
		if err != nil {
			log.Fatal("redis conn error: ", err)
		}
	})
}

const comboDuration = 5 * time.Second

// BroadcasterConfig broadcast config
type BroadcasterConfig struct {
	SendGiftID string
	Username   string
}

// Broadcaster represents a broadcast
type Broadcaster struct {
	config        *BroadcasterConfig
	sendGiftModel *models.SendGift
	giftList      []*queues.MessageSendGiftBroadcast
	errorCode     codes.ErrorCode
}

// NewBroadcaster new broadcast
func NewBroadcaster(c *BroadcasterConfig) *Broadcaster {
	b := new(Broadcaster)
	b.config = c
	b.sendGiftModel = &models.SendGift{}
	b.giftList = []*queues.MessageSendGiftBroadcast{}
	return b
}

// ErrorCode error code
func (b *Broadcaster) ErrorCode() codes.ErrorCode {
	return b.errorCode
}

// Do do the dirty work
func (b *Broadcaster) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("gifts.Broadcaster.Do error: %+v", err)
		}
	}()

	if err = b.findSendGift(); err != nil {
		b.errorCode = codes.ErrorCodeSendGiftNotFound
		return
	}

	if err = b.fetchGiftList(); err != nil {
		b.errorCode = codes.ErrorCodeSendGiftListNotFound
		return
	}

	if err = b.reSort(); err != nil {
		b.errorCode = codes.ErrorCodeSendGiftResort
		return
	}

	return
}

func (b *Broadcaster) getRedisKey() string {
	// not allowed by the redis client
	return b.sendGiftModel.BroadcastID[12:]
}

func (b *Broadcaster) getSendGiftID() int64 {
	id, _ := strconv.ParseInt(b.config.SendGiftID, 10, 0)
	return id
}

func (b *Broadcaster) findSendGift() error {
	b.sendGiftModel.ID = b.getSendGiftID()
	return b.sendGiftModel.Find()
}

func (b *Broadcaster) fetchGiftList() error {
	reply, err := redis.ByteSlices(redisConn.Do("SMEMBERS", b.getRedisKey()))
	if err != nil {
		utils.Dump(err)
		return err
	}

	for i := range reply {
		var msg queues.MessageSendGiftBroadcast
		_ = json.Unmarshal(reply[i], &msg)
		b.giftList = append(b.giftList, &msg)
	}

	return nil
}

func (b *Broadcaster) sendGiftModelToMessage() *queues.MessageSendGiftBroadcast {
	return &queues.MessageSendGiftBroadcast{
		BroadcastID:  b.sendGiftModel.BroadcastID,
		Username:     b.config.Username,
		GiftID:       b.sendGiftModel.GetGiftID(),
		Ammount:      1,
		LastSendTime: b.sendGiftModel.CreatedAt.Unix(),
	}
}

func (b *Broadcaster) reSort() error {
	if len(b.giftList) == 0 {
		msg := b.sendGiftModelToMessage()
		jsonMsg, _ := json.Marshal(msg)
		if _, err := redisConn.Do("SADD", b.getRedisKey(), string(jsonMsg)); err != nil {
			return err
		}
		b.giftList = append(b.giftList, msg)
		return nil
	}

	return nil
}

func (b *Broadcaster) notify() error {
	return notifier.Publish(b)
}

// Topic topic
func (b *Broadcaster) Topic() string {
	return fmt.Sprintf(queues.TopicBroadcastFormat.String(), b.sendGiftModel.BroadcastID)
}

// Message publish message
func (b *Broadcaster) Message() []byte {
	// fetch from redis
	return nil
}
