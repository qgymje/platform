package gifts

import (
	"encoding/json"
	"fmt"
	"log"
	"platform/commons/codes"
	"platform/commons/queues"
	"platform/commons/typeids"
	"platform/coupon_center/rpc/services/notifier"
	"platform/gift_center/rpc/models"
	"platform/utils"
	"strconv"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

const comboDuration = 5 * time.Second

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

// BroadcasterConfig broadcast config
type BroadcasterConfig struct {
	SendGiftID string
	Username   string
}

// Broadcaster represents a broadcast
type Broadcaster struct {
	config        *BroadcasterConfig
	sendGiftModel *models.SendGift
	sendGiftMsg   *queues.MessageSendGiftBroadcast
	userID        string
	hasSentBefore bool
	errorCode     codes.ErrorCode

	rankUserIDs []string
}

// NewBroadcaster new broadcast
func NewBroadcaster(c *BroadcasterConfig) *Broadcaster {
	b := new(Broadcaster)
	b.config = c
	b.sendGiftModel = &models.SendGift{}
	b.sendGiftMsg = &queues.MessageSendGiftBroadcast{}
	b.rankUserIDs = []string{}
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

	if err = b.updateRank(); err != nil {
		b.errorCode = codes.ErrorCodeSendGiftRank
		return
	}

	if err = b.notify(); err != nil {
		b.errorCode = codes.ErrorCodeSendGiftNotify
		return
	}

	if err = b.broadcastRank(); err != nil {
		b.errorCode = codes.ErrorCodeSendGiftBroadcastRank
		return
	}

	return
}

func (b *Broadcaster) getRedisScoreKey() string {
	key := "grank_" + b.sendGiftModel.BroadcastID[8:]
	utils.Dump(key)
	return key
}

func (b *Broadcaster) getRedisHashKey() string {
	key := "glive_" + b.sendGiftModel.BroadcastID[8:]
	utils.Dump(key)
	return key
}

func (b *Broadcaster) getSendGiftID() int64 {
	id, _ := strconv.ParseInt(b.config.SendGiftID, 10, 0)
	return id
}

func (b *Broadcaster) findSendGift() error {
	b.sendGiftModel.ID = b.getSendGiftID()
	if err := b.sendGiftModel.Find(); err != nil {
		return err
	}
	b.userID = b.sendGiftModel.UserID
	return nil
}

func (b *Broadcaster) fetchGiftList() error {
	value, err := redis.Bytes(redisConn.Do("HGET", b.getRedisHashKey(), b.userID))
	if err != nil {
		if err == redis.ErrNil {
			if err2 := b.initalRedisSendGiftMessage(); err2 != nil {
				return err2
			}
		} else {
			return err
		}
	} else {
		b.hasSentBefore = true
		if err = json.Unmarshal(value, &b.sendGiftMsg); err != nil {
			return err
		}
	}
	return nil
}

func (b *Broadcaster) sendGiftModelToMessage() *queues.MessageSendGiftBroadcast {
	return &queues.MessageSendGiftBroadcast{
		BroadcastID:  b.sendGiftModel.BroadcastID,
		Username:     b.config.Username,
		GiftID:       b.sendGiftModel.GetGiftID(),
		GiftName:     b.sendGiftModel.Gift.Name,
		GiftImage:    b.sendGiftModel.Gift.Image,
		Combo:        int(b.sendGiftModel.Gift.Combo),
		Amount:       1,
		TotalPrice:   b.sendGiftModel.TotalPrice(),
		LastSendTime: b.sendGiftModel.CreatedAt.Unix(),
	}
}

func (b *Broadcaster) isSameGift() bool {
	return b.sendGiftModel.Gift.GetID() == b.sendGiftMsg.GiftID
}

func (b *Broadcaster) isMissingCombo() bool {
	return time.Since(time.Unix(b.sendGiftMsg.LastSendTime, 0)) >= comboDuration
}

func (b *Broadcaster) initalRedisSendGiftMessage() error {
	return b.resetRedisSendGiftMessage()
}

func (b *Broadcaster) resetRedisSendGiftMessage() (err error) {
	b.sendGiftMsg = b.sendGiftModelToMessage()
	jsonMsg, _ := json.Marshal(b.sendGiftMsg)
	if _, err := redisConn.Do("HSET", b.getRedisHashKey(), b.userID, string(jsonMsg[:])); err != nil {
		return err
	}
	return nil
}

func reverseStrings(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (b *Broadcaster) fetchGiftListByUserIDs(userids []string) (list []*queues.MessageSendGiftBroadcast, err error) {
	reversedUserids := reverseStrings(userids)
	for _, userid := range reversedUserids {
		if err = redisConn.Send("HGET", b.getRedisHashKey(), userid); err != nil {
			return
		}
	}

	if err = redisConn.Flush(); err != nil {
		return
	}

	for i := 0; i < len(userids); i++ {
		var msg []byte
		msg, err = redis.Bytes(redisConn.Receive())
		if err != nil {
			return
		}

		var sendGiftMsg queues.MessageSendGiftBroadcast
		if err = json.Unmarshal(msg, &sendGiftMsg); err != nil {
			return
		}
		list = append(list, &sendGiftMsg)
	}

	return
}

func (b *Broadcaster) updateReidsGiftMessage() (err error) {
	jsonMsg, _ := json.Marshal(b.sendGiftMsg)
	if _, err := redisConn.Do("HSET", b.getRedisHashKey(), b.userID, string(jsonMsg[:])); err != nil {
		return err
	}
	return nil
}

func (b *Broadcaster) updateRedisScore() (err error) {
	if _, err = redisConn.Do("ZADD", b.getRedisScoreKey(), b.sendGiftMsg.TotalPrice, b.userID); err != nil {
		return
	}
	return
}

func (b *Broadcaster) fetchRank() (err error) {
	if b.rankUserIDs, err = redis.Strings(redisConn.Do("ZRANGE", b.getRedisScoreKey(), 0, 3)); err != nil {
		return
	}
	return
}

func (b *Broadcaster) updateRank() (err error) {
	if !b.hasSentBefore {
		return b.updateRedisScore()
	}

	if !b.isSameGift() || b.isMissingCombo() {
		if err = b.resetRedisSendGiftMessage(); err != nil {
			return
		}
		return b.updateRedisScore()
	}
	b.sendGiftMsg.Amount++
	b.sendGiftMsg.TotalPrice += b.sendGiftModel.TotalPrice()
	if err = b.updateReidsGiftMessage(); err != nil {
		return err
	}
	return b.updateRedisScore()
}

func (b *Broadcaster) isInRank() bool {
	for _, userid := range b.rankUserIDs {
		if userid == b.sendGiftModel.UserID {
			return true
		}
	}
	return false
}

func (b *Broadcaster) broadcastRank() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("gifts.Broadcaster.broadcastRank error:%+v", err)
		}
	}()

	if err = b.fetchRank(); err != nil {
		return
	}

	if b.isInRank() {
		rankList, err := b.fetchGiftListByUserIDs(b.rankUserIDs)
		if err != nil {
			return err
		}
		rankConfig := &RankConfig{
			BroadcastID: b.sendGiftModel.BroadcastID,
			RankList:    rankList,
		}
		rank := NewRank(rankConfig)
		return rank.Do()
	}

	return
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
	data := struct {
		Type int         `json:"type"`
		Data interface{} `json:"data"`
	}{
		int(typeids.GiftSenderInfo),
		b.sendGiftMsg,
	}
	msg, _ := json.Marshal(data)
	return msg
}
