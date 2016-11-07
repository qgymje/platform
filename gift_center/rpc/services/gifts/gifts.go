package gifts

import (
	"platform/commons/codes"
	"platform/gift_center/rpc/models"
	"platform/utils"
)

// Gifts gifts list
type Gifts struct {
	giftModels []*models.Gift

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

func (g *Gifts) findGifts() (err error) {
	g.giftModels, err = models.FindGifts()
	return
}

// Gifts get the gifts
func (g *Gifts) Gifts() []*Gift {
	gs := []*Gift{}
	for i := range g.giftModels {
		gift := modelGiftToSrvGift(g.giftModels[i])
		gs = append(gs, gift)
	}
	return gs
}
