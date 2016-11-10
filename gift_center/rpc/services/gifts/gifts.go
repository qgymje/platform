package gifts

import (
	"platform/commons/codes"
	"platform/gift_center/rpc/models"
	"platform/utils"
)

// Gifts gifts list
type Gifts struct {
	giftModels []*models.Gift
	giftID     string

	errorCode codes.ErrorCode
}

// NewGifts create  new gifts
func NewGifts() *Gifts {
	g := new(Gifts)
	g.giftModels = []*models.Gift{}
	return g
}

// ErrorCode error code
func (g *Gifts) ErrorCode() codes.ErrorCode {
	return g.errorCode
}

// Do do the dirty work
func (g *Gifts) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("gifts.Gifts.Do error: %+v", err)
		}
	}()

	if err = g.findGifts(); err != nil {
		g.errorCode = codes.ErrorCodeGiftNotFound
		return
	}

	return
}

// SetGiftID set a single gift id
func (g *Gifts) SetGiftID(id string) *Gifts {
	g.giftID = id
	return g
}

func (g *Gifts) findGifts() (err error) {
	if g.giftID != "" {
		giftModel, err := models.FindGiftByID(g.giftID)
		if err != nil {
			return err
		}
		g.giftModels = append(g.giftModels, giftModel)
	} else {
		g.giftModels, err = models.FindGifts()
	}
	return
}

// Result get the gifts
func (g *Gifts) Result() []*Gift {
	gs := []*Gift{}
	for i := range g.giftModels {
		gift := modelGiftToSrvGift(g.giftModels[i])
		gs = append(gs, gift)
	}
	return gs
}

// One returns single gift entity
func (g *Gifts) One() *Gift {
	return modelGiftToSrvGift(g.giftModels[0])
}
