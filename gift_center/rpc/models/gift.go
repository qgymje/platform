package models

import (
	"strconv"
	"time"
)

// Gift gift model
type Gift struct {
	ID        int64 `orm:"column(id)"`
	Name      string
	Image     string
	SnowFlake uint // not changeable
	SnowBall  uint // not changeable
	Combo     int8
	UpdatedAt time.Time
	DeletedAt time.Time
}

// TableName tablename
func (Gift) TableName() string {
	return TableNameGift
}

// TableUnique unique
func (g *Gift) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

// GetID wrapper id
func (g *Gift) GetID() string {
	return strconv.FormatInt(g.ID, 10)
}

// Create create a gift record
func (g *Gift) Create() (err error) {
	result, err := GetDB().Raw(`replace into gifts("name", "image", "snow_flake", "snow_ball", "updated_at") values(?,?,?,?,?)`, g.Name, g.Image, g.SnowFlake, g.SnowBall, g.UpdatedAt).Exec()
	if err != nil {
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return
}

// FindGiftByID find gift by id
func FindGiftByID(id string) (*Gift, error) {
	iid, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return nil, err
	}
	var g Gift
	err = GetDB().QueryTable(TableNameGift).Filter("id", iid).Filter("deleted_at__isnull", true).One(&g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

// FindGifts find gifts
func FindGifts() (gs []*Gift, err error) {
	_, err = GetDB().QueryTable(TableNameGift).Filter("deleted_at__isnull", true).All(&gs)
	if err != nil {
		return nil, err
	}
	return
}
