package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

// BroadcastGift broadcast gift
type BroadcastGift struct {
	BroadcastID string
	UserID      string
	GiftID      string
	Hits        int
	TotalCoin   int
	TTL         int
}

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Couldn't connect to Redis: %v\n", err)

	}
	defer conn.Close()
	// BroadcastID is the set key
	gifts := []*BroadcastGift{
		&BroadcastGift{BroadcastID: "123", UserID: "1", GiftID: "1", Hits: 100, TotalCoin: 100},
		&BroadcastGift{BroadcastID: "123", UserID: "2", GiftID: "1", Hits: 200, TotalCoin: 200},
		&BroadcastGift{BroadcastID: "123", UserID: "3", GiftID: "1", Hits: 300, TotalCoin: 200},
	}

	bro := "123"

	if _, err := conn.Do("HSET", redis.Args{bro}.AddFlat(gifts)...); err != nil {
		log.Fatal(err)
	}

	values, err := redis.Values(conn.Do("HGETALL", bro))
	if err != nil {
		log.Fatal(err)
	}
	var gift []*BroadcastGift
	if err = redis.ScanSlice(values, &gift); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s: %+v", bro, &gift)

}
