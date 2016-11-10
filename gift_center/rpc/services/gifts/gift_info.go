package gifts

import "platform/gift_center/rpc/models"

// Gift service level Gift
type Gift struct {
	GiftID    string
	Name      string
	Image     string
	SnowFlake uint
	SnowBall  uint
	Combo     int8
}

func modelGiftToSrvGift(m *models.Gift) *Gift {
	imagePath := "/v1/gift/uploads/"
	return &Gift{
		GiftID:    m.GetID(),
		Name:      m.Name,
		Image:     imagePath + m.Image,
		SnowFlake: m.SnowFlake,
		SnowBall:  m.SnowBall,
		Combo:     m.Combo,
	}
}
