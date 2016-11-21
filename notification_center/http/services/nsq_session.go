package services

import "platform/notification_center/http/services/receiver"

// Message message is b byte slice
type Message []byte

func (s Message) String() string {
	return string(s)
}

// Consumer a consumer interface
type Consumer interface {
	receiver.Receiver
	Stop() error
}

// NSQSession nsq consummer
type NSQSession struct {
	consumers []*Consumer
	receive   chan Message
}

// NewNSQSession new NSQSession
func NewNSQSession(consumers []*Consumer) *NSQSession {
	s := new(NSQSession)
	s.consumers = consuemrs
	s.receive = make(chan Message)
	return s
}

// Stop stop consumer
func (s *NSQSession) Stop() {
	for i := range s.consumers {
		s.consumers[i].Stop()
	}
}

func (s *NSQSession) combineConsume() {
	for {
		for i := range s.consumers {
			s.receive <- s.consumers[i].Handler()
		}
	}
}

// Consume combine the messages
func (s *NSQSession) Consume() <-chan Message {
	s.combineConsume()
	return s.receive
}
