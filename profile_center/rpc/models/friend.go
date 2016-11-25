package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Friend friend relation
type Friend struct {
	ID            int64          `orm:"column(id)"`
	RequestFriend *RequestFriend `orm:"column(request_id);rel(fk)"`
	FromUserID    string         `orm:"column(from_user_id)"`
	ToUserID      string         `orm:"column(to_user_id)"`
	CreatedAt     time.Time
}

// TableName table name
func (Friend) TableName() string {
	return TableNameFriend
}

// Create a friend record
func (f *Friend) Create() (err error) {
	GetDB().Begin()
	if err = f.RequestFriend.Agree(); err != nil {
		GetDB().Rollback()
		return
	}

	f.CreatedAt = time.Now()
	if _, err = GetDB().Insert(f); err != nil {
		GetDB().Rollback()
		return
	}

	GetDB().Commit()
	return
}

// Find find
func (f *Friend) Find() (err error) {
	cond1 := orm.NewCondition().And("from_user_id", f.FromUserID).And("to_user_id", f.ToUserID)
	cond2 := orm.NewCondition().And("from_user_id", f.ToUserID).And("to_user_id", f.FromUserID)
	cond := orm.NewCondition().AndCond(cond1).OrCond(cond2)
	err = GetDB().QueryTable(TableNameFriend).SetCond(cond).One(f)
	return
}
