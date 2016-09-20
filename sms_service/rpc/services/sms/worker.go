package sms

import (
	"encoding/json"
	"platform/commons/queues"
	"platform/sms_service/rpc/services/sms/receiver"
	"platform/utils"
)

// ListenRegisterSMS start a go routine to listen register sms message
func ListenRegisterSMS() {
	go func() {
		(&RegisterWorker{}).receive()
	}()
}

// RegisterWorker receive register sms code message
type RegisterWorker struct {
}

// Topic which topic to listen
func (s *RegisterWorker) Topic() string {
	return queues.TopicRegisterSMS.String()
}

// Channel default channel
func (s *RegisterWorker) Channel() string {
	return queues.ChannelDefault.String()
}

func (s *RegisterWorker) handleMessage(msg []byte) (*queues.MessageRegisterSMS, error) {
	var msgSMS queues.MessageRegisterSMS
	if err := json.Unmarshal(msg, &msgSMS); err != nil {
		return nil, err
	}
	return &msgSMS, nil
}

// Handler handle the message
func (s *RegisterWorker) Handler(msgs <-chan []byte) {
	for msg := range msgs {
		msgSMS, err := s.handleMessage(msg)
		if err != nil {
			utils.GetLog().Error("parse register sms msg error: %v", err)
		} else {
			config := &RegisterCodeConfig{
				Phone:     msgSMS.Phone,
				Code:      msgSMS.Code,
				CreatedAt: msgSMS.CreatedAt,
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
