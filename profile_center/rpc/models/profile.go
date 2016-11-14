package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	// ErrNotEnoughSnowBall not enough snow ball
	ErrNotEnoughSnowBall = errors.New("not enough snowball")
	// ErrNotEnoughSnowFlake not enough snow flake
	ErrNotEnoughSnowFlake = errors.New("not enough snowflake")
)

// Profile user profile
type Profile struct {
	ID        int64  `orm:"column(id)"`
	UserID    string `orm:"column(user_id)"`
	SnowBall  uint
	SnowFlake uint
	CreatedAt time.Time
}

// TableName table name
func (Profile) TableName() string {
	return TableNameProfile
}

// Create create a profile
func (p *Profile) Create() (err error) {
	p.SnowBall = 100
	p.SnowFlake = 1000
	p.CreatedAt = time.Now()

	_, err = GetDB().Insert(p)
	return
}

// Find find the profile
func (p *Profile) Find() (err error) {
	return GetDB().QueryTable(TableNameProfile).Filter("user_id", p.UserID).One(p)
}

// Withdraw withdraw with message
func (p *Profile) Withdraw(snowBall, snowFlake uint, typeID uint, targetID string) (msgID int64, err error) {
	if p.SnowFlake < snowFlake {
		return 0, ErrNotEnoughSnowFlake
	}

	if p.SnowBall < snowBall {
		return 0, ErrNotEnoughSnowBall
	}

	if err = GetDB().Begin(); err != nil {
		return 0, err
	}

	updatedSnowBall := p.SnowBall - snowBall
	updatedSnowFlake := p.SnowFlake - snowFlake
	params := orm.Params{"snow_ball": updatedSnowBall, "snow_flake": updatedSnowFlake}
	if _, err = GetDB().QueryTable(TableNameProfile).Filter("user_id", p.UserID).Update(params); err != nil {
		GetDB().Rollback()
		return 0, err
	}

	msgModel := &Message{
		UserID:    p.UserID,
		TypeID:    typeID,
		TargetID:  targetID,
		SnowBall:  snowBall,
		SnowFlake: snowFlake,
		Status:    int(Created),
		CreatedAt: time.Now(),
	}
	if err = msgModel.Create(); err != nil {
		GetDB().Rollback()
		return 0, err
	}

	if err = GetDB().Commit(); err != nil {
		return 0, err
	}

	return msgModel.ID, nil
}

// WithdrawRollback rollback
func (p *Profile) WithdrawRollback(msgID string) (err error) {
	imsgID, err := strconv.ParseInt(msgID, 10, 0)
	if err != nil {
		return
	}
	msgModel := &Message{ID: imsgID}
	if err = msgModel.Find(); err != nil {
		return err
	}

	rollbackedSnowBall := p.SnowBall + msgModel.SnowBall
	rollbackedSnowFlake := p.SnowBall + msgModel.SnowBall

	if err = GetDB().Begin(); err != nil {
		return
	}

	params := orm.Params{"snow_ball": rollbackedSnowBall, "snow_flake": rollbackedSnowFlake}
	if _, err = GetDB().QueryTable(TableNameProfile).Filter("user_id", p.UserID).Update(params); err != nil {
		GetDB().Rollback()
		return
	}

	if err = msgModel.Rollback(); err != nil {
		GetDB().Rollback()
		return
	}

	if err = GetDB().Commit(); err != nil {
		return
	}
	return
}

// WithdrawCommit commit
func (p *Profile) WithdrawCommit(msgID string) (err error) {
	imsgID, err := strconv.ParseInt(msgID, 10, 0)
	if err != nil {
		return
	}
	msgModel := &Message{ID: imsgID}
	if err = msgModel.Find(); err != nil {
		return err
	}
	if err = msgModel.Commit(); err != nil {
		return
	}
	return
}
