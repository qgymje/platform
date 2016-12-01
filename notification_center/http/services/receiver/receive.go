package receiver

import (
	"log"
	"sync"

	"platform/utils"

	nsq "github.com/nsqio/go-nsq"
)

// Receiver infterface
type Receiver interface {
	Topic() string
	Channel() string
	Handler(<-chan []byte)
}

// Receive represent a abstrat nsq consumer object
type Receive struct {
	consumer *nsq.Consumer
	receiver Receiver
	msg      chan []byte
}

var nsqLookupdAddr string
var once sync.Once

func getNSQLookupdAddr() {
	nsqLookupdAddr = utils.GetConf().GetString("nsq.nsqlookupd")
}

// NewReceive NSQ consumer object
func NewReceive(receiver Receiver) *Receive {
	receive := &Receive{
		receiver: receiver,
		consumer: &nsq.Consumer{},
		msg:      make(chan []byte),
	}

	var err error
	config := nsq.NewConfig()
	receive.consumer, err = nsq.NewConsumer(receive.receiver.Topic(), receive.receiver.Channel(), config)
	if err != nil {
		log.Fatalf("create consumer error: %v", err)
	}
	receive.addHanler()

	once.Do(getNSQLookupdAddr)
	err = receive.consumer.ConnectToNSQLookupd(nsqLookupdAddr)
	if err != nil {
		log.Panic("Could not connect nsqlookupd")
	}
	return receive
}

// Do start to consume
func (r *Receive) Do() (err error) {
	r.receiver.Handler(r.msg)
	return
}

// Stop stop consume
func (r *Receive) Stop() {
	r.consumer.Stop()
	//utils.DeleteChannel(r.receiver.Topic(), r.receiver.Channel())
}

func (r *Receive) addHanler() {
	r.consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("[receiver] got a message: %s", string(message.Body))
		r.msg <- message.Body
		return nil
	}))
}
