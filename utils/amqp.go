package utils

import (
	"log"

	"github.com/streadway/amqp"
)

// GetAMQP 生成一个connection以及channel
func GetAMQP() (conn *amqp.Connection, ch *amqp.Channel) {
	amqpAddr := GetConf().GetString("rabbitmq.url")
	conn, err := amqp.Dial(amqpAddr)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	ch, err = conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return
}
