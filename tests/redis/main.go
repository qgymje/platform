package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

// BroadcastGift broadcast gift
type BroadcastGift struct {
	BroadcastID  string
	UserID       string
	GiftID       string
	Hits         int
	TotalCoin    int
	LastSendTime time.Time
}

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Couldn't connect to Redis: %v\n", err)

	}
	defer conn.Close()
	// BroadcastID is the set key
	gifts := map[string]*BroadcastGift{
		"user_id_1": &BroadcastGift{BroadcastID: "123", GiftID: "1", Hits: 100, TotalCoin: 100, LastSendTime: time.Now()},
		"user_id_2": &BroadcastGift{BroadcastID: "123", GiftID: "1", Hits: 200, TotalCoin: 200, LastSendTime: time.Now()},
		"user_id_3": &BroadcastGift{BroadcastID: "123", GiftID: "1", Hits: 300, TotalCoin: 200, LastSendTime: time.Now()},
	}

	bro := "123"
	giftsJSON, err := json.Marshal(&gifts)
	if err != nil {
		log.Fatal(err)
	}
	/*
		buf := &bytes.Buffer{}
		if err = binary.Write(buf, binary.BigEndian, gifts); err != nil {
			log.Fatal(err)
		}
	*/

	if _, err := conn.Do("SET", bro, string(giftsJSON[:])); err != nil {
		log.Fatal(err)
	}

	values, err := redis.Bytes(conn.Do("GET", bro))
	if err != nil {
		log.Fatal(err)
	}

	/*
		buf2 := bytes.NewReader(values)
		var gifts2 map[string]*BroadcastGift
		err = binary.Read(buf2, binary.BigEndian, &gifts2)
		if err != nil {
			log.Fatal(err)
		}
	*/

	var gifts2 map[string]*BroadcastGift
	if err = json.Unmarshal(values, &gifts2); err != nil {
		log.Fatal(err)
	}

	for i := range gifts2 {
		log.Printf("%s: %+v", i, gifts2[i])
	}

}
