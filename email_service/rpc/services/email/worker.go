package email

import (
	"encoding/json"
	"platform/commons/queues"
	"platform/email_service/rpc/services/email/receiver"
	"platform/utils"
)

// ListenRegisterEmail start a go routine to listen register sms message
func ListenRegisterEmail() {
	go func() {
		(&RegisterWorker{}).receive()
	}()
}

// RegisterWorker receive register sms code message
type RegisterWorker struct {
}

// Topic which topic to listen
func (s *RegisterWorker) Topic() string {
	return queues.TopicRegisterEmail.String()
}

// Channel default channel
func (s *RegisterWorker) Channel() string {
	return queues.ChannelDefault.String()
}

func (s *RegisterWorker) handleMessage(msg []byte) (*queues.MessageRegisterEmail, error) {
	var msgEmail queues.MessageRegisterEmail
	if err := json.Unmarshal(msg, &msgEmail); err != nil {
		return nil, err
	}
	return &msgEmail, nil
}

// Handler handle the message
func (s *RegisterWorker) Handler(msgs <-chan []byte) {
	for msg := range msgs {
		msgEmail, err := s.handleMessage(msg)
		if err != nil {
			utils.GetLog().Error("parse register sms msg error: %v", err)
		} else {
			config := &RegisterCodeConfig{
				Email:     msgEmail.Email,
				Code:      msgEmail.Code,
				CreatedAt: msgEmail.CreatedAt,
			}
			registerCode := NewRegisterCode(config)
			if err := registerCode.Create(); err != nil {
				utils.GetLog().Error("create register code error: %v", err)
			}
		}
	}
}

func (s *RegisterWorker) receive() error {
	return receiver.NewReceive(s).Do()
}
