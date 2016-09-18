package main

import (
	"flag"
	"log"

	nsq "github.com/nsqio/go-nsq"
)

var (
	topic   = flag.String("topic", "write_test", "topic name")
	channel = flag.String("channel", "default", "channel name")
)

func main() {
	flag.Parse()

	forever := make(chan struct{})
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer(*topic, *channel, config)
	log.Println("consumer: ", q)

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %s", string(message.Body))
		return nil
	}))

	err := q.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Panic("Could not connect")
	}
	<-forever
}
