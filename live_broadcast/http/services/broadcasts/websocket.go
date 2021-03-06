package broadcasts

import (
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client represent a broadcast client
type Client struct {
	conn     *websocket.Conn
	send     chan Message
	producer *NSQSession
	consumer *NSQSession
}

// NewClient create a new client
func NewClient(conn *websocket.Conn, nsqlookupdAddr []string, nsqdAddr, topic, channel string) *Client {
	producer := NewNSQSession(nsqlookupdAddr, nsqdAddr, topic, channel)
	consumer := NewNSQSession(nsqlookupdAddr, nsqdAddr, topic, channel)

	return &Client{
		conn:     conn,
		send:     make(chan Message, 256),
		producer: producer,
		consumer: consumer,
	}
}

// ReadPump read from client
func (c *Client) ReadPump() {
	defer func() {
		c.producer.Close()
		c.consumer.Close()
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	c.producer.Publish(c.send)

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.send <- message
	}
	close(c.send)
}

func (c *Client) write(mt int, payload []byte) error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(mt, payload)
}

// WritePump  write to client
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	amqpmsgchan := c.consumer.Consume()

	for {
		select {
		case message, ok := <-amqpmsgchan:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
