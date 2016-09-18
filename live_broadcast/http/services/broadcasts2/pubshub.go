package broadcasts2

import (
	"crypto/sha1"
	"fmt"
	"log"

	nsq "github.com/nsqio/go-nsq"
)

type message []byte

func (s message) String() string {
	return string(s)
}

// NSQSession 表示一个pubsub对象
type NSQSession struct {
	producer *nsq.Producer
	consumer *nsq.Consumer
	topic    string
	channel  string
	receive  chan message
}

// NewNSQSession 创建一个基于NSQ的pubsub对象
func NewNSQSession(nsqlookupdAddr string, nsqdAddr string, topic, channel string) *NSQSession {
	log.Println(nsqlookupdAddr, nsqdAddr, topic, channel)
	session := NSQSession{
		producer: &nsq.Producer{},
		consumer: &nsq.Consumer{},
		topic:    topic,
		channel:  channel,
		receive:  make(chan message),
	}
	config := nsq.NewConfig()
	var err error
	session.producer, err = nsq.NewProducer(nsqdAddr, config)
	if err != nil {
		log.Panic("Cound not connect nsqd")
	}

	config2 := nsq.NewConfig()
	session.consumer, err = nsq.NewConsumer(topic, channel, config2)
	if err != nil {
		log.Fatalf("create consumer error: %v", err)
	}
	session.addHanler()
	err = session.consumer.ConnectToNSQLookupd(nsqlookupdAddr)
	if err != nil {
		log.Panic("Could not connect nsqlookupd")
	}
	return &session
}

// Close 关闭NSQ的连接
func (s *NSQSession) Close() {
	s.producer.Stop()
	s.consumer.Stop()
}

// Publish 发布一条消息
func (s *NSQSession) Publish(msg <-chan message) {
	go func() {
		for m := range msg {
			err := s.producer.Publish(s.topic, []byte(m))
			if err != nil {
				log.Println("publish error: %s", m.String())
			}
		}
	}()
}

func (s *NSQSession) addHanler() {
	s.consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %s", string(message.Body))
		s.receive <- message.Body
		return nil
	}))
}

// Consume 用于收消息
func (s *NSQSession) Consume() <-chan message {
	return s.receive
}

// GenChannelName 生成channel名字
func GenChannelName(id string) string {
	h := sha1.New()
	fmt.Fprint(h, id)
	return fmt.Sprintf("%x", h.Sum(nil))
}
