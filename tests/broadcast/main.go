package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var (
	host    = "localhost:8787"
	testURL = "/broadcasting/v1/%d/%d"
	room    = flag.Int("room", 2, "room number")
	user    = flag.Int("user", 2, "each room user number")
)

/*
测试性能

1. N个房间里有M个用户同时发送消息, 每个用户都能收到房间里其它用户的信息

*/

func main() {
	flag.Parse()

	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	urls := CreateTestURLs(*room, *user)
	for _, u := range urls {
		go CreateWSClient(u, interrupt)
	}

	<-done
}

func CreateWSClient(path string, interrupt chan os.Signal) {
	u := url.URL{Scheme: "ws", Host: host, Path: path}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("%s recv: %s", path, message)
		}
	}()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
			}

		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close: ", err)
				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}
}

func CreateTestURLs(room, user int) []string {
	userNum := 0
	urls := []string{}
	for i := 0; i < room; i++ {
		for j := 0; j < user; j++ {
			url := fmt.Sprintf(testURL, i, userNum)
			urls = append(urls, url)
			userNum++
		}
	}
	return urls
}
