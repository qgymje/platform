package models

import (
	"platform/utils"

	"github.com/astaxie/beego/orm"
)

// FriendFinder friend finder
type FriendFinder struct {
	userID        string
	offset, limit int

	query   orm.QuerySeter
	friends []*Friend
}

// NewFriendFinder new friend finder
func NewFriendFinder() *FriendFinder {
	f := new(FriendFinder)
	f.friends = []*Friend{}
	f.query = GetDB().QueryTable(TableNameFriend)
	return f
}

// Limit limit
func (f *FriendFinder) Limit(offset, limit int) *FriendFinder {
	f.offset = offset
	f.limit = limit

	f.query = f.query.Offset(int64(f.offset))
	f.query = f.query.Limit(f.limit)

	return f
}

// UserID set user id
func (f *FriendFinder) UserID(userID string) *FriendFinder {
	f.userID = userID
	cond1 := orm.NewCondition().And("from_user_id", userID)
	cond2 := orm.NewCondition().And("to_user_id", userID)
	f.query = f.query.SetCond(cond1.OrCond(cond2))
	return f
}

// Do the query
func (f *FriendFinder) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("models.FriendFinder.Do error: %+v", err)
		}
	}()

	n, err := f.query.All(&f.friends)
	if err != nil {
		return
	}
	if n == 0 {
		return ErrNotFound
	}

	return nil
}

// Result result
func (f *FriendFinder) Result() []*Friend {
	return f.friends
}

// One one
func (f *FriendFinder) One() *Friend {
	return f.friends[0]
}

// Count count
func (f *FriendFinder) Count() int64 {
	n, _ := f.query.Limit(-1).Count()
	return n
}
