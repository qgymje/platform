package services

// Message message is b byte slice
type Message []byte

func (s Message) String() string {
	return string(s)
}

// Consumer a consumer interface
type Consumer interface {
	Consume() <-chan Message
	Stop()
}

// NSQSession nsq consummer
type NSQSession struct {
	consumers []Consumer
	receive   chan Message
}

// NewNSQSession new NSQSession
func NewNSQSession(consumers []Consumer) *NSQSession {
	s := new(NSQSession)
	s.consumers = consumers
	s.receive = make(chan Message)
	return s
}

// Close stop consumer
func (s *NSQSession) Close() {
	for i := range s.consumers {
		s.consumers[i].Stop()
	}
}

func (s *NSQSession) combineConsume() {
	for i := range s.consumers {
		go func(i int) {
			for msg := range s.consumers[i].Consume() {
				//utils.Dump("[combineConsume]: got a message: ", msg.String())
				s.receive <- msg
			}
		}(i)
	}
}

// Consume combine the messages
func (s *NSQSession) Consume() <-chan Message {
	s.combineConsume()
	return s.receive
}
