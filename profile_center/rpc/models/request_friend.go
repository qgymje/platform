package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// RequestFriend request friend
type RequestFriend struct {
	ID         int64  `orm:"column(id)"`
	FromUserID string `orm:"column(from_user_id)"`
	ToUserID   string `orm:"column(to_user_id)"`
	Message    string
	Status     int8 // 0 unprocessed 1 success 2 deny
	CreatedAt  time.Time
}

// TableName table name
func (RequestFriend) TableName() string {
	return TableNameRequestFriend
}

// GetID get requset id
func (r *RequestFriend) GetID() string {
	return strconv.FormatInt(r.ID, 10)
}

// Create a request
func (r *RequestFriend) Create() (err error) {
	r.CreatedAt = time.Now()
	_, err = GetDB().Insert(r)
	return
}

// Find find a record
func (r *RequestFriend) Find() (err error) {
	return GetDB().QueryTable(TableNameRequestFriend).Filter("from_user_id", r.FromUserID).Filter("to_user_id", r.ToUserID).One(r)
}

// FindByID find by request id
func (r *RequestFriend) FindByID() (err error) {
	return GetDB().Read(r)
}

func (r *RequestFriend) update(params orm.Params) (err error) {
	_, err = GetDB().QueryTable(TableNameRequestFriend).Filter("id", r.ID).Update(params)
	return
}

// Agree agree
func (r *RequestFriend) Agree() (err error) {
	params := orm.Params{"status": int(Agreed)}
	return r.update(params)
}

// Refuse refuse
func (r *RequestFriend) Refuse() (err error) {
	params := orm.Params{"status": int(Agreed)}
	return r.update(params)
}
