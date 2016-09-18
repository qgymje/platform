package main

import (
	"bufio"
	"context"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/streadway/amqp"
	"tech.cloudzen/utils"
)

var url = flag.String("url", "amqp:///", "AMQP url fot both publisher and subscriber")

const exchange = "pubsub"

type message []byte

type amqpSession struct {
	*amqp.Connection
	*amqp.Channel
}

func (s amqpSession) Close() error {
	if s.Channel == nil {
		return nil
	}
	return s.Connection.Close()
}

func redial(ctx context.Context, url string) chan chan amqpSession {
	sessions := make(chan chan amqpSession)

	go func() {
		sess := make(chan amqpSession)
		defer close(sessions)

		for {
			select {
			case sessions <- sess:
			case <-ctx.Done():
				log.Println("shutting down session factory")
			}

			conn, err := amqp.Dial(url)
			if err != nil {
				log.Fatalf("cannot create channel: %v", err)
			}

			ch, err := conn.Channel()
			if err != nil {
				log.Fatalf("cannot create channel: %v", err)
			}

			//func (me *Channel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args Table) error {
			if err := ch.ExchangeDeclare(exchange, "fanout", false, true, false, false, nil); err != nil {
				log.Fatalf("cannot declare fanout exchange: %v", err)
			}

			select {
			case sess <- amqpSession{conn, ch}:
			case <-ctx.Done():
				log.Println("shutting down new session")
				return
			}
		}
	}()
	return sessions
}

func publish(sessions chan chan amqpSession, messages <-chan message) {
	var (
		running bool
		reading = messages
		pending = make(chan message, 1)
		confirm = make(chan amqp.Confirmation, 1)
	)

	for session := range sessions {
		log.Println("ranging sessions")
		pub := <-session

		if err := pub.Confirm(false); err != nil {
			log.Printf("publisher confirms not supported")
			close(confirm)
		} else {
			log.Println("notify publish")
			pub.NotifyPublish(confirm)
		}

		log.Printf("publishing...")

		for {
			log.Println("for loop.")
			var body message
			select {
			case confirmed := <-confirm:
				log.Println("received confirm.")
				utils.Dump(confirmed)
				if !confirmed.Ack {
					log.Fatalf("nack message %d, body: %q", confirmed.DeliveryTag, string(body))
				}
				reading = messages

			case body = <-pending:
				log.Println("received pending.")
				routingKey := "ignored for fanout exchanges, application dependent for other exchnages"
				err := pub.Publish(exchange, routingKey, false, false, amqp.Publishing{
					Body: body,
				})

				if err != nil {
					pending <- body
					pub.Close()
					break
				}

			case body, running = <-reading:
				log.Println("received reading.")
				if !running {
					return
				}
				pending <- body
				reading = nil
			}
		}
	}
}

func identify() string {
	hostname, err := os.Hostname()
	h := sha1.New()
	fmt.Fprint(h, hostname)
	fmt.Fprint(h, err)
	fmt.Fprint(h, os.Getpid())
	return fmt.Sprintf("%x", h.Sum(nil))
}

func subscribe(sessions chan chan amqpSession, messages chan<- message) {
	queue := identify()

	for session := range sessions {
		sub := <-session

		if _, err := sub.QueueDeclare(queue, false, true, true, false, nil); err != nil {
			log.Printf("cannot consume from exclusive queue: %q, %v", queue, err)
			return
		}

		routingKey := "application specific routing key for fancy toplogies"
		if err := sub.QueueBind(queue, routingKey, exchange, false, nil); err != nil {
			log.Printf("cannot consume without a binding to exhcnage: %q, %v", exchange, err)
			return
		}

		deliveries, err := sub.Consume(queue, "", false, true, false, false, nil)
		if err != nil {
			log.Printf("cannot consume from: %q, %v", queue, err)
			return
		}

		log.Printf("subscribed...")

		for msg := range deliveries {
			messages <- message(msg.Body)
			sub.Ack(msg.DeliveryTag, false)
		}
	}
}

func read(r io.Reader) <-chan message {
	lines := make(chan message)
	go func() {
		defer close(lines)
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			lines <- message(scan.Bytes())
		}
	}()
	return lines
}

func write(w io.Writer) chan<- message {
	lines := make(chan message)
	go func() {
		for line := range lines {
			fmt.Fprintln(w, string(line))
		}
	}()
	return lines
}

func main() {
	flag.Parse()

	ctx, done := context.WithCancel(context.Background())
	go func() {
		publish(redial(ctx, *url), read(os.Stdin))
		done()
	}()

	go func() {
		subscribe(redial(ctx, *url), write(os.Stdout))
		done()
	}()

	<-ctx.Done()
}
