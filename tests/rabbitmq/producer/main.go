package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// rabbitmqctl add_user odyssey odyssey
	// rabbitmqctrl set_user_tags odyssey administrator
	// 如果想要这个用户还可以在web页面上控制vhost, 则此用户要打上admin的tag(也就是角色)
	// rabbitmqctl add_vhost odyssey
	// rabbitmqctl set_permissions -p odyssey odyssey ".*" ".*" ".*"
	// mark: -p后是vhost_name user_name
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel() //创建一个新channel
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	confirm(ch)
	var i int
	for {
		publish(ch, "hello", "direct", "hello-key", fmt.Sprintf(`{"msg":"hello_%d"}`, i))
		i++
		time.Sleep(1 * time.Second)
	}
}

func confirm(ch *amqp.Channel) {
	// 确保消息收到
	//func (me *Channel) Confirm(noWait bool) error {
	if err := ch.Confirm(false); err != nil {
		log.Fatal(err)
	}

	//func (me *Channel) NotifyPublish(confirm chan Confirmation) chan Confirmation {
	confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))
	go func() {
		for {
			if confirmed := <-confirms; confirmed.Ack {
				log.Printf("consume confirmed ack, tag: %d", confirmed.DeliveryTag) // tag should be increased!
			} else {
				log.Printf("consume failed ack: tag: %d", confirmed.DeliveryTag)
			}
		}
	}()
}

func publish(ch *amqp.Channel, exchange string, kind string, routingKey string, body string) {
	// 发送者将消息绑定到交换机止,以提高灵活性
	//func (me *Channel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args Table) error {
	if err := ch.ExchangeDeclare(
		exchange, // exchange name
		kind,     // exchange type
		true,     // durable 指exchange是否保存
		false,    // autoDelete 是exchnage是否自动删除
		false,
		false,
		nil,
	); err != nil {
		log.Fatal(err)
	}

	//func (me *Channel) Publish(exchange, key string, mandatory, immediate bool, msg Publishing) error {
	err := ch.Publish(
		exchange,
		routingKey, // routing key is here
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/palin",
			ContentEncoding: "",
			DeliveryMode:    amqp.Transient,
			Body:            []byte(body),
			Priority:        0,
		},
	)
	log.Printf(" [x] Sent %s", body)
	if err != nil {
		log.Fatal(err)
	}
}
