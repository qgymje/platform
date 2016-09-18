package main

import (
	"flag"
	"fmt"
	"log"

	nsq "github.com/nsqio/go-nsq"
)

var (
	topic = flag.String("topic", "write_test", "topic name")
)

func main() {
	flag.Parse()
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("%d", i)
		err := w.Publish(*topic, []byte(msg))
		if err != nil {
			log.Panic("could not connect")
		}
	}

	w.Stop()
}
