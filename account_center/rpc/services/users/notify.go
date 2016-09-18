package users

import (
	"log"
	"sync"

	"platform/utils"

	nsq "github.com/nsqio/go-nsq"
)

var once sync.Once
var producer *nsq.Producer

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

func Publish(notifier Notifier) error {
	once.Do(func() {
		initProducer()
	})

	msg := notifier.Message()

	err := producer.Publish(notifier.Topic(), msg)
	if err != nil {
		utils.GetLog().Error("publish error: %v, msg: %s", err, string(msg))
	}
	return err
}
