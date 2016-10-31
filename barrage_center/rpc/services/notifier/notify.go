package notifier

import (
	"log"
	"sync"
	"time"

	"platform/utils"

	nsq "github.com/nsqio/go-nsq"
)

var once sync.Once
var producer *nsq.Producer

// Notifier notifier object
type Notifier interface {
	Topic() string
	Message() []byte
}

func initProducer() {
	var err error
	config := nsq.NewConfig()
	nsqdAddr := utils.GetConf().GetString("nsq.nsqd")
	producer, err = nsq.NewProducer(nsqdAddr, config)
	if err != nil {
		log.Fatalf("connect producer error: %v", err)
	}
}

// DeferredPublish delayed publish
func DeferredPublish(notifier Notifier, delay time.Duration) error {
	once.Do(func() {
		initProducer()
	})

	msg := notifier.Message()
	if len(msg) > 0 {
		err := producer.DeferredPublish(notifier.Topic(), delay, msg)
		if err != nil {
			utils.GetLog().Error("publish error: %v", err)
		} else {
			log.Printf("[publish]: topic: %s, message: %s\n", notifier.Topic(), string(msg))
		}
		return err
	}
	return nil

}

// Publish publish a message
func Publish(notifier Notifier) error {
	once.Do(func() {
		initProducer()
	})

	msg := notifier.Message()
	if len(msg) > 0 {
		err := producer.Publish(notifier.Topic(), msg)
		if err != nil {
			utils.GetLog().Error("publish error: %v", err)
		} else {
			log.Printf("[publish]: topic: %s, message: %s\n", notifier.Topic(), string(msg))
		}
		return err
	}
	return nil
}
