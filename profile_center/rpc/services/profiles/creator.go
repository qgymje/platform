package profiles

import (
	"encoding/json"
	"platform/commons/queues"
	"platform/profile_center/rpc/models"
	"platform/sms_service/rpc/services/sms/receiver"
	"platform/utils"
)

// ListenRegisterUser listen register_user queue
func ListenRegisterUser() {
	go func() {
		NewCreator().receive()
	}()
}

// Creator consume the user_register channel and create a profile record
type Creator struct {
	userID string
}

// NewCreator new Creator
func NewCreator() *Creator {
	c := new(Creator)
	return c
}

// Do do the dirty work
func (c *Creator) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("profiles.Creator.Do error: %+v", err)
		}
	}()

	if err = c.save(); err != nil {
		return
	}

	return
}

func (c *Creator) save() (err error) {
	profileModel := &models.Profile{
		UserID: c.userID,
	}
	return profileModel.Create()
}

// Topic topic to consume
func (c *Creator) Topic() string {
	return queues.TopicUserRegister.String()
}

// Channel channel name
func (c *Creator) Channel() string {
	return queues.ChannelUserRegisterProfileCreate.String()
}

// Handler handle the incoming msg
func (c *Creator) Handler(msgs <-chan []byte) {
	for msg := range msgs {
		msgRegister, err := c.handleMessage(msg)
		if err != nil {
			utils.GetLog().Error("parse user login sms msg error: %v", err)
		} else {
			c.userID = msgRegister.UserID
			c.Do()
		}
	}
}

func (c *Creator) handleMessage(msg []byte) (*queues.MessageUserLogin, error) {
	var msgUser queues.MessageUserLogin
	if err := json.Unmarshal(msg, &msgUser); err != nil {
		return nil, err
	}
	return &msgUser, nil
}

func (c *Creator) receive() (err error) {
	return receiver.NewReceive(c).Do()
}
