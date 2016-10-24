package barrages

import (
	"platform/barrage_center/rpc/models"
	"platform/commons/codes"
	"platform/utils"
)

// Config config
type Config struct {
	TypeID      int
	BroadcastID string
	StartTime   int64
	EndTime     int64
	PageNum     int
	PageSize    int
}

// Barrages barrages servive level object
type Barrages struct {
	config        *Config
	barrageFinder *models.BarrageFinder

	errorCode codes.ErrorCode
}

// NewBarrages create a barrages object
func NewBarrages(c *Config) *Barrages {
	b := new(Barrages)
	b.config = c
	b.barrageFinder = models.NewBarrageFinder().Limit(c.PageNum, c.PageSize).BroadcastID(c.BroadcastID).Duration(c.StartTime, c.EndTime)
	return b
}

// ErrorCode implement ErrorCoder
func (b *Barrages) ErrorCode() codes.ErrorCode {
	return b.errorCode
}

// Do do the dirty work
func (b *Barrages) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("barrages.Barrages.Do error: %+v", err)
		}
	}()

	if err = b.find(); err != nil {
		b.errorCode = codes.ErrorCodeBarrageFind
		return err
	}
	return
}

func (b *Barrages) find() (err error) {
	if err := b.barrageFinder.Do(); err != nil {
		return err
	}
	return
}

// Barrages barrages
func (b *Barrages) Barrages() []*Barrage {
	mBarrages := b.barrageFinder.Result()

	srvBarrages := []*Barrage{}
	for i := range mBarrages {
		srvBarrage := modelBarrageTosrvBarrage(mBarrages[i])
		srvBarrages = append(srvBarrages, srvBarrage)
	}
	return srvBarrages
}

// Count total result count
func (b *Barrages) Count() int64 {
	return b.barrageFinder.Count()
}
