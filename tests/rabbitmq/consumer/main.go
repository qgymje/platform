package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	//func (me *Channel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error)
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	//func (me *Channel) QueueBind(name, key, exchange string, noWait bool, args Table) error
	if err = ch.QueueBind(
		q.Name,
		"hello-key", // here is the binding key
		"hello",     // exchange name
		false,
		nil,
	); err != nil {
		log.Fatal(err)
	}

	//func (me *Channel) Qos(prefetchCount, prefetchSize int, global bool) error
	//直到ack之后才接收新的消息
	err = ch.Qos(1, 0, false)

	//func (me *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args Table) (<-chan Delivery, error)
	msgs, err := ch.Consume(
		q.Name,
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(true)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C ")
	<-forever

}
