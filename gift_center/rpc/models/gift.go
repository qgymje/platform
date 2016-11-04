package models

import "time"

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

// TableUnique unique
func (g *Gift) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

// Create create a gift record
func (g *Gift) Create() (err error) {
	result, err := GetDB().Raw(`replace into gifts("name", "image", "snow_flake", "snow_ball", "updated_at") values(?,?,?,?,?)`).Exec()
	if err != nil {
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return
}
